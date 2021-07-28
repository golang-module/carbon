package carbon

import (
	"bytes"
	"strconv"
	"strings"
	"time"
)

// ToTimestamp ToTimestampWithSecond的简称
// ToTimestamp alias for ToTimestampWithSecond
func (c Carbon) ToTimestamp() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.ToTimestampWithSecond()
}

// ToTimestampWithSecond output timestamp with second
// ToTimestampWithSecond 输出秒级时间戳
func (c Carbon) ToTimestampWithSecond() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.Time.In(c.Loc).Unix()
}

// ToTimestampWithMillisecond output timestamp with millisecond
// ToTimestampWithMillisecond 输出毫秒级时间戳
func (c Carbon) ToTimestampWithMillisecond() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.Time.UnixNano() / int64(time.Millisecond)
}

// ToTimestampWithMicrosecond output timestamp with microsecond
// ToTimestampWithMicrosecond 输出微秒级时间戳
func (c Carbon) ToTimestampWithMicrosecond() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.Time.UnixNano() / int64(time.Microsecond)
}

// ToTimestampWithNanosecond output timestamp with nanosecond
// ToTimestampWithNanosecond 输出纳秒级时间戳
func (c Carbon) ToTimestampWithNanosecond() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.Time.UnixNano()
}

// String implement Stringer interface
// String 实现 Stringer 接口
func (c Carbon) String() string {
	if c.IsInvalid() {
		return ""
	}
	return c.ToDateTimeString()
}

// ToString output "2006-01-02 15:04:05.999999999 -0700 MST" format string
// ToString 输出"2006-01-02 15:04:05.999999999 -0700 MST"格式字符串
func (c Carbon) ToString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).String()
}

// ToMonthString output month format string, support i18n
// ToMonthString 输出完整月份字符串，支持i18n
func (c Carbon) ToMonthString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	if len(c.Lang.resources) == 0 {
		c.Lang.SetLocale(defaultLocale)
	}
	if months, ok := c.Lang.resources["months"]; ok {
		slice := strings.Split(months, "|")
		return slice[c.Month()-1]
	}
	return ""
}

// ToShortMonthString output short month format string, support i18n
// ToShortMonthString 输出缩写月份字符串，支持i18n
func (c Carbon) ToShortMonthString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	if len(c.Lang.resources) == 0 {
		c.Lang.SetLocale(defaultLocale)
	}
	if months, ok := c.Lang.resources["months_short"]; ok {
		slice := strings.Split(months, "|")
		return slice[c.Month()-1]
	}
	return ""
}

// ToWeekString output week format string, support i18n
// ToWeekString 输出完整星期字符串，支持i18n
func (c Carbon) ToWeekString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	if len(c.Lang.resources) == 0 {
		c.Lang.SetLocale(defaultLocale)
	}
	if months, ok := c.Lang.resources["weeks"]; ok {
		slice := strings.Split(months, "|")
		return slice[c.Week()]
	}
	return ""
}

// ToShortWeekString output short week format string, support i18n
// ToShortWeekString 输出缩写星期字符串，支持i18n
func (c Carbon) ToShortWeekString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	if len(c.Lang.resources) == 0 {
		c.Lang.SetLocale(defaultLocale)
	}
	if months, ok := c.Lang.resources["weeks_short"]; ok {
		slice := strings.Split(months, "|")
		return slice[c.Week()]
	}
	return ""
}

// ToDayDateTimeString output day, date and time format string
// ToDayDateTimeString 输出天数日期时间字符串
func (c Carbon) ToDayDateTimeString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(DayDateTimeFormat)
}

// ToDateTimeString output date and time format string
// ToDateTimeString 输出日期时间字符串
func (c Carbon) ToDateTimeString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(DateTimeFormat)
}

// ToShortDateTimeString output short date and time format string
// ToShortDateTimeString 输出简写日期时间字符串
func (c Carbon) ToShortDateTimeString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(ShortDateTimeFormat)
}

// ToDateString output date format string
// ToDateString 输出日期字符串
func (c Carbon) ToDateString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(DateFormat)
}

// ToShortDateString output short date format string
// ToShortDateString 输出简写日期字符串
func (c Carbon) ToShortDateString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(ShortDateFormat)
}

// ToTimeString output time format string
// ToTimeString 输出时间字符串
func (c Carbon) ToTimeString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(TimeFormat)
}

// ToShortTimeString output short time format string
// ToShortTimeString 输出简写时间字符串
func (c Carbon) ToShortTimeString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(ShortTimeFormat)
}

// ToAtomString output ATOM format string
// ToAtomString 输出 ATOM 格式字符串
func (c Carbon) ToAtomString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	return c.ToRfc3339String(timezone...)
}

// ToAnsicString output ANSIC format string
// ToAnsicString 输出 ANSIC 格式字符串
func (c Carbon) ToAnsicString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(AnsicFormat)
}

// ToCookieString output COOKIE format string
// ToCookieString 输出 COOKIE 格式字符串
func (c Carbon) ToCookieString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(CookieFormat)
}

// ToRssString output RSS format string
// ToRssString 输出 RSS 格式字符串
func (c Carbon) ToRssString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(RssFormat)
}

// ToW3cString output W3C format string
// ToW3cString 输出 W3C 格式字符串
func (c Carbon) ToW3cString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	return c.ToRfc3339String(timezone...)
}

// ToUnixDateString output unix date format string
// ToUnixDateString 输出 UnixDate 格式字符串
func (c Carbon) ToUnixDateString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(UnixDateFormat)
}

