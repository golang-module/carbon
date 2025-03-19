package carbon_test

import (
	"fmt"
	"time"

	"github.com/dromara/carbon/v2"
)

func ExampleNewCarbon() {
	fmt.Println(carbon.NewCarbon().ToString())

	t1, _ := time.Parse(carbon.DateTimeLayout, "2020-08-05 13:14:15")
	fmt.Println(carbon.NewCarbon(t1).ToString())

	loc, _ := time.LoadLocation(carbon.PRC)
	t2, _ := time.ParseInLocation(carbon.DateTimeLayout, "2020-08-05 13:14:15", loc)
	fmt.Println(carbon.NewCarbon(t2).ToString())

	// Output:
	// 0001-01-01 00:00:00 +0000 UTC
	// 2020-08-05 13:14:15 +0000 UTC
	// 2020-08-05 13:14:15 +0800 CST
}

func ExampleCarbon_Copy() {
	oldCarbon := carbon.Parse("2020-08-05")
	newCarbon := oldCarbon.Copy()

	oldCarbon = oldCarbon.AddDay().SetLayout(carbon.DateTimeLayout).SetLocale("zh-CN").SetWeekStartsAt(carbon.Monday)

	fmt.Printf("old time: %s\n", oldCarbon.ToString())
	fmt.Printf("new time: %s\n", newCarbon.ToString())

	fmt.Printf("old layout: %s\n", oldCarbon.CurrentLayout())
	fmt.Printf("new layout: %s\n", newCarbon.CurrentLayout())

	fmt.Printf("old locale: %s\n", oldCarbon.Locale())
	fmt.Printf("new locale: %s\n", newCarbon.Locale())

	fmt.Printf("old week starts at: %s\n", oldCarbon.WeekStartsAt())
	fmt.Printf("new week starts at: %s\n", newCarbon.WeekStartsAt())

	// Output:
	// old time: 2020-08-06 00:00:00 +0000 UTC
	// new time: 2020-08-05 00:00:00 +0000 UTC
	// old layout: 2006-01-02 15:04:05
	// new layout: 2006-01-02
	// old locale: zh-CN
	// new locale: en
	// old week starts at: Monday
	// new week starts at: Sunday
}
