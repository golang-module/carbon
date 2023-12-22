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

	Birthday1 Carbon `json:"birthday1" carbon:"layout:2006-01-02 15:04:05"`
	Birthday2 Carbon `json:"birthday2" carbon:"layout:2006-01-02 15:04:05.999"`
	Birthday3 Carbon `json:"birthday3" carbon:"layout:2006-01-02 15:04:05.999999"`
	Birthday4 Carbon `json:"birthday4" carbon:"layout:2006-01-02 15:04:05.999999999"`

	GraduatedAt1 Carbon `json:"graduated_at1" carbon:"layout:2006-01-02"`
	GraduatedAt2 Carbon `json:"graduated_at2" carbon:"layout:2006-01-02.999"`
	GraduatedAt3 Carbon `json:"graduated_at3" carbon:"layout:2006-01-02.999999"`
	GraduatedAt4 Carbon `json:"graduated_at4" carbon:"layout:2006-01-02.999999999"`

	OperatedAt1 Carbon `json:"operated_at1" carbon:"layout:15:04:05"`
	OperatedAt2 Carbon `json:"operated_at2" carbon:"layout:15:04:05.999"`
	OperatedAt3 Carbon `json:"operated_at3" carbon:"layout:15:04:05.999999"`
	OperatedAt4 Carbon `json:"operated_at4" carbon:"layout:15:04:05.999999999"`

	CreatedAt1 Carbon `json:"created_at1" carbon:"format:Y-m-d"`
	CreatedAt2 Carbon `json:"created_at2" carbon:"format:H:i:s"`
	CreatedAt3 Carbon `json:"created_at3" carbon:"format:Y-m-d H:i:s"`
	CreatedAt4 Carbon `json:"created_at4"`
}

type Student struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Birthday Carbon `json:"birthday"`
}

func TestCarbon_MarshalJSON_LoadTag(t *testing.T) {
	c := Parse("2020-08-05 13:14:15.999999999", PRC)
	person := Person{
		Name: "gouguoyin",
		Age:  18,

		Birthday1: c,
		Birthday2: c,
		Birthday3: c,
		Birthday4: c,

		GraduatedAt1: c,
		GraduatedAt2: c,
		GraduatedAt3: c,
		GraduatedAt4: c,

		OperatedAt1: c,
		OperatedAt2: c,
		OperatedAt3: c,
		OperatedAt4: c,

		CreatedAt1: c,
		CreatedAt2: c,
		CreatedAt3: c,
		CreatedAt4: c,
	}

	loadErr := LoadTag(&person)
	assert.Nil(t, loadErr)

	data, marshalErr := json.Marshal(&person)
	assert.Nil(t, marshalErr)

	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday1.String())
	assert.Equal(t, "2020-08-05 13:14:15.999", person.Birthday2.String())
	assert.Equal(t, "2020-08-05 13:14:15.999999", person.Birthday3.String())
	assert.Equal(t, "2020-08-05 13:14:15.999999999", person.Birthday4.String())

	assert.Equal(t, "2020-08-05", person.GraduatedAt1.String())
	assert.Equal(t, "2020-08-05.999", person.GraduatedAt2.String())
	assert.Equal(t, "2020-08-05.999999", person.GraduatedAt3.String())
	assert.Equal(t, "2020-08-05.999999999", person.GraduatedAt4.String())

	assert.Equal(t, "13:14:15", person.OperatedAt1.String())
	assert.Equal(t, "13:14:15.999", person.OperatedAt2.String())
	assert.Equal(t, "13:14:15.999999", person.OperatedAt3.String())
	assert.Equal(t, "13:14:15.999999999", person.OperatedAt4.String())

	assert.Equal(t, "2020-08-05", person.CreatedAt1.String())
	assert.Equal(t, "13:14:15", person.CreatedAt2.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.CreatedAt3.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.CreatedAt4.String())

	fmt.Printf("person output by json:\n%s\n", data)
}

