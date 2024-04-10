package lunar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGregorian_ToLunar(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Date(2024, 1, 21, 0, 0, 0, 0, time.Local))},
			want: "2023-12-11 00:00:00",
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(2023, 3, 2, 0, 0, 0, 0, time.Local))},
			want: "2023-02-11 00:00:00",
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2023, 4, 1, 0, 0, 0, 0, time.Local))},
			want: "2023-02-11 00:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().String(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_ToGregorian(t *testing.T) {
	type args struct {
		l Lunar
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{Lunar{}},
			want: "",
		},
		{
			name: "case2",
			args: args{FromLunar(0, 0, 0, 0, 0, 0, false)},
			want: "",
		},
		{
			name: "case3",
			args: args{FromLunar(2023, 12, 11, 0, 0, 0, false)},
			want: "2024-01-21 00:00:00",
		},
		{
			name: "case4",
			args: args{FromLunar(2023, 2, 11, 0, 0, 0, false)},
			want: "2023-03-02 00:00:00",
		},
		{
			name: "case5",
			args: args{FromLunar(2023, 2, 11, 0, 0, 0, true)},
			want: "2023-04-01 00:00:00",
		},
		{
			name: "case6",
			args: args{FromLunar(1800, 2, 11, 0, 0, 0, false)},
			want: "",
		},
		{
			name: "case7",
			args: args{FromLunar(2500, 2, 11, 0, 0, 0, false)},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.l).ToGregorian().String(), "args{%v}", tt.args.l)
		})
	}
}

