package carbon

import (
	"errors"
	"strings"
	"time"
)

// format转layout
func format2layout(format string) string {
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

// getLocationByTimezone 通过时区获取Location实例
func getLocationByTimezone(timezone string) (*time.Location, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		err = errors.New("invalid timezone \"" + timezone + "\", please see the $GOROOT/lib/time/zoneinfo.zip file for all valid timezone")
	}
	return loc, err
}

// parseByLayout 通过布局模板解析
func parseByLayout(value string, layout string) (time.Time, error) {
	loc, _ := time.LoadLocation(Local)
	tt, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		err = errors.New("the value \"" + value + "\" and layout \"" + layout + "\" don't match")
	}
	return tt, err
}

// parseByDuration 通过持续时间解析
func parseByDuration(duration string) (time.Duration, error) {
	td, err := time.ParseDuration(duration)
	if err != nil {
		err = errors.New("invalid duration \"" + duration + "\"")
	}
	return td, err
}
