package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_SetDefault(t *testing.T) {
	defer ReSetDefault()

	SetDefault(Default{
		Layout:       DateTimeLayout,
		Timezone:     PRC,
		Locale:       "zh-CN",
		WeekStartsAt: Sunday,
	})

	assert.Equal(t, DateTimeLayout, DefaultLayout)
	assert.Equal(t, PRC, DefaultTimezone)
	assert.Equal(t, "zh-CN", DefaultLocale)
	assert.Equal(t, "Sunday", DefaultWeekStartsAt)
}
