package carbon

import (
	"testing"
	"time"
)

func TestCarbon_Timezone(t *testing.T) {
	name := "PRC"
	r := Timezone(name).loc.String()

	loc, _ := time.LoadLocation(name)
	e := loc.String()

	if e != r {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_FirstOfMonth(t *testing.T) {
	now := time.Now()
	e := now.AddDate(0, 0, -now.Day()+1).Format("2006-01-02 00:00:00")
	r := Now().FirstOfMonth().ToDateStartString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_LastOfMonth(t *testing.T) {
	now := time.Now()
	e := now.AddDate(0, 0, -now.Day()+1).AddDate(0, 1, -1).Format("2006-01-02 00:00:00")
	r := Now().LastOfMonth().ToDateStartString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_CreateFromTimestamp(t *testing.T) {
	now := time.Now()
	e := now.Format("2006-01-02 15:04:05")
	r := CreateFromTimestamp(now.Unix()).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_CreateFromDateTime(t *testing.T) {
	now := time.Now()
	e := now.Format("2006-01-02 15:04:05")
	r := CreateFromDateTime(now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second()).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_CreateFromDate(t *testing.T) {
	now := time.Now()
	e := now.Format("2006-01-02")
	r := CreateFromDate(now.Year(), int(now.Month()), now.Day()).ToDateString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_CreateFromTime(t *testing.T) {
	now := time.Now()
	e := now.Format("2006-01-02 15:04:05")
	r := CreateFromTime(now.Hour(), now.Minute(), now.Second()).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_CreateFromString(t *testing.T) {
	now := time.Now().Add(time.Hour * 2)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().ParseByDuration("2h").ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddYears(t *testing.T) {
	now := time.Now().AddDate(2, 0, 0)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().AddYears(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubYears(t *testing.T) {
	now := time.Now().AddDate(-2, 0, 0)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().SubYears(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddYear(t *testing.T) {
	now := time.Now().AddDate(1, 0, 0)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().AddYear().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubYear(t *testing.T) {
	now := time.Now().AddDate(-1, 0, 0)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().SubYear().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddMonths(t *testing.T) {
	now := time.Now().AddDate(0, 2, 0)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().AddMonths(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubMonths(t *testing.T) {
	now := time.Now().AddDate(0, -2, 0)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().SubMonths(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddMonth(t *testing.T) {
	now := time.Now().AddDate(0, 1, 0)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().AddMonth().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubMonth(t *testing.T) {
	now := time.Now().AddDate(0, -1, 0)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().SubMonth().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddDays(t *testing.T) {
	now := time.Now().AddDate(0, 0, 2)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().AddDays(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubDays(t *testing.T) {
	now := time.Now().AddDate(0, 0, -2)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().SubDays(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddDay(t *testing.T) {
	now := time.Now().AddDate(0, 0, 1)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().AddDay().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubDay(t *testing.T) {
	now := time.Now().AddDate(0, 0, -1)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().SubDay().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddHours(t *testing.T) {
	now := time.Now().Add(time.Hour * 2)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().AddHours(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubHours(t *testing.T) {
	now := time.Now().Add(time.Hour * -2)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().SubHours(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddHour(t *testing.T) {
	now := time.Now().Add(time.Hour)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().AddHour().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubHour(t *testing.T) {
	now := time.Now().Add(-time.Hour)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().SubHour().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddMinutes(t *testing.T) {
	now := time.Now().Add(time.Minute * 2)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().AddMinutes(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubMinutes(t *testing.T) {
	now := time.Now().Add(time.Minute * -2)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().SubMinutes(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddMinute(t *testing.T) {
	now := time.Now().Add(time.Minute)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().AddMinute().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubMinute(t *testing.T) {
	now := time.Now().Add(-time.Minute)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().SubMinute().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddSeconds(t *testing.T) {
	now := time.Now().Add(time.Second * 2)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().AddSeconds(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubSeconds(t *testing.T) {
	now := time.Now().Add(time.Second * -2)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().SubSeconds(2).ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_AddSecond(t *testing.T) {
	now := time.Now().Add(time.Second)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().AddSecond().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}

func TestCarbon_SubSecond(t *testing.T) {
	now := time.Now().Add(-time.Second)
	e := now.Format("2006-01-02 15:04:05")
	r := Now().SubSecond().ToDateTimeString()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}
