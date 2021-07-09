package carbon

import (
	"testing"
	"time"
)

func TestCarbon_StartOfCentury(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"00-00-00 00:00:00", ""},
		{"2020-01-01 00:00:00", "2000-01-01 00:00:00"},
		{"2020-01-31 23:59:59", "2000-01-01 00:00:00"},
		{"2020-02-01 13:14:15", "2000-01-01 00:00:00"},
		{"2020-02-28", "2000-01-01 00:00:00"},
		{"2020-02-29", "2000-01-01 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).StartOfCentury().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfCentury(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"00-00-00 00:00:00", ""},
		{"2020-01-01 00:00:00", "2999-12-31 23:59:59"},
		{"2020-01-31 23:59:59", "2999-12-31 23:59:59"},
		{"2020-02-01 13:14:15", "2999-12-31 23:59:59"},
		{"2020-02-28", "2999-12-31 23:59:59"},
		{"2020-02-29", "2999-12-31 23:59:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfCentury().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_StartOfDecade(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"00-00-00 00:00:00", ""},
		{"2020-01-31 23:59:59", "2020-01-01 00:00:00"},
		{"2021-01-01 00:00:00", "2020-01-01 00:00:00"},
		{"2029-01-31 23:59:59", "2020-01-01 00:00:00"},
		{"2020-02-01 13:14:15", "2020-01-01 00:00:00"},
		{"2020-02-28", "2020-01-01 00:00:00"},
		{"2020-02-29", "2020-01-01 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).StartOfDecade().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfDecade(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"00-00-00 00:00:00", ""},
		{"2020-01-31 23:59:59", "2029-12-31 23:59:59"},
		{"2021-01-01 00:00:00", "2029-12-31 23:59:59"},
		{"2029-01-31 23:59:59", "2029-12-31 23:59:59"},
		{"2020-02-01 13:14:15", "2029-12-31 23:59:59"},
		{"2020-02-28", "2029-12-31 23:59:59"},
		{"2020-02-29", "2029-12-31 23:59:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfDecade().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_StartOfYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"00-00-00 00:00:00", ""},
		{"2020-01-01 00:00:00", "2020-01-01 00:00:00"},
		{"2020-01-31 23:59:59", "2020-01-01 00:00:00"},
		{"2020-02-01 13:14:15", "2020-01-01 00:00:00"},
		{"2020-02-28", "2020-01-01 00:00:00"},
		{"2020-02-29", "2020-01-01 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).StartOfYear().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"00-00-00 00:00:00", ""},
		{"2020-01-01 00:00:00", "2020-12-31 23:59:59"},
		{"2020-01-31 23:59:59", "2020-12-31 23:59:59"},
		{"2020-02-01 13:14:15", "2020-12-31 23:59:59"},
		{"2020-02-28", "2020-12-31 23:59:59"},
		{"2020-02-29", "2020-12-31 23:59:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfYear().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_StartOfQuarter(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"00-00-00 00:00:00", ""},
		{"2020-01-31 00:00:00", "2020-01-01 00:00:00"},
		{"2020-02-28 00:00:00", "2020-01-01 00:00:00"},
		{"2020-04-30 23:59:59", "2020-04-01 00:00:00"},
		{"2020-07-31 23:59:59", "2020-07-01 00:00:00"},
		{"2020-08-05 13:14:15", "2020-07-01 00:00:00"},
		{"2020-10-31", "2020-10-01 00:00:00"},
		{"2020-12-31", "2020-10-01 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).StartOfQuarter().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfQuarter(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"00-00-00 00:00:00", ""},
		{"2020-01-31 00:00:00", "2020-03-31 23:59:59"},
		{"2020-02-28 00:00:00", "2020-03-31 23:59:59"},
		{"2020-04-28 23:59:59", "2020-06-30 23:59:59"},
		{"2020-07-31 23:59:59", "2020-09-30 23:59:59"},
		{"2020-08-05 13:14:15", "2020-09-30 23:59:59"},
		{"2020-10-31", "2020-12-31 23:59:59"},
		{"2020-12-31", "2020-12-31 23:59:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfQuarter().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_StartOfMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"00-00-00 00:00:00", ""},
		{"2020-01-01 00:00:00", "2020-01-01 00:00:00"},
		{"2020-01-31 23:59:59", "2020-01-01 00:00:00"},
		{"2020-02-01 13:14:15", "2020-02-01 00:00:00"},
		{"2020-02-28", "2020-02-01 00:00:00"},
		{"2020-02-29", "2020-02-01 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).StartOfMonth().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"00-00-00 00:00:00", ""},
		{"2020-01-01 00:00:00", "2020-01-31 23:59:59"},
		{"2020-01-31 23:59:59", "2020-01-31 23:59:59"},
		{"2020-02-01 13:14:15", "2020-02-29 23:59:59"},
		{"2020-02-28", "2020-02-29 23:59:59"},
		{"2020-02-29", "2020-02-29 23:59:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfMonth().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_StartOfWeek(t *testing.T) {
	Tests := []struct {
		input  string       // 输入值
		week   time.Weekday // 输入参数
		output string       // 期望输出值
	}{
		{"", time.Sunday, ""},
		{"00-00-00 00:00:00", time.Sunday, ""},
		{"", time.Monday, ""},
		{"00-00-00 00:00:00", time.Monday, ""},
		{"2021-06-13", time.Sunday, "2021-06-13 00:00:00"},
		{"2021-06-14", time.Sunday, "2021-06-13 00:00:00"},
		{"2021-06-18", time.Sunday, "2021-06-13 00:00:00"},
		{"2021-06-13", time.Monday, "2021-06-07 00:00:00"},
		{"2021-06-14", time.Monday, "2021-06-14 00:00:00"},
		{"2021-06-18", time.Monday, "2021-06-14 00:00:00"},
		{"2021-06-19", time.Monday, "2021-06-14 00:00:00"},
		{"2021-06-20", time.Monday, "2021-06-14 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).StartOfWeek(v.week).ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfWeek(t *testing.T) {
	Tests := []struct {
		input  string       // 输入值
		week   time.Weekday // 输入参数
		output string       // 期望输出值
	}{
		{"", time.Sunday, ""},
		{"00-00-00 00:00:00", time.Sunday, ""},
		{"", time.Monday, ""},
		{"00-00-00 00:00:00", time.Monday, ""},
		{"2021-06-13", time.Sunday, "2021-06-19 23:59:59"},
		{"2021-06-14", time.Sunday, "2021-06-19 23:59:59"},
		{"2021-06-18", time.Sunday, "2021-06-19 23:59:59"},
		{"2021-06-13", time.Monday, "2021-06-13 23:59:59"},
		{"2021-06-14", time.Monday, "2021-06-20 23:59:59"},
		{"2021-06-18", time.Monday, "2021-06-20 23:59:59"},
		{"2021-06-19", time.Monday, "2021-06-20 23:59:59"},
		{"2021-06-20", time.Monday, "2021-06-20 23:59:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfWeek(v.week).ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_StartOfDay(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"00-00-00 00:00:00", ""},
		{"2020-01-01 00:00:00", "2020-01-01 00:00:00"},
		{"2020-01-31 23:59:59", "2020-01-31 00:00:00"},
		{"2020-02-01 13:14:15", "2020-02-01 00:00:00"},
		{"2020-02-28", "2020-02-28 00:00:00"},
		{"2020-02-29", "2020-02-29 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).StartOfDay().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfDay(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"00-00-00 00:00:00", ""},
		{"2020-01-01 00:00:00", "2020-01-01 23:59:59"},
		{"2020-01-31 23:59:59", "2020-01-31 23:59:59"},
		{"2020-02-01 13:14:15", "2020-02-01 23:59:59"},
		{"2020-02-28", "2020-02-28 23:59:59"},
		{"2020-02-29", "2020-02-29 23:59:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfDay().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_StartOfHour(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"00-00-00 00:00:00", ""},
		{"2020-01-01 00:00:00", "2020-01-01 00:00:00"},
		{"2020-01-31 23:59:59", "2020-01-31 23:00:00"},
		{"2020-02-01 13:14:15", "2020-02-01 13:00:00"},
		{"2020-02-28", "2020-02-28 00:00:00"},
		{"2020-02-29", "2020-02-29 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).StartOfHour().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfHour(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"00-00-00 00:00:00", ""},
		{"2020-01-01 00:00:00", "2020-01-01 00:59:59"},
		{"2020-01-31 23:59:59", "2020-01-31 23:59:59"},
		{"2020-02-01 13:14:15", "2020-02-01 13:59:59"},
		{"2020-02-28", "2020-02-28 00:59:59"},
		{"2020-02-29", "2020-02-29 00:59:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfHour().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_StartOfMinute(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"00-00-00 00:00:00", ""},
		{"2020-01-01 00:00:00", "2020-01-01 00:00:00"},
		{"2020-01-31 23:59:59", "2020-01-31 23:59:00"},
		{"2020-02-01 13:14:15", "2020-02-01 13:14:00"},
		{"2020-02-28", "2020-02-28 00:00:00"},
		{"2020-02-29", "2020-02-29 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).StartOfMinute().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfMinute(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"00-00-00 00:00:00", ""},
		{"2020-01-01 00:00:00", "2020-01-01 00:00:59"},
		{"2020-01-31 23:59:59", "2020-01-31 23:59:59"},
		{"2020-02-01 13:14:15", "2020-02-01 13:14:59"},
		{"2020-02-28", "2020-02-28 00:00:59"},
		{"2020-02-29", "2020-02-29 00:00:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfMinute().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_StartOfSecond(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"00-00-00 00:00:00", ""},
		{"2020-01-01 00:00:00.123", "2020-01-01 00:00:00.0"},
		{"2020-01-31 23:59:59.123", "2020-01-31 23:59:59.0"},
		{"2020-02-01 13:14:15.123", "2020-02-01 13:14:15.0"},
		{"2020-02-28", "2020-02-28 00:00:00.0"},
		{"2020-02-29", "2020-02-29 00:00:00.0"},
	}

	for _, v := range Tests {
		output := Parse(v.input).StartOfSecond().Format("Y-m-d H:i:s.u")

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfSecond(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"00-00-00 00:00:00", ""},
		{"2020-01-01 00:00:00.123", "2020-01-01 00:00:00.999"},
		{"2020-01-31 23:59:59.123", "2020-01-31 23:59:59.999"},
		{"2020-02-01 13:14:15.123", "2020-02-01 13:14:15.999"},
		{"2020-02-28", "2020-02-28 00:00:00.999"},
		{"2020-02-29", "2020-02-29 00:00:00.999"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfSecond().Format("Y-m-d H:i:s.u")

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}
