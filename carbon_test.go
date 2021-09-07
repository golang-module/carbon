package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_Now(t *testing.T) {
	assert := assert.New(t)

	actual1 := Now().ToDateString()
	expected1 := time.Now().Format(DateFormat)
	assert.Equal(expected1, actual1)

	actual2 := Now(Local).ToDateString()
	expected2 := time.Now().In(time.Local).Format(DateFormat)
	assert.Equal(expected2, actual2)
}

func TestError_Now(t *testing.T) {
	timezone := "xxx"
	c := Now(timezone)
	assert.NotNil(t, c.Error, "It should catch an exception in Now()")
}

func TestCarbon_Yesterday(t *testing.T) {
	assert := assert.New(t)

	c1 := Yesterday()
	expected1 := time.Now().AddDate(0, 0, -1).Format(DateFormat)
	assert.Nil(c1.Error)
	assert.Equal(expected1, c1.ToDateString())

	c2 := Yesterday(Local)
	expected2 := time.Now().In(time.Local).AddDate(0, 0, -1).Format(DateFormat)
	assert.Nil(c2.Error)
	assert.Equal(expected2, c2.ToDateString())

	c3 := Parse("2020-08-05").Yesterday()
	assert.Nil(c3.Error)
	assert.Equal("2020-08-04", c3.ToDateString(), "It should be equal to 2020-08-04")
}

func TestError_Yesterday(t *testing.T) {
	timezone := "xxx"
	c := Yesterday(timezone)
	assert.NotNil(t, c.Error, "It should catch an exception in Yesterday()")
}

func TestCarbon_Tomorrow(t *testing.T) {
	assert := assert.New(t)

	c1 := Tomorrow()
	expected1 := time.Now().AddDate(0, 0, 1).Format(DateFormat)
	assert.Nil(c1.Error)
	assert.Equal(expected1, c1.ToDateString())

	c2 := Tomorrow(Local)
	expected2 := time.Now().In(time.Local).AddDate(0, 0, 1).Format(DateFormat)
	assert.Nil(c2.Error)
	assert.Equal(expected2, c2.ToDateString())

	c3 := Parse("2020-08-05").Tomorrow()
	assert.Nil(c3.Error)
	assert.Equal("2020-08-06", c3.ToDateString(), "It should be equal to 2020-08-06")
}

func TestError_Tomorrow(t *testing.T) {
	timezone := "xxx"
	c := Tomorrow(timezone)
	assert.NotNil(t, c.Error, "It should catch an exception in Tomorrow()")
}

func TestCarbon_Time2Carbon(t *testing.T) {
	assert := assert.New(t)

	expected := time.Now().Format(DateTimeFormat)
	actual := Time2Carbon(time.Now()).ToDateTimeString()
	assert.Equal(expected, actual)
}

func TestCarbon_Carbon2Time(t *testing.T) {
	assert := assert.New(t)

	expected := time.Now().Format(DateTimeFormat)
	actual := Now().Carbon2Time().Format(DateTimeFormat)
	assert.Equal(expected, actual)
}
