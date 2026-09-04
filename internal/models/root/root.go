// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package root

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cruise-org/cruise/internal/messages"
	"github.com/cruise-org/cruise/internal/models/containers"
	errorpopup "github.com/cruise-org/cruise/internal/models/error"
	"github.com/cruise-org/cruise/internal/models/home"
	"github.com/cruise-org/cruise/internal/models/images"
	"github.com/cruise-org/cruise/internal/models/monitoring"
	msgpopup "github.com/cruise-org/cruise/internal/models/msg"
	"github.com/cruise-org/cruise/internal/models/nav"
	"github.com/cruise-org/cruise/internal/models/networks"
	"github.com/cruise-org/cruise/internal/models/volumes"
	"github.com/cruise-org/cruise/internal/models/vulnerability"
	"github.com/cruise-org/cruise/pkg/enums"
	"github.com/cruise-org/cruise/pkg/page"
	overlay "github.com/rmhubbert/bubbletea-overlay"
)

type Root struct {
	Width       int
	Height      int
	CurrentPage enums.PageType
	PageItems   map[string]enums.PageType
	// Showing Variables
	IsLoading      bool
	IsShowingError bool
	IsShowingMsg   bool
	// PageMap
	PageMap map[enums.PageType]page.Page
	// Models
	ErrorPopup *errorpopup.ErrorPopup
	MsgPopup   *msgpopup.MsgPopup
	Overlay    *overlay.Model
}

func NewRoot() *Root {
	return &Root{
		CurrentPage:    enums.Home,
		IsLoading:      true,
		IsShowingError: false,
		PageItems: map[string]enums.PageType{
			"Home":          enums.Home,
			"Containers":    enums.Containers,
			"Images":        enums.Images,
			"Vulnerability": enums.Vulnerability,
			"Monitoring":    enums.Monitoring,
			"Networks":      enums.Networks,
			"Volumes":       enums.Volumes,
		},
		PageMap: map[enums.PageType]page.Page{},
	}
}

func (s *Root) Init() tea.Cmd { return nil }

func (s *Root) newPage(pt enums.PageType) page.Page {
	switch pt {
	case enums.Nav:
		return nav.NewNav(s.Width, s.Height)
	case enums.Home:
		return home.NewHome(s.Width, s.Height)
	case enums.Containers:
		return containers.NewContainers(s.Width, s.Height)
	case enums.Images:
		return images.NewImages(s.Width, s.Height)
	case enums.Vulnerability:
		return vulnerability.NewVulnerability(s.Width, s.Height)
	case enums.Monitoring:
		return monitoring.NewMonitoring(s.Width, s.Height)
	case enums.Networks:
		return networks.NewNetworks(s.Width, s.Height)
	case enums.Volumes:
		return volumes.NewVolumes(s.Width, s.Height)
	default:
		return home.NewHome(s.Width, s.Height)
	}
}

// page returns the live instance for pt, creating it on first use.
func (s *Root) page(pt enums.PageType) page.Page {
	p, ok := s.PageMap[pt]
	if !ok {
		p = s.newPage(pt)
		s.PageMap[pt] = p
	}
	return p
}

// resetPages tears down every live page so they get rebuilt fresh.
func (s *Root) resetPages() {
	for pt, p := range s.PageMap {
		p.Cleanup()
		delete(s.PageMap, pt)
	}
}

func (s *Root) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.CloseError:
		s.IsShowingError = false
		s.Overlay = nil
		return s, nil
	case messages.ErrorMsg:
		s.IsShowingError = true
		s.ErrorPopup = errorpopup.NewErrorPopup(s.Width, s.Height, msg.Msg, msg.Title, msg.Locn)

		s.Overlay = overlay.New(s.ErrorPopup, s.page(s.CurrentPage), overlay.Right, overlay.Top, 2, 2)
		return s, tea.Tick(3*time.Second, func(_ time.Time) tea.Msg { return messages.CloseError{} })
	case messages.CloseMsgPopup:
		s.IsShowingMsg = false
		s.Overlay = nil
		return s, nil
	case messages.MsgPopup:
		s.IsShowingMsg = true
		s.MsgPopup = msgpopup.NewMsgPopup(s.Width, s.Height, msg.Msg, msg.Title, msg.Locn)

		s.Overlay = overlay.New(s.MsgPopup, s.page(s.CurrentPage), overlay.Right, overlay.Top, 2, 2)
		return s, tea.Tick(3*time.Second, func(_ time.Time) tea.Msg { return messages.CloseMsgPopup{} })
	case messages.ChangePg:
		if msg.Pg != s.CurrentPage {
			if old, ok := s.PageMap[s.CurrentPage]; ok {
				old.Cleanup()
				delete(s.PageMap, s.CurrentPage)
			}
			s.CurrentPage = msg.Pg
		}
		return s, s.page(s.CurrentPage).Init()
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return s, tea.Quit
		case tea.KeyTab:
			return s, func() tea.Msg { return messages.ChangePg{Pg: enums.Nav} }
		}
	case tea.WindowSizeMsg:
		resized := s.Width != msg.Width || s.Height != msg.Height
		s.Width = msg.Width
		s.Height = msg.Height

		if resized {
			s.resetPages()
		}

		s.IsLoading = false
		return s, s.page(s.CurrentPage).Init()
	}

	updated, cmd := s.page(s.CurrentPage).Update(msg)
	s.PageMap[s.CurrentPage] = updated

	return s, cmd
}

func (s *Root) View() string {
	if s.IsLoading {
		return "\nLoading..."
	}

	if s.IsShowingError || s.IsShowingMsg {
		return s.Overlay.View()
	}

	return s.page(s.CurrentPage).View()
}
