package lunar

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-module/carbon/v2/calendar"
)

var (
	minYear, maxYear = 1900, 2100
	lunarNumbers     = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	lunarMonths      = []string{"正", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "腊"}
	lunarTimes       = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	lunarAnimals     = []string{"猴", "鸡", "狗", "猪", "鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊"}
	lunarFestivals   = []string{"春节", "元宵节", "端午节", "七夕节", "中元节", "中秋节", "重阳节", "寒衣节", "下元节", "腊八节", "小年"}

	lunarTerms = []int{
		0x04bd8, 0x04ae0, 0x0a570, 0x054d5, 0x0d260, 0x0d950, 0x16554, 0x056a0, 0x09ad0, 0x055d2, // 1900-1909
		0x04ae0, 0x0a5b6, 0x0a4d0, 0x0d250, 0x1d255, 0x0b540, 0x0d6a0, 0x0ada2, 0x095b0, 0x14977, // 1910-1919
		0x04970, 0x0a4b0, 0x0b4b5, 0x06a50, 0x06d40, 0x1ab54, 0x02b60, 0x09570, 0x052f2, 0x04970, // 1920-1929
		0x06566, 0x0d4a0, 0x0ea50, 0x16a95, 0x05ad0, 0x02b60, 0x186e3, 0x092e0, 0x1c8d7, 0x0c950, // 1930-1939
		0x0d4a0, 0x1d8a6, 0x0b550, 0x056a0, 0x1a5b4, 0x025d0, 0x092d0, 0x0d2b2, 0x0a950, 0x0b557, // 1940-1949
		0x06ca0, 0x0b550, 0x15355, 0x04da0, 0x0a5d0, 0x14573, 0x052d0, 0x0a9a8, 0x0e950, 0x06aa0, // 1950-1959
		0x0aea6, 0x0ab50, 0x04b60, 0x0aae4, 0x0a570, 0x05260, 0x0f263, 0x0d950, 0x05b57, 0x056a0, // 1960-1969
		0x096d0, 0x04dd5, 0x04ad0, 0x0a4d0, 0x0d4d4, 0x0d250, 0x0d558, 0x0b540, 0x0b5a0, 0x195a6, // 1970-1979
		0x095b0, 0x049b0, 0x0a974, 0x0a4b0, 0x0b27a, 0x06a50, 0x06d40, 0x0af46, 0x0ab60, 0x09570, // 1980-1989
		0x04af5, 0x04970, 0x064b0, 0x074a3, 0x0ea50, 0x06b58, 0x05ac0, 0x0ab60, 0x096d5, 0x092e0, // 1990-1999
		0x0c960, 0x0d954, 0x0d4a0, 0x0da50, 0x07552, 0x056a0, 0x0abb7, 0x025d0, 0x092d0, 0x0cab5, // 2000-2009
		0x0a950, 0x0b4a0, 0x0baa4, 0x0ad50, 0x055d9, 0x04ba0, 0x0a5b0, 0x15176, 0x052b0, 0x0a930, // 2010-2019
		0x07954, 0x06aa0, 0x0ad50, 0x05b52, 0x04b60, 0x0a6e6, 0x0a4e0, 0x0d260, 0x0ea65, 0x0d530, // 2020-2029
		0x05aa0, 0x076a3, 0x096d0, 0x04bd7, 0x04ad0, 0x0a4d0, 0x1d0b6, 0x0d250, 0x0d520, 0x0dd45, // 2030-2039
		0x0b5a0, 0x056d0, 0x055b2, 0x049b0, 0x0a577, 0x0a4b0, 0x0aa50, 0x1b255, 0x06d20, 0x0ada0, // 2040-2049
		0x14b63, 0x09370, 0x049f8, 0x04970, 0x064b0, 0x168a6, 0x0ea50, 0x06b20, 0x1a6c4, 0x0aae0, // 2050-2059
		0x0a2e0, 0x0d2e3, 0x0c960, 0x0d557, 0x0d4a0, 0x0da50, 0x05d55, 0x056a0, 0x0a6d0, 0x055d4, // 2060-2069
		0x052d0, 0x0a9b8, 0x0a950, 0x0b4a0, 0x0b6a6, 0x0ad50, 0x055a0, 0x0aba4, 0x0a5b0, 0x052b0, // 2070-2079
		0x0b273, 0x06930, 0x07337, 0x06aa0, 0x0ad50, 0x14b55, 0x04b60, 0x0a570, 0x054e4, 0x0d160, // 2080-2089
		0x0e968, 0x0d520, 0x0daa0, 0x16aa6, 0x056d0, 0x04ae0, 0x0a9d4, 0x0a2d0, 0x0d150, 0x0f252, // 2090-2099
		0x0d520, // 2100
	}

	invalidYearError = func() error {
		return fmt.Errorf("invalid year, currently only 200 years from 1900 to 2100 are supported")
	}
)

