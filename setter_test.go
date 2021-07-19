package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_SetTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		timezone string // 输入参数
		output   string // 期望输出值
	}{
		{1, "2020-08-05 13:14:15", PRC, "2020-08-05 13:14:15"},
		{2, "2020-08-05", Tokyo, "2020-08-05 00:00:00"},
	}

	for _, test := range tests {
		c := SetTimezone(test.timezone).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToDateTimeString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SetLanguage(t *testing.T) {
	lang := NewLanguage()
	resources := map[string]string{
		"seasons": "spring|summer|autumn|winter",
	}

	if lang.SetLocale("en") == nil {
		lang.SetResources(resources)
	}

	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		output string // 期望输出值
	}{
		{1, "", ""},
		{2, "2020-08-05", "summer"},
	}

	for _, test := range tests {
		c := SetLanguage(lang).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Season(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}

	for _, test := range tests {
		c := Parse(test.input).SetLanguage(lang)
		assert.Nil(c.Error)
		assert.Equal(c.Season(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SetYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		year   int    // 输入参数
		output string // 期望输出值
	}{
		{1, "2020-01-01", 2019, "2019-01-01"},
		{2, "2020-01-31", 2019, "2019-01-31"},
		{3, "2020-02-01", 2019, "2019-02-01"},
		{4, "2020-02-28", 2019, "2019-02-28"},
		{5, "2020-02-29", 2019, "2019-03-01"},
	}

	for _, test := range tests {
		c := Parse(test.input).SetYear(test.year)
		assert.Nil(c.Error)
		assert.Equal(c.ToDateString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SetMonth(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		month  int    // 输入参数
		output string // 期望输出值
	}{
		{1, "2020-01-01", 2, "2020-02-01"},
		{2, "2020-01-30", 2, "2020-03-01"},
		{3, "2020-01-31", 2, "2020-03-02"},
		{4, "2020-08-05", 2, "2020-02-05"},
	}

	for _, test := range tests {
		c := Parse(test.input).SetMonth(test.month)
		assert.Nil(c.Error)
		assert.Equal(c.ToDateString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SetDay(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		day    int    // 输入参数
		output string // 期望输出值
	}{
		{1, "2020-01-01", 31, "2020-01-31"},
		{2, "2020-02-01", 31, "2020-03-02"},
		{3, "2020-02-28", 31, "2020-03-02"},
		{4, "2020-02-29", 31, "2020-03-02"},
	}

	for _, test := range tests {
		c := Parse(test.input).SetDay(test.day)
		assert.Nil(c.Error)
		assert.Equal(c.ToDateString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SetHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		hour   int    // 输入参数
		output string // 期望输出值
	}{
		{1, "2020-08-05 13:14:15", 10, "2020-08-05 10:14:15"},
		{2, "2020-08-05 13:14:15", 24, "2020-08-06 00:14:15"},
	}

	for _, test := range tests {
		c := Parse(test.input).SetHour(test.hour)
		assert.Nil(c.Error)
		assert.Equal(c.ToDateTimeString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SetMinute(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		minute int    // 输入参数
		output string // 期望输出值
	}{
		{1, "2020-08-05 13:14:15", 10, "2020-08-05 13:10:15"},
		{2, "2020-08-05 13:14:15", 60, "2020-08-05 14:00:15"},
	}

	for _, test := range tests {
		c := Parse(test.input).SetMinute(test.minute)
		assert.Nil(c.Error)
		assert.Equal(c.ToDateTimeString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SetSecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		second int    // 输入参数
		output int    // 期望输出值
	}{
		{1, "2020-08-05 13:14:15", 10, 10},
		{2, "2020-08-05 13:14:15", 59, 59},
	}

	for _, test := range tests {
		c := Parse(test.input).SetSecond(test.second)
		assert.Nil(c.Error)
		assert.Equal(c.Second(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SetMillisecond(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		id          int    // 测试id
		input       string // 输入值
		millisecond int    // 输入参数
		output      int    // 期望输出值
	}{
		{1, "2020-08-05 13:14:15", 100, 100},
		{2, "2020-08-05 13:14:15", 999, 999},
	}

	for _, test := range tests {
		c := Parse(test.input).SetMillisecond(test.millisecond)
		assert.Nil(c.Error)
		assert.Equal(c.Millisecond(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SetMicrosecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id          int    // 测试id
		input       string // 输入值
		microsecond int    // 输入参数
		output      int    // 期望输出值
	}{
		{1, "2020-08-05 13:14:15", 100000, 100000},
		{2, "2020-08-05 13:14:15", 999999, 999999},
	}

	for _, test := range tests {
		c := Parse(test.input).SetMicrosecond(test.microsecond)
		assert.Nil(c.Error)
		assert.Equal(c.Microsecond(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SetNanosecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id         int    // 测试id
		input      string // 输入值
		nanosecond int    // 输入参数
		output     int    // 期望输出值
	}{
		{1, "2020-08-05 13:14:15", 100000000, 100000000},
		{2, "2020-08-05 13:14:15", 999999999, 999999999},
	}

	for _, test := range tests {
		c := Parse(test.input).SetNanosecond(test.nanosecond)
		assert.Nil(c.Error)
		assert.Equal(c.Nanosecond(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}
