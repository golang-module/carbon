package carbon_test

import (
	"fmt"

	"github.com/dromara/carbon/v2"
)

func ExampleParse() {
	fmt.Println(carbon.Parse("2020-8-5").ToString())
	fmt.Println(carbon.Parse("2020-8-05").ToString())
	fmt.Println(carbon.Parse("2020-08-05").ToString())

	fmt.Println(carbon.Parse("2020-8-5 1:2:3").ToString())
	fmt.Println(carbon.Parse("2020-08-05 1:2:03").ToString())
	fmt.Println(carbon.Parse("2020-08-05 1:02:03").ToString())
	fmt.Println(carbon.Parse("2020-08-05 01:02:03").ToString())

	fmt.Println(carbon.Parse("2023-01-08T09:02:48").ToString())
	fmt.Println(carbon.Parse("2023-1-8T09:02:48").ToString())
	fmt.Println(carbon.Parse("2023-01-08T9:2:48").ToString())
	fmt.Println(carbon.Parse("2023-01-8T9:2:48").ToString())

	fmt.Println(carbon.Parse("0000-01-01 00:00:00").ToString())
	fmt.Println(carbon.Parse("0001-01-01 00:00:00").ToString())

	// Output:
	// 2020-08-05 00:00:00 +0000 UTC
	// 2020-08-05 00:00:00 +0000 UTC
	// 2020-08-05 00:00:00 +0000 UTC
	// 2020-08-05 01:02:03 +0000 UTC
	// 2020-08-05 01:02:03 +0000 UTC
	// 2020-08-05 01:02:03 +0000 UTC
	// 2020-08-05 01:02:03 +0000 UTC
	// 2023-01-08 09:02:48 +0000 UTC
	// 2023-01-08 09:02:48 +0000 UTC
	// 2023-01-08 09:02:48 +0000 UTC
	// 2023-01-08 09:02:48 +0000 UTC
	// 0000-01-01 00:00:00 +0000 UTC
	// 0001-01-01 00:00:00 +0000 UTC
}

func ExampleParseByFormat() {
	fmt.Println(carbon.ParseByFormat("2020-08-05", carbon.DateFormat).ToString())
	fmt.Println(carbon.ParseByFormat("2020-08-05 13:14:15", carbon.DateTimeFormat).ToString())
	fmt.Println(carbon.ParseByFormat("2020|08|05 13:14:15", "Y|m|d H:i:s").ToString())

	fmt.Println(carbon.ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToString())
	fmt.Println(carbon.ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒").ToString())

	fmt.Println(carbon.ParseByFormat("1699677240", "U").ToString())
	fmt.Println(carbon.ParseByFormat("1699677240666", "V").ToString())
	fmt.Println(carbon.ParseByFormat("1699677240666666", "X").ToString())
	fmt.Println(carbon.ParseByFormat("1699677240666666666", "Z").ToString())

	// Output:
	// 2020-08-05 00:00:00 +0000 UTC
	// 2020-08-05 13:14:15 +0000 UTC
	// 2020-08-05 13:14:15 +0000 UTC
	// 2020-08-05 13:14:15 +0000 UTC
	// 2020-08-05 13:14:15 +0000 UTC
	// 2023-11-11 04:34:00 +0000 UTC
	// 2023-11-11 04:34:00.666 +0000 UTC
	// 2023-11-11 04:34:00.666666 +0000 UTC
	// 2023-11-11 04:34:00.666666666 +0000 UTC
}

func ExampleParseByLayout() {
	fmt.Println(carbon.ParseByLayout("2020-08-05", carbon.DateLayout).ToString())
	fmt.Println(carbon.ParseByLayout("2020-08-05 13:14:15", carbon.DateTimeLayout, carbon.PRC).ToString())
	fmt.Println(carbon.ParseByLayout("2020|08|05 13:14:15", "2006|01|02 15:04:05").ToString())
	fmt.Println(carbon.ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒").ToString())

	// Output:
	// 2020-08-05 00:00:00 +0000 UTC
	// 2020-08-05 13:14:15 +0800 CST
	// 2020-08-05 13:14:15 +0000 UTC
	// 2020-08-05 13:14:15 +0000 UTC
}
