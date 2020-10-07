package carbon

import (
	"strings"
	"time"
)

type Carbon struct {
	Time time.Time
	loc  *time.Location
}

// Timezone 设置时区，全局生效
func Timezone(name string) Carbon {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic("invalid timezone value: " + name)
	}
	return Carbon{loc: loc}
}

// Timezone 设置时区，临时生效
func (c Carbon) Timezone(name string) Carbon {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic("invalid timezone value: " + name)
	}
	return Carbon{Time: c.Time.In(loc), loc: loc}
}

// Now 当前
func Now() Carbon {
	return newCarbon(time.Now())
}
func (c Carbon) Now() Carbon {
	return newCarbon(Now().Time.In(c.loc))
}

// Tomorrow 明天
func Tomorrow() Carbon {
	return newCarbon(time.Now().AddDate(0, 0, 1))
}
func (c Carbon) Tomorrow() Carbon {
	return newCarbon(Tomorrow().Time.In(c.loc))
}

// Yesterday 昨天
func Yesterday() Carbon {
	return newCarbon(time.Now().AddDate(0, 0, -1))
}
func (c Carbon) Yesterday() Carbon {
	return newCarbon(Yesterday().Time.In(c.loc))
}

// CreateFromTimestamp 从时间戳创建Carbon实例
func CreateFromTimestamp(timestamp int64) Carbon {
	return newCarbon(time.Unix(timestamp, 0))
}
func (c Carbon) CreateFromTimestamp(timestamp int64) Carbon {
	return newCarbon(CreateFromTimestamp(timestamp).Time.In(c.loc))
}

// CreateFromDateTime 从年月日时分秒创建Carbon实例
func CreateFromDateTime(year int, month int, day int, hour int, minute int, second int) Carbon {
	return newCarbon(time.Date(year, time.Month(month), day, hour, minute, second, 0, getLocalByTimezone(Local)))
}
func (c Carbon) CreateFromDateTime(year int, month int, day int, hour int, minute int, second int) Carbon {
	return newCarbon(CreateFromDateTime(year, month, day, hour, minute, second).Time.In(c.loc))
}

// CreateFromDate 从年月日创建Carbon实例
func CreateFromDate(year int, month int, day int) Carbon {
	hour, minute, second := time.Now().Clock()
	return newCarbon(time.Date(year, time.Month(month), day, hour, minute, second, 0, getLocalByTimezone(Local)))
}
func (c Carbon) CreateFromDate(year int, month int, day int) Carbon {
	return newCarbon(CreateFromDate(year, month, day).Time.In(c.loc))
}

// CreateFromTime 从时分秒创建Carbon实例
func CreateFromTime(hour int, minute int, second int) Carbon {
	year, month, day := time.Now().Date()
	return newCarbon(time.Date(year, month, day, hour, minute, second, 0, getLocalByTimezone(Local)))
}
func (c Carbon) CreateFromTime(hour int, minute int, second int) Carbon {
	return newCarbon(CreateFromTime(hour, minute, second).Time.In(c.loc))
}

// Parse 解析标准格式时间字符串
func Parse(value string) Carbon {
	value = strings.Trim(value, " ")
	layout := "2006-01-02 15:04:05"

	if len(value) == 10 && strings.Count(value, "-") == 2 {
		layout = "2006-01-02"
	}

	if len(value) == 10 && strings.Count(value, "/") == 3 {
		layout = "2006/01/02"
	}

	if len(value) == 19 && strings.Count(value, "/") == 4 {
		layout = "2006/01/02 15/04/05"
	}

	if len(value) == 14 {
		layout = "20060102150405"
	}

	if len(value) == 8 {
		layout = "20060102"
	}

	t, err := time.ParseInLocation(layout, value, getLocalByTimezone(Local))
	if err != nil {
		panic("invalid format value: " + value)
	}
	return newCarbon(t)
}
func (c Carbon) Parse(value string) Carbon {
	return newCarbon(Parse(value).Time.In(c.loc))
}

// ParseByFormat 解析指定格式时间字符串
func ParseByFormat(value string, format string) Carbon {
	value = strings.Trim(value, " ")
	layout := format2layout(format)
	t, err := time.ParseInLocation(layout, value, getLocalByTimezone(Local))
	if err != nil {
		panic("invalid format value: " + value)
	}
	return newCarbon(t)
}
func (c Carbon) ParseByFormat(value string, format string) Carbon {
	return newCarbon(ParseByFormat(value, format).Time.In(c.loc))
}

