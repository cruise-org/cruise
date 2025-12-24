package compose

import (
	"fmt"

	internaltypes "github.com/NucleoFusion/cruise/internal/types"
)

func ServiceHeaders(width int) string {
	w := width / 7
	return fmt.Sprintf("%-*s %-*s %-*s %-*s %-*s %-*s",
		1*w, "Name",
		1*w, "Image",
		1*w, "Status",
		1*w, "Containers",
		2*w, "Limits (CPU - Mem)",
		1*w, "Last Updated",
	)
}

func ServiceFormatted(width int, s *internaltypes.Project, srv string) string {
	w := width / 7
	service, ok := (*s.Services)[srv]
	if !ok {
		return ""
	}

	inspectedService, ok := s.Inspect.Services[srv]
	if !ok {
		return ""
	}

	lim := inspectedService.Deploy.Resources.Limits
	limits := fmt.Sprintf("%.2f nCPUs - %d bytes", lim.NanoCPUs.Value(), int64(lim.MemoryBytes))

	startedAt := "NA"
	t, err := service.LatestStartedAt()
	if err == nil { // If error, returns NA
		startedAt = t.Format("15:04 02 Jan")
	}

	return fmt.Sprintf("%-*s %-*s %-*s %-*d %-*s %-*s",
		1*w, srv,
		1*w, s.Inspect.Services[srv].Image,
		1*w, service.Status(),
		1*w, len(*service.Containers),
		2*w, limits,
		1*w, startedAt,
	)
}
