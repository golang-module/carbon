carbon 是一个轻量级、语义化、对IDE友好的日期时间处理库，是PHP Carbon库的Golang实现版本

github:[github.com/golang-module/carbon](https://github.com/golang-module/carbon "github.com/golang-module/carbon")

gitee:[gitee.com/go-package/carbon](https://gitee.com/go-package/carbon "gitee.com/go-package/carbon")

#### 安装
```go
go get gitee.com/go-package/carbon
```

#### 用法
> 假设当前北京时间为2020-09-05 13:20:30
###### 初始化
```go
import (
    "gitee.com/go-package/carbon"
)

// 初始化
c := carbon:New()
```

###### 设置时区（不设置默认为Local，即服务器所在时区）
```go
// 设置中国时区
carbon.Timezone(carbon.PRC)

// 设置上海时区
carbon.Timezone(carbon.Shanghai)

// 设置重庆时区
carbon.Timezone(carbon.Chongqing)

// 设置香港时区
carbon.Timezone(carbon.HongKong)

// 设置澳门时区
carbon.Timezone(carbon.Macao)

// 设置台湾时区
carbon.Timezone(carbon.Taipei)

// 设置日本时区
carbon.Timezone(carbon.Japan)

// 设置东京时区
carbon.Timezone(carbon.Tokyo)

// 设置纽约时区
carbon.Timezone(carbon.NewYork)
```

###### 获取当前时间
```go
carbon.Now().Format("Y-m-d H:i:s") // 2020-09-05 13:20:30
carbon.Now().Format("y-m-d h:i:s") // 20-09-05 01:20:30
carbon.Now().Format("Y/m/d") // 2020/09/05
carbon.Now().ToDateTimeString() // 2020-09-05 13:20:30
carbon.Now().ToDateString() // 2020-09-05
carbon.Now().ToTimeString() // 13:20:30
// 获取当前时间戳
carbon.Now().ToTimestamp() // 1599272433
```
###### 获取昨天、今天、明天时间
```go
carbon.Yesterday() // 2020-09-04 00:00:00
carbon.Today() // 2020-09-05 00:00:00
carbon.Tomorrow() // 2020-09-06 00:00:00
```

###### 创建时间
```go
// 时间戳 转成 标准时间字符串
carbon.CreateFromTimestamp(1599272433).Format("Y-m-d H:i:s") // 2020-09-05 13:20:30
// 年月日时分秒 转成 标准时间字符串
carbon.CreateFromDateTime(2020, 09, 05, 13, 20, 30).Format("Y-m-d H:i:s") // 2020-09-05 13:20:30
// 年月日 转成 标准时间字符串
carbon.CreateFromDate(2020, 09, 05).Format("Y-m-d H:i:s") // 2020-09-05 00:00:00
// 时分秒 转成 标准时间字符串(年月日默认为当前年月日)
carbon.CreateFromTime(13, 20, 30).Format("Y-m-d H:i:s") // 2020-09-05 13:20:30
```

###### 解析标准格式时间字符串
```go
carbon.Parse("2020-09-05 13:20:30").Format("YmdHis") // 20200905132030
carbon.Parse("2020-09-05 13:20:30").Format("Y-m-d") // 2020-09-05
carbon.Parse("2020-09-05").Format("Y/m/d H:i:s") // 2020/09/05 00:00:00
carbon.Parse("2020-09-05").Format("Y/m/d") // 2020/09/05
carbon.Parse("2020-09-05 13:20:30").ToDateTimeString() // 2020-09-05 13:20:30
carbon.Parse("2020-09-05 13:20:30").ToDateString() // 2020-09-05
carbon.Parse("2020-09-05 13:20:30").ToTimeString() // 13:20:30
carbon.Parse("2020-09-05 13:20:30").ToTimestamp() // 1599272433

carbon.Parse("2020/09/05 13:20:30").Format("Y-m-d H:i:s") // 2020-09-05 13:20:30
carbon.Parse("2020/09/05 13:20:30").Format("Y-m-d") // 2020-09-05
carbon.Parse("2020/09/05").Format("Y-m-d H:i:s") // 2020-09-05 00:00:00
carbon.Parse("2020/09/05").Format("Y-m-d") // 2020-09-05
carbon.Parse("2020/09/05 13:20:30").ToDateTimeString() // 2020-09-05 13:20:30
carbon.Parse("2020/09/05 13:20:30").ToDateString() // 2020-09-05
carbon.Parse("2020/09/05 13:20:30").ToTimeString() // 10:20:30
carbon.Parse("2020/09/05 13:20:30").ToTimestamp() // 1599272433

carbon.Parse("20200905132030").Format("Y-m-d H:i:s") // 2020-09-05 13:20:30
carbon.Parse("20200905132030").Format("Y-m-d") // 2020-09-05
carbon.Parse("20200905").Format("Y-m-d H:i:s") // 2020-09-05 00:00:00
carbon.Parse("20200905").Format("Y-m-d") // 2020-09-05
carbon.Parse("20200905132030").ToDateTimeString() // 2020-09-05 13:20:30
carbon.Parse("20200905132030").ToDateString() // 2020-09-05
carbon.Parse("20200905132030").ToTimeString() // 13:20:30
carbon.Parse("20200905132030").ToTimestamp() // 1599272433
```

###### 解析自定义格式时间字符串
```go
carbon.ParseByCustom("2020|09|05 13:20:30", "Y|m|d H:i:s").Format("Y-m-d H:i:s") // 2020-09-05 13:20:30
carbon.ParseByCustom("2020%09%05% 01%20%30", "Y年m月d日 h%i%s").Format("Y-m-d H:i:s") // 2020-09-05 13:20:30
carbon.ParseByCustom("2020年09月05日 13:20:30", "Y年m月d日 H:i:s").Format("Y-m-d H:i:s") // 2020-09-05 13:20:30
carbon.ParseByCustom("2020年09月05日 13时20分30秒", "Y年m月d日 H时i分s秒").ToDateTimeString() // 2020-09-05 13:20:30
carbon.ParseByCustom("2020年09月05日 13时20分30秒", "Y年m月d日 H时i分s秒").ToDateString() // 2020-09-05
carbon.ParseByCustom("2020年09月05日 13时20分30秒", "Y年m月d日 H时i分s秒").ToTimeString() // 13:20:30
carbon.ParseByCustom("2020年09月05日 13时20分30秒", "Y年m月d日 H时i分s秒").ToTimestamp() // 1599272433
```

###### 增减日期
```go
// 五年后
carbon.Now().AddYears(5).ToDateTimeString() // 2025-09-05 13:20:30
// 一年后
carbon.Now().AddYear().ToDateTimeString() // 2021-09-05 13:20:30
// 五年前
carbon.Now().SubYears(5).ToDateTimeString() // 2015-09-05 13:20:30
// 一年前
carbon.Now().SubYear().ToDateTimeString() // 2019-09-05 13:20:30

// 五月后
carbon.Now().AddMonths(5).ToDateTimeString() // 2021-02-05 13:20:30
// 一月后
carbon.Now().AddMonth().ToDateTimeString() // 2020-10-05 13:20:30
// 五月前
carbon.Now().SubMonths(5).ToDateTimeString() // 2020-04-05 13:20:30
// 一月前
carbon.Now().SubMonth().ToDateTimeString() // 2020-08-05 13:20:30

// 三天后
carbon.Now().AddDays(3).ToDateTimeString() // 2020-09-08 13:20:30
// 一天后
carbon.Now().AddDay().ToDateTimeString() // 2020-09-06 13:20:30
// 三天前
carbon.Now().SubDays(3).ToDateTimeString() // 2020-09-02 13:20:30
// 一天前
carbon.Now().SubDay().ToDateTimeString() // 2020-08-04 13:20:30

// 三小时后
carbon.Now().AddHours(3).ToDateTimeString() // 2020-09-05 16:20:30
// 一小时后
carbon.Now().AddHoury().ToDateTimeString() // 2020-09-05 14:20:30
// 三小时前
carbon.Now().SubHours(3).ToDateTimeString() // 2020-09-05 10:20:30
// 一小时前
carbon.Now().SubHour().ToDateTimeString() // 2020-09-05 12:20:30

// 三分钟后
carbon.Now().AddMinutes(3).ToDateTimeString() // 2020-09-05 13:23:30
// 一分钟后
carbon.Now().AddMinute().ToDateTimeString() // 2020-09-05 13:21:30
// 三分钟前
carbon.Now().SubMinutes(3).ToDateTimeString() // 2020-09-05 13:17:30
// 一分钟前
carbon.Now().SubMinute().ToDateTimeString() // 2020-09-05 13:19:30

// 三秒钟后
carbon.Now().AddSeconds(3).ToDateTimeString() // 2020-09-05 13:20:33
// 一秒钟后
carbon.Now().AddSecond().ToDateTimeString() // 2020-09-05 13:20:31
// 三秒钟前
carbon.Now().SubSeconds(3).ToDateTimeString() // 2020-09-05 13:20:27
// 一秒钟前
carbon.Now().SubSecond().ToDateTimeString() // 2020-09-05 13:20:29
```

###### 日期判断
```go
// 是否是闰年
carbon.Now().IsLeapYear() 
// 是否是周一
carbon.Now().IsMonday() 
// 是否是周二
carbon.Now().IsTuesday() 
// 是否是周三
carbon.Now().IsWednesday() 
// 是否是周四
carbon.Now().IsThursday() 
// 是否是周五
carbon.Now().IsFriday() 
// 是否是周六
carbon.Now().IsSaturday() 
// 是否是周日
carbon.Now().IsSunday() 
```


