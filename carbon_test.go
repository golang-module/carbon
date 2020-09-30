package carbon

import (
	"fmt"
	"testing"
	"time"
)

func TestCarbon_Timezone(t *testing.T) {
	name := "PRC"
	r := New().Timezone(name).loc.String()

	loc, _ := time.LoadLocation(name)
	e := loc.String()

	if e != r {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_Yesterday(t *testing.T) {
	e := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	r := New().Now().Yesterday()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_Today(t *testing.T) {
	e := time.Now().Format("2006-01-02")
	r := New().Now().Today()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_Tomorrow(t *testing.T) {
	e := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	r := New().Now().Tomorrow()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_FirstOfMonth(t *testing.T) {
	now := time.Now()
	e := now.AddDate(0, 0, -now.Day()+1).Format("2006-01-02 00:00:00")
	r := New().Now().FirstOfMonth()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_LastOfMonth(t *testing.T) {
	now := time.Now()
	e := now.AddDate(0, 0, -now.Day()+1).AddDate(0, 1, -1).Format("2006-01-02 00:00:00")
	r := New().Now().LastOfMonth()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_CreateFromTimestamp(t *testing.T) {
	now := time.Now()
	e := now.Format("2006-01-02 15:04:05")
	r := New().CreateFromTimestamp(now.Unix()).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_CreateFromDateTime(t *testing.T) {
	now := time.Now()
	e := now.Format("2006-01-02 15:04:05")
	r := New().CreateFromDateTime(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_CreateFromDate(t *testing.T) {
	now := time.Now()
	e := now.Format("2006-01-02")
	r := New().CreateFromDate(now.Year(), now.Month(), now.Day()).ToDateString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_CreateFromTime(t *testing.T) {
	now := time.Now()
	e := now.Format("2006-01-02 15:04:05")
	r := New().CreateFromTime(now.Hour(), now.Minute(), now.Second()).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_CreateFromString(t *testing.T) {
	now := time.Now().Add(time.Hour * 2)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().CreateFromString("2h").ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddYears(t *testing.T) {
	now := time.Now().AddDate(2, 0, 0)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().AddYears(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubYears(t *testing.T) {
	now := time.Now().AddDate(-2, 0, 0)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().SubYears(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddYear(t *testing.T) {
	now := time.Now().AddDate(1, 0, 0)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().AddYear().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubYear(t *testing.T) {
	now := time.Now().AddDate(-1, 0, 0)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().SubYear().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddMonths(t *testing.T) {
	now := time.Now().AddDate(0, 2, 0)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().AddMonths(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubMonths(t *testing.T) {
	now := time.Now().AddDate(0, -2, 0)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().SubMonths(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddMonth(t *testing.T) {
	now := time.Now().AddDate(0, 1, 0)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().AddMonth().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubMonth(t *testing.T) {
	now := time.Now().AddDate(0, -1, 0)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().SubMonth().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddWeeks(t *testing.T) {
	now := time.Now().AddDate(0, 0, 14)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().AddWeeks(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubWeeks(t *testing.T) {
	now := time.Now().AddDate(0, 0, -14)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().SubWeeks(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddWeek(t *testing.T) {
	now := time.Now().AddDate(0, 0, 7)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().AddWeek().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubWeek(t *testing.T) {
	now := time.Now().AddDate(0, 0, -7)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().SubWeek().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddDays(t *testing.T) {
	now := time.Now().AddDate(0, 0, 2)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().AddDays(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubDays(t *testing.T) {
	now := time.Now().AddDate(0, 0, -2)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().SubDays(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddDay(t *testing.T) {
	now := time.Now().AddDate(0, 0, 1)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().AddDay().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubDay(t *testing.T) {
	now := time.Now().AddDate(0, 0, -1)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().SubDay().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddHours(t *testing.T) {
	now := time.Now().Add(time.Hour * 2)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().AddHours(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubHours(t *testing.T) {
	now := time.Now().Add(time.Hour * -2)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().SubHours(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddHour(t *testing.T) {
	now := time.Now().Add(time.Hour)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().AddHour().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubHour(t *testing.T) {
	now := time.Now().Add(-time.Hour)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().SubHour().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddMinutes(t *testing.T) {
	now := time.Now().Add(time.Minute * 2)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().AddMinutes(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubMinutes(t *testing.T) {
	now := time.Now().Add(time.Minute * -2)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().SubMinutes(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddMinute(t *testing.T) {
	now := time.Now().Add(time.Minute)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().AddMinute().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubMinute(t *testing.T) {
	now := time.Now().Add(-time.Minute)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().SubMinute().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddSeconds(t *testing.T) {
	now := time.Now().Add(time.Second * 2)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().AddSeconds(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubSeconds(t *testing.T) {
	now := time.Now().Add(time.Second * -2)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().SubSeconds(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddSecond(t *testing.T) {
	now := time.Now().Add(time.Second)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().AddSecond().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubSecond(t *testing.T) {
	now := time.Now().Add(-time.Second)
	e := now.Format("2006-01-02 15:04:05")
	r := New().Now().SubSecond().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_IsLeapYear(t *testing.T) {
	year := time.Now().Year()
	e := false
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		e = true
	}

	r := New().Now().IsLeapYear()
	if r != e {
		if e == true {
			t.Fatalf("Expected %d is leap year, but got %d isn't leap year", year, year)
		} else {
			t.Fatalf("Expected %d isn't leap year, but got %d is leap year", year, year)
		}
	}
}

func TestCarbon_IsMonday(t *testing.T) {
	now := time.Now()
	wd := now.Weekday().String()
	today := now.Format("2006-01-02")

	e := false
	if wd == "Monday" {
		e = true
	}

	r := New().Now().IsMonday()

	if r != e {
		if e == true {
			t.Fatalf("Expected %s is %s, but got %s isn't %s", today, wd, today, wd)
		} else {
			t.Fatalf("Expected %s isn't %s, but got %s is %s", today, wd, today, wd)
		}
	}
}

func TestCarbon_IsTuesday(t *testing.T) {
	now := time.Now()
	wd := now.Weekday().String()
	today := now.Format("2006-01-02")

	e := false
	if wd == "Tuesday" {
		e = true
	}

	r := New().Now().IsTuesday()

	if r != e {
		if e == true {
			t.Fatalf("Expected %s is %s, but got %s isn't %s", today, wd, today, wd)
		} else {
			t.Fatalf("Expected %s isn't %s, but got %s is %s", today, wd, today, wd)
		}
	}
}

func TestCarbon_IsWednesday(t *testing.T) {
	now := time.Now()
	wd := now.Weekday().String()
	today := now.Format("2006-01-02")

	e := false
	if wd == "Wednesday" {
		e = true
	}

	r := New().Now().IsWednesday()

	if r != e {
		if e == true {
			t.Fatalf("Expected %s is %s, but got %s isn't %s", today, wd, today, wd)
		} else {
			t.Fatalf("Expected %s isn't %s, but got %s is %s", today, wd, today, wd)
		}
	}
}

func TestCarbon_IsThursday(t *testing.T) {
	now := time.Now()
	wd := now.Weekday().String()
	today := now.Format("2006-01-02")

	e := false
	if wd == "Thursday" {
		e = true
	}

	r := New().Now().IsThursday()

	if r != e {
		if e == true {
			t.Fatalf("Expected %s is %s, but got %s isn't %s", today, wd, today, wd)
		} else {
			t.Fatalf("Expected %s isn't %s, but got %s is %s", today, wd, today, wd)
		}
	}
}

func TestCarbon_IsFriday(t *testing.T) {
	now := time.Now()
	wd := now.Weekday().String()
	today := now.Format("2006-01-02")

	e := false
	if wd == "Friday" {
		e = true
	}

	r := New().Now().IsFriday()

	if r != e {
		if e == true {
			t.Fatalf("Expected %s is %s, but got %s isn't %s", today, wd, today, wd)
		} else {
			t.Fatalf("Expected %s isn't %s, but got %s is %s", today, wd, today, wd)
		}
	}
}

func TestCarbon_IsSaturday(t *testing.T) {
	now := time.Now()
	wd := now.Weekday().String()
	today := now.Format("2006-01-02")

	e := false
	if wd == "Saturday" {
		e = true
	}

	r := New().Now().IsSaturday()

	if r != e {
		if e == true {
			t.Fatalf("Expected %s is %s, but got %s isn't %s", today, wd, today, wd)
		} else {
			t.Fatalf("Expected %s isn't %s, but got %s is %s", today, wd, today, wd)
		}
	}
}

func TestCarbon_IsSunday(t *testing.T) {
	now := time.Now()
	wd := now.Weekday().String()
	today := now.Format("2006-01-02")

	e := false
	if wd == "Sunday" {
		e = true
	}

	r := New().Now().IsSunday()

	if r != e {
		if e == true {
			t.Fatalf("Expected %s is %s, but got %s isn't %s", today, wd, today, wd)
		} else {
			t.Fatalf("Expected %s isn't %s, but got %s is %s", today, wd, today, wd)
		}
	}
}

