package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var korean = "kr"

func TestLang_Korean_Constellation(t *testing.T) {
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

		{"2020-01-05", "염소자리"},
		{"2020-02-05", "물병자리"},
		{"2020-03-05", "물고기자리"},
		{"2020-04-05", "양자리"},
		{"2020-05-05", "황소자리"},
		{"2020-06-05", "쌍둥이자리"},
		{"2020-07-05", "게자리"},
		{"2020-08-05", "사자자리"},
		{"2020-09-05", "처녀자리"},
		{"2020-10-05", "천칭자리"},
		{"2020-11-05", "전갈자리"},
		{"2020-12-05", "사수자리"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(korean)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Constellation(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Korean_Season(t *testing.T) {
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

		{"2020-01-05", "겨울"},
		{"2020-02-05", "겨울"},
		{"2020-03-05", "봄"},
		{"2020-04-05", "봄"},
		{"2020-05-05", "봄"},
		{"2020-06-05", "여름"},
		{"2020-07-05", "여름"},
		{"2020-08-05", "여름"},
		{"2020-09-05", "가을"},
		{"2020-10-05", "가을"},
		{"2020-11-05", "가을"},
		{"2020-12-05", "겨울"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(korean)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Season(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Korean_ToMonthString(t *testing.T) {
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

		{"2020-01-05", "일월"},
		{"2020-02-05", "이월"},
		{"2020-03-05", "삼월"},
		{"2020-04-05", "사월"},
		{"2020-05-05", "오월"},
		{"2020-06-05", "유월"},
		{"2020-07-05", "칠월"},
		{"2020-08-05", "팔월"},
		{"2020-09-05", "구월"},
		{"2020-10-05", "시월"},
		{"2020-11-05", "십일월"},
		{"2020-12-05", "십이월"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(korean)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToMonthString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Korean_ToShortMonthString(t *testing.T) {
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

		{"2020-01-05", "1월"},
		{"2020-02-05", "2월"},
		{"2020-03-05", "3월"},
		{"2020-04-05", "4월"},
		{"2020-05-05", "5월"},
		{"2020-06-05", "6월"},
		{"2020-07-05", "7월"},
		{"2020-08-05", "8월"},
		{"2020-09-05", "9월"},
		{"2020-10-05", "10월"},
		{"2020-11-05", "11월"},
		{"2020-12-05", "12월"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(korean)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortMonthString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Korean_ToWeekString(t *testing.T) {
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

		{"2020-08-01", "토요일"},
		{"2020-08-02", "일요일"},
		{"2020-08-03", "월요일"},
		{"2020-08-04", "화요일"},
		{"2020-08-05", "수요일"},
		{"2020-08-06", "목요일"},
		{"2020-08-07", "금요일"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(korean)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToWeekString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Korean_ToShortWeekString(t *testing.T) {
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

		{"2020-08-01", "토요일"},
		{"2020-08-02", "일요일"},
		{"2020-08-03", "월요일"},
		{"2020-08-04", "화요일"},
		{"2020-08-05", "수요일"},
		{"2020-08-06", "목요일"},
		{"2020-08-07", "금요일"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(korean)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortWeekString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Korean_DiffForHumans(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   string // 输入值
		input2   string // 输入值
		expected string // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", "방금"},
		{"2020-08-05 13:14:15", "2021-08-05 13:14:15", "1 년전"},
		{"2020-08-05 13:14:15", "2019-08-05 13:14:15", "1 년후"},
		{"2020-08-05 13:14:15", "2030-08-05 13:14:15", "10 년전"},
		{"2020-08-05 13:14:15", "2010-08-05 13:14:15", "10 년후"},

		{"2020-08-05 13:14:15", "2020-09-05 13:14:15", "1 개월전"},
		{"2020-08-05 13:14:15", "2020-07-05 13:14:15", "1 개월후"},
		{"2020-08-05 13:14:15", "2021-06-05 13:14:15", "10 개월전"},
		{"2020-08-05 13:14:15", "2019-10-05 13:14:15", "10 개월후"},

		{"2020-08-05 13:14:15", "2020-08-06 13:14:15", "1 일전"},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", "1 일후"},
		{"2020-08-05 13:14:15", "2020-08-15 13:14:15", "1 주전"},
		{"2020-08-05 13:14:15", "2020-07-26 13:14:15", "1 주후"},

		{"2020-08-05 13:14:15", "2020-08-05 14:14:15", "1 시간전"},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:15", "1 시간후"},
		{"2020-08-05 13:14:15", "2020-08-05 23:14:15", "10 시간전"},
		{"2020-08-05 13:14:15", "2020-08-05 03:14:15", "10 시간후"},

		{"2020-08-05 13:14:15", "2020-08-05 13:15:15", "1 분전"},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:15", "1 분후"},
		{"2020-08-05 13:14:15", "2020-08-05 13:24:15", "10 분전"},
		{"2020-08-05 13:14:15", "2020-08-05 13:04:15", "10 분후"},

		{"2020-08-05 13:14:15", "2020-08-05 13:14:16", "1 초전"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:14", "1 초후"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:25", "10 초전"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:05", "10 초후"},
	}

	for index, test := range tests {
		c1 := Parse(test.input1)
		c2 := Parse(test.input2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.SetLocale(korean).DiffForHumans(c2), "Current test index is "+strconv.Itoa(index))
	}
}
