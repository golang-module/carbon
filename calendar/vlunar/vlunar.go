package vlunar

import (
	"fmt"
	"math"
	"time"

	"github.com/golang-module/carbon/v2/calendar"
)

// see document and algorithum
// https://www.informatik.uni-leipzig.de/~duc/amlich/calrules_en.html
// https://www.informatik.uni-leipzig.de/~duc/amlich/calrules_en.html#comparison

type Festival struct {
	Day   int
	Month int
	Name  string
}

type LuckyHour struct {
	Chi  string
	From int
	To   int
}

type SolarTerm struct {
	Longitude int
	Name      string
}

// Solar defines a Solar struct.
type Gregorian struct {
	calendar.Gregorian
}

// Lunar defines a Lunar struct.
type Lunar struct {
	year, month, day, hour, minute, second, timeZone int
	zoneName                                         string
	leapMonth                                        int
	isLeapYear                                       int //0, 1
	juliusDay                                        int
	Error                                            error
}

var (
	PI         = math.Pi
	heavenStem = []string{"Giáp", "Ất", "Bính", "Đinh", "Mậu", "Kỷ", "Canh", "Tân", "Nhâm", "Quý"}
	lunarTimes = []string{"Tý", "Sửu", "Dần", "Mão", "Thìn", "Tỵ", "Ngọ", "Mùi", "Thân", "Dậu", "Tuất", "Hợi"}
	festivals  = []Festival{
		{
			Day:   1,
			Month: 1,
			Name:  "Tết Nguyên Đán",
		},
		{
			Day:   15,
			Month: 1,
			Name:  "Rằm tháng Giêng",
		},
		{
			Day:   10,
			Month: 3,
			Name:  "Giỗ Tổ Hùng Vương",
		},
		{
			Day:   15,
			Month: 4,
			Name:  "Phật Đản",
		},
		{
			Day:   5,
			Month: 5,
			Name:  "Lễ Đoan Ngọ",
		},
		{
			Day:   15,
			Month: 7,
			Name:  "Vu Lan",
		},
		{
			Day:   15,
			Month: 8,
			Name:  "Tết Trung Thu",
		},
		{
			Day:   23,
			Month: calendar.MonthsPerYear,
			Name:  "Ông Táo chầu trời",
		},
	}
	solarTerms = []SolarTerm{
		{
			Longitude: 0,
			Name:      "Xuân phân",
		},
		{
			Longitude: 15,
			Name:      "Thanh minh",
		},
		{
			Longitude: 30,
			Name:      "Cốc vũ",
		},
		{
			Longitude: 45,
			Name:      "Lập hạ",
		},
		{
			Longitude: 60,
			Name:      "Tiểu mãn",
		},
		{
			Longitude: 75,
			Name:      "Mang chủng",
		},
		{
			Longitude: 90,
			Name:      "Hạ chí",
		},
		{
			Longitude: 105,
			Name:      "Tiểu thử",
		},
		{
			Longitude: 120,
			Name:      "Đại thử",
		},
		{
			Longitude: 135,
			Name:      "Lập thu",
		},
		{
			Longitude: 150,
			Name:      "Xử thử",
		},
		{
			Longitude: 165,
			Name:      "Bạch lộ",
		},
		{
			Longitude: 180,
			Name:      "Thu phân",
		},
		{
			Longitude: 195,
			Name:      "Hàn lộ",
		},
		{
			Longitude: 210,
			Name:      "Sương giáng",
		},
		{
			Longitude: 225,
			Name:      "Lập đông",
		},
		{
			Longitude: 240,
			Name:      "Tiểu tuyết",
		},
		{
			Longitude: 255,
			Name:      "Đại tuyết",
		},
		{
			Longitude: 270,
			Name:      "Đông chí",
		},
		{
			Longitude: 285,
			Name:      "Tiểu hàn",
		},
		{
			Longitude: 300,
			Name:      "Đại hàn",
		},
		{
			Longitude: 315,
			Name:      "Lập xuân",
		},
		{
			Longitude: 330,
			Name:      "Vũ Thủy",
		},
		{
			Longitude: 345,
			Name:      "Kinh trập",
		},
	}
	luckyHours = [6]string{"110100101100", "001101001011", "110011010010", "101100110100", "001011001101", "010010110011"}
)

// NewGregorian returns a new Gregorian instance.
// 初始化 Gregorian 结构体
func NewGregorian(t time.Time) (g Gregorian) {
	g.Time = t
	return g
}

