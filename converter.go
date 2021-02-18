package carbon

import "time"

// Time2Carbon 将 time.Time 转换成 Carbon
func Time2Carbon(tt time.Time) Carbon {
	loc, _ := time.LoadLocation(Local)
	return Carbon{Time: tt, Loc: loc}
}

// Carbon2Time 将 Carbon 转换成 time.Time
func (c Carbon) Carbon2Time() time.Time {
	return c.Time
}
