package carbon

import "time"

// Timezone 设置时区
func (c Carbon) SetTimezone(name string) Carbon {
	loc, err := getLocationByTimezone(name)
	c.Loc = loc
	c.Error = err
	return c
}

// Timezone 设置时区
func SetTimezone(name string) Carbon {
	loc, err := getLocationByTimezone(name)
	return Carbon{Loc: loc, Lang: NewLanguage(), Error: err}
}

// SetLanguage 设置语言对象
func (c Carbon) SetLanguage(lang *Language) Carbon {
	if len(c.Lang.resources) == 0 {
		c.Lang = lang
		return c
	}
	for k, v := range lang.resources {
		c.Lang.resources[k] = v
	}
	return c
}

// SetLanguage 设置语言对象
func SetLanguage(lang *Language) Carbon {
	err := lang.SetLocale(lang.locale)
	return Carbon{Loc: time.Local, Lang: lang, Error: err}
}

// SetLocale 设置语言区域
func (c Carbon) SetLocale(locale string) Carbon {
	lang := NewLanguage()
	c.Error = lang.SetLocale(locale)
	c.Lang = lang
	return c
}

// SetLocale 设置语言区域
func SetLocale(locale string) Carbon {
	lang := NewLanguage()
	err := lang.SetLocale(locale)
	return Carbon{Loc: time.Local, Lang: lang, Error: err}
}

// SetYear 设置年
func (c Carbon) SetYear(year int) Carbon {
	c.Time = time.Date(year, time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Loc)
	return c
}

// SetMonth 设置月
func (c Carbon) SetMonth(month int) Carbon {
	c.Time = time.Date(c.Year(), time.Month(month), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Loc)
	return c
}

// SetDay 设置日
func (c Carbon) SetDay(day int) Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), day, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Loc)
	return c
}

// SetHour 设置时
func (c Carbon) SetHour(hour int) Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), hour, c.Minute(), c.Second(), c.Nanosecond(), c.Loc)
	return c
}

// SetMinute 设置分
func (c Carbon) SetMinute(minute int) Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), minute, c.Second(), c.Nanosecond(), c.Loc)
	return c
}

// SetSecond 设置秒
func (c Carbon) SetSecond(second int) Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), second, c.Nanosecond(), c.Loc)
	return c
}

// SetMillisecond 设置毫秒
func (c Carbon) SetMillisecond(millisecond int) Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), millisecond*1e6, c.Loc)
	return c
}

// SetMicrosecond 设置微秒
func (c Carbon) SetMicrosecond(microsecond int) Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), microsecond*1e9, c.Loc)
	return c
}

// SetNanosecond 设置纳秒
func (c Carbon) SetNanosecond(nanosecond int) Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), nanosecond, c.Loc)
	return c
}
