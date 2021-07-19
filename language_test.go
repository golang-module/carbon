package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLanguage_SetLocale(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output error  // 期望输出值
	}{
		{"en", nil},
		{"zh-CN", nil},
	}

	for _, test := range tests {
		assert.ErrorIs(NewLanguage().SetLocale(test.input), test.output)
	}
}

func TestLanguage_SetDir(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output error  // 期望输出值
	}{
		{"lang", nil},
	}

	for _, test := range tests {
		assert.ErrorIs(NewLanguage().SetDir(test.input), test.output)
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
		input  Carbon // 输入值
		output string // 期望输出值
	}{
		{Now(), "just now"},
		{Now().AddYears(1), "in 1 yr"},
		{Now().SubYears(1), "1 yr ago"},
		{Now().AddYears(10), "in 10 yrs"},
		{Now().SubYears(10), "10 yrs ago"},

		{Now().AddMonths(1), "in 1 mo"},
		{Now().SubMonths(1), "1 mo ago"},
		{Now().AddMonths(10), "in 10 mos"},
		{Now().SubMonths(10), "10 mos ago"},

		{Now().AddDays(1), "in 1d"},
		{Now().SubDays(1), "1d ago"},
		{Now().AddDays(10), "in 1w"},
		{Now().SubDays(10), "1w ago"},
	}

	for _, test := range tests {
		c := test.input.SetLanguage(lang)
		assert.Nil(c.Error)
		assert.Equal(c.DiffForHumans(), test.output)
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
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},

		{"0000-00-00 00:00:00", ""},
		{"2021-08-05 13:14:15", ""},
	}

	for _, test := range tests {
		assert.Equal(Parse(test.input).SetLanguage(lang).DiffForHumans(), test.output)
	}

	for _, test := range tests {
		assert.Equal(Parse(test.input).SetLanguage(lang).Constellation(), test.output)
	}

	for _, test := range tests {
		assert.Equal(Parse(test.input).SetLanguage(lang).Season(), test.output)
	}

	for _, test := range tests {
		assert.Equal(Parse(test.input).SetLanguage(lang).ToWeekString(), test.output)
	}

	for _, test := range tests {
		assert.Equal(Parse(test.input).SetLanguage(lang).ToShortWeekString(), test.output)
	}

	for _, test := range tests {
		assert.Equal(Parse(test.input).SetLanguage(lang).ToMonthString(), test.output)
	}

	for _, test := range tests {
		assert.Equal(Parse(test.input).SetLanguage(lang).ToShortMonthString(), test.output)
	}
}
