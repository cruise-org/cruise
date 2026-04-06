// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package config

type Keybinds struct {
	Global        GlobalKeybinds        `mapstructure:"global" yaml:"global"`
	Nav           NavKeybinds           `mapstructure:"nav" yaml:"nav"`
	Container     ContainersKeybinds    `mapstructure:"container" yaml:"container"`
	Images        ImagesKeybinds        `mapstructure:"images" yaml:"images"`
	Fzf           FzfKeybinds           `mapstructure:"fzf" yaml:"fzf"`
	Monitoring    MonitorKeybinds       `mapstructure:"monitoring" yaml:"monitoring"`
	Network       NetworkKeybinds       `mapstructure:"network" yaml:"network"`
	Volumes       VolumeKeybinds        `mapstructure:"volume" yaml:"volume"`
	Vulnerability VulnerabilityKeybinds `mapstructure:"vulnerability" yaml:"vulnerability"`
	RegistryLogin RegistryLoginKeybinds `mapstructure:"registry_login" yaml:"registry_login"`
}

type GlobalKeybinds struct {
	PageFinder    string `mapstructure:"page_finder" yaml:"page_finder"`
	ListUp        string `mapstructure:"list_up" yaml:"list_up"`
	ListDown      string `mapstructure:"list_down" yaml:"list_down"`
	FocusSearch   string `mapstructure:"focus_search" yaml:"focus_search"`
	UnfocusSearch string `mapstructure:"unfocus_search" yaml:"unfocus_search"`
	QuickQuit     string `mapstructure:"quick_quit" yaml:"quick_quit"`
}

type NavKeybinds struct {
	Exit          string `mapstructure:"exit" yaml:"exit"`
	Dashboard     string `mapstructure:"dasboard" yaml:"dasboard"`
	Containers    string `mapstructure:"containers" yaml:"containers"`
	Images        string `mapstructure:"images" yaml:"images"`
	Networks      string `mapstructure:"networks" yaml:"networks"`
	Volumes       string `mapstructure:"volumes" yaml:"volumes"`
	Monitoring    string `mapstructure:"monitoring" yaml:"monitoring"`
	Vulnerability string `mapstructure:"vulnerability" yaml:"vulnerability"`
	Registry      string `mapstructure:"registry" yaml:"registry"`
}

type ContainersKeybinds struct {
	Start       string `mapstructure:"start" yaml:"start"`
	Exec        string `mapstructure:"exec" yaml:"exec"`
	Restart     string `mapstructure:"restart" yaml:"restart"`
	Stop        string `mapstructure:"stop" yaml:"stop"`
	Remove      string `mapstructure:"remove" yaml:"remove"`
	Pause       string `mapstructure:"pause" yaml:"pause"`
	Unpause     string `mapstructure:"unpause" yaml:"unpause"`
	PortMap     string `mapstructure:"port_map" yaml:"port_map"`
	ShowDetails string `mapstructure:"show_details" yaml:"show_details"`
	ExitDetails string `mapstructure:"exit_details" yaml:"exit_details"`
}

type FzfKeybinds struct {
	Up    string `mapstructure:"up" yaml:"up"`
	Down  string `mapstructure:"down" yaml:"down"`
	Enter string `mapstructure:"enter" yaml:"enter"`
	Exit  string `mapstructure:"exit" yaml:"exit"`
}

type ImagesKeybinds struct {
	Remove string `mapstructure:"remove" yaml:"remove"`
	Prune  string `mapstructure:"prune" yaml:"prune"`
	Push   string `mapstructure:"push" yaml:"push"`
	Pull   string `mapstructure:"pull" yaml:"pull"`
	Build  string `mapstructure:"build" yaml:"build"`
	Layers string `mapstructure:"layers" yaml:"layers"`
	Exit   string `mapstructure:"exit" yaml:"exit"`
	Sync   string `mapstructure:"sync" yaml:"sync"`
}

type MonitorKeybinds struct {
	Search     string `mapstructure:"search" yaml:"search"`
	ExitSearch string `mapstructure:"exit_search" yaml:"exit_search"`
	Export     string `mapstructure:"export" yaml:"export"`
}

type NetworkKeybinds struct {
	Remove      string `mapstructure:"remove" yaml:"remove"`
	Prune       string `mapstructure:"prune" yaml:"prune"`
	ShowDetails string `mapstructure:"show_details" yaml:"show_details"`
	ExitDetails string `mapstructure:"exit_details" yaml:"exit_details"`
}

type VolumeKeybinds struct {
	Remove      string `mapstructure:"remove" yaml:"remove"`
	Prune       string `mapstructure:"prune" yaml:"prune"`
	ShowDetails string `mapstructure:"show_details" yaml:"show_details"`
	ExitDetails string `mapstructure:"exit_details" yaml:"exit_details"`
}

type VulnerabilityKeybinds struct {
	FocusScanners string `mapstructure:"focus_scanners" yaml:"focus_scanners"`
	FocusList     string `mapstructure:"focus_list" yaml:"focus_list"`
	Export        string `mapstructure:"export" yaml:"export"`
}

type RegistryLoginKeybinds struct {
	Left  string `mapstructure:"left" yaml:"left"`
	Right string `mapstructure:"right" yaml:"right"`
	Enter string `mapstructure:"enter" yaml:"enter"`
}
