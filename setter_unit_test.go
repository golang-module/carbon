package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_SetLayout(t *testing.T) {
	defer SetLayout(DateTimeLayout)

	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().SetLayout(DateLayout)
		assert.Equal(t, DateLayout, c.layout)
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetLayout(DateLayout).CurrentLayout())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetLayout(DateLayout).CurrentLayout())
		assert.Empty(t, Parse("0").SetLayout(DateLayout).CurrentLayout())
		assert.Empty(t, Parse("xxx").SetLayout(DateLayout).CurrentLayout())
	})

	t.Run("valid time", func(t *testing.T) {
		SetLayout(DateLayout)

		c := Parse("2020-08-05")
		assert.Equal(t, DateLayout, DefaultLayout)
		assert.Equal(t, DateLayout, c.CurrentLayout())

		c.SetLayout(DateTimeLayout)
		assert.Equal(t, DateLayout, DefaultLayout)
		assert.Equal(t, DateTimeLayout, c.CurrentLayout())
	})
}

func TestCarbon_SetFormat(t *testing.T) {
	defer SetFormat(DateTimeFormat)

	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().SetFormat(DateFormat)
		assert.Equal(t, DateLayout, c.layout)
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetFormat(DateFormat).CurrentLayout())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetFormat(DateFormat).CurrentLayout())
		assert.Empty(t, Parse("0").SetFormat(DateFormat).CurrentLayout())
		assert.Empty(t, Parse("xxx").SetFormat(DateFormat).CurrentLayout())
	})

	t.Run("valid time", func(t *testing.T) {
		SetFormat(DateFormat)
		c := Parse("2020-08-05")

		assert.Equal(t, DateLayout, DefaultLayout)
		assert.Equal(t, DateLayout, c.CurrentLayout())

		c.SetFormat(DateTimeFormat)
		assert.Equal(t, DateLayout, DefaultLayout)
		assert.Equal(t, DateTimeLayout, c.CurrentLayout())
	})
}

func TestCarbon_SetWeekStartsAt(t *testing.T) {
	defer SetWeekStartsAt(Sunday)

	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().SetWeekStartsAt(Sunday)
		assert.Equal(t, "0000-12-31 00:00:00 +0000 UTC", c.StartOfWeek().ToString())

		c.SetWeekStartsAt(Monday)
		assert.Equal(t, "0000-12-25 00:00:00 +0000 UTC", c.StartOfWeek().ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetWeekStartsAt(Sunday).ToString())
	})

	t.Run("invalid day", func(t *testing.T) {
		assert.Empty(t, Parse("2020-08-05").SetWeekStartsAt("").ToString())
		assert.Empty(t, Parse("2020-08-05").SetWeekStartsAt("0").ToString())
		assert.Empty(t, Parse("2020-08-05").SetWeekStartsAt("xxx").ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetWeekStartsAt(Sunday).ToString())
		assert.Empty(t, Parse("xxx").SetWeekStartsAt(Sunday).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		SetWeekStartsAt(Monday)
		c := Parse("2020-08-05")

		assert.Equal(t, Monday, DefaultWeekStartsAt)
		assert.Equal(t, Monday, c.WeekStartsAt())

		c.SetWeekStartsAt(Sunday)
		assert.Equal(t, Monday, DefaultWeekStartsAt)
		assert.Equal(t, Sunday, c.WeekStartsAt())
	})
}

func TestCarbon_SetLocale(t *testing.T) {
	defer SetLocale("en")

	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().SetLocale("en")
		assert.Equal(t, "Capricorn", c.Constellation())
		assert.Equal(t, "Winter", c.Season())
		assert.Equal(t, "January", c.ToMonthString())
		assert.Equal(t, "Jan", c.ToShortMonthString())
		assert.Equal(t, "Monday", c.ToWeekString())
		assert.Equal(t, "Mon", c.ToShortWeekString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetLocale("en").ToString())
	})

	t.Run("invalid locale", func(t *testing.T) {
		assert.Empty(t, Parse("").SetLocale("").ToString())
		assert.Empty(t, Parse("xxx").SetLocale("xxx").ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetLocale("en").ToString())
		assert.Empty(t, Parse("0").SetLocale("en").ToString())
		assert.Empty(t, Parse("xxx").SetLocale("en").ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		SetLocale("zh-CN")
		c := Parse("2020-08-05")

		assert.Equal(t, "zh-CN", DefaultLocale)
		assert.Equal(t, "zh-CN", c.Locale())

		assert.Equal(t, "狮子座", c.Constellation())
		assert.Equal(t, "夏季", c.Season())
		assert.Equal(t, "八月", c.ToMonthString())
		assert.Equal(t, "8月", c.ToShortMonthString())
		assert.Equal(t, "星期三", c.ToWeekString())
		assert.Equal(t, "周三", c.ToShortWeekString())

		c.SetLocale("en")

		assert.Equal(t, "zh-CN", DefaultLocale)
		assert.Equal(t, "en", c.Locale())

		assert.Equal(t, "Leo", c.Constellation())
		assert.Equal(t, "Summer", c.Season())
		assert.Equal(t, "August", c.ToMonthString())
		assert.Equal(t, "Aug", c.ToShortMonthString())
		assert.Equal(t, "Wednesday", c.ToWeekString())
		assert.Equal(t, "Wed", c.ToShortWeekString())
	})
}

