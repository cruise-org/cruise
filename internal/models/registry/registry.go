package registrymodel

import (
	"log"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cruise-org/cruise/internal/messages"
	"github.com/cruise-org/cruise/internal/models/registry/loginmodel"
	"github.com/cruise-org/cruise/pkg/page"
	"github.com/cruise-org/cruise/pkg/registry"
)

type RegistryModel struct {
	Width  int
	Height int
	State  int // Defines current phase in initialization, check View() for exact
	// Data
	Registries []*registry.Registry
	LoginChan  chan messages.RegistryLoginMessage
	LoginModel *loginmodel.LoginModel
	CurrLogin  *registry.Registry
}

func NewRegistryModel(w, h int) *RegistryModel {
	return &RegistryModel{
		Width:      w,
		Height:     h,
		State:      1,
		LoginModel: nil,
	}
}

func (s *RegistryModel) Cleanup() {}

func (s *RegistryModel) Init() tea.Cmd { return s.parseRegistries() }

func (s *RegistryModel) Update(msg tea.Msg) (page.Page, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.ParsedRegistries:
		s.Registries = msg.Registries
		s.State = 2
		for _, v := range s.Registries {
			log.Printf("[REG] %+v", *v)
		}
		return s, s.authenticateRegistries()
	case messages.PendingRegistryLogin:
		s.LoginChan = msg.Ch
		v, ok := <-s.LoginChan
		if ok {
			return s, func() tea.Msg { return v }
		}
		log.Printf("[REG] Pending Registry Login Came %+v \n", v)
		return s, nil
	case messages.RegistryLoginMessage:
		log.Printf("[REG] Registry Login Msg Came for: %+v", msg.Registry)
		if s.LoginModel != nil || s.CurrLogin != nil {
			return s, nil
		}

		s.LoginModel = loginmodel.NewLoginModel(s.Width, s.Height, msg.Registry)
		s.CurrLogin = msg.Registry
		return s, nil

	case messages.IgnoreLoginMessage:
		if s.CurrLogin == nil {
			return s, nil
		}
		for k, v := range s.Registries {
			if s.CurrLogin == v {
				s.Registries = slices.Delete(s.Registries, k, k+1)
				break
			}
		}

		return s, s.nextLogin()

	case messages.LoginMessage:
		err := (*s.CurrLogin).Login(msg.Pass)
		if err != nil {
			log.Printf("[LOGIN] %v \n", err)
			s.LoginModel.Update(messages.RegistryLoginError{Err: err})
			return s, nil
		}

		return s, s.nextLogin()

	case tea.KeyMsg:
		if s.LoginModel != nil {
			lm, cmd := s.LoginModel.Update(msg)
			s.LoginModel = lm
			return s, cmd
		}
	}
	return s, nil
}

func (s *RegistryModel) View() string {
	if s.LoginModel != nil {
		return lipgloss.Place(s.Width, s.Height, lipgloss.Center, lipgloss.Center, s.LoginModel.View())
	}
	switch s.State {
	case 1:
		return "Parsing Registries..."
	case 2:
		return "Checking Authentication Status..."
	case 3:
		return "Logging In..."
	}

	return "Registry Model"
}
