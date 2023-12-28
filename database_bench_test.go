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

func BenchmarkCarbon_GormDataType(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.GormDataType()
	}
}
