package carbon

import (
	"testing"
)

func TestCarbon_IsZero(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"0000-00-00 00:00:00", true},
		{"0000-00-00", true},
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsLongYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2015-01-01", true},
		{"2016-01-01", false},
		{"2017-01-01", false},
		{"2018-01-01", false},
		{"2019-01-01", false},
		{"2020-01-01", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsLongYear()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
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
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_Compare(t *testing.T) {
	now := Now()
	tomorrow := Tomorrow()
	yesterday := Yesterday()
	Tests := []struct {
		input    Carbon // 输入值
		operator string // 输入参数
		time     Carbon // 输入参数
		output   bool   // 期望输出值
	}{
		{now, ">", yesterday, true},
		{now, "<", yesterday, false},
		{now, "<", tomorrow, true},
		{now, ">", tomorrow, false},
		{now, "=", now, true},
		{now, ">=", now, true},
		{now, "<=", now, true},
		{now, "!=", now, false},
		{now, "<>", now, false},
		{now, "!=", yesterday, true},
		{now, "<>", yesterday, true},
		{now, "+", yesterday, false},
	}

	for _, v := range Tests {
		output := v.input.Compare(v.operator, v.time)

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Errorf("Input %s %s %s, expected %s, but got %s\n", v.input.ToDateTimeString(), v.operator, v.time.ToDateTimeString(), expected, reality)
		}
	}
}

func TestCarbon_Gt(t *testing.T) {
	now := Now()
	tomorrow := Tomorrow()
	yesterday := Yesterday()
	Tests := []struct {
		input  Carbon // 输入值
		time   Carbon // 输入参数
		output bool   // 期望输出值
	}{
		{now, now, false},
		{now, yesterday, true},
		{now, tomorrow, false},
	}

	for _, v := range Tests {
		output := v.input.Gt(v.time)

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Errorf("Input %s > %s, expected %s, but got %s\n", v.input.ToDateTimeString(), v.time.ToDateTimeString(), expected, reality)
		}
	}
}

func TestCarbon_Lt(t *testing.T) {
	now := Now()
	tomorrow := Tomorrow()
	yesterday := Yesterday()
	Tests := []struct {
		input  Carbon // 输入值
		time   Carbon // 输入参数
		output bool   // 期望输出值
	}{
		{now, now, false},
		{now, yesterday, false},
		{now, tomorrow, true},
	}

	for _, v := range Tests {
		output := v.input.Lt(v.time)

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Errorf("Input %s < %s, expected %s, but got %s\n", v.input.ToDateTimeString(), v.time.ToDateTimeString(), expected, reality)
		}
	}
}

func TestCarbon_Eq(t *testing.T) {
	now := Now()
	tomorrow := Tomorrow()
	yesterday := Yesterday()
	Tests := []struct {
		input  Carbon // 输入值
		time   Carbon // 输入参数
		output bool   // 期望输出值
	}{
		{now, now, true},
		{now, yesterday, false},
		{now, tomorrow, false},
	}

	for _, v := range Tests {
		output := v.input.Eq(v.time)

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Errorf("Input %s = %s, expected %s, but got %s\n", v.input.ToDateTimeString(), v.time.ToDateTimeString(), expected, reality)
		}
	}
}

func TestCarbon_Ne(t *testing.T) {
	now := Now()
	tomorrow := Tomorrow()
	yesterday := Yesterday()
	Tests := []struct {
		input  Carbon // 输入值
		time   Carbon // 输入参数
		output bool   // 期望输出值
	}{
		{now, now, false},
		{now, yesterday, true},
		{now, tomorrow, true},
	}

	for _, v := range Tests {
		output := v.input.Ne(v.time)

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Errorf("Input %s != %s, expected %s, but got %s\n", v.input.ToDateTimeString(), v.time.ToDateTimeString(), expected, reality)
		}
	}
}

