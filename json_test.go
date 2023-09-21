package carbon

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name         string         `json:"name"`
	Age          int            `json:"age"`
	Birthday1    DateTime       `json:"birthday1"`
	Birthday2    DateTimeMilli  `json:"birthday2"`
	Birthday3    DateTimeMicro  `json:"birthday3"`
	Birthday4    DateTimeNano   `json:"birthday4"`
	GraduatedAt1 Date           `json:"graduated_at1"`
	GraduatedAt2 DateMilli      `json:"graduated_at2"`
	GraduatedAt3 DateMicro      `json:"graduated_at3"`
	GraduatedAt4 DateNano       `json:"graduated_at4"`
	CreatedAt1   Timestamp      `json:"created_at1"`
	CreatedAt2   TimestampMilli `json:"created_at2"`
	CreatedAt3   TimestampMicro `json:"created_at3"`
	CreatedAt4   TimestampNano  `json:"created_at4"`
}

var person Person

func TestCarbon_MarshalJSON(t *testing.T) {
	testNow := SetTestNow(Parse("2020-08-05 13:14:15.999999999", PRC))
	person = Person{
		Name:         "gouguoyin",
		Age:          18,
		Birthday1:    testNow.Now().SubYears(18).ToDateTimeStruct(),
		Birthday2:    testNow.Now().SubYears(18).ToDateTimeMilliStruct(),
		Birthday3:    testNow.Now().SubYears(18).ToDateTimeMicroStruct(),
		Birthday4:    testNow.Now().SubYears(18).ToDateTimeNanoStruct(),
		GraduatedAt1: Parse("2020-08-05 13:14:15", PRC).ToDateStruct(),
		GraduatedAt2: Parse("2020-08-05 13:14:15.999", PRC).ToDateMilliStruct(),
		GraduatedAt3: Parse("2020-08-05 13:14:15.999999", PRC).ToDateMicroStruct(),
		GraduatedAt4: Parse("2020-08-05 13:14:15.999999999", PRC).ToDateNanoStruct(),
		CreatedAt1:   Parse("2023-08-05 13:14:15", PRC).ToTimestampStruct(),
		CreatedAt2:   Parse("2024-08-05 13:14:15.999", PRC).ToTimestampMilliStruct(),
		CreatedAt3:   Parse("2025-08-05 13:14:15.999999", PRC).ToTimestampMicroStruct(),
		CreatedAt4:   Parse("2025-08-05 13:14:15.999999999", PRC).ToTimestampNanoStruct(),
	}
	data, err := json.Marshal(&person)
	assert.Nil(t, err)

	assert.Equal(t, "2002-08-05 13:14:15", person.Birthday1.String(), "birthday1 should be \"2005-09-11 09:57:38\"")
	assert.Equal(t, "2002-08-05 13:14:15.999", person.Birthday2.String(), "birthday2 should be \"2002-08-05 13:14:15.999\"")
	assert.Equal(t, "2002-08-05 13:14:15.999999", person.Birthday3.String(), "birthday3 should be \"2002-08-05 13:14:15.999999\"")
	assert.Equal(t, "2002-08-05 13:14:15.999999999", person.Birthday4.String(), "birthday4 should be \"2002-08-05 13:14:15.999999999\"")
	assert.Equal(t, "2020-08-05", person.GraduatedAt1.String(), "graduated_at1 should be \"2020-08-05\"")
	assert.Equal(t, "2020-08-05.999", person.GraduatedAt2.String(), "graduated_at2 should be \"2020-08-05.999\"")
	assert.Equal(t, "2020-08-05.999999", person.GraduatedAt3.String(), "graduated_at3 should be \"2020-08-05.999999\"")
	assert.Equal(t, "2020-08-05.999999999", person.GraduatedAt4.String(), "graduated_at4 should be \"2020-08-05.999999999\"")

	assert.Equal(t, "1691212455", person.CreatedAt1.String(), "created_at1 should be \"1691212455\"")
	assert.Equal(t, "1722834855999", person.CreatedAt2.String(), "created_at2 should be \"1722834855999\"")
	assert.Equal(t, "1754370855999999", person.CreatedAt3.String(), "created_at3 should be `\"1754370855999999\"")
	assert.Equal(t, "1754370855999999999", person.CreatedAt4.String(), "created_at4 should be \"1754370855999999999\"")
	assert.Equal(t, int64(1691212455), person.CreatedAt1.Int64(), "created_at1 should be 1691212455")
	assert.Equal(t, int64(1722834855999), person.CreatedAt2.Int64(), "created_at2 should be 1722834855999")
	assert.Equal(t, int64(1754370855999999), person.CreatedAt3.Int64(), "created_at3 should be 1754370855999999")
	assert.Equal(t, int64(1754370855999999999), person.CreatedAt4.Int64(), "created_at4 should be 1754370855999999999")

	fmt.Printf("Person output by json:\n%s\n", data)
}

