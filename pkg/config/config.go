// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var Cfg Config

func SetCfg() error {
	cfg := GetCfgDir()
	p := filepath.Join(cfg, "config.yaml")

	if _, err := os.Stat(p); os.IsNotExist(err) {
		if err := os.MkdirAll(cfg, 0o755); err != nil {
			fmt.Printf("failed to create config dir: %s", err.Error())
			os.Exit(1)
		}

		if _, err := os.Create(p); err != nil {
			fmt.Printf("failed to create config file: %s", err.Error())
			os.Exit(1)
		}
	}

	viper.SetConfigFile(p)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	Cfg = Default()

	if err := viper.Unmarshal(&Cfg); err != nil {
		return err
	}

	return nil
}

func GetCfgDir() string {
	switch runtime.GOOS {
	case "linux", "darwin":
		home, _ := os.UserHomeDir()
		return filepath.Join(home, ".config", "cruise")
	case "windows":
		home, _ := os.UserHomeDir()
		return filepath.Join(home, ".cruise")
	default:
		cfg, _ := os.UserConfigDir()
		return cfg
	}
}
