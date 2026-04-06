package loginmodel

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cruise-org/cruise/internal/messages"
	styledhelp "github.com/cruise-org/cruise/internal/models/help"
	"github.com/cruise-org/cruise/internal/utils"
	"github.com/cruise-org/cruise/pkg/colors"
	"github.com/cruise-org/cruise/pkg/enums"
	"github.com/cruise-org/cruise/pkg/keymap"
	"github.com/cruise-org/cruise/pkg/registry"
	"github.com/cruise-org/cruise/pkg/styles"
)

type LoginModel struct {
	Width       int
	Height      int
	Registry    *registry.Registry
	PassTi      textinput.Model
	Opts        []string
	SelectedOpt int

	Keymap keymap.LoginModelKeymap
	Help   styledhelp.StyledHelp
}

func NewLoginModel(w, h int, r *registry.Registry) *LoginModel {
	ti := textinput.New()
	ti.Prompt = ""
	ti.EchoMode = textinput.EchoPassword
	ti.Width = w/3 - 4 - 4 - 11 // Taking into account border, indicators and 'Password : '
	ti.Focus()

	km := keymap.NewLoginModelMap()

	return &LoginModel{
		Width:       w / 3,
		Height:      h,
		Registry:    r,
		PassTi:      ti,
		Opts:        []string{"Login", "Return", "Ignore"},
		SelectedOpt: 0,
		Keymap:      km,
		Help:        styledhelp.NewStyledHelp(km.Bindings(), w),
	}
}

func (s *LoginModel) Init() tea.Cmd { return nil }

func (s *LoginModel) Update(msg tea.Msg) (*LoginModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.Keymap.Left):
			if s.SelectedOpt > 0 {
				s.SelectedOpt--
			}
			return s, nil
		case key.Matches(msg, s.Keymap.Right):
			if s.SelectedOpt < len(s.Opts)-1 {
				s.SelectedOpt++
			}
			return s, nil
		case key.Matches(msg, s.Keymap.Enter):
			switch s.Opts[s.SelectedOpt] {
			case "Return":
				return s, func() tea.Msg { return messages.ChangePg{Pg: enums.Home} }
			case "Ignore":
				return s, func() tea.Msg { return messages.IgnoreLoginMessage{} }
			case "Login":
				return s, func() tea.Msg {
					return messages.LoginMessage{
						Pass: s.PassTi.Value(),
					}
				}
			}
		default:
			ti, cmd := s.PassTi.Update(msg)
			s.PassTi = ti
			return s, cmd
		}
	}
	return s, nil
}

func (s *LoginModel) View() string {
	view := lipgloss.JoinVertical(lipgloss.Left,
		fmt.Sprintf("Username : %s", (*s.Registry).Username()),
		fmt.Sprintf("Provider : %s", (*s.Registry).Provider()),
		fmt.Sprintf("Domain   : %s", utils.Shorten((*s.Registry).Domain(), 30)),
		lipgloss.JoinHorizontal(lipgloss.Left,
			"Password : ", s.renderInput(),
		),
		"\n",
		s.optsView(),
	)
	style := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().FocusedBorder).
		Padding(1, 2)

	return lipgloss.JoinVertical(lipgloss.Center, lipgloss.PlaceVertical(s.Height-3, lipgloss.Center, lipgloss.JoinVertical(lipgloss.Center,
		styles.TitleStyle().Render("Login"),
		style.Render(lipgloss.PlaceHorizontal(s.Width, lipgloss.Left, view)),
	)),
		s.Help.View(),
	)
}

func (s *LoginModel) optsView() string {
	arr := make([]string, 0)
	clr := colors.Load()
	w := utils.DistributeWidth(s.Width, len(s.Opts))

	for k, v := range s.Opts {
		if s.SelectedOpt == k {
			arr = append(arr, lipgloss.PlaceHorizontal(w[k], lipgloss.Center, lipgloss.NewStyle().
				Background(clr.MenuSelectedBg).Foreground(clr.MenuSelectedText).
				Render(fmt.Sprintf(" %s ", v))))
			continue
		}

		arr = append(arr, lipgloss.PlaceHorizontal(w[k], lipgloss.Center, lipgloss.NewStyle().
			Foreground(clr.Text).
			Render(fmt.Sprintf(" %s ", v))))
	}

	return strings.Join(arr, "")
}

func (s *LoginModel) renderInput() string {
	view := s.PassTi.View()

	// If cursor is not at start → show left indicator
	if s.PassTi.Position() > s.PassTi.Width {
		view = "← " + view
	} else {
		view = "" + view
	}

	// If there's hidden text to the right
	if len(s.PassTi.Value()) > s.PassTi.Width && s.PassTi.Position() < len(s.PassTi.Value()) {
		view = view + " →"
	}

	return view
}