// ToRubyDateString output ruby date format string
// ToRubyDateString 输出 RubyDate 格式字符串
func (c Carbon) ToRubyDateString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(RubyDateFormat)
}

// ToKitchenString output KITCHEN format string
// ToKitchenString 输出 KITCHEN 格式字符串
func (c Carbon) ToKitchenString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(KitchenFormat)
}

// ToIso8601String output ISO8601 format string
// ToIso8601String 输出 ISO8601 格式字符串
func (c Carbon) ToIso8601String(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(Iso8601Format)
}

// ToRfc822String output RFC822 format string
// ToRfc822String 输出 RFC822 格式字符串
func (c Carbon) ToRfc822String(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(RFC822Format)
}

// ToRfc822zString output RFC822Z format string
// ToRfc822zString 输出 RFC822Z 格式字符串
func (c Carbon) ToRfc822zString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(RFC822ZFormat)
}

// ToRfc850String output RFC850 format string
// ToRfc850String 输出 RFC850 格式字符串
func (c Carbon) ToRfc850String(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(RFC850Format)
}

// ToRfc1036String output RFC1036 format string
// ToRfc1036String 输出 RFC1036 格式字符串
func (c Carbon) ToRfc1036String(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(RFC1036Format)
}

// ToRfc1123String output RFC1123 format string
// ToRfc1123String 输出 RFC1123 格式字符串
func (c Carbon) ToRfc1123String(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(RFC1123Format)
}

// ToRfc1123zString output RFC1123z format string
// ToRfc1123ZString 输出 RFC1123z 格式字符串
func (c Carbon) ToRfc1123zString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(RFC1123ZFormat)
}

// ToRfc2822String output RFC2822 format string
// ToRfc2822String 输出 RFC2822 格式字符串
func (c Carbon) ToRfc2822String(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(RFC2822Format)
}

// ToRfc3339String output RFC3339 format string
// ToRfc3339String 输出 RFC3339 格式字符串
func (c Carbon) ToRfc3339String(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(RFC3339Format)
}

// ToRfc7231String output RFC7231 format string
// ToRfc7231String 输出 RFC7231 格式字符串
func (c Carbon) ToRfc7231String(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(RFC7231Format)
}

// ToLayoutString output string by layout string
// ToLayoutString 输出指定布局的时间字符串
func (c Carbon) ToLayoutString(layout string, timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	return c.Time.In(c.Loc).Format(layout)
}

// Layout alias for ToLayoutString
// Layout ToLayoutString的简写
func (c Carbon) Layout(layout string, timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	return c.ToLayoutString(layout, timezone...)
}

// ToFormatString output string by format string
// ToFormatString 输出指定格式的时间字符串
func (c Carbon) ToFormatString(format string, timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	if len(timezone) == 1 {
		loc, err := getLocationByTimezone(timezone[0])
		c.Loc = loc
		c.Error = err
	}
	runes := []rune(format)
	buffer := bytes.NewBuffer(nil)
	for i := 0; i < len(runes); i++ {
		if layout, ok := formats[byte(runes[i])]; ok {
			buffer.WriteString(c.Time.In(c.Loc).Format(layout))
		} else {
			switch runes[i] {
			case '\\': // 原样输出，不解析
				buffer.WriteRune(runes[i+1])
				i++
				continue
			case 'W': // ISO-8601 格式数字表示的年份中的第几周，取值范围 1-52
				buffer.WriteString(strconv.Itoa(c.WeekOfYear()))
			case 'N': // ISO-8601 格式数字表示的星期中的第几天，取值范围 1-7
				buffer.WriteString(strconv.Itoa(c.DayOfWeek()))
			case 'S': // 月份中第几天的英文缩写后缀，如st, nd, rd, th
				suffix := "th"
				switch c.Day() {
				case 1, 21, 31:
					suffix = "st"
				case 2, 22:
					suffix = "nd"
				case 3, 23:
					suffix = "rd"
				}
				buffer.WriteString(suffix)
			case 'L': // 是否为闰年，如果是闰年为 1，否则为 0
				if c.IsLeapYear() {
					buffer.WriteString("1")
				} else {
					buffer.WriteString("0")
				}
			case 'G': // 数字表示的小时，24 小时格式，没有前导零，取值范围 0-23
				buffer.WriteString(strconv.Itoa(c.Hour()))
			case 'U': // 秒级时间戳，如 1611818268
				buffer.WriteString(strconv.FormatInt(c.ToTimestamp(), 10))
			case 'u': // 数字表示的毫秒，如 999
				buffer.WriteString(strconv.Itoa(c.Millisecond()))
			case 'w': // 数字表示的星期中的第几天，取值范围 0-6
				buffer.WriteString(strconv.Itoa(c.DayOfWeek() - 1))
			case 't': // 指定的月份有几天，取值范围 28-31
				buffer.WriteString(strconv.Itoa(c.DaysInMonth()))
			case 'z': // 年份中的第几天，取值范围 0-365
				buffer.WriteString(strconv.Itoa(c.DayOfYear() - 1))
			case 'e': // 当前位置，如 UTC，GMT，Atlantic/Azores
				buffer.WriteString(c.Location())
			case 'Q': // 当前季度，取值范围 1-4
				buffer.WriteString(strconv.Itoa(c.Quarter()))
			case 'C': // 当前世纪数，取值范围 0-99
				buffer.WriteString(strconv.Itoa(c.Century()))
			default:
				buffer.WriteRune(runes[i])
			}
		}
	}
	return buffer.String()
}

// Format alias for ToFormatString
// Format ToFormatString的简写
func (c Carbon) Format(format string, timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	return c.ToFormatString(format, timezone...)
}
