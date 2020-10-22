package carbon

import (
	"strings"
	"time"
)

type Carbon struct {
	Time time.Time
	loc  *time.Location
}

// Timezone 设置时区
func SetTimezone(name string) Carbon {
	return Carbon{loc: getLocalByTimezone(name)}
}

// Timezone 设置时区
func (c Carbon) SetTimezone(name string) Carbon {
	loc := getLocalByTimezone(name)
	return Carbon{Time: c.Time.In(c.loc), loc: loc}
}

// SetYear 设置年
func (c Carbon) SetYear(year int) Carbon {
	c.Time = time.Date(year, c.Time.Month(), c.Time.Day(), c.Time.Hour(), c.Time.Minute(), c.Time.Second(), 0, c.loc)
	return c
}

// SetMonth 设置月
func (c Carbon) SetMonth(month int) Carbon {
	c.Time = time.Date(c.Time.Year(), time.Month(month), c.Time.Day(), c.Time.Hour(), c.Time.Minute(), c.Time.Second(), 0, c.loc)
	return c
}

// SetDay 设置日
func (c Carbon) SetDay(day int) Carbon {
	c.Time = time.Date(c.Time.Year(), c.Time.Month(), day, c.Time.Hour(), c.Time.Minute(), c.Time.Second(), 0, c.loc)
	return c
}

// SetHour 设置时
func (c Carbon) SetHour(hour int) Carbon {
	c.Time = time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), hour, c.Time.Minute(), c.Time.Second(), 0, c.loc)
	return c
}

// SetMinute 设置分
func (c Carbon) SetMinute(minute int) Carbon {
	c.Time = time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), c.Time.Hour(), minute, c.Time.Second(), 0, c.loc)
	return c
}

// SetSecond 设置秒
func (c Carbon) SetSecond(second int) Carbon {
	c.Time = time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), c.Time.Hour(), c.Time.Minute(), second, 0, c.loc)
	return c
}

// Now 当前
func Now() Carbon {
	return newCarbon(time.Now())
}

// Now 当前(指定时区)
func (c Carbon) Now() Carbon {
	return newCarbon(Now().Time.In(c.loc))
}

// Tomorrow 明天
func Tomorrow() Carbon {
	return newCarbon(time.Now().AddDate(0, 0, 1))
}

// Tomorrow 明天(指定时区)
func (c Carbon) Tomorrow() Carbon {
	return newCarbon(Tomorrow().Time.In(c.loc))
}

// Yesterday 昨天
func Yesterday() Carbon {
	return newCarbon(time.Now().AddDate(0, 0, -1))
}

// Yesterday 昨天(指定时区)
func (c Carbon) Yesterday() Carbon {
	return newCarbon(Yesterday().Time.In(c.loc))
}

// CreateFromTimestamp 从时间戳创建Carbon实例
func CreateFromTimestamp(timestamp int64) Carbon {
	return newCarbon(time.Unix(timestamp, 0))
}

// CreateFromTimestamp 从时间戳创建Carbon实例(指定时区)
func (c Carbon) CreateFromTimestamp(timestamp int64) Carbon {
	return newCarbon(CreateFromTimestamp(timestamp).Time.In(c.loc))
}

// CreateFromDateTime 从年月日时分秒创建Carbon实例
func CreateFromDateTime(year int, month int, day int, hour int, minute int, second int) Carbon {
	return newCarbon(time.Date(year, time.Month(month), day, hour, minute, second, 0, getLocalByTimezone(Local)))
}

// CreateFromDateTime 从年月日时分秒创建Carbon实例(指定时区)
func (c Carbon) CreateFromDateTime(year int, month int, day int, hour int, minute int, second int) Carbon {
	return newCarbon(CreateFromDateTime(year, month, day, hour, minute, second).Time.In(c.loc))
}

// CreateFromDate 从年月日创建Carbon实例
func CreateFromDate(year int, month int, day int) Carbon {
	hour, minute, second := time.Now().Clock()
	return newCarbon(time.Date(year, time.Month(month), day, hour, minute, second, 0, getLocalByTimezone(Local)))
}

// CreateFromDate 从年月日创建Carbon实例(指定时区)
func (c Carbon) CreateFromDate(year int, month int, day int) Carbon {
	return newCarbon(CreateFromDate(year, month, day).Time.In(c.loc))
}

// CreateFromTime 从时分秒创建Carbon实例
func CreateFromTime(hour int, minute int, second int) Carbon {
	year, month, day := time.Now().Date()
	return newCarbon(time.Date(year, month, day, hour, minute, second, 0, getLocalByTimezone(Local)))
}

// CreateFromTime 从时分秒创建Carbon实例(指定时区)
func (c Carbon) CreateFromTime(hour int, minute int, second int) Carbon {
	return newCarbon(CreateFromTime(hour, minute, second).Time.In(c.loc))
}

// CreateFromGoTime 从原生time.Time创建Carbon实例
func CreateFromGoTime(t time.Time) Carbon {
	return newCarbon(t)
}

// CreateFromGoTime 从原生time.Time创建Carbon实例(指定时区)
func (c Carbon) CreateFromGoTime(t time.Time) Carbon {
	return newCarbon(CreateFromGoTime(t).Time.In(c.loc))
}