func TestCarbon_SetTimezone(t *testing.T) {
	defer SetTimezone(UTC)

	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().SetTimezone(UTC)
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", c.ToString())

		c.SetTimezone(PRC)
		assert.Equal(t, "0001-01-01 08:05:43 +0805 LMT", c.ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetTimezone(UTC).ToString())
	})

	t.Run("invalid timezone", func(t *testing.T) {
		assert.Empty(t, Parse("2020-08-05").SetTimezone("").ToString())
		assert.Empty(t, Parse("2020-08-05").SetTimezone("0").ToString())
		assert.Empty(t, Parse("2020-08-05").SetTimezone("XXX").ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetTimezone(UTC).ToString())
		assert.Empty(t, Parse("0").SetTimezone(PRC).ToString())
		assert.Empty(t, Parse("xxx").SetTimezone(PRC).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		SetTimezone(PRC)
		c := Parse("2020-08-05")

		assert.Equal(t, PRC, DefaultTimezone)
		assert.Equal(t, PRC, c.Timezone())
		assert.Equal(t, "CST", c.ZoneName())

		c.SetTimezone(UTC)
		assert.Equal(t, PRC, DefaultTimezone)
		assert.Equal(t, UTC, c.Timezone())
		assert.Equal(t, UTC, c.ZoneName())
	})
}

func TestCarbon_SetLocation(t *testing.T) {
	defer SetLocation(time.UTC)

	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().SetLocation(time.UTC)
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", c.ToString())

		loc, _ := time.LoadLocation(PRC)
		c.SetLocation(loc)
		assert.Equal(t, "0001-01-01 08:05:43 +0805 LMT", c.ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetLocation(time.UTC).ToString())
	})

	t.Run("nil location", func(t *testing.T) {
		assert.Empty(t, Parse("2020-08-05").SetLocation(nil).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetLocation(time.UTC).ToString())
		assert.Empty(t, Parse("0").SetLocation(time.UTC).ToString())
		assert.Empty(t, Parse("xxx").SetLocation(time.UTC).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		loc, _ := time.LoadLocation(PRC)
		SetLocation(loc)
		c := Parse("2020-08-05")

		assert.Equal(t, PRC, DefaultTimezone)
		assert.Equal(t, PRC, c.Timezone())
		assert.Equal(t, "CST", c.ZoneName())

		c.SetLocation(time.UTC)
		assert.Equal(t, PRC, DefaultTimezone)
		assert.Equal(t, UTC, c.Timezone())
		assert.Equal(t, UTC, c.ZoneName())
	})
}

func TestCarbon_SetLanguage(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		lang := NewLanguage()
		lang.SetLocale("en")
		c := NewCarbon().SetLanguage(lang)
		assert.Equal(t, "en", c.Locale())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		lang := NewLanguage()
		lang.SetLocale("en")
		assert.Empty(t, c.SetLanguage(lang).ToString())
	})

	t.Run("nil language", func(t *testing.T) {
		lang := NewLanguage()
		lang = nil
		assert.Empty(t, SetLanguage(lang).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		lang := NewLanguage()
		lang.SetLocale("en")
		assert.Empty(t, Parse("").SetLanguage(lang).ToString())
		assert.Empty(t, Parse("0").SetLanguage(lang).ToString())
		assert.Empty(t, Parse("xxx").SetLanguage(lang).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		lang := NewLanguage()

		lang.SetLocale("en")
		assert.Equal(t, "August", Parse("2020-08-05").SetLanguage(lang).ToMonthString())

		lang.SetLocale("zh-CN")
		assert.Equal(t, "八月", Parse("2020-08-05").SetLanguage(lang).ToMonthString())
	})
}

func TestCarbon_SetDateTime(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", NewCarbon().SetDateTime(2020, 8, 5, 13, 14, 15).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetDateTime(2020, 8, 5, 13, 14, 15).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetDateTime(2020, 8, 5, 13, 14, 15).ToString())
		assert.Empty(t, Parse("0").SetDateTime(2020, 8, 5, 13, 14, 15).ToString())
		assert.Empty(t, Parse("xxx").SetDateTime(2020, 8, 5, 13, 14, 15).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", Parse("2020-08-05").SetDateTime(2020, 8, 5, 13, 14, 15).ToString())
	})
}

