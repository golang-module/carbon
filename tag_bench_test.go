package carbon

import "testing"

func BenchmarkCarbon_Tag(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Tag()
	}
}

func BenchmarkCarbon_SetTag(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SetTag("layout:2006-01-02")
	}
}

func BenchmarkCarbon_LoadTag(b *testing.B) {
	type Student struct {
		Birthday Carbon `json:"birthday" carbon:"date"`
	}
	student := Student{
		Birthday: Now(),
	}
	for n := 0; n < b.N; n++ {
		_ = LoadTag(&student)
	}
}
