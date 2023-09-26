package carbon

import (
	"bytes"
	"fmt"
	"strconv"
)

// DateTime defines a DateTime struct.
// 定义 DateTime 结构体
type DateTime Carbon

// DateTimeMilli defines a DateTimeMilli struct.
// 定义 DateTimeMilli 结构体
type DateTimeMilli Carbon

// DateTimeMicro defines a DateTimeMicro struct.
// 定义 DateTimeMicro 结构体
type DateTimeMicro Carbon

// DateTimeNano defines a DateTimeNano struct.
// 定义 DateTimeNano 结构体
type DateTimeNano Carbon

// Date defines a Date struct.
// 定义 Date 结构体
type Date Carbon

// DateMilli defines a DateMilli struct.
// 定义 DateMilli 结构体
type DateMilli Carbon

// DateMicro defines a DateMicro struct.
// 定义 DateMicro 结构体
type DateMicro Carbon

// DateNano defines a DateNano struct.
// 定义 DateNano 结构体
type DateNano Carbon

// Time defines a Time struct.
// 定义 Time 结构体
type Time Carbon

// TimeMilli defines a TimeMilli struct.
// 定义 TimeMilli 结构体
type TimeMilli Carbon

// TimeMicro defines a TimeMicro struct.
// 定义 TimeMicro 结构体
type TimeMicro Carbon

// TimeNano defines a TimeNano struct.
// 定义 TimeNano 结构体
type TimeNano Carbon

// Timestamp defines a Timestamp struct.
// 定义 Timestamp 结构体
type Timestamp Carbon

// TimestampMilli defines a TimestampMilli struct.
// 定义 TimestampMilli 结构体
type TimestampMilli Carbon

// TimestampMicro defines a TimestampMicro struct.
// 定义 TimestampMicro 结构体
type TimestampMicro Carbon

// TimestampNano defines a TimestampNano struct.
// 定义 TimestampNano 结构体
type TimestampNano Carbon

