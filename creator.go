package carbon

import (
	"strconv"
	"time"
)

// CreateFromTimestamp 从时间戳创建 Carbon 实例
func (c Carbon) CreateFromTimestamp(timestamp int64, timezone ...string) Carbon {
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	if c.Error != nil {
		return c
	}
	ts, count := timestamp, len(strconv.FormatInt(timestamp, 10))
	if timestamp < 0 {
		count -= 1
	}
	switch count {
	case 10:
		ts = timestamp
	case 13:
		ts = timestamp / 1e3
	case 16:
		ts = timestamp / 1e6
	case 19:
		ts = timestamp / 1e9
	}
	c.Time = time.Unix(ts, 0)
	return c
}

// CreateFromTimestamp 从时间戳创建 Carbon 实例(默认时区)
func CreateFromTimestamp(timestamp int64, timezone ...string) Carbon {
	return NewCarbon().CreateFromTimestamp(timestamp, timezone...)
}

// CreateFromDateTime 从年月日时分秒创建 Carbon 实例
func (c Carbon) CreateFromDateTime(year int, month int, day int, hour int, minute int, second int, timezone ...string) Carbon {
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	if c.Error != nil {
		return c
	}
	c.Time = time.Date(year, time.Month(month), day, hour, minute, second, time.Now().Nanosecond(), c.Loc)
	return c
}

// CreateFromDateTime 从年月日时分秒创建 Carbon 实例
func CreateFromDateTime(year int, month int, day int, hour int, minute int, second int, timezone ...string) Carbon {
	return NewCarbon().CreateFromDateTime(year, month, day, hour, minute, second, timezone...)
}

// CreateFromDate 从年月日创建 Carbon 实例
func (c Carbon) CreateFromDate(year int, month int, day int, timezone ...string) Carbon {
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	if c.Error != nil {
		return c
	}
	hour, minute, second := time.Now().In(c.Loc).Clock()
	c.Time = time.Date(year, time.Month(month), day, hour, minute, second, time.Now().Nanosecond(), c.Loc)
	return c
}

// CreateFromDate 从年月日创建 Carbon 实例
func CreateFromDate(year int, month int, day int, timezone ...string) Carbon {
	return NewCarbon().CreateFromDate(year, month, day, timezone...)
}

// CreateFromTime 从时分秒创建 Carbon 实例
func (c Carbon) CreateFromTime(hour int, minute int, second int, timezone ...string) Carbon {
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	if c.Error != nil {
		return c
	}
	year, month, day := time.Now().In(c.Loc).Date()
	c.Time = time.Date(year, month, day, hour, minute, second, time.Now().Nanosecond(), c.Loc)
	return c
}

// CreateFromTime 从时分秒创建 Carbon 实例
func CreateFromTime(hour int, minute int, second int, timezone ...string) Carbon {
	return NewCarbon().CreateFromTime(hour, minute, second, timezone...)
}