func TestCarbon_SetDateTimeMilli(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15.999 +0000 UTC", NewCarbon().SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999).ToString())
		assert.Empty(t, Parse("0").SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999).ToString())
		assert.Empty(t, Parse("xxx").SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15.999 +0000 UTC", Parse("2020-08-05").SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999).ToString())
	})
}

func TestCarbon_SetDateTimeMicro(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15.999999 +0000 UTC", NewCarbon().SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString())
		assert.Empty(t, Parse("0").SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString())
		assert.Empty(t, Parse("xxx").SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15.999999 +0000 UTC", Parse("2020-08-05").SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString())
	})
}

func TestCarbon_SetDateTimeNano(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15.999999999 +0000 UTC", NewCarbon().SetDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999).ToString())
		assert.Empty(t, Parse("0").SetDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999).ToString())
		assert.Empty(t, Parse("xxx").SetDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15.999999999 +0000 UTC", Parse("2020-08-05").SetDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999).ToString())
	})
}

func TestCarbon_SetDate(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", NewCarbon().SetDate(2020, 8, 5).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetDate(2020, 8, 5).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetDate(2020, 8, 5).ToString())
		assert.Empty(t, Parse("0").SetDate(2020, 8, 5).ToString())
		assert.Empty(t, Parse("xxx").SetDate(2020, 8, 5).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", Parse("2020-08-05").SetDate(2020, 8, 5).ToString())
	})
}

func TestCarbon_SetDateMilli(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00.999 +0000 UTC", NewCarbon().SetDateMilli(2020, 8, 5, 999).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetDateMilli(2020, 8, 5, 999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetDateMilli(2020, 8, 5, 999).ToString())
		assert.Empty(t, Parse("0").SetDateMilli(2020, 8, 5, 999).ToString())
		assert.Empty(t, Parse("xxx").SetDateMilli(2020, 8, 5, 999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00.999 +0000 UTC", Parse("2020-08-05").SetDateMilli(2020, 8, 5, 999).ToString())
	})
}

func TestCarbon_SetDateMicro(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00.999999 +0000 UTC", NewCarbon().SetDateMicro(2020, 8, 5, 999999).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetDateMicro(2020, 8, 5, 999999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetDateMicro(2020, 8, 5, 999999).ToString())
		assert.Empty(t, Parse("0").SetDateMicro(2020, 8, 5, 999999).ToString())
		assert.Empty(t, Parse("xxx").SetDateMicro(2020, 8, 5, 999999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00.999999 +0000 UTC", Parse("2020-08-05").SetDateMicro(2020, 8, 5, 999999).ToString())
	})
}

func TestCarbon_SetDateNano(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00.999999999 +0000 UTC", NewCarbon().SetDateNano(2020, 8, 5, 999999999).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetDateNano(2020, 8, 5, 999999999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetDateNano(2020, 8, 5, 999999999).ToString())
		assert.Empty(t, Parse("0").SetDateNano(2020, 8, 5, 999999999).ToString())
		assert.Empty(t, Parse("xxx").SetDateNano(2020, 8, 5, 999999999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00.999999999 +0000 UTC", Parse("2020-08-05").SetDateNano(2020, 8, 5, 999999999).ToString())
	})
}

func TestCarbon_SetTime(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 13:14:15 +0000 UTC", NewCarbon().SetTime(13, 14, 15).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetTime(13, 14, 15).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetTime(13, 14, 15).ToString())
		assert.Empty(t, Parse("0").SetTime(13, 14, 15).ToString())
		assert.Empty(t, Parse("xxx").SetTime(13, 14, 15).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", Parse("2020-08-05").SetTime(13, 14, 15).ToString())
	})
}

func TestCarbon_SetTimeMilli(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 13:14:15.999 +0000 UTC", NewCarbon().SetTimeMilli(13, 14, 15, 999).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetTimeMilli(13, 14, 15, 999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetTimeMilli(13, 14, 15, 999).ToString())
		assert.Empty(t, Parse("0").SetTimeMilli(13, 14, 15, 999).ToString())
		assert.Empty(t, Parse("xxx").SetTimeMilli(13, 14, 15, 999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15.999 +0000 UTC", Parse("2020-08-05").SetTimeMilli(13, 14, 15, 999).ToString())
	})
}

