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
	loc, _ := time.LoadLocation(Local)
	return Carbon{Time: t, loc: loc}
}

// getLocalByTimezone 通过时区获取Location实例
func getLocalByTimezone(timezone string) *time.Location {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		panic("invalid timezone \"" + timezone + "\", for all valid timezone, please see the $GOROOT/lib/time/zoneinfo.zip file")
	}
	return loc
}

// parseByLayout 通过布局模板解析
func parseByLayout(value string, layout string) time.Time {
	t, err := time.ParseInLocation(layout, value, getLocalByTimezone(Local))
	if err != nil {
		panic("The value \"" + value + "\" and the layout \"" + layout + "\" don't match")
	}
	return t
}

// parseByDuration 通过持续时间解析
func parseByDuration(duration string) time.Duration {
	d, err := time.ParseDuration(duration)
	if err != nil {
		panic("invalid duration \"" + duration + "\"")
	}
	return d
}
