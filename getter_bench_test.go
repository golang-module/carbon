package carbon

import "testing"

func BenchmarkCarbon_DaysInYear(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DaysInYear()
	}
}

func BenchmarkCarbon_DaysInMonth(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DaysInMonth()
	}
}

func BenchmarkCarbon_MonthOfYear(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.MonthOfYear()
	}
}

func BenchmarkCarbon_DayOfYear(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DayOfYear()
	}
}

func BenchmarkCarbon_DayOfMonth(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DayOfMonth()
	}
}

func BenchmarkCarbon_DayOfWeek(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DayOfWeek()
	}
}

func BenchmarkCarbon_WeekOfYear(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.WeekOfYear()
	}
}

func BenchmarkCarbon_WeekOfMonth(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.WeekOfMonth()
	}
}

func BenchmarkCarbon_DateTime(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DateTime()
	}
}

func BenchmarkCarbon_DateTimeMilli(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DateTimeMilli()
	}
}

func BenchmarkCarbon_DateTimeMicro(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DateTimeMilli()
	}
}

func BenchmarkCarbon_DateTimeNano(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DateTimeNano()
	}
}

func BenchmarkCarbon_Date(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Date()
	}
}

func BenchmarkCarbon_DateMilli(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DateMilli()
	}
}

func BenchmarkCarbon_DateMicro(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DateMicro()
	}
}

func BenchmarkCarbon_DateNano(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DateNano()
	}
}

func BenchmarkCarbon_Time(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Time()
	}
}

func BenchmarkCarbon_TimeMilli(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.TimeMilli()
	}
}

func BenchmarkCarbon_TimeMicro(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.TimeMicro()
	}
}

func BenchmarkCarbon_TimeNano(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.TimeNano()
	}
}

func BenchmarkCarbon_Century(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Century()
	}
}

func BenchmarkCarbon_Decade(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Decade()
	}
}

func BenchmarkCarbon_Year(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Year()
	}
}

func BenchmarkCarbon_Quarter(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Quarter()
	}
}

func BenchmarkCarbon_Month(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Month()
	}
}

func BenchmarkCarbon_Week(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Week()
	}
}

func BenchmarkCarbon_Day(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Day()
	}
}

func BenchmarkCarbon_Hour(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Hour()
	}
}

func BenchmarkCarbon_Minute(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Minute()
	}
}

func BenchmarkCarbon_Second(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Second()
	}
}

func BenchmarkCarbon_Millisecond(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Millisecond()
	}
}

func BenchmarkCarbon_Microsecond(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Microsecond()
	}
}

func BenchmarkCarbon_Nanosecond(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Nanosecond()
	}
}

func BenchmarkCarbon_Timestamp(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Timestamp()
	}
}

func BenchmarkCarbon_TimestampMilli(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.TimestampMilli()
	}
}

func BenchmarkCarbon_TimestampMicro(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.TimestampMicro()
	}
}

func BenchmarkCarbon_TimestampNano(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.TimestampNano()
	}
}

func BenchmarkCarbon_Location(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Location()
	}
}

func BenchmarkCarbon_Timezone(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Timezone()
	}
}

func BenchmarkCarbon_Offset(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Offset()
	}
}

func BenchmarkCarbon_Locale(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Locale()
	}
}

func BenchmarkCarbon_Age(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Age()
	}
}
