package carbon

import (
	"github.com/dromara/carbon/v2"
)

var (
	// DefaultLayout default layout
	// 默认布局模板
	DefaultLayout = carbon.DateTimeLayout

	// DefaultTimezone default timezone
	// 默认时区
	DefaultTimezone = carbon.UTC

	// DefaultLocale default language locale
	// 默认语言区域
	DefaultLocale = carbon.DefaultLocale

	// DefaultWeekStartsAt default start date of the week
	// 默认一周开始日期
	DefaultWeekStartsAt = carbon.DefaultWeekStartsAt

	// DefaultWeekendDays default weekend days of the week
	// 默认周末日期
	DefaultWeekendDays = carbon.DefaultWeekendDays
)
