package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewCarbon(t *testing.T) {
	loc, _ := time.LoadLocation(PRC)

	t1, _ := time.Parse(DateTimeLayout, "2020-08-05 13:14:15")
	t2, _ := time.ParseInLocation(DateTimeLayout, "2020-08-05 13:14:15", loc)

	c1 := NewCarbon(t1)
	assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", c1.ToString())
	assert.Equal(t, t1.String(), c1.ToString())

	c2 := NewCarbon(t2)
	assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", c2.ToString())
	assert.Equal(t, t2.String(), c2.ToString())
}

func TestNow(t *testing.T) {
	assert.Equal(t, time.Now().Format(DateLayout), Now().Layout(DateLayout, Local))
	assert.Equal(t, time.Now().In(time.UTC).Format(DateLayout), Now(UTC).Layout(DateLayout))
}

func TestTomorrow(t *testing.T) {
	assert.Equal(t, time.Now().Add(time.Hour*24).Format(DateLayout), Tomorrow().Layout(DateLayout, Local))
	assert.Equal(t, time.Now().Add(time.Hour*24).In(time.UTC).Format(DateLayout), Tomorrow(UTC).Layout(DateLayout))
}

func TestYesterday(t *testing.T) {
	assert.Equal(t, time.Now().Add(time.Hour*-24).Format(DateLayout), Yesterday().Layout(DateLayout, Local))
	assert.Equal(t, time.Now().Add(time.Hour*-24).In(time.UTC).Format(DateLayout), Yesterday(UTC).Layout(DateLayout))
}

func TestParse(t *testing.T) {
	assert.Equal(t, Now().Timestamp(), Parse("now").Timestamp())
	assert.Equal(t, Yesterday().Timestamp(), Parse("yesterday").Timestamp())
	assert.Equal(t, Tomorrow().Timestamp(), Parse("tomorrow").Timestamp())
	assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", Parse("2020-8-5").ToString())
	assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", Parse("2020-8-05").ToString())
	assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", Parse("2020-08-05").ToString())
	assert.Equal(t, "2020-08-05 01:02:03 +0000 UTC", Parse("2020-8-5 1:2:3").ToString())
	assert.Equal(t, "2020-08-05 01:02:03 +0000 UTC", Parse("2020-08-05 1:2:03").ToString())
	assert.Equal(t, "2020-08-05 01:02:03 +0000 UTC", Parse("2020-08-05 1:02:03").ToString())
	assert.Equal(t, "2020-08-05 01:02:03 +0000 UTC", Parse("2020-08-05 01:02:03").ToString())
	assert.Equal(t, Parse("2022-03-08T03:01:14-07:00").ToString(), Parse("2022-03-08T10:01:14Z").ToString())

	assert.Equal(t, Now().Timestamp(), Parse("now", PRC).Timestamp())
	assert.Equal(t, Yesterday().Timestamp(), Parse("yesterday", PRC).Timestamp())
	assert.Equal(t, Tomorrow().Timestamp(), Parse("tomorrow", PRC).Timestamp())
	assert.Equal(t, "2020-08-05 00:00:00 +0800 CST", Parse("2020-8-5", PRC).ToString())
	assert.Equal(t, "2020-08-05 00:00:00 +0800 CST", Parse("2020-8-05", PRC).ToString())
	assert.Equal(t, "2020-08-05 00:00:00 +0800 CST", Parse("2020-08-05", PRC).ToString())
	assert.Equal(t, "2020-08-05 01:02:03 +0800 CST", Parse("2020-8-5 1:2:3", PRC).ToString())
	assert.Equal(t, "2020-08-05 01:02:03 +0800 CST", Parse("2020-08-05 1:2:03", PRC).ToString())
	assert.Equal(t, "2020-08-05 01:02:03 +0800 CST", Parse("2020-08-05 1:02:03", PRC).ToString())
	assert.Equal(t, "2020-08-05 01:02:03 +0800 CST", Parse("2020-08-05 01:02:03", PRC).ToString())
	assert.Equal(t, Parse("2022-03-08T03:01:14-07:00", PRC).ToString(), Parse("2022-03-08T10:01:14Z", PRC).ToString())
}

