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

	Birthday01 Carbon `json:"birthday01"`
	Birthday02 Carbon `json:"birthday02" carbon:"type:dateTime"`
	Birthday03 Carbon `json:"birthday03" tz:"PRC"`
	Birthday04 Carbon `json:"birthday04" carbon:"" tz:""`

	Birthday1 Carbon `json:"birthday1" carbon:"layout:2006-01-02"`
	Birthday2 Carbon `json:"birthday2" carbon:"layout:15:04:05"`
	Birthday3 Carbon `json:"birthday3" carbon:"layout:2006-01-02 15:04:05"`
	Birthday4 Carbon `json:"birthday4" carbon:"layout:2006-01-02" tz:"PRC"`
	Birthday5 Carbon `json:"birthday5" carbon:"layout:15:04:05" tz:"PRC"`
	Birthday6 Carbon `json:"birthday6" carbon:"layout:2006-01-02 15:04:05" tz:"PRC"`

	Birthday7  Carbon `json:"birthday7" carbon:"format:Y-m-d"`
	Birthday8  Carbon `json:"birthday8" carbon:"format:H:i:s"`
	Birthday9  Carbon `json:"birthday9" carbon:"format:Y-m-d H:i:s"`
	Birthday10 Carbon `json:"birthday10" carbon:"format:Y-m-d" tz:"PRC"`
	Birthday11 Carbon `json:"birthday11" carbon:"format:H:i:s" tz:"PRC"`
	Birthday12 Carbon `json:"birthday12" carbon:"format:Y-m-d H:i:s" tz:"PRC"`

	Birthday13 Carbon `json:"birthday13" carbon:"type:dateTime"`
	Birthday14 Carbon `json:"birthday14" carbon:"type:dateTimeMilli"`
	Birthday15 Carbon `json:"birthday15" carbon:"type:dateTimeMicro"`
	Birthday16 Carbon `json:"birthday16" carbon:"type:dateTimeNano"`
	Birthday17 Carbon `json:"birthday17" carbon:"type:shortDateTime"`
	Birthday18 Carbon `json:"birthday18" carbon:"type:shortDateTimeMilli"`
	Birthday19 Carbon `json:"birthday19" carbon:"type:shortDateTimeMicro"`
	Birthday20 Carbon `json:"birthday20" carbon:"type:shortDateTimeNano"`
	Birthday21 Carbon `json:"birthday21" carbon:"type:dayDateTime"`
	Birthday22 Carbon `json:"birthday22" carbon:"type:date"`
	Birthday23 Carbon `json:"birthday23" carbon:"type:dateMilli"`
	Birthday24 Carbon `json:"birthday24" carbon:"type:dateMicro"`
	Birthday25 Carbon `json:"birthday25" carbon:"type:dateNano"`
	Birthday26 Carbon `json:"birthday26" carbon:"type:shortDate"`
	Birthday27 Carbon `json:"birthday27" carbon:"type:shortDateMilli"`
	Birthday28 Carbon `json:"birthday28" carbon:"type:shortDateMicro"`
	Birthday29 Carbon `json:"birthday29" carbon:"type:shortDateNano"`
	Birthday30 Carbon `json:"birthday30" carbon:"type:time"`
	Birthday31 Carbon `json:"birthday31" carbon:"type:timeMilli"`
	Birthday32 Carbon `json:"birthday32" carbon:"type:timeMicro"`
	Birthday33 Carbon `json:"birthday33" carbon:"type:timeNano"`
	Birthday34 Carbon `json:"birthday34" carbon:"type:shortTime"`
	Birthday35 Carbon `json:"birthday35" carbon:"type:shortTimeMilli"`
	Birthday36 Carbon `json:"birthday36" carbon:"type:shortTimeMicro"`
	Birthday37 Carbon `json:"birthday37" carbon:"type:shortTimeNano"`
	Birthday38 Carbon `json:"birthday38" carbon:"type:atom"`
	Birthday39 Carbon `json:"birthday39" carbon:"type:ansic"`
	Birthday40 Carbon `json:"birthday40" carbon:"type:cookie"`
	Birthday41 Carbon `json:"birthday41" carbon:"type:kitchen"`
	Birthday42 Carbon `json:"birthday42" carbon:"type:rss"`
	Birthday43 Carbon `json:"birthday43" carbon:"type:rubyDate"`
	Birthday44 Carbon `json:"birthday44" carbon:"type:unixDate"`
	Birthday45 Carbon `json:"birthday45" carbon:"type:rfc1036"`
	Birthday46 Carbon `json:"birthday46" carbon:"type:rfc1123"`
	Birthday47 Carbon `json:"birthday47" carbon:"type:rfc1123Z"`
	Birthday48 Carbon `json:"birthday48" carbon:"type:rfc2822"`
	Birthday49 Carbon `json:"birthday49" carbon:"type:rfc3339"`
	Birthday50 Carbon `json:"birthday50" carbon:"type:rfc3339Milli"`
	Birthday51 Carbon `json:"birthday51" carbon:"type:rfc3339Micro"`
	Birthday52 Carbon `json:"birthday52" carbon:"type:rfc3339Nano"`
	Birthday53 Carbon `json:"birthday53" carbon:"type:rfc7231"`
	Birthday54 Carbon `json:"birthday54" carbon:"type:rfc822"`
	Birthday55 Carbon `json:"birthday55" carbon:"type:rfc822Z"`
	Birthday56 Carbon `json:"birthday56" carbon:"type:rfc850"`
	Birthday57 Carbon `json:"birthday57" carbon:"type:iso8601"`
	Birthday58 Carbon `json:"birthday58" carbon:"type:iso8601Milli"`
	Birthday59 Carbon `json:"birthday59" carbon:"type:iso8601Micro"`
	Birthday60 Carbon `json:"birthday60" carbon:"type:iso8601Nano"`

	Birthday61 Carbon `json:"birthday61" carbon:"type:timestamp"`
	Birthday62 Carbon `json:"birthday62" carbon:"type:timestampMilli"`
	Birthday63 Carbon `json:"birthday63" carbon:"type:timestampMicro"`
	Birthday64 Carbon `json:"birthday64" carbon:"type:timestampNano"`
}