// NewLunar returns a new Lunar instance.
// 初始化 Lunar 结构体
func NewLunar(year, month, day, hour, minute, second int, isLeapMonth bool) (l Lunar) {
	l.year, l.month, l.day = year, month, day
	l.hour, l.minute, l.second = hour, minute, second
	if isLeapMonth {
		l.leapMonth = month
		l.isLeapYear = 1
	}
	name, offset := time.Now().Zone()
	l.timeZone = offset / 3600
	l.zoneName = name
	l.juliusDay = lunarToJuliusDay(l.day, l.month, l.year, l.isLeapYear, l.timeZone)
	return l
}

// ToLunar Convert Gregorian calendar into Vietnamese Lunar calendar
// 将 公历 转化为 农历
func (g Gregorian) ToLunar() (l Lunar) {
	l.isLeapYear, l.leapMonth = 0, 0
	name, offset := g.Time.Zone()
	l.timeZone = offset / 3600
	l.day, l.month, l.year, l.isLeapYear, l.leapMonth, l.juliusDay = solarToLunar(g.Day(), g.Month(), g.Year(), l.timeZone)

	l.hour, l.minute, l.second = g.Clock()
	l.zoneName = name
	return l
}

// ToLunar Convert Vietnamese lunar calendar into Gregorian calendar
// 将 农历 转化为 公历
func (l Lunar) ToGregorian() (g Gregorian) {

	loc := time.FixedZone(l.zoneName, l.timeZone*3600)

	d, m, y := juliusDayToDate(l.juliusDay)
	g.Time = time.Date(y, time.Month(m), d, l.hour, l.minute, l.second, 0, loc)
	return g
}

func (l Lunar) YearHeavenStem() string {
	return heavenStem[(l.year+6)%10]
}

// Animal gets lunar animal name like "Tý".
func (l Lunar) Animal() string {
	return lunarTimes[(l.year+8)%calendar.MonthsPerYear]
}

// Festival gets lunar festival name like "Trung thu".
func (l Lunar) Festivals() (events []Festival) {
	events = []Festival{}
	month, day := l.month, l.day

	for i := 0; i < len(festivals); i++ {
		event := festivals[i]
		if event.Day == day && event.Month == month {
			events = append(events, event)
		}
	}

	return
}

func (l Lunar) SolarTerm() SolarTerm {
	solarTerm := getSolarTerm(l.juliusDay+1, 7.0)
	return solarTerms[solarTerm]
}

func (l Lunar) LuckyHour() string {
	chiOfDay := (l.juliusDay + 1) % calendar.MonthsPerYear
	luckyHour := luckyHours[chiOfDay%6]

	return luckyHour
}

func (l Lunar) LuckyHours() (ret []LuckyHour) {
	ret = []LuckyHour{}
	luckyHour := l.LuckyHour()
	for i := 0; i < calendar.MonthsPerYear; i++ {
		index := luckyHour[i]
		if index == '1' {
			detail := LuckyHour{
				Chi:  lunarTimes[i],
				From: (i*2 + 23) % 24,
				To:   (i*2 + 1) % 24,
			}
			ret = append(ret, detail)
		}
	}
	return
}

// DateTime gets lunar year, month, day, hour, minute, and second like 2020, 8, 5, 13, 14, 15.
func (l Lunar) DateTime() (year, month, day, hour, minute, second int) {
	return l.year, l.month, l.day, l.hour, l.minute, l.second
}

// Date gets lunar year, month and day like 2020, 8, 5.
func (l Lunar) Date() (year, month, day int) {
	return l.year, l.month, l.day
}

// Time gets lunar hour, minute, and second like 13, 14, 15.
func (l Lunar) Time() (hour, minute, second int) {
	return l.hour, l.minute, l.second
}

// Year gets lunar year like 2020.
func (l Lunar) Year() int {
	return l.year
}

// Month gets lunar month like 8.
func (l Lunar) Month() int {
	return l.month
}

func (l Lunar) MonthHeavenStem() string {
	return heavenStem[((l.year*calendar.MonthsPerYear)+l.month+3)%10]
}

func (l Lunar) MonthAnimal() string {
	return lunarTimes[(l.month+1)%calendar.MonthsPerYear]
}

// LeapMonth gets lunar leap month like 8.
func (l Lunar) LeapMonth() int {
	return l.leapMonth
}

