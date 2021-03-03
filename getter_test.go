package carbon

import (
	"testing"
)

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
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
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
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
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
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
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
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
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
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_DayOfWeek(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0000-00-00 00:00:00", 0},
		{"0000-00-00", 0},
		{"2020-08-03", 1},
		{"2020-08-04", 2},
		{"2020-08-05", 3},
		{"2020-08-06", 4},
		{"2020-08-07", 5},
		{"2020-08-08", 6},
		{"2020-08-09", 7},
	}

	for _, v := range Tests {
		output := Parse(v.input).DayOfWeek()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_WeekOfYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"0000-00-00", 0},
		{"2019-12-31", 1},
		{"2020-01-01", 1},
		{"2020-01-31", 5},
		{"2020-02-28", 9},
		{"2020-01-29", 5},
		{"2020-08-05", 32},
	}

	for _, v := range Tests {
		output := Parse(v.input).WeekOfYear()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
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
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_Century(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0000-00-00", 0},
		{"2020-08-05", 21},
	}

	for _, v := range Tests {
		output := Parse(v.input).Century()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_Year(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0000-00-00", 0},
		{"2020-08-05", 2020},
	}

	for _, v := range Tests {
		output := Parse(v.input).Year()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_Quarter(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0000-00-00", 0},
		{"2020-01-05", 1},
		{"2020-04-05", 2},
		{"2020-08-05", 3},
		{"2020-11-05", 4},
		{"2020-11-33", 0},
	}

	for _, v := range Tests {
		output := Parse(v.input).Quarter()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_Month(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0000-00-00 00:00:00", 0},
		{"0000-00-00", 0},
		{"2020-08-05 13:14:15", 8},
		{"2020-08-05", 8},
	}

	for _, v := range Tests {
		output := Parse(v.input).Month()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_Week(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", -1},
		{"0000-00-00 00:00:00", -1},
		{"0000-00-00", -1},
		{"2020-08-03", 1},
		{"2020-08-04", 2},
		{"2020-08-05", 3},
		{"2020-08-06", 4},
		{"2020-08-07", 5},
		{"2020-08-08", 6},
		{"2020-08-09", 0},
	}

	for _, v := range Tests {
		output := Parse(v.input).Week()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_Day(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0000-00-00 00:00:00", 0},
		{"0000-00-00", 0},
		{"2020-08-05 13:14:15", 5},
		{"2020-08-05", 5},
	}

	for _, v := range Tests {
		output := Parse(v.input).Day()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_Hour(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0000-00-00 00:00:00", 0},
		{"0000-00-00", 0},
		{"2020-08-05 13:14:15", 13},
		{"2020-08-05", 0},
	}

	for _, v := range Tests {
		output := Parse(v.input).Hour()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_Minute(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0000-00-00 00:00:00", 0},
		{"0000-00-00", 0},
		{"2020-08-05 13:14:15", 14},
		{"2020-08-05", 0},
	}

	for _, v := range Tests {
		output := Parse(v.input).Minute()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_Second(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0000-00-00 00:00:00", 0},
		{"0000-00-00", 0},
		{"2020-08-05 13:14:15", 15},
		{"2020-08-05", 0},
	}

	for _, v := range Tests {
		output := Parse(v.input).Second()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_Millisecond(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0000-00-00 00:00:00", 0},
		{"0000-00-00", 0},
		{"2020-08-05 13:14:15", 0},
		{"2020-08-05", 0},
	}

	for _, v := range Tests {
		output := Parse(v.input).Millisecond()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_Microsecond(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0000-00-00 00:00:00", 0},
		{"0000-00-00", 0},
		{"2020-08-05 13:14:15", 0},
		{"2020-08-05", 0},
	}

	for _, v := range Tests {
		output := Parse(v.input).Microsecond()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_Nanosecond(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0000-00-00 00:00:00", 0},
		{"0000-00-00", 0},
		{"2020-08-05 13:14:15", 0},
		{"2020-08-05", 0},
	}

	for _, v := range Tests {
		output := Parse(v.input).Nanosecond()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}

func TestCarbon_Timezone(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{PRC, PRC},
		{Tokyo, Tokyo},
	}

	for _, v := range Tests {
		output := SetTimezone(v.input).Timezone()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_Locale(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"en", "en"},
		{"zh-CN", "zh-CN"},
	}

	for _, v := range Tests {
		output := Now().SetLocale(v.input).Locale()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_Age(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{Now().AddYears(18).ToDateTimeString(), 0},
		{Now().SubYears(18).ToDateTimeString(), 18},
	}

	for _, v := range Tests {
		output := Parse(v.input).Age()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d", v.input, v.output, output)
		}
	}
}
