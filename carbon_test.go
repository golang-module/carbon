package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_CreateFromStdTime(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt := time.Now().In(loc)
	expected := tt.Format(DateTimeLayout)
	actual := CreateFromStdTime(tt).ToDateTimeString()
	assert.Equal(t, expected, actual)
}

func TestCarbon_FromStdTime(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt := time.Now().In(loc)
	expected := tt.Format(DateTimeLayout)
	actual := FromStdTime(tt).ToDateTimeString()
	assert.Equal(t, expected, actual)
}

func TestCarbon_ToStdTime(t *testing.T) {
	expected := time.Now().Format(DateTimeLayout)
	actual := Now().ToStdTime().Format(DateTimeLayout)
	assert.Equal(t, expected, actual)
}

func TestCarbon_Time2Carbon(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt := time.Now().In(loc)
	expected := tt.Format(DateTimeLayout)
	actual := Time2Carbon(tt).ToDateTimeString()
	assert.Equal(t, expected, actual)
}

func TestCarbon_Carbon2Time(t *testing.T) {
	expected := time.Now().Format(DateTimeLayout)
	actual := Now().Carbon2Time().Format(DateTimeLayout)
	assert.Equal(t, expected, actual)
}

func TestError_Carbon(t *testing.T) {
	timezone := "xxx"
	assert.NotNil(t, Now(timezone).Error, "It should catch an exception in Now()")
	assert.NotNil(t, Tomorrow(timezone).Error, "It should catch an exception in Tomorrow()")
	assert.NotNil(t, Yesterday(timezone).Error, "It should catch an exception in Yesterday()")
}
