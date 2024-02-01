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

func BenchmarkCarbon_Julian(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Julian()
	}
}

func BenchmarkCarbon_CreateFromJulian(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateFromJulian(60332)
	}
}

func BenchmarkCarbon_Persian(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Persian()
	}
}

func BenchmarkCarbon_CreateFromPersian(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateFromPersian(1455, 1, 1, 0, 0, 0)
	}
}
