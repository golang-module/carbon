package carbon

import "time"

// StartOfYear 本年开始时间
func (c Carbon) StartOfYear() Carbon {
	c.Time = time.Date(c.Year(), 1, 1, 0, 0, 0, 0, c.Loc)
	return c
}

// EndOfYear 本年结束时间
func (c Carbon) EndOfYear() Carbon {
	c.Time = time.Date(c.Year(), 12, 31, 23, 59, 59, 999999999, c.Loc)
	return c
}

// StartOfMonth 本月开始时间
func (c Carbon) StartOfMonth() Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), 1, 0, 0, 0, 0, c.Loc)
	return c
}

// EndOfMonth 本月结束时间
func (c Carbon) EndOfMonth() Carbon {
	t := time.Date(c.Year(), time.Month(c.Month()), 1, 23, 59, 59, 999999999, c.Loc)
	c.Time = t.AddDate(0, 1, -1)
	return c
}

// StartOfWeek 本周开始时间
func (c Carbon) StartOfWeek() Carbon {
	days := c.Week()
	if days == 0 {
		days = DaysPerWeek
	}
	t := time.Date(c.Year(), time.Month(c.Month()), c.Day(), 0, 0, 0, 0, c.Loc)
	c.Time = t.AddDate(0, 0, 1-days)
	return c
}

// EndOfWeek 本周结束时间
func (c Carbon) EndOfWeek() Carbon {
	days := c.Week()
	if days == 0 {
		days = DaysPerWeek
	}
	t := time.Date(c.Year(), time.Month(c.Month()), c.Day(), 23, 59, 59, 999999999, c.Loc)
	c.Time = t.AddDate(0, 0, DaysPerWeek-days)
	return c
}

// StartOfDay 本日开始时间
func (c Carbon) StartOfDay() Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), 0, 0, 0, 0, c.Loc)
	return c
}

// EndOfDay 本日结束时间
func (c Carbon) EndOfDay() Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), 23, 59, 59, 999999999, c.Loc)
	return c
}

// StartOfHour 小时开始时间
func (c Carbon) StartOfHour() Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), 0, 0, 0, c.Loc)
	return c
}

// EndOfHour 小时结束时间
func (c Carbon) EndOfHour() Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), 59, 59, 999999999, c.Loc)
	return c
}

// StartOfMinute 分钟开始时间
func (c Carbon) StartOfMinute() Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), 0, 0, c.Loc)
	return c
}

// EndOfMinute 分钟结束时间
func (c Carbon) EndOfMinute() Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), 59, 999999999, c.Loc)
	return c
}

// StartOfSecond 秒开始时间
func (c Carbon) StartOfSecond() Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), 0, c.Loc)
	return c
}

// EndOfSecond 秒结束时间
func (c Carbon) EndOfSecond() Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), 999999999, c.Loc)
	return c
}
