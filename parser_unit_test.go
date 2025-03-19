package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		assert.Nil(t, Parse(""))
		assert.Empty(t, Parse("").ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Error(t, Parse("0").Error)
		assert.Error(t, Parse("xxx").Error)
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, Now().Timestamp(), Parse("now").Timestamp())
		assert.Equal(t, Yesterday().Timestamp(), Parse("yesterday").Timestamp())
		assert.Equal(t, Tomorrow().Timestamp(), Parse("tomorrow").Timestamp())
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", Parse("2020-8-5").ToString())
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", Parse("2020-8-05").ToString())
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", Parse("2020-08-05").ToString())
		assert.Equal(t, "2020-08-05 01:02:03 +0000 UTC", Parse("2020-8-5 1:2:3").ToString())
		assert.Equal(t, "2020-08-05 01:02:03 +0000 UTC", Parse("2020-08-05 1:2:03").ToString())
		assert.Equal(t, "2020-08-05 01:02:03 +0000 UTC", Parse("2020-08-05 1:02:03").ToString())
		assert.Equal(t, "2020-08-05 01:02:03 +0000 UTC", Parse("2020-08-05 01:02:03").ToString())
	})

	// https://github.com/dromara/carbon/issues/202
	t.Run("issue202", func(t *testing.T) {
		assert.Equal(t, "2023-01-08 09:02:48 +0000 UTC", Parse("2023-01-08T09:02:48").ToString())
		assert.Equal(t, "2023-01-08 09:02:48 +0000 UTC", Parse("2023-1-8T09:02:48").ToString())
		assert.Equal(t, "2023-01-08 09:02:48 +0000 UTC", Parse("2023-01-08T9:2:48").ToString())
		assert.Equal(t, "2023-01-08 09:02:48 +0000 UTC", Parse("2023-01-8T9:2:48").ToString())
	})

	// https://github.com/dromara/carbon/issues/232
	t.Run("issue232", func(t *testing.T) {
		assert.Equal(t, "0000-01-01 00:00:00 +0000 UTC", Parse("0000-01-01 00:00:00").ToString())
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", Parse("0001-01-01 00:00:00").ToString())
		assert.Equal(t, "", Parse("0001-00-00 00:00:00").ToString())
	})
}

func TestParseByFormat(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		assert.Nil(t, ParseByFormat("", DateFormat))
		assert.Empty(t, ParseByFormat("", DateFormat).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Error(t, ParseByFormat("0", DateFormat).Error)
		assert.Error(t, ParseByFormat("xxx", DateFormat).Error)
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", ParseByFormat("2020-08-05", DateFormat).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByFormat("2020-08-05 13:14:15", DateTimeFormat).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByFormat("2020|08|05 13:14:15", "Y|m|d H:i:s").ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒").ToString())
	})

	// https://github.com/dromara/carbon/issues/206
	t.Run("issue206", func(t *testing.T) {
		assert.Equal(t, "2023-11-11 04:34:00 +0000 UTC", ParseByFormat("1699677240", "U").ToString())
		assert.Equal(t, "2023-11-11 04:34:00.666 +0000 UTC", ParseByFormat("1699677240666", "V").ToString())
		assert.Equal(t, "2023-11-11 04:34:00.666666 +0000 UTC", ParseByFormat("1699677240666666", "X").ToString())
		assert.Equal(t, "2023-11-11 04:34:00.666666666 +0000 UTC", ParseByFormat("1699677240666666666", "Z").ToString())
	})
}

func TestParseByLayout(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		assert.Nil(t, ParseByLayout("", DateLayout))
		assert.Empty(t, ParseByLayout("", DateLayout).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Error(t, ParseByLayout("0", DateLayout).Error)
		assert.Error(t, ParseByLayout("xxx", DateLayout).Error)
	})

	t.Run("empty layout", func(t *testing.T) {
		assert.True(t, ParseByLayout("2020-08-05", "").HasError())
		assert.Empty(t, ParseByLayout("2020-08-05", "").ToString())

		assert.False(t, ParseByLayout("2020-08-05 13:14:15", "").HasError())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByLayout("2020-08-05 13:14:15", "").ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", ParseByLayout("2020-08-05", DateLayout).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByLayout("2020-08-05 13:14:15", DateTimeLayout, PRC).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByLayout("2020|08|05 13:14:15", "2006|01|02 15:04:05").ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒").ToString())
	})
}
