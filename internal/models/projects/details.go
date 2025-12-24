package projects

import (
	"fmt"
	"time"

	"github.com/NucleoFusion/cruise/internal/compose"
	"github.com/NucleoFusion/cruise/internal/messages"
	"github.com/NucleoFusion/cruise/internal/styles"
	internaltypes "github.com/NucleoFusion/cruise/internal/types"
	"github.com/NucleoFusion/cruise/internal/utils"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ProjectDetails struct {
	Width     int
	Height    int
	Project   *internaltypes.Project
	Services  []string
	IsLoading bool
	Current   int
	// Viewports
	DetailsVp  viewport.Model
	ServicesVp viewport.Model
	ResourceVp viewport.Model
}

func NewProjectDetails(w, h int, s *internaltypes.Project) *ProjectDetails {
	// Details VP
	dvp := viewport.New(w/2, h/2)
	dvp.Style = styles.PageStyle().Padding(1, 2)
	// dvp.SetContent(getLabelView(v, labels, w))

	// Services VP
	svp := viewport.New(w, h/2)
	svp.Style = styles.PageStyle().Padding(1, 2)
	// dvp.SetContent(getLabelView(v, labels, w))

	// Resources VP
	rvp := viewport.New(w/2, h/2)
	rvp.Style = styles.PageStyle().Padding(1, 2)
	// dvp.SetContent(getLabelView(v, labels, w))

	srvc := make([]string, 0, len(s.Inspect.Services))
	for k := range s.Inspect.Services {
		srvc = append(srvc, k)
	}

	return &ProjectDetails{
		Width:      w,
		Height:     h,
		Project:    s,
		Current:    0,
		Services:   srvc,
		ResourceVp: rvp,
		ServicesVp: svp,
		DetailsVp:  dvp,
		IsLoading:  true,
	}
}

func (s *ProjectDetails) Init() tea.Cmd {
	return tea.Tick(0, func(_ time.Time) tea.Msg {
		return messages.ProjectsDetailsUpdate{}
	})
}

func (s *ProjectDetails) Update(msg tea.Msg) (*ProjectDetails, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.ProjectsDetailsUpdate:
		s.ResourceVp.SetContent(s.GetResourcesView())
		s.DetailsVp.SetContent(s.GetDetailsView())
		s.ServicesVp.SetContent(s.GetServicesView())
		return s, nil
	case tea.KeyMsg:
		// TODO: Use Keymap
		switch msg.String() {
		case "up":
			if s.Current > 0 {
				s.Current--
				s.ServicesVp.SetContent(s.GetServicesView())
			}
			return s, nil
		case "down":
			if s.Current < len(s.Services)-1 {
				s.Current++
				s.ServicesVp.SetContent(s.GetServicesView())
			}
			return s, nil
		}
	}
	return s, nil
}

func (s *ProjectDetails) View() string {
	return lipgloss.JoinVertical(lipgloss.Center,
		lipgloss.JoinHorizontal(lipgloss.Center, s.DetailsVp.View(), s.ResourceVp.View()),
		s.ServicesVp.View(),
	)
}

func (s *ProjectDetails) GetDetailsView() string {
	w := s.Width
	var start string

	t, err := s.Project.LatestStartedAt()
	if err != nil {
		start = "N/A"
	} else {
		start = t.Format("15:04 _2 Jan")
	}

	text := fmt.Sprintf("%s %s \n\n%s %s \n\n%s %s \n\n%s %s \n\n%s %s \n\n%s %s \n\n%s %s",
		styles.DetailKeyStyle().Render(" Name: "), styles.TextStyle().Render(s.Project.Inspect.Name),
		styles.DetailKeyStyle().Render(" Status: "), styles.TextStyle().Render(utils.Shorten(s.Project.Status(), w/3-15)),
		styles.DetailKeyStyle().Render(" Last Updated: "), styles.TextStyle().Render(utils.Shorten(start, w/3-15)),
		styles.DetailKeyStyle().Render(" Containers: "), styles.TextStyle().Render(utils.Shorten(fmt.Sprintf("%d", s.Project.NumContainers()), w/3-15)),
		styles.DetailKeyStyle().Render(" Services: "), styles.TextStyle().Render(utils.Shorten(fmt.Sprintf("%d", len(s.Services)), w/3-15)),
		styles.DetailKeyStyle().Render(" Networks: "), styles.TextStyle().Render(utils.Shorten(fmt.Sprintf("%d", len(s.Project.Inspect.Networks)), w/3-15)),
		styles.DetailKeyStyle().Render(" Volumes: "), styles.TextStyle().Render(utils.Shorten(fmt.Sprintf("%d", len(s.Project.Inspect.Volumes)), w/3-15)))

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w/3-4, lipgloss.Center, styles.TitleStyle().Render(" Project Details ")), "\n\n", text)
}

func (s *ProjectDetails) GetResourcesView() string {
	w := s.Width
	stats, err := s.Project.AggStats()
	if err != nil {
		return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w/3-4, lipgloss.Center, styles.TitleStyle().Render(" Project Details ")), "\n\n", "Error Geting Stats!")
	}

	text := fmt.Sprintf("%s %s \n\n%s %s \n\n%s %s \n\n%s %s \n\n%s %s",
		styles.DetailKeyStyle().Render(" CPU: "), styles.TextStyle().Render(fmt.Sprintf("%d", stats.CPU)),
		styles.DetailKeyStyle().Render(" Memory: "), styles.TextStyle().Render(fmt.Sprintf("%d", stats.Mem)),
		styles.DetailKeyStyle().Render(" Memory Limit: "), styles.TextStyle().Render(fmt.Sprintf("%d", stats.MemLimit)),
		styles.DetailKeyStyle().Render(" Net Rx: "), styles.TextStyle().Render(fmt.Sprintf("%d", stats.NetRx)),
		styles.DetailKeyStyle().Render(" Net Tx: "), styles.TextStyle().Render(fmt.Sprintf("%d", stats.NetTx)))

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w/3-4, lipgloss.Center, styles.TitleStyle().Render(" Project Details ")), "\n\n", text)
}

func (s *ProjectDetails) GetServicesView() string {
	w := s.Width

	text := compose.ServiceHeaders(w) + "\n\n"
	for k, v := range s.Services {
		if k == s.Current {
			text += styles.SelectedStyle().Render(compose.ServiceFormatted(w, s.Project, v))
		} else {
			text += styles.TextStyle().Render(compose.ServiceFormatted(w, s.Project, v))
		}

		text += "\n"
	}

	return lipgloss.JoinVertical(lipgloss.Center, lipgloss.PlaceHorizontal(w-4, lipgloss.Center, styles.TitleStyle().Render(" Service Details ")), "\n\n", text)
}
