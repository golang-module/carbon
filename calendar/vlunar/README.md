# Carbon

[![Carbon Release](https://img.shields.io/github/release/golang-module/carbon.svg)](https://github.com/golang-module/carbon/releases)
[![Go Test](https://github.com/golang-module/carbon/actions/workflows/test.yml/badge.svg)](https://github.com/golang-module/carbon/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/golang-module/carbon/v2)](https://goreportcard.com/report/github.com/golang-module/carbon/v2)
[![Go Coverage](https://codecov.io/gh/golang-module/carbon/branch/master/graph/badge.svg)](https://codecov.io/gh/golang-module/carbon)
[![Goproxy.cn](https://goproxy.cn/stats/github.com/golang-module/carbon/badges/download-count.svg)](https://goproxy.cn)
[![Carbon Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/golang-module/carbon/v2)
[![License](https://img.shields.io/github/license/golang-module/carbon)](https://github.com/golang-module/carbon/blob/master/LICENSE)

#### Introduction

A simple, semantic and developer-friendly golang package for time, has been included
by [awesome-go](https://github.com/avelino/awesome-go#date-and-time "awesome-go")

[github.com/golang-module/carbon](https://github.com/golang-module/carbon "github.com/golang-module/carbon")

[gitee.com/golang-module/carbon](https://gitee.com/golang-module/carbon "gitee.com/golang-module/carbon")

#### Vietnamese Lunar

##### Convert Solar into VLunar

```go
// Get Lunar year of animal
carbon.Parse("2024-01-25 13:14:15").VLunar().Animal() // Mão

// Get lunar festival
carbon.Parse("2021-02-12 13:14:15").VLunar().Festivals() // []{{Day: 1, Month: 1, Name: "Tết Nguyên Đán"}}

// Get lunar luck hours
carbon.Parse("2020-05-01 13:14:15").VLunar().LuckyHours() // []{{Chi: "Dần", From: 3, To: 5}, {Chi: "Thìn", From: 7, To: 9}, {Chi: "Tỵ", From: 9, To: 11}, {Chi: "Thân", From: 15, To: 17}, {Chi: "Dậu", From: 17, To: 19}, {Chi: "Hợi", From: 21, To: 23}},

// Get lunar solar term
carbon.Parse("2025-03-10 13:14:15").VLunar().SolarTerm() // {Longitude: 345, Name: "Kinh trập"}

// Get lunar year, month, day, hour, minute and second
carbon.Parse("2020-08-05 13:14:15").VLunar().DateTime() // 2020, 6, 16, 13, 14, 15
// Get lunar year, month and day
carbon.Parse("2020-08-05 13:14:15").VLunar().Date() // 2020, 6, 16
// Get lunar hour, minute and second
carbon.Parse("2020-08-05 13:14:15").VLunar().Time() // 13, 14, 15

// Get lunar year
carbon.Parse("2020-08-05 13:14:15").VLunar().Year() // 2020
// Get lunar month
carbon.Parse("2020-08-05 13:14:15").VLunar().Month() // 6
// Get lunar leap month
carbon.Parse("2020-08-05 13:14:15").VLunar().LeapMonth() // 4
// Get lunar day
carbon.Parse("2020-08-05 13:14:15").VLunar().Day() // 16
// Get lunar date as YYYY-MM-DD HH::ii::ss format string
fmt.Printf("%s", carbon.Parse("2020-08-05 13:14:15").VLunar()) // 2020-06-16 13:14:15

// Get lunar year as string
carbon.Parse("2021-05-01 13:14:15").VLunar().ToYearString() // Tân Sửu
// Get lunar month as string
carbon.Parse("2020-12-01 13:14:15").VLunar().ToMonthString() // Đinh Hợi
// Get lunar day as string
carbon.Parse("2020-08-01 13:14:15").VLunar().ToDayString() // Bính Tý
// Get lunar date as string
carbon.Parse("2020-01-01 13:14:15").VLunar().ToDateString() // Ngày 07 tháng 12 năm 2019

// Whether is a lunar leap year
carbon.Parse("2020-08-05 13:14:15").VLunar().IsLeapYear() // true
// Whether is a lunar leap month
carbon.Parse("2020-08-05 13:14:15").VLunar().IsLeapMonth() // false

// Whether is a lunar year of the rat
carbon.Parse("2020-08-05 13:14:15").VLunar().IsRatYear() // true
// Whether is a lunar year of the ox
carbon.Parse("2020-08-05 13:14:15").VLunar().IsOxYear() // false
// Whether is a lunar year of the tiger
carbon.Parse("2020-08-05 13:14:15").VLunar().IsTigerYear() // false
// Whether is a lunar year of the rabbit
carbon.Parse("2020-08-05 13:14:15").VLunar().IsRabbitYear() // false
// Whether is a lunar year of the dragon
carbon.Parse("2020-08-05 13:14:15").VLunar().IsDragonYear() // false
// Whether is a lunar year of the snake
carbon.Parse("2020-08-05 13:14:15").VLunar().IsSnakeYear() // false
// Whether is a lunar year of the horse
carbon.Parse("2020-08-05 13:14:15").VLunar().IsHorseYear() // false
// Whether is a lunar year of the goat
carbon.Parse("2020-08-05 13:14:15").VLunar().IsGoatYear() // false
// Whether is a lunar year of the monkey
carbon.Parse("2020-08-05 13:14:15").VLunar().IsMonkeyYear() // false
// Whether is a lunar year of the rooster
carbon.Parse("2020-08-05 13:14:15").VLunar().IsRoosterYear() // false
// Whether is a lunar year of the dog
carbon.Parse("2020-08-05 13:14:15").VLunar().IsDogYear() // false
// Whether is a lunar year of the dig
carbon.Parse("2020-08-05 13:14:15").VLunar().IsPigYear() // false

```

##### Convert Lunar into Solar

```go
// Convert the Vietnamese Lunar Calendar December 11, 2023 to the Solar Calendar
carbon.CreateFromVLunar(2023, 12, 11, 0, 0, 0, false).ToDateTimeString() // 2024-01-21 00:00:00
// Convert Vietnamese lunar calendar February 11, 2023 to Solar calendar
carbon.CreateFromVLunar(2023, 2, 11, 0, 0, 0, false).ToDateTimeString() // 2024-03-02 00:00:00
// Convert the Vietnamese Lunar Calendar Leap February 11, 2024 to the Solar Calendar
carbon.CreateFromVLunar(2023, 2, 11, 0, 0, 0, true).ToDateTimeString() // 2023-04-01 00:00:00
```