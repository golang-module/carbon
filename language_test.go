package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLanguage_SetLocale(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		output error  // 期望输出值
	}{
		{1, "en", nil},
		{2, "zh-CN", nil},
	}

	for _, test := range tests {
		assert.ErrorIs(NewLanguage().SetLocale(test.input), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLanguage_SetDir(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		output error  // 期望输出值
	}{
		{1, "lang", nil},
	}

	for _, test := range tests {
		assert.ErrorIs(NewLanguage().SetDir(test.input), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLanguage_SetResources1(t *testing.T) {
	assert := assert.New(t)

	lang := NewLanguage()
	resources := map[string]string{
		"seasons":  "spring|summer|autumn|winter",
		"year":     "1 yr|%d yrs",
		"month":    "1 mo|%d mos",
		"week":     "%dw",
		"day":      "%dd",
		"hour":     "%dh",
		"minute":   "%dm",
		"second":   "%ds",
		"now":      "just now",
		"ago":      "%s ago",
		"from_now": "in %s",
		"before":   "%s before",
		"after":    "%s after",
	}
	lang.SetResources(resources)

	tests := []struct {
		id     int    // 测试id
		input  Carbon // 输入值
		output string // 期望输出值
	}{
		{1, Now(), "just now"},
		{2, Now().AddYears(1), "in 1 yr"},
		{3, Now().SubYears(1), "1 yr ago"},
		{4, Now().AddYears(10), "in 10 yrs"},
		{5, Now().SubYears(10), "10 yrs ago"},

		{6, Now().AddMonths(1), "in 1 mo"},
		{7, Now().SubMonths(1), "1 mo ago"},
		{8, Now().AddMonths(10), "in 10 mos"},
		{9, Now().SubMonths(10), "10 mos ago"},

		{10, Now().AddDays(1), "in 1d"},
		{11, Now().SubDays(1), "1d ago"},
		{12, Now().AddDays(10), "in 1w"},
		{13, Now().SubDays(10), "1w ago"},
	}

	for _, test := range tests {
		c := test.input.SetLanguage(lang)
		assert.Nil(c.Error)
		assert.Equal(c.DiffForHumans(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLanguage_SetResources2(t *testing.T) {
	assert := assert.New(t)

	lang := NewLanguage()
	resources := map[string]string{
		"xxx": "xxxx",
	}
	lang.SetResources(resources)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		output string // 期望输出值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},

		{5, "0000-00-00 00:00:00", ""},
		{6, "2021-08-05 13:14:15", ""},
	}

	for _, test := range tests {
		assert.Equal(Parse(test.input).SetLanguage(lang).DiffForHumans(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}

	for _, test := range tests {
		assert.Equal(Parse(test.input).SetLanguage(lang).Constellation(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}

	for _, test := range tests {
		assert.Equal(Parse(test.input).SetLanguage(lang).Season(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}

	for _, test := range tests {
		assert.Equal(Parse(test.input).SetLanguage(lang).ToWeekString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}

	for _, test := range tests {
		assert.Equal(Parse(test.input).SetLanguage(lang).ToShortWeekString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}

	for _, test := range tests {
		assert.Equal(Parse(test.input).SetLanguage(lang).ToMonthString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}

	for _, test := range tests {
		assert.Equal(Parse(test.input).SetLanguage(lang).ToShortMonthString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}
