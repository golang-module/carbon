package carbon

import (
	"bytes"
	"fmt"
	"strconv"
)

// DateTime defines a DateTime struct.
// 定义 DateTime 结构体
type DateTime struct {
	Carbon
}

// DateTimeMilli defines a DateTimeMilli struct.
// 定义 DateTimeMilli 结构体
type DateTimeMilli struct {
	Carbon
}

// DateTimeMicro defines a DateTimeMicro struct.
// 定义 DateTimeMicro 结构体
type DateTimeMicro struct {
	Carbon
}

// DateTimeNano defines a DateTimeNano struct.
// 定义 DateTimeNano 结构体
type DateTimeNano struct {
	Carbon
}

// Date defines a Date struct.
// 定义 Date 结构体
type Date struct {
	Carbon
}

// DateMilli defines a DateMilli struct.
// 定义 DateMilli 结构体
type DateMilli struct {
	Carbon
}

// DateMicro defines a DateMicro struct.
// 定义 DateMicro 结构体
type DateMicro struct {
	Carbon
}

// DateNano defines a DateNano struct.
// 定义 DateNano 结构体
type DateNano struct {
	Carbon
}

// Timestamp defines a Timestamp struct.
// 定义 Timestamp 结构体
type Timestamp struct {
	Carbon
}

// TimestampMilli defines a TimestampMilli struct.
// 定义 TimestampMilli 结构体
type TimestampMilli struct {
	Carbon
}

// TimestampMicro defines a TimestampMicro struct.
// 定义 TimestampMicro 结构体
type TimestampMicro struct {
	Carbon
}

// TimestampNano defines a TimestampNano struct.
// 定义 TimestampNano 结构体
type TimestampNano struct {
	Carbon
}

