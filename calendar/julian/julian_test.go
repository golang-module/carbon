package julian

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGregorian_ToJulian(t *testing.T) {
	type args struct {
		g Gregorian
	}
	type want struct {
		j Julian
	}
	tests := []struct {
		args args
		want want
	}{
		{
			args: args{FromGregorian(time.Time{})},
			want: want{FromJulian(0)},
		},
		{
			args: args{FromGregorian(time.Date(2024, 1, 23, 0, 0, 0, 0, time.Local))},
			want: want{FromJulian(2460332.5)},
		},
		{
			args: args{FromGregorian(time.Date(2024, 1, 23, 0, 0, 0, 0, time.Local))},
			want: want{FromJulian(60332)},
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want.j, (tt.args.g).ToJulian(), "args{%v}", tt.args.g)
		})
	}
}

func TestJulian_ToGregorian(t *testing.T) {
	type args struct {
		j Julian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{FromJulian(0)},
			want: "",
		},
		{
			args: args{FromJulian(2460332.5)},
			want: "2024-01-23 00:00:00",
		},
		{
			args: args{FromJulian(60332)},
			want: "2024-01-23 00:00:00",
		},
		{
			args: args{FromJulian(2460333)},
			want: "2024-01-23 12:00:00",
		},
		{
			args: args{FromJulian(60332.5)},
			want: "2024-01-23 12:00:00",
		},

		{
			args: args{FromJulian(2460333.051563)},
			want: "2024-01-23 13:14:15",
		},
		{
			args: args{FromJulian(60332.551563)},
			want: "2024-01-23 13:14:15",
		},
		{
			args: args{FromJulian(60232.5)},
			want: "2023-10-15 12:00:00",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.j).ToGregorian().String(), "args{%v}", tt.args.j)
		})
	}
}

func TestGregorian_JD(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want float64
	}{
		{
			args: args{FromGregorian(time.Date(2024, 1, 23, 0, 0, 0, 0, time.Local))},
			want: 2460332.5,
		},
		{
			args: args{FromGregorian(time.Date(2024, 1, 23, 12, 0, 0, 0, time.Local))},
			want: 2460333,
		},
		{
			args: args{FromGregorian(time.Date(2024, 1, 23, 13, 14, 15, 0, time.Local))},
			want: 2460333.051563,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToJulian().JD(6), "args{%v}", tt.args.g)
		})
	}
}

func TestGregorian_MJD(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want float64
	}{
		{
			args: args{FromGregorian(time.Date(2024, 1, 23, 0, 0, 0, 0, time.Local))},
			want: 60332,
		},
		{
			args: args{FromGregorian(time.Date(2024, 1, 23, 12, 0, 0, 0, time.Local))},
			want: 60332.5,
		},
		{
			args: args{FromGregorian(time.Date(2024, 1, 23, 13, 14, 15, 0, time.Local))},
			want: 60332.551563,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToJulian().MJD(6), "args{%v}", tt.args.g)
		})
	}
}