// Calendar defines a Calendar struct.
// 定义 Calendar 结构体
type Calendar struct {
	Solar Solar
	Lunar Lunar
	Error error
}

// Solar defines a Solar struct.
// 定义 Solar 结构体
type Solar struct {
	calendar.Solar
}

// Lunar defines a Lunar struct.
// 定义 Lunar 结构体
type Lunar struct {
	year, month, day, hour, minute, second int
	isLeapMonth                            bool
}

// NewSolar returns a new Solar instance.
// 初始化 Solar 结构体
func NewSolar(t time.Time) (s Solar) {
	s.Time = t
	return s
}

// CreateFromSolar creates a Calendar instance from Solar.
// 从 公历 创建 Calendar 实例
func CreateFromSolar(s Solar) (c Calendar) {
	// leapMonths:闰月总数，daysOfYear:年天数，daysOfMonth:月天数，leapMonth:闰月月份
	daysInYear, daysInMonth, leapMonth := 365, 30, 0
	year := s.Year()
	if year < minYear || year > maxYear {
		c.Error = invalidYearError()
		return c
	}
	// 与 1900-01-31 相差多少天
	offset := s.diffInDays(time.Date(minYear, 1, 31, 0, 0, 0, 0, s.Location()))
	for c.Lunar.year = minYear; c.Lunar.year <= maxYear && offset > 0; c.Lunar.year++ {
		daysInYear = c.Lunar.getDaysInYear()
		offset -= daysInYear
	}
	if offset < 0 {
		offset += daysInYear
		c.Lunar.year--
	}
	leapMonth = c.LeapMonth()
	for c.Lunar.month = 1; c.Lunar.month <= 12 && offset > 0; c.Lunar.month++ {
		if leapMonth > 0 && c.Lunar.month == (leapMonth+1) && !c.Lunar.isLeapMonth {
			c.Lunar.month--
			c.Lunar.isLeapMonth = true
			daysInMonth = c.Lunar.getDaysInLeapMonth()
		} else {
			daysInMonth = c.Lunar.getDaysInMonth()
		}
		offset -= daysInMonth
		if c.Lunar.isLeapMonth && c.Lunar.month == (leapMonth+1) {
			c.Lunar.isLeapMonth = false
		}
	}
	// offset为0时，并且刚才计算的月份是闰月，要校正
	if offset == 0 && leapMonth > 0 && c.Lunar.month == leapMonth+1 {
		if c.Lunar.isLeapMonth {
			c.Lunar.isLeapMonth = false
		} else {
			c.Lunar.isLeapMonth = true
			c.Lunar.month--
		}
	}
	// offset小于0时，也要校正
	if offset < 0 {
		offset += daysInMonth
		c.Lunar.month--
	}
	c.Lunar.day = offset + 1
	c.Lunar.hour, c.Lunar.minute, c.Lunar.second = s.Clock()
	c.Solar = s
	return c
}

// diffInDays gets the difference in days.
// 获取相差天数
func (s Solar) diffInDays(t time.Time) int {
	return int(s.Time.Sub(t).Hours() / 24)
}

// NewLunar returns a new Lunar instance.
// 初始化 Lunar 结构体
func NewLunar(year, month, day, hour, minute, second int, isLeapMonth bool) (l Lunar) {
	l.year, l.month, l.day = year, month, day
	l.hour, l.minute, l.second = hour, minute, second
	l.isLeapMonth = isLeapMonth
	return l
}

