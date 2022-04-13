package carbon

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person1 struct {
	Name        string                   `json:"name"`
	Age         int                      `json:"age"`
	Birthday    DateTime                 `json:"birthday"`
	GraduatedAt Date                     `json:"graduated_at"`
	DateTime1   TimestampWithSecond      `json:"date_time1"`
	DateTime2   TimestampWithMillisecond `json:"date_time2"`
	DateTime3   TimestampWithMicrosecond `json:"date_time3"`
	DateTime4   TimestampWithNanosecond  `json:"date_time4"`
}

type Person2 struct {
	Name        string         `json:"name"`
	Age         int            `json:"age"`
	Birthday    DateTime       `json:"birthday"`
	GraduatedAt Date           `json:"graduated_at"`
	DateTime1   Timestamp      `json:"date_time1"`
	DateTime2   TimestampMilli `json:"date_time2"`
	DateTime3   TimestampMicro `json:"date_time3"`
	DateTime4   TimestampNano  `json:"date_time4"`
}

func TestCarbon_MarshalJSON1(t *testing.T) {
	assert := assert.New(t)

	person := Person1{
		Name:        "gouguoyin",
		Age:         18,
		Birthday:    DateTime{Now().SubYears(18)},
		GraduatedAt: Date{Parse("2020-08-05 13:14:15")},
		DateTime1:   TimestampWithSecond{Parse("2023-08-05 13:14:15")},
		DateTime2:   TimestampWithMillisecond{Parse("2024-08-05 13:14:15")},
		DateTime3:   TimestampWithMicrosecond{Parse("2025-08-05 13:14:15")},
		DateTime4:   TimestampWithNanosecond{Parse("2025-08-05 13:14:15")},
	}
	data, err := json.Marshal(&person)
	if err != nil {
		assert.NotNil(err)
	}
	fmt.Printf("Person output by json:\n%s\n", data)
}

func TestCarbon_MarshalJSON2(t *testing.T) {
	assert := assert.New(t)

	person := Person2{
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
		assert.NotNil(err)
	}
	fmt.Printf("Person output by json:\n%s\n", data)
}

func TestCarbon_UnmarshalJSON1(t *testing.T) {
	assert := assert.New(t)

	person := new(Person1)

	str1 := `{
		"name": "gouguoyin",
		"age": 18,
		"birthday": "2003-07-16 16:22:02",
		"graduated_at": "2020-08-05",
		"date_time1": 1691212455,
		"date_time2": 1722834855000,
		"date_time3": 1754370855000000,
		"date_time4": 1754370855000000000
	}`

	err1 := json.Unmarshal([]byte(str1), &person)
	if err1 != nil {
		assert.NotNil(err1)
	}
	fmt.Printf("Json string parse to person1:\n%+v\n", *person)

	str2 := `{
		"name": "",
		"age": 0,
		"birthday": "",
		"graduated_at": "",
		"date_time1": 0,
		"date_time2": 0,
		"date_time3": 0,
		"date_time4": 0
	}`
	err2 := json.Unmarshal([]byte(str2), &person)
	if err2 != nil {
		assert.NotNil(err2)
	}
	fmt.Printf("Json string parse to person2:\n%+v\n", *person)
}

func TestCarbon_UnmarshalJSON2(t *testing.T) {
	assert := assert.New(t)

	person := new(Person2)

	str1 := `{
		"name": "gouguoyin",
		"age": 18,
		"birthday": "2003-07-16 16:22:02",
		"graduated_at": "2020-08-05",
		"date_time1": 1691212455,
		"date_time2": 1722834855000,
		"date_time3": 1754370855000000,
		"date_time4": 1754370855000000000
	}`

	err1 := json.Unmarshal([]byte(str1), &person)
	if err1 != nil {
		assert.NotNil(err1)
	}
	fmt.Printf("Json string parse to person1:\n%+v\n", *person)

	str2 := `{
		"name": "",
		"age": 0,
		"birthday": "",
		"graduated_at": "",
		"date_time1": 0,
		"date_time2": 0,
		"date_time3": 0,
		"date_time4": 0
	}`
	err2 := json.Unmarshal([]byte(str2), &person)
	if err2 != nil {
		assert.NotNil(err2)
	}
	fmt.Printf("Json string parse to person2:\n%+v\n", *person)
}
