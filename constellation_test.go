package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_Constellation1(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值1
		expected string // 期望值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},
		{5, "0000-00-00 00:00:00", ""},

		{6, "2020-01-05", "Capricorn"},
		{7, "2020-02-05", "Aquarius"},
		{8, "2020-03-05", "Pisces"},
		{9, "2020-04-05", "Aries"},
		{10, "2020-05-05", "Taurus"},
		{11, "2020-06-05", "Gemini"},
		{12, "2020-07-05", "Cancer"},
		{13, "2020-08-05", "Leo"},
		{14, "2020-09-05", "Virgo"},
		{15, "2020-10-05", "Libra"},
		{16, "2020-11-05", "Scorpio"},
		{17, "2020-12-05", "Sagittarius"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Constellation(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Constellation2(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		param    string // 参数值
		expected string // 期望值
	}{
		{1, "", "en", ""},
		{2, "0", "en", ""},
		{3, "0000-00-00", "en", ""},
		{4, "2020-08-05", "en", "Leo"},
		{5, "2020-08-05", "zh-CN", "狮子座"},
		{6, "2020-08-05", "zh-TW", "獅子座"},
		{7, "2020-08-05", "jp", "ししざ"},
		{8, "2020-08-05", "kr", "처녀자리"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(test.param)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Constellation(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsAries(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected bool   // 期望值
	}{
		{1, "", false},
		{2, "0", false},
		{3, "0000-00-00", false},
		{4, "00:00:00", false},
		{5, "0000-00-00 00:00:00", false},

		{6, "2020-03-21", true},
		{7, "2020-04-19", true},
		{8, "2020-08-05", false},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsAries(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsTaurus(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected bool   // 期望值
	}{
		{1, "", false},
		{2, "0", false},
		{3, "0000-00-00", false},
		{4, "00:00:00", false},
		{5, "0000-00-00 00:00:00", false},

		{6, "2020-04-20", true},
		{7, "2020-05-20", true},
		{8, "2020-08-05", false},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsTaurus(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsGemini(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected bool   // 期望值
	}{
		{1, "", false},
		{2, "0", false},
		{3, "0000-00-00", false},
		{4, "00:00:00", false},
		{5, "0000-00-00 00:00:00", false},

		{6, "2020-05-21", true},
		{7, "2020-06-21", true},
		{8, "2020-08-05", false},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsGemini(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsCancer(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected bool   // 期望值
	}{
		{1, "", false},
		{2, "0", false},
		{3, "0000-00-00", false},
		{4, "00:00:00", false},
		{5, "0000-00-00 00:00:00", false},

		{6, "2020-06-22", true},
		{7, "2020-07-22", true},
		{8, "2020-08-05", false},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsCancer(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsLeo(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected bool   // 期望值
	}{
		{1, "", false},
		{2, "0", false},
		{3, "0000-00-00", false},
		{4, "00:00:00", false},
		{5, "0000-00-00 00:00:00", false},

		{6, "2020-07-23", true},
		{7, "2020-08-05", true},
		{8, "2020-08-22", true},
		{9, "2020-08-23", false},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsLeo(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsVirgo(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected bool   // 期望值
	}{
		{1, "", false},
		{2, "0", false},
		{3, "0000-00-00", false},
		{4, "00:00:00", false},
		{5, "0000-00-00 00:00:00", false},

		{6, "2020-08-23", true},
		{7, "2020-09-22", true},
		{8, "2020-08-05", false},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsVirgo(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsLibra(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected bool   // 期望值
	}{
		{1, "", false},
		{2, "0", false},
		{3, "0000-00-00", false},
		{4, "00:00:00", false},
		{5, "0000-00-00 00:00:00", false},

		{6, "2020-09-23", true},
		{7, "2020-10-23", true},
		{8, "2020-08-05", false},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsLibra(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsScorpio(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected bool   // 期望值
	}{
		{1, "", false},
		{2, "0", false},
		{3, "0000-00-00", false},
		{4, "00:00:00", false},
		{5, "0000-00-00 00:00:00", false},

		{6, "2020-10-24", true},
		{7, "2020-11-22", true},
		{8, "2020-08-05", false},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsScorpio(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsSagittarius(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected bool   // 期望值
	}{
		{1, "", false},
		{2, "0", false},
		{3, "0000-00-00", false},
		{4, "00:00:00", false},
		{5, "0000-00-00 00:00:00", false},

		{6, "2020-11-23", true},
		{7, "2020-12-21", true},
		{8, "2020-08-05", false},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsSagittarius(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsCapricorn(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected bool   // 期望值
	}{
		{1, "", false},
		{2, "0", false},
		{3, "0000-00-00", false},
		{4, "00:00:00", false},
		{5, "0000-00-00 00:00:00", false},

		{6, "2020-12-22", true},
		{7, "2020-01-19", true},
		{8, "2020-08-05", false},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsCapricorn(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsAquarius(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected bool   // 期望值
	}{
		{1, "", false},
		{2, "0", false},
		{3, "0000-00-00", false},
		{4, "00:00:00", false},
		{5, "0000-00-00 00:00:00", false},

		{6, "2020-01-20", true},
		{7, "2020-02-18", true},
		{8, "2020-08-05", false},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsAquarius(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsPisces(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected bool   // 期望值
	}{
		{1, "", false},
		{2, "0", false},
		{3, "0000-00-00", false},
		{4, "00:00:00", false},
		{5, "0000-00-00 00:00:00", false},

		{6, "2020-02-19", true},
		{7, "2020-03-20", true},
		{8, "2020-08-05", false},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsPisces(), "Current test id is "+strconv.Itoa(test.id))
	}
}
