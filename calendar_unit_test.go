package carbon

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFromLunar1(t *testing.T) {
	l := CreateFromLunar(2023, 12, 11, 0, 0, 0, false)
	assert.Nil(t, l.Error)
	assert.Equal(t, "2024-01-21 00:00:00", l.String())
}

func TestCarbon_Lunar(t *testing.T) {
	type args struct {
		c Carbon
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{Parse("xxx")},
			want: "",
		},
		{
			args: args{Parse("2024-01-18 00:00:00")},
			want: "2023-12-08 00:00:00",
		},
		{
			args: args{Parse("2024-01-21 00:00:00")},
			want: "2023-12-11 00:00:00",
		},
		{
			args: args{Parse("2024-01-24 12:00:00")},
			want: "2023-12-14 12:00:00",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.args.c.Lunar().String(), "args{%v}", tt.args.c)
		})
	}
}

func TestCreateFromLunar(t *testing.T) {
	type args struct {
		year, month, day, hour, minute, second int
		isLeapMonth                            bool
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{year: 2023, month: 12, day: 11, isLeapMonth: false},
			want: "2024-01-21 00:00:00",
		},
		{
			args: args{year: 2023, month: 12, day: 8, isLeapMonth: false},
			want: "2024-01-18 00:00:00",
		},
		{
			args: args{year: 2023, month: 12, day: 14, hour: 12, isLeapMonth: false},
			want: "2024-01-24 12:00:00",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, CreateFromLunar(tt.args.year, tt.args.month, tt.args.day, tt.args.hour, tt.args.minute, tt.args.second, tt.args.isLeapMonth).ToDateTimeString(), "args{%v}", tt.args)
		})
	}
}

func TestCarbon_Julian(t *testing.T) {
	type args struct {
		c Carbon
	}
	tests := []struct {
		args    args
		wantJD  float64
		wantMJD float64
	}{
		{
			args:    args{Parse("xxx")},
			wantJD:  0,
			wantMJD: 0,
		},
		{
			args:    args{Parse("2024-01-24 12:00:00")},
			wantJD:  2460334,
			wantMJD: 60333.5,
		},
		{
			args:    args{CreateFromJulian(2460334)},
			wantJD:  2460334,
			wantMJD: 60333.5,
		},
		{
			args:    args{CreateFromJulian(60333.5)},
			wantJD:  2460334,
			wantMJD: 60333.5,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.wantJD, tt.args.c.Julian().JD(), "Julian().JD()")
			assert.Equalf(t, tt.wantMJD, tt.args.c.Julian().MJD(), "Julian().MJD()")
		})
	}
}

func TestCreateFromJulian(t *testing.T) {
	type args struct {
		j float64
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{j: 2460334},
			want: "2024-01-24 12:00:00",
		},
		{
			args: args{j: 60333.5},
			want: "2024-01-24 12:00:00",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, CreateFromJulian(tt.args.j).ToDateTimeString(), "args{%v}", tt.args.j)
		})
	}
}

func TestName(t *testing.T) {
	r1 := Parse("2024-01-24 13:14:15").Julian().MJD()
	fmt.Println(r1)

	r2 := CreateFromJulian(2460334.051563).ToDateTimeString()
	fmt.Println(r2)
}
