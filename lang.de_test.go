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
		id       int    // 测试id
		input    string // 输入值
		expected string // 期望值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},
		{5, "0000-00-00 00:00:00", ""},

		{6, "2020-01-05", "Steinbock"},
		{7, "2020-02-05", "Wassermann"},
		{8, "2020-03-05", "Fisch"},
		{9, "2020-04-05", "Widder"},
		{10, "2020-05-05", "Stier"},
		{11, "2020-06-05", "Zwilling"},
		{12, "2020-07-05", "Krebs"},
		{13, "2020-08-05", "Löwe"},
		{14, "2020-09-05", "Jungfrau"},
		{15, "2020-10-05", "Waage"},
		{16, "2020-11-05", "Skorpion"},
		{17, "2020-12-05", "Schütze"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(german)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Constellation(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_German_Season(t *testing.T) {
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
		{8, "2020-03-05", "Frühling"},
		{9, "2020-04-05", "Frühling"},
		{10, "2020-05-05", "Frühling"},
		{11, "2020-06-05", "Sommer"},
		{12, "2020-07-05", "Sommer"},
		{13, "2020-08-05", "Sommer"},
		{14, "2020-09-05", "Herbst"},
		{15, "2020-10-05", "Herbst"},
		{16, "2020-11-05", "Herbst"},
		{17, "2020-12-05", "Winter"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(german)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Season(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_German_ToMonthString(t *testing.T) {
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

		{6, "2020-01-05", "Januar"},
		{7, "2020-02-05", "Februar"},
		{8, "2020-03-05", "März"},
		{9, "2020-04-05", "April"},
		{10, "2020-05-05", "Mai"},
		{11, "2020-06-05", "Juni"},
		{12, "2020-07-05", "Juli"},
		{13, "2020-08-05", "August"},
		{14, "2020-09-05", "September"},
		{15, "2020-10-05", "Oktober"},
		{16, "2020-11-05", "November"},
		{17, "2020-12-05", "Dezember"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(german)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToMonthString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_German_ToShortMonthString(t *testing.T) {
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
		{8, "2020-03-05", "Mär"},
		{9, "2020-04-05", "Apr"},
		{10, "2020-05-05", "Mai"},
		{11, "2020-06-05", "Jun"},
		{12, "2020-07-05", "Jul"},
		{13, "2020-08-05", "Aug"},
		{14, "2020-09-05", "Sep"},
		{15, "2020-10-05", "Okt"},
		{16, "2020-11-05", "Nov"},
		{17, "2020-12-05", "Dez"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(german)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortMonthString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_German_ToWeekString(t *testing.T) {
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

		{6, "2020-08-01", "Samstag"},
		{7, "2020-08-02", "Sonntag"},
		{8, "2020-08-03", "Montag"},
		{9, "2020-08-04", "Dienstag"},
		{10, "2020-08-05", "Mittwoch"},
		{11, "2020-08-06", "Donnerstag"},
		{12, "2020-08-07", "Freitag"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(german)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToWeekString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_German_ToShortWeekString(t *testing.T) {
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

		{6, "2020-08-01", "Sa"},
		{7, "2020-08-02", "So"},
		{8, "2020-08-03", "Mo"},
		{9, "2020-08-04", "Di"},
		{10, "2020-08-05", "Mi"},
		{11, "2020-08-06", "Do"},
		{12, "2020-08-07", "Fr"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(german)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortWeekString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_German_DiffForHumans(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input1   string // 输入值
		input2   string // 输入值
		expected string // 期望值
	}{
		{1, "2020-08-05 13:14:15", "2020-08-05 13:14:15", "gerade eben"},
		{2, "2020-08-05 13:14:15", "2021-08-05 13:14:15", "1 Jahr davor"},
		{3, "2020-08-05 13:14:15", "2019-08-05 13:14:15", "1 Jahr danach"},
		{4, "2020-08-05 13:14:15", "2030-08-05 13:14:15", "10 Jahre davor"},
		{5, "2020-08-05 13:14:15", "2010-08-05 13:14:15", "10 Jahre danach"},

		{6, "2020-08-05 13:14:15", "2020-09-05 13:14:15", "1 Monat davor"},
		{7, "2020-08-05 13:14:15", "2020-07-05 13:14:15", "1 Monat danach"},
		{8, "2020-08-05 13:14:15", "2021-06-05 13:14:15", "10 Monate davor"},
		{9, "2020-08-05 13:14:15", "2019-10-05 13:14:15", "10 Monate danach"},

		{10, "2020-08-05 13:14:15", "2020-08-06 13:14:15", "1 Tag davor"},
		{11, "2020-08-05 13:14:15", "2020-08-04 13:14:15", "1 Tag danach"},
		{12, "2020-08-05 13:14:15", "2020-08-15 13:14:15", "1 Woche davor"},
		{13, "2020-08-05 13:14:15", "2020-07-26 13:14:15", "1 Woche danach"},

		{14, "2020-08-05 13:14:15", "2020-08-05 14:14:15", "1 Stunde davor"},
		{15, "2020-08-05 13:14:15", "2020-08-05 12:14:15", "1 Stunde danach"},
		{16, "2020-08-05 13:14:15", "2020-08-05 23:14:15", "10 Stunden davor"},
		{17, "2020-08-05 13:14:15", "2020-08-05 03:14:15", "10 Stunden danach"},

		{18, "2020-08-05 13:14:15", "2020-08-05 13:15:15", "1 Minute davor"},
		{19, "2020-08-05 13:14:15", "2020-08-05 13:13:15", "1 Minute danach"},
		{20, "2020-08-05 13:14:15", "2020-08-05 13:24:15", "10 Minuten davor"},
		{21, "2020-08-05 13:14:15", "2020-08-05 13:04:15", "10 Minuten danach"},

		{22, "2020-08-05 13:14:15", "2020-08-05 13:14:16", "1 Sekunde davor"},
		{23, "2020-08-05 13:14:15", "2020-08-05 13:14:14", "1 Sekunde danach"},
		{24, "2020-08-05 13:14:15", "2020-08-05 13:14:25", "10 Sekunden davor"},
		{25, "2020-08-05 13:14:15", "2020-08-05 13:14:05", "10 Sekunden danach"},
	}

	for _, test := range tests {
		c1 := Parse(test.input1)
		c2 := Parse(test.input2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.SetLocale(german).DiffForHumans(c2), "Current test id is "+strconv.Itoa(test.id))
	}
}