// Day gets lunar day like 5.
func (l Lunar) Day() int {
	return l.day
}

func (l Lunar) DayHeavenStem() string {
	return heavenStem[(l.juliusDay+9)%10]
}

func (l Lunar) DayAnimal() string {
	return lunarTimes[(l.juliusDay+1)%calendar.MonthsPerYear]
}

// ToYearString outputs a string in lunar year format like "Giáp Tý".
func (l Lunar) ToYearString() string {
	year := fmt.Sprintf("%s %s", l.YearHeavenStem(), l.Animal())
	return year
}

// ToMonthString outputs a string in lunar month format like "Giáp Tý".
func (l Lunar) ToMonthString() string {
	month := fmt.Sprintf("%s %s", l.MonthHeavenStem(), l.MonthAnimal())
	return month
}

// ToDayString outputs a string in lunar day format like "Giáp Tý".
func (l Lunar) ToDayString() (day string) {
	day = fmt.Sprintf("%s %s", l.DayHeavenStem(), l.DayAnimal())
	return
}

// ToDateString outputs a string in lunar date format like "Ngày 16 tháng 9 năm 2020".
func (l Lunar) ToDateString() string {
	return fmt.Sprintf("Ngày %02d tháng %02d năm %d", l.day, l.month, l.year)
}

// String outputs a string in YYYY-MM-DD HH::ii::ss format, implement Stringer interface.
func (l Lunar) String() string {
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", l.year, l.month, l.day, l.hour, l.minute, l.second)
}

// IsLeapYear reports whether is leap year.
func (l Lunar) IsLeapYear() bool {
	return l.isLeapYear == 1
}

// IsLeapMonth reports whether is leap month.
func (l Lunar) IsLeapMonth() bool {
	return l.month == l.LeapMonth()
}

// IsRatYear reports whether is year of Rat.
func (l Lunar) IsRatYear() bool {
	return l.year%calendar.MonthsPerYear == 4
}

// IsOxYear reports whether is year of Ox.
func (l Lunar) IsOxYear() bool {
	return l.year%calendar.MonthsPerYear == 5
}

// IsTigerYear reports whether is year of Tiger.
func (l Lunar) IsTigerYear() bool {
	return l.year%calendar.MonthsPerYear == 6
}

// IsRabbitYear reports whether is year of Rabbit.
func (l Lunar) IsCatYear() bool {
	return l.year%calendar.MonthsPerYear == 7
}

// IsDragonYear reports whether is year of Dragon.
func (l Lunar) IsDragonYear() bool {
	return l.year%calendar.MonthsPerYear == 8
}

// IsSnakeYear reports whether is year of Snake.
func (l Lunar) IsSnakeYear() bool {
	return l.year%calendar.MonthsPerYear == 9
}

// IsHorseYear reports whether is year of Horse.
func (l Lunar) IsHorseYear() bool {
	return l.year%calendar.MonthsPerYear == 10
}

// IsGoatYear reports whether is year of Goat.
func (l Lunar) IsGoatYear() bool {
	return l.year%calendar.MonthsPerYear == 11
}

// IsMonkeyYear reports whether is year of Monkey.
func (l Lunar) IsMonkeyYear() bool {
	return l.year%calendar.MonthsPerYear == 0
}

// IsRoosterYear reports whether is year of Rooster.
func (l Lunar) IsRoosterYear() bool {
	return l.year%calendar.MonthsPerYear == 1
}

// IsDogYear reports whether is year of Dog.
func (l Lunar) IsDogYear() bool {
	return l.year%calendar.MonthsPerYear == 2
}

// IsPigYear reports whether is year of Pig.
func (l Lunar) IsPigYear() bool {
	return l.year%calendar.MonthsPerYear == 3
}

// private
func intToFloat(i int) float64 {
	return float64(i)
}

func mathFloor(f float64) float64 {
	return math.Floor(f)
}

func floor(f float64) int {
	return int(mathFloor(f))
}

func dateToJuliusDay(dd, mm, yy int) int {
	a := floor((14 - intToFloat(mm)) / calendar.MonthsPerYear)
	y := intToFloat(yy) + 4800 - intToFloat(a)
	m := intToFloat(mm) + calendar.MonthsPerYear*intToFloat(a) - 3
	juliusDay := intToFloat(dd) + mathFloor((153*m+2)/5) + 365*y + mathFloor(y/4) - math.Floor(y/100) + mathFloor(y/400) - 32045
	if juliusDay < 2299161 {
		// julius calendar
		juliusDay = intToFloat(dd) + mathFloor((153*m+2)/5) + 365*y + mathFloor(y/4) - 32083
	}

	return floor(juliusDay)
}

