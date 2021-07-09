package carbon

import (
	"testing"
)

func TestCarbon_Constellation1(t *testing.T) {
	Tests := []struct {
		input  string // 输入值1
		output string // 期望输出值
	}{
		{"", ""},
		{"0000-00-00", ""},
		{"2020-01-05", "Capricorn"},
		{"2020-02-05", "Aquarius"},
		{"2020-03-05", "Pisces"},
		{"2020-04-05", "Aries"},
		{"2020-05-05", "Taurus"},
		{"2020-06-05", "Gemini"},
		{"2020-07-05", "Cancer"},
		{"2020-08-05", "Leo"},
		{"2020-09-05", "Virgo"},
		{"2020-10-05", "Libra"},
		{"2020-11-05", "Scorpio"},
		{"2020-12-05", "Sagittarius"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Constellation()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_Constellation(t *testing.T) {
	Tests := []struct {
		input1 string // 输入值1
		input2 string // 输入值2
		output string // 期望输出值
	}{
		{"", "en", ""},
		{"0000-00-00", "en", ""},
		{"2020-08-05", "en", "Leo"},
		{"2020-08-05", "zh-CN", "狮子座"},
		{"2020-08-05", "zh-Tw", "獅子座"},
		{"2020-08-05", "jp", "ししざ"},
	}

	for _, v := range Tests {
		output := Parse(v.input1).SetLocale(v.input2).Constellation()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s", v.input1, v.output, output)
		}
	}
}

func TestCarbon_IsAries(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0000-00-00", false},
		{"2020-03-21", true},
		{"2020-04-19", true},
		{"2020-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsAries()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is aries, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is aries, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestCarbon_IsTaurus(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0000-00-00", false},
		{"2020-04-20", true},
		{"2020-05-20", true},
		{"2020-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsTaurus()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is taurus, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is taurus, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestCarbon_IsGemini(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0000-00-00", false},
		{"2020-05-21", true},
		{"2020-06-21", true},
		{"2020-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsGemini()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is gemini, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is gemini, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestCarbon_IsCancer(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0000-00-00", false},
		{"2020-06-22", true},
		{"2020-07-22", true},
		{"2020-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsCancer()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is cancer, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is cancer, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestCarbon_IsLeo(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0000-00-00", false},
		{"2020-07-23", true},
		{"2020-08-05", true},
		{"2020-08-22", true},
		{"2020-08-23", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsLeo()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is leo, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is leo, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestCarbon_IsVirgo(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0000-00-00", false},
		{"2020-08-23", true},
		{"2020-09-22", true},
		{"2020-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsVirgo()

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

func TestCarbon_IsLibra(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0000-00-00", false},
		{"2020-09-23", true},
		{"2020-10-23", true},
		{"2020-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsLibra()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is libra, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is libra, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestCarbon_IsScorpio(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0000-00-00", false},
		{"2020-10-24", true},
		{"2020-11-22", true},
		{"2020-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsScorpio()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is scorpio, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is scorpio, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestCarbon_IsSagittarius(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0000-00-00", false},
		{"2020-11-23", true},
		{"2020-12-21", true},
		{"2020-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsSagittarius()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is sagittarius, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is sagittarius, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestCarbon_IsCapricorn(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0000-00-00", false},
		{"2020-12-22", true},
		{"2020-01-19", true},
		{"2020-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsCapricorn()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is capricorn, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is capricorn, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestCarbon_IsAquarius(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0000-00-00", false},
		{"2020-01-20", true},
		{"2020-02-18", true},
		{"2020-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsAquarius()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is aquarius, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is aquarius, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestCarbon_IsPisces(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0000-00-00", false},
		{"2020-02-19", true},
		{"2020-03-20", true},
		{"2020-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsPisces()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is pisces, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is pisces, expected false, but got true\n", v.input)
			}
		}
	}
}
