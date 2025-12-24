package config

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func Default() Config {
	expDir := GetDefExportDir()

	if _, err := os.Stat(expDir); os.IsNotExist(err) {
		if err := os.MkdirAll(expDir, 0o755); err != nil {
			fmt.Printf("failed to create config dir: %s", err.Error())
			os.Exit(1)
		}
	}

	return Config{
		Global: Global{
			ExportDir: expDir,
			Term:      DetectTerminal(),
		},
		Keybinds: Keybinds{
			Global: GlobalKeybinds{
				PageFinder:    "tab",
				ListUp:        "up",
				ListDown:      "down",
				FocusSearch:   "/",
				UnfocusSearch: "esc",
				QuickQuit:     "q",
			},
			Nav: NavKeybinds{ // TODO: In Docs
				Exit:          "esc",
				Dashboard:     "d",
				Containers:    "c",
				Images:        "i",
				Networks:      "n",
				Volumes:       "v",
				Monitoring:    "m",
				Vulnerability: "s",
				Projects:      "p",
			},
			Fzf: FzfKeybinds{
				Up:    "up",
				Down:  "down",
				Enter: "enter",
				Exit:  "esc",
			},
			Container: ContainersKeybinds{
				Start:       "s",
				Stop:        "t",
				Remove:      "d",
				Restart:     "r",
				Pause:       "p",
				Unpause:     "u",
				Exec:        "e",
				ShowDetails: "enter",
				ExitDetails: "esc",
				PortMap:     "m",
			},
			Images: ImagesKeybinds{
				Remove: "r",
				Prune:  "d",
				Push:   "p",
				Pull:   "u",
				Build:  "b",
				Layers: "l",
				Sync:   "s",
			},
			Network: NetworkKeybinds{
				ExitDetails: "esc",
				ShowDetails: "enter",
				Prune:       "p",
				Remove:      "r",
			},
			Volumes: VolumeKeybinds{
				ExitDetails: "esc",
				ShowDetails: "enter",
				Prune:       "p",
				Remove:      "r",
			},
			Vulnerability: VulnerabilityKeybinds{
				FocusScanners: "S",
				FocusList:     "L",
			},
			Monitoring: MonitorKeybinds{
				Search:     "/",
				ExitSearch: "esc",
			},
		},
		Styles: Styles{
			Text:             "#cdd6f4",
			SubtitleText:     "#74c7ec",
			SubtitleBg:       "#313244",
			MenuSelectedBg:   "#b4befe",
			MenuSelectedText: "#1e1e2e",
			FocusedBorder:    "#b4befe",
			UnfocusedBorder:  "#45475a",
			HelpKeyBg:        "#313244",
			HelpKeyText:      "#cdd6f4",
			HelpDescText:     "#6c7086",
			ErrorText:        "#f38ba8",
			ErrorBg:          "#11111b",
			PopupBorder:      "#74c7ec",
			PlaceholderText:  "#585b70",
			MsgText:          "#74c7ec",
		},
	}
}

func GetDefExportDir() string {
	switch runtime.GOOS {
	case "linux", "darwin":
		home, _ := os.UserHomeDir()
		return filepath.Join(home, ".cruise")
	case "windows":
		home, _ := os.UserHomeDir()
		return filepath.Join(home, ".cruise", "export")
	default:
		cfg, _ := os.UserConfigDir()
		return cfg
	}
}

func DetectTerminal() string {
	if term := os.Getenv("TERMINAL"); term != "" {
		return term
	}

	switch runtime.GOOS {
	case "windows":
		// Prefer Windows Terminal (wt.exe) if installed
		if _, err := exec.LookPath("wt.exe"); err == nil {
			return "wt.exe"
		}
		if comspec := os.Getenv("ComSpec"); comspec != "" {
			return comspec
		}
		return "cmd.exe"

	case "darwin":
		return "open -a Terminal"

	case "linux":
		if _, err := exec.LookPath("x-terminal-emulator"); err == nil {
			return "x-terminal-emulator"
		}
		// Common terminals
		candidates := []string{
			"gnome-terminal",
			"konsole",
			"xfce4-terminal",
			"xterm",
		}
		for _, c := range candidates {
			if _, err := exec.LookPath(c); err == nil {
				return c
			}
		}
		if sh := os.Getenv("SHELL"); sh != "" {
			return sh
		}
		return "xterm"

	default:
		return "sh"
	}
}
