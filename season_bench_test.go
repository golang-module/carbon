package carbon

import "testing"

func BenchmarkCarbon_Season(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Season()
	}
}

func BenchmarkCarbon_StartOfSeason(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.StartOfSeason()
	}
}

func BenchmarkCarbon_EndOfSeason(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.EndOfSeason()
	}
}

func BenchmarkCarbon_IsSpring(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSpring()
	}
}

func BenchmarkCarbon_IsSummer(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsSummer()
	}
}

func BenchmarkCarbon_IsAutumn(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsAutumn()
	}
}

func BenchmarkCarbon_IsWinter(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.IsWinter()
	}
}
