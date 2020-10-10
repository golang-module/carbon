package carbon

import "time"

// ToString 输出字符串
func (c Carbon) ToString() string {
	return c.Time.String()
}

// ToTimestamp 输出时间戳
func (c Carbon) ToTimestamp() int64 {
	return c.Time.Unix()
}

// Format ToFormatString的简称
func (c Carbon) Format(format string) string {
	return c.ToFormatString(format)
}

// ToFormatString 输出指定格式时间
func (c Carbon) ToFormatString(format string) string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.loc).Format(format2layout(format))
}

// ToDayDateTimeString 输出天数日期时间字符串
func (c Carbon) ToDayDateTimeString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.Format(DayDateTimeFormat)
}

// ToDateTimeString 输出日期时间字符串
func (c Carbon) ToDateTimeString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.Format(DateTimeFormat)
}

// ToDateString 输出日期字符串
func (c Carbon) ToDateString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.Format(DateFormat)
}

// ToTimeString 输出时间字符串
func (c Carbon) ToTimeString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.Format(TimeFormat)
}

// ToAtomString 输出Atom格式字符串
func (c Carbon) ToAtomString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.Format(RFC3339Format)
}

// ToAnsicString 输出ANSIC格式字符串
func (c Carbon) ToAnsicString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.Format(AnsicFormat)
}

// ToCookieString 输出Cookie格式字符串
func (c Carbon) ToCookieString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.Format(CookieFormat)
}

// ToRssString 输出RSS格式字符串
func (c Carbon) ToRssString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.Format(RssFormat)
}

// ToW3cString 输出W3C格式字符串
func (c Carbon) ToW3cString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.Format(RFC3339Format)
}

// ToUnixDateString 输出UnixDate格式字符串
func (c Carbon) ToUnixDateString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.Format(UnixDateFormat)
}

// ToUnixDateString 输出RubyDate格式字符串
func (c Carbon) ToRubyDateString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.Format(RubyDateFormat)
}

// ToKitchenString 输出Kitchen格式字符串
func (c Carbon) ToKitchenString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.Format(KitchenFormat)
}

// ToRfc822String 输出RFC822格式字符串
func (c Carbon) ToRFC822String() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.Format(RFC822Format)
}

// ToRfc822String 输出RFC822Z格式字符串
func (c Carbon) ToRFC822zString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.Format(RFC822ZFormat)
}

// ToRfc850String 输出RFC850格式字符串
func (c Carbon) ToRFC850String() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.Format(RFC850Format)
}

// ToRfc1036String 输出RFC1036格式字符串
func (c Carbon) ToRFC1036String() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.Format(RFC1036Format)
}

// ToRfc1123String 输出RFC1123格式字符串
func (c Carbon) ToRFC1123String() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.Format(RFC1123Format)
}

// ToRFC1123ZString 输出RFC1123Z格式字符串
func (c Carbon) ToRFC1123ZString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.Format(RFC1123ZFormat)
}

// ToRFC2822String 输出RFC2822格式字符串
func (c Carbon) ToRFC2822String() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.Format(RFC2822Format)
}

// ToRfc3339String 输出RFC3339格式字符串
func (c Carbon) ToRFC3339String() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.Format(RFC3339Format)
}

// ToRfc7231String 输出RFC7231格式字符串
func (c Carbon) ToRFC7231String() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.Format(RFC7231Format)
}

// DaysInYear 获取本年的总天数
func (c Carbon) DaysInYear() int {
	if c.Time.IsZero() {
		return 0
	}
	days := DaysPerNormalYear
	if c.IsLeapYear() {
		days = DaysPerLeapYear
	}
	return days
}

// DaysInMonth 获取本月的总天数
func (c Carbon) DaysInMonth() int {
	if c.Time.IsZero() {
		return 0
	}
	return c.EndOfMonth().Time.Day()
}

// MonthOfYear 获取本年的第几月
func (c Carbon) MonthOfYear() int {
	if c.Time.IsZero() {
		return 0
	}
	return int(c.Time.Month())
}

// DayOfYear 获取本年的第几天
func (c Carbon) DayOfYear() int {
	if c.Time.IsZero() {
		return 0
	}
	return c.Time.YearDay()
}

// DayOfMonth 获取本月的第几天
func (c Carbon) DayOfMonth() int {
	if c.Time.IsZero() {
		return 0
	}
	return c.Time.Day()
}

// DayOfWeek 获取本周的第几天
func (c Carbon) DayOfWeek() int {
	if c.Time.IsZero() {
		return 0
	}
	return int(c.Time.Weekday())
}