func TestCarbon_SetTimeMicro(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 13:14:15.999999 +0000 UTC", NewCarbon().SetTimeMicro(13, 14, 15, 999999).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetTimeMicro(13, 14, 15, 9999999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetTimeMicro(13, 14, 15, 9999999).ToString())
		assert.Empty(t, Parse("0").SetTimeMicro(13, 14, 15, 9999999).ToString())
		assert.Empty(t, Parse("xxx").SetTimeMicro(13, 14, 15, 9999999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15.999999 +0000 UTC", Parse("2020-08-05").SetTimeMicro(13, 14, 15, 999999).ToString())
	})
}

func TestCarbon_SetTimeNano(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 13:14:15.999999999 +0000 UTC", NewCarbon().SetTimeNano(13, 14, 15, 999999999).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetTimeNano(13, 14, 15, 999999999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetTimeNano(13, 14, 15, 999999999).ToString())
		assert.Empty(t, Parse("0").SetTimeNano(13, 14, 15, 999999999).ToString())
		assert.Empty(t, Parse("xxx").SetTimeNano(13, 14, 15, 999999999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15.999999999 +0000 UTC", Parse("2020-08-05").SetTimeNano(13, 14, 15, 999999999).ToString())
	})
}

func TestCarbon_SetYear(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", NewCarbon().SetYear(2020).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetYear(2020).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetYear(2020).ToString())
		assert.Empty(t, Parse("0").SetYear(2020).ToString())
		assert.Empty(t, Parse("xxx").SetYear(2020).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2019-01-01 00:00:00 +0000 UTC", Parse("2020-01-01").SetYear(2019).ToString())
		assert.Equal(t, "2019-01-31 00:00:00 +0000 UTC", Parse("2020-01-31").SetYear(2019).ToString())
		assert.Equal(t, "2019-03-01 00:00:00 +0000 UTC", Parse("2020-02-29").SetYear(2019).ToString())
	})
}

func TestCarbon_SetYearNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", NewCarbon().SetYearNoOverflow(2020).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetYearNoOverflow(2020).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetYearNoOverflow(2020).ToString())
		assert.Empty(t, Parse("0").SetYearNoOverflow(2020).ToString())
		assert.Empty(t, Parse("xxx").SetYearNoOverflow(2020).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2019-01-01 00:00:00 +0000 UTC", Parse("2020-01-01").SetYearNoOverflow(2019).ToString())
		assert.Equal(t, "2019-01-31 00:00:00 +0000 UTC", Parse("2020-01-31").SetYearNoOverflow(2019).ToString())
		assert.Equal(t, "2019-02-28 00:00:00 +0000 UTC", Parse("2020-02-29").SetYearNoOverflow(2019).ToString())
	})
}

func TestCarbon_SetMonth(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-02-01 00:00:00 +0000 UTC", NewCarbon().SetMonth(2).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetMonth(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetMonth(2).ToString())
		assert.Empty(t, Parse("0").SetMonth(2).ToString())
		assert.Empty(t, Parse("xxx").SetMonth(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-02-01 00:00:00 +0000 UTC", Parse("2020-01-01").SetMonth(2).ToString())
		assert.Equal(t, "2020-03-01 00:00:00 +0000 UTC", Parse("2020-01-30").SetMonth(2).ToString())
		assert.Equal(t, "2020-03-02 00:00:00 +0000 UTC", Parse("2020-01-31").SetMonth(2).ToString())
	})
}

func TestCarbon_SetMonthNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-02-01 00:00:00 +0000 UTC", NewCarbon().SetMonthNoOverflow(2).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetMonthNoOverflow(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetMonthNoOverflow(2).ToString())
		assert.Empty(t, Parse("0").SetMonthNoOverflow(2).ToString())
		assert.Empty(t, Parse("xxx").SetMonthNoOverflow(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-02-01 00:00:00 +0000 UTC", Parse("2020-01-01").SetMonthNoOverflow(2).ToString())
		assert.Equal(t, "2020-02-29 00:00:00 +0000 UTC", Parse("2020-01-30").SetMonthNoOverflow(2).ToString())
		assert.Equal(t, "2020-02-29 00:00:00 +0000 UTC", Parse("2020-01-31").SetMonthNoOverflow(2).ToString())
	})
}

