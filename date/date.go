package date

import (
	"fmt"
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

func GetNow() time.Time {
	return time.Now()
}

func GetCurrentDate() string {
	return GetNow().Format(apiDateLayout)
}

func TimeAgo(datetimeStr string) string {
	t, err := time.Parse(time.RFC3339, datetimeStr)
	if err != nil {
		fmt.Println("Error parsing datetime:", err)
		return ""
	}

	duration := time.Since(t)
	if duration < 0 {
		duration = -duration
	}

	switch {
	case duration < time.Minute:
		return "just now"
	case duration < time.Hour:
		return fmt.Sprintf("%dmin", int(duration.Minutes()))
	case duration < 24*time.Hour:
		return fmt.Sprintf("%dh", int(duration.Hours()))
	case duration < 7*24*time.Hour:
		return fmt.Sprintf("%dd", int(duration.Hours()/24))
	case duration < 30*24*time.Hour:
		return fmt.Sprintf("%dw", int(duration.Hours()/(24*7)))
	case duration < 365*24*time.Hour:
		return fmt.Sprintf("%dmo", int(duration.Hours()/(24*30)))
	default:
		return fmt.Sprintf("%dy", int(duration.Hours()/(24*365)))
	}
}
