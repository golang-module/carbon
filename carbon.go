// @Title carbon
// @Description A simple, semantic and developer-friendly golang package for datetime
// @Page github.com/golang-module/carbon
// @Version v1.3.4
// @Author gouguoyin
// @Email mail@gouguoyin.cn

package carbon

import "time"

// 时区常量
const (
	Local = "Local"
	CET   = "CET"
	EET   = "EET"
	EST   = "EST"
	GMT   = "GMT"
	UTC   = "UTC"
	UCT   = "UCT"
	MST   = "MST"

	Cuba      = "Cuba"
	Egypt     = "Egypt"
	Eire      = "Eire"
	Greenwich = "Greenwich"
	Iceland   = "Iceland"
	Iran      = "Iran"
	Israel    = "Israel"
	Jamaica   = "Jamaica"
	Japan     = "Japan"
	Libya     = "Libya"
	Poland    = "Poland"
	Portugal  = "Portugal"
	PRC       = "PRC"
	Singapore = "Singapore"
	Turkey    = "Turkey"
	Zulu      = "Zulu"

	Shanghai   = "Asia/Shanghai"
	Chongqing  = "Asia/Chongqing"
	HongKong   = "Asia/Hong_Kong"
	Macao      = "Asia/Macao"
	Taipei     = "Asia/Taipei"
	Tokyo      = "Asia/Tokyo"
	London     = "Europe/London"
	NewYork    = "America/New_York"
	LosAngeles = "America/Los_Angeles"
)

// 月份常量
const (
	January   = "January"   // 一月
	February  = "February"  // 二月
	March     = "March"     // 三月
	April     = "April"     // 四月
	May       = "May"       // 五月
	June      = "June"      // 六月
	July      = "July"      // 七月
	August    = "August"    // 八月
	September = "September" // 九月
	October   = "October"   // 十月
	November  = "November"  // 十一月
	December  = "December"  // 十二月
)

// 星期常量
const (
	Monday    = "Monday"    // 周一
	Tuesday   = "Tuesday"   // 周二
	Wednesday = "Wednesday" // 周三
	Thursday  = "Thursday"  // 周四
	Friday    = "Friday"    // 周五
	Saturday  = "Saturday"  // 周六
	Sunday    = "Sunday"    // 周日
)

// 数字常量
const (
	YearsPerMillennium         = 1000    // 每千年1000年
	YearsPerCentury            = 100     // 每世纪100年
	YearsPerDecade             = 10      // 每十年10年
	QuartersPerYear            = 4       // 每年4季度
	MonthsPerYear              = 12      // 每年12月
	MonthsPerQuarter           = 3       // 每季度3月
	WeeksPerNormalYear         = 52      // 每常规年52周
	weeksPerLongYear           = 53      // 每长年53周
	WeeksPerMonth              = 4       // 每月4周
	DaysPerLeapYear            = 366     // 每闰年366天
	DaysPerNormalYear          = 365     // 每常规年365天
	DaysPerWeek                = 7       // 每周7天
	HoursPerWeek               = 168     // 每周168小时
	HoursPerDay                = 24      // 每天24小时
	MinutesPerDay              = 1440    // 每天1440分钟
	MinutesPerHour             = 60      // 每小时60分钟
	SecondsPerWeek             = 604800  // 每周604800秒
	SecondsPerDay              = 86400   // 每天86400秒
	SecondsPerHour             = 3600    // 每小时3600秒
	SecondsPerMinute           = 60      // 每分钟60秒
	MillisecondsPerSecond      = 1000    // 每秒1000毫秒
	MicrosecondsPerMillisecond = 1000    // 每毫秒1000微秒
	MicrosecondsPerSecond      = 1000000 // 每秒1000000微秒
)

// 时间格式化常量
const (
	AnsicFormat         = time.ANSIC
	UnixDateFormat      = time.UnixDate
	RubyDateFormat      = time.RubyDate
	RFC822Format        = time.RFC822
	RFC822ZFormat       = time.RFC822Z
	RFC850Format        = time.RFC850
	RFC1123Format       = time.RFC1123
	RFC1123ZFormat      = time.RFC1123Z
	RssFormat           = time.RFC1123Z
	RFC2822Format       = time.RFC1123Z
	RFC3339Format       = time.RFC3339
	KitchenFormat       = time.Kitchen
	CookieFormat        = "Monday, 02-Jan-2006 15:04:05 MST"
	RFC1036Format       = "Mon, 02 Jan 06 15:04:05 -0700"
	RFC7231Format       = "Mon, 02 Jan 2006 15:04:05 GMT"
	DayDateTimeFormat   = "Mon, Aug 2, 2006 3:04 PM"
	DateTimeFormat      = "2006-01-02 15:04:05"
	DateFormat          = "2006-01-02"
	TimeFormat          = "15:04:05"
	ShortDateTimeFormat = "20060102150405"
	ShortDateFormat     = "20060102"
	ShortTimeFormat     = "150405"
)

// Carbon 定义 Carbon 结构体
type Carbon struct {
	Time  time.Time
	Loc   *time.Location
	Lang  *Language
	Error error
}

// Time2Carbon 将 time.Time 转换成 Carbon
func Time2Carbon(tt time.Time) Carbon {
	loc, _ := time.LoadLocation(Local)
	return Carbon{Time: tt, Loc: loc}
}

// Carbon2Time 将 Carbon 转换成 time.Time
func (c Carbon) Carbon2Time() time.Time {
	return c.Time
}

// Now 当前
func (c Carbon) Now() Carbon {
	c.Time = time.Now()
	return c
}

// Now 当前(默认时区)
func Now() Carbon {
	return SetTimezone(Local).Now()
}

// Tomorrow 明天
func (c Carbon) Tomorrow() Carbon {
	if c.Time.IsZero() {
		c.Time = time.Now().AddDate(0, 0, 1)
	} else {
		c.Time = c.Time.AddDate(0, 0, 1)
	}
	return c
}

// Tomorrow 明天(默认时区)
func Tomorrow() Carbon {
	return SetTimezone(Local).Tomorrow()
}

// Yesterday 昨天
func (c Carbon) Yesterday() Carbon {
	if c.IsZero() {
		c.Time = time.Now().AddDate(0, 0, -1)
	} else {
		c.Time = c.Time.AddDate(0, 0, -1)
	}
	return c
}

// Yesterday 昨天(默认时区)
func Yesterday() Carbon {
	return SetTimezone(Local).Yesterday()
}
