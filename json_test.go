package carbon

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Person struct {
	ID          int64                      `json:"id"`
	Name        string                     `json:"name"`
	Age         int                        `json:"age"`
	Birthday    ToDateTimeString           `json:"birthday"`
	GraduatedAt ToDateString               `json:"graduated_at"`
	CreatedAt   ToTimeString               `json:"created_at"`
	UpdatedAt   ToTimestamp                `json:"updated_at"`
	DateTime1   ToTimestampWithSecond      `json:"date_time1"`
	DateTime2   ToTimestampWithMillisecond `json:"date_time2"`
	DateTime3   ToTimestampWithMicrosecond `json:"date_time3"`
	DateTime4   ToTimestampWithNanosecond  `json:"date_time4"`
}

func TestCarbon_MarshalJSON1(t *testing.T) {
	person := Person{
		ID:          153,
		Name:        "勾国印",
		Age:         18,
		Birthday:    ToDateTimeString{Now().SubYears(18)},
		GraduatedAt: ToDateString{Parse("2020-08-05 13:14:15")},
		CreatedAt:   ToTimeString{Parse("2021-08-05 13:14:15")},
		UpdatedAt:   ToTimestamp{Parse("2022-08-05 13:14:15")},
		DateTime1:   ToTimestampWithSecond{Parse("2023-08-05 13:14:15")},
		DateTime2:   ToTimestampWithMillisecond{Parse("2024-08-05 13:14:15")},
		DateTime3:   ToTimestampWithMicrosecond{Parse("2025-08-05 13:14:15")},
		DateTime4:   ToTimestampWithNanosecond{Parse("2025-08-05 13:14:15")},
	}
	data, err := json.Marshal(&person)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Person1 output by json:\n%s\n", data)
}

func TestCarbon_MarshalJSON2(t *testing.T) {
	person := Person{
		ID:          153,
		Name:        "勾国印",
		Age:         18,
		Birthday:    ToDateTimeString{Parse("")},
		GraduatedAt: ToDateString{Parse("0")},
		CreatedAt:   ToTimeString{Parse("0000-00-00")},
		UpdatedAt:   ToTimestamp{Parse("00:00:00")},
		DateTime1:   ToTimestampWithSecond{NewCarbon()},
		DateTime2:   ToTimestampWithMillisecond{Carbon{}},
		DateTime3:   ToTimestampWithMicrosecond{Parse("")},
		DateTime4:   ToTimestampWithNanosecond{Parse("")},
	}
	data, err := json.Marshal(&person)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Person2 output by json:\n%s\n", data)
}

func TestCarbon_UnmarshalJSON1(t *testing.T) {
	str := `{
		"id":153,
		"name":"勾国印",
		"age":18,
		"birthday":"2003-07-16 16:22:02",
		"graduated_at":"2020-08-05",
		"created_at":"13:14:15",
		"updated_at":1659676455,
		"date_time1":1691212455,
		"date_time2":1722834855000,
		"date_time3":1754370855000000,
		"date_time4":1754370855000000000
	}`
	person := new(Person)
	err := json.Unmarshal([]byte(str), &person)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Json string parse to person1:\n%+v\n", *person)
}

func TestCarbon_UnmarshalJSON2(t *testing.T) {
	str := `{
		"id":0,
		"name":"",
		"age":0,
		"birthday":"",
		"graduated_at":"",
		"created_at":"",
		"updated_at":0,
		"date_time1":0,
		"date_time2":0,
		"date_time3":0,
		"date_time4":0
	}`
	person := new(Person)
	err := json.Unmarshal([]byte(str), &person)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Json string parse to person2:\n%+v\n", *person)
}

// 错误情况
func TestCarbon_UnmarshalJSON3(t *testing.T) {
	str := `xxxx`
	person := new(Person)
	err := json.Unmarshal([]byte(str), &person)
	if err != nil {
		fmt.Println("catch an exception in UnmarshalJSON():", err)
	}
}
