package persian

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGregorian_ToPersian(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{FromGregorian(time.Time{})},
			want: "",
		},
		{
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			args: args{FromGregorian(time.Date(1800, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "1178-10-11 00:00:00",
		},
		{
			args: args{FromGregorian(time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "1402-10-11 00:00:00",
		},
		{
			args: args{FromGregorian(time.Date(2024, 8, 5, 12, 0, 0, 0, time.Local))},
			want: "1403-05-15 12:00:00",
		},
		{
			args: args{FromGregorian(time.Date(2024, 12, 31, 23, 59, 59, 0, time.Local))},
			want: "1403-10-11 23:59:59",
		},
		{
			args: args{FromGregorian(time.Date(2645, 3, 21, 0, 0, 0, 0, time.Local))},
			want: "2024-01-01 00:00:00",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToPersian().String(), "args{%v}", tt.args.g)
		})
	}
}

func TestPersian_ToGregorian(t *testing.T) {
	type args struct {
		p Persian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{Persian{}},
			want: "",
		},
		{
			args: args{FromPersian(0, 0, 0, 0, 0, 0)},
			want: "",
		},
		{
			args: args{FromPersian(1395, 1, 1, 0, 0, 0)},
			want: "2016-03-20 00:00:00",
		},
		{
			args: args{FromPersian(1402, 10, 11, 0, 0, 0)},
			want: "2024-01-01 00:00:00",
		},
		{
			args: args{FromPersian(1403, 5, 15, 12, 0, 0)},
			want: "2024-08-05 12:00:00",
		},
		{
			args: args{FromPersian(1403, 10, 11, 23, 59, 59)},
			want: "2024-12-31 23:59:59",
		},

		{
			args: args{FromPersian(2024, 1, 1, 0, 0, 0)},
			want: "2645-03-21 00:00:00",
		},
		{
			args: args{FromPersian(2100, 12, 31, 23, 59, 59)},
			want: "2722-03-22 23:59:59",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.p).ToGregorian().String(), "args{%v}", tt.args.p)
		})
	}
}

func TestPersian_Year(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{FromGregorian(time.Time{})},
			want: 0,
		},
		{
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: 0,
		},
		{
			args: args{FromGregorian(time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local))},
			want: 1402,
		},
		{
			args: args{FromGregorian(time.Date(2024, 8, 5, 12, 0, 0, 0, time.Local))},
			want: 1403,
		},
		{
			args: args{FromGregorian(time.Date(2024, 12, 31, 23, 59, 59, 0, time.Local))},
			want: 1403,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToPersian().Year(), "args{%v}", tt.args.g)
		})
	}
}

func TestPersian_Month(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{FromGregorian(time.Time{})},
			want: 0,
		},
		{
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: 0,
		},
		{
			args: args{FromGregorian(time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local))},
			want: 10,
		},
		{
			args: args{FromGregorian(time.Date(2024, 8, 5, 12, 0, 0, 0, time.Local))},
			want: 5,
		},
		{
			args: args{FromGregorian(time.Date(2024, 12, 31, 23, 59, 59, 0, time.Local))},
			want: 10,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToPersian().Month(), "args{%v}", tt.args.g)
		})
	}
}

func TestPersian_Day(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{FromGregorian(time.Time{})},
			want: 0,
		},
		{
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: 0,
		},
		{
			args: args{FromGregorian(time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local))},
			want: 11,
		},
		{
			args: args{FromGregorian(time.Date(2024, 8, 5, 12, 0, 0, 0, time.Local))},
			want: 15,
		},
		{
			args: args{FromGregorian(time.Date(2024, 12, 31, 23, 59, 59, 0, time.Local))},
			want: 11,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToPersian().Day(), "args{%v}", tt.args.g)
		})
	}
}

func TestPersian_Hour(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{FromGregorian(time.Time{})},
			want: 0,
		},
		{
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: 0,
		},
		{
			args: args{FromGregorian(time.Date(2020, 8, 19, 13, 14, 15, 0, time.Local))},
			want: 13,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToPersian().Hour(), "args{%v}", tt.args.g)
		})
	}
}

