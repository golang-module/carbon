package persian

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGregorian_ToPersian(t *testing.T) {
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
			args: args{FromGregorian(time.Date(1800, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "1178-10-11 00:00:00",
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "1402-10-11 00:00:00",
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2024, 8, 5, 12, 0, 0, 0, time.Local))},
			want: "1403-05-15 12:00:00",
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2024, 12, 31, 23, 59, 59, 0, time.Local))},
			want: "1403-10-11 23:59:59",
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2645, 3, 21, 0, 0, 0, 0, time.Local))},
			want: "2024-01-01 00:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToPersian().String(), "args{%v}", tt.args.g)
		})
	}
}

func TestPersian_ToGregorian(t *testing.T) {
	type args struct {
		p Persian
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{Persian{}},
			want: "",
		},
		{
			name: "case2",
			args: args{FromPersian(0, 0, 0, 0, 0, 0)},
			want: "",
		},
		{
			name: "case3",
			args: args{FromPersian(1395, 1, 1, 0, 0, 0)},
			want: "2016-03-20 00:00:00",
		},
		{
			name: "case4",
			args: args{FromPersian(1402, 10, 11, 0, 0, 0)},
			want: "2024-01-01 00:00:00",
		},
		{
			name: "case5",
			args: args{FromPersian(1403, 5, 15, 12, 0, 0)},
			want: "2024-08-05 12:00:00",
		},
		{
			name: "case6",
			args: args{FromPersian(1403, 10, 11, 23, 59, 59)},
			want: "2024-12-31 23:59:59",
		},

		{
			name: "case7",
			args: args{FromPersian(2024, 1, 1, 0, 0, 0)},
			want: "2645-03-21 00:00:00",
		},
		{
			name: "case8",
			args: args{FromPersian(2100, 12, 31, 23, 59, 59)},
			want: "2722-03-22 23:59:59",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.p).ToGregorian().String(), "args{%v}", tt.args.p)
		})
	}
}

func TestPersian_Year(t *testing.T) {
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
			args: args{FromGregorian(time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local))},
			want: 1402,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2024, 8, 5, 12, 0, 0, 0, time.Local))},
			want: 1403,
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2024, 12, 31, 23, 59, 59, 0, time.Local))},
			want: 1403,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToPersian().Year(), "args{%v}", tt.args.g)
		})
	}
}

func TestPersian_Month(t *testing.T) {
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
			args: args{FromGregorian(time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local))},
			want: 10,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2024, 8, 5, 12, 0, 0, 0, time.Local))},
			want: 5,
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2024, 12, 31, 23, 59, 59, 0, time.Local))},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToPersian().Month(), "args{%v}", tt.args.g)
		})
	}
}

func TestPersian_Day(t *testing.T) {
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
			args: args{FromGregorian(time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local))},
			want: 11,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2024, 8, 5, 12, 0, 0, 0, time.Local))},
			want: 15,
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2024, 12, 31, 23, 59, 59, 0, time.Local))},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToPersian().Day(), "args{%v}", tt.args.g)
		})
	}
}

func TestPersian_Hour(t *testing.T) {
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
			assert.Equalf(t, tt.want, (tt.args.g).ToPersian().Hour(), "args{%v}", tt.args.g)
		})
	}
}

func TestPersian_Minute(t *testing.T) {
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
			assert.Equalf(t, tt.want, (tt.args.g).ToPersian().Minute(), "args{%v}", tt.args.g)
		})
	}
}

func TestPersian_Second(t *testing.T) {
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
			assert.Equalf(t, tt.want, (tt.args.g).ToPersian().Second(), "args{%v}", tt.args.g)
		})
	}
}

func TestPersian_ToMonthString(t *testing.T) {
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
			want: "دی",
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local))},
			want: "بهمن",
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local))},
			want: "اسفند",
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local))},
			want: "فروردین",
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "اردیبهشت",
		},
		{
			name: "case8",
			args: args{FromGregorian(time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local))},
			want: "خرداد",
		},
		{
			name: "case9",
			args: args{FromGregorian(time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local))},
			want: "تیر",
		},
		{
			name: "case10",
			args: args{FromGregorian(time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local))},
			want: "مرداد",
		},
		{
			name: "case11",
			args: args{FromGregorian(time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local))},
			want: "شهریور",
		},
		{
			name: "case12",
			args: args{FromGregorian(time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local))},
			want: "مهر",
		},
		{
			name: "case13",
			args: args{FromGregorian(time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local))},
			want: "آبان",
		},
		{
			name: "case14",
			args: args{FromGregorian(time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local))},
			want: "آذر",
		},
		{
			name: "case15",
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "دی",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToPersian().ToMonthString(), "args{%v}", tt.args.g)
		})
	}
}

func TestPersian_ToWeekString(t *testing.T) {
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
			want: "چهارشنبه",
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local))},
			want: "شنبه",
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local))},
			want: "یکشنبه",
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local))},
			want: "چهارشنبه",
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "جمعه",
		},
		{
			name: "case8",
			args: args{FromGregorian(time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local))},
			want: "دوشنبه",
		},
		{
			name: "case9",
			args: args{FromGregorian(time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local))},
			want: "چهارشنبه",
		},
		{
			name: "case10",
			args: args{FromGregorian(time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local))},
			want: "شنبه",
		},
		{
			name: "case11",
			args: args{FromGregorian(time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local))},
			want: "سه شنبه",
		},
		{
			name: "case12",
			args: args{FromGregorian(time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local))},
			want: "پنجشنبه",
		},
		{
			name: "case13",
			args: args{FromGregorian(time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local))},
			want: "یکشنبه",
		},
		{
			name: "case14",
			args: args{FromGregorian(time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local))},
			want: "سه شنبه",
		},
		{
			name: "case15",
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "جمعه",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToPersian().ToWeekString(), "args{%v}", tt.args.g)
		})
	}
}

func TestPersian_IsLeapYear(t *testing.T) {
	type args struct {
		p Persian
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{Persian{}},
			want: false,
		},
		{
			name: "cas2",
			args: args{FromPersian(0, 0, 0, 0, 0, 0)},
			want: false,
		},
		{
			name: "case3",
			args: args{FromPersian(1394, 1, 1, 0, 0, 0)},
			want: false,
		},
		{
			name: "case4",
			args: args{FromPersian(1395, 1, 1, 0, 0, 0)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.p).IsLeapYear(), "args{%v}", tt.args.p)
		})
	}
}
