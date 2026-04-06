package keymap

import (
	"reflect"

	"github.com/charmbracelet/bubbles/key"
	"github.com/cruise-org/cruise/pkg/config"
)

type LoginModelKeymap struct {
	Left  key.Binding
	Right key.Binding
	Enter key.Binding
}

func NewLoginModelMap() LoginModelKeymap {
	m := config.Cfg.Keybinds.RegistryLogin
	return LoginModelKeymap{
		Left: key.NewBinding(
			key.WithKeys(m.Left),
			key.WithHelp(m.Left, "left"),
		),
		Right: key.NewBinding(
			key.WithKeys(m.Right),
			key.WithHelp(m.Right, "right"),
		),
		Enter: key.NewBinding(
			key.WithKeys(m.Enter),
			key.WithHelp(m.Enter, "enter"),
		),
	}
}

func (m LoginModelKeymap) Bindings() []key.Binding {
	var bindings []key.Binding

	v := reflect.ValueOf(m)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if binding, ok := field.Interface().(key.Binding); ok {
			bindings = append(bindings, binding)
		}
	}

	return bindings
}
