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
		expected  string // 期望值
	}{
		{1, -1, "1970-01-01 07:59:59"},
		{2, 0, "1970-01-01 08:00:00"},
		{4, 1577855655, "2020-01-01 13:14:15"},
		{5, 1604074084682, "2020-10-31 00:08:04"},
		{6, 1604074196366540, "2020-10-31 00:09:56"},
		{7, 1604074298500312000, "2020-10-31 00:11:38"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).CreateFromTimestamp(test.timestamp)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}

	for _, test := range tests {
		c := CreateFromTimestamp(test.timestamp, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestError_CreateFromTimestamp(t *testing.T) {
	timestamp, timezone := int64(1577855655), "xxx"
	c1 := SetTimezone(timezone).CreateFromTimestamp(timestamp)
	assert.Equal(t, invalidTimezoneError(timezone), c1.Error, "Should catch an exception in CreateFromTimestamp()")

	c2 := CreateFromTimestamp(timestamp, timezone)
	assert.Equal(t, invalidTimezoneError(timezone), c2.Error, "Should catch an exception in CreateFromTimestamp()")
}

func TestCarbon_CreateFromDateTime(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id                                     int    // 测试id
		year, month, day, hour, minute, second int    // 输入参数
		expected                               string // 期望值
	}{
		{1, 2020, 1, 1, 13, 14, 15, "2020-01-01 13:14:15"},
		{2, 2020, 1, 31, 13, 14, 15, "2020-01-31 13:14:15"},
		{3, 2020, 2, 1, 13, 14, 15, "2020-02-01 13:14:15"},
		{4, 2020, 2, 28, 13, 14, 15, "2020-02-28 13:14:15"},
		{5, 2020, 2, 29, 13, 14, 15, "2020-02-29 13:14:15"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).CreateFromDateTime(test.year, test.month, test.day, test.hour, test.minute, test.second)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}

	for _, test := range tests {
		c := CreateFromDateTime(test.year, test.month, test.day, test.hour, test.minute, test.second, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestError_CreateFromDateTime(t *testing.T) {
	year, month, day, hour, minute, second, timezone := 2020, 1, 1, 13, 14, 15, "xxx"
	c1 := SetTimezone(timezone).CreateFromDateTime(year, month, day, hour, minute, second)
	assert.Equal(t, invalidTimezoneError(timezone), c1.Error, "Should catch an exception in CreateFromDateTime()")

	c2 := CreateFromDateTime(year, month, day, hour, minute, second, timezone)
	assert.Equal(t, invalidTimezoneError(timezone), c2.Error, "Should catch an exception in CreateFromDateTime()")
}

func TestCarbon_CreateFromDate(t *testing.T) {
	assert := assert.New(t)

	clock := Now(PRC).ToTimeString()
	tests := []struct {
		id               int    // 测试id
		year, month, day int    // 输入参数
		expected         string // 期望值
	}{
		{1, 2020, 1, 1, "2020-01-01 " + clock},
		{2, 2020, 1, 31, "2020-01-31 " + clock},
		{3, 2020, 2, 1, "2020-02-01 " + clock},
		{4, 2020, 2, 28, "2020-02-28 " + clock},
		{5, 2020, 2, 29, "2020-02-29 " + clock},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).CreateFromDate(test.year, test.month, test.day)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}

	for _, test := range tests {
		c := CreateFromDate(test.year, test.month, test.day, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestError_CreateFromDate(t *testing.T) {
	year, month, day, timezone := 13, 14, 15, "xxx"
	c1 := SetTimezone(timezone).CreateFromDate(year, month, day)
	assert.Equal(t, invalidTimezoneError(timezone), c1.Error, "Should catch an exception in CreateFromDate()")

	c2 := CreateFromDate(year, month, day, timezone)
	assert.Equal(t, invalidTimezoneError(timezone), c2.Error, "Should catch an exception in CreateFromDate()")
}

func TestCarbon_CreateFromTime(t *testing.T) {
	assert := assert.New(t)

	date := Now(PRC).ToDateString()
	tests := []struct {
		id                   int    // 测试id
		hour, minute, second int    // 输入参数
		expected             string // 期望值
	}{
		{1, 0, 0, 0, date + " 00:00:00"},
		{2, 00, 00, 15, date + " 00:00:15"},
		{3, 00, 14, 15, date + " 00:14:15"},
		{4, 13, 14, 15, date + " 13:14:15"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).CreateFromTime(test.hour, test.minute, test.second)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}

	for _, test := range tests {
		c := CreateFromTime(test.hour, test.minute, test.second, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestError_CreateFromTime(t *testing.T) {
	hour, minute, second, timezone := 13, 14, 15, "xxx"
	c1 := SetTimezone(timezone).CreateFromTime(hour, minute, second)
	assert.Equal(t, invalidTimezoneError(timezone), c1.Error, "Should catch an exception in CreateFromTime()")

	c2 := CreateFromTime(hour, minute, second, timezone)
	assert.Equal(t, invalidTimezoneError(timezone), c2.Error, "Should catch an exception in CreateFromTime()")
}
