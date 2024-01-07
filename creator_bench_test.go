package carbon

import (
	"testing"
	"time"
)

func BenchmarkCarbon_CreateFromStdTime(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateFromStdTime(time.Now())
	}
	for n := 0; n < b.N; n++ {
		CreateFromStdTime(time.Now(), PRC)
	}
}

func BenchmarkCarbon_CreateFromTimestamp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateFromTimestamp(1649735755)
	}
	for n := 0; n < b.N; n++ {
		CreateFromTimestamp(1649735755, PRC)
	}
}

func BenchmarkCarbon_CreateFromTimestampMilli(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateFromTimestampMilli(1649735755981)
	}
	for n := 0; n < b.N; n++ {
		CreateFromTimestampMilli(1649735755981, PRC)
	}
}

func BenchmarkCarbon_CreateFromTimestampMicro(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateFromTimestampMicro(1649735755981566)
	}
	for n := 0; n < b.N; n++ {
		CreateFromTimestampMicro(1649735755981566, PRC)
	}
}

func BenchmarkCarbon_CreateFromTimestampNano(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateFromTimestampNano(1649735755981566000)
	}
	for n := 0; n < b.N; n++ {
		CreateFromTimestampNano(1649735755981566000, PRC)
	}
}

func BenchmarkCarbon_CreateFromDateTime(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateFromDateTime(2020, 8, 5, 13, 14, 15)
	}
	for n := 0; n < b.N; n++ {
		CreateFromDateTime(2020, 8, 5, 13, 14, 15, PRC)
	}
}

func BenchmarkCarbon_CreateFromDateTimeMilli(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateFromDateTimeMilli(2020, 8, 5, 13, 14, 15, 0)
	}
	for n := 0; n < b.N; n++ {
		CreateFromDateTimeMilli(2020, 8, 5, 13, 14, 15, 0, PRC)
	}
}

func BenchmarkCarbon_CreateFromDateTimeMicro(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateFromDateTimeMicro(2020, 8, 5, 13, 14, 15, 0)
	}
	for n := 0; n < b.N; n++ {
		CreateFromDateTimeMicro(2020, 8, 5, 13, 14, 15, 0, PRC)
	}
}

func BenchmarkCarbon_CreateFromDateTimeNano(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateFromDateTimeNano(2020, 8, 5, 13, 14, 15, 0)
	}
	for n := 0; n < b.N; n++ {
		CreateFromDateTimeNano(2020, 8, 5, 13, 14, 15, 0, PRC)
	}
}

func BenchmarkCarbon_CreateFromDate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateFromDate(2020, 8, 5)
	}
	for n := 0; n < b.N; n++ {
		CreateFromDate(2020, 8, 5, PRC)
	}
}

func BenchmarkCarbon_CreateFromDateMilli(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateFromDateMilli(2020, 8, 5, 0)
	}
	for n := 0; n < b.N; n++ {
		CreateFromDateMilli(2020, 8, 5, 0, PRC)
	}
}

func BenchmarkCarbon_CreateFromDateMicro(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateFromDateMicro(2020, 8, 5, 0)
	}
	for n := 0; n < b.N; n++ {
		CreateFromDateMicro(2020, 8, 5, 0, PRC)
	}
}

func BenchmarkCarbon_CreateFromDateNano(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateFromDateNano(2020, 8, 5, 0)
	}
	for n := 0; n < b.N; n++ {
		CreateFromDateNano(2020, 8, 5, 0, PRC)
	}
}

func BenchmarkCarbon_CreateFromTime(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateFromTime(13, 14, 15)
	}
	for n := 0; n < b.N; n++ {
		CreateFromTime(13, 14, 15, PRC)
	}
}

func BenchmarkCarbon_CreateFromTimeMilli(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateFromTimeMilli(13, 14, 15, 0)
	}
	for n := 0; n < b.N; n++ {
		CreateFromTimeMilli(13, 14, 15, 0, PRC)
	}
}

func BenchmarkCarbon_CreateFromTimeMicro(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateFromTimeMicro(13, 14, 15, 0)
	}
	for n := 0; n < b.N; n++ {
		CreateFromTimeMicro(13, 14, 15, 0, PRC)
	}
}

func BenchmarkCarbon_CreateFromTimeNano(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateFromTimeNano(13, 14, 15, 0)
	}
	for n := 0; n < b.N; n++ {
		CreateFromTimeNano(13, 14, 15, 0, PRC)
	}
}
