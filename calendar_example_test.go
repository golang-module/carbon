package carbon_test

import (
	"fmt"

	"github.com/dromara/carbon/v2"
)

func ExampleCarbon_Lunar() {
	fmt.Println(carbon.Parse("2024-01-18", carbon.PRC).Lunar().String())
	fmt.Println(carbon.Parse("2024-01-21", carbon.PRC).Lunar().String())
	fmt.Println(carbon.Parse("2024-01-24", carbon.PRC).Lunar().String())

	// Output:
	// 2023-12-08 00:00:00
	// 2023-12-11 00:00:00
	// 2023-12-14 00:00:00
}

func ExampleCreateFromLunar() {
	fmt.Println(carbon.CreateFromLunar(2023, 12, 11, 0, 0, 0, false).ToDateTimeString(carbon.PRC))
	fmt.Println(carbon.CreateFromLunar(2023, 12, 8, 0, 0, 0, false).ToDateTimeString(carbon.PRC))
	fmt.Println(carbon.CreateFromLunar(2023, 12, 14, 12, 0, 0, false).ToDateTimeString(carbon.PRC))

	// Output:
	// 2024-01-21 00:00:00
	// 2024-01-18 00:00:00
	// 2024-01-24 12:00:00
}

func ExampleCarbon_Julian() {
	fmt.Println(carbon.Parse("2024-01-24 12:00:00").Julian().JD())

	// Output:
	// 2.460334e+06
}

func ExampleCreateFromJulian() {
	fmt.Println(carbon.CreateFromJulian(2460334).ToDateTimeString())
	fmt.Println(carbon.CreateFromJulian(60333.5).ToDateTimeString())

	// Output:
	// 2024-01-24 12:00:00
	// 2024-01-24 12:00:00
}

func ExampleCarbon_Persian() {
	fmt.Println(carbon.Parse("1800-01-01 00:00:00").Persian().String())
	fmt.Println(carbon.Parse("2020-08-05 13:14:15").Persian().String())
	fmt.Println(carbon.Parse("2024-01-01 00:00:00").Persian().String())

	// Output:
	// 1178-10-11 00:00:00
	// 1399-05-15 13:14:15
	// 1402-10-11 00:00:00
}

func ExampleCreateFromPersian() {
	fmt.Println(carbon.CreateFromPersian(1178, 10, 11, 0, 0, 0).ToDateTimeString())
	fmt.Println(carbon.CreateFromPersian(1402, 10, 11, 0, 0, 0).ToDateTimeString())
	fmt.Println(carbon.CreateFromPersian(1403, 5, 15, 12, 0, 0).ToDateTimeString())

	// Output:
	// 1800-01-01 00:00:00
	// 2024-01-01 00:00:00
	// 2024-08-05 12:00:00
}
