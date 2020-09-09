package carbon

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type ToTimeString struct {
	time.Time
}

func (t ToTimeString) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("15-04-05"))
	return []byte(formatted), nil
}

func (t ToTimeString) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *ToTimeString) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = ToTimeString{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
