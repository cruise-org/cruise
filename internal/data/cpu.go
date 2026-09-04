// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package data

import (
	"errors"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

type CPUInfo struct {
	Usage         float64
	LogicCores    int
	PhysicalCores int
	Mhz           float64
	Error         error
}

func GetCPUInfo() *CPUInfo {
	c, err := cpu.Info()
	if err != nil {
		return &CPUInfo{Error: err}
	}
	if len(c) == 0 {
		return &CPUInfo{Error: errors.New("no cpu info available")}
	}
	cp := c[0]

	usage, err := cpu.Percent(time.Second, false) // Overall usage
	if err != nil {
		return &CPUInfo{Error: err}
	}
	if len(usage) == 0 {
		return &CPUInfo{Error: errors.New("no cpu usage available")}
	}

	logical, err := cpu.Counts(true)
	if err != nil {
		return &CPUInfo{Error: err}
	}

	physical, err := cpu.Counts(false)
	if err != nil {
		return &CPUInfo{Error: err}
	}

	return &CPUInfo{
		Usage:         usage[0],
		LogicCores:    logical,
		PhysicalCores: physical,
		Mhz:           cp.Mhz,
		Error:         nil,
	}
}
