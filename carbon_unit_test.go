package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewCarbon(t *testing.T) {
	t.Run("without time", func(t *testing.T) {
		assert.Equal(t, time.Time{}.Unix(), NewCarbon().Timestamp())
	})

	t.Run("without timezone", func(t *testing.T) {
		now := time.Now()
		assert.Equal(t, now.Unix(), NewCarbon(now).Timestamp())
	})

	t.Run("with timezone", func(t *testing.T) {
		now := time.Now().In(time.Local)
		assert.Equal(t, now.Unix(), NewCarbon(now).Timestamp())
	})
}

func TestCarbon_Copy(t *testing.T) {
	t.Run("copy time", func(t *testing.T) {
		oldCarbon := Parse("2020-08-05")
		newCarbon := oldCarbon.Copy()

		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", oldCarbon.ToString())
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", newCarbon.ToString())

		oldCarbon = oldCarbon.AddDay()
		assert.Equal(t, "2020-08-06 00:00:00 +0000 UTC", oldCarbon.ToString())
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", newCarbon.ToString())
	})

	t.Run("copy timezone", func(t *testing.T) {
		oldCarbon := Parse("2020-08-05")
		newCarbon := oldCarbon.Copy()

		assert.Equal(t, UTC, oldCarbon.Location())
		assert.Equal(t, UTC, newCarbon.Location())

		oldCarbon = oldCarbon.SetTimezone(PRC)
		assert.Equal(t, PRC, oldCarbon.Location())
		assert.Equal(t, UTC, newCarbon.Location())

		newCarbon = newCarbon.SetTimezone(Japan)
		assert.Equal(t, PRC, oldCarbon.Location())
		assert.Equal(t, Japan, newCarbon.Location())
	})

	t.Run("copy layout", func(t *testing.T) {
		oldCarbon := Parse("2020-08-05")
		newCarbon := oldCarbon.Copy()

		assert.Equal(t, "2006-01-02", oldCarbon.CurrentLayout())
		assert.Equal(t, "2006-01-02", newCarbon.CurrentLayout())

		oldCarbon = oldCarbon.SetLayout(DateTimeLayout)
		assert.Equal(t, DateTimeLayout, oldCarbon.CurrentLayout())
		assert.Equal(t, DateLayout, newCarbon.CurrentLayout())

		newCarbon = newCarbon.SetLayout(RFC1036Layout)
		assert.Equal(t, DateTimeLayout, oldCarbon.CurrentLayout())
		assert.Equal(t, RFC1036Layout, newCarbon.CurrentLayout())
	})

	t.Run("copy weekStartsAt", func(t *testing.T) {
		oldCarbon := Parse("2020-08-05")
		newCarbon := oldCarbon.Copy()

		assert.Equal(t, DefaultWeekStartsAt, oldCarbon.WeekStartsAt())
		assert.Equal(t, DefaultWeekStartsAt, newCarbon.WeekStartsAt())

		oldCarbon = oldCarbon.SetWeekStartsAt(Monday)
		assert.Equal(t, Monday, oldCarbon.WeekStartsAt())
		assert.Equal(t, DefaultWeekStartsAt, newCarbon.WeekStartsAt())

		newCarbon = newCarbon.SetWeekStartsAt(Sunday)
		assert.Equal(t, Monday, oldCarbon.WeekStartsAt())
		assert.Equal(t, Sunday, newCarbon.WeekStartsAt())
	})

	t.Run("copy lang", func(t *testing.T) {
		oldCarbon := Parse("2020-08-05")
		newCarbon := oldCarbon.Copy()

		assert.Equal(t, August, oldCarbon.ToMonthString())
		assert.Equal(t, August, newCarbon.ToMonthString())

		oldCarbon.SetLocale("zh-CN")
		assert.Equal(t, "八月", oldCarbon.ToMonthString())
		assert.Equal(t, "August", newCarbon.ToMonthString())

		newCarbon.SetLocale("jp")
		assert.Equal(t, "八月", oldCarbon.ToMonthString())
		assert.Equal(t, "8月", newCarbon.ToMonthString())
	})

	t.Run("copy error", func(t *testing.T) {
		oldCarbon := Parse("xxx")
		newCarbon := oldCarbon.Copy()

		assert.True(t, oldCarbon.HasError())
		assert.True(t, newCarbon.HasError())

		newCarbon = newCarbon.SetLayout("xxx")
		assert.True(t, oldCarbon.HasError())
		assert.True(t, newCarbon.HasError())
	})
}
