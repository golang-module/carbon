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
