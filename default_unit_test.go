package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_SetDefault(t *testing.T) {
	SetDefault(Default{
		Layout:       DateTimeLayout,
		Timezone:     Local,
		Locale:       "en",
		WeekStartsAt: Sunday,
	})
	assert.Equal(t, DateTimeLayout, defaultLayout)
	assert.Equal(t, Local, defaultTimezone)
	assert.Equal(t, "en", defaultLocale)
	assert.Equal(t, "Sunday", defaultWeekStartsAt)
}
