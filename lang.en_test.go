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
		id       int    // 测试id
		input    string // 输入值
		expected string // 期望值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},
		{5, "0000-00-00 00:00:00", ""},

		{6, "2020-01-05", "Capricorn"},
		{7, "2020-02-05", "Aquarius"},
		{8, "2020-03-05", "Pisces"},
		{9, "2020-04-05", "Aries"},
		{10, "2020-05-05", "Taurus"},
		{11, "2020-06-05", "Gemini"},
		{12, "2020-07-05", "Cancer"},
		{13, "2020-08-05", "Leo"},
		{14, "2020-09-05", "Virgo"},
		{15, "2020-10-05", "Libra"},
		{16, "2020-11-05", "Scorpio"},
		{17, "2020-12-05", "Sagittarius"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(english)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Constellation(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_English_Season(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected string // 期望值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},
		{5, "0000-00-00 00:00:00", ""},

		{6, "2020-01-05", "Winter"},
		{7, "2020-02-05", "Winter"},
		{8, "2020-03-05", "Spring"},
		{9, "2020-04-05", "Spring"},
		{10, "2020-05-05", "Spring"},
		{11, "2020-06-05", "Summer"},
		{12, "2020-07-05", "Summer"},
		{13, "2020-08-05", "Summer"},
		{14, "2020-09-05", "Autumn"},
		{15, "2020-10-05", "Autumn"},
		{16, "2020-11-05", "Autumn"},
		{17, "2020-12-05", "Winter"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(english)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Season(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_English_ToMonthString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected string // 期望值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},
		{5, "0000-00-00 00:00:00", ""},

		{6, "2020-01-05", "January"},
		{7, "2020-02-05", "February"},
		{8, "2020-03-05", "March"},
		{9, "2020-04-05", "April"},
		{10, "2020-05-05", "May"},
		{11, "2020-06-05", "June"},
		{12, "2020-07-05", "July"},
		{13, "2020-08-05", "August"},
		{14, "2020-09-05", "September"},
		{15, "2020-10-05", "October"},
		{16, "2020-11-05", "November"},
		{17, "2020-12-05", "December"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(english)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToMonthString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_English_ToShortMonthString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected string // 期望值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},
		{5, "0000-00-00 00:00:00", ""},

		{6, "2020-01-05", "Jan"},
		{7, "2020-02-05", "Feb"},
		{8, "2020-03-05", "Mar"},
		{9, "2020-04-05", "Apr"},
		{10, "2020-05-05", "May"},
		{11, "2020-06-05", "Jun"},
		{12, "2020-07-05", "Jul"},
		{13, "2020-08-05", "Aug"},
		{14, "2020-09-05", "Sep"},
		{15, "2020-10-05", "Oct"},
		{16, "2020-11-05", "Nov"},
		{17, "2020-12-05", "Dec"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(english)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortMonthString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_English_ToWeekString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected string // 期望值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},
		{5, "0000-00-00 00:00:00", ""},

		{6, "2020-08-01", "Saturday"},
		{7, "2020-08-02", "Sunday"},
		{8, "2020-08-03", "Monday"},
		{9, "2020-08-04", "Tuesday"},
		{10, "2020-08-05", "Wednesday"},
		{11, "2020-08-06", "Thursday"},
		{12, "2020-08-07", "Friday"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(english)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToWeekString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_English_ToShortWeekString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected string // 期望值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},
		{5, "0000-00-00 00:00:00", ""},

		{6, "2020-08-01", "Sat"},
		{7, "2020-08-02", "Sun"},
		{8, "2020-08-03", "Mon"},
		{9, "2020-08-04", "Tue"},
		{10, "2020-08-05", "Wed"},
		{11, "2020-08-06", "Thu"},
		{12, "2020-08-07", "Fri"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(english)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortWeekString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_English_DiffForHumans(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input1   string // 输入值
		input2   string // 输入值
		expected string // 期望值
	}{
		{1, "2020-08-05 13:14:15", "2020-08-05 13:14:15", "just now"},
		{2, "2020-08-05 13:14:15", "2021-08-05 13:14:15", "1 year before"},
		{3, "2020-08-05 13:14:15", "2019-08-05 13:14:15", "1 year after"},
		{4, "2020-08-05 13:14:15", "2030-08-05 13:14:15", "10 years before"},
		{5, "2020-08-05 13:14:15", "2010-08-05 13:14:15", "10 years after"},

		{6, "2020-08-05 13:14:15", "2020-09-05 13:14:15", "1 month before"},
		{7, "2020-08-05 13:14:15", "2020-07-05 13:14:15", "1 month after"},
		{8, "2020-08-05 13:14:15", "2021-06-05 13:14:15", "10 months before"},
		{9, "2020-08-05 13:14:15", "2019-10-05 13:14:15", "10 months after"},

		{10, "2020-08-05 13:14:15", "2020-08-06 13:14:15", "1 day before"},
		{11, "2020-08-05 13:14:15", "2020-08-04 13:14:15", "1 day after"},
		{12, "2020-08-05 13:14:15", "2020-08-15 13:14:15", "1 week before"},
		{13, "2020-08-05 13:14:15", "2020-07-26 13:14:15", "1 week after"},

		{14, "2020-08-05 13:14:15", "2020-08-05 14:14:15", "1 hour before"},
		{15, "2020-08-05 13:14:15", "2020-08-05 12:14:15", "1 hour after"},
		{16, "2020-08-05 13:14:15", "2020-08-05 23:14:15", "10 hours before"},
		{17, "2020-08-05 13:14:15", "2020-08-05 03:14:15", "10 hours after"},

		{18, "2020-08-05 13:14:15", "2020-08-05 13:15:15", "1 minute before"},
		{19, "2020-08-05 13:14:15", "2020-08-05 13:13:15", "1 minute after"},
		{20, "2020-08-05 13:14:15", "2020-08-05 13:24:15", "10 minutes before"},
		{21, "2020-08-05 13:14:15", "2020-08-05 13:04:15", "10 minutes after"},

		{22, "2020-08-05 13:14:15", "2020-08-05 13:14:16", "1 second before"},
		{23, "2020-08-05 13:14:15", "2020-08-05 13:14:14", "1 second after"},
		{24, "2020-08-05 13:14:15", "2020-08-05 13:14:25", "10 seconds before"},
		{25, "2020-08-05 13:14:15", "2020-08-05 13:14:05", "10 seconds after"},
	}

	for _, test := range tests {
		c1 := Parse(test.input1)
		c2 := Parse(test.input2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.SetLocale(english).DiffForHumans(c2), "Current test id is "+strconv.Itoa(test.id))
	}
}
