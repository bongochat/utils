package date

import (
	"fmt"
	"time"

	"github.com/bongochat/utils/resterrors"
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

// DateFormat takes a date string in any common format and returns it in DD-MM-YYYY format.
func DateFormat(dateStr string) (string, resterrors.RestError) {
	// List of possible date formats
	formats := []string{
		"2006-01-02",      // YYYY-MM-DD
		"02-01-2006",      // DD-MM-YYYY
		"01/02/2006",      // MM/DD/YYYY
		"02/01/2006",      // DD/MM/YYYY
		"2006/01/02",      // YYYY/MM/DD
		"January 2, 2006", // Month Day, Year
		"2 January 2006",  // Day Month Year
		"02 Jan 2006",     // DD Mon YYYY
		"Jan 2, 2006",     // Mon Day, Year
		time.RFC3339,      // RFC3339
		time.RFC1123,      // RFC1123
		time.RFC1123Z,     // RFC1123Z
		time.RFC850,       // RFC850
		time.ANSIC,        // ANSIC
		time.UnixDate,     // UnixDate
		time.RubyDate,     // RubyDate
	}

	var parsedDate time.Time
	var err error

	// Try parsing the date string with each format
	for _, format := range formats {
		parsedDate, err = time.Parse(format, dateStr)
		if err == nil {
			break
		}
	}

	if err != nil {
		return "", resterrors.NewInternalServerError("Invalid date format", "", err)
	}

	// Return the date in DD-MM-YYYY format
	return parsedDate.Format("02-01-2006"), nil
}

// DateFormatForDB takes a date string in DD-MM-YYYY format and returns it in YYYY-MM-DD format.
func DateFormatForDB(dateStr string) (string, resterrors.RestError) {
	// Parse the date using the DD-MM-YYYY format
	parsedDate, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		return "", resterrors.NewInternalServerError("Invalid date format", "", err)
	}

	// Return the date in YYYY-MM-DD format
	return parsedDate.Format("2006-01-02"), nil
}
