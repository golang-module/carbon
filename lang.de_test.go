package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var german = "de"

func TestLang_German_Constellation(t *testing.T) {
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

		{"2020-01-05", "Steinbock"},
		{"2020-02-05", "Wassermann"},
		{"2020-03-05", "Fisch"},
		{"2020-04-05", "Widder"},
		{"2020-05-05", "Stier"},
		{"2020-06-05", "Zwilling"},
		{"2020-07-05", "Krebs"},
		{"2020-08-05", "Löwe"},
		{"2020-09-05", "Jungfrau"},
		{"2020-10-05", "Waage"},
		{"2020-11-05", "Skorpion"},
		{"2020-12-05", "Schütze"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(german)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Constellation(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_German_Season(t *testing.T) {
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
		{"2020-03-05", "Frühling"},
		{"2020-04-05", "Frühling"},
		{"2020-05-05", "Frühling"},
		{"2020-06-05", "Sommer"},
		{"2020-07-05", "Sommer"},
		{"2020-08-05", "Sommer"},
		{"2020-09-05", "Herbst"},
		{"2020-10-05", "Herbst"},
		{"2020-11-05", "Herbst"},
		{"2020-12-05", "Winter"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(german)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Season(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_German_ToMonthString(t *testing.T) {
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

		{"2020-01-05", "Januar"},
		{"2020-02-05", "Februar"},
		{"2020-03-05", "März"},
		{"2020-04-05", "April"},
		{"2020-05-05", "Mai"},
		{"2020-06-05", "Juni"},
		{"2020-07-05", "Juli"},
		{"2020-08-05", "August"},
		{"2020-09-05", "September"},
		{"2020-10-05", "Oktober"},
		{"2020-11-05", "November"},
		{"2020-12-05", "Dezember"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(german)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToMonthString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_German_ToShortMonthString(t *testing.T) {
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
		{"2020-03-05", "Mär"},
		{"2020-04-05", "Apr"},
		{"2020-05-05", "Mai"},
		{"2020-06-05", "Jun"},
		{"2020-07-05", "Jul"},
		{"2020-08-05", "Aug"},
		{"2020-09-05", "Sep"},
		{"2020-10-05", "Okt"},
		{"2020-11-05", "Nov"},
		{"2020-12-05", "Dez"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(german)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortMonthString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_German_ToWeekString(t *testing.T) {
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

		{"2020-08-01", "Samstag"},
		{"2020-08-02", "Sonntag"},
		{"2020-08-03", "Montag"},
		{"2020-08-04", "Dienstag"},
		{"2020-08-05", "Mittwoch"},
		{"2020-08-06", "Donnerstag"},
		{"2020-08-07", "Freitag"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(german)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToWeekString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_German_ToShortWeekString(t *testing.T) {
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

		{"2020-08-01", "Sa"},
		{"2020-08-02", "So"},
		{"2020-08-03", "Mo"},
		{"2020-08-04", "Di"},
		{"2020-08-05", "Mi"},
		{"2020-08-06", "Do"},
		{"2020-08-07", "Fr"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(german)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortWeekString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_German_DiffForHumans(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   string // 输入值
		input2   string // 输入值
		expected string // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", "gerade eben"},
		{"2020-08-05 13:14:15", "2021-08-05 13:14:15", "1 Jahr davor"},
		{"2020-08-05 13:14:15", "2019-08-05 13:14:15", "1 Jahr danach"},
		{"2020-08-05 13:14:15", "2030-08-05 13:14:15", "10 Jahre davor"},
		{"2020-08-05 13:14:15", "2010-08-05 13:14:15", "10 Jahre danach"},

		{"2020-08-05 13:14:15", "2020-09-05 13:14:15", "1 Monat davor"},
		{"2020-08-05 13:14:15", "2020-07-05 13:14:15", "1 Monat danach"},
		{"2020-08-05 13:14:15", "2021-06-05 13:14:15", "10 Monate davor"},
		{"2020-08-05 13:14:15", "2019-10-05 13:14:15", "10 Monate danach"},

		{"2020-08-05 13:14:15", "2020-08-06 13:14:15", "1 Tag davor"},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", "1 Tag danach"},
		{"2020-08-05 13:14:15", "2020-08-15 13:14:15", "1 Woche davor"},
		{"2020-08-05 13:14:15", "2020-07-26 13:14:15", "1 Woche danach"},

		{"2020-08-05 13:14:15", "2020-08-05 14:14:15", "1 Stunde davor"},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:15", "1 Stunde danach"},
		{"2020-08-05 13:14:15", "2020-08-05 23:14:15", "10 Stunden davor"},
		{"2020-08-05 13:14:15", "2020-08-05 03:14:15", "10 Stunden danach"},

		{"2020-08-05 13:14:15", "2020-08-05 13:15:15", "1 Minute davor"},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:15", "1 Minute danach"},
		{"2020-08-05 13:14:15", "2020-08-05 13:24:15", "10 Minuten davor"},
		{"2020-08-05 13:14:15", "2020-08-05 13:04:15", "10 Minuten danach"},

		{"2020-08-05 13:14:15", "2020-08-05 13:14:16", "1 Sekunde davor"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:14", "1 Sekunde danach"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:25", "10 Sekunden davor"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:05", "10 Sekunden danach"},
	}

	for index, test := range tests {
		c1 := Parse(test.input1)
		c2 := Parse(test.input2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.SetLocale(german).DiffForHumans(c2), "Current test index is "+strconv.Itoa(index))
	}
}
