package carbon_test

import (
	"fmt"
	"time"

	"github.com/dromara/carbon/v2"
)

func ExampleSetLayout() {
	defer carbon.SetLayout(carbon.DateTimeLayout)

	fmt.Println(carbon.SetLayout(carbon.DateTimeLayout).CurrentLayout())
	fmt.Println(carbon.SetLayout(carbon.TimeLayout).CurrentLayout())

	// Output:
	// 2006-01-02 15:04:05
	// 15:04:05
}

func ExampleSetFormat() {
	defer carbon.SetFormat(carbon.DateTimeFormat)

	fmt.Println(carbon.SetFormat(carbon.DateTimeFormat).CurrentLayout())
	fmt.Println(carbon.SetFormat(carbon.TimeFormat).CurrentLayout())

	// Output:
	// 2006-01-02 15:04:05
	// 15:04:05
}

func ExampleSetWeekStartsAt() {
	defer carbon.SetWeekStartsAt(carbon.Sunday)

	fmt.Println(carbon.SetWeekStartsAt(carbon.Sunday).WeekStartsAt())
	fmt.Println(carbon.SetWeekStartsAt(carbon.Monday).WeekStartsAt())

	// Output:
	// Sunday
	// Monday
}

func ExampleSetTimezone() {
	defer carbon.SetTimezone(carbon.UTC)

	fmt.Println(carbon.SetTimezone(carbon.UTC).ZoneName())
	fmt.Println(carbon.SetTimezone(carbon.PRC).ZoneName())

	fmt.Println(carbon.SetTimezone(carbon.UTC).Timezone())
	fmt.Println(carbon.SetTimezone(carbon.PRC).Timezone())

	// Output:
	// UTC
	// LMT
	// UTC
	// PRC
}

func ExampleSetLocation() {
	defer carbon.SetLocation(time.UTC)

	loc, _ := time.LoadLocation(carbon.PRC)
	carbon.SetLocation(loc)
	c := carbon.Parse("2020-08-05")

	fmt.Println(carbon.DefaultTimezone)
	fmt.Println(c.Timezone())
	fmt.Println(c.ZoneName())

	// Output:
	// PRC
	// PRC
	// CST
}

func ExampleCarbon_SetLocation() {
	defer carbon.SetLocation(time.UTC)

	c := carbon.Parse("2020-08-05")

	loc, _ := time.LoadLocation(carbon.PRC)
	c.SetLocation(loc)
	fmt.Println(carbon.DefaultTimezone)
	fmt.Println(c.Timezone())
	fmt.Println(c.ZoneName())

	// Output:
	// UTC
	// PRC
	// CST
}

func ExampleSetLocale() {
	defer carbon.SetLocale("en")

	carbon.SetLocale("zh-CN")

	c := carbon.Parse("2020-08-05")

	fmt.Println(c.Constellation())
	fmt.Println(c.Season())
	fmt.Println(c.ToMonthString())
	fmt.Println(c.ToShortMonthString())
	fmt.Println(c.ToWeekString())
	fmt.Println(c.ToShortWeekString())

	// Output:
	// 狮子座
	// 夏季
	// 八月
	// 8月
	// 星期三
	// 周三
}

func ExampleSetLanguage() {
	lang := carbon.NewLanguage()

	lang.SetLocale("en")
	fmt.Println(carbon.Parse("2020-08-05").SetLanguage(lang).ToMonthString())

	lang.SetLocale("zh-CN")
	fmt.Println(carbon.Parse("2020-08-05").SetLanguage(lang).ToMonthString())

	// Output:
	// August
	// 八月
}

func ExampleCarbon_SetDateTime() {
	fmt.Println(carbon.Parse("2020-08-05").SetDateTime(2020, 8, 5, 13, 14, 15).ToString())

	// Output:
	// 2020-08-05 13:14:15 +0000 UTC
}

func ExampleCarbon_SetDateTimeMilli() {
	fmt.Println(carbon.Parse("2020-08-05").SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999).ToString())

	// Output:
	// 2020-08-05 13:14:15.999 +0000 UTC
}

