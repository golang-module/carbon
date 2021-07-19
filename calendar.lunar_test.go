package carbon

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLunar_Animal(t *testing.T) {
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

		{6, "2020-05-01", "鼠"},
		{7, "2020-08-05", "鼠"},
		{8, "2021-07-07", "牛"},
		{9, "2010-08-05", "虎"},
		{10, "2011-08-05", "兔"},
		{11, "2012-08-05", "龙"},
		{12, "2013-08-05", "蛇"},
		{13, "2014-08-05", "马"},
		{14, "2015-08-05", "羊"},
		{15, "2016-08-05", "猴"},
		{16, "2017-08-05", "鸡"},
		{17, "2018-08-05", "狗"},
		{18, "2019-08-05", "猪"},
		{19, "2020-04-23", "鼠"},
		{20, "2020-08-05", "鼠"},
		{21, "2021-05-12", "牛"},
		{22, "2021-08-05", "牛"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().Animal(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_Festival(t *testing.T) {
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

		{6, "2020-04-23", ""},
		{7, "2021-02-12", "春节"},
		{8, "2021-02-26", "元宵节"},
		{9, "2021-05-12", ""},
		{10, "2021-06-14", "端午节"},
		{11, "2021-08-14", "七夕节"},
		{12, "2021-08-22", "中元节"},
		{13, "2021-09-21", "中秋节"},
		{14, "2021-10-14", "重阳节"},
		{15, "2021-10-14", "重阳节"},
		{16, "2021-11-05", "寒衣节"},
		{17, "2021-11-19", "下元节"},
		{18, "2022-01-10", "腊八节"},
		{19, "2022-01-25", "小年"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().Festival(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_Year(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		output int    // 期望输出值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-04-23", 2020},
		{7, "2020-05-01", 2020},
		{8, "2020-08-05", 2020},
		{9, "2021-01-01", 2020},
		{10, "2021-05-12", 2021},
		{11, "2021-07-07", 2021},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().Year(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_Month(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		output int    // 期望输出值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-04-23", 4},
		{7, "2020-05-01", 4},
		{8, "2020-08-05", 6},
		{9, "2021-07-07", 5},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().Month(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_Day(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		output int    // 期望输出值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-04-23", 1},
		{7, "2020-05-01", 9},
		{8, "2020-08-05", 16},
		{9, "2021-07-07", 28},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().Day(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_LeapMonth(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		output int    // 期望输出值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-04-23", 4},
		{7, "2020-05-01", 4},
		{8, "2020-08-05", 4},
		{9, "2021-07-07", 0},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().LeapMonth(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_ToYearString(t *testing.T) {
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

		{6, "2020-04-23", "二零二零"},
		{7, "2020-05-01", "二零二零"},
		{8, "2020-08-05", "二零二零"},
		{9, "2021-01-01", "二零二零"},
		{10, "2021-05-12", "二零二一"},
		{11, "2021-07-07", "二零二一"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().ToYearString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_ToMonthString(t *testing.T) {
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

		{6, "2020-01-01", "腊"},
		{7, "2020-02-01", "正"},
		{8, "2020-03-01", "二"},
		{9, "2020-04-01", "三"},
		{10, "2020-04-23", "四"},
		{11, "2020-05-01", "四"},
		{12, "2020-06-01", "四"},
		{13, "2020-07-01", "五"},
		{14, "2020-07-07", "五"},
		{15, "2020-08-01", "六"},
		{16, "2020-09-01", "七"},
		{17, "2020-10-01", "八"},
		{18, "2020-11-01", "九"},
		{19, "2020-12-01", "十"},
		{20, "2021-01-01", "十一"},
		{21, "2021-02-01", "腊"},
		{22, "2021-05-12", "四"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().ToMonthString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_ToDayString(t *testing.T) {
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

		{1, "2020-01-01", "初七"},
		{2, "2020-02-01", "初八"},
		{3, "2020-03-01", "初八"},
		{4, "2020-04-01", "初九"},
		{5, "2020-04-23", "初一"},
		{6, "2020-05-01", "初九"},
		{7, "2020-06-01", "初十"},
		{8, "2020-07-01", "十一"},
		{9, "2020-08-01", "十二"},
		{10, "2020-09-01", "十四"},
		{11, "2020-10-01", "十五"},
		{12, "2020-11-01", "十六"},
		{13, "2020-12-01", "十七"},
		{14, "2021-01-03", "二十"},
		{15, "2021-04-11", "三十"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().ToDayString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_String(t *testing.T) {
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

		{6, "2020-01-01", "二零一九年腊月初七"},
		{7, "2020-02-01", "二零二零年正月初八"},
		{8, "2020-03-01", "二零二零年二月初八"},
		{9, "2020-04-01", "二零二零年三月初九"},
		{10, "2020-04-23", "二零二零年四月初一"},
		{11, "2020-05-01", "二零二零年四月初九"},
		{12, "2020-06-01", "二零二零年四月初十"},
		{13, "2020-07-01", "二零二零年五月十一"},
		{14, "2020-08-01", "二零二零年六月十二"},
		{15, "2020-09-01", "二零二零年七月十四"},
		{16, "2020-10-01", "二零二零年八月十五"},
		{17, "2020-11-01", "二零二零年九月十六"},
		{18, "2020-12-01", "二零二零年十月十七"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(fmt.Sprintf("%s", c.Lunar()), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_ToString(t *testing.T) {
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

		{6, "2020-01-01", "二零一九年腊月初七"},
		{7, "2020-02-01", "二零二零年正月初八"},
		{8, "2020-03-01", "二零二零年二月初八"},
		{9, "2020-04-01", "二零二零年三月初九"},
		{10, "2020-04-23", "二零二零年四月初一"},
		{11, "2020-05-01", "二零二零年四月初九"},
		{12, "2020-06-01", "二零二零年四月初十"},
		{13, "2020-07-01", "二零二零年五月十一"},
		{14, "2020-08-01", "二零二零年六月十二"},
		{15, "2020-09-01", "二零二零年七月十四"},
		{16, "2020-10-01", "二零二零年八月十五"},
		{17, "2020-11-01", "二零二零年九月十六"},
		{18, "2020-12-01", "二零二零年十月十七"},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().ToString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_IsLeapYear(t *testing.T) {
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

		{6, "2020-04-23", true},
		{7, "2020-05-01", true},
		{8, "2020-08-05", true},
		{9, "2021-01-01", true},
		{10, "2021-07-07", false},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().IsLeapYear(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_IsLeapMonth(t *testing.T) {
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

		{6, "2020-04-23", true},
		{7, "2020-05-01", true},
		{8, "2020-08-05", false},
		{9, "2021-01-01", false},
		{10, "2021-07-07", false},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().IsLeapMonth(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_IsRatYear(t *testing.T) {
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

		{6, "2020-05-01", true},
		{7, "2020-08-05", true},
		{8, "2021-01-01", true},
		{9, "2021-07-07", false},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().IsRatYear(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_IsOxYear(t *testing.T) {
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

		{6, "2020-05-01", false},
		{7, "2020-08-05", false},
		{8, "2021-01-01", false},
		{9, "2021-07-07", true},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().IsOxYear(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_IsTigerYear(t *testing.T) {
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

		{6, "2020-05-01", false},
		{7, "2020-08-05", false},
		{8, "2021-01-01", false},
		{9, "2021-07-07", false},
		{10, "2022-08-05", true},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().IsTigerYear(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_IsRabbitYear(t *testing.T) {
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

		{6, "2020-05-01", false},
		{7, "2020-08-05", false},
		{8, "2021-01-01", false},
		{9, "2021-07-07", false},
		{10, "2023-08-05", true},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().IsRabbitYear(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_IsDragonYear(t *testing.T) {
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

		{6, "2020-05-01", false},
		{7, "2020-08-05", false},
		{8, "2021-01-01", false},
		{9, "2021-07-07", false},
		{10, "2024-08-05", true},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().IsDragonYear(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_IsSnakeYear(t *testing.T) {
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

		{6, "2020-05-01", false},
		{7, "2020-08-05", false},
		{8, "2021-01-01", false},
		{9, "2021-07-07", false},
		{10, "2025-08-05", true},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().IsSnakeYear(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_IsHorseYear(t *testing.T) {
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

		{6, "2020-05-01", false},
		{7, "2020-08-05", false},
		{8, "2021-01-01", false},
		{9, "2021-07-07", false},
		{10, "2026-08-05", true},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().IsHorseYear(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_IsGoatYear(t *testing.T) {
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

		{6, "2020-05-01", false},
		{7, "2020-08-05", false},
		{8, "2021-01-01", false},
		{9, "2021-07-07", false},
		{10, "2027-08-05", true},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().IsGoatYear(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_IsMonkeyYear(t *testing.T) {
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

		{6, "2020-05-01", false},
		{7, "2020-08-05", false},
		{8, "2021-01-01", false},
		{9, "2021-07-07", false},
		{10, "2028-08-05", true},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().IsMonkeyYear(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_IsRoosterYear(t *testing.T) {
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

		{6, "2020-05-01", false},
		{7, "2020-08-05", false},
		{8, "2021-01-01", false},
		{9, "2021-07-07", false},
		{10, "2029-08-05", true},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().IsRoosterYear(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_IsDogYear(t *testing.T) {
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

		{6, "2020-05-01", false},
		{7, "2020-08-05", false},
		{8, "2021-01-01", false},
		{9, "2021-07-07", false},
		{10, "2030-08-05", true},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().IsDogYear(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLunar_IsPigYear(t *testing.T) {
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

		{6, "2020-05-01", false},
		{7, "2020-08-05", false},
		{8, "2021-01-01", false},
		{9, "2021-07-07", false},
		{10, "2031-08-05", true},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.Lunar().IsPigYear(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}
