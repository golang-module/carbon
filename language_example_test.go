package carbon_test

import (
	"fmt"

	"github.com/dromara/carbon/v2"
)

func ExampleLanguage_SetLocale() {
	lang := carbon.NewLanguage()

	lang.SetLocale("en")
	fmt.Println(carbon.Parse("2020-08-05").SetLanguage(lang).Constellation())
	fmt.Println(carbon.Parse("2020-08-05").SetLanguage(lang).Season())
	fmt.Println(carbon.Parse("2020-08-05").SetLanguage(lang).DiffForHumans(carbon.Parse("2024-08-05")))
	fmt.Println(carbon.Parse("2020-08-05").SetLanguage(lang).ToMonthString())
	fmt.Println(carbon.Parse("2020-08-05").SetLanguage(lang).ToShortMonthString())
	fmt.Println(carbon.Parse("2020-08-05").SetLanguage(lang).ToWeekString())
	fmt.Println(carbon.Parse("2020-08-05").SetLanguage(lang).ToShortWeekString())

	lang.SetLocale("zh-CN")
	fmt.Println(carbon.Parse("2020-08-05").SetLanguage(lang).Constellation())
	fmt.Println(carbon.Parse("2020-08-05").SetLanguage(lang).Season())
	fmt.Println(carbon.Parse("2020-08-05").SetLanguage(lang).DiffForHumans(carbon.Parse("2024-08-05")))
	fmt.Println(carbon.Parse("2020-08-05").SetLanguage(lang).ToMonthString())
	fmt.Println(carbon.Parse("2020-08-05").SetLanguage(lang).ToShortMonthString())
	fmt.Println(carbon.Parse("2020-08-05").SetLanguage(lang).ToWeekString())
	fmt.Println(carbon.Parse("2020-08-05").SetLanguage(lang).ToShortWeekString())

	// Output:
	// Leo
	// Summer
	// 4 years before
	// August
	// Aug
	// Wednesday
	// Wed
	// 狮子座
	// 夏季
	// 4 年前
	// 八月
	// 8月
	// 星期三
	// 周三
}

func ExampleLanguage_SetResources() {
	resources1 := map[string]string{
		"months":       "Ⅰ月|Ⅱ月|Ⅲ月|Ⅳ月|Ⅴ月|Ⅵ月|Ⅶ月|Ⅷ月|Ⅸ月|Ⅹ月|Ⅺ月|Ⅻ月",
		"short_months": "Ⅰ|Ⅱ|Ⅲ|Ⅳ|Ⅴ|Ⅵ|Ⅶ|Ⅷ|Ⅸ|Ⅹ|Ⅺ|Ⅻ",
	}

	lang1 := carbon.NewLanguage()
	lang1.SetLocale("en").SetResources(resources1)

	fmt.Printf("lang1:%s\n", carbon.Parse("2020-08-05").SetLanguage(lang1).Constellation())
	fmt.Printf("lang1:%s\n", carbon.Parse("2020-08-05").SetLanguage(lang1).Season())
	fmt.Printf("lang1:%s\n", carbon.Parse("2020-08-05").SetLanguage(lang1).DiffForHumans(carbon.Parse("2024-08-05")))
	fmt.Printf("lang1:%s\n", carbon.Parse("2020-08-05").SetLanguage(lang1).ToMonthString())
	fmt.Printf("lang1:%s\n", carbon.Parse("2020-08-05").SetLanguage(lang1).ToShortMonthString())
	fmt.Printf("lang1:%s\n", carbon.Parse("2020-08-05").SetLanguage(lang1).ToWeekString())
	fmt.Printf("lang1:%s\n", carbon.Parse("2020-08-05").SetLanguage(lang1).ToShortWeekString())

	resources2 := map[string]string{
		"constellations": "Aries|Taurus|Gemini|Cancer|Leo|Virgo|Libra|Scorpio|Sagittarius|Capricorn|Aquarius|Pisces",
		"seasons":        "spring|summer|autumn|winter",
		"months":         "January|February|March|April|May|June|July|August|September|October|November|December",
		"short_months":   "Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec",
		"weeks":          "Sunday|Monday|Tuesday|Wednesday|Thursday|Friday|Saturday",
		"short_weeks":    "Sun|Mon|Tue|Wed|Thu|Fri|Sat",
		"year":           "1 yr|%d yrs",
		"month":          "1 mo|%d mos",
		"week":           "%dw",
		"day":            "%dd",
		"hour":           "%dh",
		"minute":         "%dm",
		"second":         "%ds",
		"now":            "just now",
		"ago":            "%s ago",
		"from_now":       "in %s",
		"before":         "%s before",
		"after":          "%s after",
	}

	lang2 := carbon.NewLanguage()
	lang2.SetResources(resources2)

	fmt.Printf("lang2:%s\n", carbon.Parse("2020-08-05").SetLanguage(lang2).Constellation())
	fmt.Printf("lang2:%s\n", carbon.Parse("2020-08-05").SetLanguage(lang2).Season())
	fmt.Printf("lang2:%s\n", carbon.Parse("2020-08-05").SetLanguage(lang2).DiffForHumans(carbon.Parse("2024-08-05")))
	fmt.Printf("lang2:%s\n", carbon.Parse("2020-08-05").SetLanguage(lang2).ToMonthString())
	fmt.Printf("lang2:%s\n", carbon.Parse("2020-08-05").SetLanguage(lang2).ToShortMonthString())
	fmt.Printf("lang2:%s\n", carbon.Parse("2020-08-05").SetLanguage(lang2).ToWeekString())
	fmt.Printf("lang2:%s\n", carbon.Parse("2020-08-05").SetLanguage(lang2).ToShortWeekString())

	// Output:
	// lang1:Leo
	// lang1:Summer
	// lang1:4 years before
	// lang1:Ⅷ月
	// lang1:Ⅷ
	// lang1:Wednesday
	// lang1:Wed
	// lang2:Leo
	// lang2:summer
	// lang2:4 yrs before
	// lang2:August
	// lang2:Aug
	// lang2:Wednesday
	// lang2:Wed
}
