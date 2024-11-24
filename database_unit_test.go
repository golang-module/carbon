package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_Scan(t *testing.T) {
	c := NewCarbon()
	err := c.Scan(time.Now())
	assert.Nil(t, err)
	assert.Equal(t, c.ToDateTimeString(), Now().ToDateTimeString())
}

func TestCarbon_Value(t *testing.T) {
	c := Now()
	v, err := c.Value()
	assert.Nil(t, err)
	assert.Equal(t, v, c.time)
}

func TestError_Scan(t *testing.T) {
	c := NewCarbon()
	str, bytes, time, dt := "xxx", []byte("xxx"), time.Now(), DateTime{}

	err1 := c.Scan(str)
	assert.Equal(t, err1, failedScanError(str))

	err2 := c.Scan(bytes)
	assert.Equal(t, err2, failedScanError(bytes))

	err3 := c.Scan(time)
	assert.Equal(t, err3, nil)

	err4 := c.Scan(dt)
	assert.Equal(t, err4, failedScanError(dt))
}

func TestError_Value(t *testing.T) {
	c := Parse("")
	v, err := c.Value()
	assert.Nil(t, err)
	assert.Equal(t, v, nil)
}

func TestCarbon_GormDataType(t *testing.T) {
	var carbon Carbon
	assert.Equal(t, "time", carbon.GormDataType())

	var dateTime DateTime
	assert.Equal(t, "time", dateTime.GormDataType())
	var dateTimeMilli DateTimeMilli
	assert.Equal(t, "time", dateTimeMilli.GormDataType())
	var dateTimeMicro DateTimeMicro
	assert.Equal(t, "time", dateTimeMicro.GormDataType())
	var dateTimeNano DateTimeNano
	assert.Equal(t, "time", dateTimeNano.GormDataType())

	var date Date
	assert.Equal(t, "time", date.GormDataType())
	var dateMilli DateMilli
	assert.Equal(t, "time", dateMilli.GormDataType())
	var dateMicro DateMicro
	assert.Equal(t, "time", dateMicro.GormDataType())
	var dateNano DateNano
	assert.Equal(t, "time", dateNano.GormDataType())

	var time Time
	assert.Equal(t, "time", time.GormDataType())
	var timeMilli TimeMilli
	assert.Equal(t, "time", timeMilli.GormDataType())
	var timeMicro TimeMicro
	assert.Equal(t, "time", timeMicro.GormDataType())
	var timeNano TimeNano
	assert.Equal(t, "time", timeNano.GormDataType())

	var timestamp Timestamp
	assert.Equal(t, "int", timestamp.GormDataType())
	var timestampMilli TimestampMilli
	assert.Equal(t, "int", timestampMilli.GormDataType())
	var timestampMicro TimestampMicro
	assert.Equal(t, "int", timestampMicro.GormDataType())
	var timestampNano TimestampNano
	assert.Equal(t, "int", timestampNano.GormDataType())
}
