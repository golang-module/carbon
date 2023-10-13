package carbon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCarbon_SetTestNow(t *testing.T) {
	assert := assert.New(t)

	testNow := Parse("2020-08-05")

	assert.Equal("2020-08-04", SetTestNow(testNow).Yesterday().ToDateString(), "It should be equal to 2020-08-05")
	assert.Equal("2020-08-05", SetTestNow(testNow).Now().ToDateString(), "It should be equal to 2020-08-05")
	assert.Equal("2020-08-06", SetTestNow(testNow).Tomorrow().ToDateString(), "It should be equal to 2020-08-05")
}

func TestCarbon_ClearTestNow(t *testing.T) {
	assert := assert.New(t)

	datetime := "2020-08-05"
	testNow := Parse(datetime)

	assert.Equal(Now().ToDateString(), SetTestNow(testNow).ClearTestNow().Now().ToDateString())
}

func TestCarbon_HasTestNow(t *testing.T) {
	assert := assert.New(t)

	datetime := "2020-08-05"
	testNow := Parse(datetime)

	assert.Equal(false, Now().HasTestNow(), "It should be equal to false")
	assert.Equal(true, SetTestNow(testNow).HasTestNow(), "It should be equal to true")
}