func TestParseByFormat(t *testing.T) {
	assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", ParseByFormat("2020-08-05", DateFormat).ToString())
	assert.Equal(t, "0000-01-01 13:14:15 +0000 UTC", ParseByFormat("13:14:15", TimeFormat).ToString())
	assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByFormat("2020-08-05 13:14:15", DateTimeFormat).ToString())
	assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByFormat("2020|08|05 13:14:15", "Y|m|d H:i:s").ToString())
	assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToString())
	assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒").ToString())

	assert.Equal(t, "2020-08-05 00:00:00 +0800 CST", ParseByFormat("2020-08-05", DateFormat, PRC).ToString())
	assert.Equal(t, "0000-01-01 13:14:15 +0805 LMT", ParseByFormat("13:14:15", TimeFormat, PRC).ToString())
	assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByFormat("2020-08-05 13:14:15", DateTimeFormat, PRC).ToString())
	assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByFormat("2020|08|05 13:14:15", "Y|m|d H:i:s", PRC).ToString())
	assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s", PRC).ToString())
	assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒", PRC).ToString())
}

func TestParseByLayout(t *testing.T) {
	assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", ParseByLayout("2020-08-05", DateLayout).ToString())
	assert.Equal(t, "0000-01-01 13:14:15 +0000 UTC", ParseByLayout("13:14:15", TimeLayout).ToString())
	assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByLayout("2020-08-05 13:14:15", DateTimeLayout).ToString())
	assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByLayout("2020|08|05 13:14:15", "2006|01|02 15:04:05").ToString())
	assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToString())
	assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒").ToString())

	assert.Equal(t, "2020-08-05 00:00:00 +0800 CST", ParseByLayout("2020-08-05", DateLayout, PRC).ToString())
	assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByLayout("2020-08-05 13:14:15", DateTimeLayout, PRC).ToString())
	assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByLayout("2020-08-05 13:14:15", DateTimeLayout, PRC).ToString())
	assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s 2006-01-02 15:04:05", PRC).ToString())
	assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByLayout("2020|08|05 13:14:15", "2006|01|02 15:04:05", PRC).ToString())
	assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒", PRC).ToString())
}

func TestParseWithLayouts(t *testing.T) {
	c1 := ParseWithLayouts("2020|08|05 13|14|15", []string{"2006|01|02 15|04|05", "2006|1|2 3|4|5"})
	assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", c1.ToString())
	assert.Equal(t, "2006|01|02 15|04|05", c1.CurrentLayout())

	c2 := ParseWithLayouts("2020|08|05 13|14|15", []string{"2006|01|02 15|04|05", "2006|1|2 3|4|5"}, PRC)
	assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", c2.ToString())
	assert.Equal(t, "2006|01|02 15|04|05", c2.CurrentLayout())
}

func TestParseWithFormats(t *testing.T) {
	c1 := ParseWithFormats("2020|08|05 13|14|15", []string{"Y|m|d H|i|s", "y|m|d h|i|s"})
	assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", c1.ToString())
	assert.Equal(t, "2006|01|02 15|04|05", c1.CurrentLayout())

	c2 := ParseWithFormats("2020|08|05 13|14|15", []string{"Y|m|d H|i|s", "y|m|d h|i|s"}, PRC)
	assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", c2.ToString())
	assert.Equal(t, "2006|01|02 15|04|05", c2.CurrentLayout())
}

func TestCreateFromStdTime(t *testing.T) {
	now1 := time.Now()
	assert.Equal(t, now1.Unix(), CreateFromStdTime(now1).Timestamp())

	now2 := time.Now().In(time.Local)
	assert.Equal(t, now2.Unix(), CreateFromStdTime(now2).Timestamp())
	assert.Equal(t, now2.Unix(), CreateFromStdTime(now2, UTC).Timestamp())
}

