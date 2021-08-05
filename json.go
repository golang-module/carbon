package carbon

import (
	"bytes"
	"fmt"
	"strconv"
)

// ToDateTimeString define ToDateTimeString structure
type ToDateTimeString struct {
	Carbon
}

// ToDateString define ToDateTimeString structure
type ToDateString struct {
	Carbon
}

// ToTimeString define ToTimeString structure
type ToTimeString struct {
	Carbon
}

// ToTimestamp define ToTimestamp structure
type ToTimestamp struct {
	Carbon
}

// ToTimestampWithSecond define ToTimestampWithSecond structure
type ToTimestampWithSecond struct {
	Carbon
}

// ToTimestampWithMillisecond define ToTimestampWithMillisecond structure
type ToTimestampWithMillisecond struct {
	Carbon
}

// ToTimestampWithMicrosecond define ToTimestampWithMicrosecond structure
type ToTimestampWithMicrosecond struct {
	Carbon
}

// ToTimestampWithNanosecond define ToTimestampWithNanosecond structure
type ToTimestampWithNanosecond struct {
	Carbon
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t ToDateTimeString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateTimeString())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *ToDateTimeString) UnmarshalJSON(b []byte) error {
	c := Parse(string(bytes.Trim(b, `"`)))
	if c.Error == nil {
		*t = ToDateTimeString{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t ToDateString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateString())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *ToDateString) UnmarshalJSON(b []byte) error {
	c := Parse(string(bytes.Trim(b, `"`)))
	if c.Error == nil {
		*t = ToDateString{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t ToTimeString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToTimeString())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *ToTimeString) UnmarshalJSON(b []byte) error {
	c := ParseByFormat(string(bytes.Trim(b, `"`)), "H:i:s")
	if c.Error == nil {
		*t = ToTimeString{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t ToTimestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.ToTimestamp())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *ToTimestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = ToTimestamp{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t ToTimestampWithSecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.ToTimestampWithSecond())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *ToTimestampWithSecond) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = ToTimestampWithSecond{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t ToTimestampWithMillisecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.ToTimestampWithMillisecond())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *ToTimestampWithMillisecond) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = ToTimestampWithMillisecond{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t ToTimestampWithMicrosecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.ToTimestampWithMicrosecond())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *ToTimestampWithMicrosecond) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = ToTimestampWithMicrosecond{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t ToTimestampWithNanosecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.ToTimestampWithNanosecond())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *ToTimestampWithNanosecond) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = ToTimestampWithNanosecond{c}
	}
	return nil
}