func ExampleCarbon_SetDateTimeMicro() {
	fmt.Println(carbon.Parse("2020-08-05").SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString())

	// Output:
	// 2020-08-05 13:14:15.999999 +0000 UTC
}

func ExampleCarbon_SetDateTimeNano() {
	fmt.Println(carbon.Parse("2020-08-05").SetDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999).ToString())

	// Output:
	// 2020-08-05 13:14:15.999999999 +0000 UTC
}

func ExampleCarbon_SetDate() {
	fmt.Println(carbon.Parse("2020-08-05").SetDate(2020, 8, 5).ToString())

	// Output:
	// 2020-08-05 00:00:00 +0000 UTC
}

func ExampleCarbon_SetDateMilli() {
	fmt.Println(carbon.Parse("2020-08-05").SetDateMilli(2020, 8, 5, 999).ToString())

	// Output:
	// 2020-08-05 00:00:00.999 +0000 UTC
}

func ExampleCarbon_SetDateMicro() {
	fmt.Println(carbon.Parse("2020-08-05").SetDateMicro(2020, 8, 5, 999999).ToString())

	// Output:
	// 2020-08-05 00:00:00.999999 +0000 UTC
}

func ExampleCarbon_SetDateNano() {
	fmt.Println(carbon.Parse("2020-08-05").SetDateNano(2020, 8, 5, 999999999).ToString())

	// Output:
	// 2020-08-05 00:00:00.999999999 +0000 UTC
}

func ExampleCarbon_SetTime() {
	fmt.Println(carbon.Parse("2020-08-05").SetTime(13, 14, 15).ToString())

	// Output:
	// 2020-08-05 13:14:15 +0000 UTC
}

func ExampleCarbon_SetTimeMilli() {
	fmt.Println(carbon.Parse("2020-08-05").SetTimeMilli(13, 14, 15, 999).ToString())

	// Output:
	// 2020-08-05 13:14:15.999 +0000 UTC
}

func ExampleCarbon_SetTimeMicro() {
	fmt.Println(carbon.Parse("2020-08-05").SetTimeMicro(13, 14, 15, 999999).ToString())

	// Output:
	// 2020-08-05 13:14:15.999999 +0000 UTC
}

func ExampleCarbon_SetTimeNano() {
	fmt.Println(carbon.Parse("2020-08-05").SetTimeNano(13, 14, 15, 999999999).ToString())

	// Output:
	// 2020-08-05 13:14:15.999999999 +0000 UTC
}

func ExampleCarbon_SetYear() {
	fmt.Println(carbon.Parse("2020-01-01").SetYear(2019).ToString())
	fmt.Println(carbon.Parse("2020-01-31").SetYear(2019).ToString())
	fmt.Println(carbon.Parse("2020-02-29").SetYear(2019).ToString())

	// Output:
	// 2019-01-01 00:00:00 +0000 UTC
	// 2019-01-31 00:00:00 +0000 UTC
	// 2019-03-01 00:00:00 +0000 UTC
}

func ExampleCarbon_SetYearNoOverflow() {
	fmt.Println(carbon.Parse("2020-01-01").SetYearNoOverflow(2019).ToString())
	fmt.Println(carbon.Parse("2020-01-31").SetYearNoOverflow(2019).ToString())
	fmt.Println(carbon.Parse("2020-02-29").SetYearNoOverflow(2019).ToString())

	// Output:
	// 2019-01-01 00:00:00 +0000 UTC
	// 2019-01-31 00:00:00 +0000 UTC
	// 2019-02-28 00:00:00 +0000 UTC
}

func ExampleCarbon_SetMonth() {
	fmt.Println(carbon.Parse("2020-01-01").SetMonth(2).ToString())
	fmt.Println(carbon.Parse("2020-01-30").SetMonth(2).ToString())
	fmt.Println(carbon.Parse("2020-01-31").SetMonth(2).ToString())

	// Output:
	// 2020-02-01 00:00:00 +0000 UTC
	// 2020-03-01 00:00:00 +0000 UTC
	// 2020-03-02 00:00:00 +0000 UTC
}

