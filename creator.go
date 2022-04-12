package carbon

import (
	"time"
)

// CreateFromTimestamp creates a Carbon instance from a given timestamp with second.
// 从给定的秒级时间戳创建 Carbon 实例
func (c Carbon) CreateFromTimestamp(timestamp int64, timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	c.time = time.Unix(timestamp, 0)
	return c
}

// CreateFromTimestamp creates a Carbon instance from a given timestamp with second.
// 从给定的秒级时间戳创建 Carbon 实例
func CreateFromTimestamp(timestamp int64, timezone ...string) Carbon {
	return NewCarbon().CreateFromTimestamp(timestamp, timezone...)
}

// CreateFromTimestampMilli creates a Carbon instance from a given timestamp with millisecond.
// 从给定的微秒级时间戳创建 Carbon 实例
func (c Carbon) CreateFromTimestampMilli(timestamp int64, timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	c.time = time.Unix(timestamp/1e3, (timestamp%1e3)*1e6)
	return c
}

// CreateFromTimestampMilli creates a Carbon instance from a given timestamp with millisecond.
// 从给定的微秒级时间戳创建 Carbon 实例
func CreateFromTimestampMilli(timestamp int64, timezone ...string) Carbon {
	return NewCarbon().CreateFromTimestampMilli(timestamp, timezone...)
}

// CreateFromTimestampMicro creates a Carbon instance from a given timestamp with microsecond.
// 从给定的微秒级时间戳创建 Carbon 实例
func (c Carbon) CreateFromTimestampMicro(timestamp int64, timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	c.time = time.Unix(timestamp/1e6, (timestamp%1e6)*1e3)
	return c
}

// CreateFromTimestampMicro creates a Carbon instance from a given timestamp with microsecond.
// 从给定的微秒级时间戳创建 Carbon 实例
func CreateFromTimestampMicro(timestamp int64, timezone ...string) Carbon {
	return NewCarbon().CreateFromTimestampMicro(timestamp, timezone...)
}

// CreateFromTimestampNano creates a Carbon instance from a given timestamp with nanosecond.
// 从给定的纳秒级时间戳创建 Carbon 实例
func (c Carbon) CreateFromTimestampNano(timestamp int64, timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	c.time = time.Unix(timestamp/1e9, timestamp%1e9)
	return c
}

// CreateFromTimestampNano creates a Carbon instance from a given timestamp with nanosecond.
// 从给定的纳秒级时间戳创建 Carbon 实例
func CreateFromTimestampNano(timestamp int64, timezone ...string) Carbon {
	return NewCarbon().CreateFromTimestampNano(timestamp, timezone...)
}

// CreateFromDateTime creates a Carbon instance from a given date and time.
// 从给定的年月日时分秒创建 Carbon 实例
func (c Carbon) CreateFromDateTime(year, month, day, hour, minute, second int, timezone ...string) Carbon {
	return c.create(year, month, day, hour, minute, second, time.Now().Nanosecond(), timezone...)
}

// CreateFromDateTime creates a Carbon instance from a given date and time.
// 从给定的年月日时分秒创建 Carbon 实例
func CreateFromDateTime(year, month, day, hour, minute, second int, timezone ...string) Carbon {
	return NewCarbon().CreateFromDateTime(year, month, day, hour, minute, second, timezone...)
}

// CreateFromDateTimeMilli creates a Carbon instance from a given date and time with millisecond.
// 从给定的年月日时分秒创建 Carbon 实例，包含毫秒
func (c Carbon) CreateFromDateTimeMilli(year, month, day, hour, minute, second, millisecond int, timezone ...string) Carbon {
	return c.create(year, month, day, hour, minute, second, millisecond*1e6, timezone...)
}

// CreateFromDateTimeMilli creates a Carbon instance from a given date and time with millisecond.
// 从给定的年月日时分秒创建 Carbon 实例，包含毫秒
func CreateFromDateTimeMilli(year, month, day, hour, minute, second, millisecond int, timezone ...string) Carbon {
	return NewCarbon().CreateFromDateTimeMilli(year, month, day, hour, minute, second, millisecond, timezone...)
}

// CreateFromDateTimeMicro creates a Carbon instance from a given date and time with microsecond.
// 从给定的年月日时分秒创建 Carbon 实例，包含微秒
func (c Carbon) CreateFromDateTimeMicro(year, month, day, hour, minute, second, microsecond int, timezone ...string) Carbon {
	return c.create(year, month, day, hour, minute, second, microsecond*1e3, timezone...)
}

// CreateFromDateTimeMicro creates a Carbon instance from a given date and time with microsecond.
// 从给定的年月日时分秒创建 Carbon 实例，包含微秒
func CreateFromDateTimeMicro(year, month, day, hour, minute, second, microsecond int, timezone ...string) Carbon {
	return NewCarbon().CreateFromDateTimeMicro(year, month, day, hour, minute, second, microsecond, timezone...)
}

// CreateFromDateTimeNano creates a Carbon instance from a given date and time with nanosecond.
// 从给定的年月日时分秒创建 Carbon 实例，包含纳秒
func (c Carbon) CreateFromDateTimeNano(year, month, day, hour, minute, second, nanosecond int, timezone ...string) Carbon {
	return c.create(year, month, day, hour, minute, second, nanosecond, timezone...)
}

// CreateFromDateTimeNano creates a Carbon instance from a given date and time with nanosecond.
// 从给定的年月日时分秒创建 Carbon 实例，包含纳秒
func CreateFromDateTimeNano(year, month, day, hour, minute, second, nanosecond int, timezone ...string) Carbon {
	return NewCarbon().CreateFromDateTimeNano(year, month, day, hour, minute, second, nanosecond, timezone...)
}

// CreateFromDate creates a Carbon instance from a given date.
// 从给定的年月日创建 Carbon 实例
func (c Carbon) CreateFromDate(year, month, day int, timezone ...string) Carbon {
	now := Now()
	hour, minute, second := now.Time()
	return c.create(year, month, day, hour, minute, second, time.Now().Nanosecond(), timezone...)
}

// CreateFromDate creates a Carbon instance from a given date.
// 从给定的年月日创建 Carbon 实例
func CreateFromDate(year, month, day int, timezone ...string) Carbon {
	return NewCarbon().CreateFromDate(year, month, day, timezone...)
}

// CreateFromTime creates a Carbon instance from a given time.
// 从给定的时分秒创建 Carbon 实例
func (c Carbon) CreateFromTime(hour, minute, second int, timezone ...string) Carbon {
	now := Now()
	year, month, day := now.Date()
	return c.create(year, month, day, hour, minute, second, now.Nanosecond(), timezone...)
}

// CreateFromTime creates a Carbon instance from a given time.
// 从给定的时分秒创建 Carbon 实例
func CreateFromTime(hour, minute, second int, timezone ...string) Carbon {
	return NewCarbon().CreateFromTime(hour, minute, second, timezone...)
}

// create a Carbon instance from a given time.
// 从给定的时间创建 Carbon 实例
func (c Carbon) create(year, month, day, hour, minute, second, nanosecond int, timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, c.loc)
	return c
}
