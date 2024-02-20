package carbon

import "testing"

func BenchmarkCarbon_DiffInYears(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DiffInYears(Yesterday())
	}
}

func BenchmarkCarbon_DiffAbsInYears(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DiffAbsInYears(Yesterday())
	}
}

func BenchmarkCarbon_DiffInMonths(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DiffInMonths(Yesterday())
	}
}

func BenchmarkCarbon_DiffAbsInMonths(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DiffAbsInMonths(Yesterday())
	}
}

func BenchmarkCarbon_DiffInWeeks(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DiffInWeeks(Yesterday())
	}
}

func BenchmarkCarbon_DiffAbsInWeeks(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DiffAbsInWeeks(Yesterday())
	}
}

func BenchmarkCarbon_DiffInDays(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DiffInDays(Yesterday())
	}
}

func BenchmarkCarbon_DiffAbsInDays(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DiffAbsInDays(Yesterday())
	}
}

func BenchmarkCarbon_DiffInHours(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DiffInHours(Yesterday())
	}
}

func BenchmarkCarbon_DiffAbsInHours(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DiffAbsInHours(Yesterday())
	}
}

func BenchmarkCarbon_DiffInMinutes(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DiffInMinutes(Yesterday())
	}
}

func BenchmarkCarbon_DiffAbsInMinutes(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DiffAbsInMinutes(Yesterday())
	}
}

func BenchmarkCarbon_DiffInSeconds(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DiffInSeconds(Yesterday())
	}
}

func BenchmarkCarbon_DiffAbsInSeconds(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DiffAbsInSeconds(Yesterday())
	}
}

func BenchmarkCarbon_DiffInString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DiffInString(Yesterday())
	}
}

func BenchmarkCarbon_DiffAbsInString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DiffAbsInString(Yesterday())
	}
}

func BenchmarkCarbon_DiffInDuration(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DiffInDuration(Yesterday())
	}
}

func BenchmarkCarbon_DiffAbsInDuration(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DiffAbsInDuration(Yesterday())
	}
}

func BenchmarkCarbon_DiffForHumans(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DiffForHumans(Yesterday())
	}
}
