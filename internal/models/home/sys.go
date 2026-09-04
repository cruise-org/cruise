// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package home

import (
	"fmt"
	"math"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cruise-org/cruise/internal/data"
	"github.com/cruise-org/cruise/internal/messages"
	"github.com/cruise-org/cruise/pkg/styles"
)

const (
	barWidth        = 49
	sysRefreshEvery = 2 * time.Second
)

type SysRes struct {
	Width     int
	Height    int
	IsLoading bool
	CPU       *data.CPUInfo
	Mem       *data.MemInfo
	Disk      *data.DiskInfo
}

func NewSysRes(w int, h int) *SysRes {
	return &SysRes{
		Width:     w,
		Height:    h,
		IsLoading: true,
	}
}

var refresh = func(t time.Time) tea.Msg {
	cpuChan := make(chan *data.CPUInfo, 1)
	memChan := make(chan *data.MemInfo, 1)
	diskChan := make(chan *data.DiskInfo, 1)
	go func() {
		cpuChan <- data.GetCPUInfo()
	}()
	go func() {
		memChan <- data.GetMemInfo()
	}()
	go func() {
		diskChan <- data.GetDiskInfo()
	}()
	return messages.SysResReadyMsg{
		CPU:  <-cpuChan,
		Mem:  <-memChan,
		Disk: <-diskChan,
	}
}

func (s *SysRes) Init() tea.Cmd {
	return tea.Tick(0, refresh)
}

func (s *SysRes) Update(msg tea.Msg) (*SysRes, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.SysResReadyMsg:
		s.IsLoading = false
		s.CPU = msg.CPU
		s.Mem = msg.Mem
		s.Disk = msg.Disk
		return s, tea.Tick(sysRefreshEvery, refresh)
	}
	return s, nil
}

func (s *SysRes) View() string {
	return styles.SubpageStyle().PaddingTop(1).PaddingLeft(4).Render(lipgloss.JoinVertical(lipgloss.Center,
		styles.TitleStyle().Render("System Resources"),
		lipgloss.NewStyle().
			Width(s.Width-6).   //-6 from padding(4) and border(2)
			Height(s.Height-4). //-4 from title(1) border(2) and padding(1)
			Align(lipgloss.Left, lipgloss.Center).
			Render(s.FormattedView())))
}

func bar(usage float64) string {
	filled := int((usage / 100) * float64(barWidth))
	if filled < 0 {
		filled = 0
	}
	if filled > barWidth {
		filled = barWidth
	}
	return strings.Repeat("█", filled) + strings.Repeat(" ", barWidth-filled)
}

func (s SysRes) FormattedView() string {
	if s.IsLoading || s.CPU == nil || s.Mem == nil || s.Disk == nil {
		return "Querying System Data..."
	}

	cputext := ""
	if s.CPU.Error != nil {
		cputext = fmt.Sprintf("ERROR: %s", s.CPU.Error.Error())
	} else {
		cputext = fmt.Sprintf("CPU:  [%s] %.1f%% | %.1fGhz - %dL/%dP Cores", bar(s.CPU.Usage), math.Round(s.CPU.Usage*10)/10, math.Round(s.CPU.Mhz/100)/10,
			s.CPU.LogicCores, s.CPU.PhysicalCores)
	}

	memtext := ""
	if s.Mem.Err != nil {
		memtext = fmt.Sprintf("ERROR: %s", s.Mem.Err.Error())
	} else {
		memtext = fmt.Sprintf("Mem:  [%s] %.1f%% | %.1fGB / %.1fGB", bar(s.Mem.Usage), s.Mem.Usage, s.Mem.Used, s.Mem.Total)
	}

	disktext := ""
	if s.Disk.Err != nil {
		disktext = fmt.Sprintf("ERROR: %s", s.Disk.Err.Error())
	} else {
		disktext = fmt.Sprintf("Disk: [%s] %.1f%% | %.1fGB / %.1fGB", bar(s.Disk.Usage), s.Disk.Usage, s.Disk.Used, s.Disk.Total)
	}

	return fmt.Sprintf("%s\n\n%s\n\n%s", cputext, memtext, disktext)
}
