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

// Time defines a Time struct.
type Time struct {
	Carbon
}

// Timestamp defines a Timestamp struct.
type Timestamp struct {
	Carbon
}

// TimestampWithSecond defines a TimestampWithSecond struct.
type TimestampWithSecond struct {
	Carbon
}

// TimestampWithMillisecond defines a TimestampWithMillisecond struct.
type TimestampWithMillisecond struct {
	Carbon
}

// TimestampWithMicrosecond defines a TimestampWithMicrosecond struct.
type TimestampWithMicrosecond struct {
	Carbon
}

// TimestampWithNanosecond defines a TimestampWithNanosecond struct.
type TimestampWithNanosecond struct {
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
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToTimeString())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *Time) UnmarshalJSON(b []byte) error {
	c := ParseByFormat(string(bytes.Trim(b, `"`)), "H:i:s")
	if c.Error == nil {
		*t = Time{c}
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
	return []byte(fmt.Sprintf(`%d`, t.TimestampWithSecond())), nil
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
	return []byte(fmt.Sprintf(`%d`, t.TimestampWithMillisecond())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *TimestampWithMillisecond) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = TimestampWithMillisecond{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t TimestampWithMicrosecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.TimestampWithMicrosecond())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *TimestampWithMicrosecond) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = TimestampWithMicrosecond{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t TimestampWithNanosecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.TimestampWithNanosecond())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *TimestampWithNanosecond) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = TimestampWithNanosecond{c}
	}
	return nil
}