// ParseByDuration 解析持续时间字符串(基于现在时间)
// 支持正负整数/浮点数和符号ns(纳秒)、us(微妙)、ms(毫秒)、s(秒)、m(分钟)、h(小时)的组合
func ParseByDuration(duration string) Carbon {
	d, err := time.ParseDuration(duration)
	if err != nil {
		panic("invalid duration value: " + duration)
	}
	return newCarbon(time.Now().Add(d))
}
func (c Carbon) ParseByDuration(duration string) Carbon {
	return newCarbon(ParseByDuration(duration).Time.In(c.loc))
}

// ParseByTime 解析原生time.Time实例
func ParseByTime(t time.Time) Carbon {
	return newCarbon(t)
}
func (c Carbon) ParseByTime(t time.Time) Carbon {
	return newCarbon(ParseByTime(t).Time.In(c.loc))
}

// Duration 按照持续时间字符串改变时间
// 支持正负整数/浮点数和符号ns(纳秒)、us(微妙)、ms(毫秒)、s(秒)、m(分钟)、h(小时)的组合
func (c Carbon) Duration(duration string) Carbon {
	d, err := time.ParseDuration(duration)
	if err != nil {
		panic("invalid duration value: " + duration)
	}
	c.Time = c.Time.Add(d)
	return c
}

// AddYears N年后
func (c Carbon) AddYears(years int) Carbon {
	c.Time = c.Time.AddDate(years, 0, 0)
	return c
}

// AddYear 1年后
func (c Carbon) AddYear() Carbon {
	c.Time = c.Time.AddDate(1, 0, 0)
	return c
}

// SubYears N年前
func (c Carbon) SubYears(years int) Carbon {
	c.Time = c.Time.AddDate(-years, 0, 0)
	return c
}

// SubYear 1年前
func (c Carbon) SubYear() Carbon {
	c.Time = c.Time.AddDate(-1, 0, 0)
	return c
}

// AddMonths N月后
func (c Carbon) AddMonths(months int) Carbon {
	c.Time = c.Time.AddDate(0, months, 0)
	return c
}

// AddMonth 1月后
func (c Carbon) AddMonth() Carbon {
	return c.AddMonths(1)
}

// SubMonths N月前
func (c Carbon) SubMonths(months int) Carbon {
	c.Time = c.Time.AddDate(0, -months, 0)
	return c
}

// SubMonth 1月前
func (c Carbon) SubMonth() Carbon {
	return c.SubMonths(1)
}

// AddDays N天后
func (c Carbon) AddDays(days int) Carbon {
	c.Time = c.Time.AddDate(0, 0, days)
	return c
}

// AddDay 1天后
func (c Carbon) AddDay() Carbon {
	return c.AddDays(1)
}

// SubDays N天前
func (c Carbon) SubDays(days int) Carbon {
	c.Time = c.Time.AddDate(0, 0, -days)
	return c
}

// SubDay 1天前
func (c Carbon) SubDay() Carbon {
	return c.SubDays(1)
}

// AddHours N小时后
func (c Carbon) AddHours(hours int) Carbon {
	duration := time.Duration(hours) * time.Hour
	c.Time = c.Time.Add(duration)
	return c
}

// AddHour 1小时后
func (c Carbon) AddHour() Carbon {
	return c.AddHours(1)
}

// SubHours N小时前
func (c Carbon) SubHours(hours int) Carbon {
	duration := time.Duration(hours) * -time.Hour
	c.Time = c.Time.Add(duration)
	return c
}

// SubHour 1小时前
func (c Carbon) SubHour() Carbon {
	return c.SubHours(1)
}

// AddMinutes N分钟后
func (c Carbon) AddMinutes(minutes int) Carbon {
	duration := time.Duration(minutes) * time.Minute
	c.Time = c.Time.Add(duration)
	return c
}

// AddMinute 1分钟后
func (c Carbon) AddMinute() Carbon {
	return c.AddMinutes(1)
}

// SubMinutes N分钟前
func (c Carbon) SubMinutes(minutes int) Carbon {
	duration := time.Duration(minutes) * -time.Minute
	c.Time = c.Time.Add(duration)
	return c
}

// SubMinute 1分钟前
func (c Carbon) SubMinute() Carbon {
	return c.SubMinutes(1)
}

