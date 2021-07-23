package carbon

import "strings"

// DiffInYears 相差多少年
func (c Carbon) DiffInYears(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) == 1 {
		end = carbon[0]
	}
	return c.DiffInMonths(end) / MonthsPerYear
}

// DiffInYearsWithAbs 相差多少年(绝对值)
func (c Carbon) DiffInYearsWithAbs(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) == 1 {
		end = carbon[0]
	}
	return getAbsValue(c.DiffInYears(end))
}

// DiffInMonths 相差多少月
func (c Carbon) DiffInMonths(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) == 1 {
		end = carbon[0]
	}
	dy, dm, dd := end.Year()-c.Year(), end.Month()-c.Month(), end.Day()-c.Day()
	if dd < 0 {
		dm = dm - 1
	}
	if dy == 0 && dm == 0 {
		return 0
	}
	if dy == 0 && dm != 0 && dd != 0 {
		if int(end.DiffInHoursWithAbs(c)) < c.DaysInMonth()*HoursPerDay {
			return 0
		}
		return int64(dm)
	}
	return int64(dy*MonthsPerYear + dm)
}

// DiffInMonthsWithAbs 相差多少月(绝对值)
func (c Carbon) DiffInMonthsWithAbs(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) == 1 {
		end = carbon[0]
	}
	return getAbsValue(c.DiffInMonths(end))
}

// DiffInWeeks 相差多少周
func (c Carbon) DiffInWeeks(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) == 1 {
		end = carbon[0]
	}
	return c.DiffInDays(end) / DaysPerWeek
}

// DiffInWeeksWithAbs 相差多少周(绝对值)
func (c Carbon) DiffInWeeksWithAbs(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) == 1 {
		end = carbon[0]
	}
	return getAbsValue(c.DiffInWeeks(end))
}

// DiffInDays 相差多少天
func (c Carbon) DiffInDays(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) == 1 {
		end = carbon[0]
	}
	return c.DiffInSeconds(end) / SecondsPerDay
}

// DiffInDaysWithAbs 相差多少天(绝对值)
func (c Carbon) DiffInDaysWithAbs(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) == 1 {
		end = carbon[0]
	}
	return getAbsValue(c.DiffInDays(end))
}

// DiffInHours 相差多少小时
func (c Carbon) DiffInHours(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) == 1 {
		end = carbon[0]
	}
	return c.DiffInSeconds(end) / SecondsPerHour
}

// DiffInHoursWithAbs 相差多少小时(绝对值)
func (c Carbon) DiffInHoursWithAbs(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) == 1 {
		end = carbon[0]
	}
	return getAbsValue(c.DiffInHours(end))
}

// DiffInMinutes 相差多少分钟
func (c Carbon) DiffInMinutes(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) == 1 {
		end = carbon[0]
	}
	return c.DiffInSeconds(end) / SecondsPerMinute
}

// DiffInMinutesWithAbs 相差多少分钟(绝对值)
func (c Carbon) DiffInMinutesWithAbs(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) == 1 {
		end = carbon[0]
	}
	return getAbsValue(c.DiffInMinutes(end))
}

// DiffInSeconds 相差多少秒
func (c Carbon) DiffInSeconds(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) == 1 {
		end = carbon[0]
	}
	return end.ToTimestamp() - c.ToTimestamp()
}

// DiffInSecondsWithAbs 相差多少秒(绝对值)
func (c Carbon) DiffInSecondsWithAbs(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) == 1 {
		end = carbon[0]
	}
	return getAbsValue(c.DiffInSeconds(end))
}

// DiffForHumans 获取对人类友好的可读格式时间差，支持i18n
func (c Carbon) DiffForHumans(carbon ...Carbon) string {
	end := c.Now()
	if len(carbon) == 1 {
		end = carbon[0]
	}
	unit, diff := "", int64(0)
	switch true {
	case c.DiffInYearsWithAbs(end) > 0:
		unit = "year"
		diff = c.DiffInYearsWithAbs(end)
		break
	case c.DiffInMonthsWithAbs(end) > 0:
		unit = "month"
		diff = c.DiffInMonthsWithAbs(end)
		break
	case c.DiffInWeeksWithAbs(end) > 0:
		unit = "week"
		diff = c.DiffInWeeksWithAbs(end)
		break
	case c.DiffInDaysWithAbs(end) > 0:
		unit = "day"
		diff = c.DiffInDaysWithAbs(end)
		break
	case c.DiffInHoursWithAbs(end) > 0:
		unit = "hour"
		diff = c.DiffInHoursWithAbs(end)
		break
	case c.DiffInMinutesWithAbs(end) > 0:
		unit = "minute"
		diff = c.DiffInMinutesWithAbs(end)
	case c.DiffInSecondsWithAbs(end) > 0:
		unit = "second"
		diff = c.DiffInSecondsWithAbs(end)
	case c.DiffInSecondsWithAbs(end) == 0:
		unit = "now"
		diff = 0
		return c.Lang.translate(unit, diff)
	}
	translation := c.Lang.translate(unit, diff)
	if c.Lt(end) && len(carbon) == 0 {
		return strings.Replace(c.Lang.resources["ago"], "%s", translation, 1)
	}
	if c.Lt(end) && len(carbon) == 1 {
		return strings.Replace(c.Lang.resources["before"], "%s", translation, 1)
	}
	if c.Gt(end) && len(carbon) == 0 {
		return strings.Replace(c.Lang.resources["from_now"], "%s", translation, 1)
	}
	if c.Gt(end) && len(carbon) == 1 {
		return strings.Replace(c.Lang.resources["after"], "%s", translation, 1)
	}
	return translation
}
