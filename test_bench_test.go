package carbon

import "testing"

func BenchmarkCarbon_SetTestNow(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.SetTestNow(Yesterday())
	}
}

func BenchmarkCarbon_UnSetTestNow(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.UnSetTestNow()
	}
}

func BenchmarkCarbon_IsSetTestNow(b *testing.B) {
	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.IsSetTestNow()
	}
}
