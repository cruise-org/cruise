// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package config

type Styles struct {
	Text             string `mapstructure:"text" yaml:"text"`
	SubtitleText     string `mapstructure:"subtitle_text" yaml:"subtitle_text"`
	SubtitleBg       string `mapstructure:"subtitle_bg" yaml:"subtitle_bg"`
	UnfocusedBorder  string `mapstructure:"unfocused_border" yaml:"unfocused_border"`
	FocusedBorder    string `mapstructure:"focused_border" yaml:"focused_border"`
	HelpKeyBg        string `mapstructure:"help_key_bg" yaml:"help_key_bg"`
	HelpKeyText      string `mapstructure:"help_key_text" yaml:"help_key_text"`
	HelpDescText     string `mapstructure:"help_desc_text" yaml:"help_desc_text"`
	MenuSelectedBg   string `mapstructure:"menu_selected_bg" yaml:"menu_selected_bg"`
	MenuSelectedText string `mapstructure:"menu_selected_text" yaml:"menu_selected_text"`
	ErrorText        string `mapstructure:"error_text" yaml:"error_text"`
	ErrorBg          string `mapstructure:"error_bg" yaml:"error_bg"`
	PopupBorder      string `mapstructure:"popup_border" yaml:"popup_border"`
	PlaceholderText  string `mapstructure:"placeholder_text" yaml:"placeholder_text"`
	MsgText          string `mapstructure:"msg_text" yaml:"msg_text"`
}
