package internaltypes

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/compose-spec/compose-go/v2/types"
	"github.com/docker/docker/api/types/container"
)

type Project struct {
	Inspect  *types.Project
	Services *map[string]Service
}

type Service struct {
	Containers *[]ContainerDetails
}

type ContainerDetails struct {
	Inspect *container.InspectResponse
	Stats   *container.StatsResponseReader
	Decoder *json.Decoder
}

type StatsDetails struct {
	CPU      uint64
	Mem      int64
	MemLimit int64
	NetRx    int
	NetTx    int
}

func (s *Project) AggStats() (*StatsDetails, error) {
	var agg StatsDetails
	n := len(*s.Services)

	for _, v := range *s.Services {
		stats, err := v.AggStats()
		if err != nil {
			continue
		}

		agg.CPU += stats.CPU / uint64(n)
		agg.Mem += int64(stats.Mem) / int64(n)
		agg.MemLimit += int64(stats.MemLimit) / int64(n)
		agg.NetRx += stats.NetRx / n
		agg.NetTx += stats.NetTx / n
	}

	return &agg, nil
}

func (s *Service) AggStats() (*StatsDetails, error) {
	var agg StatsDetails
	n := len(*s.Containers)

	for _, v := range *s.Containers {
		stats, err := v.AggStats()
		if err != nil {
			continue
		}

		netRx, netTx := 0, 0
		for _, net := range stats.Networks {
			netRx += int(net.RxBytes)
			netTx += int(net.TxBytes)
		}

		agg.CPU += stats.CPUStats.CPUUsage.TotalUsage / uint64(n)
		agg.Mem += int64(stats.MemoryStats.Usage) / int64(n)
		agg.MemLimit += int64(stats.MemoryStats.Limit) / int64(n)
		agg.NetRx += netRx / n
		agg.NetTx += netTx / n
	}

	return &agg, nil
}

func (s *ContainerDetails) AggStats() (*container.StatsResponse, error) {
	var t container.StatsResponse
	err := s.Decoder.Decode(&t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (s *Project) NumContainers() int {
	sum := 0
	for _, v := range *s.Services {
		sum += len(*v.Containers)
	}

	return sum
}

func (s *Project) LatestStartedAt() (time.Time, error) {
	var latest time.Time

	for _, c := range *s.Services {
		t, err := c.LatestStartedAt()
		if err != nil {
			return time.Time{}, err
		}

		if t.After(latest) {
			latest = t
		}
	}

	return latest, nil
}

func (s *Service) LatestStartedAt() (time.Time, error) {
	var latest time.Time

	for _, c := range *s.Containers {
		// Parse RFC3339 string from container state
		if c.Inspect.State == nil || c.Inspect.State.StartedAt == "" {
			continue
		}

		started, err := time.Parse(time.RFC3339Nano, c.Inspect.State.StartedAt)
		if err != nil {
			return time.Time{}, err
		}

		if started.After(latest) {
			latest = started
		}
	}

	return latest, nil
}

func (s *Project) Status() string {
	running, exited, restarting, paused := 0, 0, 0, 0

	for _, c := range *s.Services {
		switch c.Status() {
		case "running":
			running++
		case "exited":
			exited++
		case "restarting":
			restarting++
		case "paused":
			paused++
		}
	}

	total := len(*s.Services)

	switch {
	case running == total:
		return "running"
	case exited == total:
		return "exited"
	case restarting > 0:
		return fmt.Sprintf("restarting (%d/%d)", running, total)
	case paused == total:
		return "paused"
	case running > 0:
		return fmt.Sprintf("degraded (%d/%d running)", running, total)
	default:
		return "unknown"
	}
}

func (s *Service) Status() string {
	if len(*s.Containers) == 0 {
		return "No Containers"
	}

	running, exited, restarting, paused := 0, 0, 0, 0

	for _, c := range *s.Containers {
		switch {
		case c.Inspect.State.Running:
			running++
		case c.Inspect.State.Dead:
			exited++
		case c.Inspect.State.Restarting:
			restarting++
		case c.Inspect.State.Paused:
			paused++
		}
	}

	total := len(*s.Containers)

	switch {
	case running == total:
		return "running"
	case exited == total:
		return "exited"
	case restarting > 0:
		return fmt.Sprintf("restarting (%d/%d)", running, total)
	case paused == total:
		return "paused"
	case running > 0:
		return fmt.Sprintf("degraded (%d/%d running)", running, total)
	default:
		return "unknown"
	}
}
