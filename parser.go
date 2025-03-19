package carbon

import (
	"strconv"
	"time"
)

// Parse parses a standard time string as a Carbon instance.
// 将标准格式时间字符串解析成 Carbon 实例
func Parse(value string, timezone ...string) *Carbon {
	if len(value) == 0 {
		return nil
	}
	c := NewCarbon()
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[0])
	}
	switch value {
	case "now":
		return Now(timezone...)
	case "yesterday":
		return Yesterday(timezone...)
	case "tomorrow":
		return Tomorrow(timezone...)
	}
	for _, layout := range layouts {
		t, err := time.ParseInLocation(layout, value, c.loc)
		if err == nil {
			c.time = t
			c.layout = layout
			return c
		}
	}
	c.Error = invalidValueError(value)
	return c
}

// ParseByFormat parses a time string as a Carbon instance by format.
// 通过格式模板将时间字符串解析成 Carbon 实例
func ParseByFormat(value, format string, timezone ...string) *Carbon {
	if len(value) == 0 {
		return nil
	}
	carbon := ParseByLayout(value, format2layout(format), timezone...)
	if carbon.Error != nil {
		carbon.Error = invalidFormatError(value, format)
	}
	return carbon
}

// ParseByLayout parses a time string as a Carbon instance by layout.
// 通过布局模板将时间字符串解析成 Carbon 实例
func ParseByLayout(value, layout string, timezone ...string) *Carbon {
	if len(value) == 0 {
		return nil
	}
	c := NewCarbon()
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[0])
	}
	if len(layout) == 0 {
		layout = DefaultLayout
	}
	if layout == "timestamp" {
		timestamp, _ := strconv.ParseInt(value, 10, 64)
		return CreateFromTimestamp(timestamp, c.Timezone())
	}
	if layout == "timestampMilli" {
		timestamp, _ := strconv.ParseInt(value, 10, 64)
		return CreateFromTimestampMilli(timestamp, c.Timezone())
	}
	if layout == "timestampMicro" {
		timestamp, _ := strconv.ParseInt(value, 10, 64)
		return CreateFromTimestampMicro(timestamp, c.Timezone())
	}
	if layout == "timestampNano" {
		timestamp, _ := strconv.ParseInt(value, 10, 64)
		return CreateFromTimestampNano(timestamp, c.Timezone())
	}
	tt, err := time.ParseInLocation(layout, value, c.loc)
	if err != nil {
		c.Error = invalidLayoutError(value, layout)
		return c
	}
	c.time = tt
	c.layout = layout
	return c
}
