package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_DoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值1
		expected string // 期望值
	}{
		{"", ""},
		{"2020-01-05 23:23:45", "子时"},
		{"2020-02-05 01:00:00", "丑时"},
		{"2020-02-05 03:40:00", "寅时"},
		{"2020-02-05 05:00:00", "卯时"},
		{"2020-02-05 07:30:00", "辰时"},
		{"2020-02-05 09:00:00", "巳时"},
		{"2020-02-05 11:00:00", "午时"},
		{"2020-02-05 13:00:00", "未时"},
		{"2020-02-05 15:00:00", "申时"},
		{"2020-02-05 17:00:00", "酉时"},
		{"2020-02-05 19:00:00", "戌时"},
		{"2020-02-05 21:00:00", "亥时"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		c.lang.SetLocale("zh-CN")
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsFirstDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 00:00:00", true},
		{"2020-04-19 00:59:59", true},
		{"2020-08-05 01:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		c.SetLocale("zh-CN")
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsFirstDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSecondDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 01:00:00", true},
		{"2020-04-19 02:59:59", true},
		{"2020-08-05 03:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		c.SetLocale("zh-CN")
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsSecondDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsThirdDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 03:00:00", true},
		{"2020-04-19 04:59:59", true},
		{"2020-08-05 05:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		c.SetLocale("zh-CN")
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsThirdDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}
func TestCarbon_IsFourthDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 05:00:00", true},
		{"2020-04-19 06:59:59", true},
		{"2020-08-05 07:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		c.SetLocale("zh-CN")
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsFourthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsFifthDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 07:00:00", true},
		{"2020-04-19 08:59:59", true},
		{"2020-08-05 09:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		c.SetLocale("zh-CN")
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsFifthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSixthDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 09:00:00", true},
		{"2020-04-19 10:59:59", true},
		{"2020-08-05 11:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		c.SetLocale("zh-CN")
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsSixthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSeventhDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 11:00:00", true},
		{"2020-04-19 12:59:59", true},
		{"2020-08-05 13:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		c.SetLocale("zh-CN")
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsSeventhDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsEighthDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 13:00:00", true},
		{"2020-04-19 14:59:59", true},
		{"2020-08-05 15:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		c.SetLocale("zh-CN")
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsEighthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsNinthDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 15:00:00", true},
		{"2020-04-19 16:59:59", true},
		{"2020-08-05 17:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		c.SetLocale("zh-CN")
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsNinthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsTenthDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 17:00:00", true},
		{"2020-04-19 18:59:59", true},
		{"2020-08-05 19:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		c.SetLocale("zh-CN")
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsTenthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsEleventhDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 19:00:00", true},
		{"2020-04-19 20:59:59", true},
		{"2020-08-05 21:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		c.SetLocale("zh-CN")
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsEleventhDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsTwelfthDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 21:00:00", true},
		{"2020-04-19 22:59:59", true},
		{"2020-08-05 23:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		c.SetLocale("zh-CN")
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsTwelfthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}