func TestCreateFromTimestamp(t *testing.T) {
	assert.Equal(t, "1969-12-31 23:59:59 +0000 UTC", CreateFromTimestamp(-1).ToString())
	assert.Equal(t, "1970-01-01 00:00:00 +0000 UTC", CreateFromTimestamp(0).ToString())
	assert.Equal(t, "1970-01-01 00:00:01 +0000 UTC", CreateFromTimestamp(1).ToString())
	assert.Equal(t, "2022-04-12 03:55:55 +0000 UTC", CreateFromTimestamp(1649735755).ToString())

	assert.Equal(t, "1970-01-01 07:59:59 +0800 CST", CreateFromTimestamp(-1, PRC).ToString())
	assert.Equal(t, "1970-01-01 08:00:00 +0800 CST", CreateFromTimestamp(0, PRC).ToString())
	assert.Equal(t, "1970-01-01 08:00:01 +0800 CST", CreateFromTimestamp(1, PRC).ToString())
	assert.Equal(t, "2022-04-12 11:55:55 +0800 CST", CreateFromTimestamp(1649735755, PRC).ToString())
}

func TestCreateFromTimestampMilli(t *testing.T) {
	assert.Equal(t, "1969-12-31 23:59:59.999 +0000 UTC", CreateFromTimestampMilli(-1).ToString())
	assert.Equal(t, "1970-01-01 00:00:00 +0000 UTC", CreateFromTimestampMilli(0).ToString())
	assert.Equal(t, "1970-01-01 00:00:00.001 +0000 UTC", CreateFromTimestampMilli(1).ToString())
	assert.Equal(t, "2022-04-12 03:55:55.981 +0000 UTC", CreateFromTimestampMilli(1649735755981).ToString())

	assert.Equal(t, "1970-01-01 07:59:59.999 +0800 CST", CreateFromTimestampMilli(-1, PRC).ToString())
	assert.Equal(t, "1970-01-01 08:00:00 +0800 CST", CreateFromTimestampMilli(0, PRC).ToString())
	assert.Equal(t, "1970-01-01 08:00:00.001 +0800 CST", CreateFromTimestampMilli(1, PRC).ToString())
	assert.Equal(t, "2022-04-12 11:55:55.981 +0800 CST", CreateFromTimestampMilli(1649735755981, PRC).ToString())
}

func TestCreateFromTimestampMicro(t *testing.T) {
	assert.Equal(t, "1969-12-31 23:59:59.999999 +0000 UTC", CreateFromTimestampMicro(-1).ToString())
	assert.Equal(t, "1970-01-01 00:00:00 +0000 UTC", CreateFromTimestampMicro(0).ToString())
	assert.Equal(t, "1970-01-01 00:00:00.000001 +0000 UTC", CreateFromTimestampMicro(1).ToString())
	assert.Equal(t, "2022-04-12 03:55:55.981566 +0000 UTC", CreateFromTimestampMicro(1649735755981566).ToString())

	assert.Equal(t, "1970-01-01 07:59:59.999999 +0800 CST", CreateFromTimestampMicro(-1, PRC).ToString())
	assert.Equal(t, "1970-01-01 08:00:00 +0800 CST", CreateFromTimestampMicro(0, PRC).ToString())
	assert.Equal(t, "1970-01-01 08:00:00.000001 +0800 CST", CreateFromTimestampMicro(1, PRC).ToString())
	assert.Equal(t, "2022-04-12 11:55:55.981566 +0800 CST", CreateFromTimestampMicro(1649735755981566, PRC).ToString())
}

func TestCreateFromTimestampNano(t *testing.T) {
	assert.Equal(t, "1969-12-31 23:59:59.999999999 +0000 UTC", CreateFromTimestampNano(-1).ToString())
	assert.Equal(t, "1970-01-01 00:00:00 +0000 UTC", CreateFromTimestampNano(0).ToString())
	assert.Equal(t, "1970-01-01 00:00:00.000000001 +0000 UTC", CreateFromTimestampNano(1).ToString())
	assert.Equal(t, "2022-04-12 03:55:55.981566888 +0000 UTC", CreateFromTimestampNano(1649735755981566888).ToString())

	assert.Equal(t, "1970-01-01 07:59:59.999999999 +0800 CST", CreateFromTimestampNano(-1, PRC).ToString())
	assert.Equal(t, "1970-01-01 08:00:00 +0800 CST", CreateFromTimestampNano(0, PRC).ToString())
	assert.Equal(t, "1970-01-01 08:00:00.000000001 +0800 CST", CreateFromTimestampNano(1, PRC).ToString())
	assert.Equal(t, "2022-04-12 11:55:55.981566888 +0800 CST", CreateFromTimestampNano(1649735755981566888, PRC).ToString())
}

