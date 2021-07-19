package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_Now(t *testing.T) {
	assert := assert.New(t)

	actual := time.Now().Format(DateFormat)
	expected := Now().ToDateString()
	assert.Equal(expected, actual)
}

func TestCarbon_Yesterday(t *testing.T) {
	assert := assert.New(t)

	c1 := Yesterday()
	actual := time.Now().AddDate(0, 0, -1).Format(DateFormat)
	assert.Nil(c1.Error)
	assert.Equal(c1.ToDateString(), actual)

	c2 := Parse("2020-08-05").Yesterday()
	assert.Nil(c2.Error)
	assert.Equal(c2.ToDateString(), "2020-08-04", "It should be equal to 2020-08-04")
}

func TestCarbon_Tomorrow(t *testing.T) {
	assert := assert.New(t)

	c1 := Tomorrow()
	actual := time.Now().AddDate(0, 0, 1).Format(DateFormat)
	assert.Nil(c1.Error)
	assert.Equal(c1.ToDateString(), actual)

	c2 := Parse("2020-08-05").Tomorrow()
	assert.Nil(c2.Error)
	assert.Equal(c2.ToDateString(), "2020-08-06", "It should be equal to 2020-08-06")
}

func TestCarbon_Time2Carbon(t *testing.T) {
	assert := assert.New(t)

	actual := time.Now().Format(DateTimeFormat)
	expected := Time2Carbon(time.Now()).ToDateTimeString()
	assert.Equal(expected, actual)
}

func TestCarbon_Carbon2Time(t *testing.T) {
	assert := assert.New(t)

	actual := time.Now().Format(DateTimeFormat)
	expected := Now().Carbon2Time().Format(DateTimeFormat)
	assert.Equal(expected, actual)
}
