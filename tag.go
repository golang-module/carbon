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

// Tag gets tag.
// 获取标签
func (c Carbon) Tag() string {
	return c.tag
}

// SetTag sets tag.
// 设置标签
func (c Carbon) SetTag(tag string) Carbon {
	if c.Error != nil {
		return c
	}
	c.tag = tag
	return c
}

// parseTag parses tag.
// 解析标签
func (c Carbon) parseTag() (key, value, tz string) {
	if c.tag == "" || len(c.tag) <= 7 {
		return "", "", Local
	}
	slice := strings.Split(strings.Trim(c.tag, ";"), ";")
	key = strings.TrimSpace(slice[0][:6])
	value = strings.TrimSpace(slice[0][7:])
	tz = strings.TrimSpace(slice[1][3:])
	return
}

// LoadTag loads tag.
// 加载标签
func LoadTag(v interface{}) error {
	typ := reflect.TypeOf(v)
	val := reflect.ValueOf(v)
	if typ.Kind() != reflect.Ptr {
		return invalidPtrError()
	}
	params := make([]reflect.Value, 1)
	for i := 0; i < val.Elem().NumField(); i++ {
		field := typ.Elem().Field(i)
		value := val.Elem().Field(i)
		if reflect.TypeOf(Carbon{}) != value.Type() {
			continue
		}

		carbonTag := field.Tag.Get("carbon")
		switch carbonTag {
		case "", "dateTime":
			carbonTag = "layout:" + DateTimeLayout
		case "dateTimeMilli":
			carbonTag = "layout:" + DateTimeMilliLayout
		case "dateTimeMicro":
			carbonTag = "layout:" + DateTimeMicroLayout
		case "dateTimeNano":
			carbonTag = "layout:" + DateTimeNanoLayout
		case "shortDateTime":
			carbonTag = "layout:" + ShortDateTimeLayout
		case "shortDateTimeMilli":
			carbonTag = "layout:" + ShortDateTimeMilliLayout
		case "shortDateTimeMicro":
			carbonTag = "layout:" + ShortDateTimeMicroLayout
		case "shortDateTimeNano":
			carbonTag = "layout:" + ShortDateTimeNanoLayout
		case "dayDateTime":
			carbonTag = "layout:" + DayDateTimeLayout

		case "date":
			carbonTag = "layout:" + DateLayout
		case "dateMilli":
			carbonTag = "layout:" + DateMilliLayout
		case "dateMicro":
			carbonTag = "layout:" + DateMicroLayout
		case "dateNano":
			carbonTag = "layout:" + DateNanoLayout
		case "shortDate":
			carbonTag = "layout:" + ShortDateLayout
		case "shortDateMilli":
			carbonTag = "layout:" + ShortDateMilliLayout
		case "shortDateMicro":
			carbonTag = "layout:" + ShortDateMicroLayout
		case "shortDateNano":
			carbonTag = "layout:" + ShortDateNanoLayout

		case "time":
			carbonTag = "layout:" + TimeLayout
		case "timeMilli":
			carbonTag = "layout:" + TimeMilliLayout
		case "timeMicro":
			carbonTag = "layout:" + TimeMicroLayout
		case "timeNano":
			carbonTag = "layout:" + TimeNanoLayout
		case "shortTime":
			carbonTag = "layout:" + ShortTimeLayout
		case "shortTimeMilli":
			carbonTag = "layout:" + ShortTimeMilliLayout
		case "shortTimeMicro":
			carbonTag = "layout:" + ShortTimeMicroLayout
		case "shortTimeNano":
			carbonTag = "layout:" + ShortTimeNanoLayout

		case "atom":
			carbonTag = "layout:" + AtomLayout
		case "ansic":
			carbonTag = "layout:" + ANSICLayout
		case "cookie":
			carbonTag = "layout:" + CookieLayout
		case "kitchen":
			carbonTag = "layout:" + KitchenLayout
		case "rss":
			carbonTag = "layout:" + RssLayout
		case "rubyDate":
			carbonTag = "layout:" + RubyDateLayout
		case "unixDate":
			carbonTag = "layout:" + UnixDateLayout

		case "rfc1036":
			carbonTag = "layout:" + RFC1036Layout
		case "rfc1123":
			carbonTag = "layout:" + RFC1123Layout
		case "rfc1123Z":
			carbonTag = "layout:" + RFC1123ZLayout
		case "rfc2822":
			carbonTag = "layout:" + RFC2822Layout
		case "rfc3339":
			carbonTag = "layout:" + RFC3339Layout
		case "rfc3339Milli":
			carbonTag = "layout:" + RFC3339MilliLayout
		case "rfc3339Micro":
			carbonTag = "layout:" + RFC3339MicroLayout
		case "rfc3339Nano":
			carbonTag = "layout:" + RFC3339NanoLayout
		case "rfc7231":
			carbonTag = "layout:" + RFC7231Layout
		case "rfc822":
			carbonTag = "layout:" + RFC822Layout
		case "rfc822Z":
			carbonTag = "layout:" + RFC822ZLayout
		case "rfc850":
			carbonTag = "layout:" + RFC850Layout

		case "iso8601":
			carbonTag = "layout:" + ISO8601Layout
		case "iso8601Milli":
			carbonTag = "layout:" + ISO8601MilliLayout
		case "iso8601Micro":
			carbonTag = "layout:" + ISO8601MicroLayout
		case "iso8601Nano":
			carbonTag = "layout:" + ISO8601NanoLayout
		}

		if !strings.Contains(carbonTag, "layout:") && !strings.Contains(carbonTag, "format:") {
			return invalidTagError()
		}

		tzTag := field.Tag.Get("tz")
		if tzTag == "" {
			tzTag = Local
		}

		params[0] = reflect.ValueOf(carbonTag + ";tz:" + tzTag)
		value.Set(value.MethodByName("SetTag").Call(params)[0])
	}
	return nil
}
