package carbon

import (
	"time"
)

// Parse parses a standard time string as a Carbon instance.
// 将标准格式时间字符串解析成 Carbon 实例
func (c Carbon) Parse(value string, timezone ...string) Carbon {
	timeLayouts := []string{
		DayDateTimeLayout,
		DateTimeLayout,
		DateTimeNanoLayout,
		ShortDateTimeLayout,
		ShortDateTimeNanoLayout,
		DateLayout,
		DateNanoLayout,
		ShortDateLayout,
		ShortDateNanoLayout,
		TimeLayout,
		TimeNanoLayout,
		ShortTimeLayout,
		ShortTimeNanoLayout,
		ANSICLayout,
		UnixDateLayout,
		RubyDateLayout,
		RFC822Layout,
		RFC822ZLayout,
		RFC850Layout,
		RFC1123Layout,
		RFC1123ZLayout,
		KitchenLayout,
		CookieLayout,
		RFC3339Layout,
		RFC3339NanoLayout,
		ISO8601Layout,
		ISO8601NanoLayout,
		RFC1036Layout,
		RFC7231Layout,
		"2006", "2006-1", "2006-1-2", "2006-1-2 15", "2006-1-2 15:4", "2006-1-2 15:4:5",
		"2006.1.2", "2006.1.2 15", "2006.1.2 15:4", "2006.1.2 15:4:5", "2006.1.2", "2006.1.2 15", "2006.1.2 15:4", "2006.1.2 15:4:5", "2006.1.2 15:4:5.999999999",
		"2006/1", "2006/1/2", "2006/1/2 15", "2006/1/2 15:4", "2006/1/2 15:4:5", "2006/1/2 15:4:5.999999999",
		"1/2/2006", "1/2/2006 15", "1/2/2006 15:4", "1/2/2006 15:4:5", "1/2/2006 15:4:5.999999999",
		"15:4:5 Jan 2, 2006 MST", "2006-1-2 15:4:5.999999999 -0700 MST",
		"2006-1-2T15:4:5-07:00", "2006-1-2T15:4:5.999999999-07:00",
		"2006-1-2T15:4:5Z07", "2006-01-02T15:04:05Z0700", "2006-1-2T15:4:5Z07:00", "2006-1-2T15:4:5.999999999Z07:00",
	}
	if value == "" || value == "0" || value == "0000-00-00 00:00:00" || value == "0000-00-00" || value == "00:00:00" {
		return c
	}
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	for _, layout := range timeLayouts {
		t, err := time.ParseInLocation(layout, value, c.loc)
		if err == nil {
			c.time = t
			return c
		}
	}
	c.Error = invalidValueError(value)
	return c
}

// Parse parses a standard time string as a Carbon instance.
// 将标准时间字符串解析成 Carbon 实例
func Parse(value string, timezone ...string) Carbon {
	return NewCarbon().Parse(value, timezone...)
}

// ParseByFormat parses a time string as a Carbon instance by format.
// 通过格式模板将时间字符串解析成 carbon 实例
func (c Carbon) ParseByFormat(value, format string, timezone ...string) Carbon {
	carbon := c.ParseByLayout(value, format2layout(format), timezone...)
	if carbon.Error != nil {
		carbon.Error = invalidFormatError(value, format)
	}
	return carbon
}

// ParseByFormat parses a time string as a Carbon instance by format.
// 通过格式模板将时间字符串解析成 carbon 实例
func ParseByFormat(value, format string, timezone ...string) Carbon {
	return NewCarbon().ParseByFormat(value, format, timezone...)
}

// ParseByLayout parses a time string as a Carbon instance by layout.
// 通过布局模板将时间字符串解析成 carbon 实例
func (c Carbon) ParseByLayout(value, layout string, timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	if value == "" || value == "0" || value == "0000-00-00 00:00:00" || value == "0000-00-00" || value == "00:00:00" {
		return c
	}
	tt, err := time.ParseInLocation(layout, value, c.loc)
	if err != nil {
		c.Error = invalidLayoutError(value, layout)
		return c
	}
	c.time = tt
	return c
}

// ParseByLayout parses a time string as a Carbon instance by layout.
// 通过布局模板将时间字符串解析成 Carbon 实例
func ParseByLayout(value, layout string, timezone ...string) Carbon {
	return NewCarbon().ParseByLayout(value, layout, timezone...)
}