func TestCreateFromDateTime(t *testing.T) {
	assert.Equal(t, "-0001-11-30 00:00:00 +0000 UTC", CreateFromDateTime(0, 0, 0, 0, 0, 0).ToString())
	assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToString())

	assert.Equal(t, "-0001-11-30 00:00:00 +0805 LMT", CreateFromDateTime(0, 0, 0, 0, 0, 0, PRC).ToString())
	assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", CreateFromDateTime(2020, 8, 5, 13, 14, 15, PRC).ToString())
}

func TestCreateFromDateTimeMilli(t *testing.T) {
	assert.Equal(t, "-0001-11-30 00:00:00 +0000 UTC", CreateFromDateTimeMilli(0, 0, 0, 0, 0, 0, 0).ToString())
	assert.Equal(t, "2020-08-05 13:14:15.999 +0000 UTC", CreateFromDateTimeMilli(2020, 8, 5, 13, 14, 15, 999).ToString())

	assert.Equal(t, "-0001-11-30 00:00:00 +0805 LMT", CreateFromDateTimeMilli(0, 0, 0, 0, 0, 0, 0, PRC).ToString())
	assert.Equal(t, "2020-08-05 13:14:15.999 +0800 CST", CreateFromDateTimeMilli(2020, 8, 5, 13, 14, 15, 999, PRC).ToString())
}

func TestCreateFromDateTimeMicro(t *testing.T) {
	assert.Equal(t, "-0001-11-30 00:00:00 +0000 UTC", CreateFromDateTimeMicro(0, 0, 0, 0, 0, 0, 0).ToString())
	assert.Equal(t, "2020-08-05 13:14:15.999999 +0000 UTC", CreateFromDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString())

	assert.Equal(t, "-0001-11-30 00:00:00 +0805 LMT", CreateFromDateTimeMicro(0, 0, 0, 0, 0, 0, 0, PRC).ToString())
	assert.Equal(t, "2020-08-05 13:14:15.999999 +0800 CST", CreateFromDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999, PRC).ToString())
}

func TestCreateFromDateTimeNano(t *testing.T) {
	assert.Equal(t, "-0001-11-30 00:00:00 +0000 UTC", CreateFromDateTimeNano(0, 0, 0, 0, 0, 0, 0).ToString())
	assert.Equal(t, "2020-08-05 13:14:15.999999999 +0000 UTC", CreateFromDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999).ToString())

	assert.Equal(t, "-0001-11-30 00:00:00 +0805 LMT", CreateFromDateTimeNano(0, 0, 0, 0, 0, 0, 0, PRC).ToString())
	assert.Equal(t, "2020-08-05 13:14:15.999999999 +0800 CST", CreateFromDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999, PRC).ToString())
}

func TestCreateFromDate(t *testing.T) {
	assert.Equal(t, "-0001-11-30 00:00:00 +0000 UTC", CreateFromDate(0, 0, 0).ToString())
	assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", CreateFromDate(2020, 8, 5).ToString())

	assert.Equal(t, "-0001-11-30 00:00:00 +0805 LMT", CreateFromDate(0, 0, 0, PRC).ToString())
	assert.Equal(t, "2020-08-05 00:00:00 +0800 CST", CreateFromDate(2020, 8, 5, PRC).ToString())
}

func TestCreateFromDateMilli(t *testing.T) {
	assert.Equal(t, "-0001-11-30 00:00:00 +0000 UTC", CreateFromDateMilli(0, 0, 0, 0).ToString())
	assert.Equal(t, "2020-08-05 00:00:00.999 +0000 UTC", CreateFromDateMilli(2020, 8, 5, 999).ToString())

	assert.Equal(t, "-0001-11-30 00:00:00 +0805 LMT", CreateFromDateMilli(0, 0, 0, 0, PRC).ToString())
	assert.Equal(t, "2020-08-05 00:00:00.999 +0800 CST", CreateFromDateMilli(2020, 8, 5, 999, PRC).ToString())
}

