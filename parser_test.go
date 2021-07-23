package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_Parse(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		output string // 期望值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},
		{5, "0000-00-00 00:00:00", ""},

		{6, "2020-08-05 13:14:15", "2020-08-05 13:14:15"},
		{7, "20200805131415", "2020-08-05 13:14:15"},
		{8, "20200805", "2020-08-05 00:00:00"},
		{9, "2020-08-05", "2020-08-05 00:00:00"},
		{10, "2020-08-05T13:14:15+08:00", "2020-08-05 13:14:15"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.output, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}

	for _, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.output, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestError_Parse(t *testing.T) {
	value, timezone := "2020-13-50", "xxx"
	c1 := Parse(value)
	assert.Equal(t, invalidValueError(value), c1.Error, "Should catch an exception in Parse()")

	c2 := Parse(value, timezone)
	assert.Equal(t, invalidTimezoneError(timezone), c2.Error, "Should catch an exception in Parse()")
}

func TestCarbon_ParseByFormat(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		format string // 输入参数
		output string // 期望值
	}{
		{1, "", "Y|m|d H:i:s", ""},
		{2, "0", "Y|m|d H:i:s", ""},
		{3, "0000-00-00", "Y|m|d H:i:s", ""},
		{4, "00:00:00", "Y|m|d H:i:s", ""},
		{5, "0000-00-00 00:00:00", "Y|m|d H:i:s", ""},

		{6, "2020|08|05 13:14:15", "Y|m|d H:i:s", "2020-08-05 13:14:15"},
		{7, "It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s", "2020-08-05 13:14:15"},
		{8, "今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒", "2020-08-05 13:14:15"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).ParseByFormat(test.input, test.format)
		assert.Nil(c.Error)
		assert.Equal(test.output, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}

	for _, test := range tests {
		c := ParseByFormat(test.input, test.format, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.output, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestError_ParseByFormat(t *testing.T) {
	value, format, timezone := "2020-08-50", "Y-m-d", "xxx"
	c1 := ParseByFormat(value, format)
	assert.Equal(t, invalidFormatError(value, format), c1.Error, "Should catch an exception in ParseByFormat()")

	c2 := ParseByFormat(value, format, timezone)
	assert.Equal(t, invalidTimezoneError(timezone), c2.Error, "Should catch an exception in ParseByFormat()")
}

func TestCarbon_ParseByLayout(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		layout string // 输入参数
		output string // 期望值
	}{
		{1, "", "2006-01-02", ""},
		{2, "0", "2006-01-02", ""},
		{3, "0000-00-00", "2006-01-02", ""},
		{4, "00:00:00", "15:04:05", ""},
		{5, "0000-00-00 00:00:00", "2006-01-02 15:04:05", ""},

		{6, "2020|08|05 13:14:15", "2006|01|02 15:04:05", "2020-08-05 13:14:15"},
		{7, "It is 2020|08|05 13:14:15", "It is 2006|01|02 15:04:05", "2020-08-05 13:14:15"},
		{8, "今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒", "2020-08-05 13:14:15"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).ParseByLayout(test.input, test.layout)
		assert.Nil(c.Error)
		assert.Equal(test.output, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}

	for _, test := range tests {
		c := ParseByLayout(test.input, test.layout, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.output, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestError_ParseByLayout(t *testing.T) {
	value, layout, timezone := "2020-08-50", "2006-01-02", "xxx"
	c1 := ParseByLayout(value, layout)
	assert.Equal(t, invalidLayoutError(value, layout), c1.Error, "Should catch an exception in ParseByLayout()")

	c2 := ParseByLayout(value, layout, timezone)
	assert.Equal(t, invalidTimezoneError(timezone), c2.Error, "Should catch an exception in ParseByLayout()")
}
