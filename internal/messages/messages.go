package messages

import (
	"encoding/json"
	"io"
	"time"

	"github.com/NucleoFusion/cruise/internal/data"
	"github.com/NucleoFusion/cruise/internal/enums"
	internaltypes "github.com/NucleoFusion/cruise/internal/types"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/volume"
)

type DashboardTick time.Time

func TickDashboard() tea.Cmd {
	return tea.Tick(5*time.Second, func(t time.Time) tea.Msg {
		return DashboardTick(t)
	})
}

type (
	ProjectsReadyMsg struct {
		Projects []*internaltypes.Project
	}

	ProjectsDetailsUpdate struct{}

	ShowProjectDetails struct {
		Project *internaltypes.Project
	}

	CloseProjectDetails struct{}

	ChangePg struct {
		Pg     enums.PageType
		Exited bool
	}

	CloseDetails struct{}

	PortMapMsg struct {
		Arr []string
		Err error
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
		Arr []any
		Err error
	}

	ScannerListMsg struct {
		Found []bool
	}

	NewContainerDetails struct {
		Stats   container.StatsResponseReader
		Decoder *json.Decoder
		Logs    *io.ReadCloser
	}

	ContainerDetailsReady struct {
		Stats   container.StatsResponseReader
		Decoder *json.Decoder
		Logs    *io.ReadCloser
	}

	ContainerDetailsTick struct{}

	ContainerReadyMsg struct {
		Items []container.Summary
		Err   error
	}

	ImagesReadyMsg struct {
		Map map[string]image.Summary
	}

	UpdateImagesMsg struct {
		Items []image.Summary
	}

	NetworksReadyMsg struct {
		Items []network.Summary
	}

	VolumesReadyMsg struct {
		Items []*volume.Volume
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
