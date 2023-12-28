package carbon

import "testing"

func BenchmarkCarbon_IsDST(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsDST()
	}
}

func BenchmarkCarbon_IsZero(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsZero()
	}
}

func BenchmarkCarbon_IsValid(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsZero()
	}
}

func BenchmarkCarbon_IsInvalid(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsInvalid()
	}
}

func BenchmarkCarbon_IsNow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsNow()
	}
}

func BenchmarkCarbon_IsFuture(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsFuture()
	}
}

func BenchmarkCarbon_IsPast(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsPast()
	}
}

func BenchmarkCarbon_IsLeapYear(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsLeapYear()
	}
}

func BenchmarkCarbon_IsLongYear(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsLongYear()
	}
}

func BenchmarkCarbon_IsJanuary(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsJanuary()
	}
}

func BenchmarkCarbon_IsFebruary(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsFebruary()
	}
}

func BenchmarkCarbon_IsMarch(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsMarch()
	}
}

func BenchmarkCarbon_IsApril(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsApril()
	}
}

func BenchmarkCarbon_IsMay(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsMay()
	}
}

func BenchmarkCarbon_IsJune(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsJune()
	}
}

func BenchmarkCarbon_IsJuly(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsJuly()
	}
}

func BenchmarkCarbon_IsAugust(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsAugust()
	}
}

func BenchmarkCarbon_IsSeptember(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSeptember()
	}
}

func BenchmarkCarbon_IsOctober(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsOctober()
	}
}

func BenchmarkCarbon_IsNovember(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsNovember()
	}
}

func BenchmarkCarbon_IsDecember(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsDecember()
	}
}

func BenchmarkCarbon_IsMonday(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsMonday()
	}
}

func BenchmarkCarbon_IsTuesday(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsTuesday()
	}
}

func BenchmarkCarbon_IsWednesday(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsWednesday()
	}
}

func BenchmarkCarbon_IsThursday(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsThursday()
	}
}

func BenchmarkCarbon_IsFriday(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsFriday()
	}
}

func BenchmarkCarbon_IsSaturday(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSaturday()
	}
}

func BenchmarkCarbon_IsSunday(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSunday()
	}
}

func BenchmarkCarbon_IsWeekday(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsWeekday()
	}
}

func BenchmarkCarbon_IsWeekend(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsWeekend()
	}
}

func BenchmarkCarbon_IsYesterday(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsYesterday()
	}
}

func BenchmarkCarbon_IsToday(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsToday()
	}
}

func BenchmarkCarbon_IsTomorrow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsTomorrow()
	}
}

func BenchmarkCarbon_IsSameCentury(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSameCentury(Yesterday())
	}
}

func BenchmarkCarbon_IsSameDecade(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSameDecade(Yesterday())
	}
}

func BenchmarkCarbon_IsSameYear(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSameYear(Yesterday())
	}
}

func BenchmarkCarbon_IsSameQuarter(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSameQuarter(Yesterday())
	}
}

func BenchmarkCarbon_IsSameMonth(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSameMonth(Yesterday())
	}
}

func BenchmarkCarbon_IsSameDay(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSameDay(Yesterday())
	}
}

func BenchmarkCarbon_IsSameHour(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSameHour(Yesterday())
	}
}

func BenchmarkCarbon_IsSameMinute(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSameMinute(Yesterday())
	}
}

func BenchmarkCarbon_IsSameSecond(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSameSecond(Yesterday())
	}
}

func BenchmarkCarbon_Compare(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Compare(">", Yesterday())
	}
}

func BenchmarkCarbon_Gt(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Gt(Yesterday())
	}
}

func BenchmarkCarbon_Lt(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Lt(Yesterday())
	}
}

func BenchmarkCarbon_Eq(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Eq(Yesterday())
	}
}

func BenchmarkCarbon_Ne(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Ne(Yesterday())
	}
}

func BenchmarkCarbon_Gte(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Gte(Yesterday())
	}
}

func BenchmarkCarbon_Lte(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Lte(Yesterday())
	}
}

func BenchmarkCarbon_Between(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Between(Tomorrow(), Yesterday())
	}
}

func BenchmarkCarbon_BetweenIncludedStart(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.BetweenIncludedStart(Tomorrow(), Yesterday())
	}
}

func BenchmarkCarbon_BetweenIncludedEnd(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.BetweenIncludedEnd(Tomorrow(), Yesterday())
	}
}

func BenchmarkCarbon_BetweenIncludedBoth(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.BetweenIncludedBoth(Tomorrow(), Yesterday())
	}
}
