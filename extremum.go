package carbon

// Closest returns the closest Carbon instance from the given Carbon instance (second-precision).
// 返回离给定 carbon 实例最近的 Carbon 实例
func (c Carbon) Closest(c1 Carbon, c2 Carbon) Carbon {
	if c1.IsZero() || c1.IsInvalid() {
		return c2
	}
	if c2.IsZero() || c2.IsInvalid() {
		return c1
	}
	if c.DiffAbsInSeconds(c1) < c.DiffAbsInSeconds(c2) {
		return c1
	}
	return c2
}

// Farthest returns the farthest Carbon instance from the given Carbon instance (second-precision).
// 返回离给定 carbon 实例最远的 Carbon 实例
func (c Carbon) Farthest(c1 Carbon, c2 Carbon) Carbon {
	if c1.IsZero() || c1.IsInvalid() {
		return c2
	}
	if c2.IsZero() || c2.IsInvalid() {
		return c1
	}
	if c.DiffAbsInSeconds(c1) > c.DiffAbsInSeconds(c2) {
		return c1
	}
	return c2
}
