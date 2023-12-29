package vlunar

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSolarToLunar(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Date(2024, 1, 21, 0, 0, 0, 0, time.Local))},
			want: "2023-12-11 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2023, 3, 2, 0, 0, 0, 0, time.Local))},
			want: "2023-02-11 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2023, 4, 1, 0, 0, 0, 0, time.Local))},
			want: "2023-02-11 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2022, 1, 25, 0, 0, 0, 0, time.Local))},
			want: "2021-12-23 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2022, 6, 13, 0, 0, 0, 0, time.Local))},
			want: "2022-05-15 00:00:00",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().String(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunarToGregorian(t *testing.T) {
	type args struct {
		l Lunar
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewLunar(2023, 12, 11, 0, 0, 0, false)},
			want: "2024-01-21 00:00:00",
		},
		{
			args: args{NewLunar(2023, 2, 11, 0, 0, 0, false)},
			want: "2023-03-02 00:00:00",
		},
		{
			args: args{NewLunar(2023, 2, 11, 0, 0, 0, true)},
			want: "2023-04-01 00:00:00",
		},
		{
			args: args{NewLunar(1800, 2, 11, 0, 0, 0, false)},
			want: "1800-03-06 00:00:00",
		},
		{
			args: args{NewLunar(2500, 2, 11, 0, 0, 0, false)},
			want: "2500-03-11 00:00:00",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.l).ToGregorian().String(), "args{%v}", tt.args.l)
		})
	}
}

