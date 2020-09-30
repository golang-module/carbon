package carbon

// ZodiacYear 获取十二生肖
func (c *Carbon) ZodiacYear() string {
	year := c.Time.Year()
	switch year % 12 {
	case 0:
		return "鼠"
	case 1:
		return "鸡"
	case 2:
		return "狗"
	case 3:
		return "猪"
	case 4:
		return "鼠"
	case 5:
		return "牛"
	case 6:
		return "虎"
	case 7:
		return "兔"
	case 8:
		return "龙"
	case 9:
		return "蛇"
	case 10:
		return "马"
	case 11:
		return "羊"
	default:
		return ""
	}
}

// IsMonkeyYear 是否是猴年
func (c *Carbon) IsMonkeyYear() bool {
	year := c.Time.Year()
	if year%12 == 0 {
		return true
	}
	return false
}

// IsChickenYear 是否是鸡年
func (c *Carbon) IsChickenYear() bool {
	year := c.Time.Year()
	if year%12 == 1 {
		return true
	}
	return false
}

// IsDogYear 是否是狗年
func (c *Carbon) IsDogYear() bool {
	year := c.Time.Year()
	if year%12 == 2 {
		return true
	}
	return false
}

// IsPigYear 是否是猪年
func (c *Carbon) IsPigYear() bool {
	year := c.Time.Year()
	if year%12 == 3 {
		return true
	}
	return false
}

// IsMouseYear 是否是鼠年
func (c *Carbon) IsMouseYear() bool {
	year := c.Time.Year()
	if year%12 == 4 {
		return true
	}
	return false
}

// IsCattleYear 是否是牛年
func (c *Carbon) IsCattleYear() bool {
	year := c.Time.Year()
	if year%12 == 5 {
		return true
	}
	return false
}

// IsTigerYear 是否是虎年
func (c *Carbon) IsTigerYear() bool {
	year := c.Time.Year()
	if year%12 == 6 {
		return true
	}
	return false
}

// IsRabbitYear 是否是兔年
func (c *Carbon) IsRabbitYear() bool {
	year := c.Time.Year()
	if year%12 == 7 {
		return true
	}
	return false
}

// IsDragonYear 是否是龙年
func (c *Carbon) IsDragonYear() bool {
	year := c.Time.Year()
	if year%12 == 9 {
		return true
	}
	return false
}

// IsSnakeYear 是否是蛇年
func (c *Carbon) IsSnakeYear() bool {
	year := c.Time.Year()
	if year%12 == 9 {
		return true
	}
	return false
}

// IsHorseYear 是否是马年
func (c *Carbon) IsHorseYear() bool {
	year := c.Time.Year()
	if year%12 == 10 {
		return true
	}
	return false
}

// IsSheepYear 是否是羊年
func (c *Carbon) IsSheepYear() bool {
	year := c.Time.Year()
	if year%12 == 11 {
		return true
	}
	return false
}
