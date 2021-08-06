# Carbon  #
[![Carbon Release](https://img.shields.io/github/release/golang-module/carbon.svg)](https://github.com/golang-module/carbon/releases)
[![Build Status](https://github.com/golang-module/carbon/workflows/Go/badge.svg)](https://github.com/golang-module/carbon/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/golang-module/carbon)](https://goreportcard.com/report/github.com/golang-module/carbon)
[![Go Code Coverage](https://codecov.io/gh/golang-module/carbon/branch/master/graph/badge.svg)](https://codecov.io/gh/golang-module/carbon)
[![Carbon Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/golang-module/carbon)
![License](https://img.shields.io/github/license/golang-module/carbon)

日本語 | [English](README.md) | [简体中文](README.cn.md) 

軽量でセマンティックで開発者に優しい golang 時間処理ライブラリ

Carbon は [awesome-go](https://github.com/avelino/awesome-go#date-and-time "awesome-go") に収録されています, よかったら, スターをください

github:[github.com/golang-module/carbon](https://github.com/golang-module/carbon "github.com/golang-module/carbon")

gitee:[gitee.com/go-package/carbon](https://gitee.com/go-package/carbon "gitee.com/go-package/carbon")

#### インストール使用
```go
// github倉庫を使う
go get -u github.com/golang-module/carbon

import (
    "github.com/golang-module/carbon"
)

// gitee倉庫を使う
go get -u gitee.com/go-package/carbon

import (
    "gitee.com/go-package/carbon"
)
```

#### 使い方の例
> デフォルトのタイムゾーンはLocalです。つまりサーバのタイムゾーンです, 現在の時間は2020-08-05 13:14:15と仮定します

##### 昨日、今日、明日
```go
// 今日の瞬間
fmt.Sprintf("%s", carbon.Now()) // 2020-08-05 13:14:15
carbon.Now().ToDateTimeString() // 2020-08-05 13:14:15
// 今日の日付
carbon.Now().ToDateString() // 2020-08-05
// 今日の時間
carbon.Now().ToTimeString() // 13:14:15
// 今日は秒タイムスタンプ
carbon.Now().ToTimestamp() // 1596604455
carbon.Now().ToTimestampWithSecond() // 1596604455
// 今日のミリ秒タイムスタンプ
carbon.Now().ToTimestampWithMillisecond() // 1596604455000
// 今日のマイクロ秒タイムスタンプ
carbon.Now().ToTimestampWithMicrosecond() // 1596604455000000
// 今日のナノ秒タイムスタンプ
carbon.Now().ToTimestampWithNanosecond() // 1596604455000000000
// タイムゾーン指定の今日
carbon.Now(Carbon.NewYork).ToDateTimeString() // 2020-08-05 14:14:15
carbon.SetTimezone(Carbon.NewYork).Now().ToDateTimeString() // 2020-08-05 14:14:15

// 昨日の今は
fmt.Sprintf("%s", carbon.Yesterday()) // 2020-08-04 13:14:15
carbon.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
// 昨日の日付
carbon.Yesterday().ToDateString() // 2020-08-04
// 昨日の時間
carbon.Yesterday().ToTimeString() // 13:14:15
// 昨日の秒タイムスタンプ
carbon.Yesterday().ToTimestamp() // 1596518055
carbon.Yesterday().ToTimestampWithSecond() // 1596518055
// 昨日のミリ秒タイムスタンプ
carbon.Yesterday().ToTimestampWithMillisecond() // 1596518055000
// 昨日のマイクロ秒タイムスタンプ
carbon.Yesterday().ToTimestampWithMicrosecond() // 1596518055000000
// 昨日のナノ秒タイムスタンプ
carbon.Yesterday().ToTimestampWithNanosecond() // 1596518055000000000
// 日付指定の昨日
carbon.Parse("2021-01-28 13:14:15").Yesterday().ToDateTimeString() // 2021-01-27 13:14:15
// タイムゾーン指定の昨日
carbon.Yesterday(Carbon.NewYork).ToDateTimeString() // 2020-08-04 14:14:15
carbon.SetTimezone(Carbon.NewYork).Yesterday().ToDateTimeString() // 2020-08-04 14:14:15

// 明日の今は
fmt.Sprintf("%s", carbon.Tomorrow()) // 2020-08-06 13:14:15
carbon.Tomorrow().ToDateTimeString() // 2020-08-06 13:14:15
// 明日の日付
carbon.Tomorrow().ToDateString() // 2020-08-06
// 明日の時間
carbon.Tomorrow().ToTimeString() // 13:14:15
// 明日の秒タイムスタンプ
carbon.Tomorrow().ToTimestamp() // 1596690855
carbon.Tomorrow().ToTimestampWithSecond() // 1596690855
// 明日のミリ秒タイムスタンプ
carbon.Tomorrow().ToTimestampWithMillisecond() // 1596690855000
// 明日のマイクロ秒タイムスタンプ
carbon.Tomorrow().ToTimestampWithMicrosecond() // 1596690855000000
// 明日のナノ秒タイムスタンプ
carbon.Tomorrow().ToTimestampWithNanosecond() // 1596690855000000000
// 日付指定の明日
carbon.Parse("2021-01-28 13:14:15").Tomorrow().ToDateTimeString() // 2021-01-29 13:14:15
// タイムゾーン指定の明日
carbon.Tomorrow(Carbon.NewYork).ToDateTimeString() // 2020-08-06 14:14:15
carbon.SetTimezone(Carbon.NewYork).Tomorrow().ToDateTimeString() // 2020-08-06 14:14:15
```

##### Carbon オブジェクトを作成する
```go
// 秒タイムスタンプから Carbon オブジェクトを作成します
carbon.CreateFromTimestamp(-1).ToDateTimeString() // 1970-01-01 07:59:59
carbon.CreateFromTimestamp(-1, carbon.Tokyo).ToDateTimeString() // 1970-01-01 08:59:59
carbon.CreateFromTimestamp(0).ToDateTimeString() // 1970-01-01 08:00:00
carbon.CreateFromTimestamp(0, carbon.Tokyo).ToDateTimeString() // 1970-01-01 09:00:00
carbon.CreateFromTimestamp(1596604455).ToDateTimeString() // 2020-08-05 13:14:15
carbon.CreateFromTimestamp(1596604455, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// ミリ秒のタイムスタンプから Carbon オブジェクトを作成します
carbon.CreateFromTimestamp(1596604455000).ToDateTimeString() // 2020-08-05 13:14:15
carbon.CreateFromTimestamp(1596604455000, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// マイクロ秒タイムスタンプから Carbon オブジェクトを作成します
carbon.CreateFromTimestamp(1596604455000000).ToDateTimeString() // 2020-08-05 13:14:15
carbon.CreateFromTimestamp(1596604455000000, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// ナノタイムスタンプから Carbon オブジェクトを作成します
carbon.CreateFromTimestamp(1596604455000000000).ToDateTimeString() // 2020-08-05 13:14:15
carbon.CreateFromTimestamp(1596604455000000000, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15

// 年月日から分秒で Carbon オブジェクトを作成します
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// 年月日から Carbon オブジェクトを作成します(時分秒はデフォルトで現在の時分秒です)
carbon.CreateFromDate(2020, 8, 5).ToDateTimeString() // 2020-08-05 13:14:15
carbon.CreateFromDate(2020, 8, 5, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// 時分秒から Carbon オブジェクトを作成します(年月日のデフォルトは現在の年月日です)
carbon.CreateFromTime(13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
carbon.CreateFromTime(13, 14, 15, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### 標準形式の時間文字列を Carbon オブジェクトに解析します
```go
carbon.Parse("").ToDateTimeString() // 空の文字列
carbon.Parse("0").ToDateTimeString() // 空の文字列
carbon.Parse("0000-00-00 00:00:00").ToDateTimeString() // 空の文字列
carbon.Parse("0000-00-00").ToDateTimeString() // 空の文字列
carbon.Parse("00:00:00").ToDateTimeString() // 空の文字列

carbon.Parse("2020-08-05 13:14:15").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("2020-08-05").ToDateTimeString() // 2020-08-05 00:00:00
carbon.Parse("20200805131415").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("20200805").ToDateTimeString() // 2020-08-05 00:00:00
carbon.Parse("2020-08-05T13:14:15+08:00").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("2020-08-05 13:14:15", carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### 文字をフォーマットして文字列を Carbon オブジェクトに解析します
```go
carbon.ParseByFormat("2020|08|05 13|14|15", "Y|m|d H|i|s").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("2020-08-05 13:14:15", "Y-m-d H:i:s", carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### レイアウト文字を使用して文字列を Carbon オブジェクトに解析します
```go
carbon.ParseByLayout("2020|08|05 13|14|15", "2006|01|02 15|04|05").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("2020-08-05 13:14:15", "2006-01-02 15:04:05", carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### Carbon と time.Time 交換
```go
// time.Time を Carbon に変換します
carbon.Time2Carbon(time.Now())
// Carbon を time.Time に変換します
carbon.Now().Carbon2Time()
```

##### 始まりと終わり
```go
// 世紀の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfCentury().ToDateTimeString() // 2000-01-01 00:00:00
// 世紀の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfCentury().ToDateTimeString() // 2999-12-31 23:59:59

// 十年の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
carbon.Parse("2021-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
carbon.Parse("2029-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
// 十年の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59
carbon.Parse("2021-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59
carbon.Parse("2029-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59

// 今年の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfYear().ToDateTimeString() // 2020-01-01 00:00:00
// 今年の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfYear().ToDateTimeString() // 2020-12-31 23:59:59

// 季度の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfQuarter().ToDateTimeString() // 2020-07-01 00:00:00
// 季度の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfQuarter().ToDateTimeString() // 2020-09-30 23:59:59

// 本月の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfMonth().ToDateTimeString() // 2020-08-01 00:00:00
// 本月の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfMonth().ToDateTimeString() // 2020-08-31 23:59:59

// 本周の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfWeek(time.Sunday).ToDateTimeString() // 2020-08-02 00:00:00
carbon.Parse("2020-08-05 13:14:15").StartOfWeek(time.Monday).ToDateTimeString() // 2020-08-03 00:00:00
// 本周の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfWeek(time.Sunday).ToDateTimeString() // 2020-08-08 23:59:59
carbon.Parse("2020-08-05 13:14:15").EndOfWeek(time.Monday).ToDateTimeString() // 2020-08-09 23:59:59

// 本日の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfDay().ToDateTimeString() // 2020-08-05 00:00:00
// 本日の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfDay().ToDateTimeString() // 2020-08-05 23:59:59

// 時間の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfHour().ToDateTimeString() // 2020-08-05 13:00:00
// 時間の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfHour().ToDateTimeString() // 2020-08-05 13:59:59

// 分钟の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfMinute().ToDateTimeString() // 2020-08-05 13:14:00
// 分钟の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfMinute().ToDateTimeString() // 2020-08-05 13:14:59

// 本秒の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfSecond().Format("Y-m-d H:i:s.u") // 2020-08-05 13:14:15.0
// 本秒の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfSecond().Format("Y-m-d H:i:s.u") // 2020-08-05 13:14:15.999
```

##### 追加と減らす
```go
// 三ヶ世紀を追加
carbon.Parse("2020-02-29 13:14:15").AddCenturies(3).ToDateTimeString() // 2320-02-29 13:14:15
// 三ヶ世紀を追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddCenturiesNoOverflow(3).ToDateTimeString() // 2320-02-29 13:14:15
// 一ヶ世紀を追加
carbon.Parse("2020-02-29 13:14:15").AddCentury().ToDateTimeString() // 2120-02-29 13:14:15
// 一ヶ世紀を追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddCenturyNoOverflow().ToDateTimeString() // 2120-02-29 13:14:15
// 三ヶ世紀を減らす
carbon.Parse("2020-02-29 13:14:15").SubCenturies(3).ToDateTimeString() // 1720-02-29 13:14:15
// 三ヶ世紀を減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubCenturiesNoOverflow(3).ToDateTimeString() // 1720-02-29 13:14:15
// 一ヶ世紀を減らす
carbon.Parse("2020-02-29 13:14:15").SubCentury().ToDateTimeString() // 1920-02-29 13:14:15
// 一ヶ世紀を減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubCenturyNoOverflow().ToDateTimeString() // 1920-02-29 13:14:15

// 三ヶ年代を追加
carbon.Parse("2020-02-29 13:14:15").Decades(3).ToDateTimeString() // 2050-03-01 13:14:15
// 三ヶ年代を追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddDecadesNoOverflow(3).ToDateTimeString() // 2050-02-28 13:14:15
// 一ヶ年代を追加
carbon.Parse("2020-02-29 13:14:15").AddDecade().ToDateTimeString() // 2030-03-01 13:14:15
// 一ヶ年代を追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddDecadeNoOverflow().ToDateTimeString() // 2030-02-28 13:14:15
// 三ヶ年代を減らす
carbon.Parse("2020-02-29 13:14:15").SubDecades(3).ToDateTimeString() // 1990-03-01 13:14:15
// 三ヶ年代を減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubDecadesNoOverflow(3).ToDateTimeString() // 1990-02-28 13:14:15
// 一ヶ年代を減らす
carbon.Parse("2020-02-29 13:14:15").SubDecade().ToDateTimeString() // 2010-03-01 13:14:15
// 一ヶ年代を減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubDecadeNoOverflow().ToDateTimeString() // 2010-02-28 13:14:15

// 三か年を追加
carbon.Parse("2020-02-29 13:14:15").AddYears(3).ToDateTimeString() // 2023-03-01 13:14:15
// 三か年を追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddYearsNoOverflow(3).ToDateTimeString() // 2023-02-28 13:14:15
// 一か年を追加
carbon.Parse("2020-02-29 13:14:15").AddYear().ToDateTimeString() // 2021-03-01 13:14:15
// 一か年を追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddYearNoOverflow().ToDateTimeString() // 2021-02-28 13:14:15
// 三か年を減らす
carbon.Parse("2020-02-29 13:14:15").SubYears(3).ToDateTimeString() // 2017-03-01 13:14:15
// 三か年を減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubYearsNoOverflow(3).ToDateTimeString() // 2017-02-28 13:14:15
// 一か年を減らす
carbon.Parse("2020-02-29 13:14:15").SubYear().ToDateTimeString() // 2019-03-01 13:14:15
// 一か年を減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubYearNoOverflow().ToDateTimeString() // 2019-02-28 13:14:15

// 三ヶ四半期を追加
carbon.Parse("2019-08-31 13:14:15").AddQuarters(3).ToDateTimeString() // 2019-03-02 13:14:15
// 三ヶ四半期を追加(オーバーフローなし)
carbon.Parse("2019-08-31 13:14:15").AddQuartersNoOverflow(3).ToDateTimeString() // 2019-02-29 13:14:15
// 一ヶ四半期を追加
carbon.Parse("2019-11-30 13:14:15").AddQuarter().ToDateTimeString() // 2020-03-01 13:14:15
// 一ヶ四半期を追加(オーバーフローなし)
carbon.Parse("2019-11-30 13:14:15").AddQuarterNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15
// 三ヶ四半期を減らす
carbon.Parse("2019-08-31 13:14:15").SubQuarters(3).ToDateTimeString() // 2019-03-03 13:14:15
// 三ヶ四半期を減らす(オーバーフローなし)
carbon.Parse("2019-08-31 13:14:15").SubQuartersNoOverflow(3).ToDateTimeString() // 2019-02-28 13:14:15
// 一ヶ四半期を減らす
carbon.Parse("2020-05-31 13:14:15").SubQuarter().ToDateTimeString() // 2020-03-02 13:14:15
// 一ヶ四半期を減らす(オーバーフローなし)
carbon.Parse("2020-05-31 13:14:15").SubQuarterNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

// 三ヶ月を追加
carbon.Parse("2020-02-29 13:14:15").AddMonths(3).ToDateTimeString() // 2020-05-29 13:14:15
// 三ヶ月を追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddMonthsNoOverflow(3).ToDateTimeString() // 2020-05-29 13:14:15
// 一ヶ月を追加
carbon.Parse("2020-01-31 13:14:15").AddMonth().ToDateTimeString() // 2020-03-02 13:14:15
// 一ヶ月を追加(オーバーフローなし)
carbon.Parse("2020-01-31 13:14:15").AddMonthNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15
// 三ヶ月を減らす
carbon.Parse("2020-02-29 13:14:15").SubMonths(3).ToDateTimeString() // 2019-11-29 13:14:15
// 三ヶ月を減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubMonthsNoOverflow(3).ToDateTimeString() // 2019-11-29 13:14:15
// 一ヶ月を減らす
carbon.Parse("2020-03-31 13:14:15").SubMonth().ToDateTimeString() // 2020-03-02 13:14:15
// 一か月を減らす(オーバーフローなし)
carbon.Parse("2020-03-31 13:14:15").SubMonthNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

// 三か週間を追加
carbon.Parse("2020-02-29 13:14:15").AddWeeks(3).ToDateTimeString() // 2020-03-21 13:14:15
// 一か週間を追加
carbon.Parse("2020-02-29 13:14:15").AddWeek().ToDateTimeString() // 2020-03-07 13:14:15
// 三か週間を減らす
carbon.Parse("2020-02-29 13:14:15").SubWeeks(3).ToDateTimeString() // 2020-02-08 13:14:15
// 一か週間を減らす
carbon.Parse("2020-02-29 13:14:15").SubWeek().ToDateTimeString() // 2020-02-22 13:14:15

// 三か日間を追加
carbon.Parse("2020-08-05 13:14:15").AddDays(3).ToDateTimeString() // 2020-08-08 13:14:15
// 一か日間を追加
carbon.Parse("2020-08-05 13:14:15").AddDay().ToDateTimeString() // 2020-08-05 13:14:15
// 三か日間を減らす
carbon.Parse("2020-08-05 13:14:15").SubDays(3).ToDateTimeString() // 2020-08-02 13:14:15
// 一か日間を減らす
carbon.Parse("2020-08-05 13:14:15").SubDay().ToDateTimeString() // 2020-08-04 13:14:15

// 三か時間を追加
carbon.Parse("2020-08-05 13:14:15").AddHours(3).ToDateTimeString() // 2020-08-05 16:14:15
// 二か時間半を追加
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5h").ToDateTimeString() // 2020-08-05 15:44:15
carbon.Parse("2020-08-05 13:14:15").AddDuration("2h30m").ToDateTimeString() // 2020-08-05 15:44:15
// 一か時間を追加
carbon.Parse("2020-08-05 13:14:15").AddHour().ToDateTimeString() // 2020-08-05 14:14:15
// 三か時間を減らす
carbon.Parse("2020-08-05 13:14:15").SubHours(3).ToDateTimeString() // 2020-08-05 10:14:15
// 二か時間半を減らす
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5h").ToDateTimeString() // 2020-08-05 10:44:15
carbon.Parse("2020-08-05 13:14:15").SubDuration("2h30m").ToDateTimeString() // 2020-08-05 10:44:15
// 一か時間を減らす
carbon.Parse("2020-08-05 13:14:15").SubHour().ToDateTimeString() // 2020-08-05 12:14:15

// 三か分钟を追加
carbon.Parse("2020-08-05 13:14:15").AddMinutes(3).ToDateTimeString() // 2020-08-05 13:17:15
// 二か分钟半を追加
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5m").ToDateTimeString() // 2020-08-05 13:16:45
carbon.Parse("2020-08-05 13:14:15").AddDuration("2m30s").ToDateTimeString() // 2020-08-05 13:16:45
// 一か分钟を追加
carbon.Parse("2020-08-05 13:14:15").AddMinute().ToDateTimeString() // 2020-08-05 13:15:15
// 三か分钟を減らす
carbon.Parse("2020-08-05 13:14:15").SubMinutes(3).ToDateTimeString() // 2020-08-05 13:11:15
// 二か分钟半を減らす
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5m").ToDateTimeString() // 2020-08-05 13:11:45
carbon.Parse("2020-08-05 13:14:15").SubDuration("2m30s").ToDateTimeString() // 2020-08-05 13:11:45
// 一か分钟を減らす
carbon.Parse("2020-08-05 13:14:15").SubMinute().ToDateTimeString() // 2020-08-05 13:13:15

// 三か秒钟を追加
carbon.Parse("2020-08-05 13:14:15").AddSeconds(3).ToDateTimeString() // 2020-08-05 13:14:18
// 二か秒钟半を追加
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5s").ToDateTimeString() // 2020-08-05 13:14:17
// 一か秒钟を追加
carbon.Parse("2020-08-05 13:14:15").AddSecond().ToDateTimeString() // 2020-08-05 13:14:16
// 三か秒钟を減らす
carbon.Parse("2020-08-05 13:14:15").SubSeconds(3).ToDateTimeString() // 2020-08-05 13:14:12
// 二か秒钟半を減らす
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5s").ToDateTimeString() // 2020-08-05 13:14:12
// 一か秒钟を減らす
carbon.Parse("2020-08-05 13:14:15").SubSecond().ToDateTimeString() // 2020-08-05 13:14:14
```

##### 时间差異
```go
// 何年違いますか
carbon.Parse("2021-08-05 13:14:15").DiffInYears(carbon.Parse("2020-08-05 13:14:15")) // -1
// 何年違いますか（絶対値）
carbon.Parse("2021-08-05 13:14:15").DiffInYearsWithAbs(carbon.Parse("2020-08-05 13:14:15")) // 1

// 何ヶ月違いますか
carbon.Parse("2020-08-05 13:14:15").DiffInMonths(carbon.Parse("2020-07-05 13:14:15")) // -1
// 何ヶ月違いますか（絶対値）
carbon.Parse("2020-08-05 13:14:15").DiffInMonthsWithAbs(carbon.Parse("2020-07-05 13:14:15")) // 1

// 何週間違いますか
carbon.Parse("2020-08-05 13:14:15").DiffInWeeks(carbon.Parse("2020-07-28 13:14:15")) // -1
// 何週間違いますか（絶対値）
carbon.Parse("2020-08-05 13:14:15").DiffInWeeksWithAbs(carbon.Parse("2020-07-28 13:14:15")) // 1

// 何日間違いますか
carbon.Parse("2020-08-05 13:14:15").DiffInDays(carbon.Parse("2020-08-04 13:14:15")) // -1
// 何日間違いますか（絶対値）
carbon.Parse("2020-08-05 13:14:15").DiffInDaysWithAbs(carbon.Parse("2020-08-04 13:14:15")) // 1

// 何時間違いますか
carbon.Parse("2020-08-05 13:14:15").DiffInHours(carbon.Parse("2020-08-05 12:14:15")) // -1
// 何時間違いますか（絶対値）
carbon.Parse("2020-08-05 13:14:15").DiffInHoursWithAbs(carbon.Parse("2020-08-05 12:14:15")) // 1

// 何分違いますか
carbon.Parse("2020-08-05 13:14:15").DiffInMinutes(carbon.Parse("2020-08-05 13:13:15")) // -1
// 何分違いますか（絶対値）
carbon.Parse("2020-08-05 13:14:15").DiffInMinutesWithAbs(carbon.Parse("2020-08-05 13:13:15")) // 1

// 何秒の差がありますか
carbon.Parse("2020-08-05 13:14:15").DiffInSeconds(carbon.Parse("2020-08-05 13:14:14")) // -1
// 何秒の差がありますか（絶対値）
carbon.Parse("2020-08-05 13:14:15").DiffInSecondsWithAbs(carbon.Parse("2020-08-05 13:14:14")) // 1

// 人間に優しい読み取り可能なフォーマットの時間差を取得します
carbon.Parse("2020-08-05 13:14:15").DiffForHumans()) // just now
carbon.Parse("2019-08-05 13:14:15").DiffForHumans() // 1 year ago
carbon.Parse("2018-08-05 13:14:15").DiffForHumans() // 2 years ago
carbon.Parse("2021-08-05 13:14:15").DiffForHumans() // 1 year from now
carbon.Parse("2022-08-05 13:14:15").DiffForHumans() // 2 years from now

carbon.Parse("2020-08-05 13:14:15").DiffForHumans(carbon.Now()) // 1 year before
carbon.Parse("2019-08-05 13:14:15").DiffForHumans(carbon.Now()) // 2 years before
carbon.Parse("2018-08-05 13:14:15").DiffForHumans(carbon.Now()) // 1 year after
carbon.Parse("2022-08-05 13:14:15").DiffForHumans(carbon.Now()) // 2 years after
```

##### 时间比較
```go
// ゼロ時間ですか
carbon.Parse("").IsZero() // true
carbon.Parse("0").IsZero() // true
carbon.Parse("0000-00-00 00:00:00").IsZero() // true
carbon.Parse("0000-00-00").IsZero() // true
carbon.Parse("00:00:00").IsZero() // true
carbon.Parse("2020-08-05 00:00:00").IsZero() // false
carbon.Parse("2020-08-05").IsZero() // false
carbon.Parse("2020-08-05").SetTimezone("xxx").IsZero() // false

// 無効な時間ですか
carbon.Parse("").IsInvalid() // true
carbon.Parse("0").IsInvalid() // true
carbon.Parse("0000-00-00 00:00:00").IsInvalid() // true
carbon.Parse("0000-00-00").IsInvalid() // true
carbon.Parse("00:00:00").IsInvalid() // true
carbon.Parse("2020-08-05 00:00:00").IsInvalid() // false
carbon.Parse("2020-08-05").IsInvalid() // false
carbon.Parse("2020-08-05").SetTimezone("xxx").IsInvalid() // true

// 現在かどうか
carbon.Now().IsNow() // true
// 未来かどうか
carbon.Tomorrow().IsFuture() // true
// 過去かどうか
carbon.Yesterday().IsPast() // true

// 閏年かどうか
carbon.Parse("2020-08-05 13:14:15").IsLeapYear() // true
// 長年ですか
carbon.Parse("2020-08-05 13:14:15").IsLongYear() // true

// 一月ですか
carbon.Parse("2020-08-05 13:14:15").IsJanuary() // false
// 二月ですか
carbon.Parse("2020-08-05 13:14:15").IsFebruary() // false
// 三月ですか
carbon.Parse("2020-08-05 13:14:15").IsMarch() // false
// 四月ですか
carbon.Parse("2020-08-05 13:14:15").IsApril()  // false
// 五月ですか
carbon.Parse("2020-08-05 13:14:15").IsMay() // false
// 六月ですか
carbon.Parse("2020-08-05 13:14:15").IsJune() // false
// 七月ですか
carbon.Parse("2020-08-05 13:14:15").IsJuly() // false
// 八月ですか
carbon.Parse("2020-08-05 13:14:15").IsAugust() // false
// 八月ですか
carbon.Parse("2020-08-05 13:14:15").IsSeptember() // true
// 十月ですか
carbon.Parse("2020-08-05 13:14:15").IsOctober() // false
// 十一月ですか
carbon.Parse("2020-08-05 13:14:15").IsNovember() // false
// 十二月ですか
carbon.Parse("2020-08-05 13:14:15").IsDecember() // false

// 月曜日ですか
carbon.Parse("2020-08-05 13:14:15").IsMonday() // false
// 火曜日ですか
carbon.Parse("2020-08-05 13:14:15").IsTuesday() // true
// 水曜日ですか
carbon.Parse("2020-08-05 13:14:15").IsWednesday() // false
// 木曜日ですか
carbon.Parse("2020-08-05 13:14:15").IsThursday()  // false
// 金曜日ですか
carbon.Parse("2020-08-05 13:14:15").IsFriday() // false
// 土曜日ですか
carbon.Parse("2020-08-05 13:14:15").IsSaturday() // false
// 日曜日ですか
carbon.Parse("2020-08-05 13:14:15").IsSunday() // false

// 営業日ですか
carbon.Parse("2020-08-05 13:14:15").IsWeekday() // false
// 週末ですか
carbon.Parse("2020-08-05 13:14:15").IsWeekend() // true

// 昨日ですか
carbon.Parse("2020-08-04 13:14:15").IsYesterday() // true
carbon.Parse("2020-08-04 00:00:00").IsYesterday() // true
carbon.Parse("2020-08-04").IsYesterday() // true
// 今日ですか
carbon.Parse("2020-08-05 13:14:15").IsToday() // true
carbon.Parse("2020-08-05 00:00:00").IsToday() // true
carbon.Parse("2020-08-05").IsToday() // true
// 明日ですか
carbon.Parse("2020-08-06 13:14:15").IsTomorrow() // true
carbon.Parse("2020-08-06 00:00:00").IsTomorrow() // true
carbon.Parse("2020-08-06").IsTomorrow() // true

// 大きいかどうか
carbon.Parse("2020-08-05 13:14:15").Gt(carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Gt(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare(">", carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">", carbon.Parse("2020-08-05 13:14:15")) // false

// 小さいかどうか
carbon.Parse("2020-08-05 13:14:15").Lt(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Lt(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("<", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<", carbon.Parse("2020-08-05 13:14:15")) // false

// 等しいかどうか
carbon.Parse("2020-08-05 13:14:15").Eq(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Eq(carbon.Parse("2020-08-05 13:14:00")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("=", carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("=", carbon.Parse("2020-08-05 13:14:00")) // false

// と等しくない
carbon.Parse("2020-08-05 13:14:15").Ne(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Ne(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("!=", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<>", carbon.Parse("2020-08-05 13:14:15")) // false

// 大きいか等しいかどうか
carbon.Parse("2020-08-05 13:14:15").Gte(carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Gte(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">=", carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">=", carbon.Parse("2020-08-05 13:14:15")) // true

// 小きいか等しいかどうか
carbon.Parse("2020-08-05 13:14:15").Lte(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Lte(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<=", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<=", carbon.Parse("2020-08-05 13:14:15")) // true

// 二つの時間の間に(この二つの時間は含まれていません)
carbon.Parse("2020-08-05 13:14:15").Between(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Between(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// 二つの時間の間に(開始時間も含めて)
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedStart(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedStart(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// 二つの時間の間に(終了時間も含めて)
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// 二つの時間の間に(この二つの時間を含めて)
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-05 13:14:15")) // true
```

##### 时间出力
```go
// 秒タイムスタンプを出力, ToTimestamp() は ToTimestampWithSecond() の略記です
carbon.Parse("2020-08-05 13:14:15").ToTimestamp() // 1596604455
carbon.Parse("2020-08-05 13:14:15").ToTimestampWithSecond() // 1596604455
// ミリ秒のタイムスタンプを出力
carbon.Parse("2020-08-05 13:14:15").ToTimestampWithMillisecond() // 1596604455000
// マイクロ秒タイムスタンプを出力
carbon.Parse("2020-08-05 13:14:15").ToTimestampWithMicrosecond() // 1596604455000000
// ナノ秒タイムスタンプを出力
carbon.Parse("2020-08-05 13:14:15").ToTimestampWithNanosecond() // 1596604455000000000

// 日期时间文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("2020-08-05 13:14:15").ToDateTimeString(carbon.Tokyo) // 2020-08-05 14:14:15
// 略語日期时间文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToShortDateTimeString() // 20200805131415
carbon.Parse("2020-08-05 13:14:15").ToShortDateTimeString(carbon.Tokyo) // 20200805141415

// 日期文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToDateString() // 2020-08-05
carbon.Parse("2020-08-05 13:14:15").ToDateString(carbon.Tokyo) // 2020-08-05
// 略語日期文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToShortDateString() // 20200805
carbon.Parse("2020-08-05 13:14:15").ToShortDateString(carbon.Tokyo) // 20200805

// 時間文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToTimeString() // 13:14:15
carbon.Parse("2020-08-05 13:14:15").ToTimeString(carbon.Tokyo) // 14:14:15
// 略語時間文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToShortTimeString() // 131415
carbon.Parse("2020-08-05 13:14:15").ToShortTimeString(carbon.Tokyo) // 141415

// Ansic フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToAnsicString() // Wed Aug  5 13:14:15 2020
carbon.Parse("2020-08-05 13:14:15").ToAnsicString(carbon.Tokyo) // Wed Aug  5 14:14:15 2020
// Atom フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToAtomString() // 2020-08-05T13:14:15+08:00
carbon.Parse("2020-08-05 13:14:15").ToAtomString(carbon.Tokyo) // 2020-08-05T14:14:15+08:00
// UnixDate フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToUnixDateString() // Wed Aug  5 13:14:15 CST 2020
carbon.Parse("2020-08-05 13:14:15").ToUnixDateString(carbon.Tokyo) // Wed Aug  5 14:14:15 JST 2020
// RubyDate フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToRubyDateString() // Wed Aug 05 13:14:15 +0800 2020
carbon.Parse("2020-08-05 13:14:15").ToRubyDateString(carbon.Tokyo) // Wed Aug 05 14:14:15 +0900 2020
// Kitchen フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToKitchenString() // 1:14PM
carbon.Parse("2020-08-05 13:14:15").ToKitchenString(carbon.Tokyo) // 2:14PM
// Cookie フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToCookieString() // Wednesday, 05-Aug-2020 13:14:15 CST
carbon.Parse("2020-08-05 13:14:15").ToCookieString(carbon.Tokyo) // Wednesday, 05-Aug-2020 14:14:15 JST
// DayDateTime フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToDayDateTimeString() // Wed, Aug 5, 2020 1:14 PM
carbon.Parse("2020-08-05 13:14:15").ToDayDateTimeString(carbon.Tokyo) // Wed, Aug 5, 2020 2:14 PM
// RSS フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToRssString() // Wed, 05 Aug 2020 13:14:15 +0800
carbon.Parse("2020-08-05 13:14:15").ToRssString(carbon.Tokyo) // Wed, 05 Aug 2020 14:14:15 +0900
// W3C フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToW3cString() // 2020-08-05T13:14:15+08:00
carbon.Parse("2020-08-05 13:14:15").ToW3cString(carbon.Tokyo) // 2020-08-05T14:14:15+09:00

// ISO8601 フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToIso8601String() // 2020-08-05T13:14:15+08:00
carbon.Parse("2020-08-05 13:14:15").ToIso8601String(carbon.Tokyo) // 2020-08-05T14:14:15+09:00
// RFC822 フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToRfc822String() // 05 Aug 20 13:14 CST
carbon.Parse("2020-08-05 13:14:15").ToRfc822String(carbon.Tokyo) // 05 Aug 20 14:14 JST
// RFC822Z フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToRfc822zString() // 05 Aug 20 13:14 +0800
carbon.Parse("2020-08-05 13:14:15").ToRfc822zString(carbon.Tokyo) // 05 Aug 20 14:14 +0900
// RFC850 フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToRfc850String() // Wednesday, 05-Aug-20 13:14:15 CST
carbon.Parse("2020-08-05 13:14:15").ToRfc850String(carbon.Tokyo) // Wednesday, 05-Aug-20 14:14:15 JST
// RFC1036 フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToRfc1036String() // Wed, 05 Aug 20 13:14:15 +0800
carbon.Parse("2020-08-05 13:14:15").ToRfc1036String(carbon.Tokyo) // Wed, 05 Aug 20 14:14:15 +0900
// RFC1123 フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToRfc1123String() // Wed, 05 Aug 2020 13:14:15 CST
carbon.Parse("2020-08-05 13:14:15").ToRfc1123String(carbon.Tokyo) // Wed, 05 Aug 2020 14:14:15 JST
// RFC1123Z フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToRfc1123zString() // Wed, 05 Aug 2020 13:14:15 +0800
carbon.Parse("2020-08-05 13:14:15").ToRfc1123zString(carbon.Tokyo) // Wed, 05 Aug 2020 14:14:15 0800
// RFC2822 フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToRfc2822String() // Wed, 05 Aug 2020 13:14:15 +0800
carbon.Parse("2020-08-05 13:14:15").ToRfc2822String(carbon.Tokyo) // Wed, 05 Aug 2020 14:14:15 +0900
// RFC3339 フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToRfc3339String() // 2020-08-05T13:14:15+08:00
carbon.Parse("2020-08-05 13:14:15").ToRfc3339String(carbon.Tokyo) // 2020-08-05T14:14:15+09:00
// RFC7231 フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToRfc7231String() // Wed, 05 Aug 2020 13:14:15 GMT
carbon.Parse("2020-08-05 13:14:15").ToRfc7231String(carbon.Tokyo) // Wed, 05 Aug 2020 14:14:15 GMT

// 日付時間文字列を出力
fmt.Sprintf("%s", carbon.Parse("2020-08-05 13:14:15")) // 2020-08-05 13:14:15
fmt.Sprintf("%s", carbon.Parse("2020-08-05 13:14:15", carbon.Tokyo)) // 2020-08-05 13:14:15

// "2006-01-02 15:04:05.999999999 -0700 MST" フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-08-05 13:14:15").ToString(carbon.Tokyo) // 2020-08-05 14:14:15 +0900 JST

// レイアウトを指定する文字列を出力, Layout() は 是ToLayoutString() の略記です
carbon.Parse("2020-08-05 13:14:15").Layout("20060102150405") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").Layout("2006年01月02日 15时04分05秒") // 2020年08月05日 13时14分15秒
carbon.Parse("2020-08-05 13:14:15").Layout("It is 2006-01-02 15:04:05") // It is 2020-08-05 13:14:15
carbon.Parse("2020-08-05 13:14:15").Layout("2006-01-02 15:04:05", carbon.Tokyo) // 2020-08-05 14:14:15

// 指定されたフォーマットの文字列を出力, Format() は 是ToFormatString() の略記です
carbon.Parse("2020-08-05 13:14:15").Format("YmdHis") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").Format("Y年m月d日 H时i分s秒") // 2020年08月05日 13时14分15秒
carbon.Parse("2020-08-05 13:14:15").Format("l jS \\o\\f F Y h:i:s A") // Wednesday 5th of August 2020 01:14:15 PM
carbon.Parse("2020-08-05 13:14:15").Format("\\I\\t \\i\\s Y-m-d H:i:s") // It is 2020-08-05 13:14:15
carbon.Parse("2020-08-05 13:14:15").Format("Y-m-d H:i:s", carbon.Tokyo) // 2020-08-05 14:14:15
```
>もっとフォーマットした出力記号は付録を見てください <a href="#format-sign-table">書式設定記号表</a>

##### 时间取得
```go
// 本年の総日数を取得
carbon.Parse("2019-08-05 13:14:15").DaysInYear() // 365
carbon.Parse("2020-08-05 13:14:15").DaysInYear() // 366
// 今月の総日数を取得
carbon.Parse("2020-02-01 13:14:15").DaysInMonth() // 29
carbon.Parse("2020-04-01 13:14:15").DaysInMonth() // 30
carbon.Parse("2020-08-01 13:14:15").DaysInMonth() // 31

// 本年の第数日を取得
carbon.Parse("2020-08-05 13:14:15").DayOfYear() // 218
// 本年の第数週を取得
carbon.Parse("2019-12-31 13:14:15").WeekOfYear() // 1
carbon.Parse("2020-08-05 13:14:15").WeekOfYear() // 32

// 今月の何日目（1から）を取得
carbon.Parse("2020-08-05 13:14:15").DayOfMonth() // 5
// 今月の何週目（1から）を取得
carbon.Parse("2020-08-05 13:14:15").WeekOfMonth() // 1
// 今月の何週目（1から）を取得
carbon.Parse("2020-08-05 13:14:15").DayOfWeek() // 3
carbon.Parse("2020-08-09 13:14:15").Week() // 7

// 現在の世紀を取得
carbon.Parse("2020-08-05 13:14:15").Century() // 21
// 現在の年代を取得
carbon.Parse("2019-08-05 13:14:15").Decade() // 10
carbon.Parse("2021-08-05 13:14:15").Decade() // 20
// 現在の年を取得
carbon.Parse("2020-08-05 13:14:15").Year() // 2020
// 現在の四半期を取得
carbon.Parse("2020-08-05 13:14:15").Quarter() // 3
// 現在の月を取得
carbon.Parse("2020-08-05 13:14:15").Month() // 8
// 現在の週を取得(0から開始)
carbon.Parse("2020-08-05 13:14:15").Week() // 3
carbon.Parse("2020-08-09 13:14:15").Week() // 0
// 現在の日数を取得
carbon.Parse("2020-08-05 13:14:15").Day() // 5
// 現在の時間を取得
carbon.Parse("2020-08-05 13:14:15").Hour() // 13
// 現在の分を取得
carbon.Parse("2020-08-05 13:14:15").Minute() // 14
// 現在の秒を取得
carbon.Parse("2020-08-05 13:14:15").Second() // 15
// 現在のミリ秒を取得
carbon.Parse("2020-08-05 13:14:15").Millisecond() // 1596604455000
// 現在のマイクロ秒を取得
carbon.Parse("2020-08-05 13:14:15").Microsecond() // 1596604455000000
// 現在のナノ秒を取得
carbon.Parse("2020-08-05 13:14:15").Nanosecond() // 1596604455000000000

// タイムゾーン名を取得
carbon.SetTimezone(carbon.PRC).Timezone() // CST
carbon.SetTimezone(carbon.Tokyo).Timezone() // JST

// ロケーション名を取得
carbon.SetTimezone(carbon.PRC).Location() // PRC
carbon.SetTimezone(carbon.Tokyo).Location() // Asia/Tokyo

// UTCタイムゾーンからのオフセットを取得、単位秒
carbon.SetTimezone(carbon.PRC).Offset() // 28800
carbon.SetTimezone(carbon.Tokyo).Offset() // 32400

// ロケール名を取得
carbon.Now().Locale() // en
carbon.Now().SetLocale("zh-CN").Locale() // zh-CN

// 星座を取得
carbon.Now().Constellation() // Leo
carbon.Now().SetLocale("en").Constellation() // Leo
carbon.Now().SetLocale("zh-CN").Constellation() // 狮子座

// 季節を取得
carbon.Now().Season() // Summer
carbon.Now().SetLocale("en").Season() // Summer
carbon.Now().SetLocale("zh-CN").Season() // 夏季

// 年齢を取得
carbon.Parse("2002-01-01 13:14:15").Age() // 17
carbon.Parse("2002-12-31 13:14:15").Age() // 18
```

##### 时间設定
```go
// タイムゾーンを設定
carbon.SetTimezone(carbon.PRC).Now().ToDateTimeString() // 2020-08-05 13:14:15
carbon.SetTimezone(carbon.Tokyo).Now().ToDateTimeString() // 2020-08-05 14:14:15
carbon.SetTimezone(carbon.Tokyo).Now().SetTimezone(carbon.PRC).ToDateTimeString() // 2020-08-05 12:14:15

// 設定ロケール
carbon.Parse("2020-07-05 13:14:15").SetLocale("en").DiffForHumans()) // 1 month ago
carbon.Parse("2020-07-05 13:14:15").SetLocale("zh-CN").DiffForHumans() // 1 月前

// 設定年
carbon.Parse("2020-02-29").SetYear(2021).ToDateString() // 2021-03-01
// 設定年(オーバーフローなし)
carbon.Parse("2020-02-29").SetYearNoOverflow(2021).ToDateString() // 2021-02-28

// 設定月
carbon.Parse("2020-01-31").SetMonth(2).ToDateString() // 2020-03-02
// 設定月(オーバーフローなし)
carbon.Parse("2020-01-31").SetMonthNoOverflow(2).ToDateString() // 2020-02-29

// 設定日
carbon.Parse("2019-08-05").SetDay(31).ToDateString() // 2020-08-31
carbon.Parse("2020-02-01").SetDay(31).ToDateString() // 2020-03-02

// 設定時
carbon.Parse("2020-08-05 13:14:15").SetHour(10).ToDateTimeString() // 2020-08-05 10:14:15
carbon.Parse("2020-08-05 13:14:15").SetHour(24).ToDateTimeString() // 2020-08-06 00:14:15

// 設定分
carbon.Parse("2020-08-05 13:14:15").SetMinute(10).ToDateTimeString() // 2020-08-05 13:10:15
carbon.Parse("2020-08-05 13:14:15").SetMinute(60).ToDateTimeString() // 2020-08-05 14:00:15

// 設定秒
carbon.Parse("2020-08-05 13:14:15").SetSecond(10).ToDateTimeString() // 2020-08-05 13:14:10
carbon.Parse("2020-08-05 13:14:15").SetSecond(60).ToDateTimeString() // 2020-08-05 13:15:00
```

##### 星座
```go
// 星座を取得
carbon.Parse("2020-08-05 13:14:15").Constellation() // Leo

// 牡羊座ですか
carbon.Parse("2020-08-05 13:14:15").IsAries() // false
// おうし座ですか
carbon.Parse("2020-08-05 13:14:15").IsTaurus() // false
// 双子座ですか
carbon.Parse("2020-08-05 13:14:15").IsGemini() // false
// かに座ですか
carbon.Parse("2020-08-05 13:14:15").IsCancer() // false
// 獅子座ですか
carbon.Parse("2020-08-05 13:14:15").IsLeo() // true
// おとめ座ですか
carbon.Parse("2020-08-05 13:14:15").IsVirgo() // false
// 天秤座ですか
carbon.Parse("2020-08-05 13:14:15").IsLibra() // false
// さそり座ですか
carbon.Parse("2020-08-05 13:14:15").IsScorpio() // false
// 射手座ですか
carbon.Parse("2020-08-05 13:14:15").IsSagittarius() // false
// 山羊座ですか
carbon.Parse("2020-08-05 13:14:15").IsCapricorn() // false
// 水瓶座ですか
carbon.Parse("2020-08-05 13:14:15").IsAquarius() // false
// 魚座ですか
carbon.Parse("2020-08-05 13:14:15").IsPisces() // false
```

##### 季節
> 気象区分によると、3-5月は春で、6-8月は夏で、9-11月は秋で、12-2月は冬です

```go
// シーズンを取得
carbon.Parse("2020-08-05 13:14:15").Season() // Summer

// この季節の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfSeason().ToDateTimeString() // 2020-06-01 00:00:00
// この季節の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfSeason().ToDateTimeString() // 2020-08-31 23:59:59

// 春かどうか
carbon.Parse("2020-08-05 13:14:15").IsSpring() // false
// 夏かどうか
carbon.Parse("2020-08-05 13:14:15").IsSummer() // true
// 秋かどうか
carbon.Parse("2020-08-05 13:14:15").IsAutumn() // false
// 冬かどうか
carbon.Parse("2020-08-05 13:14:15").IsWinter() // false
```

#####  中国の旧暦
> 現在は西暦`1900`年`2100`年の`200`年スパンだけをサポートしています

```go
// 干支を取得します
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().Animal() // 鼠

// 中国の旧暦の祝日を獲得します
carbon.Parse("2021-02-12 13:14:15", carbon.PRC).Lunar().Festival() // 春节

// 中国の旧正月を取得する
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().Year() // 2020
// 中国の太陰月を取得する
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().Month() // 6
// 中国の旧暦のうるう月を取得する
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().LeapMonth() // 4
// 中国の太陰暦を取得する
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().Day() // 16
// 中国の旧正月 YYYY-MM-DD フォーマット文字列を取得します
fmt.Sprintf("%s", carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar()) // 2020-06-16

// 中国の旧正月文字列を取得します
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().ToYearString() // 二零二零
// 中国の旧正月文字列を取得します
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().ToMonthString() // 六
// 中国の旧正月の日文字列を取得します
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().ToDayString() // 十六
// 中国の旧正月日付文字列を取得します
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().ToDateString() // 二零二零年六月十六

// 中国の旧正月の閏年ですか
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsLeapYear() // true
// 中国の旧暦の閏月かどうか
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsLeapMonth() // false

// ねずみ年かどうか
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsRatYear() // true
// 牛年かどうか
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsOxYear() // false
// 寅年かどうか
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsTigerYear() // false
// うさぎ年かどうか
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsRabbitYear() // false
// 龍年かどうか
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsDragonYear() // false
// 蛇の年かどうか
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsSnakeYear() // false
// 馬年かどうか
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsHorseYear() // false
// 羊年かどうか
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsGoatYear() // false
// 申年かどうか
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsMonkeyYear() // false
// 鶏の年かどうか
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsRoosterYear() // false
// 犬年かどうか
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsDogYear() // false
// 豚年かどうか
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsPigYear() // false
```

##### JSON 処理
###### 定義モデル
```go
type Person struct {
    ID  int64  `json:"id"`
    Name string `json:"name"`
    Age int `json:"age"`
    Birthday carbon.ToDateTimeString `json:"birthday"`
    GraduatedAt carbon.ToDateString `json:"graduated_at"`
    CreatedAt carbon.ToTimeString `json:"created_at"`
    UpdatedAt carbon.ToTimestamp `json:"updated_at"`
    DateTime1 carbon.ToTimestampWithSecond `json:"date_time1"`
    DateTime2 carbon.ToTimestampWithMillisecond `json:"date_time2"`
    DateTime3 carbon.ToTimestampWithMicrosecond `json:"date_time3"`
    DateTime4 carbon.ToTimestampWithNanosecond `json:"date_time4"`
}
```

###### 初期化モデル
```go
person := Person {
	ID:          1,
	Name:        "gouguoyin",
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
```

###### JSON コーディング
```go
data, err := json.Marshal(&person)
if err != nil {
	t.Fatal(err)
}
fmt.Printf("%s",data)
// 出力
{
	"id":1,
	"name":"gouguoyin",
	"age":18,
	"birthday":"2003-07-16 16:22:02",
	"graduated_at":"2020-08-05",
	"created_at":"13:14:15",
	"updated_at":1659676455,
	"date_time1":1691212455,
	"date_time2":1722834855000,
	"date_time3":1754370855000000,
	"date_time4":1754370855000000000
}
```

###### JSON 復号
```go
str := `{
	"id":1,
	"name":"gouguoyin",
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
fmt.Printf("%+v", *person)
// 出力
{ID:1 Name:gouguoyin Age:18 Birthday:2003-07-16 16:22:02 GraduatedAt:2020-08-05 00:00:00 CreatedAt:0000-01-01 13:14:15 UpdatedAt:2022-08-05 13:14:15 DTime1:2023-08-05 13:14:15 DateTime2:2024-08-05 13:14:15 DateTime3:2025-08-05 13:14:15 DateTime4:2025-08-05 13:14:15}
```

##### 国際化サポート

現在サポートされている言語
* [简体中国語(zh-CN)](./lang/zh-CN.json "简体中国語")
* [繁体中国語(zh-TW)](./lang/zh-TW.json "繁体中国語")
* [英語(en)](./lang/en.json "英語")
* [日本語(jp)](./lang/jp.json "日本語")
* [韓国語(kr)](./lang/kr.json "韓国語")
* [スペイン語(es)](./lang/es.json "スペイン語")：[hgisinger](https://github.com/hgisinger "hgisinger") から翻訳されます
* [German(de)](./lang/de.json "German")：[benzammour](https://github.com/benzammour "benzammour") から翻訳されます

現在サポートされている方法
* `Constellation()`：星座を取得
* `Season()`：シーズンを取得
* `DiffForHumans()`：人間に優しい読み取り可能なフォーマットの時間差を取得します
* `ToMonthString()`：月文字列を出力
* `ToShortMonthString()`：略語月文字列を出力
* `ToWeekString()`：週文字列を出力
* `ToShortWeekString()`：略語週文字列を出力

###### エリアの設定
```go
lang := NewLanguage()
if err := lang.SetLocale("zh-CN");err != nil {
	// エラー処理
    log.Fatal(err)
}

c := carbon.SetLanguage(lang)
c.Now().AddHours(1).DiffForHumans() // 1 小时后
c.Now().AddHours(1).ToMonthString() // 八月
c.Now().AddHours(1).ToShortMonthString() // 8月
c.Now().AddHours(1).ToWeekString() // 星期二
c.Now().AddHours(1).ToShortWeekString() // 周二
c.Now().AddHours(1).Constellation() // 狮子座
c.Now().AddHours(1).Season() // 夏季
```

###### 翻訳リソースの一部を書き換える(残りはまだ指定された `locale` ファイルの内容によって翻訳されます)
```go
lang := NewLanguage()

if err := lang.SetLocale("en");err != nil {
	// エラー処理
    log.Fatal(err)
}

resources := map[string]string {
    "hour":"%dh",
}
lang.SetResources(resources)

c := carbon.SetLanguage(lang)
c.Now().AddYears(1).DiffForHumans() // 1 year from now
c.Now().AddHours(1).DiffForHumans() // 1h from now
c.Now().ToMonthString() // August
c.Now().ToShortMonthString() // Aug
c.Now().ToWeekString() // Tuesday
c.Now().ToShortWeekString() // Tue
c.Now().Constellation() // Leo
c.Now().Season() // Summer
```

###### すべての翻訳リソースを書き換えます
```go
lang := NewLanguage()
resources := map[string]string {
    "seasons": "Spring|Summer|Autumn|Winter",
    "months": "January|February|March|April|May|June|July|August|September|October|November|December",
    "months_short": "Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec",
    "weeks": "Sunday|Monday|Tuesday|Wednesday|Thursday|Friday|Saturday",
    "weeks_short": "Sun|Mon|Tue|Wed|Thu|Fri|Sat",
    "constellations": "Aries|Taurus|Gemini|Cancer|Leo|Virgo|Libra|Scorpio|Sagittarius|Capricornus|Aquarius|Pisce",
    "year":"1 yr|%d yrs",
    "month":"1 mo|%d mos",
    "week":"%dw",
    "day":"%dd",
    "hour":"%dh",
    "minute":"%dm",
    "second":"%ds",
    "now": "just now",
    "ago":"%s ago",
    "from_now":"in %s",
    "before":"%s before",
    "after":"%s after",
}
lang.SetResources(resources)

c := carbon.SetLanguage(lang)
c.Now().AddYears(1).DiffForHumans() // in 1 yr
c.Now().AddHours(1).DiffForHumans() // in 1h
c.Now().ToMonthString() // August
c.Now().ToShortMonthString() // Aug
c.Now().ToWeekString() // Tuesday
c.Now().ToShortWeekString() // Tue
c.Now().Constellation() // Leo
c.Now().Season() // Summer
```

##### エラー処理
> 複数のエラーが発生した場合、最初のエラーだけを返します。前のエラーは削除された後、次のエラーに戻ります

###### シーン一
```go
c := carbon.SetTimezone(PRC).Parse("xxx")
if c.Error != nil {
    // エラー処理...
    log.Fatal(c.Error)
}
fmt.Println(c.ToDateTimeString())
// 出力
cannot parse "xxx" as carbon, please make sure the value is valid
```

###### シーン二
```go
c := carbon.SetTimezone("xxx").Parse("2020-08-05")
if c.Error != nil {
    // エラー処理...
    log.Fatal(c.Error)
}
fmt.Println(c.ToDateTimeString())
// 出力
invalid timezone "xxx", please see the file "$GOROOT/lib/time/zoneinfo.zip" for all valid timezones
```
###### シーン三
```go
c := carbon.SetTimezone("xxx").Parse("12345678")
if c.Error != nil {
    // エラー処理...
    log.Fatal(c.Error)
}
fmt.Println(c.ToDateTimeString())
// 出力
invalid timezone "xxx", please see the file "$GOROOT/lib/time/zoneinfo.zip" for all valid timezone
```

#### 付録
##### <a id="format-sign-table">書式設定記号表</a>
| 記号 | 説明 |  長さ | 範囲 | 例 |
| :------------: | :------------: | :------------: | :------------: | :------------: |
| d | 月の中の何日目ですか | 2 | 01-31 | 05 |
| D | 略語は何曜日を表しますか | 3 | Mon-Sun | Wed |
| j | 月の中の何日目ですか | - |1-31 | 5 |
| S | 何日目の英語の略語の接尾語，普通はjと協力して使います | 2 | st/nd/rd/th | th |
| l | 完全な単語は何曜日を表しますか | - | Monday-Sunday | Wednesday |
| F | 完全な単語は月を表しますか | - | January-December | August |
| m | 数字が示す月は | 2 | 01-12 | 08 |
| M | 略語の月 | 3 | Jan-Dec | Aug |
| n | 数字が示す月 | - | 1-12 | 8 |
| Y | 数字が示す年 | 4 | 0000-99999 | 2020 |
| y | 数字が示す年 | 4 | 0000-9999 | 20 |
| a | 小文字の午前と午後の標識 | 2 | am/pm | pm |
| A | 大文字の午前と午後の表示 | 2 | AM/PM | PM |
| g | 時間, 12時間のフォーマット | - | 1-12 | 1 |
| G | 時間, 24時間のフォーマット | - | 0-23 | 15 |
| h | 時間, 12時間のフォーマット | 2 | 00-11 | 03 |
| H | 時間, 24時間のフォーマット | 2 | 00-23 | 15 |
| i | 分 | 2 | 01-59 | 14 |
| s | 秒 | 2 | 01-59 | 15 |
| c | ISO8601 フォーマットの日付 | - | - | 2020-08-05T15:19:21+00:00 |
| r | RFC822 フォーマットの日付 | - | - | Thu, 21 Dec 2020 16:01:07 +0200 |
| O | グリニッジとの時間差の時間数 | - | - | +0200 |
| P | グリニッジと時間の差の時間数, 時間と分の間にコロンがあります | - | - | +02:00 |
| T | タイムゾーンの略語 | - | - | EST |
| W | ISO8601 フォーマットの数字は年の中の第数週を表します | - | 1-52 | 42 |
| N | ISO8601 フォーマットの数字は曜日の中の何日目を表しますか | 1 | 1-7 | 6 |
| L | うるう年かどうか, うるう年が1であれば, 0です | 1 | 0-1 | 1 |
| U | 秒タイムスタンプ | 10 | - | 1611818268 |
| u | ミリ秒 | 3 | 000-999 | 999 |
| w | 数字の表示の曜日 | 1 | 0-6 | 6 |
| t | 月の総日数 | 2 | 28-31 | 30 |
| z | 年の中の何日目 | - | 0-365 | 15 |
| e | 位置 | - | - | America/New_York |
| Q | 季節 | 1 | 1-4 | 1 |
| C | 世紀 | - | 0-99 | 21 |

#### 参考文献
* [briannesbitt/carbon](https://github.com/briannesbitt/Carbon)
* [jinzhu/now](https://github.com/jinzhu/now)
* [goframe/gtime](https://github.com/gogf/gf/tree/master/os/gtime)
* [arrow-py/arrow](https://github.com/arrow-py/arrow)
* [moment/moment](https://github.com/moment/moment)
* [iamkun/dayjs](https://github.com/iamkun/dayjs)