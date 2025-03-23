package carbon_test

import (
	"fmt"

	"github.com/dromara/carbon/v2"
)

func ExampleSetTestNow() {
	defer carbon.CleanTestNow()

	now := carbon.Parse("2020-08-05")
	carbon.SetTestNow(now)

	fmt.Println(carbon.Now().ToDateString())
	fmt.Println(carbon.Yesterday().ToDateString())
	fmt.Println(carbon.Tomorrow().ToDateString())
	fmt.Println(carbon.Now().DiffForHumans())
	fmt.Println(carbon.Yesterday().DiffForHumans())
	fmt.Println(carbon.Tomorrow().DiffForHumans())
	fmt.Println(carbon.Parse("2020-10-05").DiffForHumans())
	fmt.Println(now.DiffForHumans(carbon.Parse("2020-10-05")))

	// Output:
	// 2020-08-05
	// 2020-08-04
	// 2020-08-06
	// just now
	// 1 day ago
	// 1 day from now
	// 2 months from now
	// 2 months before
}

func ExampleCleanTestNow() {
	now := carbon.Parse("2020-08-05")
	carbon.SetTestNow(now)

	fmt.Println(carbon.IsTestNow())

	carbon.CleanTestNow()

	fmt.Println(carbon.IsTestNow())

	// Output:
	// true
	// false
}
