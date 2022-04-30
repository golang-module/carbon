package carbon

import (
	"bytes"
	"fmt"
	"strconv"
)

// DateTime defines a DateTime struct.
type DateTime struct {
	Carbon
}

// Date defines a Date struct.
type Date struct {
	Carbon
}

// Timestamp defines a Timestamp struct.
type Timestamp struct {
	Carbon
}

// TimestampMilli defines a TimestampMilli struct.
type TimestampMilli struct {
	Carbon
}

// TimestampMicro defines a TimestampMicro struct.
type TimestampMicro struct {
	Carbon
}

// TimestampNano defines a TimestampNano struct.
type TimestampNano struct {
	Carbon
}

// MarshalJSON implements the interface json.Marshal for Date struct.
func (t DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateTimeString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for Date struct.
func (t *DateTime) UnmarshalJSON(b []byte) error {
	c := Parse(string(bytes.Trim(b, `"`)))
	if c.Error == nil {
		*t = DateTime{c}
	}
	return nil
}

// MarshalJSON implements the interface json.Marshal for Date struct.
func (t Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for Date struct.
func (t *Date) UnmarshalJSON(b []byte) error {
	c := Parse(string(bytes.Trim(b, `"`)))
	if c.Error == nil {
		*t = Date{c}
	}
	return nil
}

// MarshalJSON implements the interface json.Marshal for Timestamp struct.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.Timestamp())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for Timestamp struct.
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, _ := strconv.ParseInt(string(b), 10, 64)
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = Timestamp{c}
	}
	return nil
}

// MarshalJSON implements the interface json.Marshal for TimestampMilli struct.
func (t TimestampMilli) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.TimestampMilli())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for TimestampMilli struct.
func (t *TimestampMilli) UnmarshalJSON(b []byte) error {
	ts, _ := strconv.ParseInt(string(b), 10, 64)
	c := CreateFromTimestampMilli(ts)
	if c.Error == nil {
		*t = TimestampMilli{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for TimestampMicro struct.
func (t TimestampMicro) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.TimestampMicro())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for TimestampMicro struct.
func (t *TimestampMicro) UnmarshalJSON(b []byte) error {
	ts, _ := strconv.ParseInt(string(b), 10, 64)
	c := CreateFromTimestampMicro(ts)
	if c.Error == nil {
		*t = TimestampMicro{c}
	}
	return nil
}

// UnmarshalJSON implements the interface json.Unmarshal for TimestampNano struct.
func (t *TimestampNano) UnmarshalJSON(b []byte) error {
	ts, _ := strconv.ParseInt(string(b), 10, 64)
	c := CreateFromTimestampNano(ts)
	if c.Error == nil {
		*t = TimestampNano{c}
	}
	return nil
}

// MarshalJSON implements the interface json.Marshal for TimestampNano struct.
func (t TimestampNano) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.TimestampNano())), nil
}
