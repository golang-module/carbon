package carbon

import "testing"

func TestCarbon_AnimalYear(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output string // 期望输出值
	}{
		{Parse("0000-00-00"), ""},
		{Parse("2020-08-05"), "鼠"},
		{Parse("2021-08-05"), "牛"},
		{Parse("2022-08-05"), "虎"},
		{Parse("2023-08-05"), "兔"},
		{Parse("2024-08-05"), "龙"},
		{Parse("2025-08-05"), "蛇"},
		{Parse("2026-08-05"), "马"},
		{Parse("2027-08-05"), "羊"},
		{Parse("2028-08-05"), "猴"},
		{Parse("2029-08-05"), "鸡"},
		{Parse("2030-08-05"), "狗"},
		{Parse("2031-08-05"), "猪"},
	}

	for _, v := range Tests {
		output := v.input.AnimalYear()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input.ToDateString(), v.output, output)
		}
	}
}

func TestCarbon_ToChineseYearStringYear(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output string // 期望输出值
	}{
		{Parse("0000-00-00"), ""},
		{Parse("1861-08-22"), "辛酉"}, // 辛酉政变发生日期
		{Parse("1894-07-25"), "甲午"}, // 甲午战争发生日期
		{Parse("1898-06-11"), "戊戌"}, // 戊戌变法发生日期
		{Parse("1901-09-07"), "辛丑"}, // 辛丑条约签署日期
		{Parse("1911-10-10"), "辛亥"}, // 辛亥革命发生日期
		{Parse("1900-08-28"), "庚子"}, // 庚子赔款发生日期
	}

	for _, v := range Tests {
		output := v.input.ToChineseYearString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input.ToDateString(), v.output, output)
		}
	}
}

func TestCarbon_IsYearOfRat(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05"), true},
		{Parse("2021-08-05"), false},
		{Parse("2022-08-05"), false},
		{Parse("2023-08-05"), false},
		{Parse("2024-08-05"), false},
		{Parse("2025-08-05"), false},
		{Parse("2026-08-05"), false},
		{Parse("2027-08-05"), false},
		{Parse("2028-08-05"), false},
		{Parse("2029-08-05"), false},
		{Parse("2030-08-05"), false},
		{Parse("2031-08-05"), false},
	}

	for _, v := range Tests {
		output := v.input.IsYearOfRat()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of rat, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of rat, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}

}

func TestCarbon_IsYearOfOx(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05"), false},
		{Parse("2021-08-05"), true},
		{Parse("2022-08-05"), false},
		{Parse("2023-08-05"), false},
		{Parse("2024-08-05"), false},
		{Parse("2025-08-05"), false},
		{Parse("2026-08-05"), false},
		{Parse("2027-08-05"), false},
		{Parse("2028-08-05"), false},
		{Parse("2029-08-05"), false},
		{Parse("2030-08-05"), false},
		{Parse("2031-08-05"), false},
	}

	for _, v := range Tests {
		output := v.input.IsYearOfOx()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of ox, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of ox, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}

}

func TestCarbon_IsYearOfTiger(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05"), false},
		{Parse("2021-08-05"), false},
		{Parse("2022-08-05"), true},
		{Parse("2023-08-05"), false},
		{Parse("2024-08-05"), false},
		{Parse("2025-08-05"), false},
		{Parse("2026-08-05"), false},
		{Parse("2027-08-05"), false},
		{Parse("2028-08-05"), false},
		{Parse("2029-08-05"), false},
		{Parse("2030-08-05"), false},
		{Parse("2031-08-05"), false},
	}

	for _, v := range Tests {
		output := v.input.IsYearOfTiger()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of tiger, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of tiger, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}

}

func TestCarbon_IsYearOfRabbit(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05"), false},
		{Parse("2021-08-05"), false},
		{Parse("2022-08-05"), false},
		{Parse("2023-08-05"), true},
		{Parse("2024-08-05"), false},
		{Parse("2025-08-05"), false},
		{Parse("2026-08-05"), false},
		{Parse("2027-08-05"), false},
		{Parse("2028-08-05"), false},
		{Parse("2029-08-05"), false},
		{Parse("2030-08-05"), false},
		{Parse("2031-08-05"), false},
	}

	for _, v := range Tests {
		output := v.input.IsYearOfRabbit()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of rabbit, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of rabbit, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}

}

func TestCarbon_IsYearOfDragon(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05"), false},
		{Parse("2021-08-05"), false},
		{Parse("2022-08-05"), false},
		{Parse("2023-08-05"), false},
		{Parse("2024-08-05"), true},
		{Parse("2025-08-05"), false},
		{Parse("2026-08-05"), false},
		{Parse("2027-08-05"), false},
		{Parse("2028-08-05"), false},
		{Parse("2029-08-05"), false},
		{Parse("2030-08-05"), false},
		{Parse("2031-08-05"), false},
	}

	for _, v := range Tests {
		output := v.input.IsYearOfDragon()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of dragon, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of dragon, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}

}

