package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_DiffInYears(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Zero(t, c.DiffInYears())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffInYears())
		assert.Zero(t, Now().DiffInYears(Parse("")))
		assert.Zero(t, Parse("xxx").DiffInYears())
		assert.Zero(t, Now().DiffInYears(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInYears(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-12-31 13:14:15").DiffInYears(Parse("2021-01-01 13:14:15")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffInYears(Parse("2021-08-28 13:14:59")))
		assert.Equal(t, int64(-1), Parse("2020-08-05 13:14:15").DiffInYears(Parse("2018-08-28 13:14:59")))
	})
}

func TestCarbon_DiffAbsInYears(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Zero(t, c.DiffAbsInYears())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffAbsInYears())
		assert.Zero(t, Now().DiffAbsInYears(Parse("")))
		assert.Zero(t, Parse("xxx").DiffAbsInYears())
		assert.Zero(t, Now().DiffAbsInYears(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInYears(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-12-31 13:14:15").DiffAbsInYears(Parse("2021-01-01 13:14:15")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInYears(Parse("2021-08-28 13:14:59")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInYears(Parse("2018-08-28 13:14:59")))
	})
}

func TestCarbon_DiffInMonths(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Zero(t, c.DiffInMonths())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffInMonths())
		assert.Zero(t, Now().DiffInMonths(Parse("")))
		assert.Zero(t, Parse("").DiffInMonths(Parse("")))
		assert.Zero(t, Parse("xxx").DiffInMonths())
		assert.Zero(t, Now().DiffInMonths(Parse("xxx")))
		assert.Zero(t, Parse("xxx").DiffInMonths(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInMonths(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInMonths(Parse("2020-07-28 13:14:00")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffInMonths(Parse("2020-09-06 13:14:59")))
		assert.Equal(t, int64(-23), Parse("2020-08-05 13:14:15").DiffInMonths(Parse("2018-08-28 13:14:59")))
	})
}

func TestCarbon_DiffAbsInMonths(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Zero(t, c.DiffAbsInMonths())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffAbsInMonths())
		assert.Zero(t, Now().DiffAbsInMonths(Parse("")))
		assert.Zero(t, Parse("xxx").DiffAbsInMonths())
		assert.Zero(t, Now().DiffAbsInMonths(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInMonths(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInMonths(Parse("2020-07-28 13:14:00")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInMonths(Parse("2020-09-06 13:14:59")))
		assert.Equal(t, int64(23), Parse("2020-08-05 13:14:15").DiffAbsInMonths(Parse("2018-08-28 13:14:59")))
	})
}

func TestCarbon_DiffInWeeks(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Zero(t, c.DiffInWeeks())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffInWeeks())
		assert.Zero(t, Now().DiffInWeeks(Parse("")))
		assert.Zero(t, Parse("xxx").DiffInWeeks())
		assert.Zero(t, Now().DiffInWeeks(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInWeeks(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(-1), Parse("2020-08-05 13:14:15").DiffInWeeks(Parse("2020-07-28 13:14:00")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffInWeeks(Parse("2020-08-12 13:14:15")))
	})
}

func TestCarbon_DiffAbsInWeeks(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Zero(t, c.DiffAbsInWeeks())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffAbsInWeeks())
		assert.Zero(t, Now().DiffAbsInWeeks(Parse("")))
		assert.Zero(t, Parse("xxx").DiffAbsInWeeks())
		assert.Zero(t, Now().DiffAbsInWeeks(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInWeeks(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInWeeks(Parse("2020-07-28 13:14:00")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInWeeks(Parse("2020-08-12 13:14:15")))
	})
}

func TestCarbon_DiffInDays(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Zero(t, c.DiffInDays())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffInDays())
		assert.Zero(t, Now().DiffInDays(Parse("")))
		assert.Zero(t, Parse("xxx").DiffInDays())
		assert.Zero(t, Now().DiffInDays(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInDays(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInDays(Parse("22020-08-04 13:14:59")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffInDays(Parse("2020-08-06 13:14:15")))
		assert.Equal(t, int64(-1), Parse("2020-08-05 13:14:15").DiffInDays(Parse("2020-08-04 13:00:00")))
	})
}

