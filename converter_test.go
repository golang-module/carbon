package carbon

import (
	"testing"
	"time"
)

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
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-08-05", "2020-08-05"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Carbon2Time().Format("2006-01-02")

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}
