package carbon

import (
	"strings"
	"time"
)

// format转layout
func format2layout(format string) string {
	format = strings.Trim(format, " ")
	layout := strings.Replace(format, "Y", "2006", 1)
	layout = strings.Replace(layout, "y", "06", 1)
	layout = strings.Replace(layout, "M", "Jan", 1)
	layout = strings.Replace(layout, "m", "01", 1)
	layout = strings.Replace(layout, "F", "January", 1)
	layout = strings.Replace(layout, "n", "1", 1)
	layout = strings.Replace(layout, "l", "Monday", 1)
	layout = strings.Replace(layout, "D", "Mon", 1)
	layout = strings.Replace(layout, "d", "02", 1)
	layout = strings.Replace(layout, "j", "2", 1)
	layout = strings.Replace(layout, "H", "15", 1)
	layout = strings.Replace(layout, "h", "03", 1)
	layout = strings.Replace(layout, "i", "04", 1)
	layout = strings.Replace(layout, "s", "05", 1)
	layout = strings.Replace(layout, "P", "PM", 1)
	layout = strings.Replace(layout, "p", "pm", 1)
	return layout
}

// newCarbon 创建一个新Carbon实例
func newCarbon(t time.Time) Carbon {
	loc, err := time.LoadLocation(Local)
	if err != nil {
		panic("invalid timezone value: " + Local)
	}
	return Carbon{Time: t, loc: loc}
}

// getLocalByTimezone 通过时区获取Location实例
func getLocalByTimezone(timezone string) *time.Location {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		panic("invalid timezone value: " + timezone)
	}
	return loc
}