func TestCarbon_UnmarshalJSON(t *testing.T) {
	str := `{
		"name": "gouguoyin",
		"age": 18,
		"birthday1": "2002-08-05 13:14:15",
		"birthday2": "2002-08-05 13:14:15.999",
		"birthday3": "2002-08-05 13:14:15.999999",
		"birthday4": "2002-08-05 13:14:15.999999999",
		"graduated_at1": "2020-08-05",
		"graduated_at2": "2020-08-05.999",
		"graduated_at3": "2020-08-05.999999",
		"graduated_at4": "2020-08-05.999999999",
		"created_at1": 1596604455,
		"created_at2": 1596604455999,
		"created_at3": 1596604455999999,
		"created_at4": 1596604455999999999
	}`

	err := json.Unmarshal([]byte(str), &person)
	assert.Nil(t, err)

	assert.Equal(t, "2002-08-05 13:14:15", person.Birthday1.String(), "birthday1 should be \"2002-08-05 13:14:15\"")
	assert.Equal(t, "2002-08-05 13:14:15.999", person.Birthday2.String(), "birthday2 should be \"2002-08-05 13:14:15.999\"")
	assert.Equal(t, "2002-08-05 13:14:15.999999", person.Birthday3.String(), "birthday3 should be \"2002-08-05 13:14:15.999999\"")
	assert.Equal(t, "2002-08-05 13:14:15.999999999", person.Birthday4.String(), "birthday4 should be \"2002-08-05 13:14:15.999999999\"")
	assert.Equal(t, "2020-08-05", person.GraduatedAt1.String(), "graduated_at1 should be \"2020-08-05\"")
	assert.Equal(t, "2020-08-05.999", person.GraduatedAt2.String(), "graduated_at2 should be \"2020-08-05.999\"")
	assert.Equal(t, "2020-08-05.999999", person.GraduatedAt3.String(), "graduated_at3 should be \"2020-08-05.999999\"")
	assert.Equal(t, "2020-08-05.999999999", person.GraduatedAt4.String(), "graduated_at4 should be \"2020-08-05.999999999\"")

	assert.Equal(t, "1596604455", person.CreatedAt1.String(), "created_at1 should be \"1596604455\"")
	assert.Equal(t, "1596604455999", person.CreatedAt2.String(), "created_at2 should be \"1596604455999\"")
	assert.Equal(t, "1596604455999999", person.CreatedAt3.String(), "created_at2 should be `\"1596604455999999\"")
	assert.Equal(t, "1596604455999999999", person.CreatedAt4.String(), "created_at2 should be \"1596604455999999999\"")
	assert.Equal(t, int64(1596604455), person.CreatedAt1.Int64(), "created_at1 should be 1596604455")
	assert.Equal(t, int64(1596604455999), person.CreatedAt2.Int64(), "created_at2 should be 1596604455999")
	assert.Equal(t, int64(1596604455999999), person.CreatedAt3.Int64(), "created_at2 should be 1596604455999999")
	assert.Equal(t, int64(1596604455999999999), person.CreatedAt4.Int64(), "created_at2 should be 1596604455999999999")

	fmt.Printf("Json string parse to person:\n%+v\n", person)
}

func TestCarbon_GormDataType(t *testing.T) {
	var dateTime DateTime
	assert.Equal(t, "time", dateTime.GormDataType())
	var dateTimeMilli DateTimeMilli
	assert.Equal(t, "time", dateTimeMilli.GormDataType())
	var dateTimeMicro DateTimeMicro
	assert.Equal(t, "time", dateTimeMicro.GormDataType())
	var dateTimeNano DateTimeNano
	assert.Equal(t, "time", dateTimeNano.GormDataType())

	var date Date
	assert.Equal(t, "time", date.GormDataType())
	var dateMilli DateMilli
	assert.Equal(t, "time", dateMilli.GormDataType())
	var dateMicro DateMicro
	assert.Equal(t, "time", dateMicro.GormDataType())
	var dateNano DateNano
	assert.Equal(t, "time", dateNano.GormDataType())

	var timestamp Timestamp
	assert.Equal(t, "int", timestamp.GormDataType())
	var timestampMilli TimestampMilli
	assert.Equal(t, "int", timestampMilli.GormDataType())
	var timestampMicro TimestampMicro
	assert.Equal(t, "int", timestampMicro.GormDataType())
	var timestampNano TimestampNano
	assert.Equal(t, "int", timestampNano.GormDataType())
}

func TestError_Json(t *testing.T) {
	str := `{
		"name": "",
		"age": 0,
		"birthday1": "",
		"birthday2": "",
		"birthday3": "",
		"birthday4": "",
		"graduated_at1": "xxx",
		"graduated_at2": "xxx",
		"graduated_at3": "xxx",
		"graduated_at4": "xxx",
		"created_at1": 0,
		"created_at2": 0,
		"created_at3": 0,
		"created_at4": 0
	}`
	err := json.Unmarshal([]byte(str), &person)
	assert.NotNil(t, err)
	fmt.Printf("Json string parse to person:\n%+v\n", person)
}
