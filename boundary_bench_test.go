package carbon

import (
	"testing"
)

func BenchmarkCarbon_StartOfCentury(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.StartOfCentury()
	}
}

func BenchmarkCarbon_EndOfCentury(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.EndOfCentury()
	}
}

func BenchmarkCarbon_StartOfDecade(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.StartOfDecade()
	}
}

func BenchmarkCarbon_EndOfDecade(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.EndOfDecade()
	}
}

func BenchmarkCarbon_StartOfYear(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.StartOfYear()
	}
}

func BenchmarkCarbon_EndOfYear(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.EndOfYear()
	}
}

func BenchmarkCarbon_StartOfQuarter(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.StartOfQuarter()
	}
}

func BenchmarkCarbon_EndOfQuarter(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.EndOfQuarter()
	}
}

func BenchmarkCarbon_StartOfMonth(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.StartOfMonth()
	}
}

func BenchmarkCarbon_EndOfMonth(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.EndOfMonth()
	}
}

func BenchmarkCarbon_StartOfWeek(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.StartOfWeek()
	}
}

func BenchmarkCarbon_EndOfWeek(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.EndOfWeek()
	}
}

func BenchmarkCarbon_StartOfDay(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.StartOfDay()
	}
}

func BenchmarkCarbon_EndOfDay(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.EndOfDay()
	}
}

func BenchmarkCarbon_StartOfHour(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.StartOfHour()
	}
}

func BenchmarkCarbon_EndOfHour(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.EndOfHour()
	}
}

func BenchmarkCarbon_StartOfMinute(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.StartOfMinute()
	}
}

func BenchmarkCarbon_EndOfMinute(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.EndOfMinute()
	}
}

func BenchmarkCarbon_StartOfSecond(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.StartOfSecond()
	}
}

func BenchmarkCarbon_EndOfSecond(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.EndOfSecond()
	}
}
