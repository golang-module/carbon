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
