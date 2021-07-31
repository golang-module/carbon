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
		id       int    // 测试id
		input    string // 输入值
		expected string // 期望值
	}{
		{1, "", ""},
		{2, "0", ""},
		{3, "0000-00-00", ""},
		{4, "00:00:00", ""},
		{5, "0000-00-00 00:00:00", ""},

		{6, "2020-01-05", "Capricornio"},
		{7, "2020-02-05", "Acuario"},
		{8, "2020-03-05", "Piscis"},
		{9, "2020-04-05", "Aries"},
		{10, "2020-05-05", "Tauro"},
		{11, "2020-06-05", "Geminis"},
		{12, "2020-07-05", "Cancer"},
		{13, "2020-08-05", "Leo"},
		{14, "2020-09-05", "Virgo"},
		{15, "2020-10-05", "Libra"},
		{16, "2020-11-05", "Escorpio"},
		{17, "2020-12-05", "Sagitario"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(spanish)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Constellation(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_Spanish_Season(t *testing.T) {
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

		{6, "2020-01-05", "Invierno"},
		{7, "2020-02-05", "Invierno"},
		{8, "2020-03-05", "Primavera"},
		{9, "2020-04-05", "Primavera"},
		{10, "2020-05-05", "Primavera"},
		{11, "2020-06-05", "Verano"},
		{12, "2020-07-05", "Verano"},
		{13, "2020-08-05", "Verano"},
		{14, "2020-09-05", "Otoño"},
		{15, "2020-10-05", "Otoño"},
		{16, "2020-11-05", "Otoño"},
		{17, "2020-12-05", "Invierno"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(spanish)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Season(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_Spanish_ToMonthString(t *testing.T) {
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

		{6, "2020-01-05", "Enero"},
		{7, "2020-02-05", "Febrero"},
		{8, "2020-03-05", "Marzo"},
		{9, "2020-04-05", "Abril"},
		{10, "2020-05-05", "Mayo"},
		{11, "2020-06-05", "Junio"},
		{12, "2020-07-05", "Julio"},
		{13, "2020-08-05", "Agosto"},
		{14, "2020-09-05", "Septiembre"},
		{15, "2020-10-05", "Octubre"},
		{16, "2020-11-05", "Noviembre"},
		{17, "2020-12-05", "Diciembre"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(spanish)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToMonthString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_Spanish_ToShortMonthString(t *testing.T) {
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

		{6, "2020-01-05", "Ene"},
		{7, "2020-02-05", "Feb"},
		{8, "2020-03-05", "Mar"},
		{9, "2020-04-05", "Abr"},
		{10, "2020-05-05", "May"},
		{11, "2020-06-05", "Jun"},
		{12, "2020-07-05", "Jul"},
		{13, "2020-08-05", "Ago"},
		{14, "2020-09-05", "Sep"},
		{15, "2020-10-05", "Oct"},
		{16, "2020-11-05", "Nov"},
		{17, "2020-12-05", "Dic"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(spanish)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortMonthString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_Spanish_ToWeekString(t *testing.T) {
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

		{6, "2020-08-01", "Sábado"},
		{7, "2020-08-02", "Domingo"},
		{8, "2020-08-03", "Lunes"},
		{9, "2020-08-04", "Martes"},
		{10, "2020-08-05", "Miércoles"},
		{11, "2020-08-06", "Jueves"},
		{12, "2020-08-07", "Viernes"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(spanish)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToWeekString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_Spanish_ToShortWeekString(t *testing.T) {
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

		{6, "2020-08-01", "Sab"},
		{7, "2020-08-02", "Dom"},
		{8, "2020-08-03", "Lun"},
		{9, "2020-08-04", "Mar"},
		{10, "2020-08-05", "Mie"},
		{11, "2020-08-06", "Jue"},
		{12, "2020-08-07", "Vie"},
	}

	for _, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).SetLocale(spanish)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortWeekString(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestLang_Spanish_DiffForHumans(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input1   string // 输入值
		input2   string // 输入值
		expected string // 期望值
	}{
		{1, "2020-08-05 13:14:15", "2020-08-05 13:14:15", "ahora"},
		{2, "2020-08-05 13:14:15", "2021-08-05 13:14:15", "1 año antes"},
		{3, "2020-08-05 13:14:15", "2019-08-05 13:14:15", "1 año después"},
		{4, "2020-08-05 13:14:15", "2030-08-05 13:14:15", "10 años antes"},
		{5, "2020-08-05 13:14:15", "2010-08-05 13:14:15", "10 años después"},

		{6, "2020-08-05 13:14:15", "2020-09-05 13:14:15", "1 mes antes"},
		{7, "2020-08-05 13:14:15", "2020-07-05 13:14:15", "1 mes después"},
		{8, "2020-08-05 13:14:15", "2021-06-05 13:14:15", "10 meses antes"},
		{9, "2020-08-05 13:14:15", "2019-10-05 13:14:15", "10 meses después"},

		{10, "2020-08-05 13:14:15", "2020-08-06 13:14:15", "1 día antes"},
		{11, "2020-08-05 13:14:15", "2020-08-04 13:14:15", "1 día después"},
		{12, "2020-08-05 13:14:15", "2020-08-15 13:14:15", "1 semana antes"},
		{13, "2020-08-05 13:14:15", "2020-07-26 13:14:15", "1 semana después"},

		{14, "2020-08-05 13:14:15", "2020-08-05 14:14:15", "1 hora antes"},
		{15, "2020-08-05 13:14:15", "2020-08-05 12:14:15", "1 hora después"},
		{16, "2020-08-05 13:14:15", "2020-08-05 23:14:15", "10 horas antes"},
		{17, "2020-08-05 13:14:15", "2020-08-05 03:14:15", "10 horas después"},

		{18, "2020-08-05 13:14:15", "2020-08-05 13:15:15", "1 minuto antes"},
		{19, "2020-08-05 13:14:15", "2020-08-05 13:13:15", "1 minuto después"},
		{20, "2020-08-05 13:14:15", "2020-08-05 13:24:15", "10 minutos antes"},
		{21, "2020-08-05 13:14:15", "2020-08-05 13:04:15", "10 minutos después"},

		{22, "2020-08-05 13:14:15", "2020-08-05 13:14:16", "1 segundo antes"},
		{23, "2020-08-05 13:14:15", "2020-08-05 13:14:14", "1 segundo después"},
		{24, "2020-08-05 13:14:15", "2020-08-05 13:14:25", "10 segundos antes"},
		{25, "2020-08-05 13:14:15", "2020-08-05 13:14:05", "10 segundos después"},
	}

	for _, test := range tests {
		c1 := Parse(test.input1)
		c2 := Parse(test.input2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.SetLocale(spanish).DiffForHumans(c2), "Current test id is "+strconv.Itoa(test.id))
	}
}
