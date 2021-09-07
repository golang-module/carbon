package carbon

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestCarbon_SetTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		timezone string // 输入参数
		expected string // 期望值
	}{
		{"0000-00-00 00:00:00", PRC, ""},
		{"2020-08-05 13:14:15", PRC, "2020-08-05 13:14:15"},
		{"2020-08-05 13:14:15", Tokyo, "2020-08-05 12:14:15"},
		{"2020-08-05 13:14:15", London, "2020-08-05 20:14:15"},
	}

	for index, test := range tests {
		c := SetTimezone(test.timezone).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_SetTimezone(t *testing.T) {
	timezone := "xxx"
	c1 := SetTimezone(timezone)
	assert.NotNil(t, c1.Error, "It should catch an exception in SetTimezone()")

	c2 := Now(timezone).SetTimezone(timezone)
	assert.NotNil(t, c2.Error, "It should catch an exception in SetTimezone()")
}

func TestCarbon_SetLocale(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		locale   string // 输入参数
		expected string // 期望值
	}{
		{"0000-00-00", "en", ""},
		{"2020-08-05", "en", "August"},
		{"2020-08-05", "jp", "はちがつ"},
		{"2020-08-05", "zh-CN", "八月"},
	}

	for index, test := range tests {
		c := SetLocale(test.locale).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToMonthString(PRC), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input).SetLocale(test.locale)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToMonthString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_SetLocale(t *testing.T) {
	locale := "xxx"

	c1 := SetLocale(locale)
	assert.NotNil(t, c1.Error, "It should catch an exception in SetLocale()")

	c2 := SetLocale(locale).SetLocale(PRC)
	assert.NotNil(t, c2.Error, "It should catch an exception in SetLocale()")
}

func TestCarbon_SetLanguage(t *testing.T) {
	lang := NewLanguage()
	resources := map[string]string{
		"seasons": "spring|summer|autumn|winter",
	}
	lang.SetLocale("en")
	if lang.Error == nil {
		lang.SetResources(resources)
	}

	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"2020-08-05", "summer"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetLanguage(lang)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Season(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := SetLanguage(lang).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Season(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_SetLanguage(t *testing.T) {
	lang, locale := NewLanguage(), "xxx"
	lang.SetLocale(locale)

	c1 := SetLanguage(lang)
	assert.NotNil(t, c1.Error, "It should catch an exception in SetLanguage()")

	c2 := SetLocale(locale).SetLanguage(lang)
	assert.NotNil(t, c2.Error, "It should catch an exception in SetLanguage()")
}

func TestCarbon_SetYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		year     int    // 输入参数
		expected string // 期望值
	}{
		{"2020-01-01", 2019, "2019-01-01"},
		{"2020-01-31", 2019, "2019-01-31"},
		{"2020-02-01", 2019, "2019-02-01"},
		{"2020-02-28", 2019, "2019-02-28"},
		{"2020-02-29", 2019, "2019-03-01"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetYear(test.year)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetYearNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		year     int    // 输入参数
		expected string // 期望值
	}{
		{"2020-01-01", 2019, "2019-01-01"},
		{"2020-01-31", 2019, "2019-01-31"},
		{"2020-02-01", 2019, "2019-02-01"},
		{"2020-02-28", 2019, "2019-02-28"},
		{"2020-02-29", 2019, "2019-02-28"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetYearNoOverflow(test.year)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_SetYear(t *testing.T) {
	value, year := "2020-08-50", 2020
	c := Parse(value).SetYear(year)
	assert.NotNil(t, c.Error, "It should catch an exception in SetYear()")
}

func TestError_SetYearNoOverflow(t *testing.T) {
	value, year := "2020-08-50", 2020
	c := Parse(value).SetYearNoOverflow(year)
	assert.NotNil(t, c.Error, "It should catch an exception in SetYearNoOverflow()")
}

func TestCarbon_SetMonth(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		month    int    // 输入参数
		expected string // 期望值
	}{
		{"2020-01-01", 2, "2020-02-01"},
		{"2020-01-30", 2, "2020-03-01"},
		{"2020-01-31", 2, "2020-03-02"},
		{"2020-08-05", 2, "2020-02-05"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetMonth(test.month)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetMonthNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		month    int    // 输入参数
		expected string // 期望值
	}{
		{"2020-01-01", 2, "2020-02-01"},
		{"2020-01-30", 2, "2020-02-29"},
		{"2020-01-31", 2, "2020-02-29"},
		{"2020-08-05", 2, "2020-02-05"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetMonthNoOverflow(test.month)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_SetMonth(t *testing.T) {
	value, month := "2020-08-50", 12
	c := Parse(value).SetMonth(month)
	assert.NotNil(t, c.Error, "It should catch an exception in SetMonth()")
}

func TestError_SetMonthNoOverflow(t *testing.T) {
	value, month := "2020-08-50", 12
	c := Parse(value).SetMonthNoOverflow(month)
	assert.NotNil(t, c.Error, "It should catch an exception in SetMonthNoOverflow()")
}

func TestCarbon_SetWeekStartsAt(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		week     string // 输入参数
		expected string // 期望值
	}{
		{"", Sunday, ""},
		{"0000-00-00 00:00:00", Sunday, ""},
		{"", Monday, ""},
		{"0000-00-00 00:00:00", Monday, ""},

		{"2021-06-13", Sunday, "2021-06-13 00:00:00"},
		{"2021-06-14", Sunday, "2021-06-13 00:00:00"},
		{"2021-06-18", Sunday, "2021-06-13 00:00:00"},

		{"2021-06-13", Monday, "2021-06-07 00:00:00"},
		{"2021-06-14", Monday, "2021-06-14 00:00:00"},
		{"2021-06-18", Monday, "2021-06-14 00:00:00"},

		{"2021-06-13", Tuesday, "2021-06-08 00:00:00"},
		{"2021-06-14", Tuesday, "2021-06-08 00:00:00"},
		{"2021-06-18", Tuesday, "2021-06-15 00:00:00"},

		{"2021-06-13", Wednesday, "2021-06-09 00:00:00"},
		{"2021-06-14", Wednesday, "2021-06-09 00:00:00"},
		{"2021-06-18", Wednesday, "2021-06-16 00:00:00"},

		{"2021-06-13", Thursday, "2021-06-10 00:00:00"},
		{"2021-06-14", Thursday, "2021-06-10 00:00:00"},
		{"2021-06-18", Thursday, "2021-06-17 00:00:00"},

		{"2021-06-13", Friday, "2021-06-11 00:00:00"},
		{"2021-06-14", Friday, "2021-06-11 00:00:00"},
		{"2021-06-18", Friday, "2021-06-18 00:00:00"},

		{"2021-06-13", Saturday, "2021-06-12 00:00:00"},
		{"2021-06-14", Saturday, "2021-06-12 00:00:00"},
		{"2021-06-18", Saturday, "2021-06-12 00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetWeekStartsAt(test.week).StartOfWeek()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetDay(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		day      int    // 输入参数
		expected string // 期望值
	}{
		{"2020-01-01", 31, "2020-01-31"},
		{"2020-02-01", 31, "2020-03-02"},
		{"2020-02-28", 31, "2020-03-02"},
		{"2020-02-29", 31, "2020-03-02"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetDay(test.day)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_SetDay(t *testing.T) {
	value, day := "2020-08-50", 30
	c := Parse(value).SetDay(day)
	assert.NotNil(t, c.Error, "It should catch an exception in SetDay()")
}

func TestCarbon_SetHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		hour     int    // 输入参数
		expected string // 期望值
	}{
		{"2020-08-05 13:14:15", 10, "2020-08-05 10:14:15"},
		{"2020-08-05 13:14:15", 24, "2020-08-06 00:14:15"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetHour(test.hour)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_SetHour(t *testing.T) {
	value, hour := "2020-08-50", 12
	c := Parse(value).SetHour(hour)
	assert.NotNil(t, c.Error, "It should catch an exception in SetHour()")
}

func TestCarbon_SetMinute(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		minute   int    // 输入参数
		expected string // 期望值
	}{
		{"2020-08-05 13:14:15", 10, "2020-08-05 13:10:15"},
		{"2020-08-05 13:14:15", 60, "2020-08-05 14:00:15"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetMinute(test.minute)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_SetMinute(t *testing.T) {
	value, minute := "2020-08-50", 30
	c := Parse(value).SetMinute(minute)
	assert.NotNil(t, c.Error, "It should catch an exception in SetMinute()")
}

func TestCarbon_SetSecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		second   int    // 输入参数
		expected int    // 期望值
	}{
		{"2020-08-05 13:14:15", 10, 10},
		{"2020-08-05 13:14:15", 59, 59},
	}

	for index, test := range tests {
		c := Parse(test.input).SetSecond(test.second)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Second(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_SetSecond(t *testing.T) {
	value, second := "2020-08-50", 30
	c := Parse(value).SetSecond(second)
	assert.NotNil(t, c.Error, "It should catch an exception in SetSecond()")
}

func TestCarbon_SetMillisecond(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input       string // 输入值
		millisecond int    // 输入参数
		expected    int    // 期望值
	}{
		{"2020-08-05 13:14:15", 100, 100},
		{"2020-08-05 13:14:15", 999, 999},
	}

	for index, test := range tests {
		c := Parse(test.input).SetMillisecond(test.millisecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Millisecond(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_SetMillisecond(t *testing.T) {
	value, millisecond := "2020-08-50", 100
	c := Parse(value).SetMillisecond(millisecond)
	assert.NotNil(t, c.Error, "It should catch an exception in SetMillisecond()")
}

func TestCarbon_SetMicrosecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input       string // 输入值
		microsecond int    // 输入参数
		expected    int    // 期望值
	}{
		{"2020-08-05 13:14:15", 100000, 100000},
		{"2020-08-05 13:14:15", 999999, 999999},
	}

	for index, test := range tests {
		c := Parse(test.input).SetMicrosecond(test.microsecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Microsecond(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_SetMicrosecond(t *testing.T) {
	value, microsecond := "2020-08-50", 100000
	c := Parse(value).SetMicrosecond(microsecond)
	assert.NotNil(t, c.Error, "It should catch an exception in SetMicrosecond()")
}

func TestCarbon_SetNanosecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input      string // 输入值
		nanosecond int    // 输入参数
		expected   int    // 期望值
	}{
		{"2020-08-05 13:14:15", 100000000, 100000000},
		{"2020-08-05 13:14:15", 999999999, 999999999},
	}

	for index, test := range tests {
		c := Parse(test.input).SetNanosecond(test.nanosecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Nanosecond(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_SetNanosecond(t *testing.T) {
	value, nanosecond := "2020-08-50", 100000000
	c := Parse(value).SetNanosecond(nanosecond)
	assert.NotNil(t, c.Error, "It should catch an exception in SetNanosecond()")
}
