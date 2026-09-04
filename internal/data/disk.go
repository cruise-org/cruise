// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package data

import (
	"github.com/shirou/gopsutil/v3/disk"
)

type DiskInfo struct {
	Total float64
	Used  float64
	Usage float64
	Err   error
}

func GetDiskInfo() *DiskInfo {
	usage, err := disk.Usage("/")
	if err != nil {
		return &DiskInfo{Err: err}
	}

	// Convert bytes to GB with 1 decimal point
	totalGB := float64(usage.Total) / (1 << 30)
	usedGB := float64(usage.Used) / (1 << 30)
	usedPercent := usage.UsedPercent

	return &DiskInfo{
		Total: totalGB,
		Used:  usedGB,
		Usage: usedPercent,
		Err:   err,
	}
}
