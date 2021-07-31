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
		id       int    // 测试id
		input    string // 输入值
		expected string // 期望值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},
		{5, "0000-00-00 00:00:00", ""},

		{6, "2020-01-05", "やぎざ"},
		{7, "2020-02-05", "みずがめざ"},
		{8, "2020-03-05", "うおざ"},
		{9, "2020-04-05", "おひつじざ"},
		{10, "2020-05-05", "おうしざ"},
		{11, "2020-06-05", "ふたござ"},
		{12, "2020-07-05", "かにざ"},
		{13, "2020-08-05", "ししざ"},
		{14, "2020-09-05", "おとめざ"},
		{15, "2020-10-05", "てんびんざ"},
		{16, "2020-11-05", "さそりざ"},
		{17, "2020-12-05", "いてざ"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(japanese)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Constellation(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_Japanese_Season(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected string // 期望值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},
		{5, "0000-00-00 00:00:00", ""},

		{6, "2020-01-05", "ふゆ"},
		{7, "2020-02-05", "ふゆ"},
		{8, "2020-03-05", "はる"},
		{9, "2020-04-05", "はる"},
		{10, "2020-05-05", "はる"},
		{11, "2020-06-05", "なつ"},
		{12, "2020-07-05", "なつ"},
		{13, "2020-08-05", "なつ"},
		{14, "2020-09-05", "あき"},
		{15, "2020-10-05", "あき"},
		{16, "2020-11-05", "あき"},
		{17, "2020-12-05", "ふゆ"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(japanese)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Season(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_Japanese_ToMonthString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected string // 期望值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},
		{5, "0000-00-00 00:00:00", ""},

		{6, "2020-01-05", "いちがつ"},
		{7, "2020-02-05", "にがつ"},
		{8, "2020-03-05", "さんがつ"},
		{9, "2020-04-05", "しがつ"},
		{10, "2020-05-05", "ごがつ"},
		{11, "2020-06-05", "ろくがつ"},
		{12, "2020-07-05", "しちがつ"},
		{13, "2020-08-05", "はちがつ"},
		{14, "2020-09-05", "くがつ"},
		{15, "2020-10-05", "じゅうがつ"},
		{16, "2020-11-05", "じゅういちがつ"},
		{17, "2020-12-05", "じゅうにがつ"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(japanese)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToMonthString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_Japanese_ToShortMonthString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected string // 期望值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},
		{5, "0000-00-00 00:00:00", ""},

		{6, "2020-01-05", "1がつ"},
		{7, "2020-02-05", "2がつ"},
		{8, "2020-03-05", "3がつ"},
		{9, "2020-04-05", "4がつ"},
		{10, "2020-05-05", "5がつ"},
		{11, "2020-06-05", "6がつ"},
		{12, "2020-07-05", "7がつ"},
		{13, "2020-08-05", "8がつ"},
		{14, "2020-09-05", "9がつ"},
		{15, "2020-10-05", "10がつ"},
		{16, "2020-11-05", "11がつ"},
		{17, "2020-12-05", "12がつ"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(japanese)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortMonthString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_Japanese_ToWeekString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected string // 期望值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},
		{5, "0000-00-00 00:00:00", ""},

		{6, "2020-08-01", "にちようび"},
		{7, "2020-08-02", "げつようび"},
		{8, "2020-08-03", "かようび"},
		{9, "2020-08-04", "すいようび"},
		{10, "2020-08-05", "もくようび"},
		{11, "2020-08-06", "きんようび"},
		{12, "2020-08-07", "どようび"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(japanese)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToWeekString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_Japanese_ToShortWeekString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected string // 期望值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},
		{5, "0000-00-00 00:00:00", ""},

		{6, "2020-08-01", "週六"},
		{7, "2020-08-02", "週日"},
		{8, "2020-08-03", "週一"},
		{9, "2020-08-04", "週二"},
		{10, "2020-08-05", "週三"},
		{11, "2020-08-06", "週四"},
		{12, "2020-08-07", "週五"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(japanese)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortWeekString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_Japanese_DiffForHumans(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input1   string // 输入值
		input2   string // 输入值
		expected string // 期望值
	}{
		{1, "2020-08-05 13:14:15", "2020-08-05 13:14:15", "たった今"},
		{2, "2020-08-05 13:14:15", "2021-08-05 13:14:15", "1 ねん前"},
		{3, "2020-08-05 13:14:15", "2019-08-05 13:14:15", "1 ねん後"},
		{4, "2020-08-05 13:14:15", "2030-08-05 13:14:15", "10 ねん前"},
		{5, "2020-08-05 13:14:15", "2010-08-05 13:14:15", "10 ねん後"},

		{6, "2020-08-05 13:14:15", "2020-09-05 13:14:15", "1 ヶがつ前"},
		{7, "2020-08-05 13:14:15", "2020-07-05 13:14:15", "1 ヶがつ後"},
		{8, "2020-08-05 13:14:15", "2021-06-05 13:14:15", "10 ヶがつ前"},
		{9, "2020-08-05 13:14:15", "2019-10-05 13:14:15", "10 ヶがつ後"},

		{10, "2020-08-05 13:14:15", "2020-08-06 13:14:15", "1 日前"},
		{11, "2020-08-05 13:14:15", "2020-08-04 13:14:15", "1 日後"},
		{12, "2020-08-05 13:14:15", "2020-08-15 13:14:15", "1 週間前"},
		{13, "2020-08-05 13:14:15", "2020-07-26 13:14:15", "1 週間後"},

		{14, "2020-08-05 13:14:15", "2020-08-05 14:14:15", "1 時間前"},
		{15, "2020-08-05 13:14:15", "2020-08-05 12:14:15", "1 時間後"},
		{16, "2020-08-05 13:14:15", "2020-08-05 23:14:15", "10 時間前"},
		{17, "2020-08-05 13:14:15", "2020-08-05 03:14:15", "10 時間後"},

		{18, "2020-08-05 13:14:15", "2020-08-05 13:15:15", "1 分钟前"},
		{19, "2020-08-05 13:14:15", "2020-08-05 13:13:15", "1 分钟後"},
		{20, "2020-08-05 13:14:15", "2020-08-05 13:24:15", "10 分钟前"},
		{21, "2020-08-05 13:14:15", "2020-08-05 13:04:15", "10 分钟後"},

		{22, "2020-08-05 13:14:15", "2020-08-05 13:14:16", "1 秒前"},
		{23, "2020-08-05 13:14:15", "2020-08-05 13:14:14", "1 秒後"},
		{24, "2020-08-05 13:14:15", "2020-08-05 13:14:25", "10 秒前"},
		{25, "2020-08-05 13:14:15", "2020-08-05 13:14:05", "10 秒後"},
	}

	for _, test := range tests {
		c1 := Parse(test.input1)
		c2 := Parse(test.input2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.SetLocale(japanese).DiffForHumans(c2), "Current test id is "+strconv.Itoa(test.id))
	}
}
