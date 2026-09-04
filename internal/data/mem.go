// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package data

import "github.com/shirou/gopsutil/v3/mem"

type MemInfo struct {
	Total float64
	Used  float64
	Usage float64
	Err   error
}

func GetMemInfo() *MemInfo {
	v, err := mem.VirtualMemory()
	if err != nil {
		return &MemInfo{Err: err}
	}

	totalGB := float64(v.Total) / (1 << 30)
	usedGB := float64(v.Used) / (1 << 30)
	usedPercent := v.UsedPercent

	return &MemInfo{
		Total: totalGB,
		Used:  usedGB,
		Usage: usedPercent,
		Err:   err,
	}
}
