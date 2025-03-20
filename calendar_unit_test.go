package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_Lunar(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Empty(t, NewCarbon().Lunar().String())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.Lunar().String())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").Lunar().String())
		assert.Empty(t, Parse("0").Lunar().String())
		assert.Empty(t, Parse("xxx").Lunar().String())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2023-12-08 00:00:00", Parse("2024-01-18", PRC).Lunar().String())
		assert.Equal(t, "2023-12-11 00:00:00", Parse("2024-01-21", PRC).Lunar().String())
		assert.Equal(t, "2023-12-14 00:00:00", Parse("2024-01-24", PRC).Lunar().String())
	})
}

func TestCreateFromLunar(t *testing.T) {
	t.Run("invalid lunar", func(t *testing.T) {
		assert.Empty(t, CreateFromLunar(2200, 12, 14, 12, 0, 0, false).ToDateTimeString())
	})

	t.Run("valid lunar", func(t *testing.T) {
		assert.Equal(t, "2024-01-21 00:00:00", CreateFromLunar(2023, 12, 11, 0, 0, 0, false).ToDateTimeString(PRC))
		assert.Equal(t, "2024-01-18 00:00:00", CreateFromLunar(2023, 12, 8, 0, 0, 0, false).ToDateTimeString(PRC))
		assert.Equal(t, "2024-01-24 12:00:00", CreateFromLunar(2023, 12, 14, 12, 0, 0, false).ToDateTimeString(PRC))
	})
}

func TestCarbon_Julian(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Zero(t, NewCarbon().Julian().JD())
		assert.Zero(t, NewCarbon().Julian().MJD())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Zero(t, c.Julian().JD())
		assert.Zero(t, c.Julian().MJD())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").Julian().JD())
		assert.Zero(t, Parse("0").Julian().JD())
		assert.Zero(t, Parse("xxx").Julian().JD())

		assert.Zero(t, Parse("").Julian().MJD())
		assert.Zero(t, Parse("0").Julian().MJD())
		assert.Zero(t, Parse("xxx").Julian().MJD())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, float64(2460334), Parse("2024-01-24 12:00:00").Julian().JD())
		assert.Equal(t, 60333.5, Parse("2024-01-24 12:00:00").Julian().MJD())
	})
}

func TestCreateFromJulian(t *testing.T) {
	t.Run("invalid julian", func(t *testing.T) {
		assert.Empty(t, CreateFromJulian(0).ToDateTimeString())
		assert.Empty(t, CreateFromJulian(-1).ToDateTimeString())
	})

	t.Run("valid julian", func(t *testing.T) {
		assert.Equal(t, "2024-01-24 12:00:00", CreateFromJulian(2460334).ToDateTimeString())
		assert.Equal(t, "2024-01-24 12:00:00", CreateFromJulian(60333.5).ToDateTimeString())
	})
}

func TestCarbon_Persian(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Empty(t, NewCarbon().Persian().String())
		assert.Empty(t, NewCarbon().Persian().String())
	})

	t.Run("nil time", func(t *testing.T) {
		c := NewCarbon()
		c = nil
		assert.Empty(t, c.Persian().String())
		assert.Empty(t, c.Persian().String())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").Persian().String())
		assert.Empty(t, Parse("0").Persian().String())
		assert.Empty(t, Parse("xxx").Persian().String())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "1178-10-11 00:00:00", Parse("1800-01-01 00:00:00").Persian().String())
		assert.Equal(t, "1399-05-15 13:14:15", Parse("2020-08-05 13:14:15").Persian().String())
		assert.Equal(t, "1402-10-11 00:00:00", Parse("2024-01-01 00:00:00").Persian().String())
	})
}

func TestCreateFromPersian(t *testing.T) {
	t.Run("invalid persian", func(t *testing.T) {
		assert.Empty(t, CreateFromPersian(9999, 12, 14, 12, 0, 0).ToDateTimeString())
	})

	t.Run("valid persian", func(t *testing.T) {
		assert.Equal(t, "1800-01-01 00:00:00", CreateFromPersian(1178, 10, 11, 0, 0, 0).ToDateTimeString())
		assert.Equal(t, "2024-01-01 00:00:00", CreateFromPersian(1402, 10, 11, 0, 0, 0).ToDateTimeString())
		assert.Equal(t, "2024-08-05 12:00:00", CreateFromPersian(1403, 5, 15, 12, 0, 0).ToDateTimeString())
	})
}
