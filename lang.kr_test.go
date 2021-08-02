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
		id       int    // 测试id
		input    string // 输入值
		expected string // 期望值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},
		{5, "0000-00-00 00:00:00", ""},

		{6, "2020-01-05", "염소자리"},
		{7, "2020-02-05", "물병자리"},
		{8, "2020-03-05", "물고기자리"},
		{9, "2020-04-05", "양자리"},
		{10, "2020-05-05", "황소자리"},
		{11, "2020-06-05", "쌍둥이자리"},
		{12, "2020-07-05", "게자리"},
		{13, "2020-08-05", "사자자리"},
		{14, "2020-09-05", "처녀자리"},
		{15, "2020-10-05", "천칭자리"},
		{16, "2020-11-05", "전갈자리"},
		{17, "2020-12-05", "사수자리"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(korean)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Constellation(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_Korean_Season(t *testing.T) {
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

		{6, "2020-01-05", "겨울"},
		{7, "2020-02-05", "겨울"},
		{8, "2020-03-05", "봄"},
		{9, "2020-04-05", "봄"},
		{10, "2020-05-05", "봄"},
		{11, "2020-06-05", "여름"},
		{12, "2020-07-05", "여름"},
		{13, "2020-08-05", "여름"},
		{14, "2020-09-05", "가을"},
		{15, "2020-10-05", "가을"},
		{16, "2020-11-05", "가을"},
		{17, "2020-12-05", "겨울"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(korean)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Season(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_Korean_ToMonthString(t *testing.T) {
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

		{6, "2020-01-05", "일월"},
		{7, "2020-02-05", "이월"},
		{8, "2020-03-05", "삼월"},
		{9, "2020-04-05", "사월"},
		{10, "2020-05-05", "오월"},
		{11, "2020-06-05", "유월"},
		{12, "2020-07-05", "칠월"},
		{13, "2020-08-05", "팔월"},
		{14, "2020-09-05", "구월"},
		{15, "2020-10-05", "시월"},
		{16, "2020-11-05", "십일월"},
		{17, "2020-12-05", "십이월"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(korean)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToMonthString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_Korean_ToShortMonthString(t *testing.T) {
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

		{6, "2020-01-05", "1월"},
		{7, "2020-02-05", "2월"},
		{8, "2020-03-05", "3월"},
		{9, "2020-04-05", "4월"},
		{10, "2020-05-05", "5월"},
		{11, "2020-06-05", "6월"},
		{12, "2020-07-05", "7월"},
		{13, "2020-08-05", "8월"},
		{14, "2020-09-05", "9월"},
		{15, "2020-10-05", "10월"},
		{16, "2020-11-05", "11월"},
		{17, "2020-12-05", "12월"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(korean)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortMonthString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_Korean_ToWeekString(t *testing.T) {
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

		{6, "2020-08-01", "토요일"},
		{7, "2020-08-02", "일요일"},
		{8, "2020-08-03", "월요일"},
		{9, "2020-08-04", "화요일"},
		{10, "2020-08-05", "수요일"},
		{11, "2020-08-06", "목요일"},
		{12, "2020-08-07", "금요일"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(korean)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToWeekString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_Korean_ToShortWeekString(t *testing.T) {
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

		{6, "2020-08-01", "토요일"},
		{7, "2020-08-02", "일요일"},
		{8, "2020-08-03", "월요일"},
		{9, "2020-08-04", "화요일"},
		{10, "2020-08-05", "수요일"},
		{11, "2020-08-06", "목요일"},
		{12, "2020-08-07", "금요일"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(korean)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortWeekString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_Korean_DiffForHumans(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input1   string // 输入值
		input2   string // 输入值
		expected string // 期望值
	}{
		{1, "2020-08-05 13:14:15", "2020-08-05 13:14:15", "방금"},
		{2, "2020-08-05 13:14:15", "2021-08-05 13:14:15", "1 년전"},
		{3, "2020-08-05 13:14:15", "2019-08-05 13:14:15", "1 년후"},
		{4, "2020-08-05 13:14:15", "2030-08-05 13:14:15", "10 년전"},
		{5, "2020-08-05 13:14:15", "2010-08-05 13:14:15", "10 년후"},

		{6, "2020-08-05 13:14:15", "2020-09-05 13:14:15", "1 개월전"},
		{7, "2020-08-05 13:14:15", "2020-07-05 13:14:15", "1 개월후"},
		{8, "2020-08-05 13:14:15", "2021-06-05 13:14:15", "10 개월전"},
		{9, "2020-08-05 13:14:15", "2019-10-05 13:14:15", "10 개월후"},

		{10, "2020-08-05 13:14:15", "2020-08-06 13:14:15", "1 일전"},
		{11, "2020-08-05 13:14:15", "2020-08-04 13:14:15", "1 일후"},
		{12, "2020-08-05 13:14:15", "2020-08-15 13:14:15", "1 주전"},
		{13, "2020-08-05 13:14:15", "2020-07-26 13:14:15", "1 주후"},

		{14, "2020-08-05 13:14:15", "2020-08-05 14:14:15", "1 시간전"},
		{15, "2020-08-05 13:14:15", "2020-08-05 12:14:15", "1 시간후"},
		{16, "2020-08-05 13:14:15", "2020-08-05 23:14:15", "10 시간전"},
		{17, "2020-08-05 13:14:15", "2020-08-05 03:14:15", "10 시간후"},

		{18, "2020-08-05 13:14:15", "2020-08-05 13:15:15", "1 분전"},
		{19, "2020-08-05 13:14:15", "2020-08-05 13:13:15", "1 분후"},
		{20, "2020-08-05 13:14:15", "2020-08-05 13:24:15", "10 분전"},
		{21, "2020-08-05 13:14:15", "2020-08-05 13:04:15", "10 분후"},

		{22, "2020-08-05 13:14:15", "2020-08-05 13:14:16", "1 초전"},
		{23, "2020-08-05 13:14:15", "2020-08-05 13:14:14", "1 초후"},
		{24, "2020-08-05 13:14:15", "2020-08-05 13:14:25", "10 초전"},
		{25, "2020-08-05 13:14:15", "2020-08-05 13:14:05", "10 초후"},
	}

	for _, test := range tests {
		c1 := Parse(test.input1)
		c2 := Parse(test.input2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.SetLocale(korean).DiffForHumans(c2), "Current test id is "+strconv.Itoa(test.id))
	}
}
