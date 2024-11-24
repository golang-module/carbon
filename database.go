package carbon

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

// returns a failed scan error.
// 失败的扫描错误
var failedScanError = func(src interface{}) error {
	return errors.New(fmt.Sprintf("failed to scan value: %v", src))
}

// Scan an interface used by Scan in package database/sql for Scanning value from database to local golang variable.
func (c *Carbon) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		*c = Parse(string(v))
	case string:
		*c = Parse(v)
	case time.Time:
		*c = CreateFromStdTime(v)
	default:
		return failedScanError(src)
	}
	if c.Error != nil {
		return failedScanError(src)
	}
	return nil
}

// Value the interface providing the Value method for package database/sql/driver.
func (c Carbon) Value() (driver.Value, error) {
	if c.IsZero() {
		return nil, nil
	}
	return c.StdTime(), nil
}

// GormDataType implements the interface GormDataTypeInterface for Carbon struct.
// 实现 GormDataTypeInterface 接口
func (c Carbon) GormDataType() string {
	return "time"
}

// GormDataType implements the interface GormDataTypeInterface for DateTime struct.
// 实现 GormDataTypeInterface 接口
func (t DateTime) GormDataType() string {
	return "time"
}

// GormDataType implements the interface GormDataTypeInterface for DateTimeMilli struct.
// 实现 GormDataTypeInterface 接口
func (t DateTimeMilli) GormDataType() string {
	return "time"
}

// GormDataType implements the interface GormDataTypeInterface for DateTimeMicro struct.
// 实现 GormDataTypeInterface 接口
func (t DateTimeMicro) GormDataType() string {
	return "time"
}

// GormDataType implements the interface GormDataTypeInterface for DateTimeNano struct.
// 实现 GormDataTypeInterface 接口
func (t DateTimeNano) GormDataType() string {
	return "time"
}

// GormDataType implements the interface GormDataTypeInterface for Date struct.
// 实现 GormDataTypeInterface 接口
func (t Date) GormDataType() string {
	return "time"
}

// GormDataType implements the interface GormDataTypeInterface for DateMilli struct.
// 实现 GormDataTypeInterface 接口
func (t DateMilli) GormDataType() string {
	return "time"
}

// GormDataType implements the interface GormDataTypeInterface for DateMicro struct.
// 实现 GormDataTypeInterface 接口
func (t DateMicro) GormDataType() string {
	return "time"
}

// GormDataType implements the interface GormDataTypeInterface for DateNano struct.
// 实现 GormDataTypeInterface 接口
func (t DateNano) GormDataType() string {
	return "time"
}

// GormDataType implements the interface GormDataTypeInterface for Time struct.
// 实现 GormDataTypeInterface 接口
func (t Time) GormDataType() string {
	return "time"
}

// GormDataType implements the interface GormDataTypeInterface for TimeMilli struct.
// 实现 GormDataTypeInterface 接口
func (t TimeMilli) GormDataType() string {
	return "time"
}

// GormDataType implements the interface GormDataTypeInterface for TimeMicro struct.
// 实现 GormDataTypeInterface 接口
func (t TimeMicro) GormDataType() string {
	return "time"
}

// GormDataType implements the interface GormDataTypeInterface for TimeNano struct.
// 实现 GormDataTypeInterface 接口
func (t TimeNano) GormDataType() string {
	return "time"
}

// GormDataType implements the interface GormDataTypeInterface for Timestamp struct.
// 实现 GormDataTypeInterface 接口
func (t Timestamp) GormDataType() string {
	return "int"
}

// GormDataType implements the interface GormDataTypeInterface for TimestampMilli struct.
// 实现 GormDataTypeInterface 接口
func (t TimestampMilli) GormDataType() string {
	return "int"
}

// GormDataType implements the interface GormDataTypeInterface for TimestampMicro struct.
// 实现 GormDataTypeInterface 接口
func (t TimestampMicro) GormDataType() string {
	return "int"
}

// GormDataType implements the interface GormDataTypeInterface for TimestampNano struct.
// 实现 GormDataTypeInterface 接口
func (t TimestampNano) GormDataType() string {
	return "int"
}