// MarshalJSON implements the interface json.Marshal for DateTime struct.
// 实现 MarshalJSON 接口
func (t DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateTimeString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for DateTime struct.
// 实现 UnmarshalJSON 接口
func (t *DateTime) UnmarshalJSON(b []byte) error {
	c := ParseByLayout(string(bytes.Trim(b, `"`)), DateTimeLayout, t.Location())
	if c.Error == nil {
		*t = DateTime{Carbon: c}
	}
	return c.Error
}

// String implements the interface Stringer for DateTime struct.
// 实现 Stringer 接口
func (t DateTime) String() string {
	return t.ToDateTimeString()
}

// ToDateTimeStruct converts Carbon to DateTime.
// 将 Carbon 结构体转换成 DateTime 结构体
func (c Carbon) ToDateTimeStruct() DateTime {
	return DateTime{Carbon: c}
}

// MarshalJSON implements the interface json.Marshal for DateTimeMilli struct.
// 实现 MarshalJSON 接口
func (t DateTimeMilli) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateTimeMilliString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for DateTimeMilli struct.
// 实现 UnmarshalJSON 接口
func (t *DateTimeMilli) UnmarshalJSON(b []byte) error {
	c := ParseByLayout(string(bytes.Trim(b, `"`)), DateTimeMilliLayout, t.Location())
	if c.Error == nil {
		*t = DateTimeMilli{c}
	}
	return c.Error
}

// String implements the interface Stringer for DateTimeMilli struct.
// 实现 Stringer 接口
func (t DateTimeMilli) String() string {
	return t.ToDateTimeMilliString()
}

// ToDateTimeMilliStruct converts Carbon to DateTimeMilli.
// 将 Carbon 结构体转换成 DateTimeMilli 结构体
func (c Carbon) ToDateTimeMilliStruct() DateTimeMilli {
	return DateTimeMilli{Carbon: c}
}

// MarshalJSON implements the interface json.Marshal for DateTimeMicro struct.
// 实现 MarshalJSON 接口
func (t DateTimeMicro) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateTimeMicroString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for DateTimeMicro struct.
// 实现 UnmarshalJSON 接口
func (t *DateTimeMicro) UnmarshalJSON(b []byte) error {
	c := ParseByLayout(string(bytes.Trim(b, `"`)), DateTimeMicroLayout, t.Location())
	if c.Error == nil {
		*t = DateTimeMicro{c}
	}
	return c.Error
}

// String implements the interface Stringer for DateTimeMicro struct.
// 实现 Stringer 接口
func (t DateTimeMicro) String() string {
	return t.ToDateTimeMicroString()
}

// ToDateTimeMicroStruct converts Carbon to DateTimeMicro.
// 将 Carbon 结构体转换成 DateTimeMicro 结构体
func (c Carbon) ToDateTimeMicroStruct() DateTimeMicro {
	return DateTimeMicro{Carbon: c}
}

// MarshalJSON implements the interface json.Marshal for DateTimeNano struct.
// 实现 MarshalJSON 接口
func (t DateTimeNano) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateTimeNanoString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for DateTimeNano struct.
// 实现 UnmarshalJSON 接口
func (t *DateTimeNano) UnmarshalJSON(b []byte) error {
	c := ParseByLayout(string(bytes.Trim(b, `"`)), DateTimeNanoLayout, t.Location())
	if c.Error == nil {
		*t = DateTimeNano{Carbon: c}
	}
	return c.Error
}

// String implements the interface Stringer for DateTimeNano struct.
// 实现 Stringer 接口
func (t DateTimeNano) String() string {
	return t.ToDateTimeNanoString()
}

// ToDateTimeNanoStruct converts Carbon to DateTimeNano.
// 将 Carbon 结构体转换成 DateTimeNano 结构体
func (c Carbon) ToDateTimeNanoStruct() DateTimeNano {
	return DateTimeNano{Carbon: c}
}

// MarshalJSON implements the interface json.Marshal for Date struct.
// 实现 MarshalJSON 接口
func (t Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for Date struct.
// 实现 UnmarshalJSON 接口
func (t *Date) UnmarshalJSON(b []byte) error {
	c := ParseByLayout(string(bytes.Trim(b, `"`)), DateLayout, t.Location())
	if c.Error == nil {
		*t = Date{Carbon: c}
	}
	return c.Error
}

// String implements the interface Stringer for Date struct.
// 实现 Stringer 接口
func (t Date) String() string {
	return t.ToDateString()
}

// ToDateStruct converts Carbon to Date.
// 将 Carbon 结构体转换成 Date 结构体
func (c Carbon) ToDateStruct() Date {
	return Date{Carbon: c}
}

// MarshalJSON implements the interface json.Marshal for DateMilli struct.
// 实现 MarshalJSON 接口
func (t DateMilli) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateMilliString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for DateMilli struct.
// 实现 UnmarshalJSON 接口
func (t *DateMilli) UnmarshalJSON(b []byte) error {
	c := ParseByLayout(string(bytes.Trim(b, `"`)), DateMilliLayout, t.Location())
	if c.Error == nil {
		*t = DateMilli{Carbon: c}
	}
	return c.Error
}

// String implements the interface Stringer for DateMilli struct.
// 实现 Stringer 接口
func (t DateMilli) String() string {
	return t.ToDateMilliString()
}

// ToDateMilliStruct converts Carbon to DateMilli.
// 将 Carbon 结构体转换成 DateMilli 结构体
func (c Carbon) ToDateMilliStruct() DateMilli {
	return DateMilli{Carbon: c}
}

// MarshalJSON implements the interface json.Marshal for DateMicro struct.
// 实现 MarshalJSON 接口
func (t DateMicro) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateMicroString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for DateMicro struct.
// 实现 UnmarshalJSON 接口
func (t *DateMicro) UnmarshalJSON(b []byte) error {
	c := ParseByLayout(string(bytes.Trim(b, `"`)), DateMicroLayout, t.Location())
	if c.Error == nil {
		*t = DateMicro{Carbon: c}
	}
	return c.Error
}

// String implements the interface Stringer for DateMicro struct.
// 实现 Stringer 接口
func (t DateMicro) String() string {
	return t.ToDateMicroString()
}

// ToDateMicroStruct converts Carbon to DateMicro.
// 将 Carbon 结构体转换成 DateMicro 结构体
func (c Carbon) ToDateMicroStruct() DateMicro {
	return DateMicro{Carbon: c}
}

// MarshalJSON implements the interface json.Marshal for DateNano struct.
// 实现 MarshalJSON 接口
func (t DateNano) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateNanoString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for DateNano struct.
// 实现 UnmarshalJSON 接口
func (t *DateNano) UnmarshalJSON(b []byte) error {
	c := ParseByLayout(string(bytes.Trim(b, `"`)), DateNanoLayout, t.Location())
	if c.Error == nil {
		*t = DateNano{Carbon: c}
	}
	return c.Error
}

// String implements the interface Stringer for DateNano struct.
// 实现 Stringer 接口
func (t DateNano) String() string {
	return t.ToDateNanoString()
}

// ToDateNanoStruct converts Carbon to DateNano.
// 将 Carbon 结构体转换成 DateNano 结构体
func (c Carbon) ToDateNanoStruct() DateNano {
	return DateNano{Carbon: c}
}

// MarshalJSON implements the interface json.Marshal for Timestamp struct.
// 实现 MarshalJSON 接口
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.Timestamp())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for Timestamp struct.
// 实现 UnmarshalJSON 接口
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, _ := strconv.ParseInt(string(b), 10, 64)
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = Timestamp{Carbon: c}
	}
	return c.Error
}

