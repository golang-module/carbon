package carbon

import (
	"database/sql/driver"
	"fmt"
	"time"
)

func (c *Carbon) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*c = Carbon{Time: value, Loc: time.Local}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (c Carbon) Value() (driver.Value, error) {
	if c.IsZero() {
		return nil, nil
	}
	return c.Time, nil
}
