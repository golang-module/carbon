package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_DiffInYears(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{1, "2020-08-05 13:14:15", "2020-07-28 13:14:00", 0},
		{2, "2020-12-31 13:14:15", "2021-01-01 13:14:15", 0},
		{3, "2020-08-05 13:14:15", "2021-08-28 13:14:59", 1},
		{4, "2020-08-05 13:14:15", "2018-08-28 13:14:59", -2},
	}

	for _, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInYears(c2), "Current test id is "+strconv.Itoa(test.id))
	}
}

func Test1Carbon_DiffInYearsWithAbs(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{1, "2020-08-05 13:14:15", "2020-07-28 13:14:00", 0},
		{2, "2020-12-31 13:14:15", "2021-01-01 13:14:15", 0},
		{3, "2020-08-05 13:14:15", "2021-08-28 13:14:59", 1},
		{4, "2020-08-05 13:14:15", "2018-08-28 13:14:59", 2},
	}

	for _, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInYearsWithAbs(c2), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_DiffInMonths(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{1, "2020-08-05 13:14:15", "2020-07-28 13:14:00", 0},
		{4, "2020-08-05 13:14:15", "2020-09-06 13:14:59", 1},
		{2, "2020-12-31 13:14:15", "2021-01-31 13:14:15", 1},
		{3, "2020-08-05 13:14:15", "2021-08-28 13:14:59", 12},
		{4, "2020-08-05 13:14:15", "2018-08-28 13:14:59", -24},
	}

	for _, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInMonths(c2), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_DiffInMonthsWithAbs(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{1, "2020-08-05 13:14:15", "2020-07-28 13:14:00", 0},
		{2, "2020-12-31 13:14:15", "2021-01-01 13:14:15", 0},
		{3, "2020-08-05 13:14:15", "2021-08-28 13:14:59", 12},
		{4, "2020-08-05 13:14:15", "2018-08-28 13:14:59", 24},
	}

	for _, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInMonthsWithAbs(c2), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_DiffInWeeks(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{1, "2020-08-05 13:14:15", "2020-07-28 13:14:00", -1},
		{2, "2020-08-05 13:14:15", "2020-07-28 13:14:15", -1},
		{3, "2020-08-05 13:14:15", "2020-07-28 13:14:59", -1},
		{4, "2020-08-05 13:14:15", "2020-08-05 13:14:00", 0},
		{5, "2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{6, "2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{7, "2020-08-05 13:14:15", "2020-08-12 13:14:00", 0},
		{8, "2020-08-05 13:14:15", "2020-08-12 13:14:15", 1},
		{9, "2020-08-05 13:14:15", "2020-08-12 13:14:59", 1},
	}

	for _, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInWeeks(c2), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_DiffInWeeksWithAbs(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{1, "2020-08-05 13:14:15", "2020-07-28 13:14:00", 1},
		{2, "2020-08-05 13:14:15", "2020-07-28 13:14:15", 1},
		{3, "2020-08-05 13:14:15", "2020-07-28 13:14:59", 1},
		{4, "2020-08-05 13:14:15", "2020-08-05 13:14:00", 0},
		{5, "2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{6, "2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{7, "2020-08-05 13:14:15", "2020-08-12 13:14:00", 0},
		{8, "2020-08-05 13:14:15", "2020-08-12 13:14:15", 1},
		{9, "2020-08-05 13:14:15", "2020-08-12 13:14:59", 1},
	}

	for _, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInWeeksWithAbs(c2), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_DiffInDays(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{1, "2020-08-05 13:14:15", "2020-08-04 13:00:00", -1},
		{2, "2020-08-05 13:14:15", "2020-08-04 13:14:15", -1},
		{3, "2020-08-05 13:14:15", "2020-08-04 13:14:59", 0},
		{4, "2020-08-05 13:14:15", "2020-08-05 13:00:00", 0},
		{5, "2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{6, "2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{7, "2020-08-05 13:14:15", "2020-08-06 13:00:00", 0},
		{8, "2020-08-05 13:14:15", "2020-08-06 13:14:15", 1},
		{9, "2020-08-05 13:14:15", "2020-08-06 13:14:59", 1},
	}

	for _, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInDays(c2), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_DiffInDaysWithAbs(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{1, "2020-08-05 13:14:15", "2020-08-04 13:00:00", 1},
		{2, "2020-08-05 13:14:15", "2020-08-04 13:14:15", 1},
		{3, "2020-08-05 13:14:15", "2020-08-04 13:14:59", 0},
		{4, "2020-08-05 13:14:15", "2020-08-05 13:00:00", 0},
		{5, "2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{6, "2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{7, "2020-08-05 13:14:15", "2020-08-06 13:00:00", 0},
		{8, "2020-08-05 13:14:15", "2020-08-06 13:14:15", 1},
		{9, "2020-08-05 13:14:15", "2020-08-06 13:14:59", 1},
	}

	for _, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInDaysWithAbs(c2), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_DiffInHours(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{1, "2020-08-05 13:14:15", "2020-08-05 12:14:00", -1},
		{2, "2020-08-05 13:14:15", "2020-08-05 12:14:15", -1},
		{3, "2020-08-05 13:14:15", "2020-08-05 12:14:59", 0},
		{4, "2020-08-05 13:14:15", "2020-08-05 13:14:00", 0},
		{5, "2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{6, "2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{7, "2020-08-05 13:14:15", "2020-08-05 14:14:00", 0},
		{8, "2020-08-05 13:14:15", "2020-08-05 14:14:15", 1},
		{9, "2020-08-05 13:14:15", "2020-08-05 14:14:59", 1},
	}

	for _, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInHours(c2), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_DiffInHoursWithAbs(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{1, "2020-08-05 13:14:15", "2020-08-05 12:14:00", 1},
		{2, "2020-08-05 13:14:15", "2020-08-05 12:14:15", 1},
		{3, "2020-08-05 13:14:15", "2020-08-05 12:14:59", 0},
		{4, "2020-08-05 13:14:15", "2020-08-05 13:14:00", 0},
		{5, "2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{6, "2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{7, "2020-08-05 13:14:15", "2020-08-05 14:14:00", 0},
		{8, "2020-08-05 13:14:15", "2020-08-05 14:14:15", 1},
		{9, "2020-08-05 13:14:15", "2020-08-05 14:14:59", 1},
	}

	for _, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInHoursWithAbs(c2), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_DiffInMinutes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{1, "2020-08-05 13:14:15", "2020-08-05 13:13:00", -1},
		{2, "2020-08-05 13:14:15", "2020-08-05 13:13:15", -1},
		{3, "2020-08-05 13:14:15", "2020-08-05 13:13:59", 0},
		{4, "2020-08-05 13:14:15", "2020-08-05 13:15:00", 0},
		{5, "2020-08-05 13:14:15", "2020-08-05 13:15:15", 1},
		{6, "2020-08-05 13:14:15", "2020-08-05 13:15:59", 1},
		{7, "2020-08-05 13:14:15", "2020-08-05 13:16:00", 1},
		{8, "2020-08-05 13:14:15", "2020-08-05 13:16:15", 2},
		{9, "2020-08-05 13:14:15", "2020-08-05 13:16:59", 2},
	}

	for _, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInMinutes(c2), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_DiffInMinutesWithAbs(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{1, "2020-08-05 13:14:15", "2020-08-05 13:13:00", 1},
		{2, "2020-08-05 13:14:15", "2020-08-05 13:13:15", 1},
		{3, "2020-08-05 13:14:15", "2020-08-05 13:13:59", 0},
		{4, "2020-08-05 13:14:15", "2020-08-05 13:15:00", 0},
		{5, "2020-08-05 13:14:15", "2020-08-05 13:15:15", 1},
		{6, "2020-08-05 13:14:15", "2020-08-05 13:15:59", 1},
		{7, "2020-08-05 13:14:15", "2020-08-05 13:16:00", 1},
		{8, "2020-08-05 13:14:15", "2020-08-05 13:16:15", 2},
		{9, "2020-08-05 13:14:15", "2020-08-05 13:16:59", 2},
	}

	for _, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInMinutesWithAbs(c2), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_DiffInSeconds(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{1, "2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{2, "2020-08-05 13:14:15", "2020-08-05 13:14:20", 5},
		{3, "2020-08-05 13:14:20", "2020-08-05 13:14:15", -5},
	}

	for _, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInSeconds(c2), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_DiffInSecondsWithAbs(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{1, "2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{2, "2020-08-05 13:14:15", "2020-08-05 13:14:20", 5},
		{3, "2020-08-05 13:14:20", "2020-08-05 13:14:15", 5},
	}

	for _, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInSecondsWithAbs(c2), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_DiffForHumans(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    Carbon // 输入值
		expected string // 期望值
	}{
		{1, Now(), "just now"},

		{2, Now().AddYears(1), "1 year from now"},
		{3, Now().SubYears(1), "1 year ago"},
		{4, Now().AddYears(10), "10 years from now"},
		{5, Now().SubYears(10), "10 years ago"},
		{2, Now().AddYearsNoOverflow(1), "1 year from now"},
		{3, Now().SubYearsNoOverflow(1), "1 year ago"},
		{4, Now().AddYearsNoOverflow(10), "10 years from now"},
		{5, Now().SubYearsNoOverflow(10), "10 years ago"},

		{6, Now().AddMonths(1), "1 month from now"},
		{7, Now().SubMonths(1), "4 weeks ago"},
		{8, Now().AddMonths(10), "10 months from now"},
		{9, Now().SubMonths(10), "9 months ago"},
		{6, Now().AddMonthsNoOverflow(1), "1 month from now"},
		{7, Now().SubMonthsNoOverflow(1), "1 month ago"},
		{8, Now().AddMonthsNoOverflow(10), "10 months from now"},
		{9, Now().SubMonthsNoOverflow(10), "10 months ago"},

		{10, Now().AddDays(1), "1 day from now"},
		{11, Now().SubDays(1), "1 day ago"},
		{12, Now().AddDays(10), "1 week from now"},
		{13, Now().SubDays(10), "1 week ago"},

		{14, Now().AddHours(1), "1 hour from now"},
		{15, Now().SubHours(1), "1 hour ago"},
		{16, Now().AddHours(10), "10 hours from now"},
		{17, Now().SubHours(10), "10 hours ago"},

		{18, Now().AddMinutes(1), "1 minute from now"},
		{19, Now().SubMinutes(1), "1 minute ago"},
		{20, Now().AddMinutes(10), "10 minutes from now"},
		{21, Now().SubMinutes(10), "10 minutes ago"},

		{22, Now().AddSeconds(1), "1 second from now"},
		{23, Now().SubSeconds(1), "1 second ago"},
		{24, Now().AddSeconds(10), "10 seconds from now"},
		{25, Now().SubSeconds(10), "10 seconds ago"},
	}

	for _, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DiffForHumans(), "Current test id is "+strconv.Itoa(test.id))
	}
}
