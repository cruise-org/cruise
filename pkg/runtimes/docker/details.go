// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package dockerruntime

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type ContainerDetails struct {
	ID   string
	data *map[string]string
	err  error
}

func NewContainerDetails(ctx context.Context, id string, cli *client.Client) *ContainerDetails {
	s := ContainerDetails{ID: id}

	insp, err := cli.ContainerInspect(ctx, s.ID)
	if err != nil {
		s.err = err
		return &s
	}

	uptime := "-"
	if startedAt, err := time.Parse(time.RFC3339Nano, insp.State.StartedAt); err == nil && !startedAt.IsZero() {
		uptime = time.Since(startedAt).String()
	}

	s.data = &map[string]string{
		"ID":            s.ID,
		"Name":          insp.Name,
		"Entrypoint":    strings.Join(insp.Config.Entrypoint, " "),
		"Command":       strings.Join(insp.Config.Cmd, " "),
		"Image":         insp.Image,
		"Status":        string(insp.State.Status),
		"RestartPolicy": string(insp.HostConfig.RestartPolicy.Name),
		"Uptime":        uptime,
	}

	return &s
}

func (s *ContainerDetails) Title() string { return "Container Details" }

func (s *ContainerDetails) Stats(ctx context.Context) (*map[string]string, error) {
	return s.data, s.err
}

type ContainerResources struct {
	ID   string
	data *map[string]string
	err  error
}

func NewContainerResourceDetails(ctx context.Context, id string, cli *client.Client) *ContainerResources {
	s := ContainerResources{ID: id}

	res, err := cli.ContainerStats(ctx, s.ID, false)
	if err != nil {
		s.err = err
		return &s
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		s.err = err
		return &s
	}

	var stats container.StatsResponse
	err = json.Unmarshal(data, &stats)
	if err != nil {
		s.data = &map[string]string{
			"error": err.Error(),
		}
		return &s
	}

	r, w := GetBlkio(stats, res.OSType)

	s.data = &map[string]string{
		"CPU":       fmt.Sprintf("%d", stats.CPUStats.CPUUsage.TotalUsage),
		"Memory":    fmt.Sprintf("%d", stats.MemoryStats.Usage),
		"Processes": fmt.Sprintf("%d", stats.NumProcs),
		"Blkio":     fmt.Sprintf("%d Rx / %d Wx", r, w),
	}

	return &s
}

func (s *ContainerResources) Title() string { return "Resources" }

func (s *ContainerResources) Stats(ctx context.Context) (*map[string]string, error) {
	return s.data, s.err
}

type ContainerNetworks struct {
	ID   string
	data *map[string]string
	err  error
}

func NewContainerNetworksDetails(ctx context.Context, id string, cli *client.Client) *ContainerNetworks {
	s := ContainerNetworks{ID: id}

	res, err := cli.ContainerInspect(ctx, s.ID)
	if err != nil {
		s.err = err
		return &s
	}

	nets := []string{}
	for k := range res.NetworkSettings.Networks {
		nets = append(nets, k)
	}

	s.data = &map[string]string{
		"IP":       res.NetworkSettings.IPAddress,
		"Mac":      res.NetworkSettings.MacAddress,
		"Networks": strings.Join(nets, "\n"),
	}

	return &s
}

func (s *ContainerNetworks) Title() string { return "Networks" }

func (s *ContainerNetworks) Stats(ctx context.Context) (*map[string]string, error) {
	return s.data, s.err
}

type ContainerVolumes struct {
	ID   string
	data *map[string]string
	err  error
}

func NewContainerVolumesDetails(ctx context.Context, id string, cli *client.Client) *ContainerVolumes {
	s := ContainerVolumes{ID: id}

	res, err := cli.ContainerInspect(ctx, s.ID)
	if err != nil {
		s.err = err
		return &s
	}

	vols := []string{}
	for _, v := range res.Mounts {
		vols = append(vols, v.Name)
	}

	s.data = &map[string]string{
		"MountLabel": res.MountLabel,
		"Mounts":     strings.Join(vols, "\n"),
	}

	return &s
}

func (s *ContainerVolumes) Title() string { return "Volumes" }

func (s *ContainerVolumes) Stats(ctx context.Context) (*map[string]string, error) {
	return s.data, s.err
}