// WeekOfYear 获取本年的第几周
func (c Carbon) WeekOfYear() int {
	if c.Time.IsZero() {
		return 0
	}
	_, week := c.Time.ISOWeek()
	return week
}

// WeekOfMonth 获取本月的第几周
func (c Carbon) WeekOfMonth() int {
	if c.Time.IsZero() {
		return 0
	}
	day := c.Time.Day()
	if day < DaysPerWeek {
		return 1
	}
	return day%DaysPerWeek + 1
}

// IsZero 是否是零值
func (c Carbon) IsZero() bool {
	return c.Time.IsZero()
}

// IsFuture 是否是未来时间
func (c Carbon) IsFuture() bool {
	if c.Time.IsZero() {
		return false
	}
	return c.ToTimestamp() > c.Now().ToTimestamp()
}

// IsPast 是否是过去时间
func (c Carbon) IsPast() bool {
	if c.Time.IsZero() {
		return true
	}
	return c.ToTimestamp() < c.Now().ToTimestamp()
}

// IsLeapYear 是否是闰年
func (c Carbon) IsLeapYear() bool {
	year := c.Time.Year()
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

// IsJanuary 是否是一月
func (c Carbon) IsJanuary() bool {
	return c.Time.In(c.loc).Month() == time.January
}

// IsMonday 是否是二月
func (c Carbon) IsFebruary() bool {
	return c.Time.In(c.loc).Month() == time.February
}

// IsMarch 是否是三月
func (c Carbon) IsMarch() bool {
	return c.Time.In(c.loc).Month() == time.March
}

// IsApril 是否是四月
func (c Carbon) IsApril() bool {
	return c.Time.In(c.loc).Month() == time.April
}

// IsMay 是否是五月
func (c Carbon) IsMay() bool {
	return c.Time.In(c.loc).Month() == time.May
}

// IsJune 是否是六月
func (c Carbon) IsJune() bool {
	return c.Time.In(c.loc).Month() == time.June
}

// IsJuly 是否是七月
func (c Carbon) IsJuly() bool {
	return c.Time.In(c.loc).Month() == time.July
}

// IsAugust 是否是八月
func (c Carbon) IsAugust() bool {
	return c.Time.In(c.loc).Month() == time.August
}

// IsSeptember 是否是九月
func (c Carbon) IsSeptember() bool {
	return c.Time.In(c.loc).Month() == time.September
}

// IsOctober 是否是十月
func (c Carbon) IsOctober() bool {
	return c.Time.In(c.loc).Month() == time.October
}

// IsNovember 是否是十一月
func (c Carbon) IsNovember() bool {
	return c.Time.In(c.loc).Month() == time.November
}

// IsDecember 是否是十二月
func (c Carbon) IsDecember() bool {
	return c.Time.In(c.loc).Month() == time.December
}

// IsMonday 是否是周一
func (c Carbon) IsMonday() bool {
	return c.Time.In(c.loc).Weekday() == time.Monday
}

// IsTuesday 是否是周二
func (c Carbon) IsTuesday() bool {
	return c.Time.In(c.loc).Weekday() == time.Tuesday
}

// IsWednesday 是否是周三
func (c Carbon) IsWednesday() bool {
	return c.Time.In(c.loc).Weekday() == time.Wednesday
}

// IsThursday 是否是周四
func (c Carbon) IsThursday() bool {
	return c.Time.In(c.loc).Weekday() == time.Thursday
}

// IsFriday 是否是周五
func (c Carbon) IsFriday() bool {
	return c.Time.In(c.loc).Weekday() == time.Friday
}

// IsSaturday 是否是周六
func (c Carbon) IsSaturday() bool {
	return c.Time.In(c.loc).Weekday() == time.Saturday
}

// IsSunday 是否是周日
func (c Carbon) IsSunday() bool {
	return c.Time.In(c.loc).Weekday() == time.Sunday
}

// IsWeekday 是否是工作日
func (c Carbon) IsWeekday() bool {
	return !c.IsSaturday() && !c.IsSunday()
}

// IsWeekend 是否是周末
func (c Carbon) IsWeekend() bool {
	return c.IsSaturday() || c.IsSunday()
}

// IsYesterday 是否是昨天
func (c Carbon) IsYesterday() bool {
	return c.ToDateString() == c.Yesterday().ToDateString()
}

// IsToday 是否是今天
func (c Carbon) IsToday() bool {
	return c.ToDateString() == c.Now().ToDateString()
}

// IsTomorrow 是否是明天
func (c Carbon) IsTomorrow() bool {
	return c.ToDateString() == c.Tomorrow().ToDateString()
}