// MarshalJSON implements the interface json.Marshal for DateTime struct.
// 实现 MarshalJSON 接口
func (t DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, Carbon(t).ToDateTimeString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for DateTime struct.
// 实现 UnmarshalJSON 接口
func (t *DateTime) UnmarshalJSON(b []byte) error {
	c := ParseByLayout(string(bytes.Trim(b, `"`)), DateTimeLayout, t.loc.String())
	if c.Error == nil {
		*t = DateTime(c)
	}
	return c.Error
}

// MarshalJSON implements the interface json.Marshal for DateTimeMilli struct.
// 实现 MarshalJSON 接口
func (t DateTimeMilli) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, Carbon(t).ToDateTimeMilliString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for DateTimeMilli struct.
// 实现 UnmarshalJSON 接口
func (t *DateTimeMilli) UnmarshalJSON(b []byte) error {
	c := ParseByLayout(string(bytes.Trim(b, `"`)), DateTimeMilliLayout, t.loc.String())
	if c.Error == nil {
		*t = DateTimeMilli(c)
	}
	return c.Error
}

// MarshalJSON implements the interface json.Marshal for DateTimeMicro struct.
// 实现 MarshalJSON 接口
func (t DateTimeMicro) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, Carbon(t).ToDateTimeMicroString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for DateTimeMicro struct.
// 实现 UnmarshalJSON 接口
func (t *DateTimeMicro) UnmarshalJSON(b []byte) error {
	c := ParseByLayout(string(bytes.Trim(b, `"`)), DateTimeMicroLayout, t.loc.String())
	if c.Error == nil {
		*t = DateTimeMicro(c)
	}
	return c.Error
}

// MarshalJSON implements the interface json.Marshal for DateTimeNano struct.
// 实现 MarshalJSON 接口
func (t DateTimeNano) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, Carbon(t).ToDateTimeNanoString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for DateTimeNano struct.
// 实现 UnmarshalJSON 接口
func (t *DateTimeNano) UnmarshalJSON(b []byte) error {
	c := ParseByLayout(string(bytes.Trim(b, `"`)), DateTimeNanoLayout, t.loc.String())
	if c.Error == nil {
		*t = DateTimeNano(c)
	}
	return c.Error
}

// MarshalJSON implements the interface json.Marshal for Date struct.
// 实现 MarshalJSON 接口
func (t Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, Carbon(t).ToDateString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for Date struct.
// 实现 UnmarshalJSON 接口
func (t *Date) UnmarshalJSON(b []byte) error {
	c := ParseByLayout(string(bytes.Trim(b, `"`)), DateLayout, t.loc.String())
	if c.Error == nil {
		*t = Date(c)
	}
	return c.Error
}

// MarshalJSON implements the interface json.Marshal for DateMilli struct.
// 实现 MarshalJSON 接口
func (t DateMilli) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, Carbon(t).ToDateMilliString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for DateMilli struct.
// 实现 UnmarshalJSON 接口
func (t *DateMilli) UnmarshalJSON(b []byte) error {
	c := ParseByLayout(string(bytes.Trim(b, `"`)), DateMilliLayout, t.loc.String())
	if c.Error == nil {
		*t = DateMilli(c)
	}
	return c.Error
}

// MarshalJSON implements the interface json.Marshal for DateMicro struct.
// 实现 MarshalJSON 接口
func (t DateMicro) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, Carbon(t).ToDateMicroString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for DateMicro struct.
// 实现 UnmarshalJSON 接口
func (t *DateMicro) UnmarshalJSON(b []byte) error {
	c := ParseByLayout(string(bytes.Trim(b, `"`)), DateMicroLayout, t.loc.String())
	if c.Error == nil {
		*t = DateMicro(c)
	}
	return c.Error
}

// MarshalJSON implements the interface json.Marshal for DateNano struct.
// 实现 MarshalJSON 接口
func (t DateNano) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, Carbon(t).ToDateNanoString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for DateNano struct.
// 实现 UnmarshalJSON 接口
func (t *DateNano) UnmarshalJSON(b []byte) error {
	c := ParseByLayout(string(bytes.Trim(b, `"`)), DateNanoLayout, t.loc.String())
	if c.Error == nil {
		*t = DateNano(c)
	}
	return c.Error
}

// MarshalJSON implements the interface json.Marshal for Time struct.
// 实现 MarshalJSON 接口
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, Carbon(t).ToTimeString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for Time struct.
// 实现 UnmarshalJSON 接口
func (t *Time) UnmarshalJSON(b []byte) error {
	c := ParseByLayout(string(bytes.Trim(b, `"`)), TimeLayout, t.loc.String())
	if c.Error == nil {
		*t = Time(c)
	}
	return c.Error
}

// MarshalJSON implements the interface json.Marshal for TimeMilli struct.
// 实现 MarshalJSON 接口
func (t TimeMilli) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, Carbon(t).ToTimeMilliString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for TimeMilli struct.
// 实现 UnmarshalJSON 接口
func (t *TimeMilli) UnmarshalJSON(b []byte) error {
	c := ParseByLayout(string(bytes.Trim(b, `"`)), TimeMilliLayout, t.loc.String())
	if c.Error == nil {
		*t = TimeMilli(c)
	}
	return c.Error
}

// MarshalJSON implements the interface json.Marshal for TimeMicro struct.
// 实现 MarshalJSON 接口
func (t TimeMicro) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, Carbon(t).ToTimeMicroString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for TimeMicro struct.
// 实现 UnmarshalJSON 接口
func (t *TimeMicro) UnmarshalJSON(b []byte) error {
	c := ParseByLayout(string(bytes.Trim(b, `"`)), TimeMicroLayout, t.loc.String())
	if c.Error == nil {
		*t = TimeMicro(c)
	}
	return c.Error
}

// MarshalJSON implements the interface json.Marshal for TimeNano struct.
// 实现 MarshalJSON 接口
func (t TimeNano) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, Carbon(t).ToTimeNanoString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for TimeNano struct.
// 实现 UnmarshalJSON 接口
func (t *TimeNano) UnmarshalJSON(b []byte) error {
	c := ParseByLayout(string(bytes.Trim(b, `"`)), TimeNanoLayout, t.loc.String())
	if c.Error == nil {
		*t = TimeNano(c)
	}
	return c.Error
}

// MarshalJSON implements the interface json.Marshal for Timestamp struct.
// 实现 MarshalJSON 接口
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, Carbon(t).Timestamp())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for Timestamp struct.
// 实现 UnmarshalJSON 接口
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, _ := strconv.ParseInt(string(b), 10, 64)
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = Timestamp(c)
	}
	return c.Error
}

