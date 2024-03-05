package carbon

import "testing"

func BenchmarkCarbon_String(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.String()
	}
}

func BenchmarkCarbon_ToString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToString()
	}
}

func BenchmarkCarbon_ToMonthString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToMonthString()
	}
}

func BenchmarkCarbon_ToShortMonthString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToShortMonthString()
	}
}

func BenchmarkCarbon_ToWeekString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToWeekString()
	}
}

func BenchmarkCarbon_ToShortWeekString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToWeekString()
	}
}

func BenchmarkCarbon_ToDayDateTimeString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToDayDateTimeString()
	}
}

func BenchmarkCarbon_ToDateTimeString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToDateTimeString()
	}
}

func BenchmarkCarbon_ToDateTimeMilliString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToDateTimeMilliString()
	}
}

func BenchmarkCarbon_ToDateTimeMicroString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToDateTimeMicroString()
	}
}

func BenchmarkCarbon_ToDateTimeNanoString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToDateTimeNanoString()
	}
}

func BenchmarkCarbon_ToShortDateTimeString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToShortDateTimeString()
	}
}

func BenchmarkCarbon_ToShortDateTimeMilliString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToShortDateTimeString()
	}
}

func BenchmarkCarbon_ToShortDateTimeMicroString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToShortDateTimeMicroString()
	}
}

func BenchmarkCarbon_ToShortDateTimeNanoString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToShortDateTimeNanoString()
	}
}

func BenchmarkCarbon_ToDateString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToDateString()
	}
}

func BenchmarkCarbon_ToDateMilliString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToDateMilliString()
	}
}

func BenchmarkCarbon_ToDateMicroString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToDateMicroString()
	}
}

func BenchmarkCarbon_ToDateNanoString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToDateNanoString()
	}
}

func BenchmarkCarbon_ToShortDateString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToShortDateString()
	}
}

func BenchmarkCarbon_ToShortDateMilliString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToShortDateMilliString()
	}
}

func BenchmarkCarbon_ToToShortDateMicroString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToShortDateMicroString()
	}
}

func BenchmarkCarbon_ToShortDateNanoString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToShortDateNanoString()
	}
}

func BenchmarkCarbon_ToTimeString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToTimeString()
	}
}

func BenchmarkCarbon_ToTimeMilliString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToTimeMilliString()
	}
}

func BenchmarkCarbon_ToTimeMicroString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToTimeMicroString()
	}
}

func BenchmarkCarbon_ToTimeNanoString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToTimeNanoString()
	}
}

func BenchmarkCarbon_ToShortTimeString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToShortTimeString()
	}
}

func BenchmarkCarbon_ToShortTimeMilliString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToShortTimeMilliString()
	}
}

func BenchmarkCarbon_ToShortTimeMicroString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToShortTimeMilliString()
	}
}

func BenchmarkCarbon_ToShortTimeNanoString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToShortTimeNanoString()
	}
}

func BenchmarkCarbon_ToAtomString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToAtomString()
	}
}

func BenchmarkCarbon_ToAnsicString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToAnsicString()
	}
}

func BenchmarkCarbon_ToCookieString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToCookieString()
	}
}

func BenchmarkCarbon_ToRssString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToRssString()
	}
}

func BenchmarkCarbon_ToW3cString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToW3cString()
	}
}

func BenchmarkCarbon_ToUnixDateString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToUnixDateString()
	}
}

func BenchmarkCarbon_ToRubyDateString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToRubyDateString()
	}
}

func BenchmarkCarbon_ToKitchenString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToKitchenString()
	}
}

func BenchmarkCarbon_ToIso8601String(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToIso8601String()
	}
}

func BenchmarkCarbon_ToIso8601MilliString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToIso8601MilliString()
	}
}

func BenchmarkCarbon_ToIso8601MicroString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToIso8601MicroString()
	}
}

func BenchmarkCarbon_ToIso8601NanoString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToIso8601NanoString()
	}
}

func BenchmarkCarbon_ToIso8601ZuluString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToIso8601ZuluString()
	}
}

func BenchmarkCarbon_ToIso8601ZuluMilliString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToIso8601ZuluMilliString()
	}
}

func BenchmarkCarbon_ToIso8601ZuluMicroString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToIso8601ZuluMicroString()
	}
}

func BenchmarkCarbon_ToIso8601ZuluNanoString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToIso8601ZuluNanoString()
	}
}

func BenchmarkCarbon_ToRfc822String(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToRfc822String()
	}
}

func BenchmarkCarbon_ToRfc822zString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToRfc822zString()
	}
}

func BenchmarkCarbon_ToRfc850String(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToRfc850String()
	}
}

func BenchmarkCarbon_ToRfc1036String(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToRfc1036String()
	}
}

func BenchmarkCarbon_ToRfc1123String(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToRfc1123String()
	}
}

func BenchmarkCarbon_ToRfc1123zString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToRfc1123zString()
	}
}

func BenchmarkCarbon_ToRfc2822String(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToRfc2822String()
	}
}

func BenchmarkCarbon_ToRfc3339String(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToRfc2822String()
	}
}

func BenchmarkCarbon_ToRfc3339MilliString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToRfc3339MilliString()
	}
}

func BenchmarkCarbon_ToRfc3339MicroString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToRfc3339MicroString()
	}
}

func BenchmarkCarbon_ToRfc3339NanoString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToRfc3339NanoString()
	}
}

func BenchmarkCarbon_ToRfc7231String(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToRfc7231String()
	}
}

func BenchmarkCarbon_ToFormattedDateString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToFormattedDateString()
	}
}

func BenchmarkCarbon_ToFormattedDayDateString(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.ToFormattedDayDateString()
	}
}

func BenchmarkCarbon_Layout(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Layout("2006-01-02", "2020-08-05")
	}
}

func BenchmarkCarbon_Format(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Format("2006-01-02", "Y-m-d")
	}
}

func BenchmarkCarbon_GoTime(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.GoString()
	}
}
