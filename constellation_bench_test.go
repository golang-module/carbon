package carbon

import "testing"

func BenchmarkCarbon_Constellation(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Constellation()
	}
}

func BenchmarkCarbon_IsAries(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsAries()
	}
}

func BenchmarkCarbon_IsTaurus(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsTaurus()
	}
}

func BenchmarkCarbon_IsGemini(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsGemini()
	}
}

func BenchmarkCarbon_IsCancer(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsCancer()
	}
}

func BenchmarkCarbon_IsLeo(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsLeo()
	}
}

func BenchmarkCarbon_IsVirgo(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsVirgo()
	}
}

func BenchmarkCarbon_IsLibra(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsLibra()
	}
}

func BenchmarkCarbon_IsScorpio(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsScorpio()
	}
}

func BenchmarkCarbon_IsSagittarius(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSagittarius()
	}
}

func BenchmarkCarbon_IsCapricorn(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsCapricorn()
	}
}

func BenchmarkCarbon_IsAquarius(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsAquarius()
	}
}

func BenchmarkCarbon_IsPisces(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsPisces()
	}
}
