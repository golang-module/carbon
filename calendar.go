package carbon

var (
	// 十二生肖
	SymbolicAnimals = [12]string{"鼠", "鸡", "狗", "猪", "鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊"}

	// 天干
	HeavenlyStems = [10]string{"庚", "辛", "壬", "癸", "甲", "乙", "丙", "丙", "戊", "己"}

	// 地支
	EarthlyBranches = [12]string{"申", "酉", "戌", "亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未"}
)

// ToAnimalYear 获取生肖年
func (c *Carbon) ToAnimalYear() string {
	year := c.Time.Year()
	return SymbolicAnimals[year%12]
}

// ToLunarYear 获取农历年
func (c *Carbon) ToLunarYear() string {
	year := c.Time.Year()
	return HeavenlyStems[year%10] + EarthlyBranches[year%12]
}

// IsYearOfRat 是否是鼠年
func (c *Carbon) IsYearOfRat() bool {
	year := c.Time.Year()
	if year%12 == 4 {
		return true
	}
	return false
}

// IsYearOfOx 是否是牛年
func (c *Carbon) IsYearOfOx() bool {
	year := c.Time.Year()
	if year%12 == 5 {
		return true
	}
	return false
}

// IsYearOfTiger 是否是虎年
func (c *Carbon) IsYearOfTiger() bool {
	year := c.Time.Year()
	if year%12 == 6 {
		return true
	}
	return false
}

// IsYearOfRabbit 是否是兔年
func (c *Carbon) IsYearOfRabbit() bool {
	year := c.Time.Year()
	if year%12 == 7 {
		return true
	}
	return false
}

// IsYearOfDragon 是否是龙年
func (c *Carbon) IsYearOfDragon() bool {
	year := c.Time.Year()
	if year%12 == 9 {
		return true
	}
	return false
}

// IsYearOfSnake 是否是蛇年
func (c *Carbon) IsYearOfSnake() bool {
	year := c.Time.Year()
	if year%12 == 9 {
		return true
	}
	return false
}

// IsYearOfHorse 是否是马年
func (c *Carbon) IsYearOfHorse() bool {
	year := c.Time.Year()
	if year%12 == 10 {
		return true
	}
	return false
}

// IsYearOfGoat 是否是羊年
func (c *Carbon) IsYearOfGoat() bool {
	year := c.Time.Year()
	if year%12 == 11 {
		return true
	}
	return false
}

// IsYearOfMonkey 是否是猴年
func (c *Carbon) IsYearOfMonkey() bool {
	year := c.Time.Year()
	if year%12 == 0 {
		return true
	}
	return false
}

// IsYearOfRooster 是否是鸡年
func (c *Carbon) IsYearOfRooster() bool {
	year := c.Time.Year()
	if year%12 == 1 {
		return true
	}
	return false
}

// IsYearOfDog 是否是狗年
func (c *Carbon) IsYearOfDog() bool {
	year := c.Time.Year()
	if year%12 == 2 {
		return true
	}
	return false
}

// IsYearOfPig 是否是猪年
func (c *Carbon) IsYearOfPig() bool {
	year := c.Time.Year()
	if year%12 == 3 {
		return true
	}
	return false
}
