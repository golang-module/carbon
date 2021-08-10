package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var spanish = "es"

func TestLang_Spanish_Constellation(t *testing.T) {
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

		{"2020-01-05", "Capricornio"},
		{"2020-02-05", "Acuario"},
		{"2020-03-05", "Piscis"},
		{"2020-04-05", "Aries"},
		{"2020-05-05", "Tauro"},
		{"2020-06-05", "Geminis"},
		{"2020-07-05", "Cancer"},
		{"2020-08-05", "Leo"},
		{"2020-09-05", "Virgo"},
		{"2020-10-05", "Libra"},
		{"2020-11-05", "Escorpio"},
		{"2020-12-05", "Sagitario"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(spanish)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Constellation(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Spanish_Season(t *testing.T) {
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

		{"2020-01-05", "Invierno"},
		{"2020-02-05", "Invierno"},
		{"2020-03-05", "Primavera"},
		{"2020-04-05", "Primavera"},
		{"2020-05-05", "Primavera"},
		{"2020-06-05", "Verano"},
		{"2020-07-05", "Verano"},
		{"2020-08-05", "Verano"},
		{"2020-09-05", "Otoño"},
		{"2020-10-05", "Otoño"},
		{"2020-11-05", "Otoño"},
		{"2020-12-05", "Invierno"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(spanish)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Season(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Spanish_ToMonthString(t *testing.T) {
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

		{"2020-01-05", "Enero"},
		{"2020-02-05", "Febrero"},
		{"2020-03-05", "Marzo"},
		{"2020-04-05", "Abril"},
		{"2020-05-05", "Mayo"},
		{"2020-06-05", "Junio"},
		{"2020-07-05", "Julio"},
		{"2020-08-05", "Agosto"},
		{"2020-09-05", "Septiembre"},
		{"2020-10-05", "Octubre"},
		{"2020-11-05", "Noviembre"},
		{"2020-12-05", "Diciembre"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(spanish)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToMonthString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Spanish_ToShortMonthString(t *testing.T) {
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

		{"2020-01-05", "Ene"},
		{"2020-02-05", "Feb"},
		{"2020-03-05", "Mar"},
		{"2020-04-05", "Abr"},
		{"2020-05-05", "May"},
		{"2020-06-05", "Jun"},
		{"2020-07-05", "Jul"},
		{"2020-08-05", "Ago"},
		{"2020-09-05", "Sep"},
		{"2020-10-05", "Oct"},
		{"2020-11-05", "Nov"},
		{"2020-12-05", "Dic"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(spanish)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortMonthString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Spanish_ToWeekString(t *testing.T) {
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

		{"2020-08-01", "Sábado"},
		{"2020-08-02", "Domingo"},
		{"2020-08-03", "Lunes"},
		{"2020-08-04", "Martes"},
		{"2020-08-05", "Miércoles"},
		{"2020-08-06", "Jueves"},
		{"2020-08-07", "Viernes"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(spanish)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToWeekString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Spanish_ToShortWeekString(t *testing.T) {
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

		{"2020-08-01", "Sab"},
		{"2020-08-02", "Dom"},
		{"2020-08-03", "Lun"},
		{"2020-08-04", "Mar"},
		{"2020-08-05", "Mie"},
		{"2020-08-06", "Jue"},
		{"2020-08-07", "Vie"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(spanish)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortWeekString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLang_Spanish_DiffForHumans(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   string // 输入值
		input2   string // 输入值
		expected string // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", "ahora"},
		{"2020-08-05 13:14:15", "2021-08-05 13:14:15", "1 año antes"},
		{"2020-08-05 13:14:15", "2019-08-05 13:14:15", "1 año después"},
		{"2020-08-05 13:14:15", "2030-08-05 13:14:15", "10 años antes"},
		{"2020-08-05 13:14:15", "2010-08-05 13:14:15", "10 años después"},

		{"2020-08-05 13:14:15", "2020-09-05 13:14:15", "1 mes antes"},
		{"2020-08-05 13:14:15", "2020-07-05 13:14:15", "1 mes después"},
		{"2020-08-05 13:14:15", "2021-06-05 13:14:15", "10 meses antes"},
		{"2020-08-05 13:14:15", "2019-10-05 13:14:15", "10 meses después"},

		{"2020-08-05 13:14:15", "2020-08-06 13:14:15", "1 día antes"},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", "1 día después"},
		{"2020-08-05 13:14:15", "2020-08-15 13:14:15", "1 semana antes"},
		{"2020-08-05 13:14:15", "2020-07-26 13:14:15", "1 semana después"},

		{"2020-08-05 13:14:15", "2020-08-05 14:14:15", "1 hora antes"},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:15", "1 hora después"},
		{"2020-08-05 13:14:15", "2020-08-05 23:14:15", "10 horas antes"},
		{"2020-08-05 13:14:15", "2020-08-05 03:14:15", "10 horas después"},

		{"2020-08-05 13:14:15", "2020-08-05 13:15:15", "1 minuto antes"},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:15", "1 minuto después"},
		{"2020-08-05 13:14:15", "2020-08-05 13:24:15", "10 minutos antes"},
		{"2020-08-05 13:14:15", "2020-08-05 13:04:15", "10 minutos después"},

		{"2020-08-05 13:14:15", "2020-08-05 13:14:16", "1 segundo antes"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:14", "1 segundo después"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:25", "10 segundos antes"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:05", "10 segundos después"},
	}

	for index, test := range tests {
		c1 := Parse(test.input1)
		c2 := Parse(test.input2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.SetLocale(spanish).DiffForHumans(c2), "Current test index is "+strconv.Itoa(index))
	}
}
