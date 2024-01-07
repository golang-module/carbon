package carbon

import "testing"

func BenchmarkCarbon_SetDefault(b *testing.B) {
	d := Default{
		Layout:   DateTimeLayout,
		Timezone: Local,
		Locale:   "en",
	}
	for n := 0; n < b.N; n++ {
		SetDefault(d)
	}
}
