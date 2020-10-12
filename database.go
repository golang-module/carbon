package carbon

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type ToDateTimeString struct {
	Carbon
}

type ToDateString struct {
	Carbon
}

type ToTimeString struct {
	Carbon
}

type ToTimestamp struct {
	Carbon
}

func (c *Carbon) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*c = Carbon{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
func (c Carbon) Value() (driver.Value, error) {
	var zeroTime time.Time
	var timeTime = c.Time
	if timeTime.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return timeTime, nil
}

func (c Carbon) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, c.ToDateTimeString())), nil
}

func (c ToDateTimeString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, c.ToDateTimeString())), nil
}

func (c ToDateString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, c.ToDateString())), nil
}

func (c ToTimeString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, c.ToTimeString())), nil
}

func (c ToTimestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, c.ToTimestamp())), nil
}
