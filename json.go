package carbon

import (
	"bytes"
	"fmt"
	"strconv"
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

type ToTimestampWithSecond struct {
	Carbon
}

type ToTimestampWithMillisecond struct {
	Carbon
}

type ToTimestampWithMicrosecond struct {
	Carbon
}

type ToTimestampWithNanosecond struct {
	Carbon
}

func (t ToDateTimeString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateTimeString())), nil
}

func (t *ToDateTimeString) UnmarshalJSON(b []byte) error {
	c := Parse(string(bytes.Trim(b, `"`)))
	if c.Error == nil {
		*t = ToDateTimeString{c}
	}
	return nil
}

func (t ToDateString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateString())), nil
}

func (t *ToDateString) UnmarshalJSON(b []byte) error {
	c := Parse(string(bytes.Trim(b, `"`)))
	if c.Error == nil {
		*t = ToDateString{c}
	}
	return nil
}

func (t ToTimeString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToTimeString())), nil
}

func (t *ToTimeString) UnmarshalJSON(b []byte) error {
	c := ParseByFormat(string(bytes.Trim(b, `"`)), "H:i:s")
	if c.Error == nil {
		*t = ToTimeString{c}
	}
	return nil
}

func (t ToTimestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.ToTimestamp())), nil
}

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

func (t ToTimestampWithSecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.ToTimestampWithSecond())), nil
}

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

func (t ToTimestampWithMillisecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.ToTimestampWithMillisecond())), nil
}

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

func (t ToTimestampWithMicrosecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.ToTimestampWithMicrosecond())), nil
}

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

func (t ToTimestampWithNanosecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.ToTimestampWithNanosecond())), nil
}

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
