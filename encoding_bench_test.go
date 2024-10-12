package carbon

import "testing"

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
	now := Now().ToDateTimeStruct()
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkDateTime_UnmarshalJSON(b *testing.B) {
	now := Now().ToDateTimeStruct()
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkDateTimeMilli_MarshalJSON(b *testing.B) {
	now := Now().ToDateTimeMilliStruct()
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkDateTimeMilli_UnmarshalJSON(b *testing.B) {
	now := Now().ToDateTimeMilliStruct()
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkDateTimeMicro_MarshalJSON(b *testing.B) {
	now := Now().ToDateTimeMicroStruct()
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkDateTimeMicro_UnmarshalJSON(b *testing.B) {
	now := Now().ToDateTimeMicroStruct()
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkDateTimeNano_MarshalJSON(b *testing.B) {
	now := Now().ToDateTimeNanoStruct()
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkDateTimeNano_UnmarshalJSON(b *testing.B) {
	now := Now().ToDateTimeNanoStruct()
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkDate_MarshalJSON(b *testing.B) {
	now := Now().ToDateStruct()
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkDate_UnmarshalJSON(b *testing.B) {
	now := Now().ToDateStruct()
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkDateMicro_MarshalJSON(b *testing.B) {
	now := Now().ToDateMicroStruct()
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkDateMicro_UnmarshalJSON(b *testing.B) {
	now := Now().ToDateMicroStruct()
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkDateNano_MarshalJSON(b *testing.B) {
	now := Now().ToDateNanoStruct()
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkDateNano_UnmarshalJSON(b *testing.B) {
	now := Now().ToDateNanoStruct()
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkTime_MarshalJSON(b *testing.B) {
	now := Now().ToTimeStruct()
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkTime_UnmarshalJSON(b *testing.B) {
	now := Now().ToTimeStruct()
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkTimeMicro_MarshalJSON(b *testing.B) {
	now := Now().ToTimeMicroStruct()
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkTimeMicro_UnmarshalJSON(b *testing.B) {
	now := Now().ToTimeMicroStruct()
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkTimeNano_MarshalJSON(b *testing.B) {
	now := Now().ToTimeNanoStruct()
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkTimeNano_UnmarshalJSON(b *testing.B) {
	now := Now().ToTimeNanoStruct()
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkTimestamp_MarshalJSON(b *testing.B) {
	now := Now().ToTimestampStruct()
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkTimestamp_UnmarshalJSON(b *testing.B) {
	now := Now().ToTimestampStruct()
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkTimestampMicro_MarshalJSON(b *testing.B) {
	now := Now().ToTimestampMicroStruct()
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkTimestampMicro_UnmarshalJSON(b *testing.B) {
	now := Now().ToTimestampMicroStruct()
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}

func BenchmarkTimestampNano_MarshalJSON(b *testing.B) {
	now := Now().ToTimestampNanoStruct()
	for n := 0; n < b.N; n++ {
		now.MarshalJSON()
	}
}

func BenchmarkTimestampNano_UnmarshalJSON(b *testing.B) {
	now := Now().ToTimestampNanoStruct()
	for n := 0; n < b.N; n++ {
		now.UnmarshalJSON(nil)
	}
}
