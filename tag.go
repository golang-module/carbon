package carbon

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	// invalid pointer error
	// 无效的指针错误
	invalidPtrError = func() error {
		return fmt.Errorf("invalid struct pointer, please make sure the struct is a pointer")
	}

	// invalid tag error
	// 无效的标签错误
	invalidTagError = func() error {
		return fmt.Errorf("invalid carbon tag, please make sure the tag is valid")
	}
)

// tag defines a tag struct.
// 定义 tag 结构体
type tag struct {
	carbon string
	tz     string
}

// SetTag sets tag.
// 设置标签
func (c Carbon) SetTag(tag tag) Carbon {
	if c.Error != nil {
		return c
	}
	c.tag = tag
	return c
}

// parseTag parses tag.
// 解析标签
func (c Carbon) parseTag() (key, value, tz string) {
	tz = strings.TrimSpace(c.tag.tz)
	if tz == "" {
		tz = c.Location()
	}
	carbon := strings.TrimSpace(c.tag.carbon)
	if carbon == "" || len(carbon) <= 7 {
		return "", "", tz
	}
	key = strings.TrimSpace(carbon[:6])
	value = strings.TrimSpace(carbon[7:])
	return
}

// LoadTag loads tag.
// 加载标签
func LoadTag(v interface{}) error {
	typeObj, valueObj := reflect.TypeOf(v), reflect.ValueOf(v)
	if typeObj.Kind() != reflect.Ptr {
		return invalidPtrError()
	}
	typeElem, valueElem := typeObj.Elem(), valueObj.Elem()
	params := make([]reflect.Value, 1)
	for i := 0; i < valueElem.NumField(); i++ {
		fieldType, fieldValue := typeElem.Field(i), valueElem.Field(i)
		if reflect.TypeOf(Carbon{}) != fieldValue.Type() {
			continue
		}

		carbon := fieldType.Tag.Get("carbon")
		switch carbon {
		case "", "dateTime":
			carbon = "layout:" + DateTimeLayout
		case "dateTimeMilli":
			carbon = "layout:" + DateTimeMilliLayout
		case "dateTimeMicro":
			carbon = "layout:" + DateTimeMicroLayout
		case "dateTimeNano":
			carbon = "layout:" + DateTimeNanoLayout
		case "shortDateTime":
			carbon = "layout:" + ShortDateTimeLayout
		case "shortDateTimeMilli":
			carbon = "layout:" + ShortDateTimeMilliLayout
		case "shortDateTimeMicro":
			carbon = "layout:" + ShortDateTimeMicroLayout
		case "shortDateTimeNano":
			carbon = "layout:" + ShortDateTimeNanoLayout
		case "dayDateTime":
			carbon = "layout:" + DayDateTimeLayout

		case "date":
			carbon = "layout:" + DateLayout
		case "dateMilli":
			carbon = "layout:" + DateMilliLayout
		case "dateMicro":
			carbon = "layout:" + DateMicroLayout
		case "dateNano":
			carbon = "layout:" + DateNanoLayout
		case "shortDate":
			carbon = "layout:" + ShortDateLayout
		case "shortDateMilli":
			carbon = "layout:" + ShortDateMilliLayout
		case "shortDateMicro":
			carbon = "layout:" + ShortDateMicroLayout
		case "shortDateNano":
			carbon = "layout:" + ShortDateNanoLayout

		case "time":
			carbon = "layout:" + TimeLayout
		case "timeMilli":
			carbon = "layout:" + TimeMilliLayout
		case "timeMicro":
			carbon = "layout:" + TimeMicroLayout
		case "timeNano":
			carbon = "layout:" + TimeNanoLayout
		case "shortTime":
			carbon = "layout:" + ShortTimeLayout
		case "shortTimeMilli":
			carbon = "layout:" + ShortTimeMilliLayout
		case "shortTimeMicro":
			carbon = "layout:" + ShortTimeMicroLayout
		case "shortTimeNano":
			carbon = "layout:" + ShortTimeNanoLayout

		case "atom":
			carbon = "layout:" + AtomLayout
		case "ansic":
			carbon = "layout:" + ANSICLayout
		case "cookie":
			carbon = "layout:" + CookieLayout
		case "kitchen":
			carbon = "layout:" + KitchenLayout
		case "rss":
			carbon = "layout:" + RssLayout
		case "rubyDate":
			carbon = "layout:" + RubyDateLayout
		case "unixDate":
			carbon = "layout:" + UnixDateLayout

		case "rfc1036":
			carbon = "layout:" + RFC1036Layout
		case "rfc1123":
			carbon = "layout:" + RFC1123Layout
		case "rfc1123Z":
			carbon = "layout:" + RFC1123ZLayout
		case "rfc2822":
			carbon = "layout:" + RFC2822Layout
		case "rfc3339":
			carbon = "layout:" + RFC3339Layout
		case "rfc3339Milli":
			carbon = "layout:" + RFC3339MilliLayout
		case "rfc3339Micro":
			carbon = "layout:" + RFC3339MicroLayout
		case "rfc3339Nano":
			carbon = "layout:" + RFC3339NanoLayout
		case "rfc7231":
			carbon = "layout:" + RFC7231Layout
		case "rfc822":
			carbon = "layout:" + RFC822Layout
		case "rfc822Z":
			carbon = "layout:" + RFC822ZLayout
		case "rfc850":
			carbon = "layout:" + RFC850Layout

		case "iso8601":
			carbon = "layout:" + ISO8601Layout
		case "iso8601Milli":
			carbon = "layout:" + ISO8601MilliLayout
		case "iso8601Micro":
			carbon = "layout:" + ISO8601MicroLayout
		case "iso8601Nano":
			carbon = "layout:" + ISO8601NanoLayout

		case "timestamp":
			carbon = "format:U"
		case "timestampMilli":
			carbon = "format:V"
		case "timestampMicro":
			carbon = "format:X"
		case "timestampNano":
			carbon = "format:Z"
		}

		if !strings.Contains(carbon, "layout:") && !strings.Contains(carbon, "format:") {
			return invalidTagError()
		}

		tz := fieldType.Tag.Get("tz")
		if tz == "" {
			tz = Local
		}

		params[0] = reflect.ValueOf(tag{
			carbon: carbon,
			tz:     tz,
		})
		fieldValue.Set(fieldValue.MethodByName("SetTag").Call(params)[0])
	}
	return nil
}
