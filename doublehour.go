package carbon

import (
	"strings"
)

// 获取当前时辰
func (c Carbon) DoubleHour() string {
	if c.IsInvalid() {
		return ""
	}
	if len(c.lang.resources) == 0 {
		c.lang.SetLocale(defaultLocale)
	}

	index := -1
	hour, minute, _ := c.Time()
	switch {
	case hour >= 23, hour == 0 && minute <= 59:
		index = 0 // 子时
	case hour >= 1 && hour < 3:
		index = 1 // 丑时
	case hour >= 3 && hour < 5:
		index = 2 //寅时
	case hour >= 5 && hour < 7:
		index = 3 //卯时
	case hour >= 7 && hour < 9:
		index = 4 // 辰时
	case hour >= 9 && hour < 11:
		index = 5 //巳时
	case hour >= 11 && hour < 13:
		index = 6 // 午时
	case hour >= 13 && hour < 15:
		index = 7 //未时
	case hour >= 15 && hour < 17:
		index = 8 // 申时
	case hour >= 17 && hour < 19:
		index = 9 //酉时
	case hour >= 19 && hour < 21:
		index = 10 //戌时
	case hour >= 21 && hour < 23:
		index = 11 //亥时
	}
	if doubleHours, ok := c.lang.resources["double_hours"]; ok {
		hours := strings.Split(doubleHours, "|")
		if len(hours) == 12 {
			return hours[index]
		}
	}
	return ""
}

// IsFirstDoubleHour reports whether is FirstDoubleHour
// 是否是子时
func (c Carbon) IsFirstDoubleHour() bool {
	if c.IsInvalid() {
		return false
	}
	hour, minute, _ := c.Time()
	if hour >= 23 {
		return true
	}
	if hour == 0 && minute <= 59 {
		return true
	}
	return false
}

// IsSecondDoubleHour reports whether is SecondDoubleHour
// 是否是丑时
func (c Carbon) IsSecondDoubleHour() bool {
	if c.IsInvalid() {
		return false
	}
	hour, _, _ := c.Time()
	if hour >= 1 && hour < 3 {
		return true
	}
	return false
}

// IsThirdDoubleHour reports whether is ThirdDoubleHour
// 是否是寅时
func (c Carbon) IsThirdDoubleHour() bool {
	if c.IsInvalid() {
		return false
	}
	hour, _, _ := c.Time()
	if hour >= 3 && hour < 5 {
		return true
	}
	return false
}

// IsFourthDoubleHour reports whether is FourthDoubleHour
// 是否是卯时
func (c Carbon) IsFourthDoubleHour() bool {
	if c.IsInvalid() {
		return false
	}
	hour, _, _ := c.Time()
	if hour >= 5 && hour < 7 {
		return true
	}
	return false
}

// IsFifthDoubleHour reports whether is FifthDoubleHour
// 是否是辰时
func (c Carbon) IsFifthDoubleHour() bool {
	if c.IsInvalid() {
		return false
	}
	hour, _, _ := c.Time()
	if hour >= 7 && hour < 9 {
		return true
	}
	return false
}

// IsSixthDoubleHour reports whether is SixthDoubleHour
// 是否是巳时
func (c Carbon) IsSixthDoubleHour() bool {
	if c.IsInvalid() {
		return false
	}
	hour, _, _ := c.Time()
	if hour >= 9 && hour < 11 {
		return true
	}
	return false
}

// IsSeventhDoubleHour reports whether is SeventhDoubleHour
// 是否是午时
func (c Carbon) IsSeventhDoubleHour() bool {
	if c.IsInvalid() {
		return false
	}
	hour, _, _ := c.Time()
	if hour >= 11 && hour < 13 {
		return true
	}
	return false
}

// IsEighthDoubleHour reports whether is EighthDoubleHour
// 是否是未时
func (c Carbon) IsEighthDoubleHour() bool {
	if c.IsInvalid() {
		return false
	}
	hour, _, _ := c.Time()
	if hour >= 13 && hour < 15 {
		return true
	}
	return false
}

// IsNinthDoubleHour reports whether is NinthDoubleHour
// 是否是申时
func (c Carbon) IsNinthDoubleHour() bool {
	if c.IsInvalid() {
		return false
	}
	hour, _, _ := c.Time()
	if hour >= 15 && hour < 17 {
		return true
	}
	return false
}

// IsTenthDoubleHour reports whether is TenthDoubleHour
// 是否是酉时
func (c Carbon) IsTenthDoubleHour() bool {
	if c.IsInvalid() {
		return false
	}
	hour, _, _ := c.Time()
	if hour >= 17 && hour < 19 {
		return true
	}
	return false
}

// IsEleventhDoubleHour reports whether is EleventhDoubleHour
// 是否是戌时
func (c Carbon) IsEleventhDoubleHour() bool {
	if c.IsInvalid() {
		return false
	}
	hour, _, _ := c.Time()
	if hour >= 19 && hour < 21 {
		return true
	}
	return false
}

// IsTwelfthDoubleHour reports whether is TwelfthDoubleHour
// 是否是亥时
func (c Carbon) IsTwelfthDoubleHour() bool {
	if c.IsInvalid() {
		return false
	}
	hour, _, _ := c.Time()
	if hour >= 21 && hour < 23 {
		return true
	}
	return false
}
