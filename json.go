package carbon

import (
	"fmt"
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

func (c ToTimestampWithSecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, c.ToTimestampWithSecond())), nil
}

func (c ToTimestampWithMillisecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, c.ToTimestampWithMillisecond())), nil
}

func (c ToTimestampWithMicrosecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, c.ToTimestampWithMicrosecond())), nil
}

func (c ToTimestampWithNanosecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, c.ToTimestampWithNanosecond())), nil
}
