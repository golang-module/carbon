package carbon

import "testing"

func TestCarbon_ToAnimalYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"0000-00-00", ""},
		{"2020-08-05", "鼠"},
		{"2021-08-05", "牛"},
		{"2022-08-05", "虎"},
		{"2023-08-05", "兔"},
		{"2024-08-05", "龙"},
		{"2025-08-05", "蛇"},
		{"2026-08-05", "马"},
		{"2027-08-05", "羊"},
		{"2028-08-05", "猴"},
		{"2029-08-05", "鸡"},
		{"2030-08-05", "狗"},
		{"2031-08-05", "猪"},
	}

	for _, v := range Tests {
		output := Parse(v.input).ToAnimalYear()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ToLunarYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"0000-00-00", ""},
		{"1861-08-22", "辛酉"}, // 辛酉政变发生日期
		{"1894-07-25", "甲午"}, // 甲午战争发生日期
		{"1898-06-11", "戊戌"}, // 戊戌变法发生日期
		{"1901-09-07", "辛丑"}, // 辛丑条约签署日期
		{"1911-10-10", "辛亥"}, // 辛亥革命发生日期
		{"1900-08-28", "庚子"}, // 庚子赔款发生日期
	}

	for _, v := range Tests {
		output := Parse(v.input).ToLunarYear()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_IsYearOfRat(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-08-05", true},
		{"2021-08-05", false},
		{"2022-08-05", false},
		{"2023-08-05", false},
		{"2024-08-05", false},
		{"2025-08-05", false},
		{"2026-08-05", false},
		{"2027-08-05", false},
		{"2028-08-05", false},
		{"2029-08-05", false},
		{"2030-08-05", false},
		{"2031-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsYearOfRat()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsYearOfOx(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-08-05", false},
		{"2021-08-05", true},
		{"2022-08-05", false},
		{"2023-08-05", false},
		{"2024-08-05", false},
		{"2025-08-05", false},
		{"2026-08-05", false},
		{"2027-08-05", false},
		{"2028-08-05", false},
		{"2029-08-05", false},
		{"2030-08-05", false},
		{"2031-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsYearOfOx()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsYearOfTiger(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-08-05", "false"},
		{"2021-08-05", "false"},
		{"2022-08-05", "true"},
		{"2023-08-05", "false"},
		{"2024-08-05", "false"},
		{"2025-08-05", "false"},
		{"2026-08-05", "false"},
		{"2027-08-05", "false"},
		{"2028-08-05", "false"},
		{"2029-08-05", "false"},
		{"2030-08-05", "false"},
		{"2031-08-05", "false"},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsYearOfTiger()

		expected := "false"
		if output == true {
			expected = "true"
		}

		if expected != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, expected)
		}
	}
}

func TestCarbon_IsYearOfRabbit(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-08-05", false},
		{"2021-08-05", false},
		{"2022-08-05", false},
		{"2023-08-05", true},
		{"2024-08-05", false},
		{"2025-08-05", false},
		{"2026-08-05", false},
		{"2027-08-05", false},
		{"2028-08-05", false},
		{"2029-08-05", false},
		{"2030-08-05", false},
		{"2031-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsYearOfRabbit()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsYearOfDragon(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-08-05", false},
		{"2021-08-05", false},
		{"2022-08-05", false},
		{"2023-08-05", false},
		{"2024-08-05", true},
		{"2025-08-05", false},
		{"2026-08-05", false},
		{"2027-08-05", false},
		{"2028-08-05", false},
		{"2029-08-05", false},
		{"2030-08-05", false},
		{"2031-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsYearOfDragon()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsYearOfSnake(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-08-05", false},
		{"2021-08-05", false},
		{"2022-08-05", false},
		{"2023-08-05", false},
		{"2024-08-05", false},
		{"2025-08-05", true},
		{"2026-08-05", false},
		{"2027-08-05", false},
		{"2028-08-05", false},
		{"2029-08-05", false},
		{"2030-08-05", false},
		{"2031-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsYearOfSnake()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsYearOfHorse(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-08-05", false},
		{"2021-08-05", false},
		{"2022-08-05", false},
		{"2023-08-05", false},
		{"2024-08-05", false},
		{"2025-08-05", false},
		{"2026-08-05", true},
		{"2027-08-05", false},
		{"2028-08-05", false},
		{"2029-08-05", false},
		{"2030-08-05", false},
		{"2031-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsYearOfHorse()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsYearOfGoat(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-08-05", false},
		{"2021-08-05", false},
		{"2022-08-05", false},
		{"2023-08-05", false},
		{"2024-08-05", false},
		{"2025-08-05", false},
		{"2026-08-05", false},
		{"2027-08-05", true},
		{"2028-08-05", false},
		{"2029-08-05", false},
		{"2030-08-05", false},
		{"2031-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsYearOfGoat()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsYearOfMonkey(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-08-05", false},
		{"2021-08-05", false},
		{"2022-08-05", false},
		{"2023-08-05", false},
		{"2024-08-05", false},
		{"2025-08-05", false},
		{"2026-08-05", false},
		{"2027-08-05", false},
		{"2028-08-05", true},
		{"2029-08-05", false},
		{"2030-08-05", false},
		{"2031-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsYearOfMonkey()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsYearOfRooster(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-08-05", false},
		{"2021-08-05", false},
		{"2022-08-05", false},
		{"2023-08-05", false},
		{"2024-08-05", false},
		{"2025-08-05", false},
		{"2026-08-05", false},
		{"2027-08-05", false},
		{"2028-08-05", false},
		{"2029-08-05", true},
		{"2030-08-05", false},
		{"2031-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsYearOfRooster()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsYearOfDog(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-08-05", false},
		{"2021-08-05", false},
		{"2022-08-05", false},
		{"2023-08-05", false},
		{"2024-08-05", false},
		{"2025-08-05", false},
		{"2026-08-05", false},
		{"2027-08-05", false},
		{"2028-08-05", false},
		{"2029-08-05", false},
		{"2030-08-05", true},
		{"2031-08-05", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsYearOfDog()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}

func TestCarbon_IsYearOfPig(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"2020-08-05", false},
		{"2021-08-05", false},
		{"2022-08-05", false},
		{"2023-08-05", false},
		{"2024-08-05", false},
		{"2025-08-05", false},
		{"2026-08-05", false},
		{"2027-08-05", false},
		{"2028-08-05", false},
		{"2029-08-05", false},
		{"2030-08-05", false},
		{"2031-08-05", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).IsYearOfPig()

		if output != v.output {
			expected := "false"
			if v.output == true {
				expected = "true"
			}

			reality := "false"
			if output == true {
				reality = "true"
			}
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, expected, reality)
		}
	}
}