func TestCreateFromDateMicro(t *testing.T) {
	assert.Equal(t, "-0001-11-30 00:00:00 +0000 UTC", CreateFromDateMicro(0, 0, 0, 0).ToString())
	assert.Equal(t, "2020-08-05 00:00:00.999999 +0000 UTC", CreateFromDateMicro(2020, 8, 5, 999999).ToString())

	assert.Equal(t, "-0001-11-30 00:00:00 +0805 LMT", CreateFromDateMicro(0, 0, 0, 0, PRC).ToString())
	assert.Equal(t, "2020-08-05 00:00:00.999999 +0800 CST", CreateFromDateMicro(2020, 8, 5, 999999, PRC).ToString())
}

func TestCreateFromDateNano(t *testing.T) {
	assert.Equal(t, "-0001-11-30 00:00:00 +0000 UTC", CreateFromDateNano(0, 0, 0, 0).ToString())
	assert.Equal(t, "2020-08-05 00:00:00.999999999 +0000 UTC", CreateFromDateNano(2020, 8, 5, 999999999).ToString())

	assert.Equal(t, "-0001-11-30 00:00:00 +0805 LMT", CreateFromDateNano(0, 0, 0, 0, PRC).ToString())
	assert.Equal(t, "2020-08-05 00:00:00.999999999 +0800 CST", CreateFromDateNano(2020, 8, 5, 999999999, PRC).ToString())
}

func TestCreateFromTime(t *testing.T) {
	assert.Equal(t, "00:00:00", CreateFromTime(0, 0, 0).ToTimeString())
	assert.Equal(t, "13:14:15", CreateFromTime(13, 14, 15).ToTimeString())

	assert.Equal(t, "00:00:00", CreateFromTime(0, 0, 0, PRC).ToTimeString())
	assert.Equal(t, "13:14:15", CreateFromTime(13, 14, 15, PRC).ToTimeString())
}

func TestCreateFromTimeMilli(t *testing.T) {
	assert.Equal(t, "00:00:00", CreateFromTimeMilli(0, 0, 0, 0).ToTimeMilliString())
	assert.Equal(t, "13:14:15.999", CreateFromTimeMilli(13, 14, 15, 999).ToTimeMilliString())

	assert.Equal(t, "00:00:00", CreateFromTimeMilli(0, 0, 0, 0, PRC).ToTimeMilliString())
	assert.Equal(t, "13:14:15.999", CreateFromTimeMilli(13, 14, 15, 999, PRC).ToTimeMilliString())
}

func TestCreateFromTimeMicro(t *testing.T) {
	assert.Equal(t, "00:00:00", CreateFromTimeMicro(0, 0, 0, 0).ToTimeMicroString())
	assert.Equal(t, "13:14:15.999999", CreateFromTimeMicro(13, 14, 15, 999999).ToTimeMicroString())

	assert.Equal(t, "00:00:00", CreateFromTimeMicro(0, 0, 0, 0, PRC).ToTimeMicroString())
	assert.Equal(t, "13:14:15.999999", CreateFromTimeMicro(13, 14, 15, 999999, PRC).ToTimeMicroString())
}

func TestCreateFromTimeNano(t *testing.T) {
	assert.Equal(t, "00:00:00", CreateFromTimeNano(0, 0, 0, 0).ToTimeNanoString())
	assert.Equal(t, "13:14:15.999999999", CreateFromTimeNano(13, 14, 15, 999999999).ToTimeNanoString())

	assert.Equal(t, "00:00:00", CreateFromTimeNano(0, 0, 0, 0, PRC).ToTimeNanoString())
	assert.Equal(t, "13:14:15.999999999", CreateFromTimeNano(13, 14, 15, 999999999, PRC).ToTimeNanoString())
}

