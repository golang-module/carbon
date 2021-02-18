package carbon

import (
	"fmt"
	"testing"
)

func TestCarbon_AddDuration(t *testing.T) {
	Tests := []struct {
		input    string // 输入值
		duration string // 输入参数
		output   string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "10h", "2020-01-01 23:14:15"},
		{"2020-01-01 13:14:15", "10.5h", "2020-01-01 23:44:15"},

		{"2020-01-01 13:14:15", "10m", "2020-01-01 13:24:15"},
		{"2020-01-01 13:14:15", "10.5m", "2020-01-01 13:24:45"},

		{"2020-01-01 13:14:15", "10s", "2020-01-01 13:14:25"},
		{"2020-01-01 13:14:15", "10.5s", "2020-01-01 13:14:25"},

		{"2020-01-01 13:14:15", "XXXX", ""}, // 异常情况
	}

	for _, v := range Tests {
		output := Parse(v.input).AddDuration(v.duration)

		if output.Error != nil {
			fmt.Println("catch an exception in AddDuration():", output.Error)
		} else {
			if output.ToDateTimeString() != v.output {
				t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output.ToDateTimeString())
			}
		}
	}

	for _, v := range Tests {
		output := SetTimezone("XXXX").Parse(v.input).AddDuration(v.duration)

		if output.Error != nil {
			fmt.Println("catch an exception in AddDuration():", output.Error)
			return
		}

		if output.ToDateTimeString() != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output.ToDateTimeString())
		}
	}
}

func TestCarbon_SubDuration(t *testing.T) {
	Tests := []struct {
		input    string // 输入值
		duration string // 输入参数
		output   string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "10h", "2020-01-01 03:14:15"},
		{"2020-01-01 13:14:15", "10.5h", "2020-01-01 02:44:15"},

		{"2020-01-01 13:14:15", "10m", "2020-01-01 13:04:15"},
		{"2020-01-01 13:14:15", "10.5m", "2020-01-01 13:03:45"},

		{"2020-01-01 13:14:15", "10s", "2020-01-01 13:14:05"},
		{"2020-01-01 13:14:15", "10.5s", "2020-01-01 13:14:04"},

		{"2020-01-01 13:14:15 XXXX", "10h", ""}, // 异常情况
		{"2020-01-01 13:14:15", "10x", ""},      // 异常情况
	}

	for _, v := range Tests {
		output := Parse(v.input).SubDuration(v.duration)

		if output.Error != nil {
			fmt.Println("catch an exception in SubDuration():", output.Error)
			return
		}

		if output.ToDateTimeString() != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output.ToDateTimeString())
		}
	}

	for _, v := range Tests {
		output := SetTimezone("XXXX").Parse(v.input).SubDuration(v.duration)

		if output.Error != nil {
			fmt.Println("catch an exception in SubDuration():", output.Error)
			return
		}

		if output.ToDateTimeString() != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output.ToDateTimeString())
		}
	}
}