// CreateFromLunar creates a Calendar instance from Lunar instance.
// 从 农历 创建 Calendar 实例
func CreateFromLunar(l Lunar) (c Calendar) {
	if l.year < minYear || l.year > maxYear {
		c.Error = invalidYearError()
		return c
	}

	days := l.getDaysInMonth()
	offset := l.getOffsetInYear()
	offset += l.getOffsetInMonth()
	// 转换闰月农历 需补充该年闰月的前一个月的时差
	if l.isLeapMonth {
		offset += days
	}

	ts := (offset+l.day)*86400 + -2206512000 + l.hour*3600 + l.minute*60 + l.second
	c.Solar.Time = time.Unix(int64(ts), 0)
	c.Lunar = l
	return c
}

func (l Lunar) getOffsetInYear() int {
	flag := false
	clone, month, offset := l, 0, 0
	for month = 1; month < l.month; month++ {
		leapMonth := l.getLeapMonth()
		if !flag {
			// 处理闰月
			if leapMonth <= month && leapMonth > 0 {
				offset += l.getDaysInLeapMonth()
				flag = true
			}
		}
		clone.month = month
		offset += clone.getDaysInMonth()
	}
	return offset
}

func (l Lunar) getOffsetInMonth() int {
	clone, year, offset := l, 0, 0
	for year = minYear; year < l.year; year++ {
		clone.year = year
		offset += clone.getDaysInYear()
	}
	return offset
}

func (l Lunar) getDaysInYear() int {
	var days = 348
	for i := 0x8000; i > 0x8; i >>= 1 {
		if (lunarTerms[l.year-minYear] & i) != 0 {
			days++
		}
	}
	return days + l.getDaysInLeapMonth()
}

func (l Lunar) getDaysInMonth() int {
	if (lunarTerms[l.year-minYear] & (0x10000 >> uint(l.month))) != 0 {
		return 30
	}
	return 29
}

func (l Lunar) getDaysInLeapMonth() int {
	if l.getLeapMonth() == 0 {
		return 0
	}
	if lunarTerms[l.year-minYear]&0x10000 != 0 {
		return 30
	}
	return 29
}

func (l Lunar) getLeapMonth() int {
	return lunarTerms[l.year-minYear] & 0xf
}

// Animal gets lunar animal name like "猴".
// 获取生肖
func (c Calendar) Animal() string {
	if c.Error != nil {
		return ""
	}
	return lunarAnimals[c.Lunar.year%calendar.MonthsPerYear]
}

// Festival gets lunar festival name like "春节".
// 获取农历节日
func (c Calendar) Festival() (festival string) {
	if c.Error != nil {
		return
	}
	month, day := c.Lunar.month, c.Lunar.day
	switch {
	case month == 1 && day == 1:
		festival = lunarFestivals[0]
	case month == 1 && day == 15:
		festival = lunarFestivals[1]
	case month == 5 && day == 5:
		festival = lunarFestivals[2]
	case month == 7 && day == 7:
		festival = lunarFestivals[3]
	case month == 7 && day == 15:
		festival = lunarFestivals[4]
	case month == 8 && day == 15:
		festival = lunarFestivals[5]
	case month == 9 && day == 9:
		festival = lunarFestivals[6]
	case month == 10 && day == 1:
		festival = lunarFestivals[7]
	case month == 10 && day == 15:
		festival = lunarFestivals[8]
	case month == 12 && day == 8:
		festival = lunarFestivals[9]
	case month == 12 && day == 23:
		festival = lunarFestivals[10]
	}
	return
}

// DateTime gets lunar year, month, day, hour, minute, and second like 2020, 8, 5, 13, 14, 15.
// 获取农历年月日时分秒
func (c Calendar) DateTime() (year, month, day, hour, minute, second int) {
	if c.Error != nil {
		return
	}
	return c.Lunar.year, c.Lunar.month, c.Lunar.day, c.Lunar.hour, c.Lunar.minute, c.Lunar.second
}

