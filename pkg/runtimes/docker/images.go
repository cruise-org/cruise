// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package dockerruntime

import (
	"context"
	"fmt"

	"github.com/cruise-org/cruise/pkg/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
)

func (s *DockerRuntime) Images(ctx context.Context) (*[]types.Image, error) {
	dockerImg, err := s.Client.ImageList(ctx, image.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	// Type Assert
	img := make([]types.Image, 0, len(dockerImg))
	for _, v := range dockerImg {
		img = append(img, types.Image{
			ID:            v.ID,
			Runtime:       "docker",
			Tags:          v.RepoTags,
			Size:          v.Size,
			CreatedAt:     v.Created,
			NumContainers: v.Containers,
		})
	}

	return &img, nil
}

func (s *DockerRuntime) RemoveImage(ctx context.Context, id string) error {
	_, err := s.Client.ImageRemove(ctx, id, image.RemoveOptions{PruneChildren: true, Force: false})
	return err
}

func (s *DockerRuntime) PruneImages(ctx context.Context) error {
	_, err := s.Client.ImagesPrune(ctx, filters.NewArgs())
	return err
}

func (s *DockerRuntime) PushImage(ctx context.Context, id string) error {
	_, err := s.Client.ImagePush(ctx, id, image.PushOptions{})
	return err
}

func (s *DockerRuntime) PullImage(ctx context.Context, id string) error {
	_, err := s.Client.ImagePull(ctx, id, image.PullOptions{})
	return err
}

func (s *DockerRuntime) ImageLayers(ctx context.Context, id string) (string, error) {
	layers, err := s.Client.ImageHistory(ctx, id)
	if err != nil {
		return "", err
	}

	text := ""
	for k, v := range layers {
		id := v.ID
		if len(id) > 19 {
			id = id[:19]
		}
		text += fmt.Sprintf("Layer: %-3d %-20s Size: %-7s %s \n", k, id, fmt.Sprintf("%dMB", v.Size/(1024*1024)), v.CreatedBy)
	}

	return text, nil
}