func TestPersian_Minute(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{FromGregorian(time.Time{})},
			want: 0,
		},
		{
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: 0,
		},
		{
			args: args{FromGregorian(time.Date(2020, 8, 19, 13, 14, 15, 0, time.Local))},
			want: 14,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToPersian().Minute(), "args{%v}", tt.args.g)
		})
	}
}

func TestPersian_Second(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{FromGregorian(time.Time{})},
			want: 0,
		},
		{
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: 0,
		},
		{
			args: args{FromGregorian(time.Date(2020, 8, 19, 13, 14, 15, 0, time.Local))},
			want: 15,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToPersian().Second(), "args{%v}", tt.args.g)
		})
	}
}

func TestPersian_ToMonthString(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{FromGregorian(time.Time{})},
			want: "",
		},
		{
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			args: args{FromGregorian(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "دی",
		},
		{
			args: args{FromGregorian(time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local))},
			want: "بهمن",
		},
		{
			args: args{FromGregorian(time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local))},
			want: "اسفند",
		},
		{
			args: args{FromGregorian(time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local))},
			want: "فروردین",
		},
		{
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "اردیبهشت",
		},
		{
			args: args{FromGregorian(time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local))},
			want: "خرداد",
		},
		{
			args: args{FromGregorian(time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local))},
			want: "تیر",
		},
		{
			args: args{FromGregorian(time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local))},
			want: "مرداد",
		},
		{
			args: args{FromGregorian(time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local))},
			want: "شهریور",
		},
		{
			args: args{FromGregorian(time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local))},
			want: "مهر",
		},
		{
			args: args{FromGregorian(time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local))},
			want: "آبان",
		},
		{
			args: args{FromGregorian(time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local))},
			want: "آذر",
		},
		{
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "دی",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToPersian().ToMonthString(), "args{%v}", tt.args.g)
		})
	}
}

func TestPersian_ToWeekString(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{FromGregorian(time.Time{})},
			want: "",
		},
		{
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "",
		},
		{
			args: args{FromGregorian(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "چهارشنبه",
		},
		{
			args: args{FromGregorian(time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local))},
			want: "شنبه",
		},
		{
			args: args{FromGregorian(time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local))},
			want: "یکشنبه",
		},
		{
			args: args{FromGregorian(time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local))},
			want: "چهارشنبه",
		},
		{
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "جمعه",
		},
		{
			args: args{FromGregorian(time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local))},
			want: "دوشنبه",
		},
		{
			args: args{FromGregorian(time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local))},
			want: "چهارشنبه",
		},
		{
			args: args{FromGregorian(time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local))},
			want: "شنبه",
		},
		{
			args: args{FromGregorian(time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local))},
			want: "سه شنبه",
		},
		{
			args: args{FromGregorian(time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local))},
			want: "پنجشنبه",
		},
		{
			args: args{FromGregorian(time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local))},
			want: "یکشنبه",
		},
		{
			args: args{FromGregorian(time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local))},
			want: "سه شنبه",
		},
		{
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "جمعه",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToPersian().ToWeekString(), "args{%v}", tt.args.g)
		})
	}
}

func TestPersian_IsLeapYear(t *testing.T) {
	type args struct {
		p Persian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{Persian{}},
			want: false,
		},
		{
			args: args{FromPersian(0, 0, 0, 0, 0, 0)},
			want: false,
		},
		{
			args: args{FromPersian(1394, 1, 1, 0, 0, 0)},
			want: false,
		},
		{
			args: args{FromPersian(1395, 1, 1, 0, 0, 0)},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.p).IsLeapYear(), "args{%v}", tt.args.p)
		})
	}
}

func TestName(t *testing.T) {
	p := FromGregorian(time.Date(1800, 1, 1, 0, 0, 0, 0, time.Local)).ToPersian()
	fmt.Println(p)
}
