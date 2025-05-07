package carbon

import "github.com/dromara/carbon/v2"

type (
	Timestamp      = carbon.Timestamp
	TimestampMilli = carbon.TimestampMilli
	TimestampMicro = carbon.TimestampMicro
	TimestampNano  = carbon.TimestampNano

	DateTime      = carbon.DateTime
	DateTimeMicro = carbon.DateTimeMicro
	DateTimeMilli = carbon.DateTimeMilli
	DateTimeNano  = carbon.DateTimeNano

	Date      = carbon.Date
	DateMilli = carbon.DateMilli
	DateMicro = carbon.DateMicro
	DateNano  = carbon.DateNano

	Time      = carbon.Time
	TimeMilli = carbon.TimeMilli
	TimeMicro = carbon.TimeMicro
	TimeNano  = carbon.TimeNano
)

func NewTimestamp(c Carbon) Timestamp {
	return carbon.NewTimestamp(c)
}
func NewTimestampMilli(c Carbon) TimestampMilli {
	return carbon.NewTimestampMilli(c)
}
func NewTimestampMicro(c Carbon) TimestampMicro {
	return carbon.NewTimestampMicro(c)
}
func NewTimestampNano(c Carbon) TimestampNano {
	return carbon.NewTimestampNano(c)
}

func NewDateTime(c Carbon) DateTime {
	return carbon.NewDateTime(c)
}
func NewDateTimeMilli(c Carbon) DateTimeMilli {
	return carbon.NewDateTimeMilli(c)
}
func NewDateTimeMicro(c Carbon) DateTimeMicro {
	return carbon.NewDateTimeMicro(c)
}
func NewDateTimeNano(c Carbon) DateTimeNano {
	return carbon.NewDateTimeNano(c)
}

func NewDate(c Carbon) Date {
	return carbon.NewDate(c)
}
func NewDateMilli(c Carbon) DateMilli {
	return carbon.NewDateMilli(c)
}
func NewDateMicro(c Carbon) DateMicro {
	return carbon.NewDateMicro(c)
}
func NewDateNano(c Carbon) DateNano {
	return carbon.NewDateNano(c)
}

func NewTime(c Carbon) Time {
	return carbon.NewTime(c)
}
func NewTimeMilli(c Carbon) TimeMilli {
	return carbon.NewTimeMilli(c)
}
func NewTimeMicro(c Carbon) TimeMicro {
	return carbon.NewTimeMicro(c)
}
func NewTimeNano(c Carbon) TimeNano {
	return carbon.NewTimeNano(c)
}
