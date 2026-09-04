// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package dockerruntime

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/cruise-org/cruise/pkg/types"
	"github.com/docker/docker/api/types/container"
)

func (s *DockerRuntime) Containers(ctx context.Context) (*[]types.Container, error) {
	dockerCtr, err := s.Client.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	// Type Assert
	cnt := make([]types.Container, 0, len(dockerCtr))
	for _, v := range dockerCtr {
		state := types.ContainerState(v.State)

		name := ""
		if len(v.Names) > 0 {
			name = v.Names[0]
		}

		// asserting ports
		ports := make([]types.ContainerPort, 0, len(v.Ports))
		for _, p := range v.Ports {
			ports = append(ports, types.ContainerPort{
				ContainerPort: p.PrivatePort,
				HostPort:      p.PublicPort,
				Protocol:      p.Type,
			})
		}

		// asserting mounts
		mounts := make([]types.ContainerMount, 0, len(v.Ports))
		for _, m := range v.Mounts {
			mounts = append(mounts, types.ContainerMount{
				Source:      m.Source,
				Destination: m.Destination,
				ReadOnly:    !m.RW,
			})
		}

		cnt = append(cnt, types.Container{
			Name:    name,
			Runtime: "docker",
			ID:      v.ID,
			Image:   v.Image,
			Created: v.Created,
			Ports:   ports,
			State:   state,
			Mounts:  mounts,
			Labels:  v.Labels,
		})
	}

	return &cnt, nil
}

func (s *DockerRuntime) ContainerDetails(ctx context.Context, id string) ([]types.StatCard, *types.StatMeta) {
	stats := []types.StatCard{
		NewContainerDetails(ctx, id, s.Client),
		NewContainerResourceDetails(ctx, id, s.Client),
		NewContainerVolumesDetails(ctx, id, s.Client),
		NewContainerNetworksDetails(ctx, id, s.Client),
	}

	return stats, &types.StatMeta{
		TotalRows:    4,
		TotalColumns: 1,
		SpanMap: &map[string]struct {
			Rows    int
			Columns int
			Index   int
		}{
			"Container Details": {Rows: 1, Columns: 1, Index: 0},
			"Resources":         {Rows: 1, Columns: 1, Index: 1},
			"Networks":          {Rows: 1, Columns: 1, Index: 2},
			"Volumes":           {Rows: 1, Columns: 1, Index: 3},
		},
	}
}

func (s *DockerRuntime) StartContainer(ctx context.Context, id string) error {
	err := s.Client.ContainerStart(ctx, id, container.StartOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (s *DockerRuntime) StopContainer(ctx context.Context, id string) error {
	err := s.Client.ContainerStop(ctx, id, container.StopOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (s *DockerRuntime) RemoveContainer(ctx context.Context, id string) error {
	err := s.Client.ContainerRemove(ctx, id, container.RemoveOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (s *DockerRuntime) PauseContainer(ctx context.Context, id string) error {
	err := s.Client.ContainerPause(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *DockerRuntime) UnpauseContainer(ctx context.Context, id string) error {
	err := s.Client.ContainerUnpause(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *DockerRuntime) RestartContainer(ctx context.Context, id string) error {
	err := s.Client.ContainerRestart(ctx, id, container.StopOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (s *DockerRuntime) ExecContainer(ctx context.Context, id string) *exec.Cmd {
	return exec.Command("docker", "exec", "-it", id, "sh")
}

func (s *DockerRuntime) PortsMap(ctx context.Context, id string) ([]string, error) {
	conts, err := s.Containers(ctx)
	if err != nil {
		return nil, err
	}

	res := map[string][]string{}
	for _, cnt := range *conts {
		inp, err := s.Client.ContainerInspect(ctx, cnt.ID)
		if err != nil {
			return nil, err
		}

		for port, bindings := range inp.NetworkSettings.Ports {
			arr := []string{}
			for _, b := range bindings {
				arr = append(arr, fmt.Sprintf("%s:%s", b.HostIP, b.HostPort))
			}

			res[port.Port()] = arr
		}
	}

	ports := make([]string, 0)
	for k, v := range res {
		ports = append(ports, fmt.Sprintf("%s -> %s", k,
			strings.Join(v, ","),
		))
	}

	return ports, nil
}
