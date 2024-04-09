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
		"birthday17":"2020-08-05 13:14:15"
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
		"birthday17":"1596604455999999999"
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