func TestCarbon_SetDay(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-31 00:00:00 +0000 UTC", NewCarbon().SetDay(31).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetDay(31).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetDay(31).ToString())
		assert.Empty(t, Parse("0").SetDay(31).ToString())
		assert.Empty(t, Parse("xxx").SetDay(31).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-31 00:00:00 +0000 UTC", Parse("2020-01-01").SetDay(31).ToString())
		assert.Equal(t, "2020-03-02 00:00:00 +0000 UTC", Parse("2020-02-01").SetDay(31).ToString())
		assert.Equal(t, "2020-03-02 00:00:00 +0000 UTC", Parse("2020-02-28").SetDay(31).ToString())
	})
}

func TestCarbon_SetHour(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 10:00:00 +0000 UTC", NewCarbon().SetHour(10).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetHour(10).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetHour(31).ToString())
		assert.Empty(t, Parse("0").SetHour(31).ToString())
		assert.Empty(t, Parse("xxx").SetHour(31).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 10:00:00 +0000 UTC", Parse("2020-01-01").SetHour(10).ToString())
		assert.Equal(t, "2020-02-02 00:00:00 +0000 UTC", Parse("2020-02-01").SetHour(24).ToString())
		assert.Equal(t, "2020-02-29 07:00:00 +0000 UTC", Parse("2020-02-28").SetHour(31).ToString())
	})
}

func TestCarbon_SetMinute(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:10:00 +0000 UTC", NewCarbon().SetMinute(10).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetMinute(10).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetMinute(31).ToString())
		assert.Empty(t, Parse("0").SetMinute(31).ToString())
		assert.Empty(t, Parse("xxx").SetMinute(31).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:10:00 +0000 UTC", Parse("2020-01-01").SetMinute(10).ToString())
		assert.Equal(t, "2020-02-01 00:24:00 +0000 UTC", Parse("2020-02-01").SetMinute(24).ToString())
		assert.Equal(t, "2020-02-28 01:00:00 +0000 UTC", Parse("2020-02-28").SetMinute(60).ToString())
	})
}

func TestCarbon_SetSecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:10 +0000 UTC", NewCarbon().SetSecond(10).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetSecond(10).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetSecond(31).ToString())
		assert.Empty(t, Parse("0").SetSecond(31).ToString())
		assert.Empty(t, Parse("xxx").SetSecond(31).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:10 +0000 UTC", Parse("2020-01-01").SetSecond(10).ToString())
		assert.Equal(t, "2020-02-01 00:00:24 +0000 UTC", Parse("2020-02-01").SetSecond(24).ToString())
		assert.Equal(t, "2020-02-28 00:01:00 +0000 UTC", Parse("2020-02-28").SetSecond(60).ToString())
	})
}

func TestCarbon_SetMillisecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00.999 +0000 UTC", NewCarbon().SetMillisecond(999).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetMillisecond(999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetMillisecond(999).ToString())
		assert.Empty(t, Parse("0").SetMillisecond(999).ToString())
		assert.Empty(t, Parse("xxx").SetMillisecond(999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00.1 +0000 UTC", Parse("2020-01-01").SetMillisecond(100).ToString())
		assert.Equal(t, "2020-01-01 00:00:00.999 +0000 UTC", Parse("2020-01-01").SetMillisecond(999).ToString())
	})
}

func TestCarbon_SetMicrosecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00.999999 +0000 UTC", NewCarbon().SetMicrosecond(999999).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetMicrosecond(999999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetMicrosecond(999999).ToString())
		assert.Empty(t, Parse("0").SetMicrosecond(999999).ToString())
		assert.Empty(t, Parse("xxx").SetMicrosecond(999999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00.1 +0000 UTC", Parse("2020-01-01").SetMicrosecond(100000).ToString())
		assert.Equal(t, "2020-01-01 00:00:00.999999 +0000 UTC", Parse("2020-01-01").SetMicrosecond(999999).ToString())
	})
}

func TestCarbon_SetNanosecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00.999999999 +0000 UTC", NewCarbon().SetNanosecond(999999999).ToString())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.SetNanosecond(999999999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetNanosecond(999999999).ToString())
		assert.Empty(t, Parse("0").SetNanosecond(999999999).ToString())
		assert.Empty(t, Parse("xxx").SetNanosecond(999999999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00.1 +0000 UTC", Parse("2020-01-01").SetNanosecond(100000000).ToString())
		assert.Equal(t, "2020-01-01 00:00:00.999999999 +0000 UTC", Parse("2020-01-01").SetNanosecond(999999999).ToString())
	})
}
