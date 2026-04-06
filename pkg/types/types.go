// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package types

type ContainerState = string

const (
	StateCreated    ContainerState = "created"    // StateCreated indicates the container is created, but not (yet) started.
	StateRunning    ContainerState = "running"    // StateRunning indicates that the container is running.
	StatePaused     ContainerState = "paused"     // StatePaused indicates that the container's current state is paused.
	StateRestarting ContainerState = "restarting" // StateRestarting indicates that the container is currently restarting.
	StateRemoving   ContainerState = "removing"   // StateRemoving indicates that the container is being removed.
	StateExited     ContainerState = "exited"     // StateExited indicates that the container exited.
	StateDead       ContainerState = "dead"       // StateDead indicates that the container failed to be deleted. Containers in this state are attempted to be cleaned up when the daemon restarts.
)

type Container struct {
	ID      string
	Runtime string
	Name    string
	Image   string
	Created int64
	Ports   []ContainerPort
	State   ContainerState
	Mounts  []ContainerMount
	Labels  map[string]string
}

type ContainerMount struct {
	Source      string
	Destination string
	ReadOnly    bool
}

type ContainerPort struct {
	ContainerPort uint16
	HostPort      uint16
	Protocol      string
}

type Image struct {
	ID            string
	Runtime       string
	Tags          []string
	Size          int64
	CreatedAt     int64
	NumContainers int64
}

type Network struct {
	ID            string
	Runtime       string
	Name          string
	Scope         string
	Driver        string
	IPv4          bool
	NumContainers int
}

type Volume struct {
	Name       string
	Runtime    string
	Scope      string
	Driver     string
	Mountpoint string
	CreatedAt  string
}

type RegistryImage struct {
	Domain      string `json:"domain"`
	Provider    string `json:"provider"`
	Name        string `json:"name"`
	Project     string `json:"project"`
	Description string `json:"description"`
	IsPrivate   bool   `json:"is_private"`
	PullCount   int64  `json:"pull_count"`
}
