package carbon

import (
	"testing"
)

func TestCarbon_ToString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15 +0800 CST"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToTimestamp(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int64  // 期望输出值
	}{
		{"2020-01-01 13:14:15", 1577855655},
		{"2020-01-31 13:14:15", 1580447655},
		{"2020-02-01 13:14:15", 1580534055},
		{"2020-02-28 13:14:15", 1582866855},
		{"2020-02-29 13:14:15", 1582953255},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToTimestamp()

		if output != v.output {
			t.Fatalf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_Format(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		format string // 输入参数
		output string // 期望输出值
	}{
		{"", "Y年m月d日", ""},
		{"0000-00-00 00:00:00", "Y年m月d日", ""},
		{"0000-00-00", "Y年m月d日", ""},
		{"00:00:00", "Y年m月d日", ""},
		{"2020-08-05 13:14:15", "Y年m月d日", "2020年08月05日"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Format(v.format)

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToFormatString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		format string // 输入参数
		output string // 期望输出值
	}{
		{"", "Y年m月d日", ""},
		{"0000-00-00 00:00:00", "Y年m月d日", ""},
		{"0000-00-00", "Y年m月d日", ""},
		{"00:00:00", "Y年m月d日", ""},
		{"2020-08-05 13:14:15", "Y年m月d日", "2020年08月05日"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToFormatString(v.format)

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToDayDateTimeString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0000-00-00 00:00:00", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"2020-08-05 13:14:15", "Wed, Aug 5, 2020 1:14 PM"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToDayDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToDateTimeString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0000-00-00 00:00:00", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15"},
		{"2020-08-05", "2020-08-05 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToDateString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0000-00-00 00:00:00", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"2020-08-05 13:14:15", "2020-08-05"},
		{"2020-08-05", "2020-08-05"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToTimeString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "13:14:15"},
		{"2020-08-05", "00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToAtomString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "2020-08-05T13:14:15+08:00"},
		{"2020-08-05", "2020-08-05T00:00:00+08:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToAtomString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToAnsicString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wed Aug  5 13:14:15 2020"},
		{"2020-08-05", "Wed Aug  5 00:00:00 2020"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToAnsicString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToCookieString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wednesday, 05-Aug-2020 13:14:15 CST"},
		{"2020-08-05", "Wednesday, 05-Aug-2020 00:00:00 CST"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToCookieString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRssString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 13:14:15 +0800"},
		{"2020-08-05", "Wed, 05 Aug 2020 00:00:00 +0800"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRssString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToW3cString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "2020-08-05T13:14:15+08:00"},
		{"2020-08-05", "2020-08-05T00:00:00+08:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToW3cString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToUnixDateString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wed Aug  5 13:14:15 CST 2020"},
		{"2020-08-05", "Wed Aug  5 00:00:00 CST 2020"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToUnixDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRubyDateString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wed Aug 05 13:14:15 +0800 2020"},
		{"2020-08-05", "Wed Aug 05 00:00:00 +0800 2020"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRubyDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToKitchenString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "1:14PM"},
		{"2020-08-05", "12:00AM"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToKitchenString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRFC822String(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "05 Aug 20 13:14 CST"},
		{"2020-08-05", "05 Aug 20 00:00 CST"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRFC822String()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRFC822zString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "05 Aug 20 13:14 +0800"},
		{"2020-08-05", "05 Aug 20 00:00 +0800"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRFC822zString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRFC850String(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wednesday, 05-Aug-20 13:14:15 CST"},
		{"2020-08-05", "Wednesday, 05-Aug-20 00:00:00 CST"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRFC850String()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRFC1036String(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wed, 05 Aug 20 13:14:15 +0800"},
		{"2020-08-05", "Wed, 05 Aug 20 00:00:00 +0800"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRFC1036String()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRFC1123String(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 13:14:15 CST"},
		{"2020-08-05", "Wed, 05 Aug 2020 00:00:00 CST"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRFC1123String()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRFC1123ZString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 13:14:15 +0800"},
		{"2020-08-05", "Wed, 05 Aug 2020 00:00:00 +0800"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRFC1123ZString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRFC2822String(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 13:14:15 +0800"},
		{"2020-08-05", "Wed, 05 Aug 2020 00:00:00 +0800"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRFC2822String()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRFC3339String(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "2020-08-05T13:14:15+08:00"},
		{"2020-08-05", "2020-08-05T00:00:00+08:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRFC3339String()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToRFC7231String(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 13:14:15 GMT"},
		{"2020-08-05", "Wed, 05 Aug 2020 00:00:00 GMT"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToRFC7231String()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_DaysInYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"0000-00-00 00:00:00", 0},
		{"2020-08-05", 366},
		{"2021-08-05", 365},
	}

	for _, v := range Tests {
		output := Parse(v.input).DaysInYear()

		if output != v.output {
			t.Fatalf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_DaysInMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"0000-00-00 00:00:00", 0},
		{"2020-01-05", 31},
		{"2020-02-05", 29},
		{"2020-03-05", 31},
		{"2020-04-05", 30},
		{"2020-05-05", 31},
		{"2021-06-05", 30},
		{"2021-07-05", 31},
		{"2021-08-05", 31},
		{"2021-09-05", 30},
		{"2021-10-05", 31},
		{"2021-11-05", 30},
		{"2021-12-05", 31},
	}

	for _, v := range Tests {
		output := Parse(v.input).DaysInMonth()

		if output != v.output {
			t.Fatalf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_MonthOfYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"0000-00-00 00:00:00", 0},
		{"2020-01-05", 1},
		{"2020-02-05", 2},
		{"2020-03-05", 3},
		{"2020-04-05", 4},
		{"2020-05-05", 5},
		{"2021-06-05", 6},
		{"2021-07-05", 7},
		{"2021-08-05", 8},
		{"2021-09-05", 9},
		{"2021-10-05", 10},
		{"2021-11-05", 11},
		{"2021-12-05", 12},
	}

	for _, v := range Tests {
		output := Parse(v.input).MonthOfYear()

		if output != v.output {
			t.Fatalf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_DayOfYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"0000-00-00", 0},
		{"2020-01-01", 1},
		{"2020-01-31", 31},
		{"2020-08-05", 218},
	}

	for _, v := range Tests {
		output := Parse(v.input).DayOfYear()

		if output != v.output {
			t.Fatalf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_DayOfMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"0000-00-00", 0},
		{"2020-01-01", 1},
		{"2020-01-31", 31},
		{"2020-08-05", 5},
	}

	for _, v := range Tests {
		output := Parse(v.input).DayOfMonth()

		if output != v.output {
			t.Fatalf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_DayOfWeek(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"0000-00-00", 0},
		{"2020-01-01", 3},
		{"2020-01-31", 5},
		{"2020-02-28", 5},
		{"2020-01-29", 3},
		{"2020-08-05", 3},
	}

	for _, v := range Tests {
		output := Parse(v.input).DayOfWeek()

		if output != v.output {
			t.Fatalf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_WeekOfYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"0000-00-00", 0},
		{"2020-01-01", 1},
		{"2020-01-31", 5},
		{"2020-02-28", 9},
		{"2020-01-29", 5},
		{"2020-08-05", 32},
	}

	for _, v := range Tests {
		output := Parse(v.input).WeekOfYear()

		if output != v.output {
			t.Fatalf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_WeekOfMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"0000-00-00", 0},
		{"2020-01-01", 1},
		{"2020-01-31", 4},
		{"2020-02-28", 1},
		{"2020-01-29", 2},
		{"2020-08-05", 1},
	}

	for _, v := range Tests {
		output := Parse(v.input).WeekOfMonth()

		if output != v.output {
			t.Fatalf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_IsZero(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"0000-00-00 00:00:00", true},
		{"2020-08-05 00:00:00", false},
		{"2020-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsZero()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsNow(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"0000-00-00 00:00:00", false},
		{Tomorrow().ToDateTimeString(), false},
		{Now().ToDateTimeString(), true},
		{Yesterday().ToDateTimeString(), false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsNow()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsFuture(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"0000-00-00 00:00:00", false},
		{Tomorrow().ToDateTimeString(), true},
		{Now().ToDateTimeString(), false},
		{Yesterday().ToDateTimeString(), false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsFuture()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsPast(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"0000-00-00 00:00:00", true},
		{Tomorrow().ToDateTimeString(), false},
		{Now().ToDateTimeString(), false},
		{Yesterday().ToDateTimeString(), true},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsPast()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsLeapYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2016-01-01", true},
		{"2017-01-01", false},
		{"2018-01-01", false},
		{"2019-01-01", false},
		{"2020-01-01", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsLeapYear()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsJanuary(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-01-01", true},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", false},
		{"2020-07-01", false},
		{"2020-08-01", false},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsJanuary()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsFebruary(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-01-01", false},
		{"2020-02-01", true},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", false},
		{"2020-07-01", false},
		{"2020-08-01", false},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsFebruary()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsMarch(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", true},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", false},
		{"2020-07-01", false},
		{"2020-08-01", false},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsMarch()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsApril(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", true},
		{"2020-05-01", false},
		{"2020-06-01", false},
		{"2020-07-01", false},
		{"2020-08-01", false},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsApril()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsMay(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", true},
		{"2020-06-01", false},
		{"2020-07-01", false},
		{"2020-08-01", false},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsMay()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsJune(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", true},
		{"2020-07-01", false},
		{"2020-08-01", false},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsJune()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsJuly(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", false},
		{"2020-07-01", true},
		{"2020-08-01", false},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsJuly()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsAugust(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", false},
		{"2020-07-01", false},
		{"2020-08-01", true},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsAugust()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsSeptember(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", false},
		{"2020-07-01", false},
		{"2020-08-01", false},
		{"2020-09-01", true},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsSeptember()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsOctober(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", false},
		{"2020-07-01", false},
		{"2020-08-01", false},
		{"2020-09-01", false},
		{"2020-10-01", true},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsOctober()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsNovember(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", false},
		{"2020-07-01", false},
		{"2020-08-01", false},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", true},
		{"2020-12-01", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsNovember()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsDecember(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", false},
		{"2020-07-01", false},
		{"2020-08-01", false},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsDecember()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsMonday(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-10-05", true},
		{"2020-10-06", false},
		{"2020-10-07", false},
		{"2020-10-08", false},
		{"2020-10-09", false},
		{"2020-10-10", false},
		{"2020-10-11", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsMonday()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsTuesday(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-10-05", false},
		{"2020-10-06", true},
		{"2020-10-07", false},
		{"2020-10-08", false},
		{"2020-10-09", false},
		{"2020-10-10", false},
		{"2020-10-11", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsTuesday()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsWednesday(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-10-05", false},
		{"2020-10-06", false},
		{"2020-10-07", true},
		{"2020-10-08", false},
		{"2020-10-09", false},
		{"2020-10-10", false},
		{"2020-10-11", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsWednesday()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsThursday(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-10-05", false},
		{"2020-10-06", false},
		{"2020-10-07", false},
		{"2020-10-08", true},
		{"2020-10-09", false},
		{"2020-10-10", false},
		{"2020-10-11", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsThursday()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsFriday(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-10-05", false},
		{"2020-10-06", false},
		{"2020-10-07", false},
		{"2020-10-08", false},
		{"2020-10-09", true},
		{"2020-10-10", false},
		{"2020-10-11", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsFriday()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsSaturday(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-10-05", false},
		{"2020-10-06", false},
		{"2020-10-07", false},
		{"2020-10-08", false},
		{"2020-10-09", false},
		{"2020-10-10", true},
		{"2020-10-11", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsSaturday()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsSunday(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-10-05", false},
		{"2020-10-06", false},
		{"2020-10-07", false},
		{"2020-10-08", false},
		{"2020-10-09", false},
		{"2020-10-10", false},
		{"2020-10-11", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsSunday()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsWeekday(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-10-05", true},
		{"2020-10-06", true},
		{"2020-10-07", true},
		{"2020-10-08", true},
		{"2020-10-09", true},
		{"2020-10-10", false},
		{"2020-10-11", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsWeekday()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsWeekend(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-10-05", false},
		{"2020-10-06", false},
		{"2020-10-07", false},
		{"2020-10-08", false},
		{"2020-10-09", false},
		{"2020-10-10", true},
		{"2020-10-11", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsWeekend()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsYesterday(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{Now().ToDateTimeString(), false},
		{Yesterday().ToDateTimeString(), true},
		{Tomorrow().ToDateTimeString(), false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsYesterday()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsToday(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{Now().ToDateTimeString(), true},
		{Yesterday().ToDateTimeString(), false},
		{Tomorrow().ToDateTimeString(), false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsToday()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsTomorrow(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{Now().ToDateTimeString(), false},
		{Yesterday().ToDateTimeString(), false},
		{Tomorrow().ToDateTimeString(), true},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsTomorrow()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}