func TestCarbon_IsYearOfSnake(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05"), false},
		{Parse("2021-08-05"), false},
		{Parse("2022-08-05"), false},
		{Parse("2023-08-05"), false},
		{Parse("2024-08-05"), false},
		{Parse("2025-08-05"), true},
		{Parse("2026-08-05"), false},
		{Parse("2027-08-05"), false},
		{Parse("2028-08-05"), false},
		{Parse("2029-08-05"), false},
		{Parse("2030-08-05"), false},
		{Parse("2031-08-05"), false},
	}

	for _, v := range Tests {
		output := v.input.IsYearOfSnake()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of snake, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of snake, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}

}

func TestCarbon_IsYearOfHorse(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05"), false},
		{Parse("2021-08-05"), false},
		{Parse("2022-08-05"), false},
		{Parse("2023-08-05"), false},
		{Parse("2024-08-05"), false},
		{Parse("2025-08-05"), false},
		{Parse("2026-08-05"), true},
		{Parse("2027-08-05"), false},
		{Parse("2028-08-05"), false},
		{Parse("2029-08-05"), false},
		{Parse("2030-08-05"), false},
		{Parse("2031-08-05"), false},
	}

	for _, v := range Tests {
		output := v.input.IsYearOfHorse()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of horse, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of horse, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}

}

func TestCarbon_IsYearOfGoat(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05"), false},
		{Parse("2021-08-05"), false},
		{Parse("2022-08-05"), false},
		{Parse("2023-08-05"), false},
		{Parse("2024-08-05"), false},
		{Parse("2025-08-05"), false},
		{Parse("2026-08-05"), false},
		{Parse("2027-08-05"), true},
		{Parse("2028-08-05"), false},
		{Parse("2029-08-05"), false},
		{Parse("2030-08-05"), false},
		{Parse("2031-08-05"), false},
	}

	for _, v := range Tests {
		output := v.input.IsYearOfGoat()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of goat, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of goat, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}

}

func TestCarbon_IsYearOfMonkey(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05"), false},
		{Parse("2021-08-05"), false},
		{Parse("2022-08-05"), false},
		{Parse("2023-08-05"), false},
		{Parse("2024-08-05"), false},
		{Parse("2025-08-05"), false},
		{Parse("2026-08-05"), false},
		{Parse("2027-08-05"), false},
		{Parse("2028-08-05"), true},
		{Parse("2029-08-05"), false},
		{Parse("2030-08-05"), false},
		{Parse("2031-08-05"), false},
	}

	for _, v := range Tests {
		output := v.input.IsYearOfMonkey()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of monkey, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of monkey, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}

}

func TestCarbon_IsYearOfRooster(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05"), false},
		{Parse("2021-08-05"), false},
		{Parse("2022-08-05"), false},
		{Parse("2023-08-05"), false},
		{Parse("2024-08-05"), false},
		{Parse("2025-08-05"), false},
		{Parse("2026-08-05"), false},
		{Parse("2027-08-05"), false},
		{Parse("2028-08-05"), false},
		{Parse("2029-08-05"), true},
		{Parse("2030-08-05"), false},
		{Parse("2031-08-05"), false},
	}

	for _, v := range Tests {
		output := v.input.IsYearOfRooster()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of rooster, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of rooster, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}

}

func TestCarbon_IsYearOfDog(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05"), false},
		{Parse("2021-08-05"), false},
		{Parse("2022-08-05"), false},
		{Parse("2023-08-05"), false},
		{Parse("2024-08-05"), false},
		{Parse("2025-08-05"), false},
		{Parse("2026-08-05"), false},
		{Parse("2027-08-05"), false},
		{Parse("2028-08-05"), false},
		{Parse("2029-08-05"), false},
		{Parse("2030-08-05"), true},
		{Parse("2031-08-05"), false},
	}

	for _, v := range Tests {
		output := v.input.IsYearOfDog()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of dog, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of dog, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}

}

func TestCarbon_IsYearOfPig(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output bool   // 期望输出值
	}{
		{Parse("2020-08-05"), false},
		{Parse("2021-08-05"), false},
		{Parse("2022-08-05"), false},
		{Parse("2023-08-05"), false},
		{Parse("2024-08-05"), false},
		{Parse("2025-08-05"), false},
		{Parse("2026-08-05"), false},
		{Parse("2027-08-05"), false},
		{Parse("2028-08-05"), false},
		{Parse("2029-08-05"), false},
		{Parse("2030-08-05"), false},
		{Parse("2031-08-05"), true},
	}

	for _, v := range Tests {
		output := v.input.IsYearOfPig()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of pig, expected true, but got false\n", v.input.ToDateString())
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of pig, expected false, but got true\n", v.input.ToDateString())
			}
		}
	}

}
