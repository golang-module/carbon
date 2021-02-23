package carbon

import (
	"testing"
)

func TestCarbon_IsZero(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("0000-00-00 00:00:00"), true},
		{Parse("0000-00-00"), true},
		{Parse("2020-08-05"), false},
	}

	for _, v := range Tests {
		output := v.input.IsZero()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is zero, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is zero, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsNow(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("0000-00-00 00:00:00"), false},
		{Tomorrow(), false},
		{Now(), true},
		{Yesterday(), false},
	}

	for _, v := range Tests {
		output := v.input.IsNow()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is now, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is now, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsFuture(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("0000-00-00 00:00:00"), false},
		{Tomorrow(), true},
		{Now(), false},
		{Yesterday(), false},
	}

	for _, v := range Tests {
		output := v.input.IsFuture()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is future, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is future, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}

}

func TestCarbon_IsPast(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("0000-00-00 00:00:00"), true},
		{Tomorrow(), false},
		{Now(), false},
		{Yesterday(), true},
	}

	for _, v := range Tests {
		output := v.input.IsPast()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is past, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is past, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsLeapYear(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2016-01-01"), true},
		{Parse("2017-01-01"), false},
		{Parse("2018-01-01"), false},
		{Parse("2019-01-01"), false},
		{Parse("2020-01-01"), true},
	}

	for _, v := range Tests {
		output := v.input.IsLeapYear()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is leap year, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is leap year, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsLongYear(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2015-01-01"), true},
		{Parse("2016-01-01"), false},
		{Parse("2017-01-01"), false},
		{Parse("2018-01-01"), false},
		{Parse("2019-01-01"), false},
		{Parse("2020-01-01"), true},
	}

	for _, v := range Tests {
		output := v.input.IsLongYear()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is long year, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is long year, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsJanuary(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-01-01"), true},
		{Parse("2020-02-01"), false},
		{Parse("2020-03-01"), false},
		{Parse("2020-04-01"), false},
		{Parse("2020-05-01"), false},
		{Parse("2020-06-01"), false},
		{Parse("2020-07-01"), false},
		{Parse("2020-08-01"), false},
		{Parse("2020-09-01"), false},
		{Parse("2020-10-01"), false},
		{Parse("2020-11-01"), false},
		{Parse("2020-12-01"), false},
	}

	for _, v := range Tests {
		output := v.input.IsJanuary()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is January, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is January, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsFebruary(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-01-01"), false},
		{Parse("2020-02-01"), true},
		{Parse("2020-03-01"), false},
		{Parse("2020-04-01"), false},
		{Parse("2020-05-01"), false},
		{Parse("2020-06-01"), false},
		{Parse("2020-07-01"), false},
		{Parse("2020-08-01"), false},
		{Parse("2020-09-01"), false},
		{Parse("2020-10-01"), false},
		{Parse("2020-11-01"), false},
		{Parse("2020-12-01"), false},
	}

	for _, v := range Tests {
		output := v.input.IsFebruary()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is February, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is February, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsMarch(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-01-01"), false},
		{Parse("2020-02-01"), false},
		{Parse("2020-03-01"), true},
		{Parse("2020-04-01"), false},
		{Parse("2020-05-01"), false},
		{Parse("2020-06-01"), false},
		{Parse("2020-07-01"), false},
		{Parse("2020-08-01"), false},
		{Parse("2020-09-01"), false},
		{Parse("2020-10-01"), false},
		{Parse("2020-11-01"), false},
		{Parse("2020-12-01"), false},
	}

	for _, v := range Tests {
		output := v.input.IsMarch()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is March, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is March, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsApril(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-01-01"), false},
		{Parse("2020-02-01"), false},
		{Parse("2020-03-01"), false},
		{Parse("2020-04-01"), true},
		{Parse("2020-05-01"), false},
		{Parse("2020-06-01"), false},
		{Parse("2020-07-01"), false},
		{Parse("2020-08-01"), false},
		{Parse("2020-09-01"), false},
		{Parse("2020-10-01"), false},
		{Parse("2020-11-01"), false},
		{Parse("2020-12-01"), false},
	}

	for _, v := range Tests {
		output := v.input.IsApril()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is April, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is April, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsMay(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-01-01"), false},
		{Parse("2020-02-01"), false},
		{Parse("2020-03-01"), false},
		{Parse("2020-04-01"), false},
		{Parse("2020-05-01"), true},
		{Parse("2020-06-01"), false},
		{Parse("2020-07-01"), false},
		{Parse("2020-08-01"), false},
		{Parse("2020-09-01"), false},
		{Parse("2020-10-01"), false},
		{Parse("2020-11-01"), false},
		{Parse("2020-12-01"), false},
	}

	for _, v := range Tests {
		output := v.input.IsMay()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is May, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is May, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsJune(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-01-01"), false},
		{Parse("2020-02-01"), false},
		{Parse("2020-03-01"), false},
		{Parse("2020-04-01"), false},
		{Parse("2020-05-01"), false},
		{Parse("2020-06-01"), true},
		{Parse("2020-07-01"), false},
		{Parse("2020-08-01"), false},
		{Parse("2020-09-01"), false},
		{Parse("2020-10-01"), false},
		{Parse("2020-11-01"), false},
		{Parse("2020-12-01"), false},
	}

	for _, v := range Tests {
		output := v.input.IsJune()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is June, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is June, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsJuly(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-01-01"), false},
		{Parse("2020-02-01"), false},
		{Parse("2020-03-01"), false},
		{Parse("2020-04-01"), false},
		{Parse("2020-05-01"), false},
		{Parse("2020-06-01"), false},
		{Parse("2020-07-01"), true},
		{Parse("2020-08-01"), false},
		{Parse("2020-09-01"), false},
		{Parse("2020-10-01"), false},
		{Parse("2020-11-01"), false},
		{Parse("2020-12-01"), false},
	}

	for _, v := range Tests {
		output := v.input.IsJuly()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is July, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is July, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsAugust(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-01-01"), false},
		{Parse("2020-02-01"), false},
		{Parse("2020-03-01"), false},
		{Parse("2020-04-01"), false},
		{Parse("2020-05-01"), false},
		{Parse("2020-06-01"), false},
		{Parse("2020-07-01"), false},
		{Parse("2020-08-01"), true},
		{Parse("2020-09-01"), false},
		{Parse("2020-10-01"), false},
		{Parse("2020-11-01"), false},
		{Parse("2020-12-01"), false},
	}

	for _, v := range Tests {
		output := v.input.IsAugust()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is August, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is August, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsSeptember(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-01-01"), false},
		{Parse("2020-02-01"), false},
		{Parse("2020-03-01"), false},
		{Parse("2020-04-01"), false},
		{Parse("2020-05-01"), false},
		{Parse("2020-06-01"), false},
		{Parse("2020-07-01"), false},
		{Parse("2020-08-01"), false},
		{Parse("2020-09-01"), true},
		{Parse("2020-10-01"), false},
		{Parse("2020-11-01"), false},
		{Parse("2020-12-01"), false},
	}

	for _, v := range Tests {
		output := v.input.IsSeptember()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is september, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is september, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsOctober(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-01-01"), false},
		{Parse("2020-02-01"), false},
		{Parse("2020-03-01"), false},
		{Parse("2020-04-01"), false},
		{Parse("2020-05-01"), false},
		{Parse("2020-06-01"), false},
		{Parse("2020-07-01"), false},
		{Parse("2020-08-01"), false},
		{Parse("2020-09-01"), false},
		{Parse("2020-10-01"), true},
		{Parse("2020-11-01"), false},
		{Parse("2020-12-01"), false},
	}

	for _, v := range Tests {
		output := v.input.IsOctober()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is October, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is October, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsNovember(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-01-01"), false},
		{Parse("2020-02-01"), false},
		{Parse("2020-03-01"), false},
		{Parse("2020-04-01"), false},
		{Parse("2020-05-01"), false},
		{Parse("2020-06-01"), false},
		{Parse("2020-07-01"), false},
		{Parse("2020-08-01"), false},
		{Parse("2020-09-01"), false},
		{Parse("2020-10-01"), false},
		{Parse("2020-11-01"), true},
		{Parse("2020-12-01"), false},
	}

	for _, v := range Tests {
		output := v.input.IsNovember()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is November, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is November, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsDecember(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-01-01"), false},
		{Parse("2020-02-01"), false},
		{Parse("2020-03-01"), false},
		{Parse("2020-04-01"), false},
		{Parse("2020-05-01"), false},
		{Parse("2020-06-01"), false},
		{Parse("2020-07-01"), false},
		{Parse("2020-08-01"), false},
		{Parse("2020-09-01"), false},
		{Parse("2020-10-01"), false},
		{Parse("2020-11-01"), false},
		{Parse("2020-12-01"), true},
	}

	for _, v := range Tests {
		output := v.input.IsDecember()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is December, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is December, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsMonday(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-10-05"), true},
		{Parse("2020-10-06"), false},
		{Parse("2020-10-07"), false},
		{Parse("2020-10-08"), false},
		{Parse("2020-10-09"), false},
		{Parse("2020-10-10"), false},
		{Parse("2020-10-11"), false},
	}

	for _, v := range Tests {
		output := v.input.IsMonday()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is Monday, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is Monday, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsTuesday(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-10-05"), false},
		{Parse("2020-10-06"), true},
		{Parse("2020-10-07"), false},
		{Parse("2020-10-08"), false},
		{Parse("2020-10-09"), false},
		{Parse("2020-10-10"), false},
		{Parse("2020-10-11"), false},
	}

	for _, v := range Tests {
		output := v.input.IsTuesday()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is Tuesday, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is Tuesday, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsWednesday(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-10-05"), false},
		{Parse("2020-10-06"), false},
		{Parse("2020-10-07"), true},
		{Parse("2020-10-08"), false},
		{Parse("2020-10-09"), false},
		{Parse("2020-10-10"), false},
		{Parse("2020-10-11"), false},
	}

	for _, v := range Tests {
		output := v.input.IsWednesday()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is Wednesday, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is Wednesday, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsThursday(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-10-05"), false},
		{Parse("2020-10-06"), false},
		{Parse("2020-10-07"), false},
		{Parse("2020-10-08"), true},
		{Parse("2020-10-09"), false},
		{Parse("2020-10-10"), false},
		{Parse("2020-10-11"), false},
	}

	for _, v := range Tests {
		output := v.input.IsThursday()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is Thursday, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is Thursday, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsFriday(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-10-05"), false},
		{Parse("2020-10-06"), false},
		{Parse("2020-10-07"), false},
		{Parse("2020-10-08"), false},
		{Parse("2020-10-09"), true},
		{Parse("2020-10-10"), false},
		{Parse("2020-10-11"), false},
	}

	for _, v := range Tests {
		output := v.input.IsFriday()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is Friday, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is Friday, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsSaturday(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-10-05"), false},
		{Parse("2020-10-06"), false},
		{Parse("2020-10-07"), false},
		{Parse("2020-10-08"), false},
		{Parse("2020-10-09"), false},
		{Parse("2020-10-10"), true},
		{Parse("2020-10-11"), false},
	}

	for _, v := range Tests {
		output := v.input.IsSaturday()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is Saturday, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is Saturday, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsSunday(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-10-05"), false},
		{Parse("2020-10-06"), false},
		{Parse("2020-10-07"), false},
		{Parse("2020-10-08"), false},
		{Parse("2020-10-09"), false},
		{Parse("2020-10-10"), false},
		{Parse("2020-10-11"), true},
	}

	for _, v := range Tests {
		output := v.input.IsSunday()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is Sunday, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is Sunday, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsWeekday(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-10-05"), true},
		{Parse("2020-10-06"), true},
		{Parse("2020-10-07"), true},
		{Parse("2020-10-08"), true},
		{Parse("2020-10-09"), true},
		{Parse("2020-10-10"), false},
		{Parse("2020-10-11"), false},
	}

	for _, v := range Tests {
		output := v.input.IsWeekday()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is weekday, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is weekday, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsWeekend(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-10-05"), false},
		{Parse("2020-10-06"), false},
		{Parse("2020-10-07"), false},
		{Parse("2020-10-08"), false},
		{Parse("2020-10-09"), false},
		{Parse("2020-10-10"), true},
		{Parse("2020-10-11"), true},
	}

	for _, v := range Tests {
		output := v.input.IsWeekend()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is weekend, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is weekend, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsYesterday(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Now(), false},
		{Yesterday(), true},
		{Tomorrow(), false},
	}

	for _, v := range Tests {
		output := v.input.IsYesterday()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is yesterday, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is yesterday, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsToday(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Now(), true},
		{Yesterday(), false},
		{Tomorrow(), false},
	}

	for _, v := range Tests {
		output := v.input.IsToday()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is today, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is today, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_IsTomorrow(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Now(), false},
		{Yesterday(), false},
		{Tomorrow(), true},
	}

	for _, v := range Tests {
		output := v.input.IsTomorrow()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is tomorrow, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is tomorrow, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}
}

func TestCarbon_Compare(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		param1 string // 输入参数1
		param2 Carbon // 输入参数2
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05"), ">", Parse("2020-08-04"), true},
		{Parse("2020-08-05"), "<", Parse("2020-08-04"), false},
		{Parse("2020-08-05"), "<", Parse("2020-08-06"), true},
		{Parse("2020-08-05"), ">", Parse("2020-08-06"), false},
		{Parse("2020-08-05"), "=", Parse("2020-08-05"), true},
		{Parse("2020-08-05"), ">=", Parse("2020-08-05"), true},
		{Parse("2020-08-05"), "<=", Parse("2020-08-05"), true},
		{Parse("2020-08-05"), "!=", Parse("2020-08-05"), false},
		{Parse("2020-08-05"), "<>", Parse("2020-08-05"), false},
		{Parse("2020-08-05"), "!=", Parse("2020-08-04"), true},
		{Parse("2020-08-05"), "<>", Parse("2020-08-04"), true},
		{Parse("2020-08-05"), "+", Parse("2020-08-04"), false},
	}

	for _, v := range Tests {
		output := v.input.Compare(v.param1, v.param2)

		if output == true {
			if v.output == false {
				t.Errorf("Input %v %s %v, expected true, but got false\n", v.input.ToDateString(), v.param1, v.param2.ToDateTimeString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %v %s %v, expected false, but got true\n", v.input, v.param1, v.param2.ToDateTimeString())
			}
		}
	}
}

func TestCarbon_Gt(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		param  Carbon // 参数值
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05"), Parse("2020-08-05"), false},
		{Parse("2020-08-05"), Parse("2020-08-04"), true},
		{Parse("2020-08-05"), Parse("2020-08-06"), false},
	}

	for _, v := range Tests {
		output := v.input.Gt(v.param)

		if output == true {
			if v.output == false {
				t.Errorf("Input %s > %s, expected true, but got false\n", v.input.ToDateString(), v.param.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s > %s, expected false, but got true\n", v.input.ToDateString(), v.param.ToDateString())
			}
		}
	}
}

func TestCarbon_Lt(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		param  Carbon // 参数值
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05"), Parse("2020-08-05"), false},
		{Parse("2020-08-05"), Parse("2020-08-04"), false},
		{Parse("2020-08-05"), Parse("2020-08-06"), true},
	}

	for _, v := range Tests {
		output := v.input.Lt(v.param)

		if output == true {
			if v.output == false {
				t.Errorf("Input %s < %s, expected true, but got false\n", v.input.ToDateString(), v.param.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s < %s, expected false, but got true\n", v.input.ToDateString(), v.param.ToDateString())
			}
		}
	}
}

func TestCarbon_Eq(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		param  Carbon // 参数值
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05"), Parse("2020-08-05"), true},
		{Parse("2020-08-05"), Parse("2020-08-04"), false},
		{Parse("2020-08-05"), Parse("2020-08-06"), false},
	}

	for _, v := range Tests {
		output := v.input.Eq(v.param)

		if output == true {
			if v.output == false {
				t.Errorf("Input %s == %s, expected true, but got false\n", v.input.ToDateString(), v.param.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s == %s, expected false, but got true\n", v.input.ToDateString(), v.param.ToDateString())
			}
		}
	}
}

func TestCarbon_Ne(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		param  Carbon // 参数值
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05"), Parse("2020-08-05"), false},
		{Parse("2020-08-05"), Parse("2020-08-04"), true},
		{Parse("2020-08-05"), Parse("2020-08-06"), true},
	}

	for _, v := range Tests {
		output := v.input.Ne(v.param)

		if output == true {
			if v.output == false {
				t.Errorf("Input %s != %s, expected true, but got false\n", v.input.ToDateString(), v.param.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s != %s, expected false, but got true\n", v.input.ToDateString(), v.param.ToDateString())
			}
		}
	}
}

func TestCarbon_Gte(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		param  Carbon // 参数值
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05"), Parse("2020-08-05"), true},
		{Parse("2020-08-05"), Parse("2020-08-04"), true},
		{Parse("2020-08-05"), Parse("2020-08-06"), false},
	}

	for _, v := range Tests {
		output := v.input.Gte(v.param)

		if output == true {
			if v.output == false {
				t.Errorf("Input %s >= %s, expected true, but got false\n", v.input.ToDateString(), v.param.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s >= %s, expected false, but got true\n", v.input.ToDateString(), v.param.ToDateString())
			}
		}
	}
}

func TestCarbon_Lte(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		param  Carbon // 参数值
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05"), Parse("2020-08-05"), true},
		{Parse("2020-08-05"), Parse("2020-08-04"), false},
		{Parse("2020-08-05"), Parse("2020-08-06"), true},
	}

	for _, v := range Tests {
		output := v.input.Lte(v.param)

		if output == true {
			if v.output == false {
				t.Errorf("Input %s <= %s, expected true, but got false\n", v.input.ToDateString(), v.param.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s <= %s, expected false, but got true\n", v.input.ToDateString(), v.param.ToDateString())
			}
		}
	}
}

func TestCarbon_Between(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		param1 Carbon // 输入参数1
		param2 Carbon // 输入参数2
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), false},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), Parse("2020-08-06 13:14:15"), false},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-04 13:14:15"), Parse("2020-08-05 13:14:15"), false},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-04 13:14:15"), Parse("2020-08-06 13:14:15"), true},
	}

	for _, v := range Tests {
		output := v.input.Between(v.param1, v.param2)

		if output == true {
			if v.output == false {
				t.Errorf("Input %s < %s < %s, expected true, but got false\n", v.param1.ToDateString(), v.input.ToDateString(), v.param2.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s < %s < %s, expected false, but got true\n", v.param1.ToDateString(), v.input.ToDateString(), v.param2.ToDateString())
			}
		}
	}
}

func TestCarbon_BetweenIncludedStartTime(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		param1 Carbon // 输入参数1
		param2 Carbon // 输入参数2
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), false},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), Parse("2020-08-06 13:14:15"), true},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-04 13:14:15"), Parse("2020-08-05 13:14:15"), false},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-04 13:14:15"), Parse("2020-08-06 13:14:15"), true},
	}

	for _, v := range Tests {
		output := v.input.BetweenIncludedStartTime(v.param1, v.param2)

		if output == true {
			if v.output == false {
				t.Errorf("Input %s <= %s < %s, expected true, but got false\n", v.param1.ToDateString(), v.input.ToDateString(), v.param2.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s <= %s < %s, expected false, but got true\n", v.param1.ToDateString(), v.input.ToDateString(), v.param2.ToDateString())
			}
		}
	}
}

func TestCarbon_BetweenIncludedEndTime(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		param1 Carbon // 输入参数1
		param2 Carbon // 输入参数2
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), false},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), Parse("2020-08-06 13:14:15"), false},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-04 13:14:15"), Parse("2020-08-05 13:14:15"), true},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-04 13:14:15"), Parse("2020-08-06 13:14:15"), true},
	}

	for _, v := range Tests {
		output := v.input.BetweenIncludedEndTime(v.param1, v.param2)

		if output == true {
			if v.output == false {
				t.Errorf("Input %s < %s <= %s, expected true, but got false\n", v.param1.ToDateString(), v.input.ToDateString(), v.param2.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s < %s <= %s, expected false, but got true\n", v.param1.ToDateString(), v.input.ToDateString(), v.param2.ToDateString())
			}
		}
	}
}

func TestCarbon_BetweenIncludedBoth(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		param1 Carbon // 输入参数1
		param2 Carbon // 输入参数2
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), true},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), Parse("2020-08-06 13:14:15"), true},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-04 13:14:15"), Parse("2020-08-05 13:14:15"), true},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-04 13:14:15"), Parse("2020-08-06 13:14:15"), true},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-06 13:14:15"), Parse("2020-08-06 13:14:15"), false},
	}

	for _, v := range Tests {
		output := v.input.BetweenIncludedBoth(v.param1, v.param2)

		if output == true {
			if v.output == false {
				t.Errorf("Input %s <= %s <= %s, expected true, but got false\n", v.param1.ToDateString(), v.input.ToDateString(), v.param2.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s <= %s <= %s, expected false, but got true\n", v.param1.ToDateString(), v.input.ToDateString(), v.param2.ToDateString())
			}
		}
	}
}
