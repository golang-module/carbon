package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_Season1(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		output string // 期望输出值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},
		{5, "0000-00-00 00:00:00", ""},

		{6, "2020-01-05", "Winter"},
		{7, "2020-02-05", "Winter"},
		{8, "2020-03-05", "Spring"},
		{9, "2020-04-05", "Spring"},
		{10, "2020-05-05", "Spring"},
		{11, "2020-06-05", "Summer"},
		{12, "2020-07-05", "Summer"},
		{13, "2020-08-05", "Summer"},
		{14, "2020-09-05", "Autumn"},
		{15, "2020-10-05", "Autumn"},
		{16, "2020-11-05", "Autumn"},
		{17, "2020-12-05", "Winter"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Season(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Season2(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		param  string // 参数
		output string // 期望输出值
	}{
		{1, "", "en", ""},
		{2, "0", "en", ""},
		{3, "0000-00-00", "en", ""},
		{4, "00:00:00", "en", ""},
		{5, "0000-00-00 00:00:00", "en", ""},

		{6, "2020-08-05", "en", "Summer"},
		{7, "2020-08-05", "zh-CN", "夏季"},
		{8, "2020-08-05", "zh-Tw", "夏季"},
		{9, "2020-08-05", "jp", "なつ"},
		{10, "2020-08-05", "kr", "여름"},
	}

	for _, test := range tests {
		c := Parse(test.input).SetLocale(test.param)
		assert.Nil(c.Error)
		assert.Equal(c.Season(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_StartOfSeason(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		output string // 期望输出值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},
		{5, "0000-00-00 00:00:00", ""},

		{6, "2020-01-15 00:00:00", "2019-12-01 00:00:00"},
		{7, "2020-02-15 00:00:00", "2019-12-01 00:00:00"},
		{8, "2020-03-15 00:00:00", "2020-03-01 00:00:00"},
		{9, "2020-04-15 23:59:59", "2020-03-01 00:00:00"},
		{10, "2020-05-15 23:59:59", "2020-03-01 00:00:00"},
		{11, "2020-06-15 23:59:59", "2020-06-01 00:00:00"},
		{12, "2020-07-15 23:59:59", "2020-06-01 00:00:00"},
		{13, "2020-08-15 13:14:15", "2020-06-01 00:00:00"},
		{14, "2020-09-15 13:14:15", "2020-09-01 00:00:00"},
		{15, "2020-10-15", "2020-09-01 00:00:00"},
		{16, "2020-11-15", "2020-09-01 00:00:00"},
		{17, "2020-12-15", "2020-12-01 00:00:00"},
		{18, "2021-01-15", "2020-12-01 00:00:00"},
		{19, "2021-01-15", "2020-12-01 00:00:00"},
	}

	for _, test := range tests {
		c := Parse(test.input).StartOfSeason()
		assert.Nil(c.Error)
		assert.Equal(c.ToDateTimeString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_EndOfSeason(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		output string // 期望输出值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},
		{5, "0000-00-00 00:00:00", ""},

		{6, "2020-01-15 00:00:00", "2020-02-29 23:59:59"},
		{7, "2020-02-15 00:00:00", "2020-02-29 23:59:59"},
		{8, "2020-03-15 00:00:00", "2020-05-31 23:59:59"},
		{9, "2020-04-15 23:59:59", "2020-05-31 23:59:59"},
		{10, "2020-05-15 23:59:59", "2020-05-31 23:59:59"},
		{11, "2020-06-15 23:59:59", "2020-08-31 23:59:59"},
		{12, "2020-07-15 23:59:59", "2020-08-31 23:59:59"},
		{13, "2020-08-15 13:14:15", "2020-08-31 23:59:59"},
		{14, "2020-09-15 13:14:15", "2020-11-30 23:59:59"},
		{15, "2020-10-15", "2020-11-30 23:59:59"},
		{16, "2020-11-15", "2020-11-30 23:59:59"},
		{17, "2020-12-15", "2021-02-28 23:59:59"},
		{18, "2021-01-15", "2021-02-28 23:59:59"},
		{19, "2021-02-15", "2021-02-28 23:59:59"},
	}

	for _, test := range tests {
		c := Parse(test.input).EndOfSeason()
		assert.Nil(c.Error)
		assert.Equal(c.ToDateTimeString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsSpring(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{1, "", false},
		{2, "0", false},
		{3, "0000-00-00", false},
		{4, "00:00:00", false},
		{5, "0000-00-00 00:00:00", false},

		{6, "2020-01-01", false},
		{7, "2020-02-01", false},
		{8, "2020-03-01", true},
		{9, "2020-04-01", true},
		{10, "2020-05-01", true},
		{11, "2020-06-01", false},
		{12, "2020-07-01", false},
		{13, "2020-08-01", false},
		{14, "2020-09-01", false},
		{15, "2020-10-01", false},
		{16, "2020-11-01", false},
		{17, "2020-12-01", false},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsSpring(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsSummer(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{1, "", false},
		{2, "0", false},
		{3, "0000-00-00", false},
		{4, "00:00:00", false},
		{5, "0000-00-00 00:00:00", false},

		{6, "2020-01-01", false},
		{7, "2020-02-01", false},
		{8, "2020-03-01", false},
		{9, "2020-04-01", false},
		{10, "2020-05-01", false},
		{11, "2020-06-01", true},
		{12, "2020-07-01", true},
		{13, "2020-08-01", true},
		{14, "2020-09-01", false},
		{15, "2020-10-01", false},
		{16, "2020-11-01", false},
		{17, "2020-12-01", false},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsSummer(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsAutumn(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{1, "", false},
		{2, "0", false},
		{3, "0000-00-00", false},
		{4, "00:00:00", false},
		{5, "0000-00-00 00:00:00", false},

		{6, "2020-01-01", false},
		{7, "2020-02-01", false},
		{8, "2020-03-01", false},
		{9, "2020-04-01", false},
		{10, "2020-05-01", false},
		{11, "2020-06-01", false},
		{12, "2020-07-01", false},
		{13, "2020-08-01", false},
		{14, "2020-09-01", true},
		{15, "2020-10-01", true},
		{16, "2020-11-01", true},
		{17, "2020-12-01", false},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsAutumn(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsWinter(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{1, "", false},
		{2, "0", false},
		{3, "0000-00-00", false},
		{4, "00:00:00", false},
		{5, "0000-00-00 00:00:00", false},

		{6, "2020-01-01", true},
		{7, "2020-02-01", true},
		{8, "2020-03-01", false},
		{9, "2020-04-01", false},
		{10, "2020-05-01", false},
		{11, "2020-06-01", false},
		{12, "2020-07-01", false},
		{13, "2020-08-01", false},
		{14, "2020-09-01", false},
		{15, "2020-10-01", false},
		{16, "2020-11-01", false},
		{17, "2020-12-01", true},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsWinter(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}
