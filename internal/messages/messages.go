// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package messages

import (
	"github.com/cruise-org/cruise/internal/data"
	"github.com/cruise-org/cruise/pkg/enums"
	"github.com/cruise-org/cruise/pkg/registry"
	"github.com/cruise-org/cruise/pkg/types"
	"github.com/docker/docker/api/types/events"
)

type (
	// Registries
	ParsedRegistries     struct{ Registries []*registry.Registry }
	RegistryLoginMessage struct{ Registry *registry.Registry }
	RegistryLoginError   struct{ Err error }
	PendingRegistryLogin struct{ Ch chan RegistryLoginMessage }
	CloseLoginMessage    struct{}
	IgnoreLoginMessage   struct{}
	LoginMessage         struct {
		Pass string
	}

	HomeStatContainer struct{ Containers *[]types.Container }
	HomeStatImage     struct{ Images *[]types.Image }
	HomeStatNetwork   struct{ Networks *[]types.Network }
	HomeStatVolume    struct{ Volumes *[]types.Volume }

	HomeLogsTick    struct{}
	HomeLogsMonitor struct{ Monitor *types.Monitor }

	MonitoringTick    struct{}
	MonitoringMonitor struct{ Monitor *types.Monitor }

	DetailRendererInit struct {
		Stats *[]types.StatCard
		Meta  *types.StatMeta
	}
	DetailRendererContent struct{ VPMap *map[string]map[string]string }

	ContainerDetailsMonitorReady struct{ Monitor *types.Monitor }

	ChangePg struct {
		Pg     enums.PageType
		Exited bool
	}

	CloseDetails struct{}

	PortMapMsg struct {
		Ports []string
		Err   error
	}

	ErrorMsg struct {
		Title string
		Msg   string
		Locn  string
	}

	CloseError struct{}

	MsgPopup struct {
		Title string
		Msg   string
		Locn  string
	}

	CloseMsgPopup struct{}

	StartScanMsg struct {
		Img string
	}

	ScanResponse struct {
		Arr *[]types.Vulnerability
	}

	ScannerListMsg struct {
		Found []bool
	}

	ContainerDetailsTick struct{}

	ContainerReadyMsg struct {
		Items *[]types.Container
		Err   error
	}

	ImagesReadyMsg struct {
		Map map[string]types.Image
	}

	UpdateImagesMsg struct {
		Items *[]types.Image
	}

	NetworksReadyMsg struct {
		Items *[]types.Network
	}

	VolumesReadyMsg struct {
		Items *[]types.Volume
	}

	UpdateNetworksMsg struct{}

	FzfSelection struct {
		Selection string
		Exited    bool
	}

	SysResReadyMsg struct {
		CPU  *data.CPUInfo
		Mem  *data.MemInfo
		Disk *data.DiskInfo
	}

	NewEvents struct {
		Events []*events.Message
	}

	DaemonReadyMsg struct {
		Err error
	}
)
