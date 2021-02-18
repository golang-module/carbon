package carbon

import (
	"fmt"
	"testing"
)

var TimezoneTests = []struct {
	input    string // 输入值
	timezone string // 输入参数
	output   string // 期望输出值
}{
	{"2020-08-05 13:14:15", PRC, "2020-08-05 13:14:15"},
	{"2020-08-05", Tokyo, "2020-08-05 01:00:00"},
	{"2020-08-05", "Hangzhou", "panic"}, // 异常情况
}

func TestCarbon_SetTimezone1(t *testing.T) {
	for _, v := range TimezoneTests {
		output := SetTimezone(v.timezone).Parse(v.input)

		if output.Error != nil {
			fmt.Println("catch an exception in SetTimezone():", output.Error)
			return
		}

		if output.ToDateTimeString() != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output.ToDateTimeString())
		}
	}
}

func TestCarbon_SetTimezone2(t *testing.T) {
	for _, v := range TimezoneTests {
		output := SetTimezone(PRC).SetTimezone(v.timezone).Parse(v.input)

		if output.Error != nil {
			fmt.Println("catch an exception in SetTimezone():", output.Error)
			return
		}

		if output.ToDateTimeString() != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output.ToDateTimeString())
		}
	}
}

var LocaleTests = []struct {
	input  string // 输入参数
	output string // 期望输出值
}{
	{"en", "en"},
	{"zh-CN", "zh-CN"},
	{"XXXX", "panic"}, // 异常情况
}

func TestCarbon_SetLocale(t *testing.T) {
	for _, v := range LocaleTests {
		output := SetLocale(v.input)

		if output.Error != nil {
			fmt.Println("catch an exception in SetLocale():", output.Error)
			return
		}

		if output.Locale() != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output.ToDateTimeString())
		}
	}
}

func TestCarbon_SetLanguage1(t *testing.T) {
	lang := NewLanguage()
	resources := map[string]string{
		"day": "%dd",
	}
	if lang.SetLocale("en") == nil {
		lang.SetResources(resources)
	}

	Tests := []struct {
		input  Carbon // 输入值
		output string // 期望输出值
	}{
		{Now(), "just now"},
		{Now().AddYears(1), "1 year from now"},
		{Now().SubYears(1), "1 year ago"},
		{Now().AddYears(10), "10 years from now"},
		{Now().SubYears(10), "10 years ago"},

		{Now().AddMonths(1), "1 month from now"},
		{Now().SubMonths(1), "1 month ago"},
		{Now().AddMonths(10), "10 months from now"},
		{Now().SubMonths(10), "10 months ago"},

		{Now().AddDays(1), "1d from now"},
		{Now().SubDays(1), "1d ago"},
		{Now().AddDays(10), "1 week from now"},
		{Now().SubDays(10), "1 week ago"},
	}

	for _, v := range Tests {
		output := (v.input).SetLanguage(lang).DiffForHumans()

		if output != v.output {
			t.Errorf("Input time %s, expected %s, but got %s", v.input.ToDateTimeString(), v.output, output)
		}
	}
}

func TestCarbon_SetLanguage2(t *testing.T) {
	lang := NewLanguage()
	resources := map[string]string{
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

	Tests := []struct {
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

	for _, v := range Tests {
		output := (v.input).SetLanguage(lang).DiffForHumans()

		if output != v.output {
			t.Errorf("Input time %s, expected %s, but got %s", v.input.ToDateTimeString(), v.output, output)
		}
	}
}

func TestCarbon_SetYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		year   int    // 输入参数
		output string // 期望输出值
	}{
		{"2020-01-01", 2019, "2019-01-01"},
		{"2020-01-31", 2019, "2019-01-31"},
		{"2020-02-01", 2019, "2019-02-01"},
		{"2020-02-28", 2019, "2019-02-28"},
		{"2020-02-29", 2019, "2019-03-01"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SetYear(v.year).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_SetMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		month  int    // 输入参数
		output string // 期望输出值
	}{
		{"2020-01-01", 2, "2020-02-01"},
		{"2020-01-30", 2, "2020-03-01"},
		{"2020-01-31", 2, "2020-03-02"},
		{"2020-08-05", 2, "2020-02-05"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SetMonth(v.month).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_SetDay(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		day    int    // 输入参数
		output string // 期望输出值
	}{
		{"2020-01-01", 31, "2020-01-31"},
		{"2020-02-01", 31, "2020-03-02"},
		{"2020-02-28", 31, "2020-03-02"},
		{"2020-02-29", 31, "2020-03-02"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SetDay(v.day).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_SetHour(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		hour   int    // 输入参数
		output string // 期望输出值
	}{
		{"2020-08-05 13:14:15", 10, "2020-08-05 10:14:15"},
		{"2020-08-05 13:14:15", 24, "2020-08-06 00:14:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SetHour(v.hour).ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_SetMinute(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		minute int    // 输入参数
		output string // 期望输出值
	}{
		{"2020-08-05 13:14:15", 10, "2020-08-05 13:10:15"},
		{"2020-08-05 13:14:15", 60, "2020-08-05 14:00:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SetMinute(v.minute).ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_SetSecond(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		second int    // 输入参数
		output string // 期望输出值
	}{
		{"2020-08-05 13:14:15", 10, "2020-08-05 13:14:10"},
		{"2020-08-05 13:14:15", 60, "2020-08-05 13:15:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SetSecond(v.second).ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}
