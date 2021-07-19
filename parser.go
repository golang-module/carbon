package carbon

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Parse 将标准格式时间字符串解析成 Carbon 实例
func (c Carbon) Parse(value string) Carbon {
	layout := DateTimeFormat
	if value == "" || value == "0" || value == "0000-00-00 00:00:00" || value == "0000-00-00" || value == "00:00:00" {
		return c
	}
	if len(value) == 10 && strings.Count(value, "-") == 2 {
		layout = DateFormat
	}
	if strings.Index(value, "T") == 10 {
		layout = RFC3339Format
	}
	if _, err := strconv.ParseInt(value, 10, 64); err == nil {
		switch len(value) {
		case 8:
			layout = ShortDateFormat
		case 14:
			layout = ShortDateTimeFormat
		}
	}
	return c.ParseByLayout(value, layout)
}

// Parse 将标准格式时间字符串解析成 Carbon 实例(默认时区)
func Parse(value string) Carbon {
	return NewCarbon().Parse(value)
}

// ParseByFormat 将特殊格式时间字符串解析成 Carbon 实例
func (c Carbon) ParseByFormat(value string, format string) Carbon {
	if value == "" || value == "0" || value == "0000-00-00 00:00:00" || value == "0000-00-00" || value == "00:00:00" {
		return c
	}
	layout := format2layout(format)
	return c.ParseByLayout(value, layout)
}

// ParseByFormat 将特殊格式时间字符串解析成 Carbon 实例(默认时区)
func ParseByFormat(value string, format string) Carbon {
	return NewCarbon().ParseByFormat(value, format)
}

// ParseByLayout 将布局时间字符串解析成 Carbon 实例
func (c Carbon) ParseByLayout(value string, layout string) Carbon {
	if value == "" || value == "0" || value == "0000-00-00 00:00:00" || value == "0000-00-00" || value == "00:00:00" {
		return c
	}
	tt, err := time.ParseInLocation(layout, value, c.Loc)
	if err != nil {
		c.Error = errors.New(fmt.Sprintf("the value %q and the layout %q don't match, so the value can't parse to carbon", value, layout))
		return c
	}
	c.Time = tt
	return c
}

// ParseByLayout 将布局时间字符串解析成 Carbon 实例(默认时区)
func ParseByLayout(value string, layout string) Carbon {
	return NewCarbon().ParseByLayout(value, layout)
}
