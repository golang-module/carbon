package julian

import (
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
		name string
		args args
		want want
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Time{})},
			want: want{FromJulian(0)},
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(1800, 1, 1, 0, 0, 0, 0, time.Local))},
			want: want{FromJulian(2378496.5)},
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2024, 1, 23, 0, 0, 0, 0, time.Local))},
			want: want{FromJulian(2460332.5)},
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2024, 1, 23, 0, 0, 0, 0, time.Local))},
			want: want{FromJulian(60332)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want.j, (tt.args.g).ToJulian(), "args{%v}", tt.args.g)
		})
	}
}

func TestJulian_ToGregorian(t *testing.T) {
	type args struct {
		j Julian
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{FromJulian(0)},
			want: "",
		},
		{
			name: "case2",
			args: args{FromJulian(2460332.5)},
			want: "2024-01-23 00:00:00",
		},
		{
			name: "case3",
			args: args{FromJulian(60332)},
			want: "2024-01-23 00:00:00",
		},
		{
			name: "case4",
			args: args{FromJulian(2460333)},
			want: "2024-01-23 12:00:00",
		},
		{
			name: "case5",
			args: args{FromJulian(60332.5)},
			want: "2024-01-23 12:00:00",
		},
		{
			name: "case6",
			args: args{FromJulian(2460333.051563)},
			want: "2024-01-23 13:14:15",
		},
		{
			name: "case7",
			args: args{FromJulian(60332.551563)},
			want: "2024-01-23 13:14:15",
		},
		{
			name: "case8",
			args: args{FromJulian(60232.5)},
			want: "2023-10-15 12:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.j).ToGregorian().String(), "args{%v}", tt.args.j)
		})
	}
}

func TestGregorian_JD(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Date(1800, 1, 1, 0, 0, 0, 0, time.Local))},
			want: 2378496.5,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(2024, 1, 23, 0, 0, 0, 0, time.Local))},
			want: 2460332.5,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2024, 1, 23, 12, 0, 0, 0, time.Local))},
			want: 2460333,
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(2024, 1, 23, 13, 14, 15, 0, time.Local))},
			want: 2460333.051563,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToJulian().JD(6), "args{%v}", tt.args.g)
		})
	}
}

func TestGregorian_MJD(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Date(2024, 1, 23, 0, 0, 0, 0, time.Local))},
			want: 60332,
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Date(2024, 1, 23, 12, 0, 0, 0, time.Local))},
			want: 60332.5,
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Date(2024, 1, 23, 13, 14, 15, 0, time.Local))},
			want: 60332.551563,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToJulian().MJD(6), "args{%v}", tt.args.g)
		})
	}
}