// Date gets lunar year, month and day like 2020, 8, 5.
// 获取农历时分秒
func (c Calendar) Date() (year, month, day int) {
	if c.Error != nil {
		return
	}
	return c.Lunar.year, c.Lunar.month, c.Lunar.day
}

// Time gets lunar hour, minute, and second like 13, 14, 15.
// 获取农历时分秒
func (c Calendar) Time() (hour, minute, second int) {
	if c.Error != nil {
		return
	}
	return c.Lunar.hour, c.Lunar.minute, c.Lunar.second
}

// Year gets lunar year like 2020.
// 获取农历年
func (c Calendar) Year() int {
	if c.Error != nil {
		return 0
	}
	return c.Lunar.year
}

// Month gets lunar month like 8.
// 获取农历月
func (c Calendar) Month() int {
	if c.Error != nil {
		return 0
	}
	return c.Lunar.month
}

// LeapMonth gets lunar leap month like 8.
// 获取农历闰月月份
func (c Calendar) LeapMonth() int {
	if c.Error != nil {
		return 0
	}
	return c.Lunar.getLeapMonth()
}

// Day gets lunar day like 5.
// 获取农历日
func (c Calendar) Day() int {
	if c.Error != nil {
		return 0
	}
	return c.Lunar.day
}

// ToYearString outputs a string in lunar year format like "二零二零".
// 获取农历年字符串
func (c Calendar) ToYearString() string {
	if c.Error != nil {
		return ""
	}
	year := fmt.Sprintf("%d", c.Lunar.year)
	for i, replace := range lunarNumbers {
		year = strings.Replace(year, fmt.Sprintf("%d", i), replace, -1)
	}
	return year
}

// ToMonthString outputs a string in lunar month format like "正月".
// 获取农历月字符串
func (c Calendar) ToMonthString() string {
	if c.Error != nil {
		return ""
	}
	month := lunarMonths[c.Lunar.month-1] + "月"
	if c.IsLeapMonth() {
		return "闰" + month
	}
	return month
}

// ToDayString outputs a string in lunar day format like "廿一".
// 获取农历日字符串
func (c Calendar) ToDayString() (day string) {
	if c.Error != nil {
		return
	}
	num := lunarNumbers[c.Lunar.day%10]
	switch {
	case c.Lunar.day == 30:
		day = "三十"
	case c.Lunar.day > 20:
		day = "廿" + num
	case c.Lunar.day == 20:
		day = "二十"
	case c.Lunar.day > 10:
		day = "十" + num
	case c.Lunar.day == 10:
		day = "初十"
	case c.Lunar.day < 10:
		day = "初" + num
	}
	return
}

// ToDateString outputs a string in lunar date format like "二零二零年腊月初五".
// 获取农历日期字符串
func (c Calendar) ToDateString() string {
	if c.Error != nil {
		return ""
	}
	return c.ToYearString() + "年" + c.ToMonthString() + c.ToDayString()
}

// String outputs a string in YYYY-MM-DD HH::ii::ss format, implement Stringer interface.
// 输出 YYYY-MM-DD HH::ii::ss 格式字符串，实现 Stringer 接口
func (c Calendar) String() string {
	if c.Error != nil {
		return ""
	}
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", c.Lunar.year, c.Lunar.month, c.Lunar.day, c.Lunar.hour, c.Lunar.minute, c.Lunar.second)
}

// IsLeapYear reports whether is leap year.
// 是否是闰年
func (c Calendar) IsLeapYear() bool {
	if c.Error != nil {
		return false
	}
	return c.LeapMonth() != 0
}

// IsLeapMonth reports whether is leap month.
// 是否是闰月
func (c Calendar) IsLeapMonth() bool {
	if c.Error != nil {
		return false
	}
	return c.Lunar.month == c.LeapMonth()
}

// IsRatYear reports whether is year of Rat.
// 是否是鼠年
func (c Calendar) IsRatYear() bool {
	if c.Error != nil {
		return false
	}
	if c.Lunar.year%calendar.MonthsPerYear == 4 {
		return true
	}
	return false
}

// IsOxYear reports whether is year of Ox.
// 是否是牛年
func (c Calendar) IsOxYear() bool {
	if c.Error != nil {
		return false
	}
	if c.Lunar.year%calendar.MonthsPerYear == 5 {
		return true
	}
	return false
}