func TestLunar_Animal(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Time{})},
			want: "",
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "鼠",
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "鼠",
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: "牛",
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2010, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "虎",
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2011, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "兔",
		},
		{
			name: "case8",
			args: args{FromGregorian(time.Date(2012, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "龙",
		},
		{
			name: "case9",
			args: args{FromGregorian(time.Date(2013, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "蛇",
		},
		{
			name: "case10",
			args: args{FromGregorian(time.Date(2014, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "马",
		},
		{
			name: "case11",
			args: args{FromGregorian(time.Date(2015, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "羊",
		},
		{
			name: "case12",
			args: args{FromGregorian(time.Date(2016, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "猴",
		},
		{
			name: "case13",
			args: args{FromGregorian(time.Date(2017, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "鸡",
		},
		{
			name: "case14",
			args: args{FromGregorian(time.Date(2018, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "狗",
		},
		{
			name: "case15",
			args: args{FromGregorian(time.Date(2019, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "猪",
		},
		{
			name: "case16",
			args: args{FromGregorian(time.Date(2020, 5, 23, 0, 0, 0, 0, time.Local))},
			want: "鼠",
		},
		{
			name: "case17",
			args: args{FromGregorian(time.Date(2020, 6, 21, 0, 0, 0, 0, time.Local))},
			want: "鼠",
		},
		{
			name: "case18",
			args: args{FromGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "鼠",
		},
		{
			name: "case19",
			args: args{FromGregorian(time.Date(2021, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "牛",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().Animal(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_Festival(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Time{})},
			want: "",
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2021, 2, 12, 0, 0, 0, 0, time.Local))},
			want: "春节",
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2021, 2, 26, 0, 0, 0, 0, time.Local))},
			want: "元宵节",
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2021, 6, 14, 0, 0, 0, 0, time.Local))},
			want: "端午节",
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2021, 8, 14, 0, 0, 0, 0, time.Local))},
			want: "七夕节",
		},
		{
			name: "case8",
			args: args{FromGregorian(time.Date(2021, 8, 22, 0, 0, 0, 0, time.Local))},
			want: "中元节",
		},
		{
			name: "case9",
			args: args{FromGregorian(time.Date(2021, 9, 21, 0, 0, 0, 0, time.Local))},
			want: "中秋节",
		},
		{
			name: "case10",
			args: args{FromGregorian(time.Date(2021, 10, 14, 0, 0, 0, 0, time.Local))},
			want: "重阳节",
		},
		{
			name: "case11",
			args: args{FromGregorian(time.Date(2021, 11, 5, 0, 0, 0, 0, time.Local))},
			want: "寒衣节",
		},
		{
			name: "case12",
			args: args{FromGregorian(time.Date(2021, 11, 19, 0, 0, 0, 0, time.Local))},
			want: "下元节",
		},
		{
			name: "case13",
			args: args{FromGregorian(time.Date(2022, 1, 10, 0, 0, 0, 0, time.Local))},
			want: "腊八节",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().Festival(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_Year(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Time{})},
			want: 0,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: 0,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: 2020,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2021, 5, 1, 0, 0, 0, 0, time.Local))},
			want: 2021,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().Year(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_Month(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Time{})},
			want: 0,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: 0,
		},
		{
			args: args{FromGregorian(time.Date(2021, 3, 5, 0, 0, 0, 0, time.Local))},
			want: 1,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2021, 4, 5, 0, 0, 0, 0, time.Local))},
			want: 2,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2021, 5, 5, 0, 0, 0, 0, time.Local))},
			want: 3,
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2021, 6, 5, 0, 0, 0, 0, time.Local))},
			want: 4,
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2021, 7, 5, 0, 0, 0, 0, time.Local))},
			want: 5,
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2021, 8, 5, 0, 0, 0, 0, time.Local))},
			want: 6,
		},
		{
			name: "case8",
			args: args{FromGregorian(time.Date(2021, 9, 5, 0, 0, 0, 0, time.Local))},
			want: 7,
		},
		{
			name: "case9",
			args: args{FromGregorian(time.Date(2021, 10, 5, 0, 0, 0, 0, time.Local))},
			want: 8,
		},
		{
			name: "case10",
			args: args{FromGregorian(time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local))},
			want: 9,
		},
		{
			name: "case11",
			args: args{FromGregorian(time.Date(2021, 11, 5, 0, 0, 0, 0, time.Local))},
			want: 10,
		},
		{
			name: "case12",
			args: args{FromGregorian(time.Date(2022, 12, 5, 0, 0, 0, 0, time.Local))},
			want: 11,
		},
		{
			name: "case13",
			args: args{FromGregorian(time.Date(2022, 1, 5, 0, 0, 0, 0, time.Local))},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().Month(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_Day(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Time{})},
			want: 0,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: 0,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 8, 19, 0, 0, 0, 0, time.Local))},
			want: 1,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 8, 20, 0, 0, 0, 0, time.Local))},
			want: 2,
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2020, 8, 21, 0, 0, 0, 0, time.Local))},
			want: 3,
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2020, 8, 22, 0, 0, 0, 0, time.Local))},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().Day(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_Hour(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Time{})},
			want: 0,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: 0,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 8, 19, 13, 14, 15, 0, time.Local))},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().Hour(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_Minute(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Time{})},
			want: 0,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: 0,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 8, 19, 13, 14, 15, 0, time.Local))},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().Minute(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_Second(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Time{})},
			want: 0,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: 0,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 8, 19, 13, 14, 15, 0, time.Local))},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().Second(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_LeapMonth(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Time{})},
			want: 0,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: 0,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: 4,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2021, 7, 1, 0, 0, 0, 0, time.Local))},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().LeapMonth(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_ToYearString(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Time{})},
			want: "",
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零",
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2021, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二一",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().ToYearString(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_ToMonthString(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Time{})},
			want: "",
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "腊月",
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local))},
			want: "正月",
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local))},
			want: "二月",
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local))},
			want: "三月",
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: "闰四月",
		},
		{
			name: "case8",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "闰四月",
		},
		{
			name: "case9",
			args: args{FromGregorian(time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local))},
			want: "闰四月",
		},
		{
			name: "case10",
			args: args{FromGregorian(time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local))},
			want: "五月",
		},
		{
			name: "case11",
			args: args{FromGregorian(time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local))},
			want: "六月",
		},
		{
			name: "case12",
			args: args{FromGregorian(time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local))},
			want: "七月",
		},
		{
			name: "case13",
			args: args{FromGregorian(time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local))},
			want: "八月",
		},
		{
			name: "case14",
			args: args{FromGregorian(time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local))},
			want: "九月",
		},
		{
			name: "case15",
			args: args{FromGregorian(time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local))},
			want: "十月",
		},
		{
			name: "case16",
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "十一月",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().ToMonthString(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_ToWeekString(t *testing.T) {
	type args struct {
		l Lunar
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{Lunar{}},
			want: "",
		},
		{
			name: "case2",
			args: args{FromLunar(0, 0, 0, 0, 0, 0, false)},
			want: "",
		},
		{
			name: "case3",
			args: args{FromLunar(2023, 12, 20, 0, 0, 0, false)},
			want: "周二",
		},
		{
			name: "case4",
			args: args{FromLunar(2023, 12, 21, 0, 0, 0, false)},
			want: "周三",
		},
		{
			name: "case5",
			args: args{FromLunar(2023, 12, 22, 0, 0, 0, false)},
			want: "周四",
		},
		{
			name: "case6",
			args: args{FromLunar(2023, 12, 23, 0, 0, 0, false)},
			want: "周五",
		},
		{
			name: "case7",
			args: args{FromLunar(2023, 12, 24, 0, 0, 0, false)},
			want: "周六",
		},
		{
			name: "case8",
			args: args{FromLunar(2023, 12, 25, 0, 0, 0, false)},
			want: "周日",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.l).ToWeekString(), "args{%v}", tt.args.l)
		})
	}
}

