# Carbon
Englsih | [Chinese](./README.md)

#### Description
A simple,semantic and developer-friendly golang package for datetime

If you feel good, please give me a star

github:[github.com/golang-module/carbon](https://github.com/golang-module/carbon "github.com/golang-module/carbon")

gitee:[gitee.com/go-package/carbon](https://gitee.com/go-package/carbon "gitee.com/go-package/carbon")

#### Installation
```go
// By github
go get -u github.com/golang-module/carbon

import (
    "github.com/golang-module/carbon"
)

// By gitee
go get -u gitee.com/go-package/carbon

import (
    "gitee.com/go-package/carbon"
)
               
```

#### Usage and example
> The default timezone is Local, assuming the current time is 2020-08-05 13:14:15

##### Set timezone
 ```go
carbon.Timezone(carbon.PRC).Now().ToDateTimeString() // 2020-08-05 13:14:15
carbon.Timezone(carbon.Tokyo).Now().ToDateTimeString() // 2020-08-05 14:14:15
carbon.Timezone(carbon.Tokyo).Timezone(carbon.PRC).Now().ToDateTimeString() // 2020-08-05 13:14:15
 ```

> For more timezone constants, please see the [const.go](./const.go) file

##### Yesterday,today and tomorrow
```go
// Datetime of today
carbon.Now().ToDateTimeString() // 2020-08-05 13:14:15
// Date of today
carbon.Now().ToDateString() // 2020-08-05
// Time of today
carbon.Now().ToTimeString() // 13:14:15
// Timestamp of today
carbon.Now().ToTimestamp() // 1596604455

// Datetime of yesterday
carbon.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
// Date of yesterday
carbon.Yesterday().ToDateString() // 2020-08-04
// Time of yesterday
carbon.Yesterday().ToTimeString() // 13:14:15
// Timestamp of yesterday
carbon.Yesterday().ToTimestamp() // 1596518055

// Datetime of tomorrow
carbon.Tomorrow().ToDateTimeString() // 2020-08-06 13:14:15
// Date of tomorrow
carbon.Tomorrow().ToDateString() // 2020-08-06
// Time of tomorrow
carbon.Tomorrow().ToTimeString() // 13:14:15
// Timestamp of tomorrow
carbon.Tomorrow().ToTimestamp() // 1596690855
```

##### Beginning and end
```go
// Beginning of the year
carbon.Parse("2020-08-05 13:14:15").BeginningOfYear().ToDateTimeString() // 2020-01-01 00:00:00
// End of the year
carbon.Parse("2020-08-05 13:14:15").EndOfYear().ToEndTimeString() // 2020-12-31 23:59:59

// Beginning of the month
carbon.Parse("2020-08-05 13:14:15").BeginningOfMonth().ToStartTimeString() // 2020-08-01 00:00:00
// End of the month
carbon.Parse("2020-08-05 13:14:15").EndOfMonth().ToEndTimeString() // 2020-08-31 23:59:59

// Beginning of the week
carbon.Parse("2020-08-05 13:14:15").FirstOfWeek().ToStartTimeString() // 2020-08-03 00:00:00
// End of the week
carbon.Parse("2020-08-05 13:14:15").LastOfWeek().ToEndTimeString() // 2020-08-09 23:59:59

// Beginning of the day
carbon.Parse("2020-08-05 13:14:15").BeginningOfDay().ToDateTimeString() // 2020-08-05 00:00:00
// End of the day
carbon.Parse("2020-08-05 13:14:15").EndOfDay().ToDateTimeString() // 2020-08-05 23:59:59

// Beginning of the hour
carbon.Parse("2020-08-05 13:14:15").BeginningOfHour().ToDateTimeString() // 2020-08-05 13:00:00
// End of the hour
carbon.Parse("2020-08-05 13:14:15").EndOfHour().ToDateTimeString() // 2020-08-05 13:59:59

// Beginning of the minute
carbon.Parse("2020-08-05 13:14:15").BeginningOfMinute().ToDateTimeString() // 2020-08-05 13:14:00
// End of the minute
carbon.Parse("2020-08-05 13:14:15").EndOfMinute().ToDateTimeString() // 2020-08-05 13:14:59
```

##### Create carbon instance
```go
// Create Carbon instance from timestamp
carbon.CreateFromTimestamp(1596604455).ToDateTimeString() // 2020-08-05 13:14:15
// Create Carbon instance from year,month,day,hour,minute and second
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
// Create Carbon instance from year,month and day
carbon.CreateFromDate(2020, 8, 5).ToDateTimeString() // 2020-08-05 13:14:15
// Create Carbon instance from hour,minute and second
carbon.CreateFromTime(13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
// Create Carbon instance from golang time.Time instance
carbon.CreateFromGoTime(time.Now()).ToTimestamp() // 1596604455
```

##### Parse standard time format string
```go
carbon.Parse("").ToDateTimeString() // empty string
carbon.Parse("0000-00-00 00:00:00").ToDateTimeString() // empty string
carbon.Parse("0000-00-00").ToDateTimeString() // empty string
carbon.Parse("2020-08-05 13:14:15").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("2020-08-05").ToDateTimeString() // 2020-08-05 00:00:00
carbon.Parse("20200805131415").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("20200805").ToDateTimeString() // 2020-08-05 00:00:00
carbon.Parse("2020-08-05T13:14:15+08:00").ToDateTimeString() // 2020-08-05 00:00:00
```

##### Parse custom time format string
```go
carbon.ParseByFormat("2020|08|05 13|14|15", "Y|m|d H|i|s").ToDateTimeString // 2020-08-05 13:14:15
carbon.ParseByFormat("2020%08%05% 13%14%15", "Y%m%d% h%i%s").ToDateTimeString // 2020-08-05 13:14:15
carbon.ParseByFormat("2020年08月05日 13时14分15秒", "Y年m月d日 H时i分s秒").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("2020年08月05日", "Y年m月d日").ToDateTimeString() // 2020-08-05 00:00:00
carbon.ParseByFormat("13时14分15秒", "H时i分s秒").ToDateTimeString() // 2020-08-05 13:14:15
```

##### Parse duration time string (base on now)
```go
// Ten hours later
carbon.ParseByDuration("10h").ToDateTimeString() // 2020-08-06 23:14:15
// Ten and a half hours ago
carbon.ParseByDuration("-10.5h").ToDateTimeString // 2020-08-05 02:44:15
// Ten minutes later
carbon.ParseByDuration("10m").ToDateTimeString // 2020-08-05 13:24:15
// Ten and a half minutes ago
carbon.ParseByDuration("-10.5m").ToDateTimeString // 2020-08-05 13:03:45
// Ten seconds later
carbon.ParseByDuration("10s").ToDateTimeString // 2020-08-05 13:14:25
// Ten seconds ago
carbon.ParseByDuration("-10.5s").ToDateTimeString // 2020-08-05 13:14:04

```

##### Time travel
```go
// After Three years
carbon.Parse("2020-02-29 13:14:15").AddYears(3).ToDateTimeString() // 2023-03-01 13:14:15
// Next three years
carbon.Parse("2020-02-29 13:14:15").NextYears(3).ToDateTimeString() // 2023-02-28 13:14:15
// After one year
carbon.Parse("2020-02-29 13:14:15").AddYear().ToDateTimeString() // 2021-03-01 13:14:15
// Next one year
carbon.Parse("2020-02-29 13:14:15").NextYear().ToDateTimeString() // 2021-02-28 13:14:15
// Before three years
carbon.Parse("2020-02-29 13:14:15").SubYears(3).ToDateTimeString() // 2017-03-01 13:14:15
// Previous three years
carbon.Parse("2020-02-29 13:14:15").PreYears(3).ToDateTimeString() // 2017-02-28 13:14:15
// After one year
carbon.Parse("2020-02-29 13:14:15").SubYear().ToDateTimeString() // 2019-03-01 13:14:15
// Previous one year
carbon.Parse("2020-02-29 13:14:15").PreYear().ToDateTimeString() // 2019-02-28 13:14:15

// After three months
carbon.Parse("2020-02-29 13:14:15").AddMonths(3).ToDateTimeString() // 2020-05-29 13:14:15
// Next three months
carbon.Parse("2020-02-29 13:14:15").NextMonths(3).ToDateTimeString() // 2020-05-29 13:14:15
// After one month
carbon.Parse("2020-01-31 13:14:15").AddMonth().ToDateTimeString() // 2020-03-02 13:14:15
// Next one month
carbon.Parse("2020-01-31 13:14:15").NextMonth().ToDateTimeString() // 2020-02-29 13:14:15
// After three months
carbon.Parse("2020-02-29 13:14:15").SubMonths(3).ToDateTimeString() // 2019-11-29 13:14:15
// Previous three months
carbon.Parse("2020-02-29 13:14:15").PreMonths(3).ToDateTimeString() // 2019-11-29 13:14:15
// After one month
carbon.Parse("2020-03-31 13:14:15").SubMonth().ToDateTimeString() // 2020-03-02 13:14:15
// Previous one month
carbon.Parse("2020-03-31 13:14:15").PreMonth().ToDateTimeString() // 2020-02-29 13:14:15

// After Three days
carbon.Parse("2020-08-05 13:14:15").AddDays(3).ToDateTimeString() // 2020-08-08 13:14:15
// After One day
carbon.Parse("2020-08-05 13:14:15").AddDay().ToDateTimeString() // 2020-08-05 13:14:15
// After three days
carbon.Parse("2020-08-05 13:14:15").SubDays(3).ToDateTimeString() // 2020-08-02 13:14:15
// After one day
carbon.Parse("2020-08-05 13:14:15").SubDay().ToDateTimeString() // 2020-08-04 13:14:15

// After three hours
carbon.Parse("2020-08-05 13:14:15").AddHours(3).ToDateTimeString() // 2020-08-05 16:14:15
// After two and a half hours
carbon.Parse("2020-08-05 13:14:15").Duration("2.5h").ToDateTimeString() // 2020-08-05 15:44:15
// After one hour
carbon.Parse("2020-08-05 13:14:15").AddHour().ToDateTimeString() // 2020-08-05 14:14:15
// Before three hours
carbon.Parse("2020-08-05 13:14:15").SubHours(3).ToDateTimeString() // 2020-08-05 10:14:15
// Before two and a half hours
carbon.Parse("2020-08-05 13:14:15").Duration("-2.5h").ToDateTimeString() // 2020-08-05 10:44:15
// Before one hour
carbon.Parse("2020-08-05 13:14:15").SubHour().ToDateTimeString() // 2020-08-05 12:14:15

// After three minutes
carbon.Parse("2020-08-05 13:14:15").AddMinutes(3).ToDateTimeString() // 2020-08-05 13:17:15
// After two and a half minutes
carbon.Parse("2020-08-05 13:14:15").Duration("2.5m").ToDateTimeString() // 2020-08-05 13:16:45
// After one minute
carbon.Parse("2020-08-05 13:14:15").AddMinute().ToDateTimeString() // 2020-08-05 13:15:15
// Before three minutes
carbon.Parse("2020-08-05 13:14:15").SubMinutes(3).ToDateTimeString() // 2020-08-05 13:11:15
// Before two and a half minutes
carbon.Parse("2020-08-05 13:14:15").Duration("-2.5m").ToDateTimeString() // 2020-08-05 13:11:45
// Before one minute
carbon.Parse("2020-08-05 13:14:15").SubMinute().ToDateTimeString() // 2020-08-05 13:13:15

// After three seconds
carbon.Parse("2020-08-05 13:14:15").AddSeconds(3).ToDateTimeString() // 2020-08-05 13:14:18
// After two and a half seconds
carbon.Parse("2020-08-05 13:14:15").Duration("2.5s").ToDateTimeString() // 2020-08-05 13:14:17
// After one second
carbon.Parse("2020-08-05 13:14:15").AddSecond().ToDateTimeString() // 2020-08-05 13:14:16
// Before three seconds
carbon.Parse("2020-08-05 13:14:15").SubSeconds(3).ToDateTimeString() // 2020-08-05 13:14:12
// Before two and a half seconds
carbon.Parse("2020-08-05 13:14:15").Duration("-2.5s").ToDateTimeString() // 2020-08-05 13:14:12
// Before one second
carbon.Parse("2020-08-05 13:14:15").SubSecond().ToDateTimeString() // 2020-08-05 13:14:14

```

##### Time output
      
```go
// To timestamp
carbon.Parse("2020-08-05 13:14:15").ToTimestamp() // 1596604455

// To string
carbon.Parse("2020-08-05 13:14:15").Time.String() // 2020-08-05 13:14:15 +0800 CST
// To string of layout format
carbon.Parse("2020-08-05 13:14:15").ToFormatString("YmdHis") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").ToFormatString("Y年m月d H时i分s秒") // 2020年08月05日 13时14分15秒
// To string of datetime format
carbon.Parse("2020-08-05 13:14:15").ToDateTimeString() // 2020-08-05 13:14:15
// To string of date format
carbon.Parse("2020-08-05 13:14:15").ToDateString() // 2020-08-05
// To string of time format
carbon.Parse("2020-08-05 13:14:15").ToTimeString() // 13:14:15

// To string of Ansic format
carbon.Parse("2020-08-05 13:14:15").ToAnsicString() // Wed Aug  5 13:14:15 2020
// To string of Atom format
carbon.Parse("2020-08-05 13:14:15").ToAtomString() // Wed Aug  5 13:14:15 2020
// To string of UnixDate format  
carbon.Parse("2020-08-05 13:14:15").ToUnixDateString() // Wed Aug  5 13:14:15 CST 2020
// To string of RubyDate format 
carbon.Parse("2020-08-05 13:14:15").ToRubyDateString() // Wed Aug 05 13:14:15 +0800 2020
// To string of Kitchen format
carbon.Parse("2020-08-05 13:14:15").ToKitchenString() // 1:14PM
// To string of Cookie format
carbon.Parse("2020-08-05 13:14:15").ToCookieString() // Wednesday, 05-Aug-2020 13:14:15 CST
// To string of DayDateTime format
carbon.Parse("2020-08-05 13:14:15").ToDayDateTimeString() // Wed, Aug 5, 2020 1:14 PM
// To string of RSS format
carbon.Parse("2020-08-05 13:14:15").ToRssString() // Wed, 05 Aug 2020 13:14:15 +0800
// To string of W3C format
carbon.Parse("2020-08-05 13:14:15").ToW3cString() // 2020-08-05T13:14:15+08:00

// To string of RFC822 format
carbon.Parse("2020-08-05 13:14:15").ToRFC822String() // 05 Aug 20 13:14 CST
// To string of RFC822Z format
carbon.Parse("2020-08-05 13:14:15").ToRFC822zString() // 05 Aug 20 13:14 +0800
// To string of RFC850 format
carbon.Parse("2020-08-05 13:14:15").ToRFC850String() // Wednesday, 05-Aug-20 13:14:15 CST
// To string of RFC1036 format
carbon.Parse("2020-08-05 13:14:15").ToRFC1036String() // Wed, 05 Aug 20 13:14:15 +0800
// To string of RFC1123 format
carbon.Parse("2020-08-05 13:14:15").ToRFC1123String() // Wed, 05 Aug 2020 13:14:15 CST
// To string of RFC2822 format
carbon.Parse("2020-08-05 13:14:15").ToRFC2822String() // Wed, 05 Aug 2020 13:14:15 +0800
// To string of RFC3339 format 
carbon.Parse("2020-08-05 13:14:15").ToRFC3339String() // 2020-08-05T13:14:15+08:00
// To string of RFC7231 format
carbon.Parse("2020-08-05 13:14:15").ToRFC7231String() // Wed, 05 Aug 2020 05:14:15 GMT

```
> For more format signs, please see the <a href="#format-sign-table">Format sign table</a>

##### Statistics
```go
// Total days of the year
carbon.Parse("2019-08-05 13:14:15").DaysInYear() // 365
carbon.Parse("2020-08-05 13:14:15").DaysInYear() // 366
// Total days of the month
carbon.Parse("2020-02-01 13:14:15").DaysInMonth() // 29
carbon.Parse("2020-04-01 13:14:15").DaysInMonth() // 30
carbon.Parse("2020-08-01 13:14:15").DaysInMonth() // 31

// current age
carbon.Parse("1990-01-01 13:14:15").Age() // 30
carbon.Parse("1990-12-31 13:14:15").Age() // 29

// Current year
carbon.Parse("2020-08-05 13:14:15").Year() // 2020
// Current month
carbon.Parse("2020-08-05 13:14:15").Month() // 8
// Current day
carbon.Parse("2020-08-05 13:14:15").Day() // 5
// Current hour
carbon.Parse("2020-08-05 13:14:15").Hour() // 13
// Current minute
carbon.Parse("2020-08-05 13:14:15").Minute() // 14
// Current second
carbon.Parse("2020-08-05 13:14:15").Second() // 15
```

##### Week and day
```go
// Day of the year
carbon.Parse("2020-08-05 13:14:15").DayOfYear() // 218
// Week of the year
carbon.Parse("2020-08-05 13:14:15").WeekOfYear() // 32
// Day of the month
carbon.Parse("2020-08-05 13:14:15").DayOfMonth() // 5
// Week of the month
carbon.Parse("2020-08-05 13:14:15").WeekOfMonth() // 1
// Day of the week
carbon.Parse("2020-08-05 13:14:15").DayOfWeek() // 3
```

##### Time judgment
```go
// Is zero time
carbon.Parse("").IsZero() // true
carbon.Parse("0").IsZero() // true
carbon.Parse("0000-00-00 00:00:00").IsZero() // true
carbon.Parse("0000-00-00").IsZero() // true
carbon.Parse("00:00:00").IsZero() // true
carbon.Parse("2020-08-05 00:00:00").IsZero() // false
carbon.Parse("2020-08-05").IsZero() // false

// Is now time
carbon.Parse(carbon.Now().ToDateTimeString()).IsNow() // true
// Is future time
carbon.Parse("2020-08-06 13:14:15").IsFuture() // true
// Is pass time
carbon.Parse("2020-08-04 13:14:15").IsPast() // true
// Is leap year
carbon.Parse("2020-08-05 13:14:15").IsLeapYear() // true

// Is January 
carbon.Parse("2020-08-05 13:14:15").IsJanuary() // false
// Is February
carbon.Parse("2020-08-05 13:14:15").IsFebruary() // false
// Is March
carbon.Parse("2020-08-05 13:14:15").IsMarch() // false
// Is April
carbon.Parse("2020-08-05 13:14:15").IsApril()  // false
// Is May
carbon.Parse("2020-08-05 13:14:15").IsMay() // false
// Is June
carbon.Parse("2020-08-05 13:14:15").IsJune() // false
// Is July
carbon.Parse("2020-08-05 13:14:15").IsJuly() // false
// Is August
carbon.Parse("2020-08-05 13:14:15").IsAugust() // false
// Is September
carbon.Parse("2020-08-05 13:14:15").IsSeptember() // true
// Is October
carbon.Parse("2020-08-05 13:14:15").IsOctober() // false
// Is November
carbon.Parse("2020-08-05 13:14:15").IsNovember() // false
// Is December
carbon.Parse("2020-08-05 13:14:15").IsDecember() // false

// Is Monday
carbon.Parse("2020-08-05 13:14:15").IsMonday() // false
// Is Tuesday
carbon.Parse("2020-08-05 13:14:15").IsTuesday() // true
// Is Wednesday
carbon.Parse("2020-08-05 13:14:15").IsWednesday() // false
// Is Thursday
carbon.Parse("2020-08-05 13:14:15").IsThursday()  // false
// Is Friday
carbon.Parse("2020-08-05 13:14:15").IsFriday() // false
// Is Saturday
carbon.Parse("2020-08-05 13:14:15").IsSaturday() // false
// Is Sunday
carbon.Parse("2020-08-05 13:14:15").IsSunday() // false
// Is Weekday
carbon.Parse("2020-08-05 13:14:15").IsWeekday() // false
// Is Weekend
carbon.Parse("2020-08-05 13:14:15").IsWeekend() // true

// Is Yesterday
carbon.Parse("2020-08-04 13:14:15").IsYesterday() // true
carbon.Parse("2020-08-04 00:00:00").IsYesterday() // true
carbon.Parse("2020-08-04").IsYesterday() // true
// Is Today
carbon.Parse("2020-08-05 13:14:15").IsToday() // true
carbon.Parse("2020-08-05 00:00:00").IsToday() // true
carbon.Parse("2020-08-05").IsToday() // true
// Is Tomorrow
carbon.Parse("2020-08-06 13:14:15").IsTomorrow() // true
carbon.Parse("2020-08-06 00:00:00").IsTomorrow() // true
carbon.Parse("2020-08-06").IsTomorrow() // true
```

##### Calendar
```go
// To year of the animal
carbon.Parse("2020-08-05 13:14:15").ToAnimalYear() // 鼠
// To lunar year
carbon.Parse("2020-08-05 13:14:15").ToLunarYear() // 庚子

// Is year of the rat
carbon.Parse("2020-08-05 13:14:15").IsYearOfRat() // true
// Is year of the ox
carbon.Parse("2020-08-05 13:14:15").IsYearOfOx() // false
// Is year of the tiger
carbon.Parse("2020-08-05 13:14:15").IsYearOfTiger() // false
// Is year of the rabbit
carbon.Parse("2020-08-05 13:14:15").IsYearOfRabbit() // false
// Is year of the dragon
carbon.Parse("2020-08-05 13:14:15").IsYearOfDragon() // false
// Is year of the snake
carbon.Parse("2020-08-05 13:14:15").IsYearOfSnake() // false
// Is year of the horse
carbon.Parse("2020-08-05 13:14:15").IsYearOfHorse() // false
// Is year of the goat
carbon.Parse("2020-08-05 13:14:15").IsYearOfGoat() // false
// Is year of the monkey
carbon.Parse("2020-08-05 13:14:15").IsYearOfMonkey() // false
// Is year of the rooster
carbon.Parse("2020-08-05 13:14:15").IsYearOfRooster() // false
// Is year of the dog
carbon.Parse("2020-08-05 13:14:15").IsYearOfDog() // false
// Is year of the dig
carbon.Parse("2020-08-05 13:14:15").IsYearOfPig() // false
```

##### Database
Assuming the database table is users, its fields have id(int), name(varchar), age(int), graduated_at(date), birthday(date), created_at(datetime), updated_at(datetime), deleted_at(datetime)

###### Define model
```go
type UserModel struct {
    ID  int64  `json:"id"`
    Name string `json:"name"`
    Age int `json:"age"`
    Birthday carbon.Carbon `json:"birthday"`
    CreatedAt carbon.ToDateTimeString `json:"created_at"`
    DeletedAt carbon.ToTimestamp `json:"deleted_at"`
    GraduatedAt carbon.ToDateString `json:"graduated_at"`
    UpdatedAt carbon.ToTimeString `json:"updated_at"`
}
```

###### Instantiate model
```go
user := UserModel {
    Name: "勾国印",
    Age: 18,
    Birthday: carbon.Now().SubYears(18),
    CreatedAt: carbon.ToDateTimeString{carbon.Now()},
    DeletedAt: carbon.ToTimestamp{carbon.Parse("2020-08-05 13:14:15")},
    GraduatedAt: carbon.ToDateString{carbon.Parse("2012-09-09")},
    UpdatedAt: carbon.ToTimeString{carbon.Now()},
}
```

###### Output fields
```go
user.ID // 18
user.Name // 勾国印
user.Birthday.ToDateString() // 2012-08-05
user.CreatedAt.ToTimestamp() // 2012-08-05 13:14:15
user.DeletedAt.ToDateTimeString() // 1596604455
user.GraduatedAt.AddDay().ToDateString() // 2012-09-10
user.UpdatedAt.ToDateString() // 2012-08-05
```

###### Output model by json
```go
data, _ := json.Marshal(&user)
fmt.Print(string(data))
// Output
{
    "id": 42,
    "name": "勾国印",
    "age": 18,
    "birthday": "2012-08-05 00:00:00",
    "created_at": "2020-08-05 13:14:15",
    "deleted_at": 1596604455
    "graduated_at": "2012-09-09",
    "updated_at": "13:14:15",
}
```

###### Output custom format
```go
// Define format
type ToRssString struct {
    carbon.Carbon
}

// Define model
type UserModel struct {
    Birthday ToRssString `json:"birthday"`
}

// Instantiate model
user := UserModel {
    Birthday: Birthday: ToRssString{carbon.Now()},
}

// Overload MarshalJSON method
func (c ToRssString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, c.ToRssString())), nil
}

// Output model by json.Marshal(&user)
{
    "birthday": "Wed, 05 Aug 2020 13:14:15 +0800",
}
```

#### Appendix
##### <a id="format-sign-table">Format sign table</a>

| sign | desc | type | length | range | example |
| ------------ | ------------ | ------------ | ------------ | ------------ | ------------ |
| Y | year | number | 4 | 0000-9999 | 2020 |
| y | year | number | 2 | 00-99 | 20 |
| M | month | letter | 3 | Jan-Dec | Aug |
| m | month | number | 2 | 01-12 | 08 |
| F | month | letter | - | January-December | August |
| n | month | number | 1/2 | 1-12 | 8 |
| l | weekday | letter | - | Monday-Sunday | Wednesday |
| D | weekday | letter | 3 | Mon-Sun | Wed |
| d | day | number | 2 | 01-31 | 05 |
| j | day | number | 1/2 |1-31 | 5 |
| H | hour | number | 2 | 00-23 | 15 |
| h | hour | number | 2 | 00-11 | 03 |
| i | minute | number | 2 | 01-59 | 14 |
| s | second | number | 2 | 01-59 | 15 |
| P | Ante Meridiem/Post Meridiem | letter | 2 | AM/PM | PM |
| p | ante meridiem/post meridiem | letter | 2 | am/pm | pm |

#### Reference
* [briannesbitt/carbon](https://github.com/briannesbitt/Carbon)
* [uniplaces/carbon](https://github.com/uniplaces/carbon)
* [jinzhu/now](https://github.com/jinzhu/now/)
* [araddon/dateparse](https://github.com/araddon/dateparse)