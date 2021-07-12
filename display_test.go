package carbon

import (
	"testing"
)

func TestCarbon_ToTimestamp(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int64  // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},
		{"2020-01-01 13:14:15", 1577855655},
		{"2020-01-31 13:14:15", 1580447655},
		{"2020-02-01 13:14:15", 1580534055},
		{"2020-02-28 13:14:15", 1582866855},
		{"2020-02-29 13:14:15", 1582953255},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToTimestamp()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToTimestampWithSecond(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int64  // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},
		{"2020-01-01 13:14:15", 1577855655},
		{"2020-01-31 13:14:15", 1580447655},
		{"2020-02-01 13:14:15", 1580534055},
		{"2020-02-28 13:14:15", 1582866855},
		{"2020-02-29 13:14:15", 1582953255},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToTimestampWithSecond()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToTimestampWithMillisecond(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int64  // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},
		{"2020-01-01 13:14:15", 1577855655000},
		{"2020-01-31 13:14:15", 1580447655000},
		{"2020-02-01 13:14:15", 1580534055000},
		{"2020-02-28 13:14:15", 1582866855000},
		{"2020-02-29 13:14:15", 1582953255000},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToTimestampWithMillisecond()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToTimestampWithMicrosecond(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int64  // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},
		{"2020-01-01 13:14:15", 1577855655000000},
		{"2020-01-31 13:14:15", 1580447655000000},
		{"2020-02-01 13:14:15", 1580534055000000},
		{"2020-02-28 13:14:15", 1582866855000000},
		{"2020-02-29 13:14:15", 1582953255000000},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToTimestampWithMicrosecond()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToTimestampWithNanosecond(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int64  // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},
		{"2020-01-01 13:14:15", 1577855655000000000},
		{"2020-01-31 13:14:15", 1580447655000000000},
		{"2020-02-01 13:14:15", 1580534055000000000},
		{"2020-02-28 13:14:15", 1582866855000000000},
		{"2020-02-29 13:14:15", 1582953255000000000},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToTimestampWithNanosecond()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0000-00-00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15 +0800 CST"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToMonthString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		param  string // 参数值
		output string // 期望输出值
	}{
		{"", "en", ""},
		{"0000-00-00", "en", ""},
		{"2020-08-05", "en", "August"},
		{"2020-08-05", "zh-CN", "八月"},
		{"2020-08-05", "zh-Tw", "八月"},
		{"2020-08-05", "jp", "八月"},
		{"2020-08-05", "kr", "팔월"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SetLocale(v.param).ToMonthString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToShortMonthString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		param  string // 参数值
		output string // 期望输出值
	}{
		{"", "en", ""},
		{"0000-00-00", "en", ""},
		{"2020-08-05", "en", "Aug"},
		{"2020-08-05", "zh-CN", "8月"},
		{"2020-08-05", "zh-Tw", "8月"},
		{"2020-08-05", "jp", "8月"},
		{"2020-08-05", "kr", "8월"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SetLocale(v.param).ToShortMonthString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToWeekString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		param  string // 参数值
		output string // 期望输出值
	}{
		{"", "en", ""},
		{"0000-00-00", "en", ""},
		{"2020-08-05", "en", "Wednesday"},
		{"2020-08-05", "zh-CN", "星期三"},
		{"2020-08-05", "zh-Tw", "星期三"},
		{"2020-08-05", "jp", "水曜日"},
		{"2020-08-05", "kr", "수요일"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SetLocale(v.param).ToWeekString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToShortWeekString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		param  string // 参数值
		output string // 期望输出值
	}{
		{"", "en", ""},
		{"0000-00-00", "en", ""},
		{"2020-08-05", "en", "Wed"},
		{"2020-08-05", "zh-CN", "周三"},
		{"2020-08-05", "zh-Tw", "週三"},
		{"2020-08-05", "jp", "週三"},
		{"2020-08-05", "kr", "수요일"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SetLocale(v.param).ToShortWeekString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_Format(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		param  string // 参数值
		output string // 期望输出值
	}{
		{"", "Y年m月d日", ""},
		{"0", "Y年m月d日", ""},
		{"0000-00-00", "Y年m月d日", ""},
		{"00:00:00", "Y年m月d日", ""},
		{"0000-00-00 00:00:00", "Y年m月d日", ""},
		{"2020-08-05 13:14:15", "Y年m月d日", "2020年08月05日"},
		{"2020-08-05 01:14:15", "j", "5"},
		{"2020-08-05 01:14:15", "W", "32"},
		{"2020-08-05 01:14:15", "M", "Aug"},
		{"2020-08-05 01:14:15", "F", "August"},
		{"2020-08-05 01:14:15", "N", "3"},
		{"2020-08-05 01:14:15", "L", "1"},
		{"2020-08-05 01:14:15", "L", "1"},
		{"2021-08-05 01:14:15", "L", "0"},
		{"2020-08-05 01:14:15", "G", "1"},
		{"2020-08-05 13:14:15", "U", "1596604455"},
		{"2020-08-05 13:14:15.999", "u", "999"},
		{"2020-08-05 13:14:15", "w", "2"},
		{"2020-08-05 13:14:15", "t", "31"},
		{"2020-08-05 13:14:15", "z", "217"},
		{"2020-08-05 13:14:15", "e", "Local"},
		{"2020-08-05 13:14:15", "jS", "5th"},
		{"2020-08-22 13:14:15", "jS", "22nd"},
		{"2020-08-23 13:14:15", "jS", "23rd"},
		{"2020-08-31 13:14:15", "jS", "31st"},
		{"2020-08-05 13:14:15", "l jS \\o\\f F Y h:i:s A", "Wednesday 5th of August 2020 01:14:15 PM"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Format(v.param)

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToDayDateTimeString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wed, Aug 5, 2020 1:14 PM"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToDayDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToDateTimeString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15"},
		{"2020-08-05", "2020-08-05 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToDateString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "2020-08-05"},
		{"2020-08-05", "2020-08-05"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToTimeString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "13:14:15"},
		{"2020-08-05", "00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToAtomString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "2020-08-05T13:14:15+08:00"},
		{"2020-08-05", "2020-08-05T00:00:00+08:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToAtomString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToAnsicString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wed Aug  5 13:14:15 2020"},
		{"2020-08-05", "Wed Aug  5 00:00:00 2020"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToAnsicString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToCookieString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wednesday, 05-Aug-2020 13:14:15 CST"},
		{"2020-08-05", "Wednesday, 05-Aug-2020 00:00:00 CST"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToCookieString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRssString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 13:14:15 +0800"},
		{"2020-08-05", "Wed, 05 Aug 2020 00:00:00 +0800"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRssString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToW3cString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "2020-08-05T13:14:15+08:00"},
		{"2020-08-05", "2020-08-05T00:00:00+08:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToW3cString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToUnixDateString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wed Aug  5 13:14:15 CST 2020"},
		{"2020-08-05", "Wed Aug  5 00:00:00 CST 2020"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToUnixDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRubyDateString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wed Aug 05 13:14:15 +0800 2020"},
		{"2020-08-05", "Wed Aug 05 00:00:00 +0800 2020"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRubyDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToKitchenString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "1:14PM"},
		{"2020-08-05", "12:00AM"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToKitchenString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRfc822String(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "05 Aug 20 13:14 CST"},
		{"2020-08-05", "05 Aug 20 00:00 CST"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRfc822String()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRfc822zString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "05 Aug 20 13:14 +0800"},
		{"2020-08-05", "05 Aug 20 00:00 +0800"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRfc822zString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRfc850String(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wednesday, 05-Aug-20 13:14:15 CST"},
		{"2020-08-05", "Wednesday, 05-Aug-20 00:00:00 CST"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRfc850String()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRfc1036String(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wed, 05 Aug 20 13:14:15 +0800"},
		{"2020-08-05", "Wed, 05 Aug 20 00:00:00 +0800"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRfc1036String()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRfc1123String(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 13:14:15 CST"},
		{"2020-08-05", "Wed, 05 Aug 2020 00:00:00 CST"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRfc1123String()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRfc1123ZString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 13:14:15 +0800"},
		{"2020-08-05", "Wed, 05 Aug 2020 00:00:00 +0800"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRfc1123ZString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRfc2822String(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 13:14:15 +0800"},
		{"2020-08-05", "Wed, 05 Aug 2020 00:00:00 +0800"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRfc2822String()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRfc3339String(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "2020-08-05T13:14:15+08:00"},
		{"2020-08-05", "2020-08-05T00:00:00+08:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRfc3339String()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRfc7231String(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 13:14:15 GMT"},
		{"2020-08-05", "Wed, 05 Aug 2020 00:00:00 GMT"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRfc7231String()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}
