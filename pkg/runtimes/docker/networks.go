// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package dockerruntime

import (
	"context"

	"github.com/cruise-org/cruise/pkg/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
)

func (s *DockerRuntime) Networks(ctx context.Context) (*[]types.Network, error) {
	dockerNet, err := s.Client.NetworkList(ctx, network.ListOptions{})
	if err != nil {
		return nil, err
	}

	// Type Assert
	net := make([]types.Network, 0, len(dockerNet))
	for _, v := range dockerNet {
		net = append(net, types.Network{
			ID:            v.ID,
			Runtime:       "docker",
			Name:          v.Name,
			Scope:         v.Scope,
			Driver:        v.Driver,
			IPv4:          v.EnableIPv4,
			NumContainers: len(v.Containers),
		})
	}

	return &net, nil
}

func (s *DockerRuntime) PruneNetworks(ctx context.Context) error {
	_, err := s.Client.NetworksPrune(ctx, filters.NewArgs())
	return err
}

func (s *DockerRuntime) RemoveNetwork(ctx context.Context, id string) error {
	return s.Client.NetworkRemove(ctx, id)
}
