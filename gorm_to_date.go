package carbon

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type ToDateString struct {
	time.Time
}

func (t ToDateString) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02"))
	return []byte(formatted), nil
}

func (t ToDateString) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}
func (t *ToDateString) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = ToDateString{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
