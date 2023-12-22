package date

import (
	"time"
)

const (
	dateLayout = "2006-01-02" // YYYY-MM-DDTHH:MM:SS
)

func GetNow() time.Time {
	return time.Now()
}

func GetCurrentDate() string {
	return GetNow().Format(dateLayout)
}