// String implements the interface Stringer for Timestamp struct.
// 实现 Stringer 接口
func (t Timestamp) String() string {
	return strconv.FormatInt(t.Timestamp(), 10)
}

// Int64 outputs timestamp with second.
// 输出秒级时间戳
func (t Timestamp) Int64() int64 {
	return t.Timestamp()
}

// ToTimestampStruct converts Carbon to Timestamp.
// 将 Carbon 结构体转换成 Timestamp 结构体
func (c Carbon) ToTimestampStruct() Timestamp {
	return Timestamp{Carbon: c}
}

// MarshalJSON implements the interface json.Marshal for TimestampMilli struct.
// 实现 MarshalJSON 接口
func (t TimestampMilli) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.TimestampMilli())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for TimestampMilli struct.
// 实现 UnmarshalJSON 接口
func (t *TimestampMilli) UnmarshalJSON(b []byte) error {
	ts, _ := strconv.ParseInt(string(b), 10, 64)
	c := CreateFromTimestampMilli(ts)
	if c.Error == nil {
		*t = TimestampMilli{Carbon: c}
	}
	return c.Error
}

// String implements the interface Stringer for TimestampMilli struct.
// 实现 Stringer 接口
func (t TimestampMilli) String() string {
	return strconv.FormatInt(t.TimestampMilli(), 10)
}

// Int64 outputs timestamp with millisecond.
// 输出豪秒级时间戳
func (t TimestampMilli) Int64() int64 {
	return t.TimestampMilli()
}

// ToTimestampMilliStruct converts Carbon to TimestampMilli.
// 将 Carbon 结构体转换成 TimestampMilli 结构体
func (c Carbon) ToTimestampMilliStruct() TimestampMilli {
	return TimestampMilli{Carbon: c}
}

// MarshalJSON implements the interface MarshalJSON for TimestampMicro struct.
// 实现 MarshalJSON 接口
func (t TimestampMicro) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.TimestampMicro())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for TimestampMicro struct.
// 实现 UnmarshalJSON 接口
func (t *TimestampMicro) UnmarshalJSON(b []byte) error {
	ts, _ := strconv.ParseInt(string(b), 10, 64)
	c := CreateFromTimestampMicro(ts)
	if c.Error == nil {
		*t = TimestampMicro{Carbon: c}
	}
	return c.Error
}

// String implements the interface Stringer for TimestampMicro struct.
// 实现 Stringer 接口
func (t TimestampMicro) String() string {
	return strconv.FormatInt(t.TimestampMicro(), 10)
}

// Int64 outputs timestamp with microsecond.
// 输出微秒级时间戳
func (t TimestampMicro) Int64() int64 {
	return t.TimestampMicro()
}

// ToTimestampMicroStruct converts Carbon to TimestampMicro.
// 将 Carbon 结构体转换成 TimestampMicro 结构体
func (c Carbon) ToTimestampMicroStruct() TimestampMicro {
	return TimestampMicro{Carbon: c}
}

// MarshalJSON implements the interface json.Marshal for TimestampNano struct.
// 实现 MarshalJSON 接口
func (t TimestampNano) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.TimestampNano())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for TimestampNano struct.
// 实现 UnmarshalJSON 接口
func (t *TimestampNano) UnmarshalJSON(b []byte) error {
	ts, _ := strconv.ParseInt(string(b), 10, 64)
	c := CreateFromTimestampNano(ts)
	if c.Error == nil {
		*t = TimestampNano{Carbon: c}
	}
	return c.Error
}

// String implements the interface Stringer for TimestampNano struct.
// 实现 Stringer 接口
func (t TimestampNano) String() string {
	return strconv.FormatInt(t.TimestampNano(), 10)
}

// Int64 outputs timestamp with nanosecond.
// 输出纳秒级时间戳
func (t TimestampNano) Int64() int64 {
	return t.TimestampNano()
}

// ToTimestampNanoStruct converts Carbon to TimestampNano.
// 将 Carbon 结构体转换成 TimestampNano 结构体
func (c Carbon) ToTimestampNanoStruct() TimestampNano {
	return TimestampNano{Carbon: c}
}
