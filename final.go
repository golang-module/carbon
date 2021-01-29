package carbon

import (
	"strconv"
	"strings"
	"time"
)

// ToGoTime 输出carbon实例
func (c Carbon) ToGoTime() time.Time {
	return c.Time
}

// ToString 输出字符串
func (c Carbon) ToString() string {
	return c.Time.String()
}

// ToTimestamp ToTimestampWithSecond的简称
func (c Carbon) ToTimestamp() int64 {
	return c.ToTimestampWithSecond()
}

// ToTimestampWithSecond 输出秒级时间戳
func (c Carbon) ToTimestampWithSecond() int64 {
	return c.Time.Unix()
}

// ToTimestampWithMillisecond 输出毫秒级时间戳
func (c Carbon) ToTimestampWithMillisecond() int64 {
	return c.Time.UnixNano() / int64(time.Millisecond)
}

// ToTimestampWithMicrosecond 输出微秒级时间戳
func (c Carbon) ToTimestampWithMicrosecond() int64 {
	return c.Time.UnixNano() / int64(time.Microsecond)
}

// ToTimestampWithNanosecond 输出纳秒级时间戳
func (c Carbon) ToTimestampWithNanosecond() int64 {
	return c.Time.UnixNano()
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

	s := c.Time.In(c.Loc).Format(format2layout(format))

	// ISO-8601 格式数字表示的星期中的第几天，取值范围 1-7
	if strings.Contains(s, "N") {
		s = strings.Replace(s, "N", strconv.Itoa(c.DayOfWeek()), 1)
	}

	// 是否为闰年，如果是闰年为 1，否则为 0
	if strings.Contains(s, "L") {
		if c.IsLeapYear() {
			s = strings.Replace(s, "L", "1", 1)
		} else {
			s = strings.Replace(s, "L", "0", 1)
		}
	}

	// 数字表示的小时，24 小时格式，没有前导零，取值范围 0-23
	if strings.Contains(s, "G") {
		s = c.Time.In(c.Loc).Format("15")
		s = strings.Replace(s, "0", "", 1)
	}

	// 秒级时间戳，如 1611818268
	if strings.Contains(s, "U") {
		s = strings.Replace(s, "U", strconv.FormatInt(c.ToTimestamp(), 10), 1)
	}

	// 数字表示的毫秒，如 999
	if strings.Contains(s, "u") {
		s = strings.Replace(s, "u", strconv.Itoa(c.Millisecond()), 1)
	}

	// 数字表示的星期中的第几天，取值范围 0-6
	if strings.Contains(s, "w") {
		s = strings.Replace(s, "w", strconv.Itoa(c.DayOfWeek()-1), 1)
	}

	// 指定的月份有几天，取值范围 28-31
	if strings.Contains(s, "t") {
		s = strings.Replace(s, "t", strconv.Itoa(c.DaysInMonth()), 1)
	}

	// 年份中的第几天，取值范围 0-365
	if strings.Contains(s, "z") {
		s = strings.Replace(s, "z", strconv.Itoa(c.DayOfYear()-1), 1)
	}

	// 时区标识，如 UTC，GMT，Atlantic/Azores
	if strings.Contains(s, "e") {
		s = strings.Replace(s, "e", c.Timezone(), 1)
	}

	return s
}

// ToDayDateTimeString 输出天数日期时间字符串
func (c Carbon) ToDayDateTimeString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(DayDateTimeFormat)
}

// ToDateTimeString 输出日期时间字符串
func (c Carbon) ToDateTimeString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(DateTimeFormat)
}

// ToDateString 输出日期字符串
func (c Carbon) ToDateString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(DateFormat)
}

// ToTimeString 输出时间字符串
func (c Carbon) ToTimeString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(TimeFormat)
}

// ToAtomString 输出Atom格式字符串
func (c Carbon) ToAtomString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RFC3339Format)
}

// ToAnsicString 输出ANSIC格式字符串
func (c Carbon) ToAnsicString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(AnsicFormat)
}

// ToCookieString 输出Cookie格式字符串
func (c Carbon) ToCookieString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(CookieFormat)
}

// ToRssString 输出RSS格式字符串
func (c Carbon) ToRssString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RssFormat)
}

// ToW3cString 输出W3C格式字符串
func (c Carbon) ToW3cString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RFC3339Format)
}

// ToUnixDateString 输出UnixDate格式字符串
func (c Carbon) ToUnixDateString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(UnixDateFormat)
}

// ToUnixDateString 输出RubyDate格式字符串
func (c Carbon) ToRubyDateString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RubyDateFormat)
}

