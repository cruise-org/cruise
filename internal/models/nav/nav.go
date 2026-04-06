// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package nav

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cruise-org/cruise/internal/messages"
	"github.com/cruise-org/cruise/pkg/enums"
	"github.com/cruise-org/cruise/pkg/keymap"
	"github.com/cruise-org/cruise/pkg/page"
	"github.com/cruise-org/cruise/pkg/styles"
)

type Nav struct {
	Width       int
	Height      int
	Pages       map[string][]enums.PageType
	Keymap      keymap.NavMap
	KeymapMap   map[enums.PageType]key.Binding
	PageNameMap map[enums.PageType]string
}

func NewNav(w, h int) *Nav {
	pgs := map[string][]enums.PageType{
		"System":    {enums.Home, enums.Registry},
		"Artifacts": {enums.Containers, enums.Images, enums.Networks, enums.Volumes},
		"Ops":       {enums.Vulnerability, enums.Monitoring},
	}

	km := keymap.NewNavMap()

	kmp := map[enums.PageType]key.Binding{
		enums.Home:          km.Dashboard,
		enums.Containers:    km.Containers,
		enums.Images:        km.Images,
		enums.Networks:      km.Networks,
		enums.Volumes:       km.Volumes,
		enums.Monitoring:    km.Monitoring,
		enums.Vulnerability: km.Vulnerability,
		enums.Registry:      km.Registry,
	}

	pgNameMap := map[enums.PageType]string{
		enums.Home:          "Dashboard",
		enums.Containers:    "Containers",
		enums.Images:        "Images",
		enums.Networks:      "Networks",
		enums.Volumes:       "Volumes",
		enums.Monitoring:    "Monitoring",
		enums.Vulnerability: "Vulnerability",
		enums.Registry:      "Registry",
	}

	return &Nav{
		Width:       w,
		Height:      h,
		Pages:       pgs,
		Keymap:      km,
		KeymapMap:   kmp,
		PageNameMap: pgNameMap,
	}
}

func (s *Nav) Init() tea.Cmd { return nil }

func (s *Nav) Update(msg tea.Msg) (page.Page, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.Keymap.Dashboard):
			return s, func() tea.Msg { return messages.ChangePg{Pg: enums.Home, Exited: false} }
		case key.Matches(msg, s.Keymap.Containers):
			return s, func() tea.Msg { return messages.ChangePg{Pg: enums.Containers, Exited: false} }
		case key.Matches(msg, s.Keymap.Images):
			return s, func() tea.Msg { return messages.ChangePg{Pg: enums.Images, Exited: false} }
		case key.Matches(msg, s.Keymap.Networks):
			return s, func() tea.Msg { return messages.ChangePg{Pg: enums.Networks, Exited: false} }
		case key.Matches(msg, s.Keymap.Volumes):
			return s, func() tea.Msg { return messages.ChangePg{Pg: enums.Volumes, Exited: false} }
		case key.Matches(msg, s.Keymap.Monitoring):
			return s, func() tea.Msg { return messages.ChangePg{Pg: enums.Monitoring, Exited: false} }
		case key.Matches(msg, s.Keymap.Vulnerability):
			return s, func() tea.Msg { return messages.ChangePg{Pg: enums.Vulnerability, Exited: false} }
		case key.Matches(msg, s.Keymap.Registry):
			return s, func() tea.Msg { return messages.ChangePg{Pg: enums.Registry, Exited: false} }
		}
	}
	return s, nil
}

func (s *Nav) View() string {
	h := s.Height - strings.Count(styles.NavText, "\n")
	return lipgloss.JoinVertical(lipgloss.Center,
		styles.TextStyle().Render(styles.NavText),
		s.GetPages(s.Width, h, "System"),
		s.GetPages(s.Width, h, "Artifacts"),
		s.GetPages(s.Width, h, "Ops"),
	)
}

func (s *Nav) GetPages(w, h int, category string) string {
	pgs := s.Pages[category]

	title := lipgloss.NewStyle().PaddingLeft(3).Render(lipgloss.PlaceHorizontal(w-3, lipgloss.Left,
		styles.TitleStyle().Render(fmt.Sprintf(" %s ", category))))

	keybinds := ""
	keybindW := (w - 10) / 4

	for k, v := range pgs {
		keybind := lipgloss.PlaceHorizontal(keybindW, lipgloss.Left, lipgloss.JoinHorizontal(lipgloss.Center,
			styles.DetailKeyStyle().Render(fmt.Sprintf(" %s ", s.KeymapMap[v].Keys()[0])),
			"  ",
			s.PageNameMap[v],
		))

		if k%4 == 0 && k != 0 {
			keybinds += "\n"
		}

		keybinds += keybind
	}

	return lipgloss.Place(w, h/5, lipgloss.Left, lipgloss.Center, lipgloss.JoinVertical(lipgloss.Left, title, "\n",
		lipgloss.NewStyle().PaddingLeft(10).Render(keybinds)))
}

func (s *Nav) Cleanup() {}
