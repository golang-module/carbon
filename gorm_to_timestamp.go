package carbon

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type ToTimestamp struct {
	time.Time
}

func (t ToTimestamp) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%d\"", t.Unix())
	return []byte(formatted), nil
}

func (t ToTimestamp) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *ToTimestamp) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = ToTimestamp{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