var c = Parse("2020-08-05 13:14:15.999999999", PRC)
var person = Person{
	Name:       "gouguoyin",
	Age:        18,
	Birthday01: c,
	Birthday02: c,
	Birthday03: c,
	Birthday04: c,
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
	Birthday18: c,
	Birthday19: c,
	Birthday20: c,
	Birthday21: c,
	Birthday22: c,
	Birthday23: c,
	Birthday24: c,
	Birthday25: c,
	Birthday26: c,
	Birthday27: c,
	Birthday28: c,
	Birthday29: c,
	Birthday30: c,
	Birthday31: c,
	Birthday32: c,
	Birthday33: c,
	Birthday34: c,
	Birthday35: c,
	Birthday36: c,
	Birthday37: c,
	Birthday38: c,
	Birthday39: c,
	Birthday40: c,
	Birthday41: c,
	Birthday42: c,
	Birthday43: c,
	Birthday44: c,
	Birthday45: c,
	Birthday46: c,
	Birthday47: c,
	Birthday48: c,
	Birthday49: c,
	Birthday50: c,
	Birthday51: c,
	Birthday52: c,
	Birthday53: c,
	Birthday54: c,
	Birthday55: c,
	Birthday56: c,
	Birthday57: c,
	Birthday58: c,
	Birthday59: c,
	Birthday60: c,
	Birthday61: c,
	Birthday62: c,
	Birthday63: c,
	Birthday64: c,
}