// MarshalJSON implements the interface json.Marshal for TimestampMilli struct.
// 实现 MarshalJSON 接口
func (t TimestampMilli) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, Carbon(t).TimestampMilli())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for TimestampMilli struct.
// 实现 UnmarshalJSON 接口
func (t *TimestampMilli) UnmarshalJSON(b []byte) error {
	ts, _ := strconv.ParseInt(string(b), 10, 64)
	c := CreateFromTimestampMilli(ts)
	if c.Error == nil {
		*t = TimestampMilli(c)
	}
	return c.Error
}

// MarshalJSON implements the interface MarshalJSON for TimestampMicro struct.
// 实现 MarshalJSON 接口
func (t TimestampMicro) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, Carbon(t).TimestampMicro())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for TimestampMicro struct.
// 实现 UnmarshalJSON 接口
func (t *TimestampMicro) UnmarshalJSON(b []byte) error {
	ts, _ := strconv.ParseInt(string(b), 10, 64)
	c := CreateFromTimestampMicro(ts)
	if c.Error == nil {
		*t = TimestampMicro(c)
	}
	return c.Error
}

// MarshalJSON implements the interface json.Marshal for TimestampNano struct.
// 实现 MarshalJSON 接口
func (t TimestampNano) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, Carbon(t).TimestampNano())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for TimestampNano struct.
// 实现 UnmarshalJSON 接口
func (t *TimestampNano) UnmarshalJSON(b []byte) error {
	ts, _ := strconv.ParseInt(string(b), 10, 64)
	c := CreateFromTimestampNano(ts)
	if c.Error == nil {
		*t = TimestampNano(c)
	}
	return c.Error
}

// ToDateTimeStruct converts Carbon to DateTime.
// 将 Carbon 结构体转换成 DateTime 结构体
func (c Carbon) ToDateTimeStruct() DateTime {
	return DateTime(c)
}

// ToDateTimeMilliStruct converts Carbon to DateTimeMilli.
// 将 Carbon 结构体转换成 DateTimeMilli 结构体
func (c Carbon) ToDateTimeMilliStruct() DateTimeMilli {
	return DateTimeMilli(c)
}

// ToDateTimeMicroStruct converts Carbon to DateTimeMicro.
// 将 Carbon 结构体转换成 DateTimeMicro 结构体
func (c Carbon) ToDateTimeMicroStruct() DateTimeMicro {
	return DateTimeMicro(c)
}

// ToDateTimeNanoStruct converts Carbon to DateTimeNano.
// 将 Carbon 结构体转换成 DateTimeNano 结构体
func (c Carbon) ToDateTimeNanoStruct() DateTimeNano {
	return DateTimeNano(c)
}

// ToDateStruct converts Carbon to Date.
// 将 Carbon 结构体转换成 Date 结构体
func (c Carbon) ToDateStruct() Date {
	return Date(c)
}

// ToDateMilliStruct converts Carbon to DateMilli.
// 将 Carbon 结构体转换成 DateMilli 结构体
func (c Carbon) ToDateMilliStruct() DateMilli {
	return DateMilli(c)
}

// ToDateMicroStruct converts Carbon to DateMicro.
// 将 Carbon 结构体转换成 DateMicro 结构体
func (c Carbon) ToDateMicroStruct() DateMicro {
	return DateMicro(c)
}

// ToDateNanoStruct converts Carbon to DateNano.
// 将 Carbon 结构体转换成 DateNano 结构体
func (c Carbon) ToDateNanoStruct() DateNano {
	return DateNano(c)
}

// ToTimeStruct converts Carbon to Time.
// 将 Carbon 结构体转换成 Time 结构体
func (c Carbon) ToTimeStruct() Time {
	return Time(c)
}

// ToTimeMilliStruct converts Carbon to TimeMilli.
// 将 Carbon 结构体转换成 TimeMilli 结构体
func (c Carbon) ToTimeMilliStruct() TimeMilli {
	return TimeMilli(c)
}

// ToTimeMicroStruct converts Carbon to TimeMicro.
// 将 Carbon 结构体转换成 TimeMicro 结构体
func (c Carbon) ToTimeMicroStruct() TimeMicro {
	return TimeMicro(c)
}

// ToTimeNanoStruct converts Carbon to TimeNano.
// 将 Carbon 结构体转换成 TimeNano 结构体
func (c Carbon) ToTimeNanoStruct() TimeNano {
	return TimeNano(c)
}

// ToTimestampStruct converts Carbon to Timestamp.
// 将 Carbon 结构体转换成 Timestamp 结构体
func (c Carbon) ToTimestampStruct() Timestamp {
	return Timestamp(c)
}

