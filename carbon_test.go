package carbon

import (
	"testing"
	"time"
)

func TestToday(t *testing.T) {
	c := New()
	e := time.Now().Format("2006-01-02 15:04:05")
	r := c.Now().Today()
	if r != e {
		t.Fatalf("Expected %s, but got %s", e, r)
	}
}
