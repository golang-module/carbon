package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var japanese = "jp"

func TestLang_Japanese_Constellation(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-05", "やぎざ"},
		{"2020-02-05", "みずがめざ"},
		{"2020-03-05", "うおざ"},
		{"2020-04-05", "おひつじざ"},
		{"2020-05-05", "おうしざ"},
		{"2020-06-05", "ふたござ"},
		{"2020-07-05", "かにざ"},
		{"2020-08-05", "ししざ"},
		{"2020-09-05", "おとめざ"},
		{"2020-10-05", "てんびんざ"},
		{"2020-11-05", "さそりざ"},
		{"2020-12-05", "いてざ"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(japanese)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Constellation(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Japanese_Season(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-05", "ふゆ"},
		{"2020-02-05", "ふゆ"},
		{"2020-03-05", "はる"},
		{"2020-04-05", "はる"},
		{"2020-05-05", "はる"},
		{"2020-06-05", "なつ"},
		{"2020-07-05", "なつ"},
		{"2020-08-05", "なつ"},
		{"2020-09-05", "あき"},
		{"2020-10-05", "あき"},
		{"2020-11-05", "あき"},
		{"2020-12-05", "ふゆ"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(japanese)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Season(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Japanese_ToMonthString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-05", "いちがつ"},
		{"2020-02-05", "にがつ"},
		{"2020-03-05", "さんがつ"},
		{"2020-04-05", "しがつ"},
		{"2020-05-05", "ごがつ"},
		{"2020-06-05", "ろくがつ"},
		{"2020-07-05", "しちがつ"},
		{"2020-08-05", "はちがつ"},
		{"2020-09-05", "くがつ"},
		{"2020-10-05", "じゅうがつ"},
		{"2020-11-05", "じゅういちがつ"},
		{"2020-12-05", "じゅうにがつ"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(japanese)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToMonthString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Japanese_ToShortMonthString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-05", "1がつ"},
		{"2020-02-05", "2がつ"},
		{"2020-03-05", "3がつ"},
		{"2020-04-05", "4がつ"},
		{"2020-05-05", "5がつ"},
		{"2020-06-05", "6がつ"},
		{"2020-07-05", "7がつ"},
		{"2020-08-05", "8がつ"},
		{"2020-09-05", "9がつ"},
		{"2020-10-05", "10がつ"},
		{"2020-11-05", "11がつ"},
		{"2020-12-05", "12がつ"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(japanese)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortMonthString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Japanese_ToWeekString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-01", "にちようび"},
		{"2020-08-02", "げつようび"},
		{"2020-08-03", "かようび"},
		{"2020-08-04", "すいようび"},
		{"2020-08-05", "もくようび"},
		{"2020-08-06", "きんようび"},
		{"2020-08-07", "どようび"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(japanese)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToWeekString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Japanese_ToShortWeekString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-01", "週六"},
		{"2020-08-02", "週日"},
		{"2020-08-03", "週一"},
		{"2020-08-04", "週二"},
		{"2020-08-05", "週三"},
		{"2020-08-06", "週四"},
		{"2020-08-07", "週五"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(japanese)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortWeekString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Japanese_DiffForHumans(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   string // 输入值
		input2   string // 输入值
		expected string // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", "たった今"},
		{"2020-08-05 13:14:15", "2021-08-05 13:14:15", "1 年前"},
		{"2020-08-05 13:14:15", "2019-08-05 13:14:15", "1 年後"},
		{"2020-08-05 13:14:15", "2030-08-05 13:14:15", "10 年前"},
		{"2020-08-05 13:14:15", "2010-08-05 13:14:15", "10 年後"},

		{"2020-08-05 13:14:15", "2020-09-05 13:14:15", "1 ヶ月前"},
		{"2020-08-05 13:14:15", "2020-07-05 13:14:15", "1 ヶ月後"},
		{"2020-08-05 13:14:15", "2021-06-05 13:14:15", "10 ヶ月前"},
		{"2020-08-05 13:14:15", "2019-10-05 13:14:15", "10 ヶ月後"},

		{"2020-08-05 13:14:15", "2020-08-06 13:14:15", "1 日前"},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", "1 日後"},
		{"2020-08-05 13:14:15", "2020-08-15 13:14:15", "1 週間前"},
		{"2020-08-05 13:14:15", "2020-07-26 13:14:15", "1 週間後"},

		{"2020-08-05 13:14:15", "2020-08-05 14:14:15", "1 時間前"},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:15", "1 時間後"},
		{"2020-08-05 13:14:15", "2020-08-05 23:14:15", "10 時間前"},
		{"2020-08-05 13:14:15", "2020-08-05 03:14:15", "10 時間後"},

		{"2020-08-05 13:14:15", "2020-08-05 13:15:15", "1 分前"},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:15", "1 分後"},
		{"2020-08-05 13:14:15", "2020-08-05 13:24:15", "10 分前"},
		{"2020-08-05 13:14:15", "2020-08-05 13:04:15", "10 分後"},

		{"2020-08-05 13:14:15", "2020-08-05 13:14:16", "1 秒前"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:14", "1 秒後"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:25", "10 秒前"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:05", "10 秒後"},
	}

	for index, test := range tests {
		c1 := Parse(test.input1)
		c2 := Parse(test.input2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.SetLocale(japanese).DiffForHumans(c2), "Current test index is "+strconv.Itoa(index))
	}
}
