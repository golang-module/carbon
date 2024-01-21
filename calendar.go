package carbon

import (
	"github.com/golang-module/v2/calendar/lunar"
)

// Lunar converts the solar calendar to the lunar calendar.
// 将公历转为农历
func (c Carbon) Lunar() (l lunar.Calendar) {
	if c.Error != nil {
		l.Error = c.Error
		return l
	}
	return lunar.CreateFromSolar(lunar.NewSolar(c.ToStdTime()))
}

// CreateFromLunar creates a Carbon instance from Lunar date and time.
// 从 农历日期 创建 Carbon 实例
func CreateFromLunar(year, month, day, hour, minute, second int, isLeapMonth bool) Carbon {
	c := NewCarbon()
	c.time = lunar.CreateFromLunar(lunar.NewLunar(year, month, day, hour, minute, second, isLeapMonth)).Solar.Time
	return c
}
