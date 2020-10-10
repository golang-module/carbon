package carbon

import (
	"encoding/json"
	"fmt"
	"testing"
)

var user = struct {
	ID          int64            `json:"id"`
	Name        string           `json:"name"`
	Age         int              `json:"age"`
	Birthday    Carbon           `json:"birthday"`
	CreatedAt   ToDateTimeString `json:"created_at"`
	DeletedAt   ToTimestamp      `json:"deleted_at"`
	GraduatedAt ToDateString     `json:"graduated_at"`
	UpdatedAt   ToTimeString     `json:"updated_at"`
}{Name: "勾国印",
	Age:         18,
	Birthday:    Now().SubYears(18),
	CreatedAt:   ToDateTimeString{Now()},
	DeletedAt:   ToTimestamp{Parse("2020-08-05 13:14:15")},
	GraduatedAt: ToDateString{Parse("2012-09-09")},
	UpdatedAt:   ToTimeString{Now()},
}

func TestCarbon_MarshalJSON(t *testing.T) {
	data, _ := json.Marshal(&user)
	fmt.Print("Model output by json:\n", string(data))
}
