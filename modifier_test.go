package carbon

import (
	"testing"
)

func TestCarbon_StartOfYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
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

func TestCarbon_StartOfMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
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
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2019-12-30 00:00:00"},
		{"2020-01-31 23:59:59", "2020-01-27 00:00:00"},
		{"2020-02-01 13:14:15", "2020-01-27 00:00:00"},
		{"2020-02-28", "2020-02-24 00:00:00"},
		{"2020-02-29", "2020-02-24 00:00:00"},
		{"2020-10-04", "2020-09-28 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).StartOfWeek().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfWeek(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-01-05 23:59:59"},
		{"2020-01-31 23:59:59", "2020-02-02 23:59:59"},
		{"2020-02-01 13:14:15", "2020-02-02 23:59:59"},
		{"2020-02-28", "2020-03-01 23:59:59"},
		{"2020-02-29", "2020-03-01 23:59:59"},
		{"2020-10-04", "2020-10-04 23:59:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfWeek().ToDateTimeString()

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