func TestSetLayout(t *testing.T) {
	defer SetLayout(DateTimeLayout)

	SetLayout(DateLayout)
	c1 := Parse("2020-08-05")
	assert.Equal(t, DateLayout, DefaultLayout)
	assert.Equal(t, DateLayout, c1.CurrentLayout())
	assert.Equal(t, "2020-08-05", c1.String())

	SetLayout(DateTimeLayout)
	c2 := Parse("2020-08-05 13:14:15")
	assert.Equal(t, DateTimeLayout, DefaultLayout)
	assert.Equal(t, DateTimeLayout, c2.CurrentLayout())
	assert.Equal(t, "2020-08-05 13:14:15", c2.String())
}

func TestSetFormat(t *testing.T) {
	defer SetFormat(DateTimeFormat)

	SetFormat(DateFormat)
	c1 := Parse("2020-08-05")
	assert.Equal(t, DateLayout, DefaultLayout)
	assert.Equal(t, DateLayout, c1.CurrentLayout())

	SetFormat(DateTimeFormat)
	c2 := Parse("2020-08-05 13:14:15")
	assert.Equal(t, DateTimeLayout, DefaultLayout)
	assert.Equal(t, DateTimeLayout, c2.CurrentLayout())
}

func TestSetTimezone(t *testing.T) {
	defer SetTimezone(UTC)

	SetTimezone(UTC)
	c1 := Parse("2020-08-05 13:14:15")
	assert.Equal(t, UTC, DefaultTimezone)
	assert.Equal(t, UTC, c1.Timezone())
	assert.Equal(t, UTC, c1.ZoneName())
	assert.Equal(t, 0, c1.ZoneOffset())
	assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", c1.ToString())

	SetTimezone(PRC)
	c2 := Parse("2020-08-05 13:14:15")
	assert.Equal(t, PRC, DefaultTimezone)
	assert.Equal(t, PRC, c2.Timezone())
	assert.Equal(t, "CST", c2.ZoneName())
	assert.Equal(t, 28800, c2.ZoneOffset())
	assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", c2.ToString())

	SetTimezone(Japan)
	c3 := Parse("2020-08-05 13:14:15")
	assert.Equal(t, Japan, DefaultTimezone)
	assert.Equal(t, Japan, c3.Timezone())
	assert.Equal(t, "JST", c3.ZoneName())
	assert.Equal(t, 32400, c3.ZoneOffset())
	assert.Equal(t, "2020-08-05 13:14:15 +0900 JST", c3.ToString())
}

func TestSetLocation(t *testing.T) {
	defer SetLocation(time.UTC)

	l1, _ := time.LoadLocation(UTC)
	SetLocation(l1)
	c1 := Parse("2020-08-05 13:14:15")
	assert.Equal(t, UTC, DefaultTimezone)
	assert.Equal(t, UTC, c1.Timezone())
	assert.Equal(t, UTC, c1.ZoneName())
	assert.Equal(t, 0, c1.ZoneOffset())
	assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", c1.ToString())

	l2, _ := time.LoadLocation(PRC)
	SetLocation(l2)
	c2 := Parse("2020-08-05 13:14:15")
	assert.Equal(t, PRC, DefaultTimezone)
	assert.Equal(t, PRC, c2.Timezone())
	assert.Equal(t, "CST", c2.ZoneName())
	assert.Equal(t, 28800, c2.ZoneOffset())
	assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", c2.ToString())

	l3, _ := time.LoadLocation(Japan)
	SetLocation(l3)
	c3 := Parse("2020-08-05 13:14:15")
	assert.Equal(t, Japan, DefaultTimezone)
	assert.Equal(t, Japan, c3.Timezone())
	assert.Equal(t, "JST", c3.ZoneName())
	assert.Equal(t, 32400, c3.ZoneOffset())
	assert.Equal(t, "2020-08-05 13:14:15 +0900 JST", c3.ToString())
}