// Parse 解析标准格式时间字符串
func Parse(value string) Carbon {
	layout := DateTimeFormat

	if value == "" || value == "0" || value == "0000-00-00 00:00:00" || value == "0000-00-00" || value == "00:00:00" {
		return Carbon{loc: getLocalByTimezone(Local)}
	}

	if len(value) == 10 && strings.Count(value, "-") == 2 {
		layout = DateFormat
	}

	if len(value) == 14 {
		layout = ShortDateTimeFormat
	}

	if len(value) == 8 {
		layout = ShortDateFormat
	}

	if strings.Index(value, "T") == 10 {
		layout = RFC3339Format
	}

	return newCarbon(parseByLayout(value, layout))
}

// Parse 解析标准格式时间字符串(指定时区)
func (c Carbon) Parse(value string) Carbon {
	return newCarbon(Parse(value).Time.In(c.loc))
}

// ParseByFormat 解析指定格式时间字符串
func ParseByFormat(value string, format string) Carbon {
	value = strings.Trim(value, " ")
	layout := format2layout(format)
	return newCarbon(parseByLayout(value, layout))
}

// ParseByFormat 解析指定格式时间字符串(指定时区)
func (c Carbon) ParseByFormat(value string, format string) Carbon {
	return newCarbon(ParseByFormat(value, format).Time.In(c.loc))
}

// ParseByDuration 解析持续时间字符串(基于现在时间)
// 支持正负整数/浮点数和符号ns(纳秒)、us(微妙)、ms(毫秒)、s(秒)、m(分钟)、h(小时)的组合
func ParseByDuration(duration string) Carbon {
	return newCarbon(time.Now().Add(parseByDuration(duration)))
}

// ParseByDuration 解析持续时间字符串(指定时区)
func (c Carbon) ParseByDuration(duration string) Carbon {
	return newCarbon(ParseByDuration(duration).Time.In(c.loc))
}

// Duration 按照持续时间字符串改变时间(指定时区)
// 支持正负整数/浮点数和符号ns(纳秒)、us(微妙)、ms(毫秒)、s(秒)、m(分钟)、h(小时)的组合
func (c Carbon) Duration(duration string) Carbon {
	c.Time = c.Time.Add(parseByDuration(duration))
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
	last := time.Date(year, month, 1, c.Time.Hour(), c.Time.Minute(), c.Time.Second(), 0, c.loc).AddDate(0, 1, -1)

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
	last := time.Date(year, month, 1, c.Time.Hour(), c.Time.Minute(), c.Time.Second(), 0, c.loc).AddDate(0, 1, -1)

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

// StartOfYear 本年开始时间
func (c Carbon) StartOfYear() Carbon {
	c.Time = time.Date(c.Time.Year(), 1, 1, 0, 0, 0, 0, c.loc)
	return c
}

// EndOfYear 本年结束时间
func (c Carbon) EndOfYear() Carbon {
	c.Time = time.Date(c.Time.Year(), 12, 31, 23, 59, 59, 0, c.loc)
	return c
}

// StartOfMonth 本月开始时间
func (c Carbon) StartOfMonth() Carbon {
	c.Time = time.Date(c.Time.Year(), c.Time.Month(), 1, 0, 0, 0, 0, c.loc)
	return c
}

// EndOfMonth 本月结束时间
func (c Carbon) EndOfMonth() Carbon {
	t := time.Date(c.Time.Year(), c.Time.Month(), 1, 23, 59, 59, 0, c.loc)
	c.Time = t.AddDate(0, 1, -1)
	return c
}

// StartOfWeek 本周开始时间
func (c Carbon) StartOfWeek() Carbon {
	days := c.Time.Weekday()
	if days == 0 {
		days = DaysPerWeek
	}
	t := time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), 0, 0, 0, 0, c.loc)
	c.Time = t.AddDate(0, 0, int(1-days))
	return c
}

// EndOfWeek 本周结束时间
func (c Carbon) EndOfWeek() Carbon {
	days := c.Time.Weekday()
	if days == 0 {
		days = DaysPerWeek
	}
	t := time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), 23, 59, 59, 0, c.loc)
	c.Time = t.AddDate(0, 0, int(DaysPerWeek-days))
	return c
}

// StartOfDay 本日开始时间
func (c Carbon) StartOfDay() Carbon {
	c.Time = time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), 0, 0, 0, 0, c.loc)
	return c
}

// EndOfDay 本日结束时间
func (c Carbon) EndOfDay() Carbon {
	c.Time = time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), 23, 59, 59, 0, c.loc)
	return c
}

// StartOfHour 小时开始时间
func (c Carbon) StartOfHour() Carbon {
	c.Time = time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), c.Time.Hour(), 0, 0, 0, c.loc)
	return c
}

// EndOfHour 小时结束时间
func (c Carbon) EndOfHour() Carbon {
	c.Time = time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), c.Time.Hour(), 59, 59, 0, c.loc)
	return c
}

// StartOfMinute 分钟开始时间
func (c Carbon) StartOfMinute() Carbon {
	c.Time = time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), c.Time.Hour(), c.Time.Minute(), 0, 0, c.loc)
	return c
}

// EndOfMinute 分钟结束时间
func (c Carbon) EndOfMinute() Carbon {
	c.Time = time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), c.Time.Hour(), c.Time.Minute(), 59, 0, c.loc)
	return c
}
