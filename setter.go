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
	return Carbon{Loc: loc, Error: err, Lang: NewLanguage()}
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
	loc, _ := getLocationByTimezone(Local)
	err := lang.SetLocale(lang.locale)
	return Carbon{Loc: loc, Lang: lang, Error: err}
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
	return Carbon{Lang: lang, Error: err}
}

// SetYear 设置年
func (c Carbon) SetYear(year int) Carbon {
	c.Time = time.Date(year, c.Time.Month(), c.Time.Day(), c.Time.Hour(), c.Time.Minute(), c.Time.Second(), c.Time.Nanosecond(), c.Loc)
	return c
}

// SetMonth 设置月
func (c Carbon) SetMonth(month int) Carbon {
	c.Time = time.Date(c.Time.Year(), time.Month(month), c.Time.Day(), c.Time.Hour(), c.Time.Minute(), c.Time.Second(), c.Time.Nanosecond(), c.Loc)
	return c
}

// SetDay 设置日
func (c Carbon) SetDay(day int) Carbon {
	c.Time = time.Date(c.Time.Year(), c.Time.Month(), day, c.Time.Hour(), c.Time.Minute(), c.Time.Second(), c.Time.Nanosecond(), c.Loc)
	return c
}

// SetHour 设置时
func (c Carbon) SetHour(hour int) Carbon {
	c.Time = time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), hour, c.Time.Minute(), c.Time.Second(), c.Time.Nanosecond(), c.Loc)
	return c
}

// SetMinute 设置分
func (c Carbon) SetMinute(minute int) Carbon {
	c.Time = time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), c.Time.Hour(), minute, c.Time.Second(), c.Time.Nanosecond(), c.Loc)
	return c
}

// SetSecond 设置秒
func (c Carbon) SetSecond(second int) Carbon {
	c.Time = time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), c.Time.Hour(), c.Time.Minute(), second, c.Time.Nanosecond(), c.Loc)
	return c
}