func juliusDayToDate(jd int) (day, month, year int) {
	var a, b, c, d, e, m int
	if jd > 2299160 {
		a = jd + 32044
		b = (4*a + 3) / 146097
		c = a - (b*146097)/4
	} else {
		b = 0
		c = jd + 32082
	}

	d = (4*c + 3) / 1461
	e = c - (1461*d)/4
	m = (5*e + 2) / 153

	day = e - (153*m+2)/5 + 1
	month = m + 3 - calendar.MonthsPerYear*(m/10)
	year = b*100 + d - 4800 + (m / 10)

	return
}

func newMoon(k int) float64 {
	var kf, t, t2, t3, dr, jd1, m, mpr, f, c1, deltat, jdNew float64
	kf = intToFloat(k)
	t = kf / 1236.85 // Time in Julian centuries from 1900 January 0.5
	t2 = t * t
	t3 = t2 * t
	dr = PI / 180
	jd1 = 2415020.75933 + 29.53058868*kf + 0.0001178*t2 - 0.000000155*t3
	jd1 = jd1 + 0.00033*math.Sin((166.56+132.87*t-0.009173*t2)*dr) // Mean new moon
	m = 359.2242 + 29.10535608*kf - 0.0000333*t2 - 0.00000347*t3   // Sun's mean anomaly
	mpr = 306.0253 + 385.81691806*kf + 0.0107306*t2 + 0.000012*t3  // Moon's mean anomaly
	f = 21.2964 + 390.67050646*kf - 0.0016528*t2 - 0.00000239*t3   // Moon's argument of latitude
	c1 = (0.1734-0.000393*t)*math.Sin(m*dr) + 0.0021*math.Sin(2*dr*m)
	c1 = c1 - 0.4068*math.Sin(mpr*dr) + 0.0161*math.Sin(dr*2*mpr)
	c1 = c1 - 0.0004*math.Sin(dr*3*mpr)
	c1 = c1 + 0.0104*math.Sin(dr*2*f) - 0.0051*math.Sin(dr*(m+mpr))
	c1 = c1 - 0.0074*math.Sin(dr*(m-mpr)) + 0.0004*math.Sin(dr*(2*f+m))
	c1 = c1 - 0.0004*math.Sin(dr*(2*f-m)) - 0.0006*math.Sin(dr*(2*f+mpr))
	c1 = c1 + 0.0010*math.Sin(dr*(2*f-mpr)) + 0.0005*math.Sin(dr*(2*mpr+m))
	if t < -11 {
		deltat = 0.001 + 0.000839*t + 0.0002261*t2 - 0.00000845*t3 - 0.000000081*t*t3
	} else {
		deltat = -0.000278 + 0.000265*t + 0.000262*t2
	}
	jdNew = jd1 + c1 - deltat

	return jdNew
}

func getNewMoonDay(k, tz int) int {
	return floor(newMoon(k) + 0.5 + intToFloat(tz)/24)
}

func sunLongitude(jdn float64) float64 {
	var t, t2, dr, m, l0, dl, l float64
	t = (jdn - 2451545.0) / 36525 // Time in Julian centuries from 2000-01-01 calendar.MonthsPerYear:00:00 GMT
	t2 = t * t
	dr = PI / 180                                                  // degree to radian
	m = 357.52910 + 35999.05030*t - 0.0001559*t2 - 0.00000048*t*t2 // mean anomaly, degree
	l0 = 280.46645 + 36000.76983*t + 0.0003032*t2                  // mean longitude, degree
	dl = (1.914600 - 0.004817*t - 0.000014*t2) * math.Sin(dr*m)
	dl = dl + (0.019993-0.000101*t)*math.Sin(dr*2*m) + 0.00029*math.Sin(dr*3*m)
	l = l0 + dl // true longitude, degree
	l = l * dr
	l = l - PI*2*(math.Floor(l/(PI*2))) // Normalize to (0, 2*PI)
	return l
}

func getSunLongitude(d, tz int) int {
	return floor((sunLongitude(intToFloat(d)-0.5-intToFloat(tz)/24) / PI) * 6)
}

