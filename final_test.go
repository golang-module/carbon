package carbon

import (
	"testing"
	"time"
)

func TestCarbon_IsLeapYear(t *testing.T) {
	year := time.Now().Year()
	e := false
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		e = true
	}

	r := Now().IsLeapYear()
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

	r := Now().IsMonday()

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

	r := Now().IsTuesday()

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

	r := Now().IsWednesday()

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

	r := Now().IsThursday()

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

	r := Now().IsFriday()

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

	r := Now().IsSaturday()

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

	r := Now().IsSunday()

	if r != e {
		if e == true {
			t.Fatalf("Expected %s is %s, but got %s isn't %s", today, wd, today, wd)
		} else {
			t.Fatalf("Expected %s isn't %s, but got %s is %s", today, wd, today, wd)
		}
	}
}