func TestCarbon_Gte(t *testing.T) {
	now := Now()
	tomorrow := Tomorrow()
	yesterday := Yesterday()
	Tests := []struct {
		input  Carbon // 输入值
		time   Carbon // 输入参数
		output bool   // 期望输出值
	}{
		{now, now, true},
		{now, yesterday, true},
		{now, tomorrow, false},
	}

	for _, v := range Tests {
		output := v.input.Gte(v.time)

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Errorf("Input %s >= %s, expected %s, but got %s\n", v.input.ToDateTimeString(), v.time.ToDateTimeString(), expected, reality)
		}
	}
}

func TestCarbon_Lte(t *testing.T) {
	now := Now()
	tomorrow := Tomorrow()
	yesterday := Yesterday()
	Tests := []struct {
		input  Carbon // 输入值
		time   Carbon // 输入参数
		output bool   // 期望输出值
	}{
		{now, now, true},
		{now, yesterday, false},
		{now, tomorrow, true},
	}

	for _, v := range Tests {
		output := v.input.Lte(v.time)

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Errorf("Input %s <= %s, expected %s, but got %s\n", v.input.ToDateTimeString(), v.time.ToDateTimeString(), expected, reality)
		}
	}
}

func TestCarbon_Between(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		time1  Carbon // 输入参数
		time2  Carbon // 输入参数
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), false},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), Parse("2020-08-06 13:14:15"), false},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-04 13:14:15"), Parse("2020-08-05 13:14:15"), false},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-04 13:14:15"), Parse("2020-08-06 13:14:15"), true},
	}

	for _, v := range Tests {
		output := v.input.Between(v.time1, v.time2)

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Errorf("Input %s < %s < %s, expected %s, but got %s\n", v.time1.ToDateTimeString(), v.input.ToDateTimeString(), v.time2.ToDateTimeString(), expected, reality)
		}
	}
}

func TestCarbon_BetweenIncludedStartTime(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		time1  Carbon // 输入参数
		time2  Carbon // 输入参数
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), false},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), Parse("2020-08-06 13:14:15"), true},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-04 13:14:15"), Parse("2020-08-05 13:14:15"), false},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-04 13:14:15"), Parse("2020-08-06 13:14:15"), true},
	}

	for _, v := range Tests {
		output := v.input.BetweenIncludedStartTime(v.time1, v.time2)

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Errorf("Input %s <= %s < %s, expected %s, but got %s\n", v.time1.ToDateTimeString(), v.input.ToDateTimeString(), v.time2.ToDateTimeString(), expected, reality)
		}
	}
}

func TestCarbon_BetweenIncludedEndTime(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		time1  Carbon // 输入参数
		time2  Carbon // 输入参数
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), false},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), Parse("2020-08-06 13:14:15"), false},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-04 13:14:15"), Parse("2020-08-05 13:14:15"), true},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-04 13:14:15"), Parse("2020-08-06 13:14:15"), true},
	}

	for _, v := range Tests {
		output := v.input.BetweenIncludedEndTime(v.time1, v.time2)

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Errorf("Input %s < %s <= %s, expected %s, but got %s\n", v.time1.ToDateTimeString(), v.input.ToDateTimeString(), v.time2.ToDateTimeString(), expected, reality)
		}
	}
}

func TestCarbon_BetweenIncludedBoth(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		time1  Carbon // 输入参数
		time2  Carbon // 输入参数
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), true},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), Parse("2020-08-06 13:14:15"), true},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-04 13:14:15"), Parse("2020-08-05 13:14:15"), true},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-04 13:14:15"), Parse("2020-08-06 13:14:15"), true},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-06 13:14:15"), Parse("2020-08-06 13:14:15"), false},
	}

	for _, v := range Tests {
		output := v.input.BetweenIncludedBoth(v.time1, v.time2)

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Errorf("Input %s <= %s <= %s, expected %s, but got %s\n", v.time1.ToDateTimeString(), v.input.ToDateTimeString(), v.time2.ToDateTimeString(), expected, reality)
		}
	}
}