func TestCarbon_MarshalJSON(t *testing.T) {
	loadErr := LoadTag(&person)
	assert.Nil(t, loadErr)

	data, marshalErr := json.Marshal(&person)
	assert.Nil(t, marshalErr)

	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday01.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday02.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday03.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday04.String())
	assert.Equal(t, "2020-08-05", person.Birthday1.String())
	assert.Equal(t, "13:14:15", person.Birthday2.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday3.String())
	assert.Equal(t, "2020-08-05", person.Birthday4.String())
	assert.Equal(t, "13:14:15", person.Birthday5.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday6.String())
	assert.Equal(t, "2020-08-05", person.Birthday7.String())
	assert.Equal(t, "13:14:15", person.Birthday8.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday9.String())
	assert.Equal(t, "2020-08-05", person.Birthday10.String())
	assert.Equal(t, "13:14:15", person.Birthday11.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday12.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday13.String())
	assert.Equal(t, "2020-08-05 13:14:15.999", person.Birthday14.String())
	assert.Equal(t, "2020-08-05 13:14:15.999999", person.Birthday15.String())
	assert.Equal(t, "2020-08-05 13:14:15.999999999", person.Birthday16.String())
	assert.Equal(t, "20200805131415", person.Birthday17.String())
	assert.Equal(t, "20200805131415.999", person.Birthday18.String())
	assert.Equal(t, "20200805131415.999999", person.Birthday19.String())
	assert.Equal(t, "20200805131415.999999999", person.Birthday20.String())
	assert.Equal(t, "Wed, Aug 5, 2020 1:14 PM", person.Birthday21.String())
	assert.Equal(t, "2020-08-05", person.Birthday22.String())
	assert.Equal(t, "2020-08-05.999", person.Birthday23.String())
	assert.Equal(t, "2020-08-05.999999", person.Birthday24.String())
	assert.Equal(t, "2020-08-05.999999999", person.Birthday25.String())
	assert.Equal(t, "20200805", person.Birthday26.String())
	assert.Equal(t, "20200805.999", person.Birthday27.String())
	assert.Equal(t, "20200805.999999", person.Birthday28.String())
	assert.Equal(t, "20200805.999999999", person.Birthday29.String())
	assert.Equal(t, "13:14:15", person.Birthday30.String())
	assert.Equal(t, "13:14:15.999", person.Birthday31.String())
	assert.Equal(t, "13:14:15.999999", person.Birthday32.String())
	assert.Equal(t, "13:14:15.999999999", person.Birthday33.String())
	assert.Equal(t, "131415", person.Birthday34.String())
	assert.Equal(t, "131415.999", person.Birthday35.String())
	assert.Equal(t, "131415.999999", person.Birthday36.String())
	assert.Equal(t, "131415.999999999", person.Birthday37.String())
	assert.Equal(t, "2020-08-05T13:14:15+08:00", person.Birthday38.String())
	assert.Equal(t, "Wed Aug  5 13:14:15 2020", person.Birthday39.String())
	assert.Equal(t, "Wednesday, 05-Aug-2020 13:14:15 CST", person.Birthday40.String())
	assert.Equal(t, "1:14PM", person.Birthday41.String())
	assert.Equal(t, "Wed, 05 Aug 2020 13:14:15 +0800", person.Birthday42.String())
	assert.Equal(t, "Wed Aug 05 13:14:15 +0800 2020", person.Birthday43.String())
	assert.Equal(t, "Wed Aug  5 13:14:15 CST 2020", person.Birthday44.String())
	assert.Equal(t, "Wed, 05 Aug 20 13:14:15 +0800", person.Birthday45.String())
	assert.Equal(t, "Wed, 05 Aug 2020 13:14:15 CST", person.Birthday46.String())
	assert.Equal(t, "Wed, 05 Aug 2020 13:14:15 +0800", person.Birthday47.String())
	assert.Equal(t, "Wed, 05 Aug 2020 13:14:15 +0800", person.Birthday48.String())
	assert.Equal(t, "2020-08-05T13:14:15+08:00", person.Birthday49.String())
	assert.Equal(t, "2020-08-05T13:14:15.999+08:00", person.Birthday50.String())
	assert.Equal(t, "2020-08-05T13:14:15.999999+08:00", person.Birthday51.String())
	assert.Equal(t, "2020-08-05T13:14:15.999999999+08:00", person.Birthday52.String())
	assert.Equal(t, "Wed, 05 Aug 2020 13:14:15 CST", person.Birthday53.String())
	assert.Equal(t, "05 Aug 20 13:14 CST", person.Birthday54.String())
	assert.Equal(t, "05 Aug 20 13:14 +0800", person.Birthday55.String())
	assert.Equal(t, "Wednesday, 05-Aug-20 13:14:15 CST", person.Birthday56.String())
	assert.Equal(t, "2020-08-05T13:14:15+08:00", person.Birthday57.String())
	assert.Equal(t, "2020-08-05T13:14:15.999+08:00", person.Birthday58.String())
	assert.Equal(t, "2020-08-05T13:14:15.999999+08:00", person.Birthday59.String())
	assert.Equal(t, "2020-08-05T13:14:15.999999999+08:00", person.Birthday60.String())
	assert.Equal(t, "1596604455", person.Birthday61.String())
	assert.Equal(t, "1596604455999", person.Birthday62.String())
	assert.Equal(t, "1596604455999999", person.Birthday63.String())
	assert.Equal(t, "1596604455999999999", person.Birthday64.String())

	fmt.Printf("person output by json:\n%s\n", data)
}

