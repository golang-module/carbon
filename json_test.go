package carbon

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name        string         `json:"name"`
	Age         int            `json:"age"`
	Birthday    DateTime       `json:"birthday"`
	GraduatedAt Date           `json:"graduated_at"`
	DateTime1   Timestamp      `json:"date_time1"`
	DateTime2   TimestampMilli `json:"date_time2"`
	DateTime3   TimestampMicro `json:"date_time3"`
	DateTime4   TimestampNano  `json:"date_time4"`
}

func TestCarbon_MarshalJSON(t *testing.T) {
	person := Person{
		Name:        "gouguoyin",
		Age:         18,
		Birthday:    DateTime{Now().SubYears(18)},
		GraduatedAt: Date{Parse("2020-08-05 13:14:15")},
		DateTime1:   Timestamp{Parse("2023-08-05 13:14:15")},
		DateTime2:   TimestampMilli{Parse("2024-08-05 13:14:15")},
		DateTime3:   TimestampMicro{Parse("2025-08-05 13:14:15")},
		DateTime4:   TimestampNano{Parse("2025-08-05 13:14:15")},
	}
	data, err := json.Marshal(&person)
	if err != nil {
		assert.NotNil(t, err)
	}
	fmt.Printf("Person output by json:\n%s\n", data)
}

func TestCarbon_UnmarshalJSON(t *testing.T) {
	person := new(Person)
	str := `{
		"name": "gouguoyin",
		"age": 18,
		"birthday": "2003-07-16 16:22:02",
		"graduated_at": "2020-08-05",
		"date_time1": 1691212455,
		"date_time2": 1722834855000,
		"date_time3": 1754370855000000,
		"date_time4": 1754370855000000000
	}`

	err := json.Unmarshal([]byte(str), &person)
	if err != nil {
		assert.NotNil(t, err)
	}
	fmt.Printf("Json string parse to person:\n%+v\n", *person)

}

func TestError_Json(t *testing.T) {
	person := new(Person)
	str := `{
		"name": "",
		"age": 0,
		"birthday": "",
		"graduated_at": "xxx",
		"date_time1": 0,
		"date_time2": 0,
		"date_time3": 0,
		"date_time4": 0
	}`
	err := json.Unmarshal([]byte(str), &person)
	if err != nil {
		assert.NotNil(t, err)
	}
	fmt.Printf("Json string parse to person:\n%+v\n", *person)
}
