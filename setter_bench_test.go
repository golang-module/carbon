package carbon

import (
	"testing"
	"time"
)

func BenchmarkCarbon_SetTimezone(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SetTimezone(PRC)
	}
}

func BenchmarkCarbon_SetLocation(b *testing.B) {
	loc, _ := time.LoadLocation(PRC)
	for n := 0; n < b.N; n++ {
		SetLocation(loc)
	}
}

func BenchmarkCarbon_SetLocale(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SetLocale("en")
	}
}

func BenchmarkCarbon_SetLanguage(b *testing.B) {
	lang := NewLanguage()
	for n := 0; n < b.N; n++ {
		SetLanguage(lang)
	}
}

func BenchmarkCarbon_SetDateTime(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetDateTime(2020, 8, 5, 0, 0, 0)
	}
}

func BenchmarkCarbon_SetDateTimeMilli(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetDateTimeMilli(2020, 8, 5, 0, 0, 0, 0)
	}
}

func BenchmarkCarbon_SetDateTimeMicro(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetDateTimeMicro(2020, 8, 5, 0, 0, 0, 0)
	}
}

func BenchmarkCarbon_SetDateTimeNano(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetDateTimeNano(2020, 8, 5, 0, 0, 0, 0)
	}
}

func BenchmarkCarbon_SetDate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetDate(2020, 8, 5)
	}
}

func BenchmarkCarbon_SetDateMilli(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetDateMilli(2020, 8, 5, 0)
	}
}

func BenchmarkCarbon_SetDateMicro(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetDateMicro(2020, 8, 5, 0)
	}
}

func BenchmarkCarbon_SetDateNano(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetDateNano(2020, 8, 5, 0)
	}
}

func BenchmarkCarbon_SetTime(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetTime(13, 14, 15)
	}
}

func BenchmarkCarbon_SetTimeMilli(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetTimeMilli(13, 14, 15, 0)
	}
}

func BenchmarkCarbon_SetTimeMicro(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetTimeMicro(13, 14, 15, 0)
	}
}

func BenchmarkCarbon_SetTimeNano(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetTimeNano(13, 14, 15, 0)
	}
}

func BenchmarkCarbon_SetYear(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetYear(2020)
	}
}

func BenchmarkCarbon_SetYearNoOverflow(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetYearNoOverflow(2020)
	}
}

func BenchmarkCarbon_SetMonth(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetMonth(8)
	}
}

func BenchmarkCarbon_SetMonthNoOverflow(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetMonthNoOverflow(8)
	}
}

func BenchmarkCarbon_SetWeekStartsAt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetWeekStartsAt(Sunday)
	}
}

func BenchmarkCarbon_SetDay(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetDay(20)
	}
}

func BenchmarkCarbon_SetHour(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetHour(20)
	}
}

func BenchmarkCarbon_SetMinute(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetMinute(20)
	}
}

func BenchmarkCarbon_SetSecond(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetSecond(20)
	}
}

func BenchmarkCarbon_SetMillisecond(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetMillisecond(20)
	}
}

func BenchmarkCarbon_SetMicrosecond(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetMicrosecond(20)
	}
}

func BenchmarkCarbon_SetNanosecond(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCarbon().SetNanosecond(20)
	}
}
