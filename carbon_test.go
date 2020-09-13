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

func TestCarbon_StartOfDay(t *testing.T) {
	a := New().Parse("2020-09-13 17:10:00").ToTimeEndString()
	fmt.Println(a)
}
