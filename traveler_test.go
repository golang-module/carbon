package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_AddDuration(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		duration string // 输入参数
		expected string // 期望值
	}{
		{1, "", "10h", ""},
		{2, "0", "10h", ""},
		{3, "0000-00-00", "10h", ""},
		{4, "00:00:00", "10h", ""},
		{5, "0000-00-00 00:00:00", "10h", ""},

		{6, "2020-01-01 13:14:15", "10h", "2020-01-01 23:14:15"},
		{7, "2020-01-01 13:14:15", "10.5h", "2020-01-01 23:44:15"},

		{8, "2020-01-01 13:14:15", "10m", "2020-01-01 13:24:15"},
		{9, "2020-01-01 13:14:15", "10.5m", "2020-01-01 13:24:45"},

		{10, "2020-01-01 13:14:15", "10s", "2020-01-01 13:14:25"},
		{11, "2020-01-01 13:14:15", "10.5s", "2020-01-01 13:14:25"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddDuration(test.duration)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestError_AddDuration(t *testing.T) {
	duration := "10x"
	c := Parse("2020-08-05").AddDuration(duration)
	assert.Equal(t, invalidDurationError(duration), c.Error, "Should catch an exception in AddDuration()")
}

func TestCarbon_SubDuration(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		duration string // 输入参数
		expected string // 期望值
	}{
		{1, "", "10h", ""},
		{2, "0", "10h", ""},
		{3, "0000-00-00", "10h", ""},
		{4, "00:00:00", "10h", ""},
		{5, "0000-00-00 00:00:00", "10h", ""},

		{6, "2020-01-01 13:14:15", "10h", "2020-01-01 03:14:15"},
		{7, "2020-01-01 13:14:15", "10.5h", "2020-01-01 02:44:15"},

		{8, "2020-01-01 13:14:15", "10m", "2020-01-01 13:04:15"},
		{9, "2020-01-01 13:14:15", "10.5m", "2020-01-01 13:03:45"},

		{10, "2020-01-01 13:14:15", "10s", "2020-01-01 13:14:05"},
		{11, "2020-01-01 13:14:15", "10.5s", "2020-01-01 13:14:04"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubDuration(test.duration)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddCenturies(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id        int    // 测试id
		input     string // 输入值
		centuries int    // 输入参数
		expected  string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01", 3, "2320-01-01"},
		{7, "2020-01-31", 3, "2320-01-31"},
		{8, "2020-02-01", 3, "2320-02-01"},
		{9, "2020-02-28", 3, "2320-02-28"},
		{10, "2020-02-29", 3, "2320-02-29"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddCenturies(test.centuries)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddCenturiesNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id        int    // 测试id
		input     string // 输入值
		centuries int    // 输入参数
		expected  string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01", 3, "2320-01-01"},
		{7, "2020-01-31", 3, "2320-01-31"},
		{8, "2020-02-01", 3, "2320-02-01"},
		{9, "2020-02-28", 3, "2320-02-28"},
		{10, "2020-02-29", 3, "2320-02-29"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddCenturiesNoOverflow(test.centuries)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubCenturies(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id        int    // 测试id
		input     string // 输入值
		centuries int
		expected  string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01", 3, "1720-01-01"},
		{7, "2020-01-31", 3, "1720-01-31"},
		{8, "2020-02-01", 3, "1720-02-01"},
		{9, "2020-02-28", 3, "1720-02-28"},
		{10, "2020-02-29", 3, "1720-02-29"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubCenturies(test.centuries)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubCenturiesNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id        int    // 测试id
		input     string // 输入值
		centuries int
		expected  string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01", 3, "1720-01-01"},
		{7, "2020-01-31", 3, "1720-01-31"},
		{8, "2020-02-01", 3, "1720-02-01"},
		{9, "2020-02-28", 3, "1720-02-28"},
		{10, "2020-02-29", 3, "1720-02-29"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubCenturiesNoOverflow(test.centuries)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddCentury(t *testing.T) {
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

		{6, "2020-01-01", "2120-01-01"},
		{7, "2020-01-31", "2120-01-31"},
		{8, "2020-02-01", "2120-02-01"},
		{9, "2020-02-28", "2120-02-28"},
		{10, "2020-02-29", "2120-02-29"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddCentury()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddCenturyNoOverflow(t *testing.T) {
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

		{6, "2020-01-01", "2120-01-01"},
		{7, "2020-01-31", "2120-01-31"},
		{8, "2020-02-01", "2120-02-01"},
		{9, "2020-02-28", "2120-02-28"},
		{10, "2020-02-29", "2120-02-29"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddCenturyNoOverflow()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubCentury(t *testing.T) {
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

		{6, "2020-01-01", "1920-01-01"},
		{7, "2020-01-31", "1920-01-31"},
		{8, "2020-02-01", "1920-02-01"},
		{9, "2020-02-28", "1920-02-28"},
		{10, "2020-02-29", "1920-02-29"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubCentury()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubCenturyNoOverflow(t *testing.T) {
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

		{6, "2020-01-01", "1920-01-01"},
		{7, "2020-01-31", "1920-01-31"},
		{8, "2020-02-01", "1920-02-01"},
		{9, "2020-02-28", "1920-02-28"},
		{10, "2020-02-29", "1920-02-29"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubCenturyNoOverflow()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddYears(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		years    int    // 输入参数
		expected string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01", 3, "2023-01-01"},
		{7, "2020-01-31", 3, "2023-01-31"},
		{8, "2020-02-01", 3, "2023-02-01"},
		{9, "2020-02-28", 3, "2023-02-28"},
		{10, "2020-02-29", 3, "2023-03-01"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddYears(test.years)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddYearsNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		years    int    // 输入参数
		expected string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01", 3, "2023-01-01"},
		{7, "2020-01-31", 3, "2023-01-31"},
		{8, "2020-02-01", 3, "2023-02-01"},
		{9, "2020-02-28", 3, "2023-02-28"},
		{10, "2020-02-29", 3, "2023-02-28"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddYearsNoOverflow(test.years)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubYears(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		years    int
		expected string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01", 3, "2017-01-01"},
		{7, "2020-01-31", 3, "2017-01-31"},
		{8, "2020-02-01", 3, "2017-02-01"},
		{9, "2020-02-28", 3, "2017-02-28"},
		{10, "2020-02-29", 3, "2017-03-01"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubYears(test.years)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubYearsNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		years    int
		expected string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01", 3, "2017-01-01"},
		{7, "2020-01-31", 3, "2017-01-31"},
		{8, "2020-02-01", 3, "2017-02-01"},
		{9, "2020-02-28", 3, "2017-02-28"},
		{10, "2020-02-29", 3, "2017-02-28"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubYearsNoOverflow(test.years)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddYear(t *testing.T) {
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

		{6, "2020-01-01", "2021-01-01"},
		{7, "2020-01-31", "2021-01-31"},
		{8, "2020-02-01", "2021-02-01"},
		{9, "2020-02-28", "2021-02-28"},
		{10, "2020-02-29", "2021-03-01"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddYear()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddYearNoOverflow(t *testing.T) {
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

		{6, "2020-01-01", "2021-01-01"},
		{7, "2020-01-31", "2021-01-31"},
		{8, "2020-02-01", "2021-02-01"},
		{9, "2020-02-28", "2021-02-28"},
		{10, "2020-02-29", "2021-02-28"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddYearNoOverflow()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubYear(t *testing.T) {
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

		{6, "2020-01-01", "2019-01-01"},
		{7, "2020-01-31", "2019-01-31"},
		{8, "2020-02-01", "2019-02-01"},
		{9, "2020-02-28", "2019-02-28"},
		{10, "2020-02-29", "2019-03-01"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubYear()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubYearNoOverflow(t *testing.T) {
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

		{6, "2020-01-01", "2019-01-01"},
		{7, "2020-01-31", "2019-01-31"},
		{8, "2020-02-01", "2019-02-01"},
		{9, "2020-02-28", "2019-02-28"},
		{10, "2020-02-29", "2019-02-28"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubYearNoOverflow()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddQuarters(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		quarters int
		expected string // 期望值
	}{
		{1, "", 2, ""},
		{2, "0", 2, ""},
		{3, "0000-00-00", 2, ""},
		{4, "00:00:00", 2, ""},
		{5, "0000-00-00 00:00:00", 2, ""},

		{6, "2019-08-01", 2, "2020-02-01"},
		{7, "2019-08-31", 2, "2020-03-02"},
		{8, "2020-01-01", 2, "2020-07-01"},
		{9, "2020-02-28", 2, "2020-08-28"},
		{10, "2020-02-29", 2, "2020-08-29"},
		{11, "2020-08-31", 2, "2021-03-03"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddQuarters(test.quarters)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddQuartersNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		quarters int
		expected string // 期望值
	}{
		{1, "", 2, ""},
		{2, "0", 2, ""},
		{3, "0000-00-00", 2, ""},
		{4, "00:00:00", 2, ""},
		{5, "0000-00-00 00:00:00", 2, ""},

		{6, "2019-08-01", 2, "2020-02-01"},
		{7, "2019-08-31", 2, "2020-02-29"},
		{8, "2020-01-01", 2, "2020-07-01"},
		{9, "2020-02-28", 2, "2020-08-28"},
		{10, "2020-02-29", 2, "2020-08-29"},
		{11, "2020-08-31", 2, "2021-02-28"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddQuartersNoOverflow(test.quarters)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubQuarters(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		quarters int
		expected string // 期望值
	}{
		{1, "", 2, ""},
		{2, "0", 2, ""},
		{3, "0000-00-00", 2, ""},
		{4, "00:00:00", 2, ""},
		{5, "0000-00-00 00:00:00", 2, ""},

		{6, "2019-08-01", 2, "2019-02-01"},
		{7, "2019-08-31", 2, "2019-03-03"},
		{8, "2020-01-01", 2, "2019-07-01"},
		{9, "2020-02-28", 2, "2019-08-28"},
		{10, "2020-02-29", 2, "2019-08-29"},
		{11, "2020-08-31", 2, "2020-03-02"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubQuarters(test.quarters)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubQuartersNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		quarters int
		expected string // 期望值
	}{
		{1, "", 2, ""},
		{2, "0", 2, ""},
		{3, "0000-00-00", 2, ""},
		{4, "00:00:00", 2, ""},
		{5, "0000-00-00 00:00:00", 2, ""},

		{6, "2019-08-01", 2, "2019-02-01"},
		{7, "2019-08-31", 2, "2019-02-28"},
		{8, "2020-01-01", 2, "2019-07-01"},
		{9, "2020-02-28", 2, "2019-08-28"},
		{10, "2020-02-29", 2, "2019-08-29"},
		{11, "2020-08-31", 2, "2020-02-29"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubQuartersNoOverflow(test.quarters)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddQuarter(t *testing.T) {
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

		{6, "2019-11-01", "2020-02-01"},
		{7, "2019-11-30", "2020-03-01"},
		{8, "2020-02-28", "2020-05-28"},
		{9, "2020-02-29", "2020-05-29"},
		{10, "2020-08-31", "2020-12-01"},
		{11, "2020-11-01", "2021-02-01"},
		{12, "2020-11-30", "2021-03-02"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddQuarter()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddQuarterNoOverflow(t *testing.T) {
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

		{6, "2019-11-01", "2020-02-01"},
		{7, "2019-11-30", "2020-02-29"},
		{8, "2020-02-28", "2020-05-28"},
		{9, "2020-02-29", "2020-05-29"},
		{10, "2020-08-31", "2020-11-30"},
		{11, "2020-11-01", "2021-02-01"},
		{12, "2020-11-30", "2021-02-28"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddQuarterNoOverflow()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubQuarter(t *testing.T) {
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

		{6, "2019-04-01", "2019-01-01"},
		{7, "2019-04-30", "2019-01-30"},
		{8, "2020-05-01", "2020-02-01"},
		{9, "2020-05-31", "2020-03-02"},
		{10, "2020-04-01", "2020-01-01"},
		{11, "2020-04-30", "2020-01-30"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubQuarter()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubQuarterNoOverflow(t *testing.T) {
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

		{6, "2019-04-01", "2019-01-01"},
		{7, "2019-04-30", "2019-01-30"},
		{8, "2020-05-01", "2020-02-01"},
		{9, "2020-05-31", "2020-02-29"},
		{10, "2020-04-01", "2020-01-01"},
		{11, "2020-04-30", "2020-01-30"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubQuarterNoOverflow()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddMonths(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		months   int
		expected string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01", 3, "2020-04-01"},
		{7, "2020-01-31", 3, "2020-05-01"},
		{8, "2020-02-01", 3, "2020-05-01"},
		{9, "2020-02-28", 3, "2020-05-28"},
		{10, "2020-02-29", 3, "2020-05-29"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddMonths(test.months)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddMonthsNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		months   int
		expected string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01", 3, "2020-04-01"},
		{7, "2020-01-31", 3, "2020-04-30"},
		{8, "2020-02-01", 3, "2020-05-01"},
		{9, "2020-02-28", 3, "2020-05-28"},
		{10, "2020-02-29", 3, "2020-05-29"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddMonthsNoOverflow(test.months)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubMonths(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		months   int
		expected string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01", 3, "2019-10-01"},
		{7, "2020-01-31", 3, "2019-10-31"},
		{8, "2020-02-01", 3, "2019-11-01"},
		{9, "2020-02-28", 3, "2019-11-28"},
		{10, "2020-02-29", 3, "2019-11-29"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubMonths(test.months)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubMonthsNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		months   int
		expected string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01", 3, "2019-10-01"},
		{7, "2020-01-31", 3, "2019-10-31"},
		{8, "2020-02-01", 3, "2019-11-01"},
		{9, "2020-02-28", 3, "2019-11-28"},
		{10, "2020-02-29", 3, "2019-11-29"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubMonthsNoOverflow(test.months)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddMonth(t *testing.T) {
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

		{6, "2020-01-01", "2020-02-01"},
		{7, "2020-01-31", "2020-03-02"},
		{8, "2020-02-01", "2020-03-01"},
		{9, "2020-02-28", "2020-03-28"},
		{10, "2020-02-29", "2020-03-29"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddMonth()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddMonthNoOverflow(t *testing.T) {
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

		{6, "2020-01-01", "2020-02-01"},
		{7, "2020-01-31", "2020-02-29"},
		{8, "2020-02-01", "2020-03-01"},
		{9, "2020-02-28", "2020-03-28"},
		{10, "2020-02-29", "2020-03-29"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddMonthNoOverflow()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubMonth(t *testing.T) {
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

		{6, "2020-01-01", "2019-12-01"},
		{7, "2020-01-31", "2019-12-31"},
		{8, "2020-02-01", "2020-01-01"},
		{9, "2020-02-28", "2020-01-28"},
		{10, "2020-02-29", "2020-01-29"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubMonth()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubMonthNoOverflow(t *testing.T) {
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

		{6, "2020-01-01", "2019-12-01"},
		{7, "2020-01-31", "2019-12-31"},
		{8, "2020-02-01", "2020-01-01"},
		{9, "2020-02-28", "2020-01-28"},
		{10, "2020-02-29", "2020-01-29"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubMonthNoOverflow()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddWeeks(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		weeks    int
		expected string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01", 3, "2020-01-22"},
		{7, "2020-01-31", 3, "2020-02-21"},
		{8, "2020-02-01", 3, "2020-02-22"},
		{9, "2020-02-28", 3, "2020-03-20"},
		{10, "2020-02-29", 3, "2020-03-21"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddWeeks(test.weeks)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubWeeks(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		weeks    int
		expected string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01", 3, "2019-12-11"},
		{7, "2020-01-31", 3, "2020-01-10"},
		{8, "2020-02-01", 3, "2020-01-11"},
		{9, "2020-02-28", 3, "2020-02-07"},
		{10, "2020-02-29", 3, "2020-02-08"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubWeeks(test.weeks)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddWeek(t *testing.T) {
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

		{6, "2020-01-01", "2020-01-08"},
		{7, "2020-01-31", "2020-02-07"},
		{8, "2020-02-01", "2020-02-08"},
		{9, "2020-02-28", "2020-03-06"},
		{10, "2020-02-29", "2020-03-07"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddWeek()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubWeek(t *testing.T) {
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

		{6, "2020-01-01", "2019-12-25"},
		{7, "2020-01-31", "2020-01-24"},
		{8, "2020-02-01", "2020-01-25"},
		{9, "2020-02-28", "2020-02-21"},
		{10, "2020-02-29", "2020-02-22"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubWeek()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddDays(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		days     int
		expected string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01", 3, "2020-01-04"},
		{7, "2020-01-31", 3, "2020-02-03"},
		{8, "2020-02-01", 3, "2020-02-04"},
		{9, "2020-02-28", 3, "2020-03-02"},
		{10, "2020-02-29", 3, "2020-03-03"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddDays(test.days)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubDays(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		days     int
		expected string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01", 3, "2019-12-29"},
		{7, "2020-01-31", 3, "2020-01-28"},
		{8, "2020-02-01", 3, "2020-01-29"},
		{9, "2020-02-28", 3, "2020-02-25"},
		{10, "2020-02-29", 3, "2020-02-26"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubDays(test.days)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddDay(t *testing.T) {
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

		{6, "2020-01-01", "2020-01-02"},
		{7, "2020-01-31", "2020-02-01"},
		{8, "2020-02-01", "2020-02-02"},
		{9, "2020-02-28", "2020-02-29"},
		{10, "2020-02-29", "2020-03-01"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddDay()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubDay(t *testing.T) {
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

		{6, "2020-01-01", "2019-12-31"},
		{7, "2020-01-31", "2020-01-30"},
		{8, "2020-02-01", "2020-01-31"},
		{9, "2020-02-28", "2020-02-27"},
		{10, "2020-02-29", "2020-02-28"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubDay()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddHours(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		hours    int
		expected string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01 13:14:15", 3, "2020-01-01 16:14:15"},
		{7, "2020-01-31 13:14:15", 3, "2020-01-31 16:14:15"},
		{8, "2020-02-01 13:14:15", 3, "2020-02-01 16:14:15"},
		{9, "2020-02-28 13:14:15", 3, "2020-02-28 16:14:15"},
		{10, "2020-02-29 13:14:15", 3, "2020-02-29 16:14:15"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddHours(test.hours)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubHours(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		hours    int
		expected string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01 13:14:15", 3, "2020-01-01 10:14:15"},
		{7, "2020-01-31 13:14:15", 3, "2020-01-31 10:14:15"},
		{8, "2020-02-01 13:14:15", 3, "2020-02-01 10:14:15"},
		{9, "2020-02-28 13:14:15", 3, "2020-02-28 10:14:15"},
		{10, "2020-02-29 13:14:15", 3, "2020-02-29 10:14:15"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubHours(test.hours)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddHour(t *testing.T) {
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

		{6, "2020-01-01 13:14:15", "2020-01-01 14:14:15"},
		{7, "2020-01-31 13:14:15", "2020-01-31 14:14:15"},
		{8, "2020-02-01 13:14:15", "2020-02-01 14:14:15"},
		{9, "2020-02-28 13:14:15", "2020-02-28 14:14:15"},
		{10, "2020-02-29 13:14:15", "2020-02-29 14:14:15"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddHour()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubHour(t *testing.T) {
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

		{6, "2020-01-01 13:14:15", "2020-01-01 12:14:15"},
		{7, "2020-01-31 13:14:15", "2020-01-31 12:14:15"},
		{8, "2020-02-01 13:14:15", "2020-02-01 12:14:15"},
		{9, "2020-02-28 13:14:15", "2020-02-28 12:14:15"},
		{10, "2020-02-29 13:14:15", "2020-02-29 12:14:15"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubHour()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddMinutes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		minutes  int
		expected string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01 13:14:15", 3, "2020-01-01 13:17:15"},
		{7, "2020-01-31 13:14:15", 3, "2020-01-31 13:17:15"},
		{8, "2020-02-01 13:14:15", 3, "2020-02-01 13:17:15"},
		{9, "2020-02-28 13:14:15", 3, "2020-02-28 13:17:15"},
		{10, "2020-02-29 13:14:15", 3, "2020-02-29 13:17:15"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddMinutes(test.minutes)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubMinutes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		minutes  int
		expected string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01 13:14:15", 3, "2020-01-01 13:11:15"},
		{7, "2020-01-31 13:14:15", 3, "2020-01-31 13:11:15"},
		{8, "2020-02-01 13:14:15", 3, "2020-02-01 13:11:15"},
		{9, "2020-02-28 13:14:15", 3, "2020-02-28 13:11:15"},
		{10, "2020-02-29 13:14:15", 3, "2020-02-29 13:11:15"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubMinutes(test.minutes)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddMinute(t *testing.T) {
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

		{6, "2020-01-01 13:14:15", "2020-01-01 13:15:15"},
		{7, "2020-01-31 13:14:15", "2020-01-31 13:15:15"},
		{8, "2020-02-01 13:14:15", "2020-02-01 13:15:15"},
		{9, "2020-02-28 13:14:15", "2020-02-28 13:15:15"},
		{10, "2020-02-29 13:14:15", "2020-02-29 13:15:15"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddMinute()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubMinute(t *testing.T) {
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

		{6, "2020-01-01 13:14:15", "2020-01-01 13:13:15"},
		{7, "2020-01-31 13:14:15", "2020-01-31 13:13:15"},
		{8, "2020-02-01 13:14:15", "2020-02-01 13:13:15"},
		{9, "2020-02-28 13:14:15", "2020-02-28 13:13:15"},
		{10, "2020-02-29 13:14:15", "2020-02-29 13:13:15"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubMinute()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddSeconds(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		seconds  int
		expected string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01 13:14:15", 3, "2020-01-01 13:14:18"},
		{7, "2020-01-31 13:14:15", 3, "2020-01-31 13:14:18"},
		{8, "2020-02-01 13:14:15", 3, "2020-02-01 13:14:18"},
		{9, "2020-02-28 13:14:15", 3, "2020-02-28 13:14:18"},
		{10, "2020-02-29 13:14:15", 3, "2020-02-29 13:14:18"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddSeconds(test.seconds)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubSeconds(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		seconds  int
		expected string // 期望值
	}{
		{1, "", 3, ""},
		{2, "0", 3, ""},
		{3, "0000-00-00", 3, ""},
		{4, "00:00:00", 3, ""},
		{5, "0000-00-00 00:00:00", 3, ""},

		{6, "2020-01-01 13:14:15", 3, "2020-01-01 13:14:12"},
		{7, "2020-01-31 13:14:15", 3, "2020-01-31 13:14:12"},
		{8, "2020-02-01 13:14:15", 3, "2020-02-01 13:14:12"},
		{9, "2020-02-28 13:14:15", 3, "2020-02-28 13:14:12"},
		{10, "2020-02-29 13:14:15", 3, "2020-02-29 13:14:12"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubSeconds(test.seconds)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_AddSecond(t *testing.T) {
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

		{6, "2020-01-01 13:14:15", "2020-01-01 13:14:16"},
		{7, "2020-01-31 13:14:15", "2020-01-31 13:14:16"},
		{8, "2020-02-01 13:14:15", "2020-02-01 13:14:16"},
		{9, "2020-02-28 13:14:15", "2020-02-28 13:14:16"},
		{10, "2020-02-29 13:14:15", "2020-02-29 13:14:16"},
	}

	for _, test := range tests {
		c := Parse(test.input).AddSecond()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_SubSecond(t *testing.T) {
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

		{6, "2020-01-01 13:14:15", "2020-01-01 13:14:14"},
		{7, "2020-01-31 13:14:15", "2020-01-31 13:14:14"},
		{8, "2020-02-01 13:14:15", "2020-02-01 13:14:14"},
		{9, "2020-02-28 13:14:15", "2020-02-28 13:14:14"},
		{10, "2020-02-29 13:14:15", "2020-02-29 13:14:14"},
	}

	for _, test := range tests {
		c := Parse(test.input).SubSecond()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}
}
