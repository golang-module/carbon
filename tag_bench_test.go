package carbon

import "testing"

func BenchmarkCarbon_SetTag(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SetTag(&tag{
			carbon: "datetime",
			tz:     Local,
		})
	}
}

func BenchmarkCarbon_LoadTag(b *testing.B) {
	type Student struct {
		Birthday Carbon `json:"birthday" carbon:"type:date"`
	}
	student := Student{
		Birthday: Now(),
	}
	for n := 0; n < b.N; n++ {
		_ = LoadTag(&student)
	}
}