func ExampleCarbon_SetMonthNoOverflow() {
	fmt.Println(carbon.Parse("2020-01-01").SetMonthNoOverflow(2).ToString())
	fmt.Println(carbon.Parse("2020-01-30").SetMonthNoOverflow(2).ToString())
	fmt.Println(carbon.Parse("2020-01-31").SetMonthNoOverflow(2).ToString())

	// Output:
	// 2020-02-01 00:00:00 +0000 UTC
	// 2020-02-29 00:00:00 +0000 UTC
	// 2020-02-29 00:00:00 +0000 UTC
}

func ExampleCarbon_SetDay() {
	fmt.Println(carbon.Parse("2020-01-01").SetDay(31).ToString())
	fmt.Println(carbon.Parse("2020-02-01").SetDay(31).ToString())
	fmt.Println(carbon.Parse("2020-02-28").SetDay(31).ToString())

	// Output:
	// 2020-01-31 00:00:00 +0000 UTC
	// 2020-03-02 00:00:00 +0000 UTC
	// 2020-03-02 00:00:00 +0000 UTC
}

func ExampleCarbon_SetHour() {
	fmt.Println(carbon.Parse("2020-01-01").SetHour(10).ToString())
	fmt.Println(carbon.Parse("2020-02-01").SetHour(24).ToString())
	fmt.Println(carbon.Parse("2020-02-28").SetHour(31).ToString())

	// Output:
	// 2020-01-01 10:00:00 +0000 UTC
	// 2020-02-02 00:00:00 +0000 UTC
	// 2020-02-29 07:00:00 +0000 UTC
}

func ExampleCarbon_SetMinute() {
	fmt.Println(carbon.Parse("2020-01-01").SetMinute(10).ToString())
	fmt.Println(carbon.Parse("2020-02-01").SetMinute(24).ToString())
	fmt.Println(carbon.Parse("2020-02-28").SetMinute(60).ToString())

	// Output:
	// 2020-01-01 00:10:00 +0000 UTC
	// 2020-02-01 00:24:00 +0000 UTC
	// 2020-02-28 01:00:00 +0000 UTC
}

func ExampleCarbon_SetSecond() {
	fmt.Println(carbon.Parse("2020-01-01").SetSecond(10).ToString())
	fmt.Println(carbon.Parse("2020-02-01").SetSecond(24).ToString())
	fmt.Println(carbon.Parse("2020-02-28").SetSecond(60).ToString())

	// Output:
	// 2020-01-01 00:00:10 +0000 UTC
	// 2020-02-01 00:00:24 +0000 UTC
	// 2020-02-28 00:01:00 +0000 UTC
}

func ExampleCarbon_SetMillisecond() {
	fmt.Println(carbon.Parse("2020-01-01").SetMillisecond(100).ToString())
	fmt.Println(carbon.Parse("2020-01-01").SetMillisecond(999).ToString())

	// Output:
	// 2020-01-01 00:00:00.1 +0000 UTC
	// 2020-01-01 00:00:00.999 +0000 UTC
}

func ExampleCarbon_SetMicrosecond() {
	fmt.Println(carbon.Parse("2020-01-01").SetMicrosecond(100000).ToString())
	fmt.Println(carbon.Parse("2020-01-01").SetMicrosecond(999999).ToString())

	// Output:
	// 2020-01-01 00:00:00.1 +0000 UTC
	// 2020-01-01 00:00:00.999999 +0000 UTC
}

func ExampleCarbon_SetNanosecond() {
	fmt.Println(carbon.Parse("2020-01-01").SetNanosecond(100000000).ToString())
	fmt.Println(carbon.Parse("2020-01-01").SetNanosecond(999999999).ToString())

	// Output:
	// 2020-01-01 00:00:00.1 +0000 UTC
	// 2020-01-01 00:00:00.999999999 +0000 UTC
}