// IsTigerYear reports whether is year of Tiger.
// 是否是虎年
func (c Calendar) IsTigerYear() bool {
	if c.Error != nil {
		return false
	}
	if c.Lunar.year%calendar.MonthsPerYear == 6 {
		return true
	}
	return false
}

// IsRabbitYear reports whether is year of Rabbit.
// 是否是兔年
func (c Calendar) IsRabbitYear() bool {
	if c.Error != nil {
		return false
	}
	if c.Lunar.year%calendar.MonthsPerYear == 7 {
		return true
	}
	return false
}

// IsDragonYear reports whether is year of Dragon.
// 是否是龙年
func (c Calendar) IsDragonYear() bool {
	if c.Error != nil {
		return false
	}
	if c.Lunar.year%calendar.MonthsPerYear == 8 {
		return true
	}
	return false
}

// IsSnakeYear reports whether is year of Snake.
// 是否是蛇年
func (c Calendar) IsSnakeYear() bool {
	if c.Error != nil {
		return false
	}
	if c.Lunar.year%calendar.MonthsPerYear == 9 {
		return true
	}
	return false
}

// IsHorseYear reports whether is year of Horse.
// 是否是马年
func (c Calendar) IsHorseYear() bool {
	if c.Error != nil {
		return false
	}
	if c.Lunar.year%calendar.MonthsPerYear == 10 {
		return true
	}
	return false
}

// IsGoatYear reports whether is year of Goat.
// 是否是羊年
func (c Calendar) IsGoatYear() bool {
	if c.Error != nil {
		return false
	}
	if c.Lunar.year%calendar.MonthsPerYear == 11 {
		return true
	}
	return false
}

// IsMonkeyYear reports whether is year of Monkey.
// 是否是猴年
func (c Calendar) IsMonkeyYear() bool {
	if c.Error != nil {
		return false
	}
	if c.Lunar.year%calendar.MonthsPerYear == 0 {
		return true
	}
	return false
}

// IsRoosterYear reports whether is year of Rooster.
// 是否是鸡年
func (c Calendar) IsRoosterYear() bool {
	if c.Error != nil {
		return false
	}
	if c.Lunar.year%calendar.MonthsPerYear == 1 {
		return true
	}
	return false
}

// IsDogYear reports whether is year of Dog.
// 是否是狗年
func (c Calendar) IsDogYear() bool {
	if c.Error != nil {
		return false
	}
	if c.Lunar.year%calendar.MonthsPerYear == 2 {
		return true
	}
	return false
}

// IsPigYear reports whether is year of Pig.
// 是否是猪年
func (c Calendar) IsPigYear() bool {
	if c.Error != nil {
		return false
	}
	if c.Lunar.year%calendar.MonthsPerYear == 3 {
		return true
	}
	return false
}

// DoubleHour gets double-hour name like "子时".
// 获取当前时辰
func (c Calendar) DoubleHour() (dh string) {
	if c.Error != nil {
		return ""
	}
	hour, minute := c.Lunar.hour, c.Lunar.minute
	switch {
	case hour >= 23, hour == 0 && minute <= 59:
		dh = lunarTimes[0] // FirstDoubleHour
	case hour >= 1 && hour < 3:
		dh = lunarTimes[1] // SecondDoubleHour
	case hour >= 3 && hour < 5:
		dh = lunarTimes[2] // ThirdDoubleHour
	case hour >= 5 && hour < 7:
		dh = lunarTimes[3] // FourthDoubleHour
	case hour >= 7 && hour < 9:
		dh = lunarTimes[4] // FifthDoubleHour
	case hour >= 9 && hour < 11:
		dh = lunarTimes[5] // SixthDoubleHour
	case hour >= 11 && hour < 13:
		dh = lunarTimes[6] // SeventhDoubleHour
	case hour >= 13 && hour < 15:
		dh = lunarTimes[7] // EighthDoubleHour
	case hour >= 15 && hour < 17:
		dh = lunarTimes[8] // NinthDoubleHour
	case hour >= 17 && hour < 19:
		dh = lunarTimes[9] // TenthDoubleHour
	case hour >= 19 && hour < 21:
		dh = lunarTimes[10] // EleventhDoubleHour
	case hour >= 21 && hour < 23:
		dh = lunarTimes[11] // TwelfthDoubleHour
	}
	return dh + "时"
}

