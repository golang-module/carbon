package carbon

import (
	"testing"
)

func TestLanguage_SetLocale(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"en", true},
		{"zh-CN", true},
		{"zh-XX", false},
	}

	for _, v := range Tests {
		output := NewLanguage().SetLocale(v.input)

		if output == nil {
			if v.output == false {
				t.Errorf("Input %s, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLanguage_SetDir(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"lang", true},
		{"xxxx", false},
	}

	for _, v := range Tests {
		output := NewLanguage().SetDir(v.input)

		if output == nil {
			if v.output == false {
				t.Errorf("Input %s, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLanguage_SetResources(t *testing.T) {
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
