package carbon_test

import (
	"fmt"

	"github.com/dromara/carbon/v2"
)

func ExampleMaxValue() {
	fmt.Println(carbon.MaxValue().ToString())
	// Output:
	// 9999-12-31 23:59:59.999999999 +0000 UTC
}

func ExampleMinValue() {
	fmt.Println(carbon.MinValue().ToString())
	// Output:
	// -9998-01-01 00:00:00 +0000 UTC
}

func ExampleMax() {
	fmt.Println(carbon.Max(carbon.Parse("2020-08-06"), carbon.Parse("2021-08-05")).ToString())
	// Output:
	// 2021-08-05 00:00:00 +0000 UTC
}

func ExampleMin() {
	fmt.Println(carbon.Min(carbon.Parse("2020-08-06"), carbon.Parse("2021-08-05")).ToString())
	// Output:
	// 2020-08-06 00:00:00 +0000 UTC
}

func ExampleCarbon_Closest() {
	fmt.Println(carbon.Parse("2020-08-05").Closest(carbon.Parse("2020-08-06"), carbon.Parse("2021-08-05")).ToString())
	// Output:
	// 2020-08-06 00:00:00 +0000 UTC
}

func ExampleCarbon_Farthest() {
	fmt.Println(carbon.Parse("2020-08-05").Farthest(carbon.Parse("2020-08-06"), carbon.Parse("2021-08-05")).ToString())
	// Output:
	// 2021-08-05 00:00:00 +0000 UTC
}