// IsFirstDoubleHour reports whether is FirstDoubleHour.
// 是否是子时
func (c Calendar) IsFirstDoubleHour() bool {
	if c.Error != nil {
		return false
	}
	hour, minute := c.Lunar.hour, c.Lunar.minute
	if hour >= 23 {
		return true
	}
	if hour == 0 && minute <= 59 {
		return true
	}
	return false
}

// IsSecondDoubleHour reports whether is SecondDoubleHour.
// 是否是丑时
func (c Calendar) IsSecondDoubleHour() bool {
	if c.Error != nil {
		return false
	}
	hour := c.Lunar.hour
	if hour >= 1 && hour < 3 {
		return true
	}
	return false
}

// IsThirdDoubleHour reports whether is ThirdDoubleHour.
// 是否是寅时
func (c Calendar) IsThirdDoubleHour() bool {
	if c.Error != nil {
		return false
	}
	hour := c.Lunar.hour
	if hour >= 3 && hour < 5 {
		return true
	}
	return false
}

// IsFourthDoubleHour reports whether is FourthDoubleHour.
// 是否是卯时
func (c Calendar) IsFourthDoubleHour() bool {
	if c.Error != nil {
		return false
	}
	hour := c.Lunar.hour
	if hour >= 5 && hour < 7 {
		return true
	}
	return false
}

// IsFifthDoubleHour reports whether is FifthDoubleHour.
// 是否是辰时
func (c Calendar) IsFifthDoubleHour() bool {
	if c.Error != nil {
		return false
	}
	hour := c.Lunar.hour
	if hour >= 7 && hour < 9 {
		return true
	}
	return false
}

// IsSixthDoubleHour reports whether is SixthDoubleHour.
// 是否是巳时
func (c Calendar) IsSixthDoubleHour() bool {
	if c.Error != nil {
		return false
	}
	hour := c.Lunar.hour
	if hour >= 9 && hour < 11 {
		return true
	}
	return false
}

// IsSeventhDoubleHour reports whether is SeventhDoubleHour.
// 是否是午时
func (c Calendar) IsSeventhDoubleHour() bool {
	if c.Error != nil {
		return false
	}
	hour := c.Lunar.hour
	if hour >= 11 && hour < 13 {
		return true
	}
	return false
}

// IsEighthDoubleHour reports whether is EighthDoubleHour.
// 是否是未时
func (c Calendar) IsEighthDoubleHour() bool {
	if c.Error != nil {
		return false
	}
	hour := c.Lunar.hour
	if hour >= 13 && hour < 15 {
		return true
	}
	return false
}

// IsNinthDoubleHour reports whether is NinthDoubleHour.
// 是否是申时
func (c Calendar) IsNinthDoubleHour() bool {
	if c.Error != nil {
		return false
	}
	hour := c.Lunar.hour
	if hour >= 15 && hour < 17 {
		return true
	}
	return false
}

// IsTenthDoubleHour reports whether is TenthDoubleHour.
// 是否是酉时
func (c Calendar) IsTenthDoubleHour() bool {
	if c.Error != nil {
		return false
	}
	hour := c.Lunar.hour
	if hour >= 17 && hour < 19 {
		return true
	}
	return false
}

// IsEleventhDoubleHour reports whether is EleventhDoubleHour.
// 是否是戌时
func (c Calendar) IsEleventhDoubleHour() bool {
	if c.Error != nil {
		return false
	}
	hour := c.Lunar.hour
	if hour >= 19 && hour < 21 {
		return true
	}
	return false
}

// IsTwelfthDoubleHour reports whether is TwelfthDoubleHour.
// 是否是亥时
func (c Calendar) IsTwelfthDoubleHour() bool {
	if c.Error != nil {
		return false
	}
	hour := c.Lunar.hour
	if hour >= 21 && hour < 23 {
		return true
	}
	return false
}
