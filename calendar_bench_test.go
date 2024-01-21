package carbon

import "testing"

func BenchmarkCarbon_Lunar(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Lunar()
	}
}

func BenchmarkCarbon_CreateFromLunar(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateFromLunar(2023, 12, 8, 0, 0, 0, false)
	}
}
