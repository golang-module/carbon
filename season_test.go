package carbon

import (
	"testing"
)

func TestCarbon_Season1(t *testing.T) {
	Tests := []struct {
		input  string // 输入值1
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"2020-01-05", "Winter"},
		{"2020-02-05", "Winter"},
		{"2020-03-05", "Spring"},
		{"2020-04-05", "Spring"},
		{"2020-05-05", "Spring"},
		{"2020-06-05", "Summer"},
		{"2020-07-05", "Summer"},
		{"2020-08-05", "Summer"},
		{"2020-09-05", "Autumn"},
		{"2020-10-05", "Autumn"},
		{"2020-11-05", "Autumn"},
		{"2020-12-05", "Winter"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Season()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_Season2(t *testing.T) {
	Tests := []struct {
		input1 string // 输入值1
		input2 string // 输入值2
		output string // 期望输出值
	}{
		{"", "en", ""},
		{"0000-00-00", "en", ""},
		{"2020-08-05", "en", "Summer"},
		{"2020-08-05", "zh-CN", "夏季"},
		{"2020-08-05", "zh-Tw", "夏季"},
		{"2020-08-05", "jp", "なつ"},
		{"2020-08-05", "kr", "여름"},
	}

	for _, v := range Tests {
		output := Parse(v.input1).SetLocale(v.input2).Season()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input1, v.output, output)
		}
	}
}

func TestCarbon_StartOfSeason(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-01-15 00:00:00", "2019-12-01 00:00:00"},
		{"2020-02-15 00:00:00", "2019-12-01 00:00:00"},
		{"2020-03-15 00:00:00", "2020-03-01 00:00:00"},
		{"2020-04-15 23:59:59", "2020-03-01 00:00:00"},
		{"2020-05-15 23:59:59", "2020-03-01 00:00:00"},
		{"2020-06-15 23:59:59", "2020-06-01 00:00:00"},
		{"2020-07-15 23:59:59", "2020-06-01 00:00:00"},
		{"2020-08-15 13:14:15", "2020-06-01 00:00:00"},
		{"2020-09-15 13:14:15", "2020-09-01 00:00:00"},
		{"2020-10-15", "2020-09-01 00:00:00"},
		{"2020-11-15", "2020-09-01 00:00:00"},
		{"2020-12-15", "2020-12-01 00:00:00"},
		{"2021-01-15", "2020-12-01 00:00:00"},
		{"2021-01-15", "2020-12-01 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).StartOfSeason().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfSeason(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-01-15 00:00:00", "2020-02-29 23:59:59"},
		{"2020-02-15 00:00:00", "2020-02-29 23:59:59"},
		{"2020-03-15 00:00:00", "2020-05-31 23:59:59"},
		{"2020-04-15 23:59:59", "2020-05-31 23:59:59"},
		{"2020-05-15 23:59:59", "2020-05-31 23:59:59"},
		{"2020-06-15 23:59:59", "2020-08-31 23:59:59"},
		{"2020-07-15 23:59:59", "2020-08-31 23:59:59"},
		{"2020-08-15 13:14:15", "2020-08-31 23:59:59"},
		{"2020-09-15 13:14:15", "2020-11-30 23:59:59"},
		{"2020-10-15", "2020-11-30 23:59:59"},
		{"2020-11-15", "2020-11-30 23:59:59"},
		{"2020-12-15", "2021-02-28 23:59:59"},
		{"2021-01-15", "2021-02-28 23:59:59"},
		{"2021-02-15", "2021-02-28 23:59:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfSeason().ToDateTimeString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_IsSpring(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", true},
		{"2020-04-01", true},
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
		output := Parse(v.input).IsSpring()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestCarbon_IsSummer(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", true},
		{"2020-07-01", true},
		{"2020-08-01", true},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsSummer()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestCarbon_IsAutumn(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", false},
		{"2020-07-01", false},
		{"2020-08-01", false},
		{"2020-09-01", true},
		{"2020-10-01", true},
		{"2020-11-01", true},
		{"2020-12-01", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsAutumn()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestCarbon_IsWinter(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"2020-01-01", true},
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
		{"2020-12-01", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsWinter()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s, expected false, but got true\n", v.input)
			}
		}
	}
}
