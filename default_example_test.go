package carbon_test

import (
	"fmt"

	"github.com/dromara/carbon/v2"
)

func ExampleSetDefault() {
	defer carbon.ReSetDefault()

	carbon.SetDefault(carbon.Default{
		Layout:       carbon.DateTimeLayout,
		Timezone:     carbon.PRC,
		Locale:       "zh-CN",
		WeekStartsAt: carbon.Sunday,
	})

	fmt.Println(carbon.DefaultLayout)
	fmt.Println(carbon.DefaultTimezone)
	fmt.Println(carbon.DefaultWeekStartsAt)
	fmt.Println(carbon.DefaultLocale)

	// Output:
	// 2006-01-02 15:04:05
	// PRC
	// Sunday
	// zh-CN
}
