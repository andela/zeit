package utility

import (
	humanize "github.com/dustin/go-humanize"
	"time"
    "strings"
    "math"
    "fmt"
)

var JavascriptISOString = "2006-01-02T15:04:05.999Z07:00"

func FormatToBasicTime(input string) string {
	t, _ := time.Parse(JavascriptISOString, input)
	return t.Format("15:04")
}

func FormatToDateTime(input string) string {
	t, _ := time.Parse(JavascriptISOString, input)
	return humanize.Time(t)
}

func FormatToCalendarDate(input time.Time) string {
    return input.Format("01 02 2006")
}

func FormatToDuration(input string) string {
    args := strings.Split(input, ".")
    return args[0] + " hours and " + args[1] + " minutes"
}

func EqualDates(t1, t2 time.Time) bool {
    timeString1 := FormatToCalendarDate(t1)
    timeString2 := FormatToCalendarDate(t2)
    return timeString1 == timeString2
}

func FormatToHourMinutes(duration time.Duration) string {
    minutes := math.Mod(duration.Minutes(), 60)
    hours := duration.Minutes() / 60
    return fmt.Sprintf("%d hours %d minutes", int(hours), int(minutes))
}