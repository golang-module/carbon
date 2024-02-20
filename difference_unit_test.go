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

	now := Now()
	tests := []struct {
		input    Carbon
		expected string
	}{
		0: {Now(), "just now"},

		1: {now.AddYearsNoOverflow(1), "-1 year"},
		2: {now.SubYearsNoOverflow(1), "1 year"},
		3: {now.AddYearsNoOverflow(10), "-10 years"},
		4: {now.SubYearsNoOverflow(10), "10 years"},

		5: {now.AddMonthsNoOverflow(1), "-1 month"},
		6: {now.SubMonthsNoOverflow(1), "1 month"},
		7: {now.AddMonthsNoOverflow(10), "-10 months"},
		8: {now.SubMonthsNoOverflow(10), "10 months"},

		9:  {now.AddDays(1), "-1 day"},
		10: {now.SubDays(1), "1 day"},
		11: {now.AddDays(10), "-1 week"},
		12: {now.SubDays(10), "1 week"},

		13: {now.AddHours(1), "-1 hour"},
		14: {now.SubHours(1), "1 hour"},
		15: {now.AddHours(10), "-10 hours"},
		16: {now.SubHours(10), "10 hours"},

		17: {now.AddMinutes(1), "-1 minute"},
		18: {now.SubMinutes(1), "1 minute"},
		19: {now.AddMinutes(10), "-10 minutes"},
		20: {now.SubMinutes(10), "10 minutes"},

		21: {now.AddSeconds(1), "-1 second"},
		22: {now.SubSeconds(1), "1 second"},
		23: {now.AddSeconds(10), "-10 seconds"},
		24: {now.SubSeconds(10), "10 seconds"},
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

	now := Now()
	tests := []struct {
		input    Carbon
		expected string
	}{
		0: {Now(), "just now"},

		1: {now.AddYearsNoOverflow(1), "1 year"},
		2: {now.SubYearsNoOverflow(1), "1 year"},
		3: {now.AddYearsNoOverflow(10), "10 years"},
		4: {now.SubYearsNoOverflow(10), "10 years"},

		5: {now.AddMonthsNoOverflow(1), "1 month"},
		6: {now.SubMonthsNoOverflow(1), "1 month"},
		7: {now.AddMonthsNoOverflow(10), "10 months"},
		8: {now.SubMonthsNoOverflow(10), "10 months"},

		9:  {now.AddDays(1), "1 day"},
		10: {now.SubDays(1), "1 day"},
		11: {now.AddDays(10), "1 week"},
		12: {now.SubDays(10), "1 week"},

		13: {now.AddHours(1), "1 hour"},
		14: {now.SubHours(1), "1 hour"},
		15: {now.AddHours(10), "10 hours"},
		16: {now.SubHours(10), "10 hours"},

		17: {now.AddMinutes(1), "1 minute"},
		18: {now.SubMinutes(1), "1 minute"},
		19: {now.AddMinutes(10), "10 minutes"},
		20: {now.SubMinutes(10), "10 minutes"},

		21: {now.AddSeconds(1), "1 second"},
		22: {now.SubSeconds(1), "1 second"},
		23: {now.AddSeconds(10), "10 seconds"},
		24: {now.SubSeconds(10), "10 seconds"},
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
	assert.NotNil(t, c.Error, "It should catch an exception in DiffAbsInString()")
	assert.Equal(t, "", c.DiffAbsInString())
}

func TestCarbon_DiffInDuration(t *testing.T) {
	assert := assert.New(t)

	now := Now()
	tests := []struct {
		input    Carbon
		expected string
	}{
		0: {now, "0s"},

		1: {now.AddYearsNoOverflow(1), "-8784h0m0s"},
		2: {now.SubYearsNoOverflow(1), "8760h0m0s"},
		3: {now.AddYearsNoOverflow(10), "-87672h0m0s"},
		4: {now.SubYearsNoOverflow(10), "87648h0m0s"},

		5: {now.AddMonthsNoOverflow(1), "-696h0m0s"},
		6: {now.SubMonthsNoOverflow(1), "744h0m0s"},
		7: {now.AddMonthsNoOverflow(10), "-7296h0m0s"},
		8: {now.SubMonthsNoOverflow(10), "7344h0m0s"},

		9:  {now.AddDays(1), "-24h0m0s"},
		10: {now.SubDays(1), "24h0m0s"},
		11: {now.AddDays(10), "-240h0m0s"},
		12: {now.SubDays(10), "240h0m0s"},

		13: {now.AddHours(1), "-1h0m0s"},
		14: {now.SubHours(1), "1h0m0s"},
		15: {now.AddHours(10), "-10h0m0s"},
		16: {now.SubHours(10), "10h0m0s"},

		17: {now.AddMinutes(1), "-1m0s"},
		18: {now.SubMinutes(1), "1m0s"},
		19: {now.AddMinutes(10), "-10m0s"},
		20: {now.SubMinutes(10), "10m0s"},

		21: {now.AddSeconds(1), "-1s"},
		22: {now.SubSeconds(1), "1s"},
		23: {now.AddSeconds(10), "-10s"},
		24: {now.SubSeconds(10), "10s"},
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DiffInDuration(now).String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffAbsInDuration(t *testing.T) {
	assert := assert.New(t)

	now := Now()
	tests := []struct {
		input    Carbon
		expected string
	}{
		0: {now, "0s"},

		1: {now.AddYearsNoOverflow(1), "8784h0m0s"},
		2: {now.SubYearsNoOverflow(1), "8760h0m0s"},
		3: {now.AddYearsNoOverflow(10), "87672h0m0s"},
		4: {now.SubYearsNoOverflow(10), "87648h0m0s"},

		5: {now.AddMonthsNoOverflow(1), "696h0m0s"},
		6: {now.SubMonthsNoOverflow(1), "744h0m0s"},
		7: {now.AddMonthsNoOverflow(10), "7296h0m0s"},
		8: {now.SubMonthsNoOverflow(10), "7344h0m0s"},

		9:  {now.AddDays(1), "24h0m0s"},
		10: {now.SubDays(1), "24h0m0s"},
		11: {now.AddDays(10), "240h0m0s"},
		12: {now.SubDays(10), "240h0m0s"},

		13: {now.AddHours(1), "1h0m0s"},
		14: {now.SubHours(1), "1h0m0s"},
		15: {now.AddHours(10), "10h0m0s"},
		16: {now.SubHours(10), "10h0m0s"},

		17: {now.AddMinutes(1), "1m0s"},
		18: {now.SubMinutes(1), "1m0s"},
		19: {now.AddMinutes(10), "10m0s"},
		20: {now.SubMinutes(10), "10m0s"},

		21: {now.AddSeconds(1), "1s"},
		22: {now.SubSeconds(1), "1s"},
		23: {now.AddSeconds(10), "10s"},
		24: {now.SubSeconds(10), "10s"},
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DiffAbsInDuration(now).String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffForHumans(t *testing.T) {
	assert := assert.New(t)

	now := Now()
	tests := []struct {
		input    Carbon
		expected string
	}{
		0: {Now(), "just now"},

		1: {now.AddYearsNoOverflow(1), "1 year from now"},
		2: {now.SubYearsNoOverflow(1), "1 year ago"},
		3: {now.AddYearsNoOverflow(10), "10 years from now"},
		4: {now.SubYearsNoOverflow(10), "10 years ago"},

		5: {now.AddMonthsNoOverflow(1), "1 month from now"},
		6: {now.SubMonthsNoOverflow(1), "1 month ago"},
		7: {now.AddMonthsNoOverflow(10), "10 months from now"},
		8: {now.SubMonthsNoOverflow(10), "10 months ago"},

		9:  {now.AddDays(1), "1 day from now"},
		10: {now.SubDays(1), "1 day ago"},
		11: {now.AddDays(10), "1 week from now"},
		12: {now.SubDays(10), "1 week ago"},

		13: {now.AddHours(1), "1 hour from now"},
		14: {now.SubHours(1), "1 hour ago"},
		15: {now.AddHours(10), "10 hours from now"},
		16: {now.SubHours(10), "10 hours ago"},

		17: {now.AddMinutes(1), "1 minute from now"},
		18: {now.SubMinutes(1), "1 minute ago"},
		19: {now.AddMinutes(10), "10 minutes from now"},
		20: {now.SubMinutes(10), "10 minutes ago"},

		21: {now.AddSeconds(1), "1 second from now"},
		22: {now.SubSeconds(1), "1 second ago"},
		23: {now.AddSeconds(10), "10 seconds from now"},
		24: {now.SubSeconds(10), "10 seconds ago"},
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
