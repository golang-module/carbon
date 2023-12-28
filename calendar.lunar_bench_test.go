package carbon

import "testing"

func BenchmarkCarbon_Lunar(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Now().Lunar()
	}
}

func BenchmarkLunar_Animal(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.Animal()
	}
}

func BenchmarkLunar_Festival(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.Festival()
	}
}

func BenchmarkLunar_DateTime(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.DateTime()
	}
}

func BenchmarkLunar_Date(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.Date()
	}
}

func BenchmarkLunar_Time(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.Time()
	}
}

func BenchmarkLunar_Year(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.Year()
	}
}

func BenchmarkLunar_Month(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.Month()
	}
}

func BenchmarkLunar_LeapMonth(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.LeapMonth()
	}
}

func BenchmarkLunar_Day(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.Day()
	}
}

func BenchmarkLunar_ToYearString(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.ToYearString()
	}
}

func BenchmarkLunar_ToMonthString(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.ToMonthString()
	}
}

func BenchmarkLunar_ToDayString(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.ToDayString()
	}
}

func BenchmarkLunar_ToDateString(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.ToDateString()
	}
}

func BenchmarkLunar_String(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.String()
	}
}

func BenchmarkLunar_IsLeapYear(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsLeapYear()
	}
}

func BenchmarkLunar_IsLeapMonth(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsLeapMonth()
	}
}

func BenchmarkLunar_IsRatYear(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsRatYear()
	}
}

func BenchmarkLunar_IsOxYear(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsOxYear()
	}
}

func BenchmarkLunar_IsTigerYear(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsTigerYear()
	}
}

func BenchmarkLunar_IsRabbitYear(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsRabbitYear()
	}
}

func BenchmarkLunar_IsDragonYear(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsDragonYear()
	}
}

func BenchmarkLunar_IsSnakeYear(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsSnakeYear()
	}
}

func BenchmarkLunar_IsHorseYear(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsHorseYear()
	}
}

func BenchmarkLunar_IsGoatYear(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsGoatYear()
	}
}

func BenchmarkLunar_IsMonkeyYear(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsMonkeyYear()
	}
}

func BenchmarkLunar_IsRoosterYear(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsRoosterYear()
	}
}

func BenchmarkLunar_IsDogYear(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsDogYear()
	}
}

func BenchmarkLunar_IsPigYear(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsPigYear()
	}
}

func BenchmarkLunar_DoubleHour(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.DoubleHour()
	}
}

func BenchmarkLunar_IsFirstDoubleHour(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsFirstDoubleHour()
	}
}

func BenchmarkLunar_IsSecondDoubleHour(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsSecondDoubleHour()
	}
}

func BenchmarkLunar_IsThirdDoubleHour(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsThirdDoubleHour()
	}
}

func BenchmarkLunar_IsFourthDoubleHour(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsFourthDoubleHour()
	}
}

func BenchmarkLunar_IsFifthDoubleHour(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsFifthDoubleHour()
	}
}

func BenchmarkLunar_IsSixthDoubleHour(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsSixthDoubleHour()
	}
}

func BenchmarkLunar_IsSeventhDoubleHour(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsSeventhDoubleHour()
	}
}

func BenchmarkLunar_IsEighthDoubleHour(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsEighthDoubleHour()
	}
}

func BenchmarkLunar_IsNinthDoubleHour(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsNinthDoubleHour()
	}
}

func BenchmarkLunar_IsTenthDoubleHour(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsTenthDoubleHour()
	}
}

func BenchmarkLunar_IsEleventhDoubleHour(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsEleventhDoubleHour()
	}
}

func BenchmarkLunar_IsTwelfthDoubleHour(b *testing.B) {
	l := Now().Lunar()
	for n := 0; n < b.N; n++ {
		l.IsTwelfthDoubleHour()
	}
}
