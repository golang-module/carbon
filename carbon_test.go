package carbon

import (
	"testing"
	"time"
)

func TestCarbon_Now(t *testing.T) {
	expected := time.Now().Format(DateTimeFormat)

	output := Now().Time.Format(DateTimeFormat)

	if expected != output {
		t.Fatalf("Expected %s, but got %s", expected, output)
	}

	output = SetTimezone(PRC).Now().Time.Format(DateTimeFormat)

	if expected != output {
		t.Fatalf("Expected %s, but got %s", expected, output)
	}

}

func TestCarbon_Yesterday(t *testing.T) {
	expected := time.Now().AddDate(0, 0, -1).Format(DateTimeFormat)

	output := Yesterday().Time.Format(DateTimeFormat)

	if expected != output {
		t.Fatalf("Expected %s, but got %s", expected, output)
	}

	output = SetTimezone(PRC).Parse("2020-08-05").Yesterday().Time.Format(DateFormat)

	if "2020-08-04" != output {
		t.Fatalf("Expected %s, but got %s", expected, output)
	}
}

func TestCarbon_Tomorrow(t *testing.T) {
	expected := time.Now().AddDate(0, 0, 1).Format(DateTimeFormat)

	output := Tomorrow().Time.Format(DateTimeFormat)

	if expected != output {
		t.Errorf("Expected %s, but got %s", expected, output)
	}

	output = SetTimezone(PRC).Parse("2020-08-05").Tomorrow().Time.Format(DateFormat)

	if "2020-08-06" != output {
		t.Errorf("Expected %s, but got %s", expected, output)
	}
}
