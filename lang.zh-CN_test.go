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
		id       int    // 测试id
		input    string // 输入值
		expected string // 期望值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},
		{5, "0000-00-00 00:00:00", ""},

		{6, "2020-01-05", "摩羯座"},
		{7, "2020-02-05", "水瓶座"},
		{8, "2020-03-05", "双鱼座"},
		{9, "2020-04-05", "白羊座"},
		{10, "2020-05-05", "金牛座"},
		{11, "2020-06-05", "双子座"},
		{12, "2020-07-05", "巨蟹座"},
		{13, "2020-08-05", "狮子座"},
		{14, "2020-09-05", "处女座"},
		{15, "2020-10-05", "天秤座"},
		{16, "2020-11-05", "天蝎座"},
		{17, "2020-12-05", "射手座"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(chs)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Constellation(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_CHS_Season(t *testing.T) {
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

		{6, "2020-01-05", "冬季"},
		{7, "2020-02-05", "冬季"},
		{8, "2020-03-05", "春季"},
		{9, "2020-04-05", "春季"},
		{10, "2020-05-05", "春季"},
		{11, "2020-06-05", "夏季"},
		{12, "2020-07-05", "夏季"},
		{13, "2020-08-05", "夏季"},
		{14, "2020-09-05", "秋季"},
		{15, "2020-10-05", "秋季"},
		{16, "2020-11-05", "秋季"},
		{17, "2020-12-05", "冬季"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(chs)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Season(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_CHS_ToMonthString(t *testing.T) {
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

		{6, "2020-01-05", "一月"},
		{7, "2020-02-05", "二月"},
		{8, "2020-03-05", "三月"},
		{9, "2020-04-05", "四月"},
		{10, "2020-05-05", "五月"},
		{11, "2020-06-05", "六月"},
		{12, "2020-07-05", "七月"},
		{13, "2020-08-05", "八月"},
		{14, "2020-09-05", "九月"},
		{15, "2020-10-05", "十月"},
		{16, "2020-11-05", "十一月"},
		{17, "2020-12-05", "十二月"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(chs)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToMonthString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_CHS_ToShortMonthString(t *testing.T) {
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

		{6, "2020-01-05", "1月"},
		{7, "2020-02-05", "2月"},
		{8, "2020-03-05", "3月"},
		{9, "2020-04-05", "4月"},
		{10, "2020-05-05", "5月"},
		{11, "2020-06-05", "6月"},
		{12, "2020-07-05", "7月"},
		{13, "2020-08-05", "8月"},
		{14, "2020-09-05", "9月"},
		{15, "2020-10-05", "10月"},
		{16, "2020-11-05", "11月"},
		{17, "2020-12-05", "12月"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(chs)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortMonthString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_CHS_ToWeekString(t *testing.T) {
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

		{6, "2020-08-01", "星期六"},
		{7, "2020-08-02", "星期日"},
		{8, "2020-08-03", "星期一"},
		{9, "2020-08-04", "星期二"},
		{10, "2020-08-05", "星期三"},
		{11, "2020-08-06", "星期四"},
		{12, "2020-08-07", "星期五"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(chs)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToWeekString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_CHS_ToShortWeekString(t *testing.T) {
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

		{6, "2020-08-01", "周六"},
		{7, "2020-08-02", "周日"},
		{8, "2020-08-03", "周一"},
		{9, "2020-08-04", "周二"},
		{10, "2020-08-05", "周三"},
		{11, "2020-08-06", "周四"},
		{12, "2020-08-07", "周五"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(chs)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortWeekString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_CHS_DiffForHumans(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input1   string // 输入值
		input2   string // 输入值
		expected string // 期望值
	}{
		{1, "2020-08-05 13:14:15", "2020-08-05 13:14:15", "刚刚"},
		{2, "2020-08-05 13:14:15", "2021-08-05 13:14:15", "1 年前"},
		{3, "2020-08-05 13:14:15", "2019-08-05 13:14:15", "1 年后"},
		{4, "2020-08-05 13:14:15", "2030-08-05 13:14:15", "10 年前"},
		{5, "2020-08-05 13:14:15", "2010-08-05 13:14:15", "10 年后"},

		{6, "2020-08-05 13:14:15", "2020-09-05 13:14:15", "1 个月前"},
		{7, "2020-08-05 13:14:15", "2020-07-05 13:14:15", "1 个月后"},
		{8, "2020-08-05 13:14:15", "2021-06-05 13:14:15", "10 个月前"},
		{9, "2020-08-05 13:14:15", "2019-10-05 13:14:15", "10 个月后"},

		{10, "2020-08-05 13:14:15", "2020-08-06 13:14:15", "1 天前"},
		{11, "2020-08-05 13:14:15", "2020-08-04 13:14:15", "1 天后"},
		{12, "2020-08-05 13:14:15", "2020-08-15 13:14:15", "1 周前"},
		{13, "2020-08-05 13:14:15", "2020-07-26 13:14:15", "1 周后"},

		{14, "2020-08-05 13:14:15", "2020-08-05 14:14:15", "1 小时前"},
		{15, "2020-08-05 13:14:15", "2020-08-05 12:14:15", "1 小时后"},
		{16, "2020-08-05 13:14:15", "2020-08-05 23:14:15", "10 小时前"},
		{17, "2020-08-05 13:14:15", "2020-08-05 03:14:15", "10 小时后"},

		{18, "2020-08-05 13:14:15", "2020-08-05 13:15:15", "1 分钟前"},
		{19, "2020-08-05 13:14:15", "2020-08-05 13:13:15", "1 分钟后"},
		{20, "2020-08-05 13:14:15", "2020-08-05 13:24:15", "10 分钟前"},
		{21, "2020-08-05 13:14:15", "2020-08-05 13:04:15", "10 分钟后"},

		{22, "2020-08-05 13:14:15", "2020-08-05 13:14:16", "1 秒前"},
		{23, "2020-08-05 13:14:15", "2020-08-05 13:14:14", "1 秒后"},
		{24, "2020-08-05 13:14:15", "2020-08-05 13:14:25", "10 秒前"},
		{25, "2020-08-05 13:14:15", "2020-08-05 13:14:05", "10 秒后"},
	}

	for _, test := range tests {
		c1 := Parse(test.input1)
		c2 := Parse(test.input2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.SetLocale(chs).DiffForHumans(c2), "Current test id is "+strconv.Itoa(test.id))
	}
}
