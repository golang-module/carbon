package carbon

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type ToDateTimeString struct {
	time.Time
}

func (t ToDateTimeString) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil

}
func (t ToDateTimeString) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}
func (t *ToDateTimeString) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = ToDateTimeString{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