func TestSetLocale(t *testing.T) {
	defer SetLocale("en")

	SetLocale("zh-CN")
	c1 := Parse("2020-08-05")
	assert.Equal(t, "zh-CN", DefaultLocale)
	assert.Equal(t, "zh-CN", c1.Locale())
	assert.Equal(t, "狮子座", c1.Constellation())
	assert.Equal(t, "夏季", c1.Season())
	assert.Equal(t, "八月", c1.ToMonthString())
	assert.Equal(t, "8月", c1.ToShortMonthString())
	assert.Equal(t, "星期三", c1.ToWeekString())
	assert.Equal(t, "周三", c1.ToShortWeekString())

	SetLocale("en")
	c2 := Parse("2020-08-05")
	assert.Equal(t, "en", DefaultLocale)
	assert.Equal(t, "en", c2.Locale())
	assert.Equal(t, "Leo", c2.Constellation())
	assert.Equal(t, "Summer", c2.Season())
	assert.Equal(t, "August", c2.ToMonthString())
	assert.Equal(t, "Aug", c2.ToShortMonthString())
	assert.Equal(t, "Wednesday", c2.ToWeekString())
	assert.Equal(t, "Wed", c2.ToShortWeekString())
}

func TestSetWeekStartsAt(t *testing.T) {
	defer SetWeekStartsAt(Monday)

	SetWeekStartsAt(Monday)
	c1 := Parse("2020-08-05")
	assert.Equal(t, Monday, DefaultWeekStartsAt)
	assert.Equal(t, Monday, c1.WeekStartsAt())
	assert.Equal(t, "2020-08-03 00:00:00 +0000 UTC", c1.StartOfWeek().ToString())

	SetWeekStartsAt(Sunday)
	c2 := Parse("2020-08-05")
	assert.Equal(t, Sunday, DefaultWeekStartsAt)
	assert.Equal(t, Sunday, c2.WeekStartsAt())
	assert.Equal(t, "2020-08-02 00:00:00 +0000 UTC", c2.StartOfWeek().ToString())
}

func TestSetWeekendDays(t *testing.T) {
	defer SetWeekendDays([]Weekday{
		Saturday, Sunday,
	})

	SetWeekendDays([]Weekday{
		Saturday,
	})
	assert.True(t, Parse("2025-04-12").IsWeekend())
	assert.False(t, Parse("2025-04-13").IsWeekend())

	SetWeekendDays([]Weekday{
		Sunday,
	})
	assert.False(t, Parse("2025-04-12").IsWeekend())
	assert.True(t, Parse("2025-04-13").IsWeekend())
}

func TestSetTestNow(t *testing.T) {
	now := Parse("2020-08-05")

	SetTestNow(now)
	assert.Equal(t, "2020-08-05", Now().ToDateString())
	assert.Equal(t, "2020-08-04", Yesterday().ToDateString())
	assert.Equal(t, "2020-08-06", Tomorrow().ToDateString())
	assert.Equal(t, "just now", Now().DiffForHumans())
	assert.Equal(t, "1 day ago", Yesterday().DiffForHumans())
	assert.Equal(t, "1 day from now", Tomorrow().DiffForHumans())
	assert.Equal(t, "2 months from now", Parse("2020-10-05").DiffForHumans())
	assert.Equal(t, "2 months before", now.DiffForHumans(Parse("2020-10-05")))
	assert.True(t, IsTestNow())

	CleanTestNow()
	assert.Equal(t, time.Now().In(time.UTC).Format(DateLayout), Now().ToDateString())
	assert.Equal(t, time.Now().In(time.UTC).Add(time.Hour*-24).Format(DateLayout), Yesterday().ToDateString())
	assert.Equal(t, time.Now().In(time.UTC).Add(time.Hour*24).Format(DateLayout), Tomorrow().ToDateString())
	assert.False(t, IsTestNow())
}

func TestSetDefault(t *testing.T) {
	defer ResetDefault()

	SetDefault(Default{
		Layout:       DateTimeLayout,
		Timezone:     PRC,
		Locale:       "zh-CN",
		WeekStartsAt: Monday,
		WeekendDays: []Weekday{
			Saturday, Sunday,
		},
	})

	assert.Equal(t, DateTimeLayout, DefaultLayout)
	assert.Equal(t, PRC, DefaultTimezone)
	assert.Equal(t, "zh-CN", DefaultLocale)
	assert.Equal(t, Monday, DefaultWeekStartsAt)
	assert.Equal(t, []Weekday{
		Saturday, Sunday,
	}, DefaultWeekendDays)
}

