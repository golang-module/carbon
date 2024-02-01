package calendar

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGregorian_Year(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{NewGregorian(time.Time{})},
			want: 0,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local))},
			want: 2020,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.args.g.Year(), "args(%v)", tt.args.g)
		})
	}
}

func TestGregorian_Month(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{NewGregorian(time.Time{})},
			want: 0,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local))},
			want: 8,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.args.g.Month(), "args(%v)", tt.args.g)
		})
	}
}

func TestGregorian_Week(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{NewGregorian(time.Time{})},
			want: 0,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local))},
			want: 3,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.args.g.Week(), "args(%v)", tt.args.g)
		})
	}
}

func TestGregorian_Day(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{NewGregorian(time.Time{})},
			want: 0,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local))},
			want: 5,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.args.g.Day(), "args(%v)", tt.args.g)
		})
	}
}

func TestGregorian_Hour(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{NewGregorian(time.Time{})},
			want: 0,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local))},
			want: 13,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.args.g.Hour(), "args(%v)", tt.args.g)
		})
	}
}

func TestGregorian_Minute(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{NewGregorian(time.Time{})},
			want: 0,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local))},
			want: 14,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.args.g.Minute(), "args(%v)", tt.args.g)
		})
	}
}

func TestGregorian_Second(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{NewGregorian(time.Time{})},
			want: 0,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local))},
			want: 15,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.args.g.Second(), "args(%v)", tt.args.g)
		})
	}
}

func TestGregorian_String(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Time{})},
			want: "",
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local))},
			want: "2020-08-05 13:14:15",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.args.g.String(), "args(%v)", tt.args.g)
		})
	}
}

func TestGregorian_Location(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Time{})},
			want: "UTC",
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local))},
			want: "Local",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.args.g.Location().String(), "args(%v)", tt.args.g)
		})
	}
}

func TestGregorian_Date(t *testing.T) {
	type args struct {
		Time time.Time
	}
	tests := []struct {
		args      args
		wantYear  int
		wantMonth int
		wantDay   int
	}{
		{
			args:      args{time.Time{}},
			wantYear:  0,
			wantMonth: 0,
			wantDay:   0,
		},
		{
			args:      args{time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local)},
			wantYear:  2020,
			wantMonth: 8,
			wantDay:   5,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			g := NewGregorian(tt.args.Time)
			gotYear, gotMonth, gotDay := g.Date()
			assert.Equalf(t, tt.wantYear, gotYear, "Date()")
			assert.Equalf(t, tt.wantMonth, gotMonth, "Date()")
			assert.Equalf(t, tt.wantDay, gotDay, "Date()")
		})
	}
}

func TestGregorian_Clock(t *testing.T) {
	type args struct {
		Time time.Time
	}
	tests := []struct {
		args       args
		wantHour   int
		wantMinute int
		wantSecond int
	}{
		{
			args:       args{time.Time{}},
			wantHour:   0,
			wantMinute: 0,
			wantSecond: 0,
		},
		{
			args:       args{time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local)},
			wantHour:   13,
			wantMinute: 14,
			wantSecond: 15,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			g := NewGregorian(tt.args.Time)
			gotHour, goMinute, gotSecond := g.Clock()
			assert.Equalf(t, tt.wantHour, gotHour, "Clock()")
			assert.Equalf(t, tt.wantMinute, goMinute, "Clock()")
			assert.Equalf(t, tt.wantSecond, gotSecond, "Clock()")
		})
	}
}

func TestGregorian_IsZero(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Time{})},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local))},
			want: false,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.args.g.IsZero(), "args(%v)", tt.args.g)
		})
	}
}

func TestGregorian_IsLeapYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Time{})},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2021, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.args.g.IsLeapYear(), "args(%v)", tt.args.g)
		})
	}
}