// AddSeconds N秒钟后
func (c Carbon) AddSeconds(seconds int) Carbon {
	duration := time.Duration(seconds) * time.Second
	c.Time = c.Time.Add(duration)
	return c
}

// AddSecond 1秒钟后
func (c Carbon) AddSecond() Carbon {
	return c.AddSeconds(1)
}

// SubSeconds N秒钟前
func (c Carbon) SubSeconds(seconds int) Carbon {
	duration := time.Duration(seconds) * -time.Second
	c.Time = c.Time.Add(duration)
	return c
}

// SubSecond 1秒钟前
func (c Carbon) SubSecond() Carbon {
	return c.SubSeconds(1)
}

// NextYears N年后
func (c Carbon) NextYears(years int) Carbon {
	year := c.Time.Year() + years
	month := c.Time.Month()
	day := c.Time.Day()

	// 获取N年后本月的最后一天
	last := time.Date(year, month, 1, c.Time.Hour(), c.Time.Minute(), c.Time.Second(), c.Time.Nanosecond(), c.loc).AddDate(0, 1, -1)

	if day > last.Day() {
		day = last.Day()
	}

	c.Time = time.Date(last.Year(), last.Month(), day, c.Time.Hour(), c.Time.Minute(), c.Time.Second(), 0, c.loc)
	return c
}

// NextYear 1年后
func (c Carbon) NextYear() Carbon {
	return c.NextYears(1)
}

// N月后
func (c Carbon) NextMonths(months int) Carbon {
	year := c.Time.Year()
	month := c.Time.Month() + time.Month(months)
	day := c.Time.Day()

	// 获取N月后的最后一天
	last := time.Date(year, month, 1, c.Time.Hour(), c.Time.Minute(), c.Time.Second(), c.Time.Nanosecond(), c.loc).AddDate(0, 1, -1)

	if day > last.Day() {
		day = last.Day()
	}

	c.Time = time.Date(last.Year(), last.Month(), day, c.Time.Hour(), c.Time.Minute(), c.Time.Second(), 0, c.loc)
	return c
}

// NextMonth 1月后
func (c Carbon) NextMonth() Carbon {
	return c.NextMonths(1)
}

// PreYears N年前
func (c Carbon) PreYears(years int) Carbon {
	return c.NextYears(-years)
}

// PreYear 1年前
func (c Carbon) PreYear() Carbon {
	return c.NextYears(-1)
}

// NextMonths N月后
func (c Carbon) PreMonths(months int) Carbon {
	return c.NextMonths(-months)
}

// NextMonth 1月后
func (c Carbon) PreMonth() Carbon {
	return c.NextMonths(-1)
}

// FirstOfYear 年初
func (c Carbon) FirstOfYear() Carbon {
	c.Time = time.Date(c.Time.Year(), 01, 01, c.Time.Hour(), c.Time.Minute(), c.Time.Second(), c.Time.Nanosecond(), c.loc)
	return c
}

// LastOfYear 年末
func (c Carbon) LastOfYear() Carbon {
	c.Time = time.Date(c.Time.Year(), 12, 31, c.Time.Hour(), c.Time.Minute(), c.Time.Second(), c.Time.Nanosecond(), c.loc)
	return c
}

// FirstOfMonth 月初
func (c Carbon) FirstOfMonth() Carbon {
	c.Time = time.Date(c.Time.Year(), c.Time.Month(), 01, c.Time.Hour(), c.Time.Minute(), c.Time.Second(), c.Time.Nanosecond(), c.loc)
	return c
}

// LastOfMonth 月末
func (c Carbon) LastOfMonth() Carbon {
	first := time.Date(c.Time.Year(), c.Time.Month(), 1, c.Time.Hour(), c.Time.Minute(), c.Time.Second(), c.Time.Nanosecond(), c.loc)
	c.Time = first.AddDate(0, 1, -1)
	return c
}

// FirstOfWeek 周初
func (c Carbon) FirstOfWeek() Carbon {
	days := c.Time.Weekday()
	if days == 0 {
		days = DaysPerWeek
	}
	c.Time = c.Time.AddDate(0, 0, int(1-days))
	return c
}

// LastOfWeek 周末(一周的最后一天)
func (c Carbon) LastOfWeek() Carbon {
	days := c.Time.Weekday()
	if days == 0 {
		days = DaysPerWeek
	}
	c.Time = c.Time.AddDate(0, 0, int(DaysPerWeek-days))
	return c
}
