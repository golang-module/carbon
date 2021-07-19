package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_IsZero(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{1, "", true},
		{2, "0", true},
		{3, "0000-00-00", true},
		{4, "00:00:00", true},
		{5, "0000-00-00 00:00:00", true},

		{6, "2020-08-05", false},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsZero(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsNow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{1, Parse(""), false},
		{2, Parse("0"), false},
		{3, Parse("0000-00-00"), false},
		{4, Parse("00:00:00"), false},
		{5, Parse("0000-00-00 00:00:00"), false},

		{6, Tomorrow(), false},
		{7, Now(), true},
		{8, Yesterday(), false},
	}

	for _, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(c.IsNow(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsFuture(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{1, Parse(""), false},
		{2, Parse("0"), false},
		{3, Parse("0000-00-00"), false},
		{4, Parse("00:00:00"), false},
		{5, Parse("0000-00-00 00:00:00"), false},

		{6, Tomorrow(), true},
		{7, Now(), false},
		{8, Yesterday(), false},
	}

	for _, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(c.IsFuture(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsPast(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{1, Parse(""), false},
		{2, Parse("0"), false},
		{3, Parse("0000-00-00"), false},
		{4, Parse("00:00:00"), false},
		{5, Parse("0000-00-00 00:00:00"), false},

		{6, Tomorrow(), false},
		{7, Now(), false},
		{8, Yesterday(), true},
	}

	for _, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(c.IsPast(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsLeapYear(t *testing.T) {
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

		{6, "2016-01-01", true},
		{7, "2017-01-01", false},
		{8, "2018-01-01", false},
		{9, "2019-01-01", false},
		{10, "2020-01-01", true},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsLeapYear(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsLongYear(t *testing.T) {
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

		{6, "2015-01-01", true},
		{7, "2016-01-01", false},
		{8, "2017-01-01", false},
		{9, "2018-01-01", false},
		{10, "2019-01-01", false},
		{11, "2020-01-01", true},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsLongYear(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsJanuary(t *testing.T) {
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
		{7, "2020-02-01", false},
		{8, "2020-03-01", false},
		{9, "2020-04-01", false},
		{10, "2020-05-01", false},
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
		assert.Equal(c.IsJanuary(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsFebruary(t *testing.T) {
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
		{17, "2020-12-01", false},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsFebruary(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsMarch(t *testing.T) {
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
		{9, "2020-04-01", false},
		{10, "2020-05-01", false},
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
		assert.Equal(c.IsMarch(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsApril(t *testing.T) {
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
		{9, "2020-04-01", true},
		{10, "2020-05-01", false},
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
		assert.Equal(c.IsApril(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsMay(t *testing.T) {
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
		assert.Equal(c.IsMay(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsJune(t *testing.T) {
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
		assert.Equal(c.IsJune(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsJuly(t *testing.T) {
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
		{12, "2020-07-01", true},
		{13, "2020-08-01", false},
		{14, "2020-09-01", false},
		{15, "2020-10-01", false},
		{16, "2020-11-01", false},
		{17, "2020-12-01", false},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsJuly(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsAugust(t *testing.T) {
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
		{13, "2020-08-01", true},
		{14, "2020-09-01", false},
		{15, "2020-10-01", false},
		{16, "2020-11-01", false},
		{17, "2020-12-01", false},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsAugust(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsSeptember(t *testing.T) {
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
		{15, "2020-10-01", false},
		{16, "2020-11-01", false},
		{17, "2020-12-01", false},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsSeptember(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsOctober(t *testing.T) {
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
		{14, "2020-09-01", false},
		{15, "2020-10-01", true},
		{16, "2020-11-01", false},
		{17, "2020-12-01", false},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsOctober(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsNovember(t *testing.T) {
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
		{14, "2020-09-01", false},
		{15, "2020-10-01", false},
		{16, "2020-11-01", true},
		{17, "2020-12-01", false},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsNovember(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsDecember(t *testing.T) {
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
		{14, "2020-09-01", false},
		{15, "2020-10-01", false},
		{16, "2020-11-01", false},
		{17, "2020-12-01", true},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsDecember(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsMonday(t *testing.T) {
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

		{6, "2020-10-05", true},
		{7, "2020-10-06", false},
		{8, "2020-10-07", false},
		{9, "2020-10-08", false},
		{10, "2020-10-09", false},
		{11, "2020-10-10", false},
		{12, "2020-10-11", false},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsMonday(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsTuesday(t *testing.T) {
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

		{6, "2020-10-05", false},
		{7, "2020-10-06", true},
		{8, "2020-10-07", false},
		{9, "2020-10-08", false},
		{10, "2020-10-09", false},
		{11, "2020-10-10", false},
		{12, "2020-10-11", false},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsTuesday(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsWednesday(t *testing.T) {
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

		{6, "2020-10-05", false},
		{7, "2020-10-06", false},
		{8, "2020-10-07", true},
		{9, "2020-10-08", false},
		{10, "2020-10-09", false},
		{11, "2020-10-10", false},
		{12, "2020-10-11", false},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsWednesday(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsThursday(t *testing.T) {
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

		{6, "2020-10-05", false},
		{7, "2020-10-06", false},
		{8, "2020-10-07", false},
		{9, "2020-10-08", true},
		{10, "2020-10-09", false},
		{11, "2020-10-10", false},
		{12, "2020-10-11", false},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsThursday(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsFriday(t *testing.T) {
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

		{6, "2020-10-05", false},
		{7, "2020-10-06", false},
		{8, "2020-10-07", false},
		{9, "2020-10-08", false},
		{10, "2020-10-09", true},
		{11, "2020-10-10", false},
		{12, "2020-10-11", false},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsFriday(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsSaturday(t *testing.T) {
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

		{6, "2020-10-05", false},
		{7, "2020-10-06", false},
		{8, "2020-10-07", false},
		{9, "2020-10-08", false},
		{10, "2020-10-09", false},
		{11, "2020-10-10", true},
		{12, "2020-10-11", false},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsSaturday(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsSunday(t *testing.T) {
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

		{6, "2020-10-05", false},
		{7, "2020-10-06", false},
		{8, "2020-10-07", false},
		{9, "2020-10-08", false},
		{10, "2020-10-09", false},
		{11, "2020-10-10", false},
		{12, "2020-10-11", true},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsSunday(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsWeekday(t *testing.T) {
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

		{6, "2020-10-05", true},
		{7, "2020-10-06", true},
		{8, "2020-10-07", true},
		{9, "2020-10-08", true},
		{10, "2020-10-09", true},
		{11, "2020-10-10", false},
		{12, "2020-10-11", false},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsWeekday(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsWeekend(t *testing.T) {
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

		{6, "2020-10-05", false},
		{7, "2020-10-06", false},
		{8, "2020-10-07", false},
		{9, "2020-10-08", false},
		{10, "2020-10-09", false},
		{11, "2020-10-10", true},
		{12, "2020-10-11", true},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.IsWeekend(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsYesterday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{1, NewCarbon(), false},
		{2, Now(), false},
		{3, Yesterday(), true},
		{4, Tomorrow(), false},
	}

	for _, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(c.IsYesterday(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsToday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{1, NewCarbon(), false},
		{2, Now(), true},
		{3, Yesterday(), false},
		{4, Tomorrow(), false},
	}

	for _, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(c.IsToday(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_IsTomorrow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{1, NewCarbon(), false},
		{2, Now(), false},
		{3, Yesterday(), false},
		{4, Tomorrow(), true},
	}

	for _, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(c.IsTomorrow(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Compare(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		param1 string // 输入参数1
		param2 string // 输入参数2
		output bool   // 期望输出值
	}{
		{1, "2020-08-05", ">", "2020-08-04", true},
		{2, "2020-08-05", "<", "2020-08-04", false},
		{3, "2020-08-05", "<", "2020-08-06", true},
		{4, "2020-08-05", ">", "2020-08-06", false},
		{5, "2020-08-05", "=", "2020-08-05", true},
		{6, "2020-08-05", ">=", "2020-08-05", true},
		{7, "2020-08-05", "<=", "2020-08-05", true},
		{8, "2020-08-05", "!=", "2020-08-05", false},
		{9, "2020-08-05", "<>", "2020-08-05", false},
		{10, "2020-08-05", "!=", "2020-08-04", true},
		{11, "2020-08-05", "<>", "2020-08-04", true},
		{12, "2020-08-05", "+", "2020-08-04", false},
	}

	for _, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(c1.Compare(test.param1, c2), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Gt(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		param  string // 参数值
		output bool   // 期望输出值
	}{
		{1, "2020-08-05", "2020-08-05", false},
		{2, "2020-08-05", "2020-08-04", true},
		{3, "2020-08-05", "2020-08-06", false},
	}

	for _, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(c1.Gt(c2), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Lt(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		param  string // 参数值
		output bool   // 期望输出值
	}{
		{1, "2020-08-05", "2020-08-05", false},
		{2, "2020-08-05", "2020-08-04", false},
		{3, "2020-08-05", "2020-08-06", true},
	}

	for _, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(c1.Lt(c2), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Eq(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		param  string // 参数值
		output bool   // 期望输出值
	}{
		{1, "2020-08-05", "2020-08-05", true},
		{2, "2020-08-05", "2020-08-04", false},
		{3, "2020-08-05", "2020-08-06", false},
	}

	for _, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(c1.Eq(c2), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Ne(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		param  string // 参数值
		output bool   // 期望输出值
	}{
		{1, "2020-08-05", "2020-08-05", false},
		{2, "2020-08-05", "2020-08-04", true},
		{3, "2020-08-05", "2020-08-06", true},
	}

	for _, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(c1.Ne(c2), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Gte(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		param  string // 参数值
		output bool   // 期望输出值
	}{
		{1, "2020-08-05", "2020-08-05", true},
		{2, "2020-08-05", "2020-08-04", true},
		{3, "2020-08-05", "2020-08-06", false},
	}

	for _, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(c1.Gte(c2), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Lte(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		param  string // 参数值
		output bool   // 期望输出值
	}{
		{1, "2020-08-05", "2020-08-05", true},
		{2, "2020-08-05", "2020-08-04", false},
		{3, "2020-08-05", "2020-08-06", true},
	}

	for _, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(c1.Lte(c2), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Between(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		param1 string // 输入参数1
		param2 string // 输入参数2
		output bool   // 期望输出值
	}{
		{1, "2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-05 13:14:15", false},
		{2, "2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-06 13:14:15", false},
		{3, "2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-05 13:14:15", false},
		{4, "2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-06 13:14:15", true},
	}

	for _, test := range tests {
		c1, c2, c3 := Parse(test.input), Parse(test.param1), Parse(test.param2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Nil(c3.Error)
		assert.Equal(c1.Between(c2, c3), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_BetweenIncludedStartTime(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		param1 string // 输入参数1
		param2 string // 输入参数2
		output bool   // 期望输出值
	}{
		{1, "2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-05 13:14:15", false},
		{2, "2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-06 13:14:15", true},
		{3, "2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-05 13:14:15", false},
		{4, "2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-06 13:14:15", true},
	}

	for _, test := range tests {
		c1, c2, c3 := Parse(test.input), Parse(test.param1), Parse(test.param2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Nil(c3.Error)
		assert.Equal(c1.BetweenIncludedStartTime(c2, c3), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_BetweenIncludedEndTime(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		param1 string // 输入参数1
		param2 string // 输入参数2
		output bool   // 期望输出值
	}{
		{1, "2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-05 13:14:15", false},
		{2, "2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-06 13:14:15", false},
		{3, "2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-05 13:14:15", true},
		{4, "2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-06 13:14:15", true},
	}

	for _, test := range tests {
		c1, c2, c3 := Parse(test.input), Parse(test.param1), Parse(test.param2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Nil(c3.Error)
		assert.Equal(c1.BetweenIncludedEndTime(c2, c3), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_BetweenIncludedBoth(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		param1 string // 输入参数1
		param2 string // 输入参数2
		output bool   // 期望输出值
	}{
		{1, "2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-05 13:14:15", true},
		{2, "2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-06 13:14:15", true},
		{3, "2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-05 13:14:15", true},
		{4, "2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-06 13:14:15", true},
		{5, "2020-08-05 13:14:15", "2020-08-06 13:14:15", "2020-08-06 13:14:15", false},
	}

	for _, test := range tests {
		c1, c2, c3 := Parse(test.input), Parse(test.param1), Parse(test.param2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Nil(c3.Error)
		assert.Equal(c1.BetweenIncludedBoth(c2, c3), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}
