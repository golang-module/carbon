# Persian(Jalaali) Calendar

English | [简体中文](README.cn.md) | [日本語](README.jp.md)

#### Usage and example

##### Convert `Gregorian` calendar to `Persian` calendar

```go
// Get persian year
carbon.Parse("2020-08-05 13:14:15").Persian().Year() // 1399
// Get persian month
carbon.Parse("2020-08-05 13:14:15").Persian().Month() // 5
// Get persian day
carbon.Parse("2020-08-05 13:14:15").Persian().Day() // 15
// Get persian hour
carbon.Parse("2020-08-05 13:14:15").Persian().Hour() // 13
// Get persian minute
carbon.Parse("2020-08-05 13:14:15").Lunar().Minute() // 14
// Get persian second
carbon.Parse("2020-08-05 13:14:15").Lunar().Second() // 15

// Get persian date and time string
carbon.Parse("2020-08-05 13:14:15").Lunar().String() // 1399-05-15 13:14:15
fmt.Printf("%s", carbon.Parse("2020-08-05 13:14:15").Lunar()) // 1399-05-15 13:14:15
// // Get persian month as string
carbon.Parse("2020-08-05 13:14:15").Persian().ToMonthString() // مرداد
// // Get persian week as string
carbon.Parse("2020-08-05 13:14:15").Persian().ToWeekString() // چهارشنبه

```

##### Convert `Persian` calendar to `Gregorian` calendar

```go
carbon.CreateFromPersian(1, 1, 1, 0, 0, 0).ToDateTimeString() // 2016-03-20 00:00:00
carbon.CreateFromPersian(622, 1, 1, 0, 0, 0).ToDateTimeString() // 1243-03-21 00:00:00
carbon.CreateFromPersian(1395, 1, 1, 0, 0, 0).ToDateTimeString() // 2016-03-20 00:00:00
carbon.CreateFromPersian(9377, 1, 1, 0, 0, 0).ToDateTimeString() // 9998-03-19 00:00:00
```

##### Comparison
```go
// Whether is a valid persian date
carbon.CreateFromPersian(1, 1, 1, 0, 0, 0).IsValid() // true
carbon.CreateFromPersian(622, 1, 1, 0, 0, 0).IsValid() // true
carbon.CreateFromPersian(9377, 1, 1, 0, 0, 0).IsValid() // true
carbon.CreateFromPersian(9999, 1, 1, 0, 0, 0).IsValid() // false

// Whether is a persian leap year
carbon.CreateFromPersian(1395, 1, 1, 0, 0, 0).IsLeapYear() // true
carbon.CreateFromPersian(9377, 1, 1, 0, 0, 0).IsLeapYear() // true
carbon.CreateFromPersian(622, 1, 1, 0, 0, 0).IsLeapYear() // false
carbon.CreateFromPersian(9999, 1, 1, 0, 0, 0).IsLeapYear() // false

```