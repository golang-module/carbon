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

func TestCarbon_Time2Carbon(t *testing.T) {
	Tests := []struct {
		time   time.Time // // 输入参数
		output string    // 期望输出值
	}{
		{time.Now(), time.Now().Format(DateTimeFormat)},
	}

	for _, v := range Tests {
		output := Time2Carbon(v.time).ToDateTimeString()

		if output != v.output {
			t.Errorf("Expected %s, but got %s", v.output, output)
		}
	}
}

func TestCarbon_Carbon2Time(t *testing.T) {
	Tests := []struct {
		input  Carbon // 输入值
		output string // 期望输出值
	}{
		{Parse("2020-08-05"), "2020-08-05"},
	}

	for _, v := range Tests {
		output := v.input.Carbon2Time().Format("2006-01-02")

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input.ToDateString(), v.output, output)
		}
	}
}