func TestMaxValue(t *testing.T) {
	assert.Equal(t, "9999-12-31 23:59:59.999999999 +0000 UTC", MaxValue().ToString())
	assert.Equal(t, "December", MaxValue().ToMonthString())
	assert.Equal(t, "十二月", MaxValue().SetLocale("zh-CN").ToMonthString())
}

func TestMinValue(t *testing.T) {
	assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", MinValue().ToString())
	assert.Equal(t, "January", MinValue().ToMonthString())
	assert.Equal(t, "一月", MinValue().SetLocale("zh-CN").ToMonthString())
}

func TestMaxDuration(t *testing.T) {
	assert.Equal(t, 9.223372036854776e+09, MaxDuration().Seconds())
}

func TestMinDuration(t *testing.T) {
	assert.Equal(t, -9.223372036854776e+09, MinDuration().Seconds())
}

func TestMax(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", Max(NewCarbon(), NewCarbon()).ToString())
		assert.Equal(t, "2021-08-05 00:00:00 +0000 UTC", Max(NewCarbon(), Parse("2021-08-05")).ToString())
		assert.Equal(t, "2021-08-05 00:00:00 +0000 UTC", Max(Parse("2021-08-05"), NewCarbon()).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Max(Parse(""), Tomorrow()).ToString())
		assert.Empty(t, Max(Parse("0"), Tomorrow()).ToString())
		assert.Empty(t, Max(Now(), Parse("xxx")).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-06 00:00:00 +0000 UTC", Max(Parse("2020-08-06")).ToString())
		assert.Equal(t, "2021-08-05 00:00:00 +0000 UTC", Max(Parse("2020-08-06"), Parse("2021-08-05")).ToString())
	})
}

func TestMin(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", Min(NewCarbon(), NewCarbon()).ToString())
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", Min(NewCarbon(), Parse("2021-08-05")).ToString())
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", Min(Parse("2021-08-05"), NewCarbon()).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Min(Parse(""), Tomorrow()).ToString())
		assert.Empty(t, Min(Parse("0"), Tomorrow()).ToString())
		assert.Empty(t, Min(Now(), Parse("xxx")).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-06 00:00:00 +0000 UTC", Min(Parse("2020-08-06")).ToString())
		assert.Equal(t, "2020-08-06 00:00:00 +0000 UTC", Min(Parse("2020-08-06"), Parse("2021-08-05")).ToString())
	})
}

func TestLanguage_NewLanguage(t *testing.T) {
	lang := NewLanguage()

	lang.SetLocale("en")
	assert.Equal(t, "Leo", Parse("2020-08-05").SetLanguage(lang).Constellation())
	assert.Equal(t, "Summer", Parse("2020-08-05").SetLanguage(lang).Season())
	assert.Equal(t, "4 years before", Parse("2020-08-05").SetLanguage(lang).DiffForHumans(Parse("2024-08-05")))
	assert.Equal(t, "August", Parse("2020-08-05").SetLanguage(lang).ToMonthString())
	assert.Equal(t, "Aug", Parse("2020-08-05").SetLanguage(lang).ToShortMonthString())
	assert.Equal(t, "Wednesday", Parse("2020-08-05").SetLanguage(lang).ToWeekString())
	assert.Equal(t, "Wed", Parse("2020-08-05").SetLanguage(lang).ToShortWeekString())

	lang.SetLocale("zh-CN")
	assert.Equal(t, "狮子座", Parse("2020-08-05").SetLanguage(lang).Constellation())
	assert.Equal(t, "夏季", Parse("2020-08-05").SetLanguage(lang).Season())
	assert.Equal(t, "4 年前", Parse("2020-08-05").SetLanguage(lang).DiffForHumans(Parse("2024-08-05")))
	assert.Equal(t, "八月", Parse("2020-08-05").SetLanguage(lang).ToMonthString())
	assert.Equal(t, "8月", Parse("2020-08-05").SetLanguage(lang).ToShortMonthString())
	assert.Equal(t, "星期三", Parse("2020-08-05").SetLanguage(lang).ToWeekString())
	assert.Equal(t, "周三", Parse("2020-08-05").SetLanguage(lang).ToShortWeekString())
}
