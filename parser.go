package carbon

import (
	"strconv"
	"strings"
	"time"
)

// Parse parses a standard string as a Carbon instance.
// 将标准格式时间字符串解析成 Carbon 实例
func (c Carbon) Parse(value string, timezone ...string) Carbon {
	layout := DateTimeFormat
	if _, err := strconv.ParseInt(value, 10, 64); err == nil {
		switch {
		case len(value) == 8:
			layout = ShortDateFormat
		case len(value) == 14:
			layout = ShortDateTimeFormat
		}
	} else {
		switch {
		case len(value) == 10 && strings.Count(value, "-") == 2:
			layout = DateFormat
		case len(value) == 18 && strings.Index(value, ".") == 14:
			layout = ShortDateTimeMilliFormat
		case len(value) == 21 && strings.Index(value, ".") == 14:
			layout = ShortDateTimeMicroFormat
		case len(value) == 24 && strings.Index(value, ".") == 14:
			layout = ShortDateTimeNanoFormat
		case len(value) == 25 && strings.Index(value, "T") == 10:
			layout = RFC3339Format
		case len(value) == 29 && strings.Index(value, "T") == 10 && strings.Index(value, ".") == 19:
			layout = RFC3339MilliFormat
		case len(value) == 32 && strings.Index(value, "T") == 10 && strings.Index(value, ".") == 19:
			layout = RFC3339MicroFormat
		case len(value) == 35 && strings.Index(value, "T") == 10 && strings.Index(value, ".") == 19:
			layout = RFC3339NanoFormat
		}
	}
	carbon := c.ParseByLayout(value, layout, timezone...)
	if carbon.Error != nil {
		carbon.Error = invalidValueError(value)
	}
	return carbon
}

// Parse parses a standard string as a Carbon instance.
// 将标准时间字符串解析成 Carbon 实例
func Parse(value string, timezone ...string) Carbon {
	return NewCarbon().Parse(value, timezone...)
}

// ParseByFormat parses a string as a Carbon instance by format.
// 通过格式化字符将字符串解析成 carbon 实例
func (c Carbon) ParseByFormat(value, format string, timezone ...string) Carbon {
	carbon := c.ParseByLayout(value, format2layout(format), timezone...)
	if carbon.Error != nil {
		carbon.Error = invalidFormatError(value, format)
	}
	return carbon
}

// ParseByFormat parses a string as a Carbon instance by format.
// 通过布局字符将字符串解析成 carbon 实例
func ParseByFormat(value, format string, timezone ...string) Carbon {
	return NewCarbon().ParseByFormat(value, format, timezone...)
}

// ParseByLayout parses a string as a Carbon instance by layout.
// 通过布局字符将字符串解析成 carbon 实例
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

// ParseByLayout parses a string as a Carbon instance by layout.
// 将布局时间字符串解析成 Carbon 实例
func ParseByLayout(value, layout string, timezone ...string) Carbon {
	return NewCarbon().ParseByLayout(value, layout, timezone...)
}
