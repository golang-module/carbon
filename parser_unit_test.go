package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_Parse(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		0: {"", ""},
		1: {"0", ""},
		2: {"0000-00-00", ""},
		3: {"00:00:00", ""},
		4: {"0000-00-00 00:00:00", ""},

		5: {Parse("now").ToDateTimeString(), Now().ToDateTimeString() + " +0800 CST"},
		6: {Parse("yesterday").ToDateTimeString(), Yesterday().ToDateTimeString() + " +0800 CST"},
		7: {Parse("tomorrow").ToDateTimeString(), Tomorrow().ToDateTimeString() + " +0800 CST"},

		8:  {"2020-8-5", "2020-08-05 00:00:00 +0800 CST"},
		9:  {"2020-8-05", "2020-08-05 00:00:00 +0800 CST"},
		10: {"2020-08-05", "2020-08-05 00:00:00 +0800 CST"},
		11: {"2020-8-5 1:2:3", "2020-08-05 01:02:03 +0800 CST"},
		12: {"2020-08-05 1:2:03", "2020-08-05 01:02:03 +0800 CST"},
		13: {"2020-08-05 1:02:03", "2020-08-05 01:02:03 +0800 CST"},
		14: {"2020-08-05 01:02:03", "2020-08-05 01:02:03 +0800 CST"},

		15: {"2020-8-5 13:14:15PM EST", "2020-08-05 21:14:15 +0800 CST"},
		16: {"2020-8-05 13:14:15PM EST", "2020-08-05 21:14:15 +0800 CST"},
		17: {"2020-08-05 13:14:15PM EST", "2020-08-05 21:14:15 +0800 CST"},
		18: {"2020-08-05 13:14:15.999999999PM EST", "2020-08-05 21:14:15.999999999 +0800 CST"},

		19: {"2020-8-5 13:14:15 PM EST", "2020-08-05 21:14:15 +0800 CST"},
		20: {"2020-8-05 13:14:15 PM EST", "2020-08-05 21:14:15 +0800 CST"},
		21: {"2020-08-05 13:14:15 PM EST", "2020-08-05 21:14:15 +0800 CST"},
		22: {"2020-08-05 13:14:15.999999999 PM EST", "2020-08-05 21:14:15.999999999 +0800 CST"},

		23: {"2020-8-5 13:14:15 PM EST", "2020-08-05 21:14:15 +0800 CST"},
		24: {"2020-8-05 13:14:15 PM EST", "2020-08-05 21:14:15 +0800 CST"},
		25: {"2020-08-05 13:14:15 PM EST", "2020-08-05 21:14:15 +0800 CST"},
		26: {"2020-08-05 13:14:15.999999999 PM EST", "2020-08-05 21:14:15.999999999 +0800 CST"},

		27: {"2020-08-05 1:2:3.999", "2020-08-05 01:02:03.999 +0800 CST"},
		28: {"2020-08-05 1:2:3.999999", "2020-08-05 01:02:03.999999 +0800 CST"},
		29: {"2020-08-05 1:2:3.999999999", "2020-08-05 01:02:03.999999999 +0800 CST"},

		30: {"2020-08-05 1:2:3.999 +0800 CST", "2020-08-05 01:02:03.999 +0800 CST"},
		31: {"2020-08-05 1:2:3.999999 +0800 CST", "2020-08-05 01:02:03.999999 +0800 CST"},
		32: {"2020-08-05 1:2:3.999999999 +0800 CST", "2020-08-05 01:02:03.999999999 +0800 CST"},

		33: {"2020-8-5T13:14:15Z", "2020-08-05 21:14:15 +0800 CST"},
		34: {"2020-8-5T13:14:15.999Z", "2020-08-05 21:14:15.999 +0800 CST"},
		35: {"2020-8-5T13:14:15.999999Z", "2020-08-05 21:14:15.999999 +0800 CST"},
		36: {"2020-8-5T13:14:15.999999999Z", "2020-08-05 21:14:15.999999999 +0800 CST"},

		37: {"2020.8.5", "2020-08-05 00:00:00 +0800 CST"},
		38: {"2020.8.05", "2020-08-05 00:00:00 +0800 CST"},
		39: {"2020.08.05", "2020-08-05 00:00:00 +0800 CST"},
		40: {"2020.8.5 13:14:15", "2020-08-05 13:14:15 +0800 CST"},
		41: {"2020.8.05 13:14:15", "2020-08-05 13:14:15 +0800 CST"},
		42: {"2020.08.05 13:14:15", "2020-08-05 13:14:15 +0800 CST"},

		43: {"2020/8/5", "2020-08-05 00:00:00 +0800 CST"},
		44: {"2020/8/05", "2020-08-05 00:00:00 +0800 CST"},
		45: {"2020/08/05", "2020-08-05 00:00:00 +0800 CST"},
		46: {"2020/8/5 13:14:15", "2020-08-05 13:14:15 +0800 CST"},
		47: {"2020/8/05 13:14:15", "2020-08-05 13:14:15 +0800 CST"},
		48: {"2020/08/05 13:14:15", "2020-08-05 13:14:15 +0800 CST"},
		49: {"2020/8/5 13:14:15.999", "2020-08-05 13:14:15.999 +0800 CST"},
		50: {"2020/8/05 13:14:15.999999", "2020-08-05 13:14:15.999999 +0800 CST"},
		51: {"2020/08/5 13:14:15.999999999", "2020-08-05 13:14:15.999999999 +0800 CST"},

		52: {"2020-8-5T13:14:15+08:00", "2020-08-05 13:14:15 +0800 CST"},
		53: {"2020-8-05T13:14:15+08:00", "2020-08-05 13:14:15 +0800 CST"},
		54: {"2020-08-05T13:14:15+08:00", "2020-08-05 13:14:15 +0800 CST"},
		55: {"2020-08-05T13:14:15.999+08:00", "2020-08-05 13:14:15.999 +0800 CST"},
		56: {"2020-08-05T13:14:15.999999+08:00", "2020-08-05 13:14:15.999999 +0800 CST"},
		57: {"2020-08-05T13:14:15.999999999+08:00", "2020-08-05 13:14:15.999999999 +0800 CST"},

		58: {"20200805", "2020-08-05 00:00:00 +0800 CST"},
		59: {"20200805131415", "2020-08-05 13:14:15 +0800 CST"},
		60: {"20200805131415+08:00", "2020-08-05 13:14:15 +0800 CST"},
		61: {"20200805131415.999", "2020-08-05 13:14:15.999 +0800 CST"},
		62: {"20200805131415.999999", "2020-08-05 13:14:15.999999 +0800 CST"},
		63: {"20200805131415.999999999", "2020-08-05 13:14:15.999999999 +0800 CST"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ParseByFormat(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		format   string
		expected string
	}{
		0: {"", "Y|m|d H:i:s", ""},
		1: {"0", "Y|m|d H:i:s", ""},
		2: {"0000-00-00", "Y|m|d H:i:s", ""},
		3: {"00:00:00", "Y|m|d H:i:s", ""},
		4: {"0000-00-00 00:00:00", "Y|m|d H:i:s", ""},

		5: {"2020|08|05 13:14:15", "Y|m|d H:i:s", "2020-08-05 13:14:15"},
		6: {"It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s", "2020-08-05 13:14:15"},
		7: {"今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒", "2020-08-05 13:14:15"},
		8: {"上次上报时间:2020-08-05 13:14:15，请每日按时打卡", "上次上报时间:Y-m-d H:i:s，请每日按时打卡", "2020-08-05 13:14:15"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).ParseByFormat(test.input, test.format)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := ParseByFormat(test.input, test.format, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ParseByLayout(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		layout   string
		expected string
	}{
		0: {"", "2006-01-02", ""},
		1: {"0", "2006-01-02", ""},
		2: {"0000-00-00", "2006-01-02", ""},
		3: {"00:00:00", "15:04:05", ""},
		4: {"0000-00-00 00:00:00", "2006-01-02 15:04:05", ""},

		5: {"2020|08|05 13:14:15", "2006|01|02 15:04:05", "2020-08-05 13:14:15"},
		6: {"It is 2020|08|05 13:14:15", "It is 2006|01|02 15:04:05", "2020-08-05 13:14:15"},
		7: {"今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒", "2020-08-05 13:14:15"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).ParseByLayout(test.input, test.layout)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := ParseByLayout(test.input, test.layout, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

// https://github.com/golang-module/carbon/issues/202
func TestCarbon_Issue202(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		0: {"2023-01-08T09:02:48", "2023-01-08 09:02:48 +0800 CST"},
		1: {"2023-1-8T09:02:48", "2023-01-08 09:02:48 +0800 CST"},
		2: {"2023-01-08T9:2:48", "2023-01-08 09:02:48 +0800 CST"},
		3: {"2023-01-8T9:2:48", "2023-01-08 09:02:48 +0800 CST"},
		4: {"2023-1-8T9:2:48", "2023-01-08 09:02:48 +0800 CST"},
		5: {"2023-01-08T09:02:48.000+0000", "2023-01-08 17:02:48 +0800 CST"},
		6: {"2023-1-8T09:02:48.000+0000", "2023-01-08 17:02:48 +0800 CST"},
		7: {"2023-1-8T9:2:48.000+0000", "2023-01-08 17:02:48 +0800 CST"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_ParseByLayout(t *testing.T) {
	assert.NotNil(t, ParseByLayout("2020-08-05", "2006-03-04", "xxx").Error, "It should catch an exception in ParseByLayout")
	assert.NotNil(t, ParseByLayout("xxx", "2006-03-04", PRC).Error, "It should catch an exception in ParseByLayout")
}

func TestError_ParseByFormat(t *testing.T) {
	assert.NotNil(t, ParseByFormat("2020-08-05", "Y-m-d", "xxx").Error, "It should catch an exception in ParseByFormat()")
	assert.NotNil(t, ParseByFormat("xxx", "Y-m-d", PRC).Error, "It should catch an exception in ParseByFormat")
}

// https://github.com/golang-module/carbon/issues/206
func TestCarbon_Issue206(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		format   string
		expected string
	}{
		0: {"1699677240", "U", "2023-11-11 12:34:00 +0800 CST"},
		1: {"1699677240666", "V", "2023-11-11 12:34:00.666 +0800 CST"},
		2: {"1699677240666666", "X", "2023-11-11 12:34:00.666666 +0800 CST"},
		3: {"1699677240666666666", "Z", "2023-11-11 12:34:00.666666666 +0800 CST"},
	}

	for index, test := range tests {
		c := ParseByFormat(test.input, test.format, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}
