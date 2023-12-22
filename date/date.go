package date

import (
	"time"
)

const (
	dateLayout = "2023-12-25T15:04:05Z" // YYYY-MM-DDTHH:MM:SS
)

func GetNow() time.Time {
	return time.Now()
}

func GetCurrentDate() string {
	return GetNow().Format(dateLayout)
}
