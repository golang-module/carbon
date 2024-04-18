package carbon

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`

	Birthday1 Carbon `json:"birthday1"`
	Birthday2 Carbon `json:"birthday2" carbon:"type:dateTime"`
	Birthday3 Carbon `json:"birthday3" tz:"PRC"`
	Birthday4 Carbon `json:"birthday4" carbon:"" tz:""`

	Birthday5 Carbon `json:"birthday5" carbon:"layout:2006-01-02" tz:"PRC"`
	Birthday6 Carbon `json:"birthday6" carbon:"layout:15:04:05" tz:"PRC"`
	Birthday7 Carbon `json:"birthday7" carbon:"layout:2006-01-02 15:04:05" tz:"PRC"`

	Birthday8  Carbon `json:"birthday8" carbon:"format:Y-m-d" tz:"PRC"`
	Birthday9  Carbon `json:"birthday9" carbon:"format:H:i:s" tz:"PRC"`
	Birthday10 Carbon `json:"birthday10" carbon:"format:Y-m-d H:i:s" tz:"PRC"`

	Birthday11 Carbon `json:"birthday11" carbon:"type:date" tz:"PRC"`
	Birthday12 Carbon `json:"birthday12" carbon:"type:time" tz:"PRC"`
	Birthday13 Carbon `json:"birthday13" carbon:"type:dateTime" tz:"PRC"`

	Birthday14 Carbon `json:"birthday14" carbon:"type:timestamp" tz:"PRC"`
	Birthday15 Carbon `json:"birthday15" carbon:"type:timestampMilli" tz:"PRC"`
	Birthday16 Carbon `json:"birthday16" carbon:"type:timestampMicro" tz:"PRC"`
	Birthday17 Carbon `json:"birthday17" carbon:"type:timestampNano" tz:"PRC"`

	Birthday18 DateTime      `json:"birthday18"`
	Birthday19 DateTimeMilli `json:"birthday19"`
	Birthday20 DateTimeMicro `json:"birthday20"`
	Birthday21 DateTimeNano  `json:"birthday21"`

	Birthday22 Date      `json:"birthday22"`
	Birthday23 DateMilli `json:"birthday23"`
	Birthday24 DateMicro `json:"birthday24"`
	Birthday25 DateNano  `json:"birthday25"`

	Birthday26 Time      `json:"birthday26"`
	Birthday27 TimeMilli `json:"birthday27"`
	Birthday28 TimeMicro `json:"birthday28"`
	Birthday29 TimeNano  `json:"birthday29"`

	Birthday30 Timestamp      `json:"birthday30"`
	Birthday31 TimestampMilli `json:"birthday31"`
	Birthday32 TimestampMicro `json:"birthday32"`
	Birthday33 TimestampNano  `json:"birthday33"`
}

func TestCarbon_MarshalJSON(t *testing.T) {
	var c = Parse("2020-08-05 13:14:15.999999999", PRC)
	var person = Person{
		Name:       "gouguoyin",
		Age:        18,
		Birthday1:  c,
		Birthday2:  c,
		Birthday3:  c,
		Birthday4:  c,
		Birthday5:  c,
		Birthday6:  c,
		Birthday7:  c,
		Birthday8:  c,
		Birthday9:  c,
		Birthday10: c,
		Birthday11: c,
		Birthday12: c,
		Birthday13: c,
		Birthday14: c,
		Birthday15: c,
		Birthday16: c,
		Birthday17: c,

		Birthday18: c.ToDateTimeStruct(),
		Birthday19: c.ToDateTimeMilliStruct(),
		Birthday20: c.ToDateTimeMicroStruct(),
		Birthday21: c.ToDateTimeNanoStruct(),

		Birthday22: c.ToDateStruct(),
		Birthday23: c.ToDateMilliStruct(),
		Birthday24: c.ToDateMicroStruct(),
		Birthday25: c.ToDateNanoStruct(),

		Birthday26: c.ToTimeStruct(),
		Birthday27: c.ToTimeMilliStruct(),
		Birthday28: c.ToTimeMicroStruct(),
		Birthday29: c.ToTimeNanoStruct(),

		Birthday30: c.ToTimestampStruct(),
		Birthday31: c.ToTimestampMilliStruct(),
		Birthday32: c.ToTimestampMicroStruct(),
		Birthday33: c.ToTimestampNanoStruct(),
	}

	data, marshalErr := json.Marshal(&person)
	assert.Nil(t, marshalErr)
	fmt.Printf("json encode:\n%s\n", data)

	tests := []struct {
		name   string
		want   string
		actual string
	}{
		{
			name:   "birthday1",
			actual: person.Birthday1.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday2",
			actual: person.Birthday2.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday3",
			actual: person.Birthday3.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday4",
			actual: person.Birthday4.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday5",
			actual: person.Birthday5.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday6",
			actual: person.Birthday6.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday7",
			actual: person.Birthday7.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday8",
			actual: person.Birthday8.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday9",
			actual: person.Birthday9.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday10",
			actual: person.Birthday10.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday11",
			actual: person.Birthday11.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday12",
			actual: person.Birthday12.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday13",
			actual: person.Birthday13.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday14",
			actual: person.Birthday14.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday15",
			actual: person.Birthday15.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday16",
			actual: person.Birthday16.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday17",
			actual: person.Birthday17.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday18",
			actual: person.Birthday18.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday19",
			actual: person.Birthday19.String(),
			want:   "2020-08-05 13:14:15.999",
		},
		{
			name:   "birthday20",
			actual: person.Birthday20.String(),
			want:   "2020-08-05 13:14:15.999999",
		},
		{
			name:   "birthday21",
			actual: person.Birthday21.String(),
			want:   "2020-08-05 13:14:15.999999999",
		},

		{
			name:   "birthday22",
			actual: person.Birthday22.String(),
			want:   "2020-08-05",
		},
		{
			name:   "birthday23",
			actual: person.Birthday23.String(),
			want:   "2020-08-05.999",
		},
		{
			name:   "birthday24",
			actual: person.Birthday24.String(),
			want:   "2020-08-05.999999",
		},
		{
			name:   "birthday25",
			actual: person.Birthday25.String(),
			want:   "2020-08-05.999999999",
		},

		{
			name:   "birthday26",
			actual: person.Birthday26.String(),
			want:   "13:14:15",
		},
		{
			name:   "birthday27",
			actual: person.Birthday27.String(),
			want:   "13:14:15.999",
		},
		{
			name:   "birthday28",
			actual: person.Birthday28.String(),
			want:   "13:14:15.999999",
		},
		{
			name:   "birthday29",
			actual: person.Birthday29.String(),
			want:   "13:14:15.999999999",
		},

		{
			name:   "birthday30",
			actual: person.Birthday30.String(),
			want:   "1596604455",
		},
		{
			name:   "birthday31",
			actual: person.Birthday31.String(),
			want:   "1596604455999",
		},
		{
			name:   "birthday32",
			actual: person.Birthday32.String(),
			want:   "1596604455999999",
		},
		{
			name:   "birthday33",
			actual: person.Birthday33.String(),
			want:   "1596604455999999999",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "json.Marshal()")
		})
	}
}

func TestCarbon_UnmarshalJSON(t *testing.T) {
	jsonStr := `{
		"name":"gouguoyin",
		"age":18,
		"birthday1":"2020-08-05 13:14:15",
		"birthday2":"2020-08-05 13:14:15",
		"birthday3":"2020-08-05 13:14:15",
		"birthday4":"2020-08-05 13:14:15",
		"birthday5":"2020-08-05 13:14:15",
		"birthday6":"2020-08-05 13:14:15",
		"birthday7":"2020-08-05 13:14:15",
		"birthday8":"2020-08-05 13:14:15",
		"birthday9":"2020-08-05 13:14:15",
		"birthday10":"2020-08-05 13:14:15",
		"birthday11":"2020-08-05 13:14:15",
		"birthday12":"2020-08-05 13:14:15",
		"birthday13":"2020-08-05 13:14:15",
		"birthday14":"2020-08-05 13:14:15",
		"birthday15":"2020-08-05 13:14:15",
		"birthday16":"2020-08-05 13:14:15",
		"birthday17":"2020-08-05 13:14:15",
		"birthday18":"2020-08-05 13:14:15",
		"birthday19":"2020-08-05 13:14:15.999",
		"birthday20":"2020-08-05 13:14:15.999999",
		"birthday21":"2020-08-05 13:14:15.999999999",
		"birthday22":"2020-08-05",
		"birthday23":"2020-08-05.999",
		"birthday24":"2020-08-05.999999",
		"birthday25":"2020-08-05.999999999",
		"birthday26":"13:14:15",
		"birthday27":"13:14:15.999",
		"birthday28":"13:14:15.999999",
		"birthday29":"13:14:15.999999999",
		"birthday30":1596604455,
		"birthday31":1596604455999,
		"birthday32":1596604455999999,
		"birthday33":1596604455999999999
	}`

	var person Person
	unmarshalErr := json.Unmarshal([]byte(jsonStr), &person)
	assert.Nil(t, unmarshalErr)
	fmt.Printf("json decode:\n%+v\n", person)

	tests := []struct {
		name   string
		want   string
		actual string
	}{
		{
			name:   "birthday1",
			actual: person.Birthday1.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday2",
			actual: person.Birthday2.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday3",
			actual: person.Birthday3.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday4",
			actual: person.Birthday4.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday5",
			actual: person.Birthday5.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday6",
			actual: person.Birthday6.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday7",
			actual: person.Birthday7.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday8",
			actual: person.Birthday8.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday9",
			actual: person.Birthday9.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday10",
			actual: person.Birthday10.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday11",
			actual: person.Birthday11.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday12",
			actual: person.Birthday12.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday13",
			actual: person.Birthday13.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday14",
			actual: person.Birthday14.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday15",
			actual: person.Birthday15.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday16",
			actual: person.Birthday16.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday17",
			actual: person.Birthday17.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday18",
			actual: person.Birthday18.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday19",
			actual: person.Birthday19.String(),
			want:   "2020-08-05 13:14:15.999",
		},
		{
			name:   "birthday20",
			actual: person.Birthday20.String(),
			want:   "2020-08-05 13:14:15.999999",
		},
		{
			name:   "birthday21",
			actual: person.Birthday21.String(),
			want:   "2020-08-05 13:14:15.999999999",
		},
		{
			name:   "birthday22",
			actual: person.Birthday22.String(),
			want:   "2020-08-05",
		},
		{
			name:   "birthday23",
			actual: person.Birthday23.String(),
			want:   "2020-08-05.999",
		},
		{
			name:   "birthday24",
			actual: person.Birthday24.String(),
			want:   "2020-08-05.999999",
		},
		{
			name:   "birthday25",
			actual: person.Birthday25.String(),
			want:   "2020-08-05.999999999",
		},

		{
			name:   "birthday26",
			actual: person.Birthday26.String(),
			want:   "13:14:15",
		},
		{
			name:   "birthday27",
			actual: person.Birthday27.String(),
			want:   "13:14:15.999",
		},
		{
			name:   "birthday28",
			actual: person.Birthday28.String(),
			want:   "13:14:15.999999",
		},
		{
			name:   "birthday29",
			actual: person.Birthday29.String(),
			want:   "13:14:15.999999999",
		},

		{
			name:   "birthday30",
			actual: person.Birthday30.String(),
			want:   "1596604455",
		},
		{
			name:   "birthday31",
			actual: person.Birthday31.String(),
			want:   "1596604455999",
		},
		{
			name:   "birthday32",
			actual: person.Birthday32.String(),
			want:   "1596604455999999",
		},
		{
			name:   "birthday33",
			actual: person.Birthday33.String(),
			want:   "1596604455999999999",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "json.Unmarshal()")
		})
	}
}

func TestCarbon_MarshalJSON_LoadTag(t *testing.T) {
	var c = Parse("2020-08-05 13:14:15.999999999", PRC)
	var person = Person{
		Name:       "gouguoyin",
		Age:        18,
		Birthday1:  c,
		Birthday2:  c,
		Birthday3:  c,
		Birthday4:  c,
		Birthday5:  c,
		Birthday6:  c,
		Birthday7:  c,
		Birthday8:  c,
		Birthday9:  c,
		Birthday10: c,
		Birthday11: c,
		Birthday12: c,
		Birthday13: c,
		Birthday14: c,
		Birthday15: c,
		Birthday16: c,
		Birthday17: c,

		Birthday18: DateTime{Carbon: c},
		Birthday19: DateTimeMilli{Carbon: c},
		Birthday20: DateTimeMicro{Carbon: c},
		Birthday21: DateTimeNano{Carbon: c},

		Birthday22: Date{Carbon: c},
		Birthday23: DateMilli{Carbon: c},
		Birthday24: DateMicro{Carbon: c},
		Birthday25: DateNano{Carbon: c},

		Birthday26: Time{Carbon: c},
		Birthday27: TimeMilli{Carbon: c},
		Birthday28: TimeMicro{Carbon: c},
		Birthday29: TimeNano{Carbon: c},

		Birthday30: Timestamp{Carbon: c},
		Birthday31: TimestampMilli{Carbon: c},
		Birthday32: TimestampMicro{Carbon: c},
		Birthday33: TimestampNano{Carbon: c},
	}

	loadErr := LoadTag(&person)
	assert.Nil(t, loadErr)

	data, marshalErr := json.Marshal(&person)
	assert.Nil(t, marshalErr)

	fmt.Printf("json encode:\n%s\n", data)

	tests := []struct {
		name   string
		want   string
		actual string
	}{
		{
			name:   "birthday1",
			actual: person.Birthday1.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday2",
			actual: person.Birthday2.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday3",
			actual: person.Birthday3.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday4",
			actual: person.Birthday4.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday5",
			actual: person.Birthday5.String(),
			want:   "2020-08-05",
		},
		{
			name:   "birthday6",
			actual: person.Birthday6.String(),
			want:   "13:14:15",
		},
		{
			name:   "birthday7",
			actual: person.Birthday7.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday8",
			actual: person.Birthday8.String(),
			want:   "2020-08-05",
		},
		{
			name:   "birthday9",
			actual: person.Birthday9.String(),
			want:   "13:14:15",
		},
		{
			name:   "birthday10",
			actual: person.Birthday10.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday11",
			actual: person.Birthday11.String(),
			want:   "2020-08-05",
		},
		{
			name:   "birthday12",
			actual: person.Birthday12.String(),
			want:   "13:14:15",
		},
		{
			name:   "birthday13",
			actual: person.Birthday13.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday14",
			actual: person.Birthday14.String(),
			want:   "1596604455",
		},
		{
			name:   "birthday15",
			actual: person.Birthday15.String(),
			want:   "1596604455999",
		},
		{
			name:   "birthday16",
			actual: person.Birthday16.String(),
			want:   "1596604455999999",
		},
		{
			name:   "birthday17",
			actual: person.Birthday17.String(),
			want:   "1596604455999999999",
		},
		{
			name:   "birthday18",
			actual: person.Birthday18.String(),
			want:   "2020-08-05 13:14:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "json.Marshal()")
		})
	}
}

func TestCarbon_UnmarshalJSON_LoadTag(t *testing.T) {
	jsonStr := `{
		"name":"gouguoyin",
		"age":18,
		"birthday1":"2020-08-05 13:14:15",
		"birthday2":"2020-08-05 13:14:15",
		"birthday3":"2020-08-05 13:14:15",
		"birthday4":"2020-08-05 13:14:15",
		"birthday5":"2020-08-05",
		"birthday6":"13:14:15",
		"birthday7":"2020-08-05 13:14:15",
		"birthday8":"2020-08-05",
		"birthday9":"13:14:15",
		"birthday10":"2020-08-05 13:14:15",
		"birthday11":"2020-08-05",
		"birthday12":"13:14:15",
		"birthday13":"2020-08-05 13:14:15",
		"birthday14":"1596604455",
		"birthday15":"1596604455999",
		"birthday16":"1596604455999999",
		"birthday17":"1596604455999999999",
		"birthday18":"2020-08-05 13:14:15"
	}`
	var person Person
	loadErr := LoadTag(&person)
	assert.Nil(t, loadErr)
	unmarshalErr := json.Unmarshal([]byte(jsonStr), &person)
	assert.Nil(t, unmarshalErr)
	fmt.Printf("json decode:\n%+v\n", person)

	tests := []struct {
		name   string
		want   string
		actual string
	}{
		{
			name:   "birthday1",
			actual: person.Birthday1.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday2",
			actual: person.Birthday2.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday3",
			actual: person.Birthday3.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday4",
			actual: person.Birthday4.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday5",
			actual: person.Birthday5.String(),
			want:   "2020-08-05",
		},
		{
			name:   "birthday6",
			actual: person.Birthday6.String(),
			want:   "13:14:15",
		},
		{
			name:   "birthday7",
			actual: person.Birthday7.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday8",
			actual: person.Birthday8.String(),
			want:   "2020-08-05",
		},
		{
			name:   "birthday9",
			actual: person.Birthday9.String(),
			want:   "13:14:15",
		},
		{
			name:   "birthday10",
			actual: person.Birthday10.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday11",
			actual: person.Birthday11.String(),
			want:   "2020-08-05",
		},
		{
			name:   "birthday12",
			actual: person.Birthday12.String(),
			want:   "13:14:15",
		},
		{
			name:   "birthday13",
			actual: person.Birthday13.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday14",
			actual: person.Birthday14.String(),
			want:   "1596604455",
		},
		{
			name:   "birthday15",
			actual: person.Birthday15.String(),
			want:   "1596604455999",
		},
		{
			name:   "birthday16",
			actual: person.Birthday16.String(),
			want:   "1596604455999999",
		},
		{
			name:   "birthday17",
			actual: person.Birthday17.String(),
			want:   "1596604455999999999",
		},
		{
			name:   "birthday18",
			actual: person.Birthday18.String(),
			want:   "2020-08-05 13:14:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "json.Unmarshal()")
		})
	}
}

func TestCarbon_TimestampToInt64(t *testing.T) {
	jsonStr := `{
		"birthday30":1596604455,
		"birthday31":1596604455999,
		"birthday32":1596604455999999,
		"birthday33":1596604455999999999
	}`
	var person Person
	unmarshalErr := json.Unmarshal([]byte(jsonStr), &person)
	assert.Nil(t, unmarshalErr)
	fmt.Printf("json decode:\n%+v\n", person)

	tests := []struct {
		name   string
		want   int64
		actual int64
	}{
		{
			name:   "birthday30",
			actual: person.Birthday30.Int64(),
			want:   1596604455,
		},
		{
			name:   "birthday31",
			actual: person.Birthday31.Int64(),
			want:   1596604455999,
		},
		{
			name:   "birthday32",
			actual: person.Birthday32.Int64(),
			want:   1596604455999999,
		},
		{
			name:   "birthday33",
			actual: person.Birthday33.Int64(),
			want:   1596604455999999999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "json.Unmarshal()")
		})
	}
}
func TestError_Json(t *testing.T) {
	type Student struct {
		Birthday1 Carbon `json:"birthday1" carbon:"dateTime"`
		Birthday2 Carbon `json:"birthday2" carbon:"layout:xxx"`
		Birthday3 Carbon `json:"birthday3" carbon:"format:xxx"`
		Birthday4 Carbon `json:"birthday4" carbon:"xxx"`
		Birthday5 Carbon `json:"birthday5" tz:"xxx"`
	}

	student := Student{
		Birthday1: Parse("XXX"),
		Birthday2: Parse("2020-08-05"),
		Birthday3: Parse("2020-08-05"),
		Birthday4: Parse("2020-08-05"),
		Birthday5: Parse("2020-08-05"),
	}

	_, marshalErr := json.Marshal(student)
	fmt.Println("marshal error:", marshalErr.Error())
	assert.NotNil(t, marshalErr)

	str := `{
		"name": "gouguoyin",
		"age": 18,
		"birthday1": "2020-08-05 13:14:15",
		"birthday2": "2020-08-05 13:14:15",
		"birthday3": "2020-08-05 13:14:15",
		"birthday4": "2020-08-05 13:14:15",
		"birthday5": "2020-08-05 13:14:15"
	}`

	unmarshalErr := json.Unmarshal([]byte(str), &student)
	fmt.Println("unmarshal error:", unmarshalErr.Error())
	assert.NotNil(t, unmarshalErr)
}

// https://github.com/golang-module/carbon/issues/225
func TestCarbon_Issue225(t *testing.T) {
	str := `{
		"birthday1":"",
		"birthday2":null
	}`

	var person Person
	unmarshalErr := json.Unmarshal([]byte(str), &person)
	assert.Nil(t, unmarshalErr)

	assert.Equal(t, "", person.Birthday1.String())
	assert.Equal(t, "", person.Birthday2.String())
}
