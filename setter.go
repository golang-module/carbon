package carbon

import (
	"time"
)

// SetTimezone sets timezone.
// 设置时区
func (c Carbon) SetTimezone(name string) Carbon {
	if c.Error != nil {
		return c
	}
	loc, err := getLocationByTimezone(name)
	c.Loc = loc
	c.Error = err
	return c
}

// SetTimezone sets timezone.
// 设置时区
func SetTimezone(name string) Carbon {
	return NewCarbon().SetTimezone(name)
}

// SetLocale sets locale.
// 设置语言区域
func (c Carbon) SetLocale(locale string) Carbon {
	if c.Error != nil {
		return c
	}
	c.Error = c.Lang.SetLocale(locale)
	return c
}

// SetLocale sets locale.
// 设置语言区域
func SetLocale(locale string) Carbon {
	c := NewCarbon()
	c.Error = c.Lang.SetLocale(locale)
	return c
}

// SetLanguage sets language.
// 设置语言对象
func (c Carbon) SetLanguage(lang *Language) Carbon {
	if c.Error != nil {
		return c
	}
	c.Lang = lang
	return c
}

// SetLanguage sets language.
// 设置语言对象
func SetLanguage(lang *Language) Carbon {
	c, err := NewCarbon(), lang.SetLocale(lang.locale)
	c.Lang = lang
	c.Error = err
	return c
}

// SetYear sets year.
// 设置年份
func (c Carbon) SetYear(year int) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.Time = time.Date(year, time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Loc)
	return c
}

// SetYearNoOverflow sets year without overflowing month.
// 设置年份(月份不溢出)
func (c Carbon) SetYearNoOverflow(year int) Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.AddYearsNoOverflow(year - c.Year())
}

// SetMonth sets month.
// 设置月份
func (c Carbon) SetMonth(month int) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(month), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Loc)
	return c
}

// SetMonthNoOverflow sets month without overflowing month.
// 设置月份(月份不溢出)
func (c Carbon) SetMonthNoOverflow(month int) Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.AddMonthsNoOverflow(month - c.Month())
}

// SetDay sets day.
// 设置日期
func (c Carbon) SetDay(day int) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), day, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Loc)
	return c
}

// SetHour sets hour.
// 设置小时
func (c Carbon) SetHour(hour int) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), hour, c.Minute(), c.Second(), c.Nanosecond(), c.Loc)
	return c
}

// SetMinute sets minute.
// 设置分钟
func (c Carbon) SetMinute(minute int) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), minute, c.Second(), c.Nanosecond(), c.Loc)
	return c
}

// SetSecond sets second.
// 设置秒数
func (c Carbon) SetSecond(second int) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), second, c.Nanosecond(), c.Loc)
	return c
}

// SetMillisecond sets millisecond.
// 设置毫秒
func (c Carbon) SetMillisecond(millisecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), millisecond*1e6, c.Loc)
	return c
}

// SetMicrosecond sets microsecond.
// 设置微秒
func (c Carbon) SetMicrosecond(microsecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), microsecond*1e3, c.Loc)
	return c
}

// SetNanosecond sets nanosecond.
// 设置纳秒
func (c Carbon) SetNanosecond(nanosecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), nanosecond, c.Loc)
	return c
}
