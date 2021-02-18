package carbon

import (
	"testing"
)

func TestCarbon_CreateFromTimestamp(t *testing.T) {
	Tests := []struct {
		timestamp int64  // 输入参数
		output    string // 期望输出值
	}{
		{123456, "1970-01-01 08:00:00"},
		{1577855655, "2020-01-01 13:14:15"},
		{1604074084682, "2020-10-31 00:08:04"},
		{1604074196366540, "2020-10-31 00:09:56"},
		{1604074298500312000, "2020-10-31 00:11:38"},
	}

	for _, v := range Tests {
		output := CreateFromTimestamp(v.timestamp).ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %d, expected %s, but got %s", v.timestamp, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).CreateFromTimestamp(v.timestamp).ToDateTimeString()

		if output != v.output {
			t.Errorf("Expected %s, but got %s", v.output, output)
		}
	}
}

func TestCarbon_CreateFromDateTime(t *testing.T) {
	Tests := []struct {
		year, month, day, hour, minute, second int    // 输入参数
		output                                 string // 期望输出值
	}{
		{2020, 01, 01, 13, 14, 15, "2020-01-01 13:14:15"},
		{2020, 1, 31, 13, 14, 15, "2020-01-31 13:14:15"},
		{2020, 2, 1, 13, 14, 15, "2020-02-01 13:14:15"},
		{2020, 2, 28, 13, 14, 15, "2020-02-28 13:14:15"},
		{2020, 2, 29, 13, 14, 15, "2020-02-29 13:14:15"},
	}

	for _, v := range Tests {
		output := CreateFromDateTime(v.year, v.month, v.day, v.hour, v.minute, v.second).ToDateTimeString()

		if output != v.output {
			t.Errorf("Expected %s, but got %s", v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).CreateFromDateTime(v.year, v.month, v.day, v.hour, v.minute, v.second).ToDateTimeString()

		if output != v.output {
			t.Errorf("Expected %s, but got %s", v.output, output)
		}
	}
}

func TestCarbon_CreateFromDate(t *testing.T) {
	clock := Now().ToTimeString()

	Tests := []struct {
		year, month, day int    // 输入参数
		output           string // 期望输出值
	}{
		{2020, 01, 01, "2020-01-01 " + clock},
		{2020, 1, 31, "2020-01-31 " + clock},
		{2020, 2, 1, "2020-02-01 " + clock},
		{2020, 2, 28, "2020-02-28 " + clock},
		{2020, 2, 29, "2020-02-29 " + clock},
	}

	for _, v := range Tests {
		output := CreateFromDate(v.year, v.month, v.day).ToDateTimeString()

		if output != v.output {
			t.Errorf("Expected %s, but got %s", v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).CreateFromDate(v.year, v.month, v.day).ToDateTimeString()

		if output != v.output {
			t.Errorf("Expected %s, but got %s", v.output, output)
		}
	}
}

func TestCarbon_CreateFromTime(t *testing.T) {
	date := Now().ToDateString()

	Tests := []struct {
		hour, minute, second int    // 输入参数
		output               string // 期望输出值
	}{
		{0, 0, 0, date + " 00:00:00"},
		{00, 00, 15, date + " 00:00:15"},
		{00, 14, 15, date + " 00:14:15"},
		{13, 14, 15, date + " 13:14:15"},
	}

	for _, v := range Tests {
		output := CreateFromTime(v.hour, v.minute, v.second).ToDateTimeString()

		if output != v.output {
			t.Errorf("Expected %s, but got %s", v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).CreateFromTime(v.hour, v.minute, v.second).ToDateTimeString()

		if output != v.output {
			t.Errorf("Expected %s, but got %s", v.output, output)
		}
	}
}
