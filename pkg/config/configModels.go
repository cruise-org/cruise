// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package config

type Config struct {
	Global   Global   `mapstructure:"global" yaml:"global"`
	Styles   Styles   `mapstructure:"styles" yaml:"styles"`
	Keybinds Keybinds `mapstructure:"keybinds" yaml:"keybinds"`
}

type Global struct {
	ExportDir string           `mapstructure:"export_dir" yaml:"export_dir"`
	Term      string           `mapstructure:"term" yaml:"term"`
	Runtimes  []string         `mapstructure:"runtimes" yaml:"runtimes"`
	Registry  []RegistryConfig `mapstructure:"registry" yaml:"registry"`
}

type RegistryConfig struct {
	Provider string `mapstructure:"provider" yaml:"provider"`
	Username string `mapstructure:"username" yaml:"username"`
	Domain   string `mapstructure:"domain" yaml:"domain"`
	Ignore   bool   `mapstructure:"ignore" yaml:"ignore"`
}
