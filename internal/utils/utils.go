// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package utils

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cruise-org/cruise/internal/messages"
	"github.com/cruise-org/cruise/pkg/enums"
	"github.com/docker/docker/api/types/container"
)

func DistributeWidth(w, parts int) []int {
	if parts <= 0 {
		return nil
	}

	base := w / parts
	rem := w % parts

	result := make([]int, parts)

	// Fill with base width
	for i := range result {
		result[i] = base
	}

	// Distribute remainder centered
	mid := parts / 2

	if parts%2 == 1 {
		// Odd → start exactly at center
		left := mid
		right := mid + 1

		for rem > 0 {
			result[left]++
			rem--
			if rem == 0 {
				break
			}

			if right < parts {
				result[right]++
				rem--
			}

			left--
			right++
			if left < 0 {
				left = 0
			}
		}
	} else {
		// Even → start from middle pair
		left := mid - 1
		right := mid

		for rem > 0 {
			if left >= 0 {
				result[left]++
				rem--
			}
			if rem == 0 {
				break
			}
			if right < parts {
				result[right]++
				rem--
			}

			left--
			right++
		}
	}

	return result
}

func ReturnError(loc, title string, err error) tea.Cmd {
	return func() tea.Msg {
		return messages.ErrorMsg{
			Locn:  loc,
			Title: title,
			Msg:   err.Error(),
		}
	}
}

func ReturnMsg(loc, title, msg string) tea.Cmd {
	return func() tea.Msg {
		return messages.MsgPopup{
			Locn:  loc,
			Title: title,
			Msg:   msg,
		}
	}
}

func FormatDuration(d time.Duration) string {
	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	if days > 0 {
		return fmt.Sprintf("%dd %dh %dm", days, hours, minutes)
	} else if hours > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	} else if minutes > 0 {
		return fmt.Sprintf("%dm %ds", minutes, seconds)
	}
	return fmt.Sprintf("%ds", seconds)
}

func ShortID(id string) string {
	if len(id) > 12 {
		return id[:12]
	}
	return id
}

func CreatedAgo(ts int64) string {
	return time.Since(time.Unix(ts, 0)).Round(time.Second).String() + " ago"
}

func FormatPorts(ports []container.Port) string {
	if len(ports) == 0 {
		return "-"
	}
	var result []string
	for _, p := range ports {
		if p.PublicPort != 0 {
			result = append(result, fmt.Sprintf("%d->%d/%s", p.PublicPort, p.PrivatePort, p.Type))
		} else {
			result = append(result, fmt.Sprintf("%d/%s", p.PrivatePort, p.Type))
		}
	}
	return strings.Join(result, ",")
}

func FormatMounts(mounts []container.MountPoint) string {
	if len(mounts) == 0 {
		return "-"
	}
	var result []string
	for _, m := range mounts {
		result = append(result, m.Destination)
	}
	return strings.Join(result, ",")
}

func FormatSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func Shorten(s string, max int) string {
	if max <= 0 {
		return ""
	}

	if len(s) <= max {
		return s
	}

	if max <= 3 {
		return s[:max]
	}

	return s[:max-3] + "..."
}

func CalculateCPUPercent(stats container.StatsResponse) float64 {
	cpuDelta := float64(stats.CPUStats.CPUUsage.TotalUsage - stats.PreCPUStats.CPUUsage.TotalUsage)
	systemDelta := float64(stats.CPUStats.SystemUsage - stats.PreCPUStats.SystemUsage)
	onlineCPUs := float64(stats.CPUStats.OnlineCPUs)
	if onlineCPUs == 0 {
		onlineCPUs = float64(len(stats.CPUStats.CPUUsage.PercpuUsage)) // Fallback
	}

	if systemDelta > 0.0 && cpuDelta > 0.0 {
		return (cpuDelta / systemDelta) * onlineCPUs * 100.0
	}
	return 0.0
}

func CalculateMemoryPercent(stats container.StatsResponse) float64 {
	used := float64(stats.MemoryStats.Usage - stats.MemoryStats.Stats["cache"])
	total := float64(stats.MemoryStats.Limit)

	percent := (used / total) * 100.0

	return percent
}

// Wraps the text to the given length and also limits no. of lines, adds a "..." line if exceeding.
func WrapAndLimit(s string, maxLen, maxLines int) string {
	var lines []string

	for i := 0; i < len(s); i += maxLen {
		end := i + maxLen
		if end > len(s) {
			end = len(s)
		}
		lines = append(lines, s[i:end])
	}

	if len(lines) > maxLines {
		lines = append(lines[:maxLines], "...")
	}

	return strings.Join(lines, "\n")
}

func ToAnySlice[T any](in []T) []any {
	out := make([]any, len(in))
	for i := range in {
		out[i] = in[i]
	}
	return out
}

func GetSeverity(scanner, severity string) enums.Severity {
	switch severity {
	case "CRITICAL":
		return enums.Critical
	case "HIGH":
		return enums.High
	case "MEDIUM":
		return enums.Medium
	case "LOW":
		return enums.Low
	default:
		return enums.Unknown
	}
}

func SeverityText(sev enums.Severity) string {
	switch sev {
	case enums.Critical:
		return "CRITICAL"
	case enums.High:
		return "HIGH"
	case enums.Medium:
		return "MEDIUM"
	case enums.Low:
		return "LOW"
	default:
		return "UKNOWN"
	}
}
