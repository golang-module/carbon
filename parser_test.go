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
		output string // 期望输出值
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
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(c.ToDateTimeStringWithTimezone(PRC), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_ParseByFormat(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		format string // 输入参数
		output string // 期望输出值
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
		c := ParseByFormat(test.input, test.format)
		assert.Nil(c.Error)
		assert.Equal(c.ToDateTimeString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_ParseByLayout(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id     int    // 测试id
		input  string // 输入值
		format string // 输入参数
		output string // 期望输出值
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
		c := ParseByLayout(test.input, test.format)
		assert.Nil(c.Error)
		assert.Equal(c.ToDateTimeString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}
