package carbon

import "testing"

func BenchmarkCarbon_Scan(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		_ = now.Scan(nil)
	}
}

func BenchmarkCarbon_Value(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Value()
	}
}

func BenchmarkDateTime_Scan(b *testing.B) {
	now := NewDateTime(Now())
	for n := 0; n < b.N; n++ {
		_ = now.Scan(nil)
	}
}

func BenchmarkDateTime_Value(b *testing.B) {
	now := NewDateTime(Now())
	for n := 0; n < b.N; n++ {
		now.Value()
	}
}

func BenchmarkDateTimeMilli_Scan(b *testing.B) {
	now := NewDateTimeMilli(Now())
	for n := 0; n < b.N; n++ {
		_ = now.Scan(nil)
	}
}

func BenchmarkDateTimeMilli_Value(b *testing.B) {
	now := NewDateTimeMilli(Now())
	for n := 0; n < b.N; n++ {
		now.Value()
	}
}

func BenchmarkNewDateTimeMicro_Scan(b *testing.B) {
	now := NewDateTimeMicro(Now())
	for n := 0; n < b.N; n++ {
		_ = now.Scan(nil)
	}
}

func BenchmarkDateTimeMicro_Value(b *testing.B) {
	now := NewDateTimeMicro(Now())
	for n := 0; n < b.N; n++ {
		now.Value()
	}
}

func BenchmarkNewDateTimeNano_Scan(b *testing.B) {
	now := NewDateTimeNano(Now())
	for n := 0; n < b.N; n++ {
		_ = now.Scan(nil)
	}
}

func BenchmarkDateTimeNano_Value(b *testing.B) {
	now := NewDateTimeNano(Now())
	for n := 0; n < b.N; n++ {
		now.Value()
	}
}

func BenchmarkNewDate_Scan(b *testing.B) {
	now := NewDate(Now())
	for n := 0; n < b.N; n++ {
		_ = now.Scan(nil)
	}
}

func BenchmarkDate_Value(b *testing.B) {
	now := NewDate(Now())
	for n := 0; n < b.N; n++ {
		now.Value()
	}
}

func BenchmarkNewDateMilli_Scan(b *testing.B) {
	now := NewDateMilli(Now())
	for n := 0; n < b.N; n++ {
		_ = now.Scan(nil)
	}
}

func BenchmarkDateMilli_Value(b *testing.B) {
	now := NewDateMilli(Now())
	for n := 0; n < b.N; n++ {
		now.Value()
	}
}

func BenchmarkNewDateMicro_Scan(b *testing.B) {
	now := NewDateMicro(Now())
	for n := 0; n < b.N; n++ {
		_ = now.Scan(nil)
	}
}

func BenchmarkNewDateMicro_Value(b *testing.B) {
	now := NewDateMicro(Now())
	for n := 0; n < b.N; n++ {
		now.Value()
	}
}

func BenchmarkNewDateNano_Scan(b *testing.B) {
	now := NewDateNano(Now())
	for n := 0; n < b.N; n++ {
		_ = now.Scan(nil)
	}
}

func BenchmarkNewDateNano_Value(b *testing.B) {
	now := NewDateNano(Now())
	for n := 0; n < b.N; n++ {
		now.Value()
	}
}

func BenchmarkNewTime_Scan(b *testing.B) {
	now := NewTime(Now())
	for n := 0; n < b.N; n++ {
		_ = now.Scan(nil)
	}
}

func BenchmarkNewTime_Value(b *testing.B) {
	now := NewTime(Now())
	for n := 0; n < b.N; n++ {
		now.Value()
	}
}

func BenchmarkNewTimeMilli_Scan(b *testing.B) {
	now := NewTimeMilli(Now())
	for n := 0; n < b.N; n++ {
		_ = now.Scan(nil)
	}
}

func BenchmarkNewTimeMilli_Value(b *testing.B) {
	now := NewTimeMilli(Now())
	for n := 0; n < b.N; n++ {
		now.Value()
	}
}

func BenchmarkNewTimeMicro_Scan(b *testing.B) {
	now := NewTimeMicro(Now())
	for n := 0; n < b.N; n++ {
		_ = now.Scan(nil)
	}
}

func BenchmarkNewTimeMicro_Value(b *testing.B) {
	now := NewTimeMicro(Now())
	for n := 0; n < b.N; n++ {
		now.Value()
	}
}