// ToKitchenString 输出Kitchen格式字符串
func (c Carbon) ToKitchenString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(KitchenFormat)
}

// ToRFC822String 输出RFC822格式字符串
func (c Carbon) ToRFC822String() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RFC822Format)
}

// ToRFC822zString 输出RFC822Z格式字符串
func (c Carbon) ToRFC822zString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RFC822ZFormat)
}

// ToRFC850String 输出RFC850格式字符串
func (c Carbon) ToRFC850String() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RFC850Format)
}

// ToRFC1036String 输出RFC1036格式字符串
func (c Carbon) ToRFC1036String() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RFC1036Format)
}

// ToRFC1123String 输出RFC1123格式字符串
func (c Carbon) ToRFC1123String() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RFC1123Format)
}

// ToRFC1123ZString 输出RFC1123Z格式字符串
func (c Carbon) ToRFC1123ZString() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RFC1123ZFormat)
}

// ToRFC2822String 输出RFC2822格式字符串
func (c Carbon) ToRFC2822String() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RFC2822Format)
}

// ToRFC3339String 输出RFC3339格式字符串
func (c Carbon) ToRFC3339String() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RFC3339Format)
}

// ToRFC7231String 输出RFC7231格式字符串
func (c Carbon) ToRFC7231String() string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RFC7231Format)
}

// DiffInWeeks 相差多少周
func (start Carbon) DiffInWeeks(end Carbon) int64 {
	return start.DiffInSeconds(end) / SecondsPerWeek
}

// DiffInWeeksWithAbs 相差多少周（绝对值）
func (start Carbon) DiffInWeeksWithAbs(end Carbon) int64 {
	diff := start.DiffInWeeks(end)
	if diff < 0 {
		diff = -diff
	}
	return diff
}

// DiffInDays 相差多少天
func (start Carbon) DiffInDays(end Carbon) int64 {
	return start.DiffInSeconds(end) / SecondsPerDay
}

// DiffInDaysWithAbs 相差多少天（绝对值）
func (start Carbon) DiffInDaysWithAbs(end Carbon) int64 {
	diff := start.DiffInDays(end)
	if diff < 0 {
		diff = -diff
	}
	return diff
}

// DiffInHours 相差多少小时
func (start Carbon) DiffInHours(end Carbon) int64 {
	return start.DiffInSeconds(end) / SecondsPerHour
}

// DiffInHoursWithAbs 相差多少小时（绝对值）
func (start Carbon) DiffInHoursWithAbs(end Carbon) int64 {
	diff := start.DiffInHours(end)
	if diff < 0 {
		diff = -diff
	}
	return diff
}

// DiffInMinutes 相差多少分钟
func (start Carbon) DiffInMinutes(end Carbon) int64 {
	return start.DiffInSeconds(end) / SecondsPerMinute
}

// DiffInMinutesWithAbs 相差多少分钟（绝对值）
func (start Carbon) DiffInMinutesWithAbs(end Carbon) int64 {
	diff := start.DiffInMinutes(end)
	if diff < 0 {
		diff = -diff
	}
	return diff
}

// DiffInSeconds 相差多少秒
func (start Carbon) DiffInSeconds(end Carbon) int64 {
	if start.Time.IsZero() && end.Time.IsZero() {
		return 0
	}
	if end.Time.IsZero() {
		return -start.ToTimestamp()
	}
	if start.Time.IsZero() {
		return end.ToTimestamp()
	}

	return end.ToTimestamp() - start.ToTimestamp()
}

// DiffInSecondsWithAbs 相差多少秒（绝对值）
func (start Carbon) DiffInSecondsWithAbs(end Carbon) int64 {
	diff := start.DiffInSeconds(end)
	if diff < 0 {
		diff = -diff
	}
	return diff
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
	return c.EndOfMonth().Time.In(c.Loc).Day()
}

// MonthOfYear 获取本年的第几月
func (c Carbon) MonthOfYear() int {
	if c.Time.IsZero() {
		return 0
	}
	return int(c.Time.In(c.Loc).Month())
}

// DayOfYear 获取本年的第几天
func (c Carbon) DayOfYear() int {
	if c.Time.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).YearDay()
}

// DayOfMonth 获取本月的第几天
func (c Carbon) DayOfMonth() int {
	if c.Time.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Day()
}

// DayOfWeek 获取本周的第几天
func (c Carbon) DayOfWeek() int {
	if c.Time.IsZero() {
		return 0
	}
	return int(c.Time.In(c.Loc).Weekday())
}

// WeekOfYear 获取本年的第几周
func (c Carbon) WeekOfYear() int {
	if c.Time.IsZero() {
		return 0
	}
	_, week := c.Time.In(c.Loc).ISOWeek()
	return week
}

