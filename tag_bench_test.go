package carbon

import "testing"

func BenchmarkTag_SetTag(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.SetTag(&Tag{
			carbon: "datetime",
			tz:     Local,
		})
	}
}

func BenchmarkTag_SetLayout(b *testing.B) {
	tag := NewTag()
	for n := 0; n < b.N; n++ {
		tag.SetLayout(DateTimeLayout)
	}
}

func BenchmarkTag_SetFormat(b *testing.B) {
	tag := NewTag()
	for n := 0; n < b.N; n++ {
		tag.SetFormat(DateTimeFormat)
	}
}

func BenchmarkTag_SetType(b *testing.B) {
	tag := NewTag()
	for n := 0; n < b.N; n++ {
		tag.SetType("dateTime")
	}
}

func BenchmarkTag_SetTimezone(b *testing.B) {
	tag := NewTag()
	for n := 0; n < b.N; n++ {
		tag.SetTimezone(PRC)
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
