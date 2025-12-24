package projects

import (
	"log"
	"sort"
	"time"

	"github.com/NucleoFusion/cruise/internal/colors"
	"github.com/NucleoFusion/cruise/internal/compose"
	"github.com/NucleoFusion/cruise/internal/config"
	"github.com/NucleoFusion/cruise/internal/messages"
	"github.com/NucleoFusion/cruise/internal/styles"
	internaltypes "github.com/NucleoFusion/cruise/internal/types"
	"github.com/NucleoFusion/cruise/internal/utils"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type ProjectList struct {
	Width         int
	Height        int
	Items         []*internaltypes.Project
	Err           error
	FilteredItems []*internaltypes.Project
	SelectedIndex int
	Ti            textinput.Model
	Vp            viewport.Model
	IsLoading     bool
}

func NewProjectList(w int, h int) *ProjectList {
	ti := textinput.New()
	ti.Width = w - 10
	ti.Prompt = " Search: "
	ti.Placeholder = "Press '/' to search..."

	ti.PromptStyle = lipgloss.NewStyle().Foreground(colors.Load().FocusedBorder)
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(colors.Load().PlaceholderText)
	ti.TextStyle = styles.TextStyle()

	vp := viewport.New(w+2, h-4)
	vp.Style = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().FocusedBorder).
		Padding(1).Foreground(colors.Load().Text)

	return &ProjectList{
		Width:         w,
		Height:        h,
		Ti:            ti,
		SelectedIndex: 0,
		Vp:            vp,
		IsLoading:     true,
	}
}

func (s *ProjectList) Init() tea.Cmd {
	return tea.Tick(0, func(_ time.Time) tea.Msg {
		log.Println("Calc ProjectsReadyMsg")
		items, err := compose.GetProjects()
		if err != nil {
			log.Println("Showing Error", err.Error())
			return utils.ReturnError("Projects List", "Error Getting Project Summaries", err)
		}

		log.Println("Error Sending ProjectsReadyMsg")
		return messages.ProjectsReadyMsg{
			Projects: items,
		}
	})
}

func (s *ProjectList) Update(msg tea.Msg) (*ProjectList, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.ProjectsReadyMsg:
		s.Items = msg.Projects
		s.FilteredItems = msg.Projects
		s.IsLoading = false
		return s, nil

	case tea.KeyMsg:
		if s.Ti.Focused() {
			if msg.String() == config.Cfg.Keybinds.Global.UnfocusSearch {
				s.Ti.Blur()
				return s, nil
			}
			var cmd tea.Cmd
			s.Ti, cmd = s.Ti.Update(msg)
			s.Filter(s.Ti.Value())
			s.UpdateList()
			return s, cmd
		}
		switch msg.String() {
		case config.Cfg.Keybinds.Global.FocusSearch:
			s.Ti.Focus()
			return s, nil
		case config.Cfg.Keybinds.Global.ListDown:
			if len(s.FilteredItems)-1 > s.SelectedIndex {
				s.SelectedIndex += 1
			}
			if s.SelectedIndex > s.Vp.Height+s.Vp.YOffset-3 { // -2 for border and sosething else, idk breaks otherwise
				s.Vp.YOffset += 1
			}
			s.UpdateList()
			return s, nil
		case config.Cfg.Keybinds.Global.ListUp:
			if 0 < s.SelectedIndex {
				s.SelectedIndex -= 1
			}
			if s.SelectedIndex < s.Vp.YOffset {
				s.Vp.YOffset -= 1
			}
			s.UpdateList()
			return s, nil
		case "enter": // TODO: use keymap
			return s, func() tea.Msg {
				return messages.ShowProjectDetails{
					Project: s.GetCurrentItem(),
				}
			}
		}
	}
	return s, nil
}

func (s *ProjectList) View() string {
	if s.IsLoading {
		return styles.PageStyle().Render(lipgloss.Place(s.Width, s.Height-2, lipgloss.Center, lipgloss.Center, "Loading..."))
	}

	if s.Err != nil {
		return styles.PageStyle().Render(lipgloss.Place(s.Width, s.Height-2, lipgloss.Center, lipgloss.Center, "Error: "+s.Err.Error()))
	}

	if len(s.Items) == 0 {
		return styles.PageStyle().Render(lipgloss.Place(s.Width, s.Height-2, lipgloss.Center, lipgloss.Center, "No Containers Found!"))
	}

	log.Println("Showing List")

	style := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().FocusedBorder)

	s.UpdateList()

	log.Println("Updated List")

	return lipgloss.JoinVertical(lipgloss.Center,
		style.Render(s.Ti.View()),
		s.Vp.View())
}

func (s *ProjectList) UpdateList() {
	text := lipgloss.NewStyle().Bold(true).Render(compose.ProjectHeaders(s.Width)+"\n") + "\n"
	//
	for k, v := range s.FilteredItems {
		line := compose.ProjectFormatted(v, s.Width)

		if k == s.SelectedIndex {
			line = styles.SelectedStyle().Render(line)
		} else {
			line = styles.TextStyle().Render(line)
		}

		text += line + "\n"
	}
	//
	s.Vp.SetContent(text)
}

func (s *ProjectList) Filter(val string) {
	w := (s.Width)/9 - 1

	formatted := make([]string, len(s.Items))
	originals := make([]*internaltypes.Project, len(s.Items))

	for i, v := range s.Items {
		str := compose.ProjectFormatted(v, w)
		formatted[i] = str
		originals[i] = v
	}

	ranked := fuzzy.RankFindFold(val, formatted)
	sort.Sort(ranked)

	result := make([]*internaltypes.Project, len(ranked))
	for i, r := range ranked {
		result[i] = originals[r.OriginalIndex]
	}

	s.FilteredItems = result

	if len(s.FilteredItems) <= s.SelectedIndex {
		s.SelectedIndex = len(s.FilteredItems) - 1
	}
}

func (s *ProjectList) GetCurrentItem() *internaltypes.Project {
	return s.FilteredItems[s.SelectedIndex]
}
