package carbon

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday1.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday2.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday3.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday4.String())

	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday5.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday6.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday7.String())

	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday8.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday9.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday10.String())

	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday11.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday12.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday13.String())

	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday14.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday15.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday16.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday17.String())

	fmt.Printf("person output by json:\n%s\n", data)
}

func TestCarbon_UnmarshalJSON(t *testing.T) {
	str := `{
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
	unmarshalErr := json.Unmarshal([]byte(str), &person)
	assert.Nil(t, unmarshalErr)

	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday1.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday2.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday3.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday4.String())

	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday5.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday6.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday7.String())

	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday8.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday9.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday10.String())

	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday11.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday12.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday13.String())

	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday14.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday15.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday16.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday17.String())

	fmt.Printf("Json string parse to person:\n%+v\n", person)
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

	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday1.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday2.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday3.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday4.String())

	assert.Equal(t, "2020-08-05", person.Birthday5.String())
	assert.Equal(t, "13:14:15", person.Birthday6.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday7.String())

	assert.Equal(t, "2020-08-05", person.Birthday8.String())
	assert.Equal(t, "13:14:15", person.Birthday9.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday10.String())

	assert.Equal(t, "2020-08-05", person.Birthday11.String())
	assert.Equal(t, "13:14:15", person.Birthday12.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday13.String())

	assert.Equal(t, "1596604455", person.Birthday14.String())
	assert.Equal(t, "1596604455999", person.Birthday15.String())
	assert.Equal(t, "1596604455999999", person.Birthday16.String())
	assert.Equal(t, "1596604455999999999", person.Birthday17.String())

	fmt.Printf("person output by json:\n%s\n", data)
}

func TestCarbon_UnmarshalJSON_LoadTag(t *testing.T) {
	str := `{
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

	unmarshalErr := json.Unmarshal([]byte(str), &person)
	assert.Nil(t, unmarshalErr)

	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday1.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday2.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday3.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday4.String())

	assert.Equal(t, "2020-08-05", person.Birthday5.String())
	assert.Equal(t, "13:14:15", person.Birthday6.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday7.String())

	assert.Equal(t, "2020-08-05", person.Birthday8.String())
	assert.Equal(t, "13:14:15", person.Birthday9.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday10.String())

	assert.Equal(t, "2020-08-05", person.Birthday11.String())
	assert.Equal(t, "13:14:15", person.Birthday12.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday13.String())

	assert.Equal(t, "1596604455", person.Birthday14.String())
	assert.Equal(t, "1596604455999", person.Birthday15.String())
	assert.Equal(t, "1596604455999999", person.Birthday16.String())
	assert.Equal(t, "1596604455999999999", person.Birthday17.String())

	fmt.Printf("Json string parse to person:\n%+v\n", person)
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
