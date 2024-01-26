package julian

import (
	"math"
	"strconv"
	"time"

	"github.com/golang-module/carbon/v2/calendar"
)

var (
	// julian day or modified julian day decimal precision
	// 儒略日或简化儒略日小数精度
	decimalPrecision = 6
)

// Gregorian defines a Gregorian struct.
// 定义 Gregorian 结构体
type Gregorian struct {
	calendar.Gregorian
}

// Julian defines a Julian struct.
// 定义 Julian 结构体
type Julian struct {
	jd, mjd float64
}

// FromGregorian creates from a Gregorian instance.
// 创建 Gregorian 结构体
func FromGregorian(t time.Time) (g Gregorian) {
	g.Time = t
	return g
}

// FromJulian creates from a new Julian instance.
// 创建 Julian 结构体
func FromJulian(f float64) (j Julian) {
	// get length of the integer part of f
	n := len(strconv.Itoa(int(math.Ceil(f))))
	switch n {
	// modified julian day
	case 5:
		j.mjd = f
		j.jd = f + 2400000.5
	// julian day
	case 7:
		j.jd = f
		j.mjd = f - 2400000.5
	default:
		j.jd = 0
		j.mjd = 0
	}
	return
}

// ToJulian converts Gregorian calendar to Julian calendar.
// 将 公历 转化为 儒略历
func (g Gregorian) ToJulian() (j Julian) {
	if g.IsZero() {
		return j
	}
	y := g.Year()
	m := g.Month()
	d := float64(g.Day()) + ((float64(g.Second())/60+float64(g.Minute()))/60+float64(g.Hour()))/24
	n := 0
	if y*372+m*31+int(d) >= 588829 {
		n = y / 100
		n = 2 - n + n/4
	}
	if m <= 2 {
		m += 12
		y--
	}
	jd := float64(int(365.25*(float64(y)+4716))) + float64(int(30.6001*(float64(m)+1))) + d + float64(n) - 1524.5
	return FromJulian(jd)
}

// ToGregorian converts Julian calendar to Gregorian calendar.
// 将 儒略历 转化为 公历
func (j Julian) ToGregorian() (g Gregorian) {
	if j.jd == 0 || j.mjd == 0 {
		return g
	}
	d := int(j.jd + 0.5)
	f := j.jd + 0.5 - float64(d)
	if d >= 2299161 {
		c := int((float64(d) - 1867216.25) / 36524.25)
		d += 1 + c - c/4
	}
	d += 1524
	year := int((float64(d) - 122.1) / 365.25)
	d -= int(365.25 * float64(year))
	month := int(float64(d) / 30.601)
	d -= int(30.601 * float64(month))
	day := d
	if month > 13 {
		month -= 13
		year -= 4715
	} else {
		month -= 1
		year -= 4716
	}
	f *= 24
	hour := int(f)

	f -= float64(hour)
	f *= 60
	minute := int(f)

	f -= float64(minute)
	f *= 60
	second := int(math.Round(f))
	return FromGregorian(time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local))
}

// JD gets julian day like 2460332.5
// 获取儒略日
func (j Julian) JD(precision ...int) float64 {
	if len(precision) > 0 {
		decimalPrecision = precision[0]
	}
	return parseFloat64(j.jd, decimalPrecision)
}

// MJD gets modified julian day like 60332
// 获取简化儒略日
func (j Julian) MJD(precision ...int) float64 {
	if len(precision) > 0 {
		decimalPrecision = precision[0]
	}
	return parseFloat64(j.mjd, decimalPrecision)
}

// parseFloat64 round to n decimal places
// 四舍五入保留 n 位小数点
func parseFloat64(f float64, n int) float64 {
	p := math.Pow10(n)
	return math.Round(f*p) / p
}
