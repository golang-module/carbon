package carbon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCarbon_SetTestNow(t *testing.T) {
	assert := assert.New(t)

	testNow := Parse("2020-08-05")

	assert.Equal("2020-08-04", SetTestNow(testNow).Yesterday().ToDateString(), "It should be equal to 2020-08-04")
	assert.Equal("2020-08-05", SetTestNow(testNow).Now().ToDateString(), "It should be equal to 2020-08-05")
	assert.Equal("2020-08-06", SetTestNow(testNow).Tomorrow().ToDateString(), "It should be equal to 2020-08-06")
	assert.Equal(31, SetTestNow(testNow).Parse("1989-08-05").Age(), "It should be equal to 31")
	assert.Equal(int64(1), SetTestNow(testNow).Parse("2020-07-05").DiffInMonths(), "It should be equal to 1")
	assert.Equal(int64(4), SetTestNow(testNow).Parse("2020-07-05").DiffInWeeks(), "It should be equal to 4")
	assert.Equal(int64(31), SetTestNow(testNow).Parse("2020-07-05").DiffInDays(), "It should be equal to 31")
	assert.Equal(int64(-13), SetTestNow(testNow).Parse("2020-08-05 13:14:15").DiffInHours(), "It should be equal to -13")
	assert.Equal(int64(-794), SetTestNow(testNow).Parse("2020-08-05 13:14:15").DiffInMinutes(), "It should be equal to -794")
	assert.Equal(int64(-47655), SetTestNow(testNow).Parse("2020-08-05 13:14:15").DiffInSeconds(), "It should be equal to -47655")
	assert.Equal("-13 hours", SetTestNow(testNow).Parse("2020-08-05 13:14:15").DiffInString(), "It should be -13 hours")
	assert.Equal("13 hours", SetTestNow(testNow).Parse("2020-08-05 13:14:15").DiffAbsInString(), "It should be -13 hours")
	assert.Equal("13 hours from now", SetTestNow(testNow).Parse("2020-08-05 13:14:15").DiffForHumans(), "It should be 13 hours from now")
}

func TestCarbon_UnSetTestNow(t *testing.T) {
	assert := assert.New(t)

	datetime := "2020-08-05"
	testNow := Parse(datetime)

	assert.Equal(Now().ToDateString(), SetTestNow(testNow).UnSetTestNow().Now().ToDateString())
}

func TestCarbon_IsSetTestNow(t *testing.T) {
	assert := assert.New(t)

	datetime := "2020-08-05"
	testNow := Parse(datetime)

	assert.Equal(false, Now().IsSetTestNow(), "It should be equal to false")
	assert.Equal(true, SetTestNow(testNow).IsSetTestNow(), "It should be equal to true")
}