func TestCarbon_UnmarshalJSON(t *testing.T) {
	str := `{
		"name": "gouguoyin",
		"age": 18,
		"birthday01": "2020-08-05 13:14:15",
		"birthday02": "2020-08-05 13:14:15",
		"birthday03": "2020-08-05 13:14:15",
		"birthday04": "2020-08-05 13:14:15",
		"birthday1": "2020-08-05",
		"birthday2": "13:14:15",
		"birthday3": "2020-08-05 13:14:15",
		"birthday4": "2020-08-05",
		"birthday5": "13:14:15",
		"birthday6": "2020-08-05 13:14:15",
		"birthday7": "2020-08-05",
		"birthday8": "13:14:15",
		"birthday9": "2020-08-05 13:14:15",
		"birthday10": "2020-08-05",
		"birthday11": "13:14:15",
		"birthday12": "2020-08-05 13:14:15",
		"birthday13": "2020-08-05 13:14:15",
		"birthday14": "2020-08-05 13:14:15.999",
		"birthday15": "2020-08-05 13:14:15.999999",
		"birthday16": "2020-08-05 13:14:15.999999999",
		"birthday17": "20200805131415",
		"birthday18": "20200805131415.999",
		"birthday19": "20200805131415.999999",
		"birthday20": "20200805131415.999999999",
		"birthday21": "Wed, Aug 5, 2020 1:14 PM",
		"birthday22": "2020-08-05",
		"birthday23": "2020-08-05.999",
		"birthday24": "2020-08-05.999999",
		"birthday25": "2020-08-05.999999999",
		"birthday26": "20200805",
		"birthday27": "20200805.999",
		"birthday28": "20200805.999999",
		"birthday29": "20200805.999999999",
		"birthday30": "13:14:15",
		"birthday31": "13:14:15.999",
		"birthday32": "13:14:15.999999",
		"birthday33": "13:14:15.999999999",
		"birthday34": "131415",
		"birthday35": "131415.999",
		"birthday36": "131415.999999",
		"birthday37": "131415.999999999",
		"birthday38": "2020-08-05T13:14:15+08:00",
		"birthday39": "Wed Aug  5 13:14:15 2020",
		"birthday40": "Wednesday, 05-Aug-2020 13:14:15 CST",
		"birthday41": "1:14PM",
		"birthday42": "Wed, 05 Aug 2020 13:14:15 +0800",
		"birthday43": "Wed Aug 05 13:14:15 +0800 2020",
		"birthday44": "Wed Aug  5 13:14:15 CST 2020",
		"birthday45": "Wed, 05 Aug 20 13:14:15 +0800",
		"birthday46": "Wed, 05 Aug 2020 13:14:15 CST",
		"birthday47": "Wed, 05 Aug 2020 13:14:15 +0800",
		"birthday48": "Wed, 05 Aug 2020 13:14:15 +0800",
		"birthday49": "2020-08-05T13:14:15+08:00",
		"birthday50": "2020-08-05T13:14:15.999+08:00",
		"birthday51": "2020-08-05T13:14:15.999999+08:00",
		"birthday52": "2020-08-05T13:14:15.999999999+08:00",
		"birthday53": "Wed, 05 Aug 2020 13:14:15 CST",
		"birthday54": "05 Aug 20 13:14 CST",
		"birthday55": "05 Aug 20 13:14 +0800",
		"birthday56": "Wednesday, 05-Aug-20 13:14:15 CST",
		"birthday57": "2020-08-05T13:14:15+08:00",
		"birthday58": "2020-08-05T13:14:15.999+08:00",
		"birthday59": "2020-08-05T13:14:15.999999+08:00",
		"birthday60": "2020-08-05T13:14:15.999999999+08:00",
		"birthday61": 1596604455,
		"birthday62": 1596604455999,
		"birthday63": 1596604455999999,
		"birthday64": 1596604455999999999
	}`

	loadErr := LoadTag(&person)
	assert.Nil(t, loadErr)

	unmarshalErr := json.Unmarshal([]byte(str), &person)
	assert.Nil(t, unmarshalErr)

	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday01.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday02.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday03.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday04.String())
	assert.Equal(t, "2020-08-05", person.Birthday1.String())
	assert.Equal(t, "13:14:15", person.Birthday2.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday3.String())
	assert.Equal(t, "2020-08-05", person.Birthday4.String())
	assert.Equal(t, "13:14:15", person.Birthday5.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday6.String())
	assert.Equal(t, "2020-08-05", person.Birthday7.String())
	assert.Equal(t, "13:14:15", person.Birthday8.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday9.String())
	assert.Equal(t, "2020-08-05", person.Birthday10.String())
	assert.Equal(t, "13:14:15", person.Birthday11.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday12.String())
	assert.Equal(t, "2020-08-05 13:14:15", person.Birthday13.String())
	assert.Equal(t, "2020-08-05 13:14:15.999", person.Birthday14.String())
	assert.Equal(t, "2020-08-05 13:14:15.999999", person.Birthday15.String())
	assert.Equal(t, "2020-08-05 13:14:15.999999999", person.Birthday16.String())
	assert.Equal(t, "20200805131415", person.Birthday17.String())
	assert.Equal(t, "20200805131415.999", person.Birthday18.String())
	assert.Equal(t, "20200805131415.999999", person.Birthday19.String())
	assert.Equal(t, "20200805131415.999999999", person.Birthday20.String())
	assert.Equal(t, "Wed, Aug 5, 2020 1:14 PM", person.Birthday21.String())
	assert.Equal(t, "2020-08-05", person.Birthday22.String())
	assert.Equal(t, "2020-08-05.999", person.Birthday23.String())
	assert.Equal(t, "2020-08-05.999999", person.Birthday24.String())
	assert.Equal(t, "2020-08-05.999999999", person.Birthday25.String())
	assert.Equal(t, "20200805", person.Birthday26.String())
	assert.Equal(t, "20200805.999", person.Birthday27.String())
	assert.Equal(t, "20200805.999999", person.Birthday28.String())
	assert.Equal(t, "20200805.999999999", person.Birthday29.String())
	assert.Equal(t, "13:14:15", person.Birthday30.String())
	assert.Equal(t, "13:14:15.999", person.Birthday31.String())
	assert.Equal(t, "13:14:15.999999", person.Birthday32.String())
	assert.Equal(t, "13:14:15.999999999", person.Birthday33.String())
	assert.Equal(t, "131415", person.Birthday34.String())
	assert.Equal(t, "131415.999", person.Birthday35.String())
	assert.Equal(t, "131415.999999", person.Birthday36.String())
	assert.Equal(t, "131415.999999999", person.Birthday37.String())
	assert.Equal(t, "2020-08-05T13:14:15+08:00", person.Birthday38.String())
	assert.Equal(t, "Wed Aug  5 13:14:15 2020", person.Birthday39.String())
	assert.Equal(t, "Wednesday, 05-Aug-2020 13:14:15 CST", person.Birthday40.String())
	assert.Equal(t, "1:14PM", person.Birthday41.String())
	assert.Equal(t, "Wed, 05 Aug 2020 13:14:15 +0800", person.Birthday42.String())
	assert.Equal(t, "Wed Aug 05 13:14:15 +0800 2020", person.Birthday43.String())
	assert.Equal(t, "Wed Aug  5 13:14:15 CST 2020", person.Birthday44.String())
	assert.Equal(t, "Wed, 05 Aug 20 13:14:15 +0800", person.Birthday45.String())
	assert.Equal(t, "Wed, 05 Aug 2020 13:14:15 CST", person.Birthday46.String())
	assert.Equal(t, "Wed, 05 Aug 2020 13:14:15 +0800", person.Birthday47.String())
	assert.Equal(t, "Wed, 05 Aug 2020 13:14:15 +0800", person.Birthday48.String())
	assert.Equal(t, "2020-08-05T13:14:15+08:00", person.Birthday49.String())
	assert.Equal(t, "2020-08-05T13:14:15.999+08:00", person.Birthday50.String())
	assert.Equal(t, "2020-08-05T13:14:15.999999+08:00", person.Birthday51.String())
	assert.Equal(t, "2020-08-05T13:14:15.999999999+08:00", person.Birthday52.String())
	assert.Equal(t, "Wed, 05 Aug 2020 13:14:15 CST", person.Birthday53.String())
	assert.Equal(t, "05 Aug 20 13:14 CST", person.Birthday54.String())
	assert.Equal(t, "05 Aug 20 13:14 +0800", person.Birthday55.String())
	assert.Equal(t, "Wednesday, 05-Aug-20 13:14:15 CST", person.Birthday56.String())
	assert.Equal(t, "2020-08-05T13:14:15+08:00", person.Birthday57.String())
	assert.Equal(t, "2020-08-05T13:14:15.999+08:00", person.Birthday58.String())
	assert.Equal(t, "2020-08-05T13:14:15.999999+08:00", person.Birthday59.String())
	assert.Equal(t, "2020-08-05T13:14:15.999999999+08:00", person.Birthday60.String())
	assert.Equal(t, "1596604455", person.Birthday61.String())
	assert.Equal(t, "1596604455999", person.Birthday62.String())
	assert.Equal(t, "1596604455999999", person.Birthday63.String())
	assert.Equal(t, "1596604455999999999", person.Birthday64.String())

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
