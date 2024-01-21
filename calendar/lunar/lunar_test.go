package lunar

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateFromSolar(t *testing.T) {
	type args struct {
		s Solar
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normalMonth",
			args: args{NewSolar(time.Date(2024, 1, 21, 0, 0, 0, 0, time.Local))},
			want: "2023-12-11 00:00:00",
		},
		{
			name: "notLeapMonth",
			args: args{NewSolar(time.Date(2023, 3, 2, 0, 0, 0, 0, time.Local))},
			want: "2023-02-11 00:00:00",
		},
		{
			name: "isLeapMonth",
			args: args{NewSolar(time.Date(2023, 4, 1, 0, 0, 0, 0, time.Local))},
			want: "2023-02-11 00:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CreateFromSolar(tt.args.s).String(), "CreateFromSolar(%v)", tt.args.s)
		})
	}
}

func TestCreateFromLunar(t *testing.T) {
	type args struct {
		l Lunar
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normalMonth",
			args: args{NewLunar(2023, 12, 11, 0, 0, 0, false)},
			want: "2024-01-21 00:00:00",
		},
		{
			name: "notLeapMonth",
			args: args{NewLunar(2023, 2, 11, 0, 0, 0, false)},
			want: "2023-03-02 00:00:00",
		},
		{
			name: "isLeapMonth",
			args: args{NewLunar(2023, 2, 11, 0, 0, 0, true)},
			want: "2023-04-01 00:00:00",
		},
		{
			name: "ltMinYear",
			args: args{NewLunar(1800, 2, 11, 0, 0, 0, false)},
			want: "",
		},
		{
			name: "gtMaxYear",
			args: args{NewLunar(2500, 2, 11, 0, 0, 0, false)},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CreateFromLunar(tt.args.l).Solar.String(), "CreateFromLunar(%v)", tt.args.l)
		})
	}
}

