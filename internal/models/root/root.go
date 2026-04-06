// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package root

import (
	"log"
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
	registrymodel "github.com/cruise-org/cruise/internal/models/registry"
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
			"Registry":      enums.Registry,
		},
	}
}

func (s *Root) Init() tea.Cmd { return nil }

func (s *Root) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("[ROOT MSG] %T\n", msg)
	switch msg := msg.(type) {
	case messages.CloseError:
		s.IsShowingError = false
		s.Overlay = nil
		return s, nil
	case messages.ErrorMsg:
		s.IsShowingError = true
		s.ErrorPopup = errorpopup.NewErrorPopup(s.Width, s.Height, msg.Msg, msg.Title, msg.Locn)

		var curr page.Page
		curr = s.PageMap[s.CurrentPage]

		s.Overlay = overlay.New(s.ErrorPopup, curr, overlay.Right, overlay.Top, 2, 2)
		return s, tea.Tick(3*time.Second, func(_ time.Time) tea.Msg { return messages.CloseError{} })
	case messages.CloseMsgPopup:
		s.IsShowingMsg = false
		s.Overlay = nil
		return s, nil
	case messages.MsgPopup:
		s.IsShowingMsg = true
		s.MsgPopup = msgpopup.NewMsgPopup(s.Width, s.Height, msg.Msg, msg.Title, msg.Locn)

		var curr page.Page
		curr = s.PageMap[s.CurrentPage]

		s.Overlay = overlay.New(s.MsgPopup, curr, overlay.Right, overlay.Top, 2, 2)
		return s, tea.Tick(3*time.Second, func(_ time.Time) tea.Msg { return messages.CloseMsgPopup{} })
	case messages.ChangePg:
		s.CurrentPage = msg.Pg
		cmd := s.PageMap[s.CurrentPage].Init()
		return s, cmd
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return s, tea.Quit
		case tea.KeyTab:
			return s, func() tea.Msg { return messages.ChangePg{Pg: enums.Nav} }
		}
	case tea.WindowSizeMsg:
		s.Width = msg.Width
		s.Height = msg.Height

		w := s.Width
		h := s.Height

		s.PageMap = map[enums.PageType]page.Page{
			enums.Nav:           nav.NewNav(w, h),
			enums.Home:          home.NewHome(w, h),
			enums.Containers:    containers.NewContainers(w, h),
			enums.Images:        images.NewImages(w, h),
			enums.Vulnerability: vulnerability.NewVulnerability(w, h),
			enums.Monitoring:    monitoring.NewMonitoring(w, h),
			enums.Networks:      networks.NewNetworks(w, h),
			enums.Volumes:       volumes.NewVolumes(w, h),
			enums.Registry:      registrymodel.NewRegistryModel(w, h),
		}

		s.IsLoading = false
		return s, s.PageMap[s.CurrentPage].Init()
	}

	var cmd tea.Cmd
	s.PageMap[s.CurrentPage], cmd = s.PageMap[s.CurrentPage].Update(msg)

	return s, cmd
}

func (s *Root) View() string {
	if s.IsLoading {
		return "\nLoading..."
	}

	if s.IsShowingError || s.IsShowingMsg {
		return s.Overlay.View()
	}

	return s.PageMap[s.CurrentPage].View()
}
