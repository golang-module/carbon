package carbon

import "testing"

func BenchmarkCarbon_Parse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parse("2020-08-05")
	}
}

func BenchmarkCarbon_ParseByFormat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ParseByFormat("2020-08-05", "Y-m-d")
	}
}

func BenchmarkCarbon_ParseByLayout(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ParseByLayout("2020-08-05", "2006-01-02")
	}
}
