package carbon

import "testing"

func BenchmarkCarbon_Now(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Now()
	}
}

func BenchmarkCarbon_Yesterday(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Yesterday()
	}
}

func BenchmarkCarbon_Tomorrow(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Tomorrow()
	}
}

func BenchmarkCarbon_AddDuration(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddDuration("1s")
	}
}

func BenchmarkCarbon_SubDuration(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubDuration("1s")
	}
}

func BenchmarkCarbon_AddCenturies(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddCenturies(1)
	}
}

func BenchmarkCarbon_AddCenturiesNoOverflow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddCenturiesNoOverflow(1)
	}
}

func BenchmarkCarbon_AddCentury(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddCentury()
	}
}

func BenchmarkCarbon_AddCenturyNoOverflow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddCenturyNoOverflow()
	}
}

func BenchmarkCarbon_SubCenturies(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubCenturies(1)
	}
}

func BenchmarkCarbon_SubCenturiesNoOverflow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubCenturiesNoOverflow(1)
	}
}

func BenchmarkCarbon_SubCentury(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubCentury()
	}
}

func BenchmarkCarbon_SubCenturyNoOverflow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubCenturyNoOverflow()
	}
}

func BenchmarkCarbon_AddDecades(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddDecades(1)
	}
}

func BenchmarkCarbon_AddDecadesNoOverflow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddDecadesNoOverflow(1)
	}
}

func BenchmarkCarbon_AddDecade(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddDecade()
	}
}

func BenchmarkCarbon_AddDecadeNoOverflow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddDecadeNoOverflow()
	}
}

func BenchmarkCarbon_SubDecades(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubDecades(1)
	}
}

func BenchmarkCarbon_SubDecadesNoOverflow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubDecadesNoOverflow(1)
	}
}

func BenchmarkCarbon_SubDecade(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubDecade()
	}
}

func BenchmarkCarbon_SubDecadeNoOverflow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubDecadeNoOverflow()
	}
}

func BenchmarkCarbon_AddYears(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddYears(1)
	}
}

func BenchmarkCarbon_AddYearsNoOverflow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddYearsNoOverflow(1)
	}
}

func BenchmarkCarbon_AddYear(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddYear()
	}
}

func BenchmarkCarbon_AddYearNoOverflow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddYearNoOverflow()
	}
}

func BenchmarkCarbon_SubYears(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubYears(1)
	}
}

func BenchmarkCarbon_SubYearsNoOverflow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubYearsNoOverflow(1)
	}
}

func BenchmarkCarbon_SubYear(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubYear()
	}
}

func BenchmarkCarbon_SubYearNoOverflow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubYearNoOverflow()
	}
}

func BenchmarkCarbon_AddQuarters(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddQuarters(1)
	}
}

func BenchmarkCarbon_AddQuartersNoOverflow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddQuartersNoOverflow(1)
	}
}

func BenchmarkCarbon_AddQuarter(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddQuarter()
	}
}

func BenchmarkCarbon_AddQuarterNoOverflow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddQuarter()
	}
}

func BenchmarkCarbon_SubQuarters(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubQuarters(1)
	}
}

func BenchmarkCarbon_SubQuartersNoOverflow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubQuartersNoOverflow(1)
	}
}

func BenchmarkCarbon_SubQuarter(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubQuarter()
	}
}

func BenchmarkCarbon_SubQuarterNoOverflow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubQuarterNoOverflow()
	}
}

func BenchmarkCarbon_AddMonths(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddMonths(1)
	}
}

func BenchmarkCarbon_AddMonthsNoOverflow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddMonthsNoOverflow(1)
	}
}

func BenchmarkCarbon_AddMonth(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddMonth()
	}
}

func BenchmarkCarbon_AddMonthNoOverflow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddMonthNoOverflow()
	}
}

func BenchmarkCarbon_SubMonths(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubMonths(1)
	}
}

func BenchmarkCarbon_SubMonthsNoOverflow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubMonthsNoOverflow(1)
	}
}

func BenchmarkCarbon_SubMonth(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubMonth()
	}
}

func BenchmarkCarbon_SubMonthNoOverflow(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubMonthNoOverflow()
	}
}

func BenchmarkCarbon_AddWeeks(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddWeeks(1)
	}
}

func BenchmarkCarbon_AddWeek(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddWeek()
	}
}

func BenchmarkCarbon_SubWeeks(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubWeeks(1)
	}
}

func BenchmarkCarbon_SubWeek(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubWeek()
	}
}

func BenchmarkCarbon_AddDays(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddDays(1)
	}
}

func BenchmarkCarbon_AddDay(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddDay()
	}
}

func BenchmarkCarbon_SubDays(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubDays(1)
	}
}

func BenchmarkCarbon_SubDay(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubDay()
	}
}

func BenchmarkCarbon_AddHours(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddHours(1)
	}
}

func BenchmarkCarbon_AddHour(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddHour()
	}
}

func BenchmarkCarbon_SubHours(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubHours(1)
	}
}

func BenchmarkCarbon_SubHour(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubHour()
	}
}

func BenchmarkCarbon_AddMinutes(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddMinutes(1)
	}
}

func BenchmarkCarbon_AddMinute(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddMinute()
	}
}

func BenchmarkCarbon_SubMinutes(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubMinutes(1)
	}
}

func BenchmarkCarbon_SubMinute(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubMinute()
	}
}

func BenchmarkCarbon_AddSeconds(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddSeconds(1)
	}
}

func BenchmarkCarbon_AddSecond(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddSecond()
	}
}

func BenchmarkCarbon_SubSeconds(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubSeconds(1)
	}
}

func BenchmarkCarbon_SubSecond(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubSecond()
	}
}

func BenchmarkCarbon_AddMilliseconds(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddMilliseconds(1)
	}
}

func BenchmarkCarbon_AddMillisecond(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddMillisecond()
	}
}

func BenchmarkCarbon_SubMilliseconds(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubMilliseconds(1)
	}
}

func BenchmarkCarbon_SubMillisecond(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubMillisecond()
	}
}

func BenchmarkCarbon_AddMicroseconds(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddMicroseconds(1)
	}
}

func BenchmarkCarbon_AddMicrosecond(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddMicrosecond()
	}
}

func BenchmarkCarbon_SubMicroseconds(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubMicroseconds(1)
	}
}

func BenchmarkCarbon_SubMicrosecond(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubMicrosecond()
	}
}

func BenchmarkCarbon_AddNanoseconds(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddNanoseconds(1)
	}
}

func BenchmarkCarbon_AddNanosecond(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.AddNanosecond()
	}
}

func BenchmarkCarbon_SubNanoseconds(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubNanoseconds(1)
	}
}

func BenchmarkCarbon_SubNanosecond(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SubNanosecond()
	}
}
