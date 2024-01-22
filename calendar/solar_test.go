package calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSolar_Year(t *testing.T) {
	type args struct {
		s Solar
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{NewSolar(time.Time{})},
			want: 0,
		},
		{
			name: "2020",
			args: args{NewSolar(time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local))},
			want: 2020,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.args.s.Year(), "NewSolar(%v)", tt.args.s)
		})
	}
}

func TestSolar_Month(t *testing.T) {
	type args struct {
		s Solar
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{NewSolar(time.Time{})},
			want: 0,
		},
		{
			name: "2020",
			args: args{NewSolar(time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local))},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.args.s.Month(), "NewSolar(%v)", tt.args.s)
		})
	}
}

func TestSolar_Day(t *testing.T) {
	type args struct {
		s Solar
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{NewSolar(time.Time{})},
			want: 0,
		},
		{
			name: "2020",
			args: args{NewSolar(time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local))},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.args.s.Day(), "NewSolar(%v)", tt.args.s)
		})
	}
}

func TestSolar_Hour(t *testing.T) {
	type args struct {
		s Solar
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{NewSolar(time.Time{})},
			want: 0,
		},
		{
			name: "2020",
			args: args{NewSolar(time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local))},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.args.s.Hour(), "NewSolar(%v)", tt.args.s)
		})
	}
}

func TestSolar_Minute(t *testing.T) {
	type args struct {
		s Solar
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{NewSolar(time.Time{})},
			want: 0,
		},
		{
			name: "2020",
			args: args{NewSolar(time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local))},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.args.s.Minute(), "NewSolar(%v)", tt.args.s)
		})
	}
}

func TestSolar_Second(t *testing.T) {
	type args struct {
		s Solar
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{NewSolar(time.Time{})},
			want: 0,
		},
		{
			name: "2020",
			args: args{NewSolar(time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local))},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.args.s.Second(), "NewSolar(%v)", tt.args.s)
		})
	}
}

func TestSolar_String(t *testing.T) {
	type args struct {
		s Solar
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0",
			args: args{NewSolar(time.Time{})},
			want: "",
		},
		{
			name: "2020",
			args: args{NewSolar(time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local))},
			want: "2020-08-05 13:14:15",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.args.s.String(), "NewSolar(%v)", tt.args.s)
		})
	}
}

func TestSolar_Location(t *testing.T) {
	type args struct {
		s Solar
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "UTC",
			args: args{NewSolar(time.Time{})},
			want: "UTC",
		},
		{
			name: "Local",
			args: args{NewSolar(time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local))},
			want: "Local",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.args.s.Location().String(), "NewSolar(%v)", tt.args.s)
		})
	}
}

func TestSolar_Date(t *testing.T) {
	type args struct {
		Time time.Time
	}
	tests := []struct {
		name      string
		args      args
		wantYear  int
		wantMonth int
		wantDay   int
	}{
		{
			name:      "zeroTime",
			args:      args{time.Time{}},
			wantYear:  0,
			wantMonth: 0,
			wantDay:   0,
		},
		{
			name:      "normalDate",
			args:      args{time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local)},
			wantYear:  2020,
			wantMonth: 8,
			wantDay:   5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Solar{
				Time: tt.args.Time,
			}
			gotYear, gotMonth, gotDay := s.Date()
			assert.Equalf(t, tt.wantYear, gotYear, "Date()")
			assert.Equalf(t, tt.wantMonth, gotMonth, "Date()")
			assert.Equalf(t, tt.wantDay, gotDay, "Date()")
		})
	}
}

func TestSolar_Clock(t *testing.T) {
	type args struct {
		Time time.Time
	}
	tests := []struct {
		name       string
		args       args
		wantHour   int
		wantMinute int
		wantSecond int
	}{
		{
			name:       "zeroTime",
			args:       args{time.Time{}},
			wantHour:   0,
			wantMinute: 0,
			wantSecond: 0,
		},
		{
			name:       "normalTime",
			args:       args{time.Date(2020, 8, 5, 13, 14, 15, 0, time.Local)},
			wantHour:   13,
			wantMinute: 14,
			wantSecond: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Solar{
				Time: tt.args.Time,
			}
			gotHour, goMinute, gotSecond := s.Clock()
			assert.Equalf(t, tt.wantHour, gotHour, "Clock()")
			assert.Equalf(t, tt.wantMinute, goMinute, "Clock()")
			assert.Equalf(t, tt.wantSecond, gotSecond, "Clock()")
		})
	}
}
