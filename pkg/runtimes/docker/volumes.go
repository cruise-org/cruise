// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package dockerruntime

import (
	"context"

	"github.com/cruise-org/cruise/pkg/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
)

func (s *DockerRuntime) Volumes(ctx context.Context) (*[]types.Volume, error) {
	dockerVol, err := s.Client.VolumeList(ctx, volume.ListOptions{})
	if err != nil {
		return nil, err
	}

	// Type Assert
	vol := make([]types.Volume, 0, len(dockerVol.Volumes))
	for _, v := range dockerVol.Volumes {
		vol = append(vol, types.Volume{
			Name:       v.Name,
			Runtime:    "docker",
			Scope:      v.Scope,
			Driver:     v.Driver,
			Mountpoint: v.Mountpoint,
			CreatedAt:  v.CreatedAt,
		})
	}

	return &vol, nil
}

func (s *DockerRuntime) PruneVolumes(ctx context.Context) error {
	_, err := s.Client.VolumesPrune(ctx, filters.NewArgs())
	return err
}

func (s *DockerRuntime) RemoveVolume(ctx context.Context, id string) error {
	return s.Client.VolumeRemove(ctx, id, false)
}
