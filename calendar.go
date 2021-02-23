package carbon

var (
	// 生肖
	SymbolicAnimals = [12]string{"猴", "鸡", "狗", "猪", "鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊"}

	// 天干
	HeavenlyStems = [10]string{"庚", "辛", "壬", "癸", "甲", "乙", "丙", "丁", "戊", "己"}

	// 地支
	EarthlyBranches = [12]string{"申", "酉", "戌", "亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未"}
)

// ToAnimalYear 获取生肖年
func (c Carbon) ToAnimalYear() string {
	if c.IsZero() {
		return ""
	}
	return SymbolicAnimals[c.Year()%MonthsPerYear]
}

// ToLunarYear 获取农历年
func (c Carbon) ToLunarYear() string {
	if c.IsZero() {
		return ""
	}
	return HeavenlyStems[c.Year()%YearsPerDecade] + EarthlyBranches[c.Year()%MonthsPerYear]
}

// IsYearOfRat 是否是鼠年
func (c Carbon) IsYearOfRat() bool {
	if c.Year()%MonthsPerYear == 4 {
		return true
	}
	return false
}

// IsYearOfOx 是否是牛年
func (c Carbon) IsYearOfOx() bool {
	if c.Year()%MonthsPerYear == 5 {
		return true
	}
	return false
}

// IsYearOfTiger 是否是虎年
func (c Carbon) IsYearOfTiger() bool {
	if c.Year()%MonthsPerYear == 6 {
		return true
	}
	return false
}

// IsYearOfRabbit 是否是兔年
func (c Carbon) IsYearOfRabbit() bool {
	if c.Year()%MonthsPerYear == 7 {
		return true
	}
	return false
}

// IsYearOfDragon 是否是龙年
func (c Carbon) IsYearOfDragon() bool {
	if c.Year()%MonthsPerYear == 8 {
		return true
	}
	return false
}

// IsYearOfSnake 是否是蛇年
func (c Carbon) IsYearOfSnake() bool {
	if c.Year()%MonthsPerYear == 9 {
		return true
	}
	return false
}

// IsYearOfHorse 是否是马年
func (c Carbon) IsYearOfHorse() bool {
	if c.Year()%MonthsPerYear == 10 {
		return true
	}
	return false
}

// IsYearOfGoat 是否是羊年
func (c Carbon) IsYearOfGoat() bool {
	if c.Year()%MonthsPerYear == 11 {
		return true
	}
	return false
}

// IsYearOfMonkey 是否是猴年
func (c Carbon) IsYearOfMonkey() bool {
	if c.Year()%MonthsPerYear == 0 {
		return true
	}
	return false
}

// IsYearOfRooster 是否是鸡年
func (c Carbon) IsYearOfRooster() bool {
	if c.Year()%MonthsPerYear == 1 {
		return true
	}
	return false
}

// IsYearOfDog 是否是狗年
func (c Carbon) IsYearOfDog() bool {
	if c.Year()%MonthsPerYear == 2 {
		return true
	}
	return false
}

// IsYearOfPig 是否是猪年
func (c Carbon) IsYearOfPig() bool {
	if c.Year()%MonthsPerYear == 3 {
		return true
	}
	return false
}