func TestLunar_Animal(t *testing.T) {
	type args struct {
		s Solar
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "zeroTime",
			args: args{NewSolar(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			name: "ratYear",
			args: args{NewSolar(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "鼠",
		},
		{
			name: "ratYear",
			args: args{NewSolar(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "鼠",
		},
		{
			name: "oxYear",
			args: args{NewSolar(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: "牛",
		},
		{
			name: "tigerYear",
			args: args{NewSolar(time.Date(2010, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "虎",
		},
		{
			name: "rabbitYear",
			args: args{NewSolar(time.Date(2011, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "兔",
		},
		{
			name: "dragonYear",
			args: args{NewSolar(time.Date(2012, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "龙",
		},
		{
			name: "snakeYear",
			args: args{NewSolar(time.Date(2013, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "蛇",
		},
		{
			name: "horseYear",
			args: args{NewSolar(time.Date(2014, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "马",
		},
		{
			name: "goatYear",
			args: args{NewSolar(time.Date(2015, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "羊",
		},
		{
			name: "monkeyYear",
			args: args{NewSolar(time.Date(2016, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "猴",
		},
		{
			name: "roosterYear",
			args: args{NewSolar(time.Date(2017, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "鸡",
		},
		{
			name: "dogYear",
			args: args{NewSolar(time.Date(2018, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "狗",
		},
		{
			name: "pigYear",
			args: args{NewSolar(time.Date(2019, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "猪",
		},
		{
			name: "ratYear",
			args: args{NewSolar(time.Date(2020, 5, 23, 0, 0, 0, 0, time.Local))},
			want: "鼠",
		},
		{
			name: "ratYear",
			args: args{NewSolar(time.Date(2020, 6, 21, 0, 0, 0, 0, time.Local))},
			want: "鼠",
		},
		{
			name: "ratYear",
			args: args{NewSolar(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "鼠",
		},
		{
			name: "goatYear",
			args: args{NewSolar(time.Date(2021, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "牛",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CreateFromSolar(tt.args.s).Animal(), "CreateFromSolar(%v)", tt.args.s)
		})
	}
}

func TestLunar_Festival(t *testing.T) {
	type args struct {
		s Solar
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "zeroTime",
			args: args{NewSolar(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			name: "notFestival",
			args: args{NewSolar(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			name: "springFestival",
			args: args{NewSolar(time.Date(2021, 2, 12, 0, 0, 0, 0, time.Local))},
			want: "春节",
		},
		{
			name: "springFestival",
			args: args{NewSolar(time.Date(2021, 2, 26, 0, 0, 0, 0, time.Local))},
			want: "元宵节",
		},
		{
			name: "dragonBoatFestival",
			args: args{NewSolar(time.Date(2021, 6, 14, 0, 0, 0, 0, time.Local))},
			want: "端午节",
		},
		{
			name: "chineseValentine'sDay",
			args: args{NewSolar(time.Date(2021, 8, 14, 0, 0, 0, 0, time.Local))},
			want: "七夕节",
		},
		{
			name: "zhongyuanFestival",
			args: args{NewSolar(time.Date(2021, 8, 22, 0, 0, 0, 0, time.Local))},
			want: "中元节",
		},
		{
			name: "zhongyuanFestival",
			args: args{NewSolar(time.Date(2021, 9, 21, 0, 0, 0, 0, time.Local))},
			want: "中秋节",
		},
		{
			name: "doubleNinthFestival",
			args: args{NewSolar(time.Date(2021, 10, 14, 0, 0, 0, 0, time.Local))},
			want: "重阳节",
		},
		{
			name: "hanyiFestival",
			args: args{NewSolar(time.Date(2021, 11, 5, 0, 0, 0, 0, time.Local))},
			want: "寒衣节",
		},
		{
			name: "xiaYuanFestival",
			args: args{NewSolar(time.Date(2021, 11, 19, 0, 0, 0, 0, time.Local))},
			want: "下元节",
		},
		{
			name: "labaFestival",
			args: args{NewSolar(time.Date(2022, 1, 10, 0, 0, 0, 0, time.Local))},
			want: "腊八节",
		},
		{
			name: "lunarYear",
			args: args{NewSolar(time.Date(2022, 1, 25, 0, 0, 0, 0, time.Local))},
			want: "小年",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CreateFromSolar(tt.args.s).Festival(), "CreateFromSolar(%v)", tt.args.s)
		})
	}
}

func TestLunar_DateTime(t *testing.T) {
	tests := []struct {
		input                                  string
		year, month, day, hour, minute, second int
	}{
		0: {"", 0, 0, 0, 0, 0, 0},
		1: {"0", 0, 0, 0, 0, 0, 0},
		2: {"0000-00-00", 0, 0, 0, 0, 0, 0},
		3: {"00:00:00", 0, 0, 0, 0, 0, 0},
		4: {"0000-00-00 00:00:00", 0, 0, 0, 0, 0, 0},

		5: {"2020-08-05 13:14:15", 2020, 6, 16, 13, 14, 15},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02 15:04:05", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		year, month, day, hour, minute, second := lunar.DateTime()
		assert.Equal(t, test.year, year, "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, test.month, month, "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, test.day, day, "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, test.hour, hour, "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, test.minute, minute, "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, test.second, second, "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_Date(t *testing.T) {
	tests := []struct {
		input            string
		year, month, day int
	}{
		0: {"", 0, 0, 0},
		1: {"0", 0, 0, 0},
		2: {"0000-00-00", 0, 0, 0},
		3: {"00:00:00", 0, 0, 0},
		4: {"0000-00-00 00:00:00", 0, 0, 0},

		5: {"2020-08-05 13:14:15", 2020, 6, 16},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02 15:04:05", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		year, month, day := lunar.Date()
		assert.Equal(t, test.year, year, "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, test.month, month, "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, test.day, day, "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_Time(t *testing.T) {
	tests := []struct {
		input                string
		hour, minute, second int
	}{
		0: {"", 0, 0, 0},
		1: {"0", 0, 0, 0},
		2: {"0000-00-00", 0, 0, 0},
		3: {"00:00:00", 0, 0, 0},
		4: {"0000-00-00 00:00:00", 0, 0, 0},

		5: {"2020-08-05 13:14:15", 13, 14, 15},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02 15:04:05", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		hour, minute, second := lunar.Time()
		assert.Equal(t, test.hour, hour, "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, test.minute, minute, "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, test.second, second, "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_Year(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		0: {"", 0},
		1: {"0", 0},
		2: {"0000-00-00", 0},
		3: {"00:00:00", 0},
		4: {"0000-00-00 00:00:00", 0},

		5:  {"2020-04-23", 2020},
		6:  {"2020-05-01", 2020},
		7:  {"2020-08-05", 2020},
		8:  {"2021-01-01", 2020},
		9:  {"2021-05-12", 2021},
		10: {"2021-07-07", 2021},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.Year(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_Month(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		0: {"", 0},
		1: {"0", 0},
		2: {"0000-00-00", 0},
		3: {"00:00:00", 0},
		4: {"0000-00-00 00:00:00", 0},

		5:  {"2021-01-05", 11},
		6:  {"2021-02-05", 12},
		7:  {"2021-03-05", 1},
		8:  {"2021-04-05", 2},
		9:  {"2021-05-05", 3},
		10: {"2021-06-05", 4},
		11: {"2021-07-05", 5},
		12: {"2021-08-05", 6},
		13: {"2021-09-05", 7},
		14: {"2021-10-05", 8},
		15: {"2021-11-05", 10},
		16: {"2021-12-05", 11},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.Month(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_Day(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		0: {"", 0},
		1: {"0", 0},
		2: {"0000-00-00", 0},
		3: {"00:00:00", 0},
		4: {"0000-00-00 00:00:00", 0},

		5:  {"2020-08-01", 12},
		6:  {"2020-08-02", 13},
		7:  {"2020-08-03", 14},
		8:  {"2020-08-04", 15},
		9:  {"2020-08-05", 16},
		10: {"2020-08-06", 17},
		11: {"2020-08-07", 18},
		12: {"2020-08-08", 19},
		13: {"2020-08-09", 20},
		14: {"2020-08-10", 21},
		15: {"2020-08-11", 22},
		16: {"2020-08-12", 23},
		17: {"2020-08-13", 24},
		18: {"2020-08-14", 25},
		19: {"2020-08-15", 26},
		20: {"2020-08-16", 27},
		21: {"2020-08-17", 28},
		22: {"2020-08-18", 29},
		23: {"2020-08-19", 1},
		24: {"2020-08-20", 2},
		25: {"2020-08-21", 3},
		26: {"2020-08-22", 4},
		27: {"2020-08-23", 5},
		28: {"2020-08-24", 6},
		29: {"2020-08-25", 7},
		30: {"2020-08-26", 8},
		31: {"2020-08-27", 9},
		32: {"2020-08-28", 10},
		33: {"2020-08-29", 11},
		34: {"2020-08-30", 12},
		35: {"2020-08-31", 13},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.Day(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_LeapMonth(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		0: {"", 0},
		1: {"0", 0},
		2: {"0000-00-00", 0},
		3: {"00:00:00", 0},
		4: {"0000-00-00 00:00:00", 0},

		5: {"2020-04-23", 4},
		6: {"2020-05-01", 4},
		7: {"2020-08-05", 4},
		8: {"2021-07-07", 0},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.LeapMonth(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_ToYearString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		0: {"", ""},
		1: {"0", ""},
		2: {"0000-00-00", ""},
		3: {"00:00:00", ""},
		4: {"0000-00-00 00:00:00", ""},

		5:  {"2020-04-23", "二零二零"},
		6:  {"2020-05-01", "二零二零"},
		7:  {"2020-08-05", "二零二零"},
		8:  {"2021-01-01", "二零二零"},
		9:  {"2021-05-12", "二零二一"},
		10: {"2021-07-07", "二零二一"},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.ToYearString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_ToMonthString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		0: {"", ""},
		1: {"0", ""},
		2: {"0000-00-00", ""},
		3: {"00:00:00", ""},
		4: {"0000-00-00 00:00:00", ""},

		5:  {"2020-01-01", "腊月"},
		6:  {"2020-02-01", "正月"},
		7:  {"2020-03-01", "二月"},
		8:  {"2020-04-01", "三月"},
		9:  {"2020-04-23", "闰四月"},
		10: {"2020-05-01", "闰四月"},
		11: {"2020-06-01", "闰四月"},
		12: {"2020-07-01", "五月"},
		13: {"2020-07-07", "五月"},
		14: {"2020-08-01", "六月"},
		15: {"2020-09-01", "七月"},
		16: {"2020-10-01", "八月"},
		17: {"2020-11-01", "九月"},
		18: {"2020-12-01", "十月"},
		19: {"2021-01-01", "十一月"},
		20: {"2021-02-01", "腊月"},
		21: {"2021-05-12", "四月"},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.ToMonthString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_ToDayString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		0: {"", ""},
		1: {"0", ""},
		2: {"0000-00-00", ""},
		3: {"00:00:00", ""},
		4: {"0000-00-00 00:00:00", ""},

		5:  {"2020-01-01", "初七"},
		6:  {"2020-02-01", "初八"},
		7:  {"2020-03-01", "初八"},
		8:  {"2020-04-01", "初九"},
		9:  {"2020-04-23", "初一"},
		10: {"2020-05-01", "初九"},
		11: {"2020-06-01", "初十"},
		12: {"2020-07-01", "十一"},
		13: {"2020-08-01", "十二"},
		14: {"2020-09-01", "十四"},
		15: {"2020-10-01", "十五"},
		16: {"2020-11-01", "十六"},
		17: {"2020-12-01", "十七"},
		18: {"2021-01-03", "二十"},
		19: {"2021-01-05", "廿二"},
		20: {"2021-04-11", "三十"},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.ToDayString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_String(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		0: {"", ""},
		1: {"0", ""},
		2: {"0000-00-00", ""},
		3: {"00:00:00", ""},
		4: {"0000-00-00 00:00:00", ""},

		5:  {"2020-01-01", "2019-12-07 00:00:00"},
		6:  {"2020-02-01", "2020-01-08 00:00:00"},
		7:  {"2020-03-01", "2020-02-08 00:00:00"},
		8:  {"2020-04-01", "2020-03-09 00:00:00"},
		9:  {"2020-04-23", "2020-04-01 00:00:00"},
		10: {"2020-05-01", "2020-04-09 00:00:00"},
		11: {"2020-06-01", "2020-04-10 00:00:00"},
		12: {"2020-07-01", "2020-05-11 00:00:00"},
		13: {"2020-08-01", "2020-06-12 00:00:00"},
		14: {"2020-09-01", "2020-07-14 00:00:00"},
		15: {"2020-10-01", "2020-08-15 00:00:00"},
		16: {"2020-11-01", "2020-09-16 00:00:00"},
		17: {"2020-12-01", "2020-10-17 00:00:00"},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, fmt.Sprintf("%s", lunar), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_ToDateString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		0: {"", ""},
		1: {"0", ""},
		2: {"0000-00-00", ""},
		3: {"00:00:00", ""},
		4: {"0000-00-00 00:00:00", ""},

		5:  {"2020-01-01", "二零一九年腊月初七"},
		6:  {"2020-02-01", "二零二零年正月初八"},
		7:  {"2020-03-01", "二零二零年二月初八"},
		8:  {"2020-04-01", "二零二零年三月初九"},
		9:  {"2020-04-23", "二零二零年闰四月初一"},
		10: {"2020-05-01", "二零二零年闰四月初九"},
		11: {"2020-06-01", "二零二零年闰四月初十"},
		12: {"2020-07-01", "二零二零年五月十一"},
		13: {"2020-08-01", "二零二零年六月十二"},
		14: {"2020-09-01", "二零二零年七月十四"},
		15: {"2020-10-01", "二零二零年八月十五"},
		16: {"2020-11-01", "二零二零年九月十六"},
		17: {"2020-12-01", "二零二零年十月十七"},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsLeapYear(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-04-23", true},
		6: {"2020-05-01", true},
		7: {"2020-08-05", true},
		8: {"2021-01-01", true},
		9: {"2021-07-07", false},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsLeapYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsLeapMonth(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-04-23", true},
		6: {"2020-05-01", true},
		7: {"2020-08-05", false},
		8: {"2021-01-01", false},
		9: {"2021-07-07", false},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsLeapMonth(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsRatYear(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", true},
		6: {"2020-08-05", true},
		7: {"2021-01-01", true},
		8: {"2021-07-07", false},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsRatYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsOxYear(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", true},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsOxYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsTigerYear(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", false},
		9: {"2022-08-05", true},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsTigerYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsRabbitYear(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", false},
		9: {"2023-08-05", true},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsRabbitYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsDragonYear(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", false},
		9: {"2024-08-05", true},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsDragonYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsSnakeYear(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", false},
		9: {"2025-08-05", true},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsSnakeYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsHorseYear(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", false},
		9: {"2026-08-05", true},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsHorseYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsGoatYear(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", false},
		9: {"2027-08-05", true},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsGoatYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsMonkeyYear(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", false},
		9: {"2028-08-05", true},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsMonkeyYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsRoosterYear(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", false},
		9: {"2029-08-05", true},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsRoosterYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsDogYear(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", false},
		9: {"2030-08-05", true},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsDogYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsPigYear(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", false},
		9: {"2031-08-05", true},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsPigYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_DoubleHour(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{

		5:  {"2020-01-05 23:23:45", "子时"},
		6:  {"2020-01-05 00:59:45", "子时"},
		7:  {"2020-02-05 01:00:00", "丑时"},
		8:  {"2020-02-05 03:40:00", "寅时"},
		9:  {"2020-02-05 05:00:00", "卯时"},
		10: {"2020-02-05 07:30:00", "辰时"},
		11: {"2020-02-05 09:00:00", "巳时"},
		12: {"2020-02-05 11:00:00", "午时"},
		13: {"2020-02-05 13:00:00", "未时"},
		14: {"2020-02-05 15:00:00", "申时"},
		15: {"2020-02-05 14:59:00", "未时"},
		16: {"2020-02-05 17:00:00", "酉时"},
		17: {"2020-02-05 19:00:00", "戌时"},
		18: {"2020-02-05 21:00:00", "亥时"},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02 15:04:05", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.DoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsFirstDoubleHour(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 00:00:00", true},
		6: {"2020-04-19 00:59:59", true},
		7: {"2020-08-05 23:00:00", true},
		8: {"2020-08-05 01:00:00", false},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02 15:04:05", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsFirstDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsSecondDoubleHour(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 01:00:00", true},
		6: {"2020-04-19 02:59:59", true},
		7: {"2020-08-05 03:00:00", false},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02 15:04:05", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsSecondDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsThirdDoubleHour(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 03:00:00", true},
		6: {"2020-04-19 04:59:59", true},
		7: {"2020-08-05 05:00:00", false},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02 15:04:05", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsThirdDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsFourthDoubleHour(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 05:00:00", true},
		6: {"2020-04-19 06:59:59", true},
		7: {"2020-08-05 07:00:00", false},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02 15:04:05", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsFourthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsFifthDoubleHour(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 07:00:00", true},
		6: {"2020-04-19 08:59:59", true},
		7: {"2020-08-05 09:00:00", false},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02 15:04:05", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsFifthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsSixthDoubleHour(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 09:00:00", true},
		6: {"2020-04-19 10:59:59", true},
		7: {"2020-08-05 11:00:00", false},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02 15:04:05", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsSixthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsSeventhDoubleHour(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 11:00:00", true},
		6: {"2020-04-19 12:59:59", true},
		7: {"2020-08-05 13:00:00", false},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02 15:04:05", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsSeventhDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsEighthDoubleHour(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 13:00:00", true},
		6: {"2020-04-19 14:59:59", true},
		7: {"2020-08-05 15:00:00", false},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02 15:04:05", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsEighthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsNinthDoubleHour(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 15:00:00", true},
		6: {"2020-04-19 16:59:59", true},
		7: {"2020-08-05 17:00:00", false},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02 15:04:05", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsNinthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsTenthDoubleHour(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 17:00:00", true},
		6: {"2020-04-19 18:59:59", true},
		7: {"2020-08-05 19:00:00", false},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02 15:04:05", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsTenthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsEleventhDoubleHour(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 19:00:00", true},
		6: {"2020-04-19 20:59:59", true},
		7: {"2020-08-05 21:00:00", false},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02 15:04:05", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsEleventhDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsTwelfthDoubleHour(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 21:00:00", true},
		6: {"2020-04-19 22:59:59", true},
		7: {"2020-08-05 23:00:00", false},
	}

	for index, test := range tests {
		tm, _ := time.Parse("2006-01-02 15:04:05", test.input)
		lunar := CreateFromSolar(NewSolar(tm))

		assert.Equal(t, test.expected, lunar.IsTwelfthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}