func TestLunar_Animal(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "Mùi",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "Tý",
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "Tý",
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: "Sửu",
		},
		{
			args: args{NewGregorian(time.Date(2010, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "Dần",
		},
		{
			args: args{NewGregorian(time.Date(2011, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "Mão",
		},
		{
			args: args{NewGregorian(time.Date(2012, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "Thìn",
		},
		{
			args: args{NewGregorian(time.Date(2013, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "Tỵ",
		},
		{
			args: args{NewGregorian(time.Date(2014, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "Ngọ",
		},
		{
			args: args{NewGregorian(time.Date(2015, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "Mùi",
		},
		{
			args: args{NewGregorian(time.Date(2016, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "Thân",
		},
		{
			args: args{NewGregorian(time.Date(2017, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "Dậu",
		},
		{
			args: args{NewGregorian(time.Date(2018, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "Tuất",
		},
		{
			args: args{NewGregorian(time.Date(2019, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "Hợi",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 23, 0, 0, 0, 0, time.Local))},
			want: "Tý",
		},
		{
			args: args{NewGregorian(time.Date(2020, 6, 21, 0, 0, 0, 0, time.Local))},
			want: "Tý",
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "Tý",
		},
		{
			args: args{NewGregorian(time.Date(2021, 8, 5, 0, 0, 0, 0, time.Local))},
			want: "Sửu",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().Animal(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_SolarTerm(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want SolarTerm
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 240, Name: "Tiểu tuyết"},
		},
		{
			args: args{NewGregorian(time.Date(2024, 3, 20, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 0, Name: "Xuân phân"},
		},
		{
			args: args{NewGregorian(time.Date(2024, 4, 16, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 15, Name: "Thanh minh"},
		},
		{
			args: args{NewGregorian(time.Date(2024, 5, 2, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 30, Name: "Cốc vũ"},
		},
		{
			args: args{NewGregorian(time.Date(2024, 5, 16, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 45, Name: "Lập hạ"},
		},
		{
			args: args{NewGregorian(time.Date(2024, 5, 29, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 60, Name: "Tiểu mãn"},
		},
		{
			args: args{NewGregorian(time.Date(2024, 6, 11, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 75, Name: "Mang chủng"},
		},
		{
			args: args{NewGregorian(time.Date(2024, 6, 27, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 90, Name: "Hạ chí"},
		},
		{
			args: args{NewGregorian(time.Date(2024, 7, 10, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 105, Name: "Tiểu thử"},
		},
		{
			args: args{NewGregorian(time.Date(2024, 7, 25, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 120, Name: "Đại thử"},
		},
		{
			args: args{NewGregorian(time.Date(2024, 8, 8, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 135, Name: "Lập thu"},
		},
		{
			args: args{NewGregorian(time.Date(2024, 8, 29, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 150, Name: "Xử thử"},
		},
		{
			args: args{NewGregorian(time.Date(2024, 9, 10, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 165, Name: "Bạch lộ"},
		},
		{
			args: args{NewGregorian(time.Date(2024, 9, 26, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 180, Name: "Thu phân"},
		},
		{
			args: args{NewGregorian(time.Date(2024, 10, 9, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 195, Name: "Hàn lộ"},
		},
		{
			args: args{NewGregorian(time.Date(2024, 10, 31, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 210, Name: "Sương giáng"},
		},
		{
			args: args{NewGregorian(time.Date(2024, 11, 13, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 225, Name: "Lập đông"},
		},
		{
			args: args{NewGregorian(time.Date(2024, 11, 22, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 240, Name: "Tiểu tuyết"},
		},
		{
			args: args{NewGregorian(time.Date(2024, 11, 22, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 240, Name: "Tiểu tuyết"},
		},
		{
			args: args{NewGregorian(time.Date(2024, 12, 6, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 255, Name: "Đại tuyết"},
		},
		{
			args: args{NewGregorian(time.Date(2024, 12, 26, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 270, Name: "Đông chí"},
		},
		{
			args: args{NewGregorian(time.Date(2025, 1, 8, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 285, Name: "Tiểu hàn"},
		},
		{
			args: args{NewGregorian(time.Date(2025, 1, 23, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 300, Name: "Đại hàn"},
		},
		{
			args: args{NewGregorian(time.Date(2025, 2, 3, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 315, Name: "Lập xuân"},
		},
		{
			args: args{NewGregorian(time.Date(2025, 2, 19, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 330, Name: "Vũ Thủy"},
		},
		{
			args: args{NewGregorian(time.Date(2025, 3, 10, 0, 0, 0, 0, time.Local))},
			want: SolarTerm{Longitude: 345, Name: "Kinh trập"},
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().SolarTerm(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_LuckyHour(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "010010110011",
		},
		{
			args: args{NewGregorian(time.Date(2025, 3, 2, 0, 0, 0, 0, time.Local))},
			want: "110100101100",
		},
		{
			args: args{NewGregorian(time.Date(2025, 3, 3, 0, 0, 0, 0, time.Local))},
			want: "001101001011",
		},
		{
			args: args{NewGregorian(time.Date(2025, 3, 4, 0, 0, 0, 0, time.Local))},
			want: "110011010010",
		},
		{
			args: args{NewGregorian(time.Date(2025, 3, 5, 0, 0, 0, 0, time.Local))},
			want: "101100110100",
		},
		{
			args: args{NewGregorian(time.Date(2025, 3, 6, 0, 0, 0, 0, time.Local))},
			want: "001011001101",
		},
		{
			args: args{NewGregorian(time.Date(2025, 3, 7, 0, 0, 0, 0, time.Local))},
			want: "010010110011",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().LuckyHour(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_LuckyHours(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want []LuckyHour
	}{
		{
			args: args{NewGregorian(time.Date(2025, 3, 2, 0, 0, 0, 0, time.Local))},
			want: []LuckyHour{{Chi: "Tý", From: 23, To: 1}, {Chi: "Sửu", From: 1, To: 3}, {Chi: "Mão", From: 5, To: 7}, {Chi: "Ngọ", From: 11, To: 13}, {Chi: "Thân", From: 15, To: 17}, {Chi: "Dậu", From: 17, To: 19}},
		},
		{
			args: args{NewGregorian(time.Date(2025, 3, 3, 0, 0, 0, 0, time.Local))},
			want: []LuckyHour{{Chi: "Dần", From: 3, To: 5}, {Chi: "Mão", From: 5, To: 7}, {Chi: "Tỵ", From: 9, To: 11}, {Chi: "Thân", From: 15, To: 17}, {Chi: "Tuất", From: 19, To: 21}, {Chi: "Hợi", From: 21, To: 23}},
		},
		{
			args: args{NewGregorian(time.Date(2025, 3, 4, 0, 0, 0, 0, time.Local))},
			want: []LuckyHour{{Chi: "Tý", From: 23, To: 1}, {Chi: "Sửu", From: 1, To: 3}, {Chi: "Thìn", From: 7, To: 9}, {Chi: "Tỵ", From: 9, To: 11}, {Chi: "Mùi", From: 13, To: 15}, {Chi: "Tuất", From: 19, To: 21}},
		},
		{
			args: args{NewGregorian(time.Date(2025, 3, 5, 0, 0, 0, 0, time.Local))},
			want: []LuckyHour{{Chi: "Tý", From: 23, To: 1}, {Chi: "Dần", From: 3, To: 5}, {Chi: "Mão", From: 5, To: 7}, {Chi: "Ngọ", From: 11, To: 13}, {Chi: "Mùi", From: 13, To: 15}, {Chi: "Dậu", From: 17, To: 19}},
		},
		{
			args: args{NewGregorian(time.Date(2025, 3, 6, 0, 0, 0, 0, time.Local))},
			want: []LuckyHour{{Chi: "Dần", From: 3, To: 5}, {Chi: "Thìn", From: 7, To: 9}, {Chi: "Tỵ", From: 9, To: 11}, {Chi: "Thân", From: 15, To: 17}, {Chi: "Dậu", From: 17, To: 19}, {Chi: "Hợi", From: 21, To: 23}},
		},
		{
			args: args{NewGregorian(time.Date(2025, 3, 7, 0, 0, 0, 0, time.Local))},
			want: []LuckyHour{{Chi: "Sửu", From: 1, To: 3}, {Chi: "Thìn", From: 7, To: 9}, {Chi: "Ngọ", From: 11, To: 13}, {Chi: "Mùi", From: 13, To: 15}, {Chi: "Tuất", From: 19, To: 21}, {Chi: "Hợi", From: 21, To: 23}},
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: []LuckyHour{{Chi: "Dần", From: 3, To: 5}, {Chi: "Thìn", From: 7, To: 9}, {Chi: "Tỵ", From: 9, To: 11}, {Chi: "Thân", From: 15, To: 17}, {Chi: "Dậu", From: 17, To: 19}, {Chi: "Hợi", From: 21, To: 23}},
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().LuckyHours(), "args{%v}", tt.args.g)
		})
	}
}
func TestLunar_Festival(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want []Festival
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: []Festival{},
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: []Festival{},
		},
		{
			args: args{NewGregorian(time.Date(2021, 2, 12, 0, 0, 0, 0, time.Local))},
			want: []Festival{{Day: 1, Month: 1, Name: "Tết Nguyên Đán"}},
		},
		{
			args: args{NewGregorian(time.Date(2021, 2, 26, 0, 0, 0, 0, time.Local))},
			want: []Festival{{Day: 15, Month: 1, Name: "Rằm tháng Giêng"}},
		},
		{
			args: args{NewGregorian(time.Date(2025, 4, 7, 0, 0, 0, 0, time.Local))},
			want: []Festival{{Day: 10, Month: 3, Name: "Giỗ Tổ Hùng Vương"}},
		},
		{
			args: args{NewGregorian(time.Date(2025, 5, 12, 0, 0, 0, 0, time.Local))},
			want: []Festival{{Day: 15, Month: 4, Name: "Phật Đản"}},
		},
		{
			args: args{NewGregorian(time.Date(2021, 6, 14, 0, 0, 0, 0, time.Local))},
			want: []Festival{{Day: 5, Month: 5, Name: "Lễ Đoan Ngọ"}},
		},
		{
			args: args{NewGregorian(time.Date(2021, 8, 14, 0, 0, 0, 0, time.Local))},
			want: []Festival{},
		},
		{
			args: args{NewGregorian(time.Date(2021, 8, 22, 0, 0, 0, 0, time.Local))},
			want: []Festival{{Day: 15, Month: 7, Name: "Vu Lan"}},
		},
		{
			args: args{NewGregorian(time.Date(2021, 9, 21, 0, 0, 0, 0, time.Local))},
			want: []Festival{{Day: 15, Month: 8, Name: "Tết Trung Thu"}},
		},
		{
			args: args{NewGregorian(time.Date(2021, 10, 14, 0, 0, 0, 0, time.Local))},
			want: []Festival{},
		},
		{
			args: args{NewGregorian(time.Date(2021, 11, 5, 0, 0, 0, 0, time.Local))},
			want: []Festival{},
		},
		{
			args: args{NewGregorian(time.Date(2021, 11, 19, 0, 0, 0, 0, time.Local))},
			want: []Festival{},
		},
		{
			args: args{NewGregorian(time.Date(2022, 1, 10, 0, 0, 0, 0, time.Local))},
			want: []Festival{},
		},
		{
			args: args{NewGregorian(time.Date(2026, 2, 10, 0, 0, 0, 0, time.Local))},
			want: []Festival{{Day: 23, Month: 12, Name: "Ông Táo chầu trời"}},
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().Festivals(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_Year(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: -1,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: 2020,
		},
		{
			args: args{NewGregorian(time.Date(2021, 5, 1, 0, 0, 0, 0, time.Local))},
			want: 2021,
		},
		{
			args: args{NewGregorian(time.Date(2025, 1, 28, 0, 0, 0, 0, time.Local))},
			want: 2024,
		},
		{
			args: args{NewGregorian(time.Date(2025, 1, 29, 0, 0, 0, 0, time.Local))},
			want: 2025,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().Year(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_Month(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: 11,
		},
		{
			args: args{NewGregorian(time.Date(2021, 3, 5, 0, 0, 0, 0, time.Local))},
			want: 1,
		},
		{
			args: args{NewGregorian(time.Date(2021, 4, 5, 0, 0, 0, 0, time.Local))},
			want: 2,
		},
		{
			args: args{NewGregorian(time.Date(2021, 5, 5, 0, 0, 0, 0, time.Local))},
			want: 3,
		},
		{
			args: args{NewGregorian(time.Date(2021, 6, 5, 0, 0, 0, 0, time.Local))},
			want: 4,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 5, 0, 0, 0, 0, time.Local))},
			want: 5,
		},
		{
			args: args{NewGregorian(time.Date(2021, 8, 5, 0, 0, 0, 0, time.Local))},
			want: 6,
		},
		{
			args: args{NewGregorian(time.Date(2021, 9, 5, 0, 0, 0, 0, time.Local))},
			want: 7,
		},
		{
			args: args{NewGregorian(time.Date(2021, 10, 5, 0, 0, 0, 0, time.Local))},
			want: 8,
		},
		{
			args: args{NewGregorian(time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local))},
			want: 9,
		},
		{
			args: args{NewGregorian(time.Date(2021, 11, 5, 0, 0, 0, 0, time.Local))},
			want: 10,
		},
		{
			args: args{NewGregorian(time.Date(2022, 12, 5, 0, 0, 0, 0, time.Local))},
			want: 11,
		},
		{
			args: args{NewGregorian(time.Date(2022, 1, 5, 0, 0, 0, 0, time.Local))},
			want: 12,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().Month(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_LeapMonth(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: 0,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: 4,
		},
		{
			args: args{NewGregorian(time.Date(2023, 7, 1, 0, 0, 0, 0, time.Local))},
			want: 2,
		},
		{
			args: args{NewGregorian(time.Date(2025, 7, 1, 0, 0, 0, 0, time.Local))},
			want: 6,
		},
		{
			args: args{NewGregorian(time.Date(2028, 7, 1, 0, 0, 0, 0, time.Local))},
			want: 5,
		},
		{
			args: args{NewGregorian(time.Date(2031, 7, 1, 0, 0, 0, 0, time.Local))},
			want: 3,
		},
		{
			args: args{NewGregorian(time.Date(2033, 7, 1, 0, 0, 0, 0, time.Local))},
			want: 0,
		},
		{
			args: args{NewGregorian(time.Date(2036, 7, 1, 0, 0, 0, 0, time.Local))},
			want: 6,
		},
		{
			args: args{NewGregorian(time.Date(2039, 7, 1, 0, 0, 0, 0, time.Local))},
			want: 5,
		},
		{
			args: args{NewGregorian(time.Date(2042, 7, 1, 0, 0, 0, 0, time.Local))},
			want: 2,
		},
		{
			args: args{NewGregorian(time.Date(2044, 7, 1, 0, 0, 0, 0, time.Local))},
			want: 7,
		},
		{
			args: args{NewGregorian(time.Date(2047, 7, 1, 0, 0, 0, 0, time.Local))},
			want: 5,
		},
		{
			args: args{NewGregorian(time.Date(2050, 7, 1, 0, 0, 0, 0, time.Local))},
			want: 3,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().LeapMonth(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_Day(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: 4,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 19, 0, 0, 0, 0, time.Local))},
			want: 1,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 20, 0, 0, 0, 0, time.Local))},
			want: 2,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 21, 0, 0, 0, 0, time.Local))},
			want: 3,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 22, 0, 0, 0, 0, time.Local))},
			want: 4,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().Day(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_ToYearString(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "Kỷ Mùi",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "Canh Tý",
		},
		{
			args: args{NewGregorian(time.Date(2021, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "Tân Sửu",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().ToYearString(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_ToMonthString(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "Bính Tý",
		},
		{
			args: args{NewGregorian(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "Đinh Sửu",
		},
		{
			args: args{NewGregorian(time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local))},
			want: "Mậu Dần",
		},
		{
			args: args{NewGregorian(time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local))},
			want: "Kỷ Mão",
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local))},
			want: "Canh Thìn",
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: "Tân Tỵ",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "Tân Tỵ",
		},
		{
			args: args{NewGregorian(time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local))},
			want: "Tân Tỵ",
		},
		{
			args: args{NewGregorian(time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local))},
			want: "Nhâm Ngọ",
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local))},
			want: "Quý Mùi",
		},
		{
			args: args{NewGregorian(time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local))},
			want: "Giáp Thân",
		},
		{
			args: args{NewGregorian(time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local))},
			want: "Ất Dậu",
		},
		{
			args: args{NewGregorian(time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local))},
			want: "Bính Tuất",
		},
		{
			args: args{NewGregorian(time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local))},
			want: "Đinh Hợi",
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "Mậu Tý",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().ToMonthString(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_ToDayString(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "Kỷ Hợi",
		},
		{
			args: args{NewGregorian(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "Quý Mão",
		},
		{
			args: args{NewGregorian(time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local))},
			want: "Giáp Tuất",
		},
		{
			args: args{NewGregorian(time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local))},
			want: "Quý Mão",
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local))},
			want: "Giáp Tuất",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "Giáp Thìn",
		},
		{
			args: args{NewGregorian(time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local))},
			want: "Ất Hợi",
		},
		{
			args: args{NewGregorian(time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local))},
			want: "Ất Tỵ",
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local))},
			want: "Bính Tý",
		},
		{
			args: args{NewGregorian(time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local))},
			want: "Đinh Mùi",
		},
		{
			args: args{NewGregorian(time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local))},
			want: "Đinh Sửu",
		},
		{
			args: args{NewGregorian(time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local))},
			want: "Mậu Thân",
		},
		{
			args: args{NewGregorian(time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local))},
			want: "Mậu Dần",
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 3, 0, 0, 0, 0, time.Local))},
			want: "Tân Hợi",
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 5, 0, 0, 0, 0, time.Local))},
			want: "Quý Sửu",
		},
		{
			args: args{NewGregorian(time.Date(2021, 4, 11, 0, 0, 0, 0, time.Local))},
			want: "Kỷ Sửu",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().ToDayString(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_DateTime(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "-1-11-04 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-04-09 00:00:00",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			y, m, d, h, i, s := (tt.args.g).ToLunar().DateTime()
			assert.Equalf(t, tt.want, fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", y, m, d, h, i, s), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_Date(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "-1-11-04",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-04-09",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			y, m, d := (tt.args.g).ToLunar().Date()
			assert.Equalf(t, tt.want, fmt.Sprintf("%d-%02d-%02d", y, m, d), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_Time(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 1, 1, 1, 0, time.Local))},
			want: "01:01:01",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			h, i, s := (tt.args.g).ToLunar().Time()
			assert.Equalf(t, tt.want, fmt.Sprintf("%02d:%02d:%02d", h, i, s), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_String(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "-1-11-04 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "2019-12-07 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-01-08 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-02-08 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-03-09 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-04-09 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-04-10 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-05-11 00:00:00",
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local))},
			want: "2020-06-12 00:00:00",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().String(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_ToDateString(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: "Ngày 04 tháng 11 năm -1",
		},
		{
			args: args{NewGregorian(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local))},
			want: "Ngày 07 tháng 12 năm 2019",
		},
		{
			args: args{NewGregorian(time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local))},
			want: "Ngày 08 tháng 01 năm 2020",
		},
		{
			args: args{NewGregorian(time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local))},
			want: "Ngày 08 tháng 02 năm 2020",
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local))},
			want: "Ngày 09 tháng 03 năm 2020",
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: "Ngày 09 tháng 04 năm 2020",
		},
		{
			args: args{NewGregorian(time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local))},
			want: "Ngày 10 tháng 04 năm 2020",
		},
		{
			args: args{NewGregorian(time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local))},
			want: "Ngày 11 tháng 05 năm 2020",
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local))},
			want: "Ngày 12 tháng 06 năm 2020",
		},
		{
			args: args{NewGregorian(time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local))},
			want: "Ngày 14 tháng 07 năm 2020",
		},
		{
			args: args{NewGregorian(time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local))},
			want: "Ngày 15 tháng 08 năm 2020",
		},
		{
			args: args{NewGregorian(time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local))},
			want: "Ngày 16 tháng 09 năm 2020",
		},
		{
			args: args{NewGregorian(time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local))},
			want: "Ngày 17 tháng 10 năm 2020",
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().ToDateString(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsLeapYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2022, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsLeapYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsLeapMonth(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2022, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsLeapMonth(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsRatYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2022, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsRatYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsOxYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: true,
		},
		{
			args: args{NewGregorian(time.Date(2022, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsOxYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsTigerYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2022, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsTigerYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsCatYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2023, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsCatYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsDragonYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2024, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsDragonYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsSnakeYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2025, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsSnakeYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsHorseYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2026, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsHorseYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsGoatYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2027, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsGoatYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsMonkeyYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2028, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsMonkeyYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsRoosterYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2029, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsRoosterYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsDogYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2030, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsDogYear(), "args{%v}", tt.args.g)
		})
	}
}

func TestLunar_IsPigYear(t *testing.T) {
	type args struct {
		g Gregorian
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{NewGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 4, 23, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2020, 8, 5, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2021, 7, 7, 0, 0, 0, 0, time.Local))},
			want: false,
		},
		{
			args: args{NewGregorian(time.Date(2031, 8, 5, 0, 0, 0, 0, time.Local))},
			want: true,
		},
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToLunar().IsPigYear(), "args{%v}", tt.args.g)
		})
	}
}