// WeekOfMonth 获取本月的第几周
func (c Carbon) WeekOfMonth() int {
	if c.Time.IsZero() {
		return 0
	}
	day := c.Time.In(c.Loc).Day()
	if day < DaysPerWeek {
		return 1
	}
	return day%DaysPerWeek + 1
}

// Timezone 获取时区
func (c Carbon) Timezone() string {
	return c.Loc.String()
}

// Age 获取年龄
func (c Carbon) Age() int {
	if c.Time.IsZero() {
		return 0
	}
	if c.ToTimestamp() > Now().ToTimestamp() {
		return 0
	}
	age := time.Now().Year() - c.Time.In(c.Loc).Year()
	if int(time.Now().Month())*100+time.Now().Day() < int(c.Time.In(c.Loc).Month())*100+c.Time.In(c.Loc).Day() {
		age = age - 1
	}
	return age
}

// 获取当前年
func (c Carbon) Year() int {
	if c.Time.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Year()
}

// 获取当前季度
func (c Carbon) Quarter() int {
	if c.Time.IsZero() {
		return 0
	}
	month := c.Time.In(c.Loc).Month()
	quarter := 0
	switch {
	case month >= 1 && month <= 3:
		quarter = 1
	case month >= 4 && month <= 6:
		quarter = 2
	case month >= 7 && month <= 9:
		quarter = 3
	case month >= 10 && month <= 12:
		quarter = 4
	}
	return quarter
}

// 获取当前月
func (c Carbon) Month() int {
	if c.Time.IsZero() {
		return 0
	}
	return int(c.Time.In(c.Loc).Month())
}

// 获取当前日
func (c Carbon) Day() int {
	if c.Time.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Day()
}

// 获取当前小时
func (c Carbon) Hour() int {
	if c.Time.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Hour()
}

// 获取当前分钟数
func (c Carbon) Minute() int {
	if c.Time.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Minute()
}

// 获取当前秒数
func (c Carbon) Second() int {
	if c.Time.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Second()
}

// 获取当前毫秒数
func (c Carbon) Millisecond() int {
	if c.Time.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Nanosecond() / 1e6
}

// 获取当前微秒数
func (c Carbon) Microsecond() int {
	if c.Time.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Nanosecond() / 1e9
}

// 获取当前纳秒数
func (c Carbon) Nanosecond() int {
	if c.Time.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Nanosecond()
}

// IsZero 是否是零值
func (c Carbon) IsZero() bool {
	return c.Time.In(c.Loc).IsZero()
}

// IsNow 是否是当前时间
func (c Carbon) IsNow() bool {
	return c.ToTimestamp() == c.Now().ToTimestamp()
}

// IsFuture 是否是未来时间
func (c Carbon) IsFuture() bool {
	return c.ToTimestamp() > c.Now().ToTimestamp()
}

// IsPast 是否是过去时间
func (c Carbon) IsPast() bool {
	return c.ToTimestamp() < c.Now().ToTimestamp()
}

