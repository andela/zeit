package utility

import (
	"time"
	humanize "github.com/dustin/go-humanize"
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