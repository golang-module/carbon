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

// TimestampWithSecond defines a TimestampWithSecond struct, it will be replaced by Timestamp.
type TimestampWithSecond struct {
	Carbon
}

// Timestamp defines a Timestamp struct.
type Timestamp struct {
	Carbon
}

// TimestampWithMillisecond defines a TimestampWithMillisecond struct, it will be replaced by TimestampMilli.
type TimestampWithMillisecond struct {
	Carbon
}

// TimestampMilli defines a TimestampMilli struct.
type TimestampMilli struct {
	Carbon
}

// TimestampWithMicrosecond defines a TimestampWithMicrosecond struct, it will be replaced by TimestampMicro.
type TimestampWithMicrosecond struct {
	Carbon
}

// TimestampMicro defines a TimestampMicro struct.
type TimestampMicro struct {
	Carbon
}

// TimestampWithNanosecond defines a TimestampWithNanosecond struct, it will be replaced by TimestampNano.
type TimestampWithNanosecond struct {
	Carbon
}

// TimestampNano defines a TimestampNano struct.
type TimestampNano struct {
	Carbon
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateTimeString())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *DateTime) UnmarshalJSON(b []byte) error {
	c := Parse(string(bytes.Trim(b, `"`)))
	if c.Error == nil {
		*t = DateTime{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateString())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *Date) UnmarshalJSON(b []byte) error {
	c := Parse(string(bytes.Trim(b, `"`)))
	if c.Error == nil {
		*t = Date{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.Timestamp())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = Timestamp{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t TimestampWithSecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.Timestamp())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *TimestampWithSecond) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = TimestampWithSecond{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t TimestampWithMillisecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.TimestampMilli())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *TimestampWithMillisecond) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestampMilli(ts)
	if c.Error == nil {
		*t = TimestampWithMillisecond{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t TimestampMilli) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.TimestampMilli())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *TimestampMilli) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestampMilli(ts)
	if c.Error == nil {
		*t = TimestampMilli{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t TimestampWithMicrosecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.TimestampMicro())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *TimestampWithMicrosecond) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestampMicro(ts)
	if c.Error == nil {
		*t = TimestampWithMicrosecond{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t TimestampMicro) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.TimestampMicro())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *TimestampMicro) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestampMicro(ts)
	if c.Error == nil {
		*t = TimestampMicro{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t TimestampWithNanosecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.TimestampNano())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *TimestampWithNanosecond) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestampNano(ts)
	if c.Error == nil {
		*t = TimestampWithNanosecond{c}
	}
	return nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *TimestampNano) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestampNano(ts)
	if c.Error == nil {
		*t = TimestampNano{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t TimestampNano) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.TimestampNano())), nil
}
