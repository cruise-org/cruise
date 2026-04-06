package registrymodel

import (
	"fmt"
	"log"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cruise-org/cruise/internal/messages"
	"github.com/cruise-org/cruise/pkg/config"
	"github.com/cruise-org/cruise/pkg/registry"
)

func (s *RegistryModel) nextLogin() tea.Cmd {
	s.CurrLogin = nil
	s.LoginModel = nil

	v, ok := <-s.LoginChan
	if ok {
		return func() tea.Msg { return v }
	}

	s.State = 4
	return nil
}

func (s *RegistryModel) parseRegistries() tea.Cmd {
	return func() tea.Msg {
		cfgRegistries := config.Cfg.Global.Registry
		registries := make([]*registry.Registry, 0)
		for _, v := range cfgRegistries {
			if v.Ignore {
				continue
			}

			r, err := registry.GetRegistry(&v)
			if err != nil {
				continue
			}

			registries = append(registries, &r)
		}

		return messages.ParsedRegistries{Registries: registries}
	}
}

func (s *RegistryModel) authenticateRegistries() tea.Cmd {
	return func() tea.Msg {
		var wg sync.WaitGroup
		ch := make(chan messages.RegistryLoginMessage, len(s.Registries))

		log.Println("[REG AUTH] Begin")

		for _, v := range s.Registries {
			wg.Add(1)
			go func(v *registry.Registry) {
				defer wg.Done()
				if !registry.IsLoggedIn(fmt.Sprintf("%s/%s", (*v).Provider(), (*v).Domain())) {
					ch <- messages.RegistryLoginMessage{Registry: v}
				}
				log.Printf("[REG AUTH] Checked Auth for %+v \n", v)
			}(v)
		}

		wg.Wait()
		close(ch)
		log.Println("[REG AUTH] Waited Success")

		return messages.PendingRegistryLogin{Ch: ch}
	}
}
