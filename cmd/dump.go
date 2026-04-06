// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cruise-org/cruise/pkg/config"
	"github.com/spf13/cobra"
	"go.yaml.in/yaml/v3"
)

var dumpPath string

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump the config file at the given filepath.",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		fmt.Println("Dumping config to:", path)

		if filepath.Ext(path) == "" {
			path = filepath.Join(path, "dump.yaml") // changed
		}

		f, err := os.Create(path)
		if err != nil {
			fmt.Println("Error Creating File: ", err.Error())
			return
		}
		defer f.Close()

		encoder := yaml.NewEncoder(f)
		err = encoder.Encode(config.Default())
		if err != nil {
			fmt.Println("Failed to write to file: ", err.Error())
			return
		}
		fmt.Println("Successfully Dumped file to " + path)
	},
}

func init() {
	dumpCmd.Flags().StringVarP(&dumpPath, "path", "p",
		filepath.Join(config.GetCfgDir(), "dump.yaml"), // changed
		"filepath to dump the config")
	rootCmd.AddCommand(dumpCmd)
}
