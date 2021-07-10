package carbon

import "testing"

func TestLunar_Animal(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-05-01", "鼠"},
		{"2020-08-05", "鼠"},
		{"2021-07-07", "牛"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().Animal()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestLunar_Year(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},
		{"2020-05-01", 2020},
		{"2020-08-05", 2020},
		{"2021-07-07", 2021},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().Year()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d\n", v.input, v.output, output)
		}
	}
}

func TestLunar_Month(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},
		{"2020-05-01", 4},
		{"2020-08-05", 6},
		{"2021-07-07", 5},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().Month()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d\n", v.input, v.output, output)
		}
	}
}

func TestLunar_Day(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},
		{"2020-05-01", 9},
		{"2020-08-05", 16},
		{"2021-07-07", 28},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().Day()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d\n", v.input, v.output, output)
		}
	}
}

func TestLunar_LeapMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},
		{"2020-05-01", 4},
		{"2020-08-05", 4},
		{"2021-07-07", 0},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().LeapMonth()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d\n", v.input, v.output, output)
		}
	}
}

func TestLunar_ToChinaYearString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-05-01", "二零二零"},
		{"2020-08-05", "二零二零"},
		{"2021-07-07", "二零二一"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().ToChineseYearString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestLunar_ToChinaMonthString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-05-01", "四月"},
		{"2020-08-05", "六月"},
		{"2021-07-07", "五月"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().ToChineseMonthString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestLunar_ToChinaDayString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-05-01", "初九"},
		{"2020-08-05", "十六"},
		{"2021-07-07", "廿八"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().ToChineseDayString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestLunar_ToGanZhiYearString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-05-01", "庚子"},
		{"2020-08-05", "庚子"},
		{"2021-07-07", "辛丑"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().ToGanZhiYearString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestLunar_IsLeapYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", true},
		{"2020-08-05", true},
		{"2021-07-07", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsLeapYear()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is leap year, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is leap month, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsLeapMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", true},
		{"2020-08-05", false},
		{"2021-07-07", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsLeapMonth()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is leap month, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is leap month, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsYearOfRat(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", true},
		{"2020-08-05", true},
		{"2021-07-07", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsYearOfRat()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of rat, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of rat, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsYearOfOx(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-07-07", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsYearOfOx()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of ox, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of ox, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsYearOfTiger(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-07-07", false},
		{"2022-08-05", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsYearOfTiger()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of tiger, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of tiger, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsYearOfRabbit(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-07-07", false},
		{"2023-08-05", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsYearOfRabbit()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of rabbit, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of rabbit, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsYearOfDragon(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-07-07", false},
		{"2024-08-05", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsYearOfDragon()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of dragon, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of dragon, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsYearOfSnake(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-07-07", false},
		{"2025-08-05", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsYearOfSnake()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of snake, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of snake, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsYearOfHorse(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-07-07", false},
		{"2026-08-05", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsYearOfHorse()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of horse, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of horse, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsYearOfGoat(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-07-07", false},
		{"2027-08-05", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsYearOfGoat()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of goat, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of goat, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsYearOfMonkey(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-07-07", false},
		{"2028-08-05", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsYearOfMonkey()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of monkey, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of monkey, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsYearOfRooster(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-07-07", false},
		{"2029-08-05", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsYearOfRooster()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of rooster, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of rooster, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsYearOfDog(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-07-07", false},
		{"2030-08-05", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsYearOfDog()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of dog, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of dog, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsYearOfPig(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-07-07", false},
		{"2031-08-05", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsYearOfPig()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of pig, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of pig, expected false, but got true\n", v.input)
			}
		}
	}
}