// IsLeapYear 是否是闰年
func (c Carbon) IsLeapYear() bool {
	year := c.Time.In(c.Loc).Year()
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

// IsLongYear 是否是长年
func (c Carbon) IsLongYear() bool {
	t := time.Date(c.Year(), time.December, 31, 0, 0, 0, 0, c.Loc)
	_, w := t.ISOWeek()
	return w == weeksPerLongYear
}

// IsJanuary 是否是一月
func (c Carbon) IsJanuary() bool {
	return c.Time.In(c.Loc).Month() == time.January
}

// IsMonday 是否是二月
func (c Carbon) IsFebruary() bool {
	return c.Time.In(c.Loc).Month() == time.February
}

// IsMarch 是否是三月
func (c Carbon) IsMarch() bool {
	return c.Time.In(c.Loc).Month() == time.March
}

// IsApril 是否是四月
func (c Carbon) IsApril() bool {
	return c.Time.In(c.Loc).Month() == time.April
}

// IsMay 是否是五月
func (c Carbon) IsMay() bool {
	return c.Time.In(c.Loc).Month() == time.May
}

// IsJune 是否是六月
func (c Carbon) IsJune() bool {
	return c.Time.In(c.Loc).Month() == time.June
}

// IsJuly 是否是七月
func (c Carbon) IsJuly() bool {
	return c.Time.In(c.Loc).Month() == time.July
}

// IsAugust 是否是八月
func (c Carbon) IsAugust() bool {
	return c.Time.In(c.Loc).Month() == time.August
}

// IsSeptember 是否是九月
func (c Carbon) IsSeptember() bool {
	return c.Time.In(c.Loc).Month() == time.September
}

// IsOctober 是否是十月
func (c Carbon) IsOctober() bool {
	return c.Time.In(c.Loc).Month() == time.October
}

// IsNovember 是否是十一月
func (c Carbon) IsNovember() bool {
	return c.Time.In(c.Loc).Month() == time.November
}

// IsDecember 是否是十二月
func (c Carbon) IsDecember() bool {
	return c.Time.In(c.Loc).Month() == time.December
}

// IsMonday 是否是周一
func (c Carbon) IsMonday() bool {
	return c.Time.In(c.Loc).Weekday() == time.Monday
}

// IsTuesday 是否是周二
func (c Carbon) IsTuesday() bool {
	return c.Time.In(c.Loc).Weekday() == time.Tuesday
}

// IsWednesday 是否是周三
func (c Carbon) IsWednesday() bool {
	return c.Time.In(c.Loc).Weekday() == time.Wednesday
}

// IsThursday 是否是周四
func (c Carbon) IsThursday() bool {
	return c.Time.In(c.Loc).Weekday() == time.Thursday
}

// IsFriday 是否是周五
func (c Carbon) IsFriday() bool {
	return c.Time.In(c.Loc).Weekday() == time.Friday
}

// IsSaturday 是否是周六
func (c Carbon) IsSaturday() bool {
	return c.Time.In(c.Loc).Weekday() == time.Saturday
}

// IsSunday 是否是周日
func (c Carbon) IsSunday() bool {
	return c.Time.In(c.Loc).Weekday() == time.Sunday
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
	return c.ToDateString() == Now().SubDay().ToDateString()
}

// IsToday 是否是今天
func (c Carbon) IsToday() bool {
	return c.ToDateString() == c.Now().ToDateString()
}

// IsTomorrow 是否是明天
func (c Carbon) IsTomorrow() bool {
	return c.ToDateString() == Now().AddDay().ToDateString()
}

// Compare 时间比较
func (c Carbon) Compare(operator string, t Carbon) bool {
	switch operator {
	case "=":
		return c.Time.Equal(t.Time)
	case "<>":
		return !c.Time.Equal(t.Time)
	case "!=":
		return !c.Time.Equal(t.Time)
	case ">":
		return c.Time.After(t.Time)
	case ">=":
		return c.Time.After(t.Time) || c.Time.Equal(t.Time)
	case "<":
		return c.Time.Before(t.Time)
	case "<=":
		return c.Time.Before(t.Time) || c.Time.Equal(t.Time)
	}

	return false
}

// Gt 大于
func (c Carbon) Gt(t Carbon) bool {
	return c.Time.After(t.Time)
}

// Lt 小于
func (c Carbon) Lt(t Carbon) bool {
	return c.Time.Before(t.Time)
}

// Eq 等于
func (c Carbon) Eq(t Carbon) bool {
	return c.Time.Equal(t.Time)
}

// Ne 不等于
func (c Carbon) Ne(t Carbon) bool {
	return !c.Time.Equal(t.Time)
}

// Gte 大于等于
func (c Carbon) Gte(t Carbon) bool {
	return c.Time.After(t.Time) || c.Time.Equal(t.Time)
}

// Lte 小于等于
func (c Carbon) Lte(t Carbon) bool {
	return c.Time.Before(t.Time) || c.Time.Equal(t.Time)
}

// Between 是否在两个时间之间(不包括这两个时间)
func (c Carbon) Between(start Carbon, end Carbon) bool {
	if c.Time.After(start.Time) && c.Time.Before(end.Time) {
		return true
	}
	return false
}

// BetweenIncludedStartTime 是否在两个时间之间(包括开始时间)
func (c Carbon) BetweenIncludedStartTime(start Carbon, end Carbon) bool {
	if (c.Time.After(start.Time) || c.Time.Equal(start.Time)) && c.Time.Before(end.Time) {
		return true
	}
	return false
}

// BetweenIncludedEndTime 是否在两个时间之间(包括结束时间)
func (c Carbon) BetweenIncludedEndTime(start Carbon, end Carbon) bool {
	if c.Time.After(start.Time) && (c.Time.Before(end.Time) || c.Time.Equal(end.Time)) {
		return true
	}
	return false
}

// BetweenIncludedBoth 是否在两个时间之间(包括这两个时间)
func (c Carbon) BetweenIncludedBoth(start Carbon, end Carbon) bool {
	if (c.Time.After(start.Time) || c.Time.Equal(start.Time)) && (c.Time.Before(end.Time) || c.Time.Equal(end.Time)) {
		return true
	}
	return false
}
