// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package home

import (
	"context"
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cruise-org/cruise/internal/messages"
	"github.com/cruise-org/cruise/pkg/runtimes"
	"github.com/cruise-org/cruise/pkg/styles"
	"github.com/cruise-org/cruise/pkg/types"
)

const statsQueryTimeout = 15 * time.Second

type QuickStats struct {
	Width      int
	Height     int
	Containers *[]types.Container
	Images     *[]types.Image
	Networks   *[]types.Network
	Volumes    *[]types.Volume
}

func NewQuickStats(w int, h int) *QuickStats {
	return &QuickStats{
		Width:  w,
		Height: h,
	}
}

func getCmds() []tea.Cmd {
	return []tea.Cmd{
		// Containers
		func() tea.Msg {
			ctx, cancel := context.WithTimeout(context.Background(), statsQueryTimeout)
			defer cancel()
			cnts, err := runtimes.RuntimeSrv.Containers(ctx)
			if err != nil {
				return messages.ErrorMsg{Msg: err.Error()}
			}
			return messages.HomeStatContainer{Containers: cnts}
		},

		// Images
		func() tea.Msg {
			ctx, cancel := context.WithTimeout(context.Background(), statsQueryTimeout)
			defer cancel()
			cnts, err := runtimes.RuntimeSrv.Images(ctx)
			if err != nil {
				return messages.ErrorMsg{Msg: err.Error()}
			}
			return messages.HomeStatImage{Images: cnts}
		},

		// Networks
		func() tea.Msg {
			ctx, cancel := context.WithTimeout(context.Background(), statsQueryTimeout)
			defer cancel()
			cnts, err := runtimes.RuntimeSrv.Networks(ctx)
			if err != nil {
				return messages.ErrorMsg{Msg: err.Error()}
			}
			return messages.HomeStatNetwork{Networks: cnts}
		},

		// Volumes
		func() tea.Msg {
			ctx, cancel := context.WithTimeout(context.Background(), statsQueryTimeout)
			defer cancel()
			cnts, err := runtimes.RuntimeSrv.Volumes(ctx)
			if err != nil {
				return messages.ErrorMsg{Msg: err.Error()}
			}
			return messages.HomeStatVolume{Volumes: cnts}
		},
	}
}

func (s *QuickStats) Init() tea.Cmd {
	return tea.Batch(getCmds()...)
}

func (s *QuickStats) Update(msg tea.Msg) (*QuickStats, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.HomeStatContainer:
		s.Containers = msg.Containers
	case messages.HomeStatImage:
		s.Images = msg.Images
	case messages.HomeStatNetwork:
		s.Networks = msg.Networks
	case messages.HomeStatVolume:
		s.Volumes = msg.Volumes
	}
	return s, nil
}

func (s QuickStats) View() string {
	return styles.SubpageStyle().PaddingTop(1).PaddingLeft(4).Render(lipgloss.JoinVertical(lipgloss.Center,
		styles.TitleStyle().Render("Docker Stats"),
		lipgloss.NewStyle().
			Width(s.Width-6).   //-6 from padding(4) and border(2)
			Height(s.Height-4). //-4 from title(1) border(2) and padding(1)
			Align(lipgloss.Left, lipgloss.Center).
			Render(s.GetFormattedView())))
}

func (s QuickStats) GetFormattedView() string {
	if s.Containers == nil && s.Images == nil && s.Volumes == nil && s.Networks == nil {
		return "Loading"
	}

	cnts := "..."
	if s.Containers != nil {
		cnts = fmt.Sprintf("%d", len(*s.Containers))
	}

	imgs := "..."
	if s.Images != nil {
		imgs = fmt.Sprintf("%d", len(*s.Images))
	}

	nets := "..."
	if s.Networks != nil {
		nets = fmt.Sprintf("%d", len(*s.Networks))
	}

	vols := "..."
	if s.Volumes != nil {
		vols = fmt.Sprintf("%d", len(*s.Volumes))
	}

	return fmt.Sprintf(`
	Containers: %s

	Images:     %s

	Volumes:    %s

	Networks:   %s
			`, cnts, imgs, vols, nets)
}
