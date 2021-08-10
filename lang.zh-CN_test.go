package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var chs = "zh-CN"

func TestLang_CHS_Constellation(t *testing.T) {
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

		{"2020-01-05", "摩羯座"},
		{"2020-02-05", "水瓶座"},
		{"2020-03-05", "双鱼座"},
		{"2020-04-05", "白羊座"},
		{"2020-05-05", "金牛座"},
		{"2020-06-05", "双子座"},
		{"2020-07-05", "巨蟹座"},
		{"2020-08-05", "狮子座"},
		{"2020-09-05", "处女座"},
		{"2020-10-05", "天秤座"},
		{"2020-11-05", "天蝎座"},
		{"2020-12-05", "射手座"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(chs)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Constellation(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_CHS_Season(t *testing.T) {
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

		{"2020-01-05", "冬季"},
		{"2020-02-05", "冬季"},
		{"2020-03-05", "春季"},
		{"2020-04-05", "春季"},
		{"2020-05-05", "春季"},
		{"2020-06-05", "夏季"},
		{"2020-07-05", "夏季"},
		{"2020-08-05", "夏季"},
		{"2020-09-05", "秋季"},
		{"2020-10-05", "秋季"},
		{"2020-11-05", "秋季"},
		{"2020-12-05", "冬季"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(chs)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Season(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_CHS_ToMonthString(t *testing.T) {
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

		{"2020-01-05", "一月"},
		{"2020-02-05", "二月"},
		{"2020-03-05", "三月"},
		{"2020-04-05", "四月"},
		{"2020-05-05", "五月"},
		{"2020-06-05", "六月"},
		{"2020-07-05", "七月"},
		{"2020-08-05", "八月"},
		{"2020-09-05", "九月"},
		{"2020-10-05", "十月"},
		{"2020-11-05", "十一月"},
		{"2020-12-05", "十二月"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(chs)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToMonthString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_CHS_ToShortMonthString(t *testing.T) {
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

		{"2020-01-05", "1月"},
		{"2020-02-05", "2月"},
		{"2020-03-05", "3月"},
		{"2020-04-05", "4月"},
		{"2020-05-05", "5月"},
		{"2020-06-05", "6月"},
		{"2020-07-05", "7月"},
		{"2020-08-05", "8月"},
		{"2020-09-05", "9月"},
		{"2020-10-05", "10月"},
		{"2020-11-05", "11月"},
		{"2020-12-05", "12月"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(chs)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortMonthString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_CHS_ToWeekString(t *testing.T) {
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

		{"2020-08-01", "星期六"},
		{"2020-08-02", "星期日"},
		{"2020-08-03", "星期一"},
		{"2020-08-04", "星期二"},
		{"2020-08-05", "星期三"},
		{"2020-08-06", "星期四"},
		{"2020-08-07", "星期五"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(chs)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToWeekString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_CHS_ToShortWeekString(t *testing.T) {
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

		{"2020-08-01", "周六"},
		{"2020-08-02", "周日"},
		{"2020-08-03", "周一"},
		{"2020-08-04", "周二"},
		{"2020-08-05", "周三"},
		{"2020-08-06", "周四"},
		{"2020-08-07", "周五"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(chs)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortWeekString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_CHS_DiffForHumans(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   string // 输入值
		input2   string // 输入值
		expected string // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", "刚刚"},
		{"2020-08-05 13:14:15", "2021-08-05 13:14:15", "1 年前"},
		{"2020-08-05 13:14:15", "2019-08-05 13:14:15", "1 年后"},
		{"2020-08-05 13:14:15", "2030-08-05 13:14:15", "10 年前"},
		{"2020-08-05 13:14:15", "2010-08-05 13:14:15", "10 年后"},

		{"2020-08-05 13:14:15", "2020-09-05 13:14:15", "1 个月前"},
		{"2020-08-05 13:14:15", "2020-07-05 13:14:15", "1 个月后"},
		{"2020-08-05 13:14:15", "2021-06-05 13:14:15", "10 个月前"},
		{"2020-08-05 13:14:15", "2019-10-05 13:14:15", "10 个月后"},

		{"2020-08-05 13:14:15", "2020-08-06 13:14:15", "1 天前"},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", "1 天后"},
		{"2020-08-05 13:14:15", "2020-08-15 13:14:15", "1 周前"},
		{"2020-08-05 13:14:15", "2020-07-26 13:14:15", "1 周后"},

		{"2020-08-05 13:14:15", "2020-08-05 14:14:15", "1 小时前"},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:15", "1 小时后"},
		{"2020-08-05 13:14:15", "2020-08-05 23:14:15", "10 小时前"},
		{"2020-08-05 13:14:15", "2020-08-05 03:14:15", "10 小时后"},

		{"2020-08-05 13:14:15", "2020-08-05 13:15:15", "1 分钟前"},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:15", "1 分钟后"},
		{"2020-08-05 13:14:15", "2020-08-05 13:24:15", "10 分钟前"},
		{"2020-08-05 13:14:15", "2020-08-05 13:04:15", "10 分钟后"},

		{"2020-08-05 13:14:15", "2020-08-05 13:14:16", "1 秒前"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:14", "1 秒后"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:25", "10 秒前"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:05", "10 秒后"},
	}

	for index, test := range tests {
		c1 := Parse(test.input1)
		c2 := Parse(test.input2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.SetLocale(chs).DiffForHumans(c2), "Current test index is "+strconv.Itoa(index))
	}
}
