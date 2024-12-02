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
		l string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Time{}), ""},
			want: "",
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Time{}), "en"},
			want: "",
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Time{}), "fa"},
			want: "",
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)), ""},
			want: "",
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)), "en"},
			want: "",
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)), "fa"},
			want: "",
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Dey",
		},
		{
			name: "case8",
			args: args{FromGregorian(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "دی",
		},
		{
			name: "case9",
			args: args{FromGregorian(time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Bahman",
		},
		{
			name: "case10",
			args: args{FromGregorian(time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "بهمن",
		},
		{
			name: "case11",
			args: args{FromGregorian(time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Esfand",
		},
		{
			name: "case12",
			args: args{FromGregorian(time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "اسفند",
		},
		{
			name: "case13",
			args: args{FromGregorian(time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Farvardin",
		},
		{
			name: "case14",
			args: args{FromGregorian(time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "فروردین",
		},
		{
			name: "case15",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Ordibehesht",
		},
		{
			name: "case16",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "اردیبهشت",
		},
		{
			name: "case17",
			args: args{FromGregorian(time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Khordad",
		},
		{
			name: "case18",
			args: args{FromGregorian(time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "خرداد",
		},
		{
			name: "case19",
			args: args{FromGregorian(time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Tir",
		},
		{
			name: "case20",
			args: args{FromGregorian(time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "تیر",
		},
		{
			name: "case21",
			args: args{FromGregorian(time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Mordad",
		},
		{
			name: "case22",
			args: args{FromGregorian(time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "مرداد",
		},
		{
			name: "case23",
			args: args{FromGregorian(time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Shahrivar",
		},
		{
			name: "case24",
			args: args{FromGregorian(time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "شهریور",
		},
		{
			name: "case25",
			args: args{FromGregorian(time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Mehr",
		},
		{
			name: "case26",
			args: args{FromGregorian(time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "مهر",
		},
		{
			name: "case27",
			args: args{FromGregorian(time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Aban",
		},
		{
			name: "case28",
			args: args{FromGregorian(time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "آبان",
		},
		{
			name: "case29",
			args: args{FromGregorian(time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Azar",
		},
		{
			name: "case30",
			args: args{FromGregorian(time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "آذر",
		},
		{
			name: "case31",
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Dey",
		},
		{
			name: "case32",
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "دی",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToPersian().ToMonthString(tt.args.l), "args{%v}", tt.args.g)
		})
	}
}

func TestPersian_ToShortMonthString(t *testing.T) {
	type args struct {
		g Gregorian
		l string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{FromGregorian(time.Time{}), ""},
			want: "",
		},
		{
			name: "case2",
			args: args{FromGregorian(time.Time{}), "en"},
			want: "",
		},
		{
			name: "case3",
			args: args{FromGregorian(time.Time{}), "fa"},
			want: "",
		},
		{
			name: "case4",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)), ""},
			want: "",
		},
		{
			name: "case5",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)), "en"},
			want: "",
		},
		{
			name: "case6",
			args: args{FromGregorian(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)), "fa"},
			want: "",
		},
		{
			name: "case7",
			args: args{FromGregorian(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Dey",
		},
		{
			name: "case8",
			args: args{FromGregorian(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "دی",
		},
		{
			name: "case9",
			args: args{FromGregorian(time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Bah",
		},
		{
			name: "case10",
			args: args{FromGregorian(time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "بهم",
		},
		{
			name: "case11",
			args: args{FromGregorian(time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Esf",
		},
		{
			name: "case12",
			args: args{FromGregorian(time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "اسف",
		},
		{
			name: "case13",
			args: args{FromGregorian(time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Far",
		},
		{
			name: "case14",
			args: args{FromGregorian(time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "فرو",
		},
		{
			name: "case15",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Ord",
		},
		{
			name: "case16",
			args: args{FromGregorian(time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "ارد",
		},
		{
			name: "case17",
			args: args{FromGregorian(time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Kho",
		},
		{
			name: "case18",
			args: args{FromGregorian(time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "خرد",
		},
		{
			name: "case19",
			args: args{FromGregorian(time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Tir",
		},
		{
			name: "case20",
			args: args{FromGregorian(time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "تیر",
		},
		{
			name: "case21",
			args: args{FromGregorian(time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Mor",
		},
		{
			name: "case22",
			args: args{FromGregorian(time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "مرد",
		},
		{
			name: "case23",
			args: args{FromGregorian(time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Sha",
		},
		{
			name: "case24",
			args: args{FromGregorian(time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "شهر",
		},
		{
			name: "case25",
			args: args{FromGregorian(time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Meh",
		},
		{
			name: "case26",
			args: args{FromGregorian(time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "مهر",
		},
		{
			name: "case27",
			args: args{FromGregorian(time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Aba",
		},
		{
			name: "case28",
			args: args{FromGregorian(time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "آبا",
		},
		{
			name: "case29",
			args: args{FromGregorian(time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Aza",
		},
		{
			name: "case30",
			args: args{FromGregorian(time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "آذر",
		},
		{
			name: "case31",
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)), "en"},
			want: "Dey",
		},
		{
			name: "case32",
			args: args{FromGregorian(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)), "fa"},
			want: "دی",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.g).ToPersian().ToShortMonthString(tt.args.l), "args{%v}", tt.args.g)
		})
	}
}
func TestPersian_ToWeekString(t *testing.T) {
	type args struct {
		p Persian
		l string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{FromPersian(0, 0, 0, 0, 0, 0), ""},
			want: "",
		},
		{
			name: "case2",
			args: args{FromPersian(0, 0, 0, 0, 0, 0), "en"},
			want: "",
		},
		{
			name: "case3",
			args: args{FromPersian(0, 0, 0, 0, 0, 0), "fa"},
			want: "",
		},
		{
			name: "case4",
			args: args{FromPersian(2020, 1, 1, 0, 0, 0), "en"},
			want: "Yekshanbeh",
		},
		{
			name: "case5",
			args: args{FromPersian(2020, 1, 1, 0, 0, 0), "fa"},
			want: "نجشنبه",
		},
		{
			name: "case6",
			args: args{FromPersian(2020, 1, 2, 0, 0, 0), "en"},
			want: "Doshanbeh",
		},
		{
			name: "case7",
			args: args{FromPersian(2020, 1, 2, 0, 0, 0), "fa"},
			want: "دوشنبه",
		},
		{
			name: "case8",
			args: args{FromPersian(2020, 1, 3, 0, 0, 0), "en"},
			want: "Seshanbeh",
		},
		{
			name: "case9",
			args: args{FromPersian(2020, 1, 3, 0, 0, 0), "fa"},
			want: "سه شنبه",
		},
		{
			name: "case10",
			args: args{FromPersian(2020, 1, 4, 0, 0, 0), "en"},
			want: "Chaharshanbeh",
		},
		{
			name: "case11",
			args: args{FromPersian(2020, 1, 4, 0, 0, 0), "fa"},
			want: "چهارشنبه",
		},
		{
			name: "case12",
			args: args{FromPersian(2020, 1, 5, 0, 0, 0), "en"},
			want: "Panjshanbeh",
		},
		{
			name: "case13",
			args: args{FromPersian(2020, 1, 5, 0, 0, 0), "fa"},
			want: "پنجشنبه",
		},
		{
			name: "case14",
			args: args{FromPersian(2020, 1, 6, 0, 0, 0), "en"},
			want: "Jomeh",
		},
		{
			name: "case15",
			args: args{FromPersian(2020, 1, 6, 0, 0, 0), "fa"},
			want: "جمعه",
		},
		{
			name: "case16",
			args: args{FromPersian(2020, 1, 7, 0, 0, 0), "en"},
			want: "Shanbeh",
		},
		{
			name: "case17",
			args: args{FromPersian(2020, 1, 7, 0, 0, 0), "fa"},
			want: "شنبه",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.p).ToWeekString(tt.args.l), "args{%v}", tt.args.p)
		})
	}
}

func TestPersian_ToShortWeekString(t *testing.T) {
	type args struct {
		p Persian
		l string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{FromPersian(0, 0, 0, 0, 0, 0), ""},
			want: "",
		},
		{
			name: "case2",
			args: args{FromPersian(0, 0, 0, 0, 0, 0), "en"},
			want: "",
		},
		{
			name: "case3",
			args: args{FromPersian(0, 0, 0, 0, 0, 0), "fa"},
			want: "",
		},
		{
			name: "case4",
			args: args{FromPersian(2020, 1, 1, 0, 0, 0), "en"},
			want: "Yek",
		},
		{
			name: "case5",
			args: args{FromPersian(2020, 1, 1, 0, 0, 0), "fa"},
			want: "پ",
		},
		{
			name: "case6",
			args: args{FromPersian(2020, 1, 2, 0, 0, 0), "en"},
			want: "Dos",
		},
		{
			name: "case7",
			args: args{FromPersian(2020, 1, 2, 0, 0, 0), "fa"},
			want: "چ",
		},
		{
			name: "case8",
			args: args{FromPersian(2020, 1, 3, 0, 0, 0), "en"},
			want: "Ses",
		},
		{
			name: "case9",
			args: args{FromPersian(2020, 1, 3, 0, 0, 0), "fa"},
			want: "س",
		},
		{
			name: "case10",
			args: args{FromPersian(2020, 1, 4, 0, 0, 0), "en"},
			want: "Cha",
		},
		{
			name: "case11",
			args: args{FromPersian(2020, 1, 4, 0, 0, 0), "fa"},
			want: "د",
		},
		{
			name: "case12",
			args: args{FromPersian(2020, 1, 5, 0, 0, 0), "en"},
			want: "Pan",
		},
		{
			name: "case13",
			args: args{FromPersian(2020, 1, 5, 0, 0, 0), "fa"},
			want: "ی",
		},
		{
			name: "case14",
			args: args{FromPersian(2020, 1, 6, 0, 0, 0), "en"},
			want: "Jom",
		},
		{
			name: "case15",
			args: args{FromPersian(2020, 1, 6, 0, 0, 0), "fa"},
			want: "ش",
		},
		{
			name: "case16",
			args: args{FromPersian(2020, 1, 7, 0, 0, 0), "en"},
			want: "Sha",
		},
		{
			name: "case17",
			args: args{FromPersian(2020, 1, 7, 0, 0, 0), "fa"},
			want: "ج",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.p).ToShortWeekString(tt.args.l), "args{%v}", tt.args.p)
		})
	}
}
func TestPersian_IsValid(t *testing.T) {
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
			args: args{FromPersian(620, 1, 1, 0, 0, 0)},
			want: true,
		},
		{
			name: "case4",
			args: args{FromPersian(9378, 1, 1, 0, 0, 0)},
			want: false,
		},
		{
			name: "case5",
			args: args{FromPersian(622, 1, 1, 0, 0, 0)},
			want: true,
		},
		{
			name: "case6",
			args: args{FromPersian(9377, 1, 1, 0, 0, 0)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, (tt.args.p).IsValid(), "args{%v}", tt.args.p)
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

func TestGregorian_Year_Error(t *testing.T) {
	type args struct {
		p Persian
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "case1",
			args: args{FromPersian(0, 0, 0, 0, 0, 0)},
			want: InvalidDateError(),
		},
		{
			name: "case2",
			args: args{FromPersian(621, 10, 10, 0, 0, 0)},
			want: InvalidDateError(),
		},

		{
			name: "case3",
			args: args{FromPersian(9378, 10, 10, 0, 0, 0)},
			want: InvalidDateError(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, InvalidDateError(), "args(%v)", tt.args.p)
		})
	}
}

func TestError_Locale_Error(t *testing.T) {
	p, l := FromPersian(622, 1, 1, 0, 0, 0), "xx"

	assert.Equal(t, "", p.ToMonthString(l))
	assert.Equal(t, "", p.ToShortMonthString(l))
	assert.Equal(t, "", p.ToWeekString(l))
	assert.Equal(t, "", p.ToShortWeekString(l))
}
