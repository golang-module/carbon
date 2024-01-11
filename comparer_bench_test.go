package carbon

import "testing"

func BenchmarkCarbon_IsDST(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsDST()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsDST()
	}
}

func BenchmarkCarbon_IsZero(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsZero()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsZero()
	}
}

func BenchmarkCarbon_IsValid(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsZero()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsZero()
	}
}

func BenchmarkCarbon_IsInvalid(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsInvalid()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsInvalid()
	}
}

func BenchmarkCarbon_IsAM(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsAM()
	}
}

func BenchmarkCarbon_IsPM(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsPM()
	}
}

func BenchmarkCarbon_IsNow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsNow()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsNow()
	}
}

func BenchmarkCarbon_IsFuture(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsFuture()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsFuture()
	}
}

func BenchmarkCarbon_IsPast(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsPast()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsPast()
	}
}

func BenchmarkCarbon_IsLeapYear(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsLeapYear()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsLeapYear()
	}
}

func BenchmarkCarbon_IsLongYear(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsLongYear()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsLongYear()
	}
}

func BenchmarkCarbon_IsJanuary(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsJanuary()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsJanuary()
	}
}

func BenchmarkCarbon_IsFebruary(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsFebruary()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsFebruary()
	}
}

func BenchmarkCarbon_IsMarch(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsMarch()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsMarch()
	}
}

func BenchmarkCarbon_IsApril(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsApril()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsApril()
	}
}

func BenchmarkCarbon_IsMay(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsMay()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsMay()
	}
}

func BenchmarkCarbon_IsJune(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsJune()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsJune()
	}
}

func BenchmarkCarbon_IsJuly(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsJuly()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsJuly()
	}
}

func BenchmarkCarbon_IsAugust(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsAugust()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsAugust()
	}
}

func BenchmarkCarbon_IsSeptember(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSeptember()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsSeptember()
	}
}

func BenchmarkCarbon_IsOctober(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsOctober()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsOctober()
	}
}

func BenchmarkCarbon_IsNovember(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsNovember()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsNovember()
	}
}

func BenchmarkCarbon_IsDecember(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsDecember()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsDecember()
	}
}

func BenchmarkCarbon_IsMonday(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsMonday()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsMonday()
	}
}

func BenchmarkCarbon_IsTuesday(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsTuesday()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsTuesday()
	}
}

func BenchmarkCarbon_IsWednesday(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsWednesday()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsWednesday()
	}
}

func BenchmarkCarbon_IsThursday(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsThursday()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsThursday()
	}
}

func BenchmarkCarbon_IsFriday(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsFriday()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsFriday()
	}
}

func BenchmarkCarbon_IsSaturday(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSaturday()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsSaturday()
	}
}

func BenchmarkCarbon_IsSunday(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSunday()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsSunday()
	}
}

func BenchmarkCarbon_IsWeekday(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsWeekday()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsWeekday()
	}
}

func BenchmarkCarbon_IsWeekend(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsWeekend()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsWeekend()
	}
}

func BenchmarkCarbon_IsYesterday(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsYesterday()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsYesterday()
	}
}

func BenchmarkCarbon_IsToday(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsToday()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsToday()
	}
}

func BenchmarkCarbon_IsTomorrow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsTomorrow()
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsTomorrow()
	}
}

func BenchmarkCarbon_IsSameCentury(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSameCentury(Yesterday())
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsSameCentury(Yesterday())
	}
}

func BenchmarkCarbon_IsSameDecade(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSameDecade(Yesterday())
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsSameDecade(Yesterday())
	}
}

func BenchmarkCarbon_IsSameYear(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSameYear(Yesterday())
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsSameYear(Yesterday())
	}
}

func BenchmarkCarbon_IsSameQuarter(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSameQuarter(Yesterday())
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsSameQuarter(Yesterday())
	}
}

func BenchmarkCarbon_IsSameMonth(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSameMonth(Yesterday())
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsSameMonth(Yesterday())
	}
}

func BenchmarkCarbon_IsSameDay(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSameDay(Yesterday())
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsSameDay(Yesterday())
	}
}

func BenchmarkCarbon_IsSameHour(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSameHour(Yesterday())
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsSameHour(Yesterday())
	}
}

func BenchmarkCarbon_IsSameMinute(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSameMinute(Yesterday())
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsSameMinute(Yesterday())
	}
}

func BenchmarkCarbon_IsSameSecond(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSameSecond(Yesterday())
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().IsSameSecond(Yesterday())
	}
}

func BenchmarkCarbon_Compare(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Compare(">", Yesterday())
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().Compare(">", Yesterday())
	}
}

func BenchmarkCarbon_Gt(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Gt(Yesterday())
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().Gt(Yesterday())
	}
}

func BenchmarkCarbon_Lt(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Lt(Yesterday())
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().Lt(Yesterday())
	}
}

func BenchmarkCarbon_Eq(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Eq(Yesterday())
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().Eq(Yesterday())
	}
}

func BenchmarkCarbon_Ne(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Ne(Yesterday())
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().Ne(Yesterday())
	}
}

func BenchmarkCarbon_Gte(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Gte(Yesterday())
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().Gte(Yesterday())
	}
}

func BenchmarkCarbon_Lte(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Lte(Yesterday())
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().Lte(Yesterday())
	}
}

func BenchmarkCarbon_Between(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Between(Tomorrow(), Yesterday())
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().Between(Tomorrow(), Yesterday())
	}
}

func BenchmarkCarbon_BetweenIncludedStart(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.BetweenIncludedStart(Tomorrow(), Yesterday())
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().BetweenIncludedStart(Tomorrow(), Yesterday())
	}
}

func BenchmarkCarbon_BetweenIncludedEnd(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.BetweenIncludedEnd(Tomorrow(), Yesterday())
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().BetweenIncludedEnd(Tomorrow(), Yesterday())
	}
}

func BenchmarkCarbon_BetweenIncludedBoth(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.BetweenIncludedBoth(Tomorrow(), Yesterday())
	}
	for n := 0; n < b.N; n++ {
		NewCarbon().BetweenIncludedBoth(Tomorrow(), Yesterday())
	}
}
