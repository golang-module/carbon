package carbon

import (
	"encoding/json"
	"fmt"
	"testing"
)

var user = struct {
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
}{
	Name:        "勾国印",
	Age:         18,
	Birthday:    ToDateTimeString{Now().SubYears(18)},
	GraduatedAt: ToDateString{Parse("2012-09-09")},
	CreatedAt:   ToTimeString{Now()},
	UpdatedAt:   ToTimestamp{Now()},
	DateTime1:   ToTimestampWithSecond{Now()},
	DateTime2:   ToTimestampWithMillisecond{Now()},
	DateTime3:   ToTimestampWithMicrosecond{Now()},
	DateTime4:   ToTimestampWithNanosecond{Now()},
}

func TestCarbon_MarshalJSON(*testing.T) {
	data, _ := json.Marshal(&user)
	fmt.Print("Model output by json:\n", string(data)+"\n")
}
