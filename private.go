package carbon

import (
	"errors"
	"strings"
	"time"
)

// format转layout
func format2layout(format string) string {
	layout := format

	// 月份中的第几天，有前导零，取值范围 01-31
	if strings.Contains(layout, "d") {
		layout = strings.Replace(format, "d", "02", 1)
	}

	// 星期中的第几天，3 个字母，取值范围 Mon-Sun
	if strings.Contains(layout, "D") {
		layout = strings.Replace(layout, "D", "Mon", 1)
	}

	// 月份中的第几天，没有前导零，取值范围 1-31
	if strings.Contains(layout, "j") {
		layout = strings.Replace(layout, "j", "2", 1)
	}

	// 完整单词表示的星期几，取值范围 Sunday-Saturday
	if strings.Contains(layout, "l") {
		layout = strings.Replace(layout, "l", "Monday", 1)
	}

	// 完整单词表示的月份，取值范围 January-December
	if strings.Contains(layout, "F") {
		layout = strings.Replace(layout, "F", "January", 1)
	}

	// 数字表示的月份，有前导零，取值范围 01-12
	if strings.Contains(layout, "m") {
		layout = strings.Replace(layout, "m", "01", 1)
	}

	// 缩写单词表示的月份，取值范围 Jan-Dec
	if strings.Contains(layout, "M") {
		layout = strings.Replace(layout, "M", "Jan", 1)
	}

	// 数字表示的月份，没有前导零，取值范围 1-12
	if strings.Contains(layout, "n") {
		layout = strings.Replace(layout, "n", "1", 1)
	}

	// 数字表示的简写年份，有前导零，如 21
	if strings.Contains(layout, "y") {
		layout = strings.Replace(layout, "y", "06", 1)
	}

	// 数字表示的完整年份，如 2021
	if strings.Contains(layout, "Y") {
		layout = strings.Replace(layout, "Y", "2006", 1)
	}

	// 小写的上午和下午缩写字母，如 am 或 pm
	if strings.Contains(layout, "a") {
		layout = strings.Replace(layout, "a", "pm", 1)
	}

	// 大写的上午和下午缩写字母，如 AM 或 PM
	if strings.Contains(layout, "A") {
		layout = strings.Replace(layout, "A", "PM", 1)
	}

	// 数字表示的小时，12 小时格式，没有前导零，取值范围 1-12
	if strings.Contains(layout, "g") {
		layout = strings.Replace(layout, "g", "3", 1)
	}

	// 数字表示的小时，12 小时格式，有前导零，取值范围 01-12
	if strings.Contains(layout, "h") {
		layout = strings.Replace(layout, "h", "03", 1)
	}

	// 数字表示的小时，24 小时格式，有前导零，取值范围 00-23
	if strings.Contains(layout, "H") {
		layout = strings.Replace(layout, "H", "15", 1)
	}

	// 数字表示的分钟，有前导零，取值范围 00-59
	if strings.Contains(layout, "i") {
		layout = strings.Replace(layout, "i", "04", 1)
	}

	// 数字表示的秒数，有前导零，取值范围 00-59
	if strings.Contains(layout, "s") {
		layout = strings.Replace(layout, "s", "05", 1)
	}

	// ISO8601 格式的日期，如 2020-08-05T15:19:21+00:00
	if strings.Contains(layout, "c") {
		layout = strings.Replace(layout, "c", "2006-01-02T15:04:05-07:00", 1)
	}

	// RFC822 格式的日期，如 Thu, 21 Dec 2020 16:01:07 +0200
	if strings.Contains(layout, "r") {
		layout = strings.Replace(layout, "r", "Mon, 02 Jan 06 15:04 MST", 1)
	}

	// 与格林威治时间相差的小时数，如 +0200
	if strings.Contains(layout, "O") {
		layout = strings.Replace(layout, "O", "-0700", 1)
	}

	// 与格林威治时间相差的小时数，小时和分钟之间有冒号分隔，如 +02:00
	if strings.Contains(layout, "P") {
		layout = strings.Replace(layout, "P", "-07:00", 1)
	}

	// 本机所在的时区，如 EST，MDT
	if strings.Contains(layout, "T") {
		layout = strings.Replace(layout, "T", "MST", 1)
	}

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