func getSolarTerm(dayNumber int, timeZone int) int {
	return floor(sunLongitude(intToFloat(dayNumber)-0.5-intToFloat(timeZone)/24.0) / PI * calendar.MonthsPerYear)
}

/* Find the day that starts the lunar month 11 of the given year for the given time zone */
func getLunarStartNovember(yy, tz int) int {
	var off, k, nm int

	off = dateToJuliusDay(31, calendar.MonthsPerYear, yy) - 2415021
	k = floor(intToFloat(off) / 29.530588853)
	nm = getNewMoonDay(k, tz)
	sunLong := getSunLongitude(nm, tz) // sun longitude at local midnight
	if sunLong >= 9 {
		nm = getNewMoonDay(k-1, tz)
	}
	return nm
}

func getLeapMonthOffset(a11 float64, tz int) int {
	k := floor((a11-2415021.076998695)/29.530588853 + 0.5)
	last := 0
	i := 1 // We start with the month following lunar month 11
	arc := getSunLongitude(getNewMoonDay(k+i, tz), tz)
	for {
		last = arc
		i++
		newmoon := getNewMoonDay(k+i, tz)
		arc = getSunLongitude(newmoon, tz)

		if arc == last || i >= 14 {
			break
		}
	}
	return i - 1
}

// The function SolarToLunar converts a given Solar date (dd, mm, yy) into
// the corresponding Lunar date (lunarD, lunarM, lunarY, lunarLeap, lunarMonthLeap, juliusDayNumber) in the Vietnamese Lunar calendar.
// The time zone (tz) is used to calculate the date in the Lunar calendar.
func solarToLunar(dd, mm, yy, tz int) (lunarD, lunarM, lunarY, lunarLeap, lunarMonthLeap, juliusDayNumber int) {
	juliusDayNumber = dateToJuliusDay(dd, mm, yy)
	k := floor((intToFloat(juliusDayNumber) - 2415021.076998695) / 29.530588853)
	monthStart := getNewMoonDay(k+1, tz)
	if monthStart > juliusDayNumber {
		monthStart = getNewMoonDay(k, tz)
	}
	a11 := getLunarStartNovember(yy, tz)
	b11 := a11
	if a11 >= monthStart {
		lunarY = yy
		a11 = getLunarStartNovember(yy-1, tz)
	} else {
		lunarY = yy + 1
		b11 = getLunarStartNovember(yy+1, tz)
	}
	lunarD = juliusDayNumber - monthStart + 1
	diff := floor((intToFloat(monthStart) - intToFloat(a11)) / 29)

	lunarLeap = 0
	lunarMonthLeap = 0
	lunarM = diff + 11
	if b11-a11 > 365 {
		leapMonthOffset := getLeapMonthOffset(intToFloat(a11), tz)
		if diff >= leapMonthOffset {
			lunarM = diff + 10
		}
		lunarMonthLeap = leapMonthOffset - 2

		if lunarMonthLeap < 0 {
			lunarMonthLeap += 12
		}

		lunarLeap = 1
	}

	if lunarM > calendar.MonthsPerYear {
		lunarM = lunarM - calendar.MonthsPerYear
	}

	if lunarM >= 11 && diff < 4 {
		lunarY -= 1
	}

	return
}

func lunarToJuliusDay(lunarDay, lunarMonth, lunarYear, lunarLeap, tz int) (juliusDay int) {
	var k, a11, b11, off, leapOff, leapMonth, monthStart int
	if lunarMonth < 11 {
		a11 = getLunarStartNovember(lunarYear-1, tz)
		b11 = getLunarStartNovember(lunarYear, tz)
	} else {
		a11 = getLunarStartNovember(lunarYear, tz)
		b11 = getLunarStartNovember(lunarYear+1, tz)
	}
	k = floor(0.5 + (intToFloat(a11)-2415021.076998695)/29.530588853)
	off = lunarMonth - 11
	if off < 0 {
		off += 12
	}
	if b11-a11 > 365 {
		leapOff = getLeapMonthOffset(intToFloat(a11), tz)
		leapMonth = leapOff - 2
		if leapMonth < 0 {
			leapMonth += 12
		}
		if lunarLeap != 0 && lunarMonth != leapMonth {
			return
		} else if lunarLeap != 0 || off >= leapOff {
			off += 1
		}
	}
	monthStart = getNewMoonDay(k+off, tz)
	juliusDay = monthStart + lunarDay - 1
	return
}
