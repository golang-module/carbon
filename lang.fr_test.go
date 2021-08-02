package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var french = "fr"

func TestLang_French_Constellation(t *testing.T) {
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

		{6, "2020-01-05", "Capricorne"},
		{7, "2020-02-05", "Verseau"},
		{8, "2020-03-05", "Poisson"},
		{9, "2020-04-05", "Bélier"},
		{10, "2020-05-05", "Taureau"},
		{11, "2020-06-05", "Gémeaux"},
		{12, "2020-07-05", "Cancer"},
		{13, "2020-08-05", "Lion"},
		{14, "2020-09-05", "Vierge"},
		{15, "2020-10-05", "Balance"},
		{16, "2020-11-05", "Scorpion"},
		{17, "2020-12-05", "Sagittaire"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(french)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Constellation(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_French_Season(t *testing.T) {
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

		{6, "2020-01-05", "Hiver"},
		{7, "2020-02-05", "Hiver"},
		{8, "2020-03-05", "Printemps"},
		{9, "2020-04-05", "Printemps"},
		{10, "2020-05-05", "Printemps"},
		{11, "2020-06-05", "Été"},
		{12, "2020-07-05", "Été"},
		{13, "2020-08-05", "Été"},
		{14, "2020-09-05", "Automne"},
		{15, "2020-10-05", "Automne"},
		{16, "2020-11-05", "Automne"},
		{17, "2020-12-05", "Hiver"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(french)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Season(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_French_ToMonthString(t *testing.T) {
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

		{6, "2020-01-05", "Janvier"},
		{7, "2020-02-05", "Février"},
		{8, "2020-03-05", "Mars"},
		{9, "2020-04-05", "Avril"},
		{10, "2020-05-05", "Mai"},
		{11, "2020-06-05", "Juin"},
		{12, "2020-07-05", "Juillet"},
		{13, "2020-08-05", "Août"},
		{14, "2020-09-05", "Septembre"},
		{15, "2020-10-05", "Octobre"},
		{16, "2020-11-05", "Novembre"},
		{17, "2020-12-05", "Décembre"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(french)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToMonthString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_French_ToShortMonthString(t *testing.T) {
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

		{6, "2020-01-05", "Janv"},
		{7, "2020-02-05", "Févr"},
		{8, "2020-03-05", "Mars"},
		{9, "2020-04-05", "Avr"},
		{10, "2020-05-05", "Mai"},
		{11, "2020-06-05", "Juin"},
		{12, "2020-07-05", "Juil"},
		{13, "2020-08-05", "Août"},
		{14, "2020-09-05", "Sept"},
		{15, "2020-10-05", "Oct"},
		{16, "2020-11-05", "Nov"},
		{17, "2020-12-05", "Déc"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(french)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortMonthString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_French_ToWeekString(t *testing.T) {
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

		{6, "2020-08-01", "Samedi"},
		{7, "2020-08-02", "Dimanche"},
		{8, "2020-08-03", "Lundi"},
		{9, "2020-08-04", "Mardi"},
		{10, "2020-08-05", "Mercredi"},
		{11, "2020-08-06", "Jeudi"},
		{12, "2020-08-07", "Vendredi"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(french)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToWeekString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_French_ToShortWeekString(t *testing.T) {
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

		{6, "2020-08-01", "Sam"},
		{7, "2020-08-02", "Dim"},
		{8, "2020-08-03", "Lun"},
		{9, "2020-08-04", "Mar"},
		{10, "2020-08-05", "Mer"},
		{11, "2020-08-06", "Jeu"},
		{12, "2020-08-07", "Ven"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(french)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortWeekString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_French_DiffForHumans(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input1   string // 输入值
		input2   string // 输入值
		expected string // 期望值
	}{
		{1, "2020-08-05 13:14:15", "2020-08-05 13:14:15", "juste"},
		{2, "2020-08-05 13:14:15", "2021-08-05 13:14:15", "1 an avant"},
		{3, "2020-08-05 13:14:15", "2019-08-05 13:14:15", "1 an après"},
		{4, "2020-08-05 13:14:15", "2030-08-05 13:14:15", "10 ans avant"},
		{5, "2020-08-05 13:14:15", "2010-08-05 13:14:15", "10 ans après"},

		{6, "2020-08-05 13:14:15", "2020-09-05 13:14:15", "1 mois avant"},
		{7, "2020-08-05 13:14:15", "2020-07-05 13:14:15", "1 mois après"},
		{8, "2020-08-05 13:14:15", "2021-06-05 13:14:15", "10 mois avant"},
		{9, "2020-08-05 13:14:15", "2019-10-05 13:14:15", "10 mois après"},

		{10, "2020-08-05 13:14:15", "2020-08-06 13:14:15", "1 jour avant"},
		{11, "2020-08-05 13:14:15", "2020-08-04 13:14:15", "1 jour après"},
		{12, "2020-08-05 13:14:15", "2020-08-15 13:14:15", "1 semaine avant"},
		{13, "2020-08-05 13:14:15", "2020-07-26 13:14:15", "1 semaine après"},

		{14, "2020-08-05 13:14:15", "2020-08-05 14:14:15", "1 heure avant"},
		{15, "2020-08-05 13:14:15", "2020-08-05 12:14:15", "1 heure après"},
		{16, "2020-08-05 13:14:15", "2020-08-05 23:14:15", "10 heures avant"},
		{17, "2020-08-05 13:14:15", "2020-08-05 03:14:15", "10 heures après"},

		{18, "2020-08-05 13:14:15", "2020-08-05 13:15:15", "1 minute avant"},
		{19, "2020-08-05 13:14:15", "2020-08-05 13:13:15", "1 minute après"},
		{20, "2020-08-05 13:14:15", "2020-08-05 13:24:15", "10 minutes avant"},
		{21, "2020-08-05 13:14:15", "2020-08-05 13:04:15", "10 minutes après"},

		{22, "2020-08-05 13:14:15", "2020-08-05 13:14:16", "1 seconde avant"},
		{23, "2020-08-05 13:14:15", "2020-08-05 13:14:14", "1 seconde après"},
		{24, "2020-08-05 13:14:15", "2020-08-05 13:14:25", "10 secondes avant"},
		{25, "2020-08-05 13:14:15", "2020-08-05 13:14:05", "10 secondes après"},
	}

	for _, test := range tests {
		c1 := Parse(test.input1)
		c2 := Parse(test.input2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.SetLocale(french).DiffForHumans(c2), "Current test id is "+strconv.Itoa(test.id))
	}
}