// ToTimestampMilliStruct converts Carbon to TimestampMilli.
// 将 Carbon 结构体转换成 TimestampMilli 结构体
func (c Carbon) ToTimestampMilliStruct() TimestampMilli {
	return TimestampMilli(c)
}

// ToTimestampMicroStruct converts Carbon to TimestampMicro.
// 将 Carbon 结构体转换成 TimestampMicro 结构体
func (c Carbon) ToTimestampMicroStruct() TimestampMicro {
	return TimestampMicro(c)
}

// ToTimestampNanoStruct converts Carbon to TimestampNano.
// 将 Carbon 结构体转换成 TimestampNano 结构体
func (c Carbon) ToTimestampNanoStruct() TimestampNano {
	return TimestampNano(c)
}

// Int64 outputs timestamp with second.
// 输出秒级时间戳
func (t Timestamp) Int64() int64 {
	return Carbon(t).Timestamp()
}

// Int64 outputs timestamp with millisecond.
// 输出豪秒级时间戳
func (t TimestampMilli) Int64() int64 {
	return Carbon(t).TimestampMilli()
}

// Int64 outputs timestamp with microsecond.
// 输出微秒级时间戳
func (t TimestampMicro) Int64() int64 {
	return Carbon(t).TimestampMicro()
}

// Int64 outputs timestamp with nanosecond.
// 输出纳秒级时间戳
func (t TimestampNano) Int64() int64 {
	return Carbon(t).TimestampNano()
}

// String implements the interface Stringer for DateTime struct.
// 实现 Stringer 接口
func (t DateTime) String() string {
	return Carbon(t).ToDateTimeString()
}

// String implements the interface Stringer for DateTimeMilli struct.
// 实现 Stringer 接口
func (t DateTimeMilli) String() string {
	return Carbon(t).ToDateTimeMilliString()
}

// String implements the interface Stringer for DateTimeMicro struct.
// 实现 Stringer 接口
func (t DateTimeMicro) String() string {
	return Carbon(t).ToDateTimeMicroString()
}

// String implements the interface Stringer for DateTimeNano struct.
// 实现 Stringer 接口
func (t DateTimeNano) String() string {
	return Carbon(t).ToDateTimeNanoString()
}

// String implements the interface Stringer for Date struct.
// 实现 Stringer 接口
func (t Date) String() string {
	return Carbon(t).ToDateString()
}

// String implements the interface Stringer for DateMilli struct.
// 实现 Stringer 接口
func (t DateMilli) String() string {
	return Carbon(t).ToDateMilliString()
}

// String implements the interface Stringer for DateMicro struct.
// 实现 Stringer 接口
func (t DateMicro) String() string {
	return Carbon(t).ToDateMicroString()
}

// String implements the interface Stringer for DateNano struct.
// 实现 Stringer 接口
func (t DateNano) String() string {
	return Carbon(t).ToDateNanoString()
}

// String implements the interface Stringer for Time struct.
// 实现 Stringer 接口
func (t Time) String() string {
	return Carbon(t).ToTimeString()
}

// String implements the interface Stringer for TimeMilli struct.
// 实现 Stringer 接口
func (t TimeMilli) String() string {
	return Carbon(t).ToTimeMilliString()
}

// String implements the interface Stringer for TimeMicro struct.
// 实现 Stringer 接口
func (t TimeMicro) String() string {
	return Carbon(t).ToTimeMicroString()
}

// String implements the interface Stringer for TimeNano struct.
// 实现 Stringer 接口
func (t TimeNano) String() string {
	return Carbon(t).ToTimeNanoString()
}

// String implements the interface Stringer for Timestamp struct.
// 实现 Stringer 接口
func (t Timestamp) String() string {
	return strconv.FormatInt(Carbon(t).Timestamp(), 10)
}

// String implements the interface Stringer for TimestampMilli struct.
// 实现 Stringer 接口
func (t TimestampMilli) String() string {
	return strconv.FormatInt(Carbon(t).TimestampMilli(), 10)
}

// String implements the interface Stringer for TimestampMicro struct.
// 实现 Stringer 接口
func (t TimestampMicro) String() string {
	return strconv.FormatInt(Carbon(t).TimestampMicro(), 10)
}

// String implements the interface Stringer for TimestampNano struct.
// 实现 Stringer 接口
func (t TimestampNano) String() string {
	return strconv.FormatInt(Carbon(t).TimestampNano(), 10)
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
