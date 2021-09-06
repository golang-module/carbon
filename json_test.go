package carbon

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	ID          int64                    `json:"id"`
	Name        string                   `json:"name"`
	Age         int                      `json:"age"`
	Birthday    DateTime                 `json:"birthday"`
	GraduatedAt Date                     `json:"graduated_at"`
	CreatedAt   Time                     `json:"created_at"`
	UpdatedAt   Timestamp                `json:"updated_at"`
	DateTime1   TimestampWithSecond      `json:"date_time1"`
	DateTime2   TimestampWithMillisecond `json:"date_time2"`
	DateTime3   TimestampWithMicrosecond `json:"date_time3"`
	DateTime4   TimestampWithNanosecond  `json:"date_time4"`
}

func TestCarbon_MarshalJSON(t *testing.T) {
	assert := assert.New(t)

	person := Person{
		ID:          1,
		Name:        "gouguoyin",
		Age:         18,
		Birthday:    DateTime{Now().SubYears(18)},
		GraduatedAt: Date{Parse("2020-08-05 13:14:15")},
		CreatedAt:   Time{Parse("2021-08-05 13:14:15")},
		UpdatedAt:   Timestamp{Parse("2022-08-05 13:14:15")},
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

func TestCarbon_UnmarshalJSON(t *testing.T) {
	assert := assert.New(t)

	person := new(Person)

	str1 := `{
		"id": 1,
		"name": "gouguoyin",
		"age": 18,
		"birthday": "2003-07-16 16:22:02",
		"graduated_at": "2020-08-05",
		"created_at": "13:14:15",
		"updated_at": 1659676455,
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
		"id": 0,
		"name": "",
		"age": 0,
		"birthday": "",
		"graduated_at": "",
		"created_at": "",
		"updated_at": 0,
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
