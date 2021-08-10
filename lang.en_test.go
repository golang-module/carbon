package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var english = "en"

func TestLang_English_Constellation(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-05", "Capricorn"},
		{"2020-02-05", "Aquarius"},
		{"2020-03-05", "Pisces"},
		{"2020-04-05", "Aries"},
		{"2020-05-05", "Taurus"},
		{"2020-06-05", "Gemini"},
		{"2020-07-05", "Cancer"},
		{"2020-08-05", "Leo"},
		{"2020-09-05", "Virgo"},
		{"2020-10-05", "Libra"},
		{"2020-11-05", "Scorpio"},
		{"2020-12-05", "Sagittarius"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(english)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Constellation(), "test index id is "+strconv.Itoa(index))
	}
}

func TestLang_English_Season(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-05", "Winter"},
		{"2020-02-05", "Winter"},
		{"2020-03-05", "Spring"},
		{"2020-04-05", "Spring"},
		{"2020-05-05", "Spring"},
		{"2020-06-05", "Summer"},
		{"2020-07-05", "Summer"},
		{"2020-08-05", "Summer"},
		{"2020-09-05", "Autumn"},
		{"2020-10-05", "Autumn"},
		{"2020-11-05", "Autumn"},
		{"2020-12-05", "Winter"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(english)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Season(), "test index id is "+strconv.Itoa(index))
	}
}

func TestLang_English_ToMonthString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-05", "January"},
		{"2020-02-05", "February"},
		{"2020-03-05", "March"},
		{"2020-04-05", "April"},
		{"2020-05-05", "May"},
		{"2020-06-05", "June"},
		{"2020-07-05", "July"},
		{"2020-08-05", "August"},
		{"2020-09-05", "September"},
		{"2020-10-05", "October"},
		{"2020-11-05", "November"},
		{"2020-12-05", "December"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(english)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToMonthString(), "test index id is "+strconv.Itoa(index))
	}
}

func TestLang_English_ToShortMonthString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-05", "Jan"},
		{"2020-02-05", "Feb"},
		{"2020-03-05", "Mar"},
		{"2020-04-05", "Apr"},
		{"2020-05-05", "May"},
		{"2020-06-05", "Jun"},
		{"2020-07-05", "Jul"},
		{"2020-08-05", "Aug"},
		{"2020-09-05", "Sep"},
		{"2020-10-05", "Oct"},
		{"2020-11-05", "Nov"},
		{"2020-12-05", "Dec"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(english)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortMonthString(), "test index id is "+strconv.Itoa(index))
	}
}

func TestLang_English_ToWeekString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-01", "Saturday"},
		{"2020-08-02", "Sunday"},
		{"2020-08-03", "Monday"},
		{"2020-08-04", "Tuesday"},
		{"2020-08-05", "Wednesday"},
		{"2020-08-06", "Thursday"},
		{"2020-08-07", "Friday"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(english)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToWeekString(), "test index id is "+strconv.Itoa(index))
	}
}

func TestLang_English_ToShortWeekString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-01", "Sat"},
		{"2020-08-02", "Sun"},
		{"2020-08-03", "Mon"},
		{"2020-08-04", "Tue"},
		{"2020-08-05", "Wed"},
		{"2020-08-06", "Thu"},
		{"2020-08-07", "Fri"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(english)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortWeekString(), "test index id is "+strconv.Itoa(index))
	}
}

func TestLang_English_DiffForHumans(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   string // 输入值
		input2   string // 输入值
		expected string // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", "just now"},
		{"2020-08-05 13:14:15", "2021-08-05 13:14:15", "1 year before"},
		{"2020-08-05 13:14:15", "2019-08-05 13:14:15", "1 year after"},
		{"2020-08-05 13:14:15", "2030-08-05 13:14:15", "10 years before"},
		{"2020-08-05 13:14:15", "2010-08-05 13:14:15", "10 years after"},

		{"2020-08-05 13:14:15", "2020-09-05 13:14:15", "1 month before"},
		{"2020-08-05 13:14:15", "2020-07-05 13:14:15", "1 month after"},
		{"2020-08-05 13:14:15", "2021-06-05 13:14:15", "10 months before"},
		{"2020-08-05 13:14:15", "2019-10-05 13:14:15", "10 months after"},

		{"2020-08-05 13:14:15", "2020-08-06 13:14:15", "1 day before"},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", "1 day after"},
		{"2020-08-05 13:14:15", "2020-08-15 13:14:15", "1 week before"},
		{"2020-08-05 13:14:15", "2020-07-26 13:14:15", "1 week after"},

		{"2020-08-05 13:14:15", "2020-08-05 14:14:15", "1 hour before"},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:15", "1 hour after"},
		{"2020-08-05 13:14:15", "2020-08-05 23:14:15", "10 hours before"},
		{"2020-08-05 13:14:15", "2020-08-05 03:14:15", "10 hours after"},

		{"2020-08-05 13:14:15", "2020-08-05 13:15:15", "1 minute before"},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:15", "1 minute after"},
		{"2020-08-05 13:14:15", "2020-08-05 13:24:15", "10 minutes before"},
		{"2020-08-05 13:14:15", "2020-08-05 13:04:15", "10 minutes after"},

		{"2020-08-05 13:14:15", "2020-08-05 13:14:16", "1 second before"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:14", "1 second after"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:25", "10 seconds before"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:05", "10 seconds after"},
	}

	for index, test := range tests {
		c1 := Parse(test.input1)
		c2 := Parse(test.input2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.SetLocale(english).DiffForHumans(c2), "test index id is "+strconv.Itoa(index))
	}
}
