package carbon_test

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dromara/carbon/v2"
)

func TestLayoutType_Scan(t *testing.T) {
	c := carbon.NewLayoutType[carbon.DateTime](carbon.Now())

	t.Run("[]byte type", func(t *testing.T) {
		e := c.Scan([]byte(carbon.Now().ToDateString()))
		assert.Nil(t, e)
	})

	t.Run("string type", func(t *testing.T) {
		e := c.Scan(carbon.Now().ToDateString())
		assert.Nil(t, e)
	})

	t.Run("int64 type", func(t *testing.T) {
		e := c.Scan(carbon.Now().Timestamp())
		assert.Nil(t, e)
	})

	t.Run("time type", func(t *testing.T) {
		e := c.Scan(carbon.Now().StdTime())
		assert.Nil(t, e)
	})

	t.Run("unsupported type", func(t *testing.T) {
		e := c.Scan(nil)
		assert.Error(t, e)
	})
}

func TestLayoutType_Value(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		v, err := carbon.NewLayoutType[carbon.DateTime](nil).Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("zero time", func(t *testing.T) {
		c := carbon.NewCarbon()
		v, err := carbon.NewLayoutType[carbon.DateTime](c).Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("invalid time", func(t *testing.T) {
		c := carbon.Parse("xxx")
		v, err := carbon.NewLayoutType[carbon.DateTime](c).Value()
		assert.Nil(t, v)
		assert.Error(t, err)
	})

	t.Run("valid time", func(t *testing.T) {
		c := carbon.Parse("2020-08-05")
		v, err := carbon.NewLayoutType[carbon.DateTime](c).Value()
		assert.Equal(t, c.StdTime(), v)
		assert.Nil(t, err)
	})
}

func TestLayoutType_MarshalJSON(t *testing.T) {
	type User struct {
		Date     carbon.LayoutType[carbon.Date]     `json:"date"`
		Time     carbon.LayoutType[carbon.Time]     `json:"time"`
		DateTime carbon.LayoutType[carbon.DateTime] `json:"date_time"`
	}

	t.Run("nil time", func(t *testing.T) {
		var user User

		user.Date = carbon.NewLayoutType[carbon.Date](nil)
		user.Time = carbon.NewLayoutType[carbon.Time](nil)
		user.DateTime = carbon.NewLayoutType[carbon.DateTime](nil)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":"","time":"","date_time":""}`, string(data))
	})

	t.Run("zero time", func(t *testing.T) {
		var user User

		c := carbon.NewCarbon()

		user.Date = carbon.NewLayoutType[carbon.Date](c)
		user.Time = carbon.NewLayoutType[carbon.Time](c)
		user.DateTime = carbon.NewLayoutType[carbon.DateTime](c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":"","time":"","date_time":""}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("xxx")

		user.Date = carbon.NewLayoutType[carbon.Date](c)
		user.Time = carbon.NewLayoutType[carbon.Time](c)
		user.DateTime = carbon.NewLayoutType[carbon.DateTime](c)
		data, err := json.Marshal(&user)
		assert.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("valid time", func(t *testing.T) {
		var user User

		now := carbon.Parse("2020-08-05 13:14:15")

		user.Date = carbon.NewLayoutType[carbon.Date](now)
		user.Time = carbon.NewLayoutType[carbon.Time](now.AddDay())
		user.DateTime = carbon.NewLayoutType[carbon.DateTime](now.SubDay())

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":"2020-08-05","time":"13:14:15","date_time":"2020-08-05 13:14:15"}`, string(data))
	})
}

func TestLayoutType_UnmarshalJSON(t *testing.T) {
	type User struct {
		Date     carbon.LayoutType[carbon.Date]     `json:"date"`
		Time     carbon.LayoutType[carbon.Time]     `json:"time"`
		DateTime carbon.LayoutType[carbon.DateTime] `json:"date_time"`
	}

	t.Run("empty value", func(t *testing.T) {
		var user User

		value := `{"date":"", "time":"", "date_time": ""}`
		err := json.Unmarshal([]byte(value), &user)
		assert.NoError(t, err)
		assert.Empty(t, user.Date.String())
		assert.Empty(t, user.Time.String())
		assert.Empty(t, user.DateTime.String())
	})

	t.Run("null value", func(t *testing.T) {
		var user User

		value := `{"date":"null", "time":"null", "date_time": "null"}`
		err := json.Unmarshal([]byte(value), &user)
		assert.NoError(t, err)
		assert.Empty(t, user.Date.String())
		assert.Empty(t, user.Time.String())
		assert.Empty(t, user.DateTime.String())
	})

	t.Run("zero value", func(t *testing.T) {
		var user User

		value := `{"date":"0", "time":"0", "date_time": "0"}`
		err := json.Unmarshal([]byte(value), &user)
		assert.NoError(t, err)
		assert.Empty(t, user.Date.String())
		assert.Empty(t, user.Time.String())
		assert.Empty(t, user.DateTime.String())
	})

	t.Run("valid value", func(t *testing.T) {
		var user User

		value := `{"date":"2020-08-05","time":"13:14:15","date_time":"2020-08-05 13:14:15"}`
		err := json.Unmarshal([]byte(value), &user)
		assert.NoError(t, err)
		assert.Equal(t, "2020-08-05", user.Date.String())
		assert.Equal(t, "13:14:15", user.Time.String())
		assert.Equal(t, "2020-08-05 13:14:15", user.DateTime.String())
	})
}

func TestLayoutType_GormDataType(t *testing.T) {
	c := carbon.Now()
	dataType := "carbonLayout"

	c1 := carbon.NewLayoutType[carbon.Date](c)
	assert.Equal(t, dataType, c1.GormDataType())

	c2 := carbon.NewLayoutType[carbon.Time](c)
	assert.Equal(t, dataType, c2.GormDataType())

	c3 := carbon.NewLayoutType[carbon.DateTime](c)
	assert.Equal(t, dataType, c3.GormDataType())
}

func TestFormatType_Scan(t *testing.T) {
	c := carbon.NewFormatType[carbon.DateTime](carbon.Now())

	t.Run("[]byte type", func(t *testing.T) {
		e := c.Scan([]byte(carbon.Now().ToDateString()))
		assert.Nil(t, e)
	})

	t.Run("string type", func(t *testing.T) {
		e := c.Scan(carbon.Now().ToDateString())
		assert.Nil(t, e)
	})

	t.Run("int64 type", func(t *testing.T) {
		e := c.Scan(carbon.Now().Timestamp())
		assert.Nil(t, e)
	})

	t.Run("time type", func(t *testing.T) {
		e := c.Scan(carbon.Now().StdTime())
		assert.Nil(t, e)
	})

	t.Run("unsupported type", func(t *testing.T) {
		e := c.Scan(nil)
		assert.Error(t, e)
	})
}

func TestFormatType_Value(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		v, err := carbon.NewFormatType[carbon.DateTime](nil).Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("zero time", func(t *testing.T) {
		c := carbon.NewCarbon()
		v, err := carbon.NewFormatType[carbon.DateTime](c).Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("invalid time", func(t *testing.T) {
		c := carbon.Parse("xxx")
		v, err := carbon.NewFormatType[carbon.DateTime](c).Value()
		assert.Nil(t, v)
		assert.Error(t, err)
	})

	t.Run("valid time", func(t *testing.T) {
		c := carbon.Parse("2020-08-05")
		v, err := carbon.NewFormatType[carbon.DateTime](c).Value()
		assert.Equal(t, c.StdTime(), v)
		assert.Nil(t, err)
	})
}

func TestFormatType_MarshalJSON(t *testing.T) {
	type User struct {
		Date     carbon.FormatType[carbon.Date]     `json:"date"`
		Time     carbon.FormatType[carbon.Time]     `json:"time"`
		DateTime carbon.FormatType[carbon.DateTime] `json:"date_time"`
	}

	t.Run("nil time", func(t *testing.T) {
		var user User

		user.Date = carbon.NewFormatType[carbon.Date](nil)
		user.Time = carbon.NewFormatType[carbon.Time](nil)
		user.DateTime = carbon.NewFormatType[carbon.DateTime](nil)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":"","time":"","date_time":""}`, string(data))
	})

	t.Run("zero time", func(t *testing.T) {
		var user User

		user.Date = carbon.NewFormatType[carbon.Date](carbon.NewCarbon())
		user.Time = carbon.NewFormatType[carbon.Time](carbon.NewCarbon())
		user.DateTime = carbon.NewFormatType[carbon.DateTime](carbon.NewCarbon())

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":"","time":"","date_time":""}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("xxx")

		user.Date = carbon.NewFormatType[carbon.Date](c)
		user.Time = carbon.NewFormatType[carbon.Time](c)
		user.DateTime = carbon.NewFormatType[carbon.DateTime](c)
		data, err := json.Marshal(&user)
		assert.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("valid time", func(t *testing.T) {
		var user User

		now := carbon.Parse("2020-08-05 13:14:15")

		user.Date = carbon.NewFormatType[carbon.Date](now)
		user.Time = carbon.NewFormatType[carbon.Time](now.AddDay())
		user.DateTime = carbon.NewFormatType[carbon.DateTime](now.SubDay())

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":"2020-08-05","time":"13:14:15","date_time":"2020-08-05 13:14:15"}`, string(data))
	})
}

func TestFormatType_UnmarshalJSON(t *testing.T) {
	type User struct {
		Date     carbon.FormatType[carbon.Date]     `json:"date"`
		Time     carbon.FormatType[carbon.Time]     `json:"time"`
		DateTime carbon.FormatType[carbon.DateTime] `json:"date_time"`
	}
	t.Run("empty value", func(t *testing.T) {
		var user User

		value := `{"date":"", "time":"", "date_time": ""}`
		err := json.Unmarshal([]byte(value), &user)
		assert.NoError(t, err)
		assert.Empty(t, user.Date.String())
		assert.Empty(t, user.Time.String())
		assert.Empty(t, user.DateTime.String())
	})

	t.Run("null value", func(t *testing.T) {
		var user User

		value := `{"date":"null", "time":"null", "date_time": "null"}`
		err := json.Unmarshal([]byte(value), &user)
		assert.NoError(t, err)
		assert.Empty(t, user.Date.String())
		assert.Empty(t, user.Time.String())
		assert.Empty(t, user.DateTime.String())
	})

	t.Run("zero value", func(t *testing.T) {
		var user User

		value := `{"date":"0", "time":"0", "date_time": "0"}`
		err := json.Unmarshal([]byte(value), &user)
		assert.NoError(t, err)
		assert.Empty(t, user.Date.String())
		assert.Empty(t, user.Time.String())
		assert.Empty(t, user.DateTime.String())
	})

	t.Run("valid value", func(t *testing.T) {
		var user User

		value := `{"date":"2020-08-05","time":"13:14:15","date_time":"2020-08-05 13:14:15"}`
		err := json.Unmarshal([]byte(value), &user)
		assert.NoError(t, err)
		assert.Equal(t, "2020-08-05", user.Date.String())
		assert.Equal(t, "13:14:15", user.Time.String())
		assert.Equal(t, "2020-08-05 13:14:15", user.DateTime.String())
	})
}

func TestFormatType_GormDataType(t *testing.T) {
	c := carbon.Now()
	dataType := "carbonFormat"

	c1 := carbon.NewFormatType[carbon.Date](c)
	assert.Equal(t, dataType, c1.GormDataType())

	c2 := carbon.NewFormatType[carbon.Time](c)
	assert.Equal(t, dataType, c2.GormDataType())

	c3 := carbon.NewFormatType[carbon.DateTime](c)
	assert.Equal(t, dataType, c3.GormDataType())
}

func TestTimestampType_Scan(t *testing.T) {
	t.Run("[]byte type", func(t *testing.T) {
		c1 := carbon.NewTimestampType[carbon.Timestamp](carbon.Now())
		e1 := c1.Scan([]byte(strconv.Itoa(int(c1.Timestamp()))))
		assert.Nil(t, e1)

		c2 := carbon.NewTimestampType[carbon.TimestampMilli](carbon.Now())
		e2 := c2.Scan([]byte(strconv.Itoa(int(c2.Timestamp()))))
		assert.Nil(t, e2)

		c3 := carbon.NewTimestampType[carbon.TimestampMicro](carbon.Now())
		e3 := c3.Scan([]byte(strconv.Itoa(int(c3.Timestamp()))))
		assert.Nil(t, e3)

		c4 := carbon.NewTimestampType[carbon.TimestampNano](carbon.Now())
		e4 := c4.Scan([]byte(strconv.Itoa(int(c4.Timestamp()))))
		assert.Nil(t, e4)

		e5 := c4.Scan([]byte("xxx"))
		assert.Error(t, e5)
	})

	t.Run("string type", func(t *testing.T) {
		c1 := carbon.NewTimestampType[carbon.Timestamp](carbon.Now())
		e1 := c1.Scan(strconv.Itoa(int(c1.Timestamp())))
		assert.Nil(t, e1)

		c2 := carbon.NewTimestampType[carbon.TimestampMilli](carbon.Now())
		e2 := c2.Scan(strconv.Itoa(int(c2.Timestamp())))
		assert.Nil(t, e2)

		c3 := carbon.NewTimestampType[carbon.TimestampMicro](carbon.Now())
		e3 := c3.Scan(strconv.Itoa(int(c3.Timestamp())))
		assert.Nil(t, e3)

		c4 := carbon.NewTimestampType[carbon.TimestampNano](carbon.Now())
		e4 := c4.Scan(strconv.Itoa(int(c4.Timestamp())))
		assert.Nil(t, e4)

		e5 := c4.Scan("xxx")
		assert.Error(t, e5)
	})

	t.Run("int64 type", func(t *testing.T) {
		c1 := carbon.NewTimestampType[carbon.Timestamp](carbon.Now())
		e1 := c1.Scan(carbon.Now().Timestamp())
		assert.Nil(t, e1)

		c2 := carbon.NewTimestampType[carbon.TimestampMilli](carbon.Now())
		e2 := c2.Scan(carbon.Now().TimestampMilli())
		assert.Nil(t, e2)

		c3 := carbon.NewTimestampType[carbon.TimestampMicro](carbon.Now())
		e3 := c3.Scan(carbon.Now().TimestampMicro())
		assert.Nil(t, e3)

		c4 := carbon.NewTimestampType[carbon.TimestampNano](carbon.Now())
		e4 := c4.Scan(carbon.Now().TimestampNano())
		assert.Nil(t, e4)
	})

	t.Run("time type", func(t *testing.T) {
		c1 := carbon.NewTimestampType[carbon.Timestamp](carbon.Now())
		e1 := c1.Scan(carbon.Now().StdTime())
		assert.Nil(t, e1)

		c2 := carbon.NewTimestampType[carbon.TimestampMilli](carbon.Now())
		e2 := c2.Scan(carbon.Now().StdTime())
		assert.Nil(t, e2)

		c3 := carbon.NewTimestampType[carbon.TimestampMicro](carbon.Now())
		e3 := c3.Scan(carbon.Now().StdTime())
		assert.Nil(t, e3)

		c4 := carbon.NewTimestampType[carbon.TimestampNano](carbon.Now())
		e4 := c4.Scan(carbon.Now().StdTime())
		assert.Nil(t, e4)
	})

	t.Run("unsupported type", func(t *testing.T) {
		c := carbon.Now()

		c1 := carbon.NewTimestampType[carbon.Timestamp](c)
		e1 := c1.Scan(nil)
		assert.Error(t, e1)

		c2 := carbon.NewTimestampType[carbon.TimestampMilli](c)
		e2 := c2.Scan(nil)
		assert.Error(t, e2)

		c3 := carbon.NewTimestampType[carbon.TimestampMicro](c)
		e3 := c3.Scan(nil)
		assert.Error(t, e3)

		c4 := carbon.NewTimestampType[carbon.TimestampNano](c)
		e4 := c4.Scan(nil)
		assert.Error(t, e4)
	})
}

func TestTimestampType_Value(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		v1, e1 := carbon.NewTimestampType[carbon.Timestamp](nil).Value()
		assert.Nil(t, v1)
		assert.Nil(t, e1)

		v2, e2 := carbon.NewTimestampType[carbon.TimestampMilli](nil).Value()
		assert.Nil(t, v2)
		assert.Nil(t, e2)

		v3, e3 := carbon.NewTimestampType[carbon.TimestampMicro](nil).Value()
		assert.Nil(t, v3)
		assert.Nil(t, e3)

		v4, e4 := carbon.NewTimestampType[carbon.TimestampNano](nil).Value()
		assert.Nil(t, v4)
		assert.Nil(t, e4)
	})

	t.Run("zero time", func(t *testing.T) {
		c := carbon.NewCarbon()

		v1, e1 := carbon.NewTimestampType[carbon.Timestamp](c).Value()
		assert.Nil(t, v1)
		assert.Nil(t, e1)

		v2, e2 := carbon.NewTimestampType[carbon.TimestampMilli](c).Value()
		assert.Nil(t, v2)
		assert.Nil(t, e2)

		v3, e3 := carbon.NewTimestampType[carbon.TimestampMicro](c).Value()
		assert.Nil(t, v3)
		assert.Nil(t, e3)

		v4, e4 := carbon.NewTimestampType[carbon.TimestampNano](c).Value()
		assert.Nil(t, v4)
		assert.Nil(t, e4)
	})

	t.Run("invalid time", func(t *testing.T) {
		c := carbon.Parse("xxx")

		v1, e1 := carbon.NewTimestampType[carbon.Timestamp](c).Value()
		assert.Nil(t, v1)
		assert.Error(t, e1)

		v2, e2 := carbon.NewTimestampType[carbon.TimestampMilli](c).Value()
		assert.Nil(t, v2)
		assert.Error(t, e2)

		v3, e3 := carbon.NewTimestampType[carbon.TimestampMicro](c).Value()
		assert.Nil(t, v3)
		assert.Error(t, e3)

		v4, e4 := carbon.NewTimestampType[carbon.TimestampNano](c).Value()
		assert.Nil(t, v4)
		assert.Error(t, e4)
	})

	t.Run("valid time", func(t *testing.T) {
		c := carbon.Parse("2020-08-05")

		v1, e1 := carbon.NewTimestampType[carbon.Timestamp](c).Value()
		assert.Equal(t, c.Timestamp(), v1)
		assert.Nil(t, e1)

		v2, e2 := carbon.NewTimestampType[carbon.TimestampMilli](c).Value()
		assert.Equal(t, c.TimestampMilli(), v2)
		assert.Nil(t, e2)

		v3, e3 := carbon.NewTimestampType[carbon.TimestampMicro](c).Value()
		assert.Equal(t, c.TimestampMicro(), v3)
		assert.Nil(t, e3)

		v4, e4 := carbon.NewTimestampType[carbon.TimestampNano](c).Value()
		assert.Equal(t, c.TimestampNano(), v4)
		assert.Nil(t, e4)
	})
}

func TestTimestamp_MarshalJSON(t *testing.T) {
	type User struct {
		Timestamp      carbon.TimestampType[carbon.Timestamp]      `json:"timestamp"`
		TimestampMilli carbon.TimestampType[carbon.TimestampMilli] `json:"timestamp_milli"`
		TimestampMicro carbon.TimestampType[carbon.TimestampMicro] `json:"timestamp_micro"`
		TimestampNano  carbon.TimestampType[carbon.TimestampNano]  `json:"timestamp_nano"`
	}

	t.Run("nil time", func(t *testing.T) {
		var user User

		user.Timestamp = carbon.NewTimestampType[carbon.Timestamp](nil)
		user.TimestampMilli = carbon.NewTimestampType[carbon.TimestampMilli](nil)
		user.TimestampMicro = carbon.NewTimestampType[carbon.TimestampMicro](nil)
		user.TimestampNano = carbon.NewTimestampType[carbon.TimestampNano](nil)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"timestamp":0,"timestamp_milli":0,"timestamp_micro":0,"timestamp_nano":0}`, string(data))
	})

	t.Run("zero time", func(t *testing.T) {
		var user User

		c := carbon.NewCarbon()

		user.Timestamp = carbon.NewTimestampType[carbon.Timestamp](c)
		user.TimestampMilli = carbon.NewTimestampType[carbon.TimestampMilli](c)
		user.TimestampMicro = carbon.NewTimestampType[carbon.TimestampMicro](c)
		user.TimestampNano = carbon.NewTimestampType[carbon.TimestampNano](c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"timestamp":0,"timestamp_milli":0,"timestamp_micro":0,"timestamp_nano":0}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("xxx")

		user.Timestamp = carbon.NewTimestampType[carbon.Timestamp](c)
		user.TimestampMilli = carbon.NewTimestampType[carbon.TimestampMilli](c)
		user.TimestampMicro = carbon.NewTimestampType[carbon.TimestampMicro](c)
		user.TimestampNano = carbon.NewTimestampType[carbon.TimestampNano](c)

		data, err := json.Marshal(&user)
		assert.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("valid time", func(t *testing.T) {
		var user User

		now := carbon.Parse("2020-08-05 13:14:15")

		user.Timestamp = carbon.NewTimestampType[carbon.Timestamp](now)
		user.TimestampMilli = carbon.NewTimestampType[carbon.TimestampMilli](now.AddDay())
		user.TimestampMicro = carbon.NewTimestampType[carbon.TimestampMicro](now.SubDay())
		user.TimestampNano = carbon.NewTimestampType[carbon.TimestampNano](now.EndOfDay())

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"timestamp":1596633255,"timestamp_milli":1596633255000,"timestamp_micro":1596633255000000,"timestamp_nano":1596671999999999999}`, string(data))
	})
}

func TestTimestamp_UnmarshalJSON(t *testing.T) {
	type User struct {
		Timestamp      carbon.TimestampType[carbon.Timestamp]      `json:"timestamp"`
		TimestampMilli carbon.TimestampType[carbon.TimestampMilli] `json:"timestamp_milli"`
		TimestampMicro carbon.TimestampType[carbon.TimestampMicro] `json:"timestamp_micro"`
		TimestampNano  carbon.TimestampType[carbon.TimestampNano]  `json:"timestamp_nano"`
	}

	t.Run("empty value", func(t *testing.T) {
		var user User

		value := `{"timestamp":0,"timestamp_milli":0,"timestamp_micro":0,"timestamp_nano":0}`
		err := json.Unmarshal([]byte(value), &user)

		assert.NoError(t, err)
		assert.Equal(t, "0", user.Timestamp.String())
		assert.Equal(t, "0", user.TimestampMilli.String())
		assert.Equal(t, "0", user.TimestampMicro.String())
		assert.Equal(t, "0", user.TimestampNano.String())

		assert.Zero(t, user.Timestamp.Int64())
		assert.Zero(t, user.TimestampMilli.Int64())
		assert.Zero(t, user.TimestampMicro.Int64())
		assert.Zero(t, user.TimestampNano.Int64())
	})

	t.Run("null value", func(t *testing.T) {
		var user User

		value := `{"timestamp":0,"timestamp_milli":0,"timestamp_micro":0,"timestamp_nano":0}`
		err := json.Unmarshal([]byte(value), &user)

		assert.NoError(t, err)
		assert.Equal(t, "0", user.Timestamp.String())
		assert.Equal(t, "0", user.TimestampMilli.String())
		assert.Equal(t, "0", user.TimestampMicro.String())
		assert.Equal(t, "0", user.TimestampNano.String())

		assert.Zero(t, user.Timestamp.Int64())
		assert.Zero(t, user.TimestampMilli.Int64())
		assert.Zero(t, user.TimestampMicro.Int64())
		assert.Zero(t, user.TimestampNano.Int64())
	})

	t.Run("zero value", func(t *testing.T) {
		var user User

		value := `{"timestamp":0,"timestamp_milli":0,"timestamp_micro":0,"timestamp_nano":0}`
		err := json.Unmarshal([]byte(value), &user)

		assert.NoError(t, err)
		assert.Equal(t, "0", user.Timestamp.String())
		assert.Equal(t, "0", user.TimestampMilli.String())
		assert.Equal(t, "0", user.TimestampMicro.String())
		assert.Equal(t, "0", user.TimestampNano.String())

		assert.Zero(t, user.Timestamp.Int64())
		assert.Zero(t, user.TimestampMilli.Int64())
		assert.Zero(t, user.TimestampMicro.Int64())
		assert.Zero(t, user.TimestampNano.Int64())
	})

	t.Run("valid value", func(t *testing.T) {
		var user User

		value := `{"timestamp":1596633255,"timestamp_milli":1596633255000,"timestamp_micro":1596633255000000,"timestamp_nano":1596671999999999999}`
		err := json.Unmarshal([]byte(value), &user)

		assert.NoError(t, err)
		assert.Equal(t, "1596633255", user.Timestamp.String())
		assert.Equal(t, "1596633255000", user.TimestampMilli.String())
		assert.Equal(t, "1596633255000000", user.TimestampMicro.String())
		assert.Equal(t, "1596671999999999999", user.TimestampNano.String())

		assert.Equal(t, int64(1596633255), user.Timestamp.Int64())
		assert.Equal(t, int64(1596633255000), user.TimestampMilli.Int64())
		assert.Equal(t, int64(1596633255000000), user.TimestampMicro.Int64())
		assert.Equal(t, int64(1596671999999999999), user.TimestampNano.Int64())
	})
}

func TestTimestamp_GormDataType(t *testing.T) {
	c := carbon.Now()
	dataType := "carbonTimestamp"

	c1 := carbon.NewTimestampType[carbon.Timestamp](c)
	assert.Equal(t, dataType, c1.GormDataType())

	c2 := carbon.NewTimestampType[carbon.TimestampMilli](c)
	assert.Equal(t, dataType, c2.GormDataType())

	c3 := carbon.NewTimestampType[carbon.TimestampMicro](c)
	assert.Equal(t, dataType, c3.GormDataType())

	c4 := carbon.NewTimestampType[carbon.TimestampNano](c)
	assert.Equal(t, dataType, c4.GormDataType())
}

type RFC3339Layout struct{}

func (t RFC3339Layout) SetLayout() string {
	return carbon.RFC3339Layout
}

func TestLayoutType_Customer(t *testing.T) {
	type User struct {
		Customer carbon.LayoutType[RFC3339Layout] `json:"customer"`
	}

	var user User

	c := carbon.Parse("2020-08-05 13:14:15")
	user.Customer = carbon.NewLayoutType[RFC3339Layout](c)

	data, err := json.Marshal(&user)
	assert.NoError(t, err)
	assert.Equal(t, `{"customer":"2020-08-05T13:14:15Z"}`, string(data))

	var person User
	err = json.Unmarshal(data, &person)

	assert.NoError(t, err)
	assert.Equal(t, "2020-08-05T13:14:15Z", person.Customer.String())
}

type ISO8601Format struct{}

func (t ISO8601Format) SetFormat() string {
	return carbon.ISO8601Format
}

func TestFormatType_Customer(t *testing.T) {
	type User struct {
		Customer carbon.FormatType[ISO8601Format] `json:"customer"`
	}

	var user User

	c := carbon.Parse("2020-08-05 13:14:15")
	user.Customer = carbon.NewFormatType[ISO8601Format](c)

	data, err := json.Marshal(&user)
	assert.NoError(t, err)
	assert.Equal(t, `{"customer":"2020-08-05T13:14:15+00:00"}`, string(data))

	var person User
	err = json.Unmarshal(data, &person)

	assert.NoError(t, err)
	assert.Equal(t, "2020-08-05T13:14:15+00:00", person.Customer.String())
}