func TestCarbon_DiffAbsInDays(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Zero(t, c.DiffAbsInDays())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffAbsInDays())
		assert.Zero(t, Now().DiffAbsInDays(Parse("")))
		assert.Zero(t, Parse("xxx").DiffAbsInDays())
		assert.Zero(t, Now().DiffAbsInDays(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInDays(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInDays(Parse("22020-08-04 13:14:59")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInDays(Parse("2020-08-06 13:14:15")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInDays(Parse("2020-08-04 13:00:00")))
	})
}

func TestCarbon_DiffInHours(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Zero(t, c.DiffInHours())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffInHours())
		assert.Zero(t, Now().DiffInHours(Parse("")))
		assert.Zero(t, Parse("xxx").DiffInHours())
		assert.Zero(t, Now().DiffInHours(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInHours(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInHours(Parse("2020-08-05 14:13:15")))
		assert.Equal(t, int64(-1), Parse("2020-08-05 13:14:15").DiffInHours(Parse("2020-08-05 12:14:00")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffInHours(Parse("2020-08-05 14:14:15")))
	})
}

func TestCarbon_DiffAbsInHours(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Zero(t, c.DiffAbsInHours())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffAbsInHours())
		assert.Zero(t, Now().DiffAbsInHours(Parse("")))
		assert.Zero(t, Parse("xxx").DiffAbsInHours())
		assert.Zero(t, Now().DiffAbsInHours(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInHours(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInHours(Parse("2020-08-05 14:13:15")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInHours(Parse("2020-08-05 12:14:00")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInHours(Parse("2020-08-05 14:14:15")))
	})
}

func TestCarbon_DiffInMinutes(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Zero(t, c.DiffInMinutes())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffInMinutes())
		assert.Zero(t, Now().DiffInMinutes(Parse("")))
		assert.Zero(t, Parse("xxx").DiffInMinutes())
		assert.Zero(t, Now().DiffInMinutes(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInMinutes(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInMinutes(Parse("2020-08-05 13:15:10")))
		assert.Equal(t, int64(-1), Parse("2020-08-05 13:14:15").DiffInMinutes(Parse("2020-08-05 13:13:00")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffInMinutes(Parse("2020-08-05 13:15:15")))
	})
}

func TestCarbon_DiffAbsInMinutes(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Zero(t, c.DiffAbsInMinutes())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffAbsInMinutes())
		assert.Zero(t, Now().DiffAbsInMinutes(Parse("")))
		assert.Zero(t, Parse("xxx").DiffAbsInMinutes())
		assert.Zero(t, Now().DiffAbsInMinutes(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInMinutes(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInMinutes(Parse("2020-08-05 13:15:10")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInMinutes(Parse("2020-08-05 13:13:00")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInMinutes(Parse("2020-08-05 13:15:15")))
	})
}

func TestCarbon_DiffInSeconds(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Zero(t, c.DiffInSeconds())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffInSeconds())
		assert.Zero(t, Now().DiffInSeconds(Parse("")))
		assert.Zero(t, Parse("xxx").DiffInSeconds())
		assert.Zero(t, Now().DiffInSeconds(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInSeconds(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInSeconds(Parse("2020-08-05 13:14:15.999999")))
		assert.Equal(t, int64(5), Parse("2020-08-05 13:14:15").DiffInSeconds(Parse("2020-08-05 13:14:20")))
		assert.Equal(t, int64(-5), Parse("2020-08-05 13:14:15").DiffInSeconds(Parse("2020-08-05 13:14:10")))
	})
}

func TestCarbon_DiffAbsInSeconds(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Zero(t, c.DiffAbsInSeconds())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffAbsInSeconds())
		assert.Zero(t, Now().DiffAbsInSeconds(Parse("")))
		assert.Zero(t, Parse("xxx").DiffAbsInSeconds())
		assert.Zero(t, Now().DiffAbsInSeconds(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInSeconds(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInSeconds(Parse("2020-08-05 13:14:15.999999")))
		assert.Equal(t, int64(5), Parse("2020-08-05 13:14:15").DiffAbsInSeconds(Parse("2020-08-05 13:14:20")))
		assert.Equal(t, int64(5), Parse("2020-08-05 13:14:15").DiffAbsInSeconds(Parse("2020-08-05 13:14:10")))
	})
}

func TestCarbon_DiffInString(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.DiffInString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").DiffInString())
		assert.Empty(t, Now().DiffInString(Parse("")))
		assert.Empty(t, Parse("xxx").DiffInString())
		assert.Empty(t, Now().DiffInString(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		now := Now()
		assert.Equal(t, "just now", Parse("2020-08-05 13:14:15").DiffInString(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, "-1 year", now.Copy().AddYearsNoOverflow(1).DiffInString(now))
		assert.Equal(t, "1 year", now.Copy().SubYearsNoOverflow(1).DiffInString(now))
		assert.Equal(t, "-1 month", now.Copy().AddMonthsNoOverflow(1).DiffInString(now))
		assert.Equal(t, "1 month", now.Copy().SubMonthsNoOverflow(1).DiffInString(now))
		assert.Equal(t, "-1 week", now.Copy().AddWeeks(1).DiffInString(now))
		assert.Equal(t, "1 week", now.Copy().SubWeeks(1).DiffInString(now))
		assert.Equal(t, "-1 day", now.Copy().AddDays(1).DiffInString(now))
		assert.Equal(t, "1 day", now.Copy().SubDays(1).DiffInString(now))
		assert.Equal(t, "-1 hour", now.Copy().AddHours(1).DiffInString(now))
		assert.Equal(t, "1 hour", now.Copy().SubHours(1).DiffInString(now))
		assert.Equal(t, "-1 minute", now.Copy().AddMinutes(1).DiffInString(now))
		assert.Equal(t, "1 minute", now.Copy().SubMinutes(1).DiffInString(now))
		assert.Equal(t, "-1 second", now.Copy().AddSeconds(1).DiffInString(now))
		assert.Equal(t, "1 second", now.Copy().SubSeconds(1).DiffInString(now))
	})
}

func TestCarbon_DiffAbsInString(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.DiffAbsInString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").DiffAbsInString())
		assert.Empty(t, Now().DiffAbsInString(Parse("")))
		assert.Empty(t, Parse("xxx").DiffAbsInString())
		assert.Empty(t, Now().DiffAbsInString(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		now := Now()
		assert.Equal(t, "just now", Parse("2020-08-05 13:14:15").DiffAbsInString(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, "1 year", now.Copy().AddYearsNoOverflow(1).DiffAbsInString(now))
		assert.Equal(t, "1 year", now.Copy().SubYearsNoOverflow(1).DiffAbsInString(now))
		assert.Equal(t, "1 month", now.Copy().AddMonthsNoOverflow(1).DiffAbsInString(now))
		assert.Equal(t, "1 month", now.Copy().SubMonthsNoOverflow(1).DiffAbsInString(now))
		assert.Equal(t, "1 day", now.Copy().AddDays(1).DiffAbsInString(now))
		assert.Equal(t, "1 day", now.Copy().SubDays(1).DiffAbsInString(now))
		assert.Equal(t, "1 hour", now.Copy().AddHours(1).DiffAbsInString(now))
		assert.Equal(t, "1 hour", now.Copy().SubHours(1).DiffAbsInString(now))
		assert.Equal(t, "1 minute", now.Copy().AddMinutes(1).DiffAbsInString(now))
		assert.Equal(t, "1 minute", now.Copy().SubMinutes(1).DiffAbsInString(now))
		assert.Equal(t, "1 second", now.Copy().AddSeconds(1).DiffAbsInString(now))
		assert.Equal(t, "1 second", now.Copy().SubSeconds(1).DiffAbsInString(now))
	})
}

func TestCarbon_DiffInDuration(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.DiffInDuration())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").DiffInDuration())
		assert.Empty(t, Now().DiffInDuration(Parse("")))
		assert.Empty(t, Parse("xxx").DiffInDuration())
		assert.Empty(t, Now().DiffInDuration(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		now := Now()
		assert.Equal(t, "0s", Parse("2020-08-05 13:14:15").DiffInDuration(Parse("2020-08-05 13:14:15")).String())
		assert.Equal(t, "-8760h0m0s", now.Copy().AddYearsNoOverflow(1).DiffInDuration(now).String())
		assert.Equal(t, "8760h0m0s", now.Copy().SubYearsNoOverflow(1).DiffInDuration(now).String())
		assert.Equal(t, "-744h0m0s", now.Copy().AddMonthsNoOverflow(1).DiffInDuration(now).String())
		assert.Equal(t, "672h0m0s", now.Copy().SubMonthsNoOverflow(1).DiffInDuration(now).String())
		assert.Equal(t, "-24h0m0s", now.Copy().AddDays(1).DiffInDuration(now).String())
		assert.Equal(t, "24h0m0s", now.Copy().SubDays(1).DiffInDuration(now).String())
	})
}

func TestCarbon_DiffAbsInDuration(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.DiffAbsInDuration())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").DiffAbsInDuration())
		assert.Empty(t, Now().DiffAbsInDuration(Parse("")))
		assert.Empty(t, Parse("xxx").DiffAbsInDuration())
		assert.Empty(t, Now().DiffAbsInDuration(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		now := Now()
		assert.Equal(t, "0s", Parse("2020-08-05 13:14:15").DiffAbsInDuration(Parse("2020-08-05 13:14:15")).String())
		assert.Equal(t, "8760h0m0s", now.Copy().AddYearsNoOverflow(1).DiffAbsInDuration(now).String())
		assert.Equal(t, "8760h0m0s", now.Copy().SubYearsNoOverflow(1).DiffAbsInDuration(now).String())
		assert.Equal(t, "744h0m0s", now.Copy().AddMonthsNoOverflow(1).DiffAbsInDuration(now).String())
		assert.Equal(t, "672h0m0s", now.Copy().SubMonthsNoOverflow(1).DiffAbsInDuration(now).String())
		assert.Equal(t, "24h0m0s", now.Copy().AddDays(1).DiffAbsInDuration(now).String())
		assert.Equal(t, "24h0m0s", now.Copy().SubDays(1).DiffAbsInDuration(now).String())
	})
}

func TestCarbon_DiffForHumans(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.DiffForHumans())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").DiffForHumans())
		assert.Empty(t, Now().DiffForHumans(Parse("")))
		assert.Empty(t, Parse("xxx").DiffForHumans())
		assert.Empty(t, Now().DiffForHumans(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		SetTestNow(Parse("2020-08-05 13:14:15"))
		now := Now()
		assert.Equal(t, "just now", Parse("2020-08-05 13:14:15").DiffForHumans(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, "2 days ago", Parse("2020-08-03 13:14:15").DiffForHumans())
		assert.Equal(t, "2 days from now", Parse("2020-08-07 13:14:15").DiffForHumans())
		assert.Equal(t, "1 year after", now.Copy().AddYearsNoOverflow(1).DiffForHumans(now))
		assert.Equal(t, "1 year before", now.Copy().SubYearsNoOverflow(1).DiffForHumans(now))
		assert.Equal(t, "1 month after", now.Copy().AddMonthsNoOverflow(1).DiffForHumans(now))
		assert.Equal(t, "1 month before", now.Copy().SubMonthsNoOverflow(1).DiffForHumans(now))
		assert.Equal(t, "1 day after", now.Copy().AddDays(1).DiffForHumans(now))
		assert.Equal(t, "1 day before", now.Copy().SubDays(1).DiffForHumans(now))
	})
}

func TestCarbon_getDiffInMonths(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Zero(t, getDiffInMonths(c, c))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, getDiffInMonths(Parse(""), Parse("")))
		assert.Zero(t, getDiffInMonths(Parse(""), Parse("xxx")))
		assert.Zero(t, getDiffInMonths(Parse("xxx"), Parse("xxx")))
		assert.Zero(t, getDiffInMonths(Parse("xxx"), Parse("")))
	})

	t.Run("valid time", func(t *testing.T) {
		SetTestNow(Parse("2020-08-05 13:14:15"))
		assert.Equal(t, int64(0), getDiffInMonths(Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(1), getDiffInMonths(Parse("2020-07-05 13:14:15"), Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(-1), getDiffInMonths(Parse("2020-09-05 13:14:15"), Parse("2020-08-05 13:14:15")))
	})
}

// https://github.com/dromara/carbon/issues/255
func TestCarbon_Issue255(t *testing.T) {
	assert.Equal(t, int64(0), Parse("2024-10-11").DiffInMonths(Parse("2024-11-10")))
	assert.Equal(t, int64(0), Parse("2024-11-10").DiffInMonths(Parse("2024-10-11")))
	assert.Equal(t, int64(1), Parse("2024-10-11").DiffInMonths(Parse("2024-11-11")))
	assert.Equal(t, int64(-1), Parse("2024-11-11").DiffInMonths(Parse("2024-10-11")))
	assert.Equal(t, int64(0), Parse("2024-10-11 23:59:00").DiffInMonths(Parse("2024-11-11 00:00:00")))
	assert.Equal(t, int64(-23), Parse("2020-08-05 13:14:15").DiffInMonths(Parse("2018-08-28 13:14:59")))
	assert.Equal(t, int64(23), Parse("2018-08-28 13:14:59").DiffInMonths(Parse("2020-08-05 13:14:15")))
	assert.Equal(t, int64(11999), Parse("1024-12-25 13:14:20").DiffInMonths(Parse("2024-12-25 13:14:19")))
}
