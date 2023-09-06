// The methods in this file will be removed in the future

package carbon

import "time"

// FromStdTime converts standard time.Time to Carbon.
// Deprecated: It will be removed in the future, CreateFromStdTime is recommended.
// 将标准 time.Time 转换成 Carbon，未来将移除，推荐使用 CreateFromStdTime
func FromStdTime(tt time.Time) Carbon {
	return CreateFromStdTime(tt)
}

// Time2Carbon converts standard time.Time to Carbon.
// Deprecated: It will be removed in the future, CreateFromStdTime is recommended.
// 将标准 time.Time 转换成 Carbon，未来将移除，推荐使用 CreateFromStdTime
func Time2Carbon(tt time.Time) Carbon {
	return CreateFromStdTime(tt)
}

// Carbon2Time converts Carbon to standard time.Time.
// Deprecated: It will be removed in the future, ToStdTime is recommended.
// 将 Carbon 转换成标准 time.Time，未来将移除，推荐使用 ToStdTime
func (c Carbon) Carbon2Time() time.Time {
	return c.ToStdTime()
}
