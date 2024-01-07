package carbon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCarbon_SetDefault(t *testing.T) {
	SetDefault(Default{
		Layout:   DateTimeLayout,
		Timezone: Local,
		Locale:   "en",
	})
	assert.Equal(t, DateTimeLayout, defaultLayout)
	assert.Equal(t, Local, defaultTimezone)
	assert.Equal(t, "en", defaultLocale)
}
