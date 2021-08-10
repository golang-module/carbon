package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var turkish = "tr"

func TestLang_Turkish_Constellation(t *testing.T) {
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

		{"2020-01-05", "Oğlak"},
		{"2020-02-05", "Kova"},
		{"2020-03-05", "Balık"},
		{"2020-04-05", "Koç"},
		{"2020-05-05", "Boğa"},
		{"2020-06-05", "İkizler"},
		{"2020-07-05", "Yengeç"},
		{"2020-08-05", "Aslan"},
		{"2020-09-05", "Başak"},
		{"2020-10-05", "Terazi"},
		{"2020-11-05", "Akrep"},
		{"2020-12-05", "Yay"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(turkish)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Constellation(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Turkish_Season(t *testing.T) {
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

		{"2020-01-05", "Kış"},
		{"2020-02-05", "Kış"},
		{"2020-03-05", "İlkbahar"},
		{"2020-04-05", "İlkbahar"},
		{"2020-05-05", "İlkbahar"},
		{"2020-06-05", "Yaz"},
		{"2020-07-05", "Yaz"},
		{"2020-08-05", "Yaz"},
		{"2020-09-05", "Sonbahar"},
		{"2020-10-05", "Sonbahar"},
		{"2020-11-05", "Sonbahar"},
		{"2020-12-05", "Kış"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(turkish)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Season(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Turkish_ToMonthString(t *testing.T) {
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

		{"2020-01-05", "Ocak"},
		{"2020-02-05", "Şubat"},
		{"2020-03-05", "Mart"},
		{"2020-04-05", "Nisan"},
		{"2020-05-05", "Mayıs"},
		{"2020-06-05", "Haziran"},
		{"2020-07-05", "Temmuz"},
		{"2020-08-05", "Ağustos"},
		{"2020-09-05", "Eylül"},
		{"2020-10-05", "Ekim"},
		{"2020-11-05", "Kasım"},
		{"2020-12-05", "Aralık"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(turkish)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToMonthString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Turkish_ToShortMonthString(t *testing.T) {
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

		{"2020-01-05", "Oca"},
		{"2020-02-05", "Şub"},
		{"2020-03-05", "Mar"},
		{"2020-04-05", "Nis"},
		{"2020-05-05", "May"},
		{"2020-06-05", "Haz"},
		{"2020-07-05", "Tem"},
		{"2020-08-05", "Ağu"},
		{"2020-09-05", "Eyl"},
		{"2020-10-05", "Eki"},
		{"2020-11-05", "Kas"},
		{"2020-12-05", "Ara"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(turkish)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortMonthString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Turkish_ToWeekString(t *testing.T) {
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

		{"2020-08-01", "Cumartesi"},
		{"2020-08-02", "Pazar"},
		{"2020-08-03", "Pazartesi"},
		{"2020-08-04", "Salı"},
		{"2020-08-05", "Çarşamba"},
		{"2020-08-06", "Perşembe"},
		{"2020-08-07", "Cuma"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(turkish)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToWeekString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Turkish_ToShortWeekString(t *testing.T) {
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

		{"2020-08-01", "Cts"},
		{"2020-08-02", "Paz"},
		{"2020-08-03", "Pts"},
		{"2020-08-04", "Sal"},
		{"2020-08-05", "Çrş"},
		{"2020-08-06", "Per"},
		{"2020-08-07", "Cum"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(turkish)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortWeekString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Turkish_DiffForHumans(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   string // 输入值
		input2   string // 输入值
		expected string // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", "az önce"},
		{"2020-08-05 13:14:15", "2021-08-05 13:14:15", "bir yıl önce"},
		{"2020-08-05 13:14:15", "2019-08-05 13:14:15", "bir yıl sonra"},
		{"2020-08-05 13:14:15", "2030-08-05 13:14:15", "10 yıl önce"},
		{"2020-08-05 13:14:15", "2010-08-05 13:14:15", "10 yıl sonra"},

		{"2020-08-05 13:14:15", "2020-09-05 13:14:15", "bir ay önce"},
		{"2020-08-05 13:14:15", "2020-07-05 13:14:15", "bir ay sonra"},
		{"2020-08-05 13:14:15", "2021-06-05 13:14:15", "10 ay önce"},
		{"2020-08-05 13:14:15", "2019-10-05 13:14:15", "10 ay sonra"},

		{"2020-08-05 13:14:15", "2020-08-06 13:14:15", "bir gün önce"},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", "bir gün sonra"},
		{"2020-08-05 13:14:15", "2020-08-15 13:14:15", "bir hafta önce"},
		{"2020-08-05 13:14:15", "2020-07-26 13:14:15", "bir hafta sonra"},

		{"2020-08-05 13:14:15", "2020-08-05 14:14:15", "bir saat önce"},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:15", "bir saat sonra"},
		{"2020-08-05 13:14:15", "2020-08-05 23:14:15", "10 saat önce"},
		{"2020-08-05 13:14:15", "2020-08-05 03:14:15", "10 saat sonra"},

		{"2020-08-05 13:14:15", "2020-08-05 13:15:15", "bir dakika önce"},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:15", "bir dakika sonra"},
		{"2020-08-05 13:14:15", "2020-08-05 13:24:15", "10 dakika önce"},
		{"2020-08-05 13:14:15", "2020-08-05 13:04:15", "10 dakika sonra"},

		{"2020-08-05 13:14:15", "2020-08-05 13:14:16", "bir saniye önce"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:14", "bir saniye sonra"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:25", "10 saniye önce"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:05", "10 saniye sonra"},
	}

	for index, test := range tests {
		c1 := Parse(test.input1)
		c2 := Parse(test.input2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.SetLocale(turkish).DiffForHumans(c2), "Current test index is "+strconv.Itoa(index))
	}
}
