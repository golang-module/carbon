package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_CreateFromTimestamp(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id        int    // 测试id
		timestamp int64  // 输入参数
		output    string // 期望输出值
	}{
		{1, 0, ""},
		{2, 123456, "1970-01-01 08:00:00"},
		{3, 1577855655, "2020-01-01 13:14:15"},
		{4, 1604074084682, "2020-10-31 00:08:04"},
		{5, 1604074196366540, "2020-10-31 00:09:56"},
		{6, 1604074298500312000, "2020-10-31 00:11:38"},
	}

	for _, test := range tests {
		c := CreateFromTimestamp(test.timestamp)
		assert.Nil(c.Error)
		assert.Equal(c.ToDateTimeString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_CreateFromDateTime(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id                                     int    // 测试id
		year, month, day, hour, minute, second int    // 输入参数
		output                                 string // 期望输出值
	}{
		{1, 2020, 1, 1, 13, 14, 15, "2020-01-01 13:14:15"},
		{2, 2020, 1, 31, 13, 14, 15, "2020-01-31 13:14:15"},
		{3, 2020, 2, 1, 13, 14, 15, "2020-02-01 13:14:15"},
		{4, 2020, 2, 28, 13, 14, 15, "2020-02-28 13:14:15"},
		{5, 2020, 2, 29, 13, 14, 15, "2020-02-29 13:14:15"},
	}

	for _, test := range tests {
		c := CreateFromDateTime(test.year, test.month, test.day, test.hour, test.minute, test.second)
		assert.Nil(c.Error)
		assert.Equal(c.ToDateTimeString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_CreateFromDate(t *testing.T) {
	assert := assert.New(t)

	clock := Now().ToTimeString()
	tests := []struct {
		id               int    // 测试id
		year, month, day int    // 输入参数
		output           string // 期望输出值
	}{
		{1, 2020, 1, 1, "2020-01-01 " + clock},
		{2, 2020, 1, 31, "2020-01-31 " + clock},
		{3, 2020, 2, 1, "2020-02-01 " + clock},
		{4, 2020, 2, 28, "2020-02-28 " + clock},
		{5, 2020, 2, 29, "2020-02-29 " + clock},
	}

	for _, test := range tests {
		c := CreateFromDate(test.year, test.month, test.day)
		assert.Nil(c.Error)
		assert.Equal(c.ToDateTimeString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_CreateFromTime(t *testing.T) {
	assert := assert.New(t)

	date := Now().ToDateString()
	tests := []struct {
		id                   int    // 测试id
		hour, minute, second int    // 输入参数
		output               string // 期望输出值
	}{
		{1, 0, 0, 0, date + " 00:00:00"},
		{2, 00, 00, 15, date + " 00:00:15"},
		{3, 00, 14, 15, date + " 00:14:15"},
		{4, 13, 14, 15, date + " 13:14:15"},
	}

	for _, test := range tests {
		c := CreateFromTime(test.hour, test.minute, test.second)
		assert.Nil(c.Error)
		assert.Equal(c.ToDateTimeString(), test.output, "Current test id is "+strconv.Itoa(test.id))
	}
}