func BenchmarkNewTimeNano_Scan(b *testing.B) {
	now := NewTimeNano(Now())
	for n := 0; n < b.N; n++ {
		_ = now.Scan(nil)
	}
}

func BenchmarkNewTimeNano_Value(b *testing.B) {
	now := NewTimeNano(Now())
	for n := 0; n < b.N; n++ {
		now.Value()
	}
}

func BenchmarkNewTimestamp_Scan(b *testing.B) {
	now := NewTimestamp(Now())
	for n := 0; n < b.N; n++ {
		_ = now.Scan(nil)
	}
}

func BenchmarkNewTimestamp_Value(b *testing.B) {
	now := NewTimestamp(Now())
	for n := 0; n < b.N; n++ {
		now.Value()
	}
}

func BenchmarkNewTimestampMilli_Scan(b *testing.B) {
	now := NewTimestampMilli(Now())
	for n := 0; n < b.N; n++ {
		_ = now.Scan(nil)
	}
}

func BenchmarkNewTimestampMilli_Value(b *testing.B) {
	now := NewTimestampMilli(Now())
	for n := 0; n < b.N; n++ {
		now.Value()
	}
}

func BenchmarkNewTimestampNano_Scan(b *testing.B) {
	now := NewTimestampNano(Now())
	for n := 0; n < b.N; n++ {
		_ = now.Scan(nil)
	}
}

func BenchmarkNewTimestampNano_Value(b *testing.B) {
	now := NewTimestampNano(Now())
	for n := 0; n < b.N; n++ {
		now.Value()
	}
}

func BenchmarkCarbon_MarshalJSON(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkCarbon_UnmarshalJSON(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkDateTime_MarshalJSON(b *testing.B) {
	now := NewDateTime(Now())
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkDateTime_UnmarshalJSON(b *testing.B) {
	now := NewDateTime(Now())
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkDateTimeMilli_MarshalJSON(b *testing.B) {
	now := NewDateTimeMilli(Now())
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkDateTimeMilli_UnmarshalJSON(b *testing.B) {
	now := NewDateTimeMilli(Now())
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkDateTimeMicro_MarshalJSON(b *testing.B) {
	now := NewDateTimeMicro(Now())
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkDateTimeMicro_UnmarshalJSON(b *testing.B) {
	now := NewDateTimeMicro(Now())
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkDateTimeNano_MarshalJSON(b *testing.B) {
	now := NewDateTimeNano(Now())
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkDateTimeNano_UnmarshalJSON(b *testing.B) {
	now := NewDateTimeNano(Now())
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkDate_MarshalJSON(b *testing.B) {
	now := NewDate(Now())
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkDate_UnmarshalJSON(b *testing.B) {
	now := NewDate(Now())
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkDateMicro_MarshalJSON(b *testing.B) {
	now := NewDateMicro(Now())
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkDateMicro_UnmarshalJSON(b *testing.B) {
	now := NewDateMicro(Now())
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkDateNano_MarshalJSON(b *testing.B) {
	now := NewDateNano(Now())
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkDateNano_UnmarshalJSON(b *testing.B) {
	now := NewDateNano(Now())
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkTime_MarshalJSON(b *testing.B) {
	now := NewTime(Now())
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkTime_UnmarshalJSON(b *testing.B) {
	now := NewTime(Now())
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkTimeMicro_MarshalJSON(b *testing.B) {
	now := NewTimeMicro(Now())
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkTimeMicro_UnmarshalJSON(b *testing.B) {
	now := NewTimeMicro(Now())
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkTimeNano_MarshalJSON(b *testing.B) {
	now := NewTimeNano(Now())
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkTimeNano_UnmarshalJSON(b *testing.B) {
	now := NewTimeNano(Now())
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkTimestamp_MarshalJSON(b *testing.B) {
	now := NewTimestamp(Now())
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkTimestamp_UnmarshalJSON(b *testing.B) {
	now := NewTimestamp(Now())
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkTimestampMicro_MarshalJSON(b *testing.B) {
	now := NewTimestampMicro(Now())
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkTimestampMicro_UnmarshalJSON(b *testing.B) {
	now := NewTimestampMicro(Now())
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkTimestampNano_MarshalJSON(b *testing.B) {
	now := NewTimestampNano(Now())
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkTimestampNano_UnmarshalJSON(b *testing.B) {
	now := NewTimestampNano(Now())
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}
