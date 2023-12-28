package carbon

import "testing"

func BenchmarkCarbon_NewLanguage(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewLanguage()
	}
}

func BenchmarkLanguage_SetLocale(b *testing.B) {
	l := NewLanguage()
	for n := 0; n < b.N; n++ {
		l.SetLocale("zh-CN")
	}
}

func BenchmarkLanguage_SetResources(b *testing.B) {
	l := NewLanguage()
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
	for n := 0; n < b.N; n++ {
		l.SetResources(resources)
	}
}
