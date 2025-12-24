package keymap

import (
	"reflect"

	"github.com/NucleoFusion/cruise/internal/config"
	"github.com/charmbracelet/bubbles/key"
)

type NavMap struct {
	Exit          key.Binding
	Dashboard     key.Binding
	Containers    key.Binding
	Images        key.Binding
	Networks      key.Binding
	Volumes       key.Binding
	Monitoring    key.Binding
	Vulnerability key.Binding
	Projects      key.Binding
}

func NewNavMap() NavMap {
	m := config.Cfg.Keybinds.Nav
	return NavMap{
		Exit: key.NewBinding(
			key.WithKeys(m.Exit),
			key.WithHelp(m.Exit, "exit"),
		),
		Dashboard: key.NewBinding(
			key.WithKeys(m.Dashboard),
			key.WithHelp(m.Dashboard, "dashboard"),
		),
		Containers: key.NewBinding(
			key.WithKeys(m.Containers),
			key.WithHelp(m.Containers, "containers"),
		),
		Images: key.NewBinding(
			key.WithKeys(m.Images),
			key.WithHelp(m.Images, "images"),
		),
		Networks: key.NewBinding(
			key.WithKeys(m.Networks),
			key.WithHelp(m.Networks, "networks"),
		),
		Volumes: key.NewBinding(
			key.WithKeys(m.Volumes),
			key.WithHelp(m.Volumes, "volumes"),
		),
		Monitoring: key.NewBinding(
			key.WithKeys(m.Monitoring),
			key.WithHelp(m.Monitoring, "monitoring"),
		),
		Vulnerability: key.NewBinding(
			key.WithKeys(m.Vulnerability),
			key.WithHelp(m.Vulnerability, "vulnerability"),
		),
		Projects: key.NewBinding(
			key.WithKeys(m.Projects),
			key.WithHelp(m.Projects, "projects"),
		),
	}
}

func (m NavMap) Bindings() []key.Binding {
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
