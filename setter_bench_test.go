package carbon

import (
	"testing"
	"time"
)

func BenchmarkCarbon_SetWeekStartsAt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SetWeekStartsAt(Sunday)
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetWeekStartsAt(Sunday)
	}
}

func BenchmarkCarbon_SetTimezone(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SetTimezone(PRC)
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetTimezone(PRC)
	}
}

func BenchmarkCarbon_SetLocation(b *testing.B) {
	loc, _ := time.LoadLocation(PRC)
	for n := 0; n < b.N; n++ {
		SetLocation(loc)
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetLocation(loc)
	}
}

func BenchmarkCarbon_SetLocale(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SetLocale("en")
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetLocale("en")
	}
}

func BenchmarkCarbon_SetDateTime(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetDateTime(2020, 8, 5, 0, 0, 0)
	}
}

func BenchmarkCarbon_SetDateTimeMilli(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetDateTimeMilli(2020, 8, 5, 0, 0, 0, 0)
	}
}

func BenchmarkCarbon_SetDateTimeMicro(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetDateTimeMicro(2020, 8, 5, 0, 0, 0, 0)
	}
}

func BenchmarkCarbon_SetDateTimeNano(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetDateTimeNano(2020, 8, 5, 0, 0, 0, 0)
	}
}

func BenchmarkCarbon_SetDate(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetDate(2020, 8, 5)
	}
}

func BenchmarkCarbon_SetDateMilli(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetDateMilli(2020, 8, 5, 0)
	}
}

func BenchmarkCarbon_SetDateMicro(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetDateMicro(2020, 8, 5, 0)
	}
}

func BenchmarkCarbon_SetDateNano(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetDateNano(2020, 8, 5, 0)
	}
}

func BenchmarkCarbon_SetTime(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetTime(13, 14, 15)
	}
}

func BenchmarkCarbon_SetTimeMilli(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetTimeMilli(13, 14, 15, 0)
	}
}

func BenchmarkCarbon_SetTimeMicro(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetTimeMicro(13, 14, 15, 0)
	}
}

func BenchmarkCarbon_SetTimeNano(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetTimeNano(13, 14, 15, 0)
	}
}

func BenchmarkCarbon_SetYear(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetYear(2020)
	}
}

func BenchmarkCarbon_SetYearNoOverflow(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetYearNoOverflow(2020)
	}
}

func BenchmarkCarbon_SetMonth(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetMonth(8)
	}
}

func BenchmarkCarbon_SetMonthNoOverflow(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetMonthNoOverflow(8)
	}
}

func BenchmarkCarbon_SetDay(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetDay(20)
	}
}

func BenchmarkCarbon_SetHour(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetHour(20)
	}
}

func BenchmarkCarbon_SetMinute(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetMinute(20)
	}
}

func BenchmarkCarbon_SetSecond(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetSecond(20)
	}
}

func BenchmarkCarbon_SetMillisecond(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetMillisecond(20)
	}
}

func BenchmarkCarbon_SetMicrosecond(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetMicrosecond(20)
	}
}

func BenchmarkCarbon_SetNanosecond(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetNanosecond(20)
	}
}