func TestCarbon_AddCenturies(t *testing.T) {
	type Test struct {
		input     string // 输入值
		centuries int    // 输入参数
		output    string // 期望输出值
	}

	Tests := []Test{
		{"2020-01-01", 3, "2320-01-01"},
		{"2020-01-31", 3, "2320-01-31"},
		{"2020-02-01", 3, "2320-02-01"},
		{"2020-02-28", 3, "2320-02-28"},
		{"2020-02-29", 3, "2320-02-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddCenturies(v.centuries).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddCenturies(v.centuries).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddCenturiesNoOverflow(t *testing.T) {
	Tests := []struct {
		input     string // 输入值
		centuries int    // 输入参数
		output    string // 期望输出值
	}{
		{"2020-01-01", 3, "2320-01-01"},
		{"2020-01-31", 3, "2320-01-31"},
		{"2020-02-01", 3, "2320-02-01"},
		{"2020-02-28", 3, "2320-02-28"},
		{"2020-02-29", 3, "2320-02-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddCenturiesNoOverflow(v.centuries).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddCenturiesNoOverflow(v.centuries).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubCenturies(t *testing.T) {
	type Test struct {
		input     string // 输入值
		centuries int    // 输入参数
		output    string // 期望输出值
	}

	Tests := []Test{
		{"2020-01-01", 3, "1720-01-01"},
		{"2020-01-31", 3, "1720-01-31"},
		{"2020-02-01", 3, "1720-02-01"},
		{"2020-02-28", 3, "1720-02-28"},
		{"2020-02-29", 3, "1720-02-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubCenturies(v.centuries).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubCenturies(v.centuries).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubCenturiesNoOverflow(t *testing.T) {
	Tests := []struct {
		input     string // 输入值
		centuries int
		output    string // 期望输出值
	}{
		{"2020-01-01", 3, "1720-01-01"},
		{"2020-01-31", 3, "1720-01-31"},
		{"2020-02-01", 3, "1720-02-01"},
		{"2020-02-28", 3, "1720-02-28"},
		{"2020-02-29", 3, "1720-02-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubCenturiesNoOverflow(v.centuries).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubCenturiesNoOverflow(v.centuries).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddCentury(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2120-01-01"},
		{"2020-01-31", "2120-01-31"},
		{"2020-02-01", "2120-02-01"},
		{"2020-02-28", "2120-02-28"},
		{"2020-02-29", "2120-02-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddCentury().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddCentury().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddCenturyNoOverflow(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2120-01-01"},
		{"2020-01-31", "2120-01-31"},
		{"2020-02-01", "2120-02-01"},
		{"2020-02-28", "2120-02-28"},
		{"2020-02-29", "2120-02-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddCenturyNoOverflow().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddCenturyNoOverflow().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubCentury(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "1920-01-01"},
		{"2020-01-31", "1920-01-31"},
		{"2020-02-01", "1920-02-01"},
		{"2020-02-28", "1920-02-28"},
		{"2020-02-29", "1920-02-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubCentury().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubCentury().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubCenturyNoOverflow(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "1920-01-01"},
		{"2020-01-31", "1920-01-31"},
		{"2020-02-01", "1920-02-01"},
		{"2020-02-28", "1920-02-28"},
		{"2020-02-29", "1920-02-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubCenturyNoOverflow().ToDateString()

		if output != v.output {
			t.Errorf("Expected %s, but got %s", v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubCenturyNoOverflow().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddYears(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		years  int    // 输入参数
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2023-01-01"},
		{"2020-01-31", 3, "2023-01-31"},
		{"2020-02-01", 3, "2023-02-01"},
		{"2020-02-28", 3, "2023-02-28"},
		{"2020-02-29", 3, "2023-03-01"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddYears(v.years).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddYears(v.years).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddYearsNoOverflow(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		years  int    // 输入参数
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2023-01-01"},
		{"2020-01-31", 3, "2023-01-31"},
		{"2020-02-01", 3, "2023-02-01"},
		{"2020-02-28", 3, "2023-02-28"},
		{"2020-02-29", 3, "2023-02-28"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddYearsNoOverflow(v.years).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddYearsNoOverflow(v.years).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubYears(t *testing.T) {
	type Test struct {
		input  string // 输入值
		years  int    // 输入参数
		output string // 期望输出值
	}

	Tests := []Test{
		{"2020-01-01", 3, "2017-01-01"},
		{"2020-01-31", 3, "2017-01-31"},
		{"2020-02-01", 3, "2017-02-01"},
		{"2020-02-28", 3, "2017-02-28"},
		{"2020-02-29", 3, "2017-03-01"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubYears(v.years).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubYears(v.years).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubYearsNoOverflow(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		years  int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2017-01-01"},
		{"2020-01-31", 3, "2017-01-31"},
		{"2020-02-01", 3, "2017-02-01"},
		{"2020-02-28", 3, "2017-02-28"},
		{"2020-02-29", 3, "2017-02-28"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubYearsNoOverflow(v.years).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubYearsNoOverflow(v.years).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2021-01-01"},
		{"2020-01-31", "2021-01-31"},
		{"2020-02-01", "2021-02-01"},
		{"2020-02-28", "2021-02-28"},
		{"2020-02-29", "2021-03-01"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddYear().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddYear().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddYearNoOverflow(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2021-01-01"},
		{"2020-01-31", "2021-01-31"},
		{"2020-02-01", "2021-02-01"},
		{"2020-02-28", "2021-02-28"},
		{"2020-02-29", "2021-02-28"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddYearNoOverflow().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddYearNoOverflow().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2019-01-01"},
		{"2020-01-31", "2019-01-31"},
		{"2020-02-01", "2019-02-01"},
		{"2020-02-28", "2019-02-28"},
		{"2020-02-29", "2019-03-01"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubYear().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubYear().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubYearNoOverflow(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2019-01-01"},
		{"2020-01-31", "2019-01-31"},
		{"2020-02-01", "2019-02-01"},
		{"2020-02-28", "2019-02-28"},
		{"2020-02-29", "2019-02-28"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubYearNoOverflow().ToDateString()

		if output != v.output {
			t.Errorf("Expected %s, but got %s", v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubYearNoOverflow().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddQuarters(t *testing.T) {
	Tests := []struct {
		input    string // 输入值
		quarters int
		output   string // 期望输出值
	}{
		{"2019-08-01", 2, "2020-02-01"},
		{"2019-08-31", 2, "2020-03-02"},
		{"2020-01-01", 2, "2020-07-01"},
		{"2020-02-28", 2, "2020-08-28"},
		{"2020-02-29", 2, "2020-08-29"},
		{"2020-08-31", 2, "2021-03-03"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddQuarters(v.quarters).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddQuarters(v.quarters).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddQuartersNoOverflow(t *testing.T) {
	Tests := []struct {
		input    string // 输入值
		quarters int
		output   string // 期望输出值
	}{
		{"2019-08-01", 2, "2020-02-01"},
		{"2019-08-31", 2, "2020-02-29"},
		{"2020-01-01", 2, "2020-07-01"},
		{"2020-02-28", 2, "2020-08-28"},
		{"2020-02-29", 2, "2020-08-29"},
		{"2020-08-31", 2, "2021-02-28"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddQuartersNoOverflow(v.quarters).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddQuartersNoOverflow(v.quarters).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubQuarters(t *testing.T) {
	Tests := []struct {
		input    string // 输入值
		quarters int
		output   string // 期望输出值
	}{
		{"2019-08-01", 2, "2019-02-01"},
		{"2019-08-31", 2, "2019-03-03"},
		{"2020-01-01", 2, "2019-07-01"},
		{"2020-02-28", 2, "2019-08-28"},
		{"2020-02-29", 2, "2019-08-29"},
		{"2020-08-31", 2, "2020-03-02"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubQuarters(v.quarters).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubQuarters(v.quarters).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubQuartersNoOverflow(t *testing.T) {
	Tests := []struct {
		input    string // 输入值
		quarters int
		output   string // 期望输出值
	}{
		{"2019-08-01", 2, "2019-02-01"},
		{"2019-08-31", 2, "2019-02-28"},
		{"2020-01-01", 2, "2019-07-01"},
		{"2020-02-28", 2, "2019-08-28"},
		{"2020-02-29", 2, "2019-08-29"},
		{"2020-08-31", 2, "2020-02-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubQuartersNoOverflow(v.quarters).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubQuartersNoOverflow(v.quarters).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddQuarter(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2019-11-01", "2020-02-01"},
		{"2019-11-30", "2020-03-01"},
		{"2020-02-28", "2020-05-28"},
		{"2020-02-29", "2020-05-29"},
		{"2020-08-31", "2020-12-01"},
		{"2020-11-01", "2021-02-01"},
		{"2020-11-30", "2021-03-02"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddQuarter().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddQuarter().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddQuarterNoOverflow(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2019-11-01", "2020-02-01"},
		{"2019-11-30", "2020-02-29"},
		{"2020-02-28", "2020-05-28"},
		{"2020-02-29", "2020-05-29"},
		{"2020-08-31", "2020-11-30"},
		{"2020-11-01", "2021-02-01"},
		{"2020-11-30", "2021-02-28"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddQuarterNoOverflow().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddQuarterNoOverflow().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubQuarter(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2019-04-01", "2019-01-01"},
		{"2019-04-30", "2019-01-30"},
		{"2020-05-01", "2020-02-01"},
		{"2020-05-31", "2020-03-02"},
		{"2020-04-01", "2020-01-01"},
		{"2020-04-30", "2020-01-30"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubQuarter().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubQuarter().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubQuarterNoOverflow(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2019-04-01", "2019-01-01"},
		{"2019-04-30", "2019-01-30"},
		{"2020-05-01", "2020-02-01"},
		{"2020-05-31", "2020-02-29"},
		{"2020-04-01", "2020-01-01"},
		{"2020-04-30", "2020-01-30"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubQuarterNoOverflow().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubQuarterNoOverflow().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddMonths(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		months int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2020-04-01"},
		{"2020-01-31", 3, "2020-05-01"},
		{"2020-02-01", 3, "2020-05-01"},
		{"2020-02-28", 3, "2020-05-28"},
		{"2020-02-29", 3, "2020-05-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddMonths(v.months).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddMonths(v.months).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddMonthsNoOverflow(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		months int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2020-04-01"},
		{"2020-01-31", 3, "2020-04-30"},
		{"2020-02-01", 3, "2020-05-01"},
		{"2020-02-28", 3, "2020-05-28"},
		{"2020-02-29", 3, "2020-05-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddMonthsNoOverflow(v.months).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddMonthsNoOverflow(v.months).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubMonths(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		months int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2019-10-01"},
		{"2020-01-31", 3, "2019-10-31"},
		{"2020-02-01", 3, "2019-11-01"},
		{"2020-02-28", 3, "2019-11-28"},
		{"2020-02-29", 3, "2019-11-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubMonths(v.months).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubMonths(v.months).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubMonthsNoOverflow(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		months int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2019-10-01"},
		{"2020-01-31", 3, "2019-10-31"},
		{"2020-02-01", 3, "2019-11-01"},
		{"2020-02-28", 3, "2019-11-28"},
		{"2020-02-29", 3, "2019-11-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubMonthsNoOverflow(v.months).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubMonthsNoOverflow(v.months).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2020-02-01"},
		{"2020-01-31", "2020-03-02"},
		{"2020-02-01", "2020-03-01"},
		{"2020-02-28", "2020-03-28"},
		{"2020-02-29", "2020-03-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddMonth().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddMonth().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddMonthNoOverflow(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2020-02-01"},
		{"2020-01-31", "2020-02-29"},
		{"2020-02-01", "2020-03-01"},
		{"2020-02-28", "2020-03-28"},
		{"2020-02-29", "2020-03-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddMonthNoOverflow().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddMonthNoOverflow().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2019-12-01"},
		{"2020-01-31", "2019-12-31"},
		{"2020-02-01", "2020-01-01"},
		{"2020-02-28", "2020-01-28"},
		{"2020-02-29", "2020-01-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubMonth().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubMonth().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubMonthNoOverflow(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2019-12-01"},
		{"2020-01-31", "2019-12-31"},
		{"2020-02-01", "2020-01-01"},
		{"2020-02-28", "2020-01-28"},
		{"2020-02-29", "2020-01-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubMonthNoOverflow().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubMonthNoOverflow().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddWeeks(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		weeks  int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2020-01-22"},
		{"2020-01-31", 3, "2020-02-21"},
		{"2020-02-01", 3, "2020-02-22"},
		{"2020-02-28", 3, "2020-03-20"},
		{"2020-02-29", 3, "2020-03-21"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddWeeks(v.weeks).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddWeeks(v.weeks).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubWeeks(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		weeks  int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2019-12-11"},
		{"2020-01-31", 3, "2020-01-10"},
		{"2020-02-01", 3, "2020-01-11"},
		{"2020-02-28", 3, "2020-02-07"},
		{"2020-02-29", 3, "2020-02-08"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubWeeks(v.weeks).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubWeeks(v.weeks).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddWeek(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2020-01-08"},
		{"2020-01-31", "2020-02-07"},
		{"2020-02-01", "2020-02-08"},
		{"2020-02-28", "2020-03-06"},
		{"2020-02-29", "2020-03-07"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddWeek().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddWeek().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubWeek(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2019-12-25"},
		{"2020-01-31", "2020-01-24"},
		{"2020-02-01", "2020-01-25"},
		{"2020-02-28", "2020-02-21"},
		{"2020-02-29", "2020-02-22"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubWeek().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubWeek().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddDays(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		days   int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2020-01-04"},
		{"2020-01-31", 3, "2020-02-03"},
		{"2020-02-01", 3, "2020-02-04"},
		{"2020-02-28", 3, "2020-03-02"},
		{"2020-02-29", 3, "2020-03-03"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddDays(v.days).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddDays(v.days).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubDays(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		days   int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2019-12-29"},
		{"2020-01-31", 3, "2020-01-28"},
		{"2020-02-01", 3, "2020-01-29"},
		{"2020-02-28", 3, "2020-02-25"},
		{"2020-02-29", 3, "2020-02-26"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubDays(v.days).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubDays(v.days).ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddDay(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2020-01-02"},
		{"2020-01-31", "2020-02-01"},
		{"2020-02-01", "2020-02-02"},
		{"2020-02-28", "2020-02-29"},
		{"2020-02-29", "2020-03-01"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddDay().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddDay().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubDay(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2019-12-31"},
		{"2020-01-31", "2020-01-30"},
		{"2020-02-01", "2020-01-31"},
		{"2020-02-28", "2020-02-27"},
		{"2020-02-29", "2020-02-28"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubDay().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubDay().ToDateString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddHours(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		hours  int
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", 3, "2020-01-01 16:14:15"},
		{"2020-01-31 13:14:15", 3, "2020-01-31 16:14:15"},
		{"2020-02-01 13:14:15", 3, "2020-02-01 16:14:15"},
		{"2020-02-28 13:14:15", 3, "2020-02-28 16:14:15"},
		{"2020-02-29 13:14:15", 3, "2020-02-29 16:14:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddHours(v.hours).ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddHours(v.hours).ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubHours(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		hours  int
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", 3, "2020-01-01 10:14:15"},
		{"2020-01-31 13:14:15", 3, "2020-01-31 10:14:15"},
		{"2020-02-01 13:14:15", 3, "2020-02-01 10:14:15"},
		{"2020-02-28 13:14:15", 3, "2020-02-28 10:14:15"},
		{"2020-02-29 13:14:15", 3, "2020-02-29 10:14:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubHours(v.hours).ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubHours(v.hours).ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddHour(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "2020-01-01 14:14:15"},
		{"2020-01-31 13:14:15", "2020-01-31 14:14:15"},
		{"2020-02-01 13:14:15", "2020-02-01 14:14:15"},
		{"2020-02-28 13:14:15", "2020-02-28 14:14:15"},
		{"2020-02-29 13:14:15", "2020-02-29 14:14:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddHour().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddHour().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubHour(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "2020-01-01 12:14:15"},
		{"2020-01-31 13:14:15", "2020-01-31 12:14:15"},
		{"2020-02-01 13:14:15", "2020-02-01 12:14:15"},
		{"2020-02-28 13:14:15", "2020-02-28 12:14:15"},
		{"2020-02-29 13:14:15", "2020-02-29 12:14:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubHour().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubHour().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddMinutes(t *testing.T) {
	Tests := []struct {
		input   string // 输入值
		minutes int
		output  string // 期望输出值
	}{
		{"2020-01-01 13:14:15", 3, "2020-01-01 13:17:15"},
		{"2020-01-31 13:14:15", 3, "2020-01-31 13:17:15"},
		{"2020-02-01 13:14:15", 3, "2020-02-01 13:17:15"},
		{"2020-02-28 13:14:15", 3, "2020-02-28 13:17:15"},
		{"2020-02-29 13:14:15", 3, "2020-02-29 13:17:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddMinutes(v.minutes).ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddMinutes(v.minutes).ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubMinutes(t *testing.T) {
	Tests := []struct {
		input   string // 输入值
		minutes int
		output  string // 期望输出值
	}{
		{"2020-01-01 13:14:15", 3, "2020-01-01 13:11:15"},
		{"2020-01-31 13:14:15", 3, "2020-01-31 13:11:15"},
		{"2020-02-01 13:14:15", 3, "2020-02-01 13:11:15"},
		{"2020-02-28 13:14:15", 3, "2020-02-28 13:11:15"},
		{"2020-02-29 13:14:15", 3, "2020-02-29 13:11:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubMinutes(v.minutes).ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubMinutes(v.minutes).ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddMinute(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "2020-01-01 13:15:15"},
		{"2020-01-31 13:14:15", "2020-01-31 13:15:15"},
		{"2020-02-01 13:14:15", "2020-02-01 13:15:15"},
		{"2020-02-28 13:14:15", "2020-02-28 13:15:15"},
		{"2020-02-29 13:14:15", "2020-02-29 13:15:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddMinute().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddMinute().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubMinute(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "2020-01-01 13:13:15"},
		{"2020-01-31 13:14:15", "2020-01-31 13:13:15"},
		{"2020-02-01 13:14:15", "2020-02-01 13:13:15"},
		{"2020-02-28 13:14:15", "2020-02-28 13:13:15"},
		{"2020-02-29 13:14:15", "2020-02-29 13:13:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubMinute().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubMinute().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddSeconds(t *testing.T) {
	Tests := []struct {
		input   string // 输入值
		seconds int
		output  string // 期望输出值
	}{
		{"2020-01-01 13:14:15", 3, "2020-01-01 13:14:18"},
		{"2020-01-31 13:14:15", 3, "2020-01-31 13:14:18"},
		{"2020-02-01 13:14:15", 3, "2020-02-01 13:14:18"},
		{"2020-02-28 13:14:15", 3, "2020-02-28 13:14:18"},
		{"2020-02-29 13:14:15", 3, "2020-02-29 13:14:18"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddSeconds(v.seconds).ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddSeconds(v.seconds).ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubSeconds(t *testing.T) {
	Tests := []struct {
		input   string // 输入值
		seconds int
		output  string // 期望输出值
	}{
		{"2020-01-01 13:14:15", 3, "2020-01-01 13:14:12"},
		{"2020-01-31 13:14:15", 3, "2020-01-31 13:14:12"},
		{"2020-02-01 13:14:15", 3, "2020-02-01 13:14:12"},
		{"2020-02-28 13:14:15", 3, "2020-02-28 13:14:12"},
		{"2020-02-29 13:14:15", 3, "2020-02-29 13:14:12"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubSeconds(v.seconds).ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubSeconds(v.seconds).ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddSecond(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "2020-01-01 13:14:16"},
		{"2020-01-31 13:14:15", "2020-01-31 13:14:16"},
		{"2020-02-01 13:14:15", "2020-02-01 13:14:16"},
		{"2020-02-28 13:14:15", "2020-02-28 13:14:16"},
		{"2020-02-29 13:14:15", "2020-02-29 13:14:16"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddSecond().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddSecond().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubSecond(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "2020-01-01 13:14:14"},
		{"2020-01-31 13:14:15", "2020-01-31 13:14:14"},
		{"2020-02-01 13:14:15", "2020-02-01 13:14:14"},
		{"2020-02-28 13:14:15", "2020-02-28 13:14:14"},
		{"2020-02-29 13:14:15", "2020-02-29 13:14:14"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubSecond().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubSecond().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}