func TestCarbon_MarshalJSON_UnLoadTag(t *testing.T) {
	c := Parse("2020-08-05 13:14:15.999999999", PRC)
	student := Student{
		Name:     "gouguoyin",
		Age:      18,
		Birthday: c,
	}

	data, marshalErr := json.Marshal(&student)
	assert.Nil(t, marshalErr)

	assert.Equal(t, "2020-08-05 13:14:15", student.Birthday.String())
	fmt.Printf("student output by json:\n%s\n", data)
}

func TestCarbon_UnmarshalJSON_LoadTag(t *testing.T) {
	str := `{
		"name": "gouguoyin",
		"age": 18,
		"birthday1": "2020-08-05 13:14:15",
		"birthday2": "2020-08-05 13:14:15.999",
		"birthday3": "2020-08-05 13:14:15.999999",
		"birthday4": "2020-08-05 13:14:15.999999999",
		"graduated_at1": "2020-08-05",
		"graduated_at2": "2020-08-05.999",
		"graduated_at3": "2020-08-05.999999",
		"graduated_at4": "2020-08-05.999999999",
		"operated_at1": "13:14:15",
		"operated_at2": "13:14:15.999",
		"operated_at3": "13:14:15.999999",
		"operated_at4": "13:14:15.999999999",
		"created_at1": "2020-08-05",
		"created_at2": "13:14:15",
		"created_at3": "2020-08-05 13:14:15",
		"created_at4": "2020-08-05 13:14:15"
	}`

	var person Person

	loadErr := LoadTag(&person)
	assert.Nil(t, loadErr)

	unmarshalErr := json.Unmarshal([]byte(str), &person)
	assert.Nil(t, unmarshalErr)

	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday1.String())
	assert.Equal(t, "2020-08-05 13:14:15.999", person.Birthday2.String())
	assert.Equal(t, "2020-08-05 13:14:15.999999", person.Birthday3.String())
	assert.Equal(t, "2020-08-05 13:14:15.999999999", person.Birthday4.String())

	assert.Equal(t, "2020-08-05", person.GraduatedAt1.String())
	assert.Equal(t, "2020-08-05.999", person.GraduatedAt2.String())
	assert.Equal(t, "2020-08-05.999999", person.GraduatedAt3.String())
	assert.Equal(t, "2020-08-05.999999999", person.GraduatedAt4.String())

	assert.Equal(t, "13:14:15", person.OperatedAt1.String())
	assert.Equal(t, "13:14:15.999", person.OperatedAt2.String())
	assert.Equal(t, "13:14:15.999999", person.OperatedAt3.String())
	assert.Equal(t, "13:14:15.999999999", person.OperatedAt4.String())

	assert.Equal(t, "2020-08-05", person.CreatedAt1.String())
	assert.Equal(t, "13:14:15", person.CreatedAt2.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.CreatedAt3.String())

	fmt.Printf("Json string parse to person:\n%+v\n", person)
}

func TestCarbon_UnmarshalJSON_UnLoadTag(t *testing.T) {
	str := `{
		"name": "gouguoyin",
		"age": 18,
		"birthday": "2020-08-05 13:14:15.999"
	}`

	var student Student
	unmarshalErr := json.Unmarshal([]byte(str), &student)
	assert.Nil(t, unmarshalErr)

	assert.Equal(t, "2020-08-05 13:14:15", student.Birthday.String())
	fmt.Printf("Json string parse to student:\n%+v\n", student)
}

func TestError_JSON(t *testing.T) {
	student := Student{
		Name:     "gouguoyin",
		Age:      18,
		Birthday: Parse("XXX"),
	}

	_, marshalErr := json.Marshal(&student)
	assert.NotNil(t, marshalErr)

	str := `{
		"name": "gouguoyin",
		"age": 18,
		"birthday": "2020-08-05 13:14:15.999"
	}`
	unmarshalErr := json.Unmarshal([]byte(str), &student)
	assert.NotNil(t, unmarshalErr)
}
