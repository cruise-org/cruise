package config

type Keybinds struct {
	Global        GlobalKeybinds        `mapstructure:"global" toml:"global"`
	Nav           NavKeybinds           `mapstructure:"nav" toml:"nav"`
	Container     ContainersKeybinds    `mapstructure:"container" toml:"container"`
	Images        ImagesKeybinds        `mapstructure:"images" toml:"images"`
	Fzf           FzfKeybinds           `mapstructure:"fzf" toml:"fzf"`
	Monitoring    MonitorKeybinds       `mapstructure:"monitoring" toml:"monitoring"`
	Network       NetworkKeybinds       `mapstructure:"network" toml:"network"`
	Volumes       VolumeKeybinds        `mapstructure:"volume" toml:"volume"`
	Vulnerability VulnerabilityKeybinds `mapstructure:"vulnerability" toml:"vulnerability"`
}

type GlobalKeybinds struct {
	PageFinder    string `mapstructure:"page_finder" toml:"page_finder"`
	ListUp        string `mapstructure:"list_up" toml:"list_up"`
	ListDown      string `mapstructure:"list_down" toml:"list_down"`
	FocusSearch   string `mapstructure:"focus_search" toml:"focus_search"`
	UnfocusSearch string `mapstructure:"unfocus_search" toml:"unfocus_search"`
	QuickQuit     string `mapstructure:"quick_quit" toml:"quick_quit"`
}

type NavKeybinds struct {
	Exit          string `mapstructure:"exit" toml:"exit"`
	Dashboard     string `mapstructure:"dasboard" toml:"dasboard"`
	Containers    string `mapstructure:"containers" toml:"containers"`
	Images        string `mapstructure:"images" toml:"images"`
	Networks      string `mapstructure:"networks" toml:"networks"`
	Volumes       string `mapstructure:"volumes" toml:"volumes"`
	Monitoring    string `mapstructure:"monitoring" toml:"monitoring"`
	Vulnerability string `mapstructure:"vulnerability" toml:"vulnerability"`
	Projects      string `mapstructure:"projects" toml:"projects"`
}

type ContainersKeybinds struct {
	Start       string `mapstructure:"start" toml:"start"`
	Exec        string `mapstructure:"exec" toml:"exec"`
	Restart     string `mapstructure:"restart" toml:"restart"`
	Stop        string `mapstructure:"stop" toml:"stop"`
	Remove      string `mapstructure:"remove" toml:"remove"`
	Pause       string `mapstructure:"pause" toml:"pause"`
	Unpause     string `mapstructure:"unpause" toml:"unpause"`
	PortMap     string `mapstructure:"port_map" toml:"port_map"`
	ShowDetails string `mapstructure:"show_details" toml:"show_details"`
	ExitDetails string `mapstructure:"exit_details" toml:"exit_details"`
}

type FzfKeybinds struct {
	Up    string `mapstructure:"up" toml:"up"`
	Down  string `mapstructure:"down" toml:"down"`
	Enter string `mapstructure:"enter" toml:"enter"`
	Exit  string `mapstructure:"exit" toml:"exit"`
}

type ImagesKeybinds struct {
	Remove string `mapstructure:"remove" toml:"remove"`
	Prune  string `mapstructure:"prune" toml:"prune"`
	Push   string `mapstructure:"push" toml:"push"`
	Pull   string `mapstructure:"pull" toml:"pull"`
	Build  string `mapstructure:"build" toml:"build"`
	Layers string `mapstructure:"layers" toml:"layers"`
	Exit   string `mapstructure:"exit" toml:"exit"`
	Sync   string `mapstructure:"sync" toml:"sync"`
}

type MonitorKeybinds struct {
	Search     string `mapstructure:"search" toml:"search"`
	ExitSearch string `mapstructure:"exit_search" toml:"exit_search"`
	Export     string `mapstructure:"export" toml:"export"`
}

type NetworkKeybinds struct {
	Remove      string `mapstructure:"remove" toml:"remove"`
	Prune       string `mapstructure:"prune" toml:"prune"`
	ShowDetails string `mapstructure:"show_details" toml:"show_details"`
	ExitDetails string `mapstructure:"exit_details" toml:"exit_details"`
}

type VolumeKeybinds struct {
	Remove      string `mapstructure:"remove" toml:"remove"`
	Prune       string `mapstructure:"prune" toml:"prune"`
	ShowDetails string `mapstructure:"show_details" toml:"show_details"`
	ExitDetails string `mapstructure:"exit_details" toml:"exit_details"`
}

type VulnerabilityKeybinds struct {
	FocusScanners string `mapstructure:"focus_scanners" toml:"focus_scanners"`
	FocusList     string `mapstructure:"focus_list" toml:"focus_list"`
	Export        string `mapstructure:"export" toml:"export"`
}
