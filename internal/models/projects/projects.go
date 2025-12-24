package projects

import (
	"log"
	"strings"

	"github.com/NucleoFusion/cruise/internal/keymap"
	"github.com/NucleoFusion/cruise/internal/messages"
	styledhelp "github.com/NucleoFusion/cruise/internal/models/help"
	"github.com/NucleoFusion/cruise/internal/styles"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Projects struct {
	Width       int
	Height      int
	List        *ProjectList
	Keymap      keymap.ContainersMap
	Help        styledhelp.StyledHelp
	DetailsPg   *ProjectDetails
	ShowDetails bool
}

func NewProjects(w int, h int) *Projects {
	return &Projects{
		Width:       w,
		Height:      h,
		List:        NewProjectList(w-4, h-3-strings.Count(styles.ProjectsText, "\n")),
		ShowDetails: false,
		Help:        styledhelp.NewStyledHelp(keymap.NewContainersMap().Bindings(), w),
	}
}

func (s *Projects) Init() tea.Cmd {
	return s.List.Init()
}

func (s *Projects) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.ProjectsReadyMsg:
		var cmd tea.Cmd
		s.List, cmd = s.List.Update(msg)
		return s, cmd
	case messages.ProjectsDetailsUpdate:
		if s.ShowDetails {
			var cmd tea.Cmd
			s.DetailsPg, cmd = s.DetailsPg.Update(msg)
			return s, cmd
		}
	case messages.ShowProjectDetails:
		s.ShowDetails = true
		s.List.Ti.Blur()
		s.DetailsPg = NewProjectDetails(s.Width, s.Height, msg.Project)
		return s, s.DetailsPg.Init()
	case messages.CloseProjectDetails:
		s.ShowDetails = false
		s.List.Ti.Blur()
		return s, nil
	case tea.KeyMsg:
		if s.ShowDetails {
			if msg.String() == "esc" { // TODO: Use keymap
				s.ShowDetails = false
				s.List.Ti.Focus()
				return s, nil
			}
			var cmd tea.Cmd
			s.DetailsPg, cmd = s.DetailsPg.Update(msg)
			return s, cmd
		} else if s.List.Ti.Focused() {
			var cmd tea.Cmd
			s.List, cmd = s.List.Update(msg)
			return s, cmd
		}
	}

	var cmd tea.Cmd
	s.List, cmd = s.List.Update(msg)
	return s, cmd
}

func (s *Projects) View() string {
	log.Println("Showing Projects")
	if s.ShowDetails {
		return s.DetailsPg.View()
	}

	return lipgloss.JoinVertical(lipgloss.Center,
		styles.TextStyle().Render(styles.ProjectsText), s.GetListText(), s.Help.View())
}

func (s *Projects) GetListText() string {
	return lipgloss.NewStyle().PaddingLeft(1).Render(s.List.View())
}
