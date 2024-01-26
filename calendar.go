package carbon

import (
	"github.com/golang-module/carbon/v2/calendar/julian"
	"github.com/golang-module/carbon/v2/calendar/lunar"
)

// Lunar converts gregorian calendar to lunar calendar.
// 将 公历 转化为 农历
func (c Carbon) Lunar() (l lunar.Lunar) {
	if c.Error != nil {
		l.Error = c.Error
		return l
	}
	return lunar.FromGregorian(c.ToStdTime()).ToLunar()
}

// CreateFromLunar creates a Carbon instance from Lunar date and time.
// 从 农历日期 创建 Carbon 实例
func CreateFromLunar(year, month, day, hour, minute, second int, isLeapMonth bool) Carbon {
	c := NewCarbon()
	c.time = lunar.FromLunar(year, month, day, hour, minute, second, isLeapMonth).ToGregorian().Time
	return c
}

// Julian converts gregorian calendar to Julian calendar
// 将 公历 转化为 儒略历
func (c Carbon) Julian() (j julian.Julian) {
	if c.Error != nil {
		return j
	}
	return julian.FromGregorian(c.ToStdTime()).ToJulian()
}

// CreateFromJulian creates a Carbon instance from Julian Day.or Modified Julian Day
// 从 儒略日/简化儒略日 创建 Carbon 实例
func CreateFromJulian(f float64) Carbon {
	c := NewCarbon()
	c.time = julian.FromJulian(f).ToGregorian().Time
	return c
}