func TestLunar_ToDayString(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Time{})},
			want: "",
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "初七",
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local))},
			want: "初八",
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local))},
			want: "初八",
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local))},
			want: "初九",
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "初九",
		},
		{
			name: "case8",
			args: args{FromGregorian(time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local))},
			want: "初十",
		},
		{
			name: "case9",
			args: args{FromGregorian(time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local))},
			want: "十一",
		},
		{
			name: "case10",
			args: args{FromGregorian(time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local))},
			want: "十二",
		},
		{
			name: "case11",
			args: args{FromGregorian(time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local))},
			want: "十四",
		},
		{
			name: "case12",
			args: args{FromGregorian(time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local))},
			want: "十五",
		},
		{
			name: "case13",
			args: args{FromGregorian(time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local))},
			want: "十六",
		},
		{
			name: "case14",
			args: args{FromGregorian(time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local))},
			want: "十七",
		},
		{
			name: "case15",
			args: args{FromGregorian(time.Date(2021, 1, 3, 0, 0, 0, 0, time.Local))},
			want: "二十",
		},
		{
			name: "case16",
			args: args{FromGregorian(time.Date(2021, 1, 5, 0, 0, 0, 0, time.Local))},
			want: "廿二",
		},
		{
			name: "case17",
			args: args{FromGregorian(time.Date(2021, 4, 11, 0, 0, 0, 0, time.Local))},
			want: "三十",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().ToDayString(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_String(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Time{})},
			want: "",
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "2019-12-07 00:00:00",
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-01-08 00:00:00",
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-02-08 00:00:00",
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-03-09 00:00:00",
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-04-09 00:00:00",
		},
		{
			name: "case8",
			args: args{FromGregorian(time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-04-10 00:00:00",
		},
		{
			name: "case9",
			args: args{FromGregorian(time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-05-11 00:00:00",
		},
		{
			name: "case10",
			args: args{FromGregorian(time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-06-12 00:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().String(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_ToDateString(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Time{})},
			want: "",
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "二零一九年腊月初七",
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年正月初八",
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年二月初八",
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年三月初九",
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年闰四月初九",
		},
		{
			name: "case8",
			args: args{FromGregorian(time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年闰四月初十",
		},
		{
			name: "case9",
			args: args{FromGregorian(time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年五月十一",
		},
		{
			name: "case10",
			args: args{FromGregorian(time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年六月十二",
		},
		{
			name: "case11",
			args: args{FromGregorian(time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年七月十四",
		},
		{
			name: "case12",
			args: args{FromGregorian(time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年八月十五",
		},
		{
			name: "case13",
			args: args{FromGregorian(time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年九月十六",
		},
		{
			name: "case14",
			args: args{FromGregorian(time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local))},
			want: "二零二零年十月十七",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().ToDateString(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsLeapYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Time{})},
			want: false,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			name: "case8",
			args: args{FromGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case9",
			args: args{FromGregorian(time.Date(2022, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsLeapYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsLeapMonth(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2022, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsLeapMonth(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsRatYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2022, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsRatYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsOxYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2022, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsOxYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsTigerYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2022, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsTigerYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsRabbitYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2023, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsRabbitYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsDragonYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2024, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsDragonYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsSnakeYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2025, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsSnakeYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsHorseYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2026, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsHorseYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsGoatYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2027, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsGoatYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsMonkeyYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2028, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsMonkeyYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsRoosterYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2029, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsRoosterYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsDogYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case1",
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2030, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsDogYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsPigYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2031, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsPigYear(), "args{%v}", tt.args.g)
		})
	}
}
