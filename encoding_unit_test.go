package carbon

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`

	Birthday1 Carbon `json:"birthday1"`

	Birthday2 DateTime      `json:"birthday2"`
	Birthday3 DateTimeMilli `json:"birthday3"`
	Birthday4 DateTimeMicro `json:"birthday4"`
	Birthday5 DateTimeNano  `json:"birthday5"`

	Birthday6 Date      `json:"birthday6"`
	Birthday7 DateMilli `json:"birthday7"`
	Birthday8 DateMicro `json:"birthday8"`
	Birthday9 DateNano  `json:"birthday9"`

	Birthday10 Time      `json:"birthday10"`
	Birthday11 TimeMilli `json:"birthday11"`
	Birthday12 TimeMicro `json:"birthday12"`
	Birthday13 TimeNano  `json:"birthday13"`

	Birthday14 Timestamp      `json:"birthday14"`
	Birthday15 TimestampMilli `json:"birthday15"`
	Birthday16 TimestampMicro `json:"birthday16"`
	Birthday17 TimestampNano  `json:"birthday17"`
}

func TestCarbon_MarshalJSON(t *testing.T) {
	var c = Parse("2020-08-05 13:14:15.999999999", PRC)
	var person = Person{
		Name: "gouguoyin",
		Age:  18,

		Birthday1: c,
		Birthday2: c.ToDateTimeStruct(),
		Birthday3: c.ToDateTimeMilliStruct(),
		Birthday4: c.ToDateTimeMicroStruct(),
		Birthday5: c.ToDateTimeNanoStruct(),

		Birthday6: c.ToDateStruct(),
		Birthday7: c.ToDateMilliStruct(),
		Birthday8: c.ToDateMicroStruct(),
		Birthday9: c.ToDateNanoStruct(),

		Birthday10: c.ToTimeStruct(),
		Birthday11: c.ToTimeMilliStruct(),
		Birthday12: c.ToTimeMicroStruct(),
		Birthday13: c.ToTimeNanoStruct(),

		Birthday14: c.ToTimestampStruct(),
		Birthday15: c.ToTimestampMilliStruct(),
		Birthday16: c.ToTimestampMicroStruct(),
		Birthday17: c.ToTimestampNanoStruct(),
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
			name:   "Birthday3",
			actual: person.Birthday3.String(),
			want:   "2020-08-05 13:14:15.999",
		},
		{
			name:   "Birthday4",
			actual: person.Birthday4.String(),
			want:   "2020-08-05 13:14:15.999999",
		},
		{
			name:   "Birthday5",
			actual: person.Birthday5.String(),
			want:   "2020-08-05 13:14:15.999999999",
		},

		{
			name:   "Birthday6",
			actual: person.Birthday6.String(),
			want:   "2020-08-05",
		},
		{
			name:   "Birthday7",
			actual: person.Birthday7.String(),
			want:   "2020-08-05.999",
		},
		{
			name:   "Birthday8",
			actual: person.Birthday8.String(),
			want:   "2020-08-05.999999",
		},
		{
			name:   "Birthday9",
			actual: person.Birthday9.String(),
			want:   "2020-08-05.999999999",
		},

		{
			name:   "Birthday10",
			actual: person.Birthday10.String(),
			want:   "13:14:15",
		},
		{
			name:   "Birthday11",
			actual: person.Birthday11.String(),
			want:   "13:14:15.999",
		},
		{
			name:   "Birthday12",
			actual: person.Birthday12.String(),
			want:   "13:14:15.999999",
		},
		{
			name:   "Birthday13",
			actual: person.Birthday13.String(),
			want:   "13:14:15.999999999",
		},

		{
			name:   "Birthday14",
			actual: person.Birthday14.String(),
			want:   "1596604455",
		},
		{
			name:   "Birthday15",
			actual: person.Birthday15.String(),
			want:   "1596604455999",
		},
		{
			name:   "Birthday16",
			actual: person.Birthday16.String(),
			want:   "1596604455999999",
		},
		{
			name:   "Birthday17",
			actual: person.Birthday17.String(),
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
		"Birthday2":"2020-08-05 13:14:15",
		"Birthday3":"2020-08-05 13:14:15.999",
		"Birthday4":"2020-08-05 13:14:15.999999",
		"Birthday5":"2020-08-05 13:14:15.999999999",
		"Birthday6":"2020-08-05",
		"Birthday7":"2020-08-05.999",
		"Birthday8":"2020-08-05.999999",
		"Birthday9":"2020-08-05.999999999",
		"Birthday10":"13:14:15",
		"Birthday11":"13:14:15.999",
		"Birthday12":"13:14:15.999999",
		"Birthday13":"13:14:15.999999999",
		"Birthday14":1596604455,
		"Birthday15":1596604455999,
		"Birthday16":1596604455999999,
		"Birthday17":1596604455999999999
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
			name:   "Birthday2",
			actual: person.Birthday2.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "Birthday3",
			actual: person.Birthday3.String(),
			want:   "2020-08-05 13:14:15.999",
		},
		{
			name:   "Birthday4",
			actual: person.Birthday4.String(),
			want:   "2020-08-05 13:14:15.999999",
		},
		{
			name:   "Birthday5",
			actual: person.Birthday5.String(),
			want:   "2020-08-05 13:14:15.999999999",
		},
		{
			name:   "Birthday6",
			actual: person.Birthday6.String(),
			want:   "2020-08-05",
		},
		{
			name:   "Birthday7",
			actual: person.Birthday7.String(),
			want:   "2020-08-05.999",
		},
		{
			name:   "Birthday8",
			actual: person.Birthday8.String(),
			want:   "2020-08-05.999999",
		},
		{
			name:   "Birthday9",
			actual: person.Birthday9.String(),
			want:   "2020-08-05.999999999",
		},

		{
			name:   "Birthday10",
			actual: person.Birthday10.String(),
			want:   "13:14:15",
		},
		{
			name:   "Birthday11",
			actual: person.Birthday11.String(),
			want:   "13:14:15.999",
		},
		{
			name:   "Birthday12",
			actual: person.Birthday12.String(),
			want:   "13:14:15.999999",
		},
		{
			name:   "Birthday13",
			actual: person.Birthday13.String(),
			want:   "13:14:15.999999999",
		},

		{
			name:   "Birthday14",
			actual: person.Birthday14.String(),
			want:   "1596604455",
		},
		{
			name:   "Birthday15",
			actual: person.Birthday15.String(),
			want:   "1596604455999",
		},
		{
			name:   "Birthday16",
			actual: person.Birthday16.String(),
			want:   "1596604455999999",
		},
		{
			name:   "Birthday17",
			actual: person.Birthday17.String(),
			want:   "1596604455999999999",
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
		"Birthday14":1596604455,
		"Birthday15":1596604455999,
		"Birthday16":1596604455999999,
		"Birthday17":1596604455999999999
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
			name:   "Birthday14",
			actual: person.Birthday14.Int64(),
			want:   1596604455,
		},
		{
			name:   "Birthday15",
			actual: person.Birthday15.Int64(),
			want:   1596604455999,
		},
		{
			name:   "Birthday16",
			actual: person.Birthday16.Int64(),
			want:   1596604455999999,
		},
		{
			name:   "Birthday17",
			actual: person.Birthday17.Int64(),
			want:   1596604455999999999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "json.Unmarshal()")
		})
	}
}

// https://github.com/dromara/carbon/issues/225
func TestCarbon_Issue225(t *testing.T) {
	emptyStr := `{
		"birthday1":"",
		"birthday2":"",
		"birthday3":"",
		"birthday4":"",
		"birthday5":"",
		"birthday6":"",
		"birthday7":"",
		"birthday8":"",
		"birthday9":"",
		"birthday10":"",
		"birthday11":"",
		"birthday12":"",
		"birthday13":"",
		"birthday14":"",
		"birthday15":"",
		"birthday16":"",
		"birthday17":""
	}`
	nullStr := `{
		"birthday1":null,
		"birthday2":null,
		"birthday3":null,
		"birthday4":null,
		"birthday5":null,
		"birthday6":null,
		"birthday7":null,
		"birthday8":null,
		"birthday9":null,
		"birthday10":null,
		"birthday11":null,
		"birthday12":null,
		"birthday13":null,
		"birthday14":null,
		"birthday15":null,
		"birthday16":null,
		"birthday17":null
	}`

	var person Person
	emptyErr := json.Unmarshal([]byte(emptyStr), &person)
	assert.Nil(t, emptyErr)

	nullErr := json.Unmarshal([]byte(nullStr), &person)
	assert.Nil(t, nullErr)
}

// https://github.com/dromara/carbon/issues/240
func TestCarbon_Issue240(t *testing.T) {
	jsonStr := `{
		"birthday1":"",
		"birthday2":null
	}`

	var person Person
	emptyErr := json.Unmarshal([]byte(jsonStr), &person)
	assert.Nil(t, emptyErr)
	assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", person.Birthday1.StdTime().String())
	assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", person.Birthday2.StdTime().String())
	assert.Equal(t, true, person.Birthday1.IsZero())
	assert.Equal(t, true, person.Birthday2.IsZero())
	assert.Equal(t, false, person.Birthday1.IsValid())
	assert.Equal(t, false, person.Birthday2.IsValid())
}

// https://github.com/dromara/carbon/issues/243
func TestCarbon_Issue243(t *testing.T) {
	type Project struct {
		StartDate DateTime `gorm:"column:start_date" json:"startDate"`
		EndDate   DateTime `gorm:"column:end_date" json:"endDate"`
	}

	project := new(Project)
	jsonStr := `{"startDate":"2024-10-01 00:00:00","endDate":"2024-10-31 23:59:59"}`
	err := json.Unmarshal([]byte(jsonStr), &project)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, "Local", project.StartDate.Location())
	assert.Equal(t, "Local", project.EndDate.Location())
}
