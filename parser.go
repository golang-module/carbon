package carbon

import (
	"strconv"
	"strings"
	"time"
)

// Parse parse as a Carbon instance by a standard string
// 将标准格式时间字符串解析成 Carbon 实例
func (c Carbon) Parse(value string, timezone ...string) Carbon {
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	if c.Error != nil {
		return c
	}
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
	if c.ParseByLayout(value, layout).Error != nil {
		c.Error = invalidValueError(value)
		return c
	}
	return c.ParseByLayout(value, layout)
}

// Parse parse as a Carbon instance by a standard string
// 将标准时间字符串解析成 Carbon 实例
func Parse(value string, timezone ...string) Carbon {
	return NewCarbon().Parse(value, timezone...)
}

// ParseByFormat parse as a Carbon instance by a format string
// 通过格式化字符将字符串解析成 carbon 实例
func (c Carbon) ParseByFormat(value string, format string, timezone ...string) Carbon {
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	if c.Error != nil {
		return c
	}
	if value == "" || value == "0" || value == "0000-00-00 00:00:00" || value == "0000-00-00" || value == "00:00:00" {
		return c
	}
	layout := format2layout(format)
	if c.ParseByLayout(value, layout).Error != nil {
		c.Error = invalidFormatError(value, format)
		return c
	}
	return c.ParseByLayout(value, layout)
}

// ParseByFormat parse as a Carbon instance by a format string
// 通过布局字符将字符串解析成 carbon 实例
func ParseByFormat(value string, format string, timezone ...string) Carbon {
	return NewCarbon().ParseByFormat(value, format, timezone...)
}

// ParseByLayout parse as a Carbon instance by a layout string
// 通过布局字符将字符串解析成 carbon 实例
func (c Carbon) ParseByLayout(value string, layout string, timezone ...string) Carbon {
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	if c.Error != nil {
		return c
	}
	if value == "" || value == "0" || value == "0000-00-00 00:00:00" || value == "0000-00-00" || value == "00:00:00" {
		return c
	}
	tt, err := time.ParseInLocation(layout, value, c.Loc)
	if err != nil {
		c.Error = invalidLayoutError(value, layout)
		return c
	}
	c.Time = tt
	return c
}

// ParseByLayout parse as a Carbon instance by a layout string
// 将布局时间字符串解析成 Carbon 实例
func ParseByLayout(value string, layout string, timezone ...string) Carbon {
	return NewCarbon().ParseByLayout(value, layout, timezone...)
}
