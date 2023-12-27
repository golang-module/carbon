package carbon

// SetTestNow sets a test Carbon instance (real or mock) to be returned when a "now" instance is created.
// 设置当前测试时间
func SetTestNow(carbon Carbon) Carbon {
	return NewCarbon().SetTestNow(carbon)
}

// SetTestNow sets a test Carbon instance (real or mock) to be returned when a "now" instance is created.
// 设置当前测试时间
func (c Carbon) SetTestNow(carbon Carbon) Carbon {
	c.testNow = carbon.TimestampNano()
	if carbon.loc != nil {
		c.loc = carbon.loc
	}
	return c
}

// UnSetTestNow clears a test Carbon instance (real or mock) to be returned when a "now" instance is created.
// 清除当前测试时间
func (c Carbon) UnSetTestNow() Carbon {
	c.testNow = 0
	return c
}

// IsSetTestNow report whether there is testing time now.
// 是否设置过当前测试时间
func (c Carbon) IsSetTestNow() bool {
	return c.testNow > 0
}
