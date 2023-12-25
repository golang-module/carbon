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

	Birthday0 Carbon `json:"birthday0"`

	Birthday1 Carbon `json:"birthday1" carbon:"layout:2006-01-02"`
	Birthday2 Carbon `json:"birthday2" carbon:"layout:15:04:05"`
	Birthday3 Carbon `json:"birthday3" carbon:"layout:2006-01-02 15:04:05"`

	Birthday4 Carbon `json:"birthday4" carbon:"format:Y-m-d"`
	Birthday5 Carbon `json:"birthday5" carbon:"format:H:i:s"`
	Birthday6 Carbon `json:"birthday6" carbon:"format:Y-m-d H:i:s"`
}

func TestCarbon_MarshalJSON_LoadTag(t *testing.T) {
	c := Parse("2020-08-05 13:14:15.999999999", PRC)
	person := Person{
		Name:      "gouguoyin",
		Age:       18,
		Birthday0: c,
		Birthday1: c,
		Birthday2: c,
		Birthday3: c,
		Birthday4: c,
		Birthday5: c,
		Birthday6: c,
	}

	loadErr := LoadTag(&person)
	assert.Nil(t, loadErr)

	data, marshalErr := json.Marshal(&person)
	assert.Nil(t, marshalErr)

	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday0.String())

	assert.Equal(t, "2020-08-05", person.Birthday1.String())
	assert.Equal(t, "13:14:15", person.Birthday2.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday3.String())

	assert.Equal(t, "2020-08-05", person.Birthday4.String())
	assert.Equal(t, "13:14:15", person.Birthday5.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday6.String())

	fmt.Printf("person output by json:\n%s\n", data)
}

func TestCarbon_MarshalJSON_UnLoadTag(t *testing.T) {
	c := Parse("2020-08-05 13:14:15.999999999", PRC)
	person := Person{
		Name:      "gouguoyin",
		Age:       18,
		Birthday1: c,
		Birthday2: c,
		Birthday3: c,
		Birthday4: c,
		Birthday5: c,
		Birthday6: c,
	}

	data, marshalErr := json.Marshal(&person)
	assert.Nil(t, marshalErr)

	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday1.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday2.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday3.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday4.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday5.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday6.String())
	fmt.Printf("student output by json:\n%s\n", data)
}

func TestCarbon_UnmarshalJSON_LoadTag(t *testing.T) {
	str := `{
		"name":"gouguoyin",
		"age":18,
		"birthday0":"2020-08-05 13:14:15",
		"birthday1":"2020-08-05",
		"birthday2":"13:14:15",
		"birthday3":"2020-08-05 13:14:15",
		"birthday4":"2020-08-05",
		"birthday5":"13:14:15",
		"birthday6":"2020-08-05 13:14:15"
	}`

	var person Person

	loadErr := LoadTag(&person)
	assert.Nil(t, loadErr)

	unmarshalErr := json.Unmarshal([]byte(str), &person)
	assert.Nil(t, unmarshalErr)

	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday0.String())
	assert.Equal(t, "2020-08-05", person.Birthday1.String())
	assert.Equal(t, "13:14:15", person.Birthday2.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday3.String())
	assert.Equal(t, "2020-08-05", person.Birthday4.String())
	assert.Equal(t, "13:14:15", person.Birthday5.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday6.String())

	fmt.Printf("Json string parse to person:\n%+v\n", person)
}

func TestCarbon_UnmarshalJSON_UnLoadTag(t *testing.T) {
	str := `{
		"name":"gouguoyin",
		"age":18,
		"birthday0":"2020-08-05 13:14:15",
		"birthday1":"2020-08-05"
	}`

	var person Person
	unmarshalErr := json.Unmarshal([]byte(str), &person)
	assert.Nil(t, unmarshalErr)

	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday0.String())
	assert.Equal(t, "2020-08-05 00:00:00", person.Birthday1.String())
	fmt.Printf("Json string parse to person:\n%+v\n", person)
}

func TestError_JSON(t *testing.T) {
	person := Person{
		Name:      "gouguoyin",
		Age:       18,
		Birthday1: Parse("XXX"),
	}

	_, marshalErr := json.Marshal(person)
	fmt.Println("marshal error:", marshalErr.Error())
	assert.NotNil(t, marshalErr)

	str := `{
		"name": "gouguoyin",
		"age": 18,
		"birthday1": "2020-08-05 13:14:15"
	}`
	unmarshalErr := json.Unmarshal([]byte(str), &person)
	fmt.Println("unmarshal error:", unmarshalErr.Error())
	assert.NotNil(t, unmarshalErr)
}
