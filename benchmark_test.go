package carbon

import (
	"testing"
	"time"
)

func BenchmarkCarbon_SetLocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, loc := range locationList {
			Now().SetLocation(loc)
		}
	}
}

func BenchmarkCarbon_SetTimezone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, timezone := range timezoneList {
			Now().SetTimezone(timezone)
		}
	}
}

var timezoneList = getTimezoneList()
var locationList = getLocationList()

func getTimezoneList() []string {
	return []string{
		Local, UTC, GMT, CST, EET, WET, CET, EST, MST, Cuba, Egypt, Eire, Greenwich, Iceland,
		Iran, Israel, Jamaica, Japan, Libya, Poland, Portugal, PRC, Singapore, Turkey,
	}
}

func getLocationList() []*time.Location {
	timezoneList := getTimezoneList()
	result := make([]*time.Location, 0, len(timezoneList))
	for _, v := range timezoneList {
		loc, _ := time.LoadLocation(v)
		result = append(result, loc)
	}
	return result
}
