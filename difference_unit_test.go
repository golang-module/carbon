package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_DiffInYears(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param    string
		expected int64
	}{
		{"2020-08-05 13:14:15", "2020-07-28 13:14:00", 0},
		{"2020-12-31 13:14:15", "2021-01-01 13:14:15", 0},
		{"2020-08-05 13:14:15", "2021-08-28 13:14:59", 1},
		{"2018-08-28 13:14:59", "2020-08-05 13:14:15", 1},
		{"2020-08-05 13:14:15", "2018-08-28 13:14:59", -1},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInYears(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffAbsInYears(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param    string
		expected int64
	}{
		{"2020-08-05 13:14:15", "2020-07-28 13:14:00", 0},
		{"2020-12-31 13:14:15", "2021-01-01 13:14:15", 0},
		{"2020-08-05 13:14:15", "2021-08-28 13:14:59", 1},
		{"2020-08-05 13:14:15", "2018-08-28 13:14:59", 1},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffAbsInYears(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffInMonths(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param    string
		expected int64
	}{
		{"2020-08-05 13:14:15", "2020-07-28 13:14:00", 0},
		{"2020-08-05 13:14:15", "2020-09-06 13:14:59", 1},
		{"2020-12-31 13:14:15", "2021-01-31 13:14:15", 1},
		{"2020-08-05 13:14:15", "2021-08-28 13:14:59", 12},
		{"2020-08-05 13:14:15", "2018-08-28 13:14:59", -24},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInMonths(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffAbsInMonths(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param    string
		expected int64
	}{
		{"2020-08-05 13:14:15", "2020-07-28 13:14:00", 0},
		{"2020-12-31 13:14:15", "2021-01-01 13:14:15", 0},
		{"2020-08-05 13:14:15", "2021-08-28 13:14:59", 12},
		{"2020-08-05 13:14:15", "2018-08-28 13:14:59", 24},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffAbsInMonths(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffInWeeks(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param    string
		expected int64
	}{
		{"2020-08-05 13:14:15", "2020-07-28 13:14:00", -1},
		{"2020-08-05 13:14:15", "2020-07-28 13:14:15", -1},
		{"2020-08-05 13:14:15", "2020-07-28 13:14:59", -1},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-12 13:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-12 13:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-12 13:14:59", 1},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInWeeks(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffAbsInWeeks(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param    string
		expected int64
	}{
		{"2020-08-05 13:14:15", "2020-07-28 13:14:00", 1},
		{"2020-08-05 13:14:15", "2020-07-28 13:14:15", 1},
		{"2020-08-05 13:14:15", "2020-07-28 13:14:59", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-12 13:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-12 13:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-12 13:14:59", 1},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffAbsInWeeks(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffInDays(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param    string
		expected int64
	}{
		{"2020-08-05 13:14:15", "2020-08-04 13:00:00", -1},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", -1},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:00:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-06 13:00:00", 0},
		{"2020-08-05 13:14:15", "2020-08-06 13:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-06 13:14:59", 1},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInDays(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffAbsInDays(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param    string
		expected int64
	}{
		{"2020-08-05 13:14:15", "2020-08-04 13:00:00", 1},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:00:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-06 13:00:00", 0},
		{"2020-08-05 13:14:15", "2020-08-06 13:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-06 13:14:59", 1},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffAbsInDays(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffInHours(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param    string
		expected int64
	}{
		{"2020-08-05 13:14:15", "2020-08-05 12:14:00", -1},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:15", -1},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 14:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 14:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-05 14:14:59", 1},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInHours(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffAbsInHours(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param    string
		expected int64
	}{
		{"2020-08-05 13:14:15", "2020-08-05 12:14:00", 1},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 14:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 14:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-05 14:14:59", 1},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffAbsInHours(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffInMinutes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param    string
		expected int64
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:13:00", -1},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:15", -1},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:15:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:15:15", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:15:59", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:16:00", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:16:15", 2},
		{"2020-08-05 13:14:15", "2020-08-05 13:16:59", 2},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInMinutes(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffAbsInMinutes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param    string
		expected int64
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:13:00", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:15", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:15:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:15:15", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:15:59", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:16:00", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:16:15", 2},
		{"2020-08-05 13:14:15", "2020-08-05 13:16:59", 2},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffAbsInMinutes(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffInSeconds(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param    string
		expected int64
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:20", 5},
		{"2020-08-05 13:14:20", "2020-08-05 13:14:15", -5},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInSeconds(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffAbsInSeconds(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		param    string
		expected int64
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:20", 5},
		{"2020-08-05 13:14:20", "2020-08-05 13:14:15", 5},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffAbsInSeconds(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffInString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    Carbon
		expected string
	}{
		0: {Now(), "just now"},

		1: {Now().AddYearsNoOverflow(1), "-1 year"},
		2: {Now().SubYearsNoOverflow(1), "1 year"},
		3: {Now().AddYearsNoOverflow(10), "-10 years"},
		4: {Now().SubYearsNoOverflow(10), "10 years"},

		5: {Now().AddMonthsNoOverflow(1), "-1 month"},
		6: {Now().SubMonthsNoOverflow(1), "1 month"},
		7: {Now().AddMonthsNoOverflow(10), "-10 months"},
		8: {Now().SubMonthsNoOverflow(10), "10 months"},

		9:  {Now().AddDays(1), "-1 day"},
		10: {Now().SubDays(1), "1 day"},
		11: {Now().AddDays(10), "-1 week"},
		12: {Now().SubDays(10), "1 week"},

		13: {Now().AddHours(1), "-1 hour"},
		14: {Now().SubHours(1), "1 hour"},
		15: {Now().AddHours(10), "-10 hours"},
		16: {Now().SubHours(10), "10 hours"},

		17: {Now().AddMinutes(1), "-1 minute"},
		18: {Now().SubMinutes(1), "1 minute"},
		19: {Now().AddMinutes(10), "-10 minutes"},
		20: {Now().SubMinutes(10), "10 minutes"},

		21: {Now().AddSeconds(1), "-1 second"},
		22: {Now().SubSeconds(1), "1 second"},
		23: {Now().AddSeconds(10), "-10 seconds"},
		24: {Now().SubSeconds(10), "10 seconds"},
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DiffInString(), "Current test index is "+strconv.Itoa(index))
	}
	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DiffInString(Now()), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLangError_DiffInString(t *testing.T) {
	lang := NewLanguage()
	lang.SetLocale("xxx")
	c := SetLanguage(lang).Now().AddMonths(1)
	assert.NotNil(t, c.Error, "It should catch an exception in DiffInString()")
	assert.Equal(t, "", c.DiffInString())
}

func TestCarbon_DiffAbsInString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    Carbon
		expected string
	}{
		0: {Now(), "just now"},

		1: {Now().AddYearsNoOverflow(1), "1 year"},
		2: {Now().SubYearsNoOverflow(1), "1 year"},
		3: {Now().AddYearsNoOverflow(10), "10 years"},
		4: {Now().SubYearsNoOverflow(10), "10 years"},

		5: {Now().AddMonthsNoOverflow(1), "1 month"},
		6: {Now().SubMonthsNoOverflow(1), "1 month"},
		7: {Now().AddMonthsNoOverflow(10), "10 months"},
		8: {Now().SubMonthsNoOverflow(10), "10 months"},

		9:  {Now().AddDays(1), "1 day"},
		10: {Now().SubDays(1), "1 day"},
		11: {Now().AddDays(10), "1 week"},
		12: {Now().SubDays(10), "1 week"},

		13: {Now().AddHours(1), "1 hour"},
		14: {Now().SubHours(1), "1 hour"},
		15: {Now().AddHours(10), "10 hours"},
		16: {Now().SubHours(10), "10 hours"},

		17: {Now().AddMinutes(1), "1 minute"},
		18: {Now().SubMinutes(1), "1 minute"},
		19: {Now().AddMinutes(10), "10 minutes"},
		20: {Now().SubMinutes(10), "10 minutes"},

		21: {Now().AddSeconds(1), "1 second"},
		22: {Now().SubSeconds(1), "1 second"},
		23: {Now().AddSeconds(10), "10 seconds"},
		24: {Now().SubSeconds(10), "10 seconds"},
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DiffAbsInString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DiffAbsInString(Now()), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLangError_DiffAbsInString(t *testing.T) {
	lang := NewLanguage()
	lang.SetLocale("xxx")
	c := SetLanguage(lang).Now().AddMonths(1)
	assert.NotNil(t, c.Error, "It should catch an exception in DiffInStringWithAbs()")
	assert.Equal(t, "", c.DiffAbsInString())
}

func TestCarbon_DiffForHumans(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    Carbon
		expected string
	}{
		0: {Now(), "just now"},

		1: {Now().AddYearsNoOverflow(1), "1 year from now"},
		2: {Now().SubYearsNoOverflow(1), "1 year ago"},
		3: {Now().AddYearsNoOverflow(10), "10 years from now"},
		4: {Now().SubYearsNoOverflow(10), "10 years ago"},

		5: {Now().AddMonthsNoOverflow(1), "1 month from now"},
		6: {Now().SubMonthsNoOverflow(1), "1 month ago"},
		7: {Now().AddMonthsNoOverflow(10), "10 months from now"},
		8: {Now().SubMonthsNoOverflow(10), "10 months ago"},

		9:  {Now().AddDays(1), "1 day from now"},
		10: {Now().SubDays(1), "1 day ago"},
		11: {Now().AddDays(10), "1 week from now"},
		12: {Now().SubDays(10), "1 week ago"},

		13: {Now().AddHours(1), "1 hour from now"},
		14: {Now().SubHours(1), "1 hour ago"},
		15: {Now().AddHours(10), "10 hours from now"},
		16: {Now().SubHours(10), "10 hours ago"},

		17: {Now().AddMinutes(1), "1 minute from now"},
		18: {Now().SubMinutes(1), "1 minute ago"},
		19: {Now().AddMinutes(10), "10 minutes from now"},
		20: {Now().SubMinutes(10), "10 minutes ago"},

		21: {Now().AddSeconds(1), "1 second from now"},
		22: {Now().SubSeconds(1), "1 second ago"},
		23: {Now().AddSeconds(10), "10 seconds from now"},
		24: {Now().SubSeconds(10), "10 seconds ago"},
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DiffForHumans(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLangError_DiffForHumans(t *testing.T) {
	lang := NewLanguage()
	lang.SetLocale("xxx")
	c := SetLanguage(lang).Now().AddMonths(1)
	assert.NotNil(t, c.Error, "It should catch an exception in DiffForHumans()")
	assert.Equal(t, "", c.DiffForHumans())
}
