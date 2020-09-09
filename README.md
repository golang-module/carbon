carbon 是一个轻量级、语义化、对IDE友好的日期时间处理库，是PHP Carbon库的Golang实现版本，初衷是为了摆脱Golang反人类的2006-01-02 15:04:05格式化时间设计

github:[github.com/golang-module/carbon](https://github.com/golang-module/carbon "github.com/golang-module/carbon")

gitee:[gitee.com/go-package/carbon](https://gitee.com/go-package/carbon "gitee.com/go-package/carbon")

#### 安装
```go
go get -u gitee.com/go-package/carbon
```

#### 用法
###### 初始化
```go
import (
    "gitee.com/go-package/carbon"
)

// 初始化
c := carbon.New()
```

###### 设置时区（不设置默认为Local，即服务器所在时区）
```go
// 设置中国时区
c = c.Timezone(carbon.PRC)

// 设置上海时区
c = c.Timezone(carbon.Shanghai)

// 设置重庆时区
c = c.Timezone(carbon.Chongqing)

// 设置香港时区
c = c.Timezone(carbon.HongKong)

// 设置澳门时区
c = c.Timezone(carbon.Macao)

// 设置台湾时区
c = c.Timezone(carbon.Taipei)

// 设置日本时区
c = c.Timezone(carbon.Japan)

// 设置东京时区
c = c.Timezone(carbon.Tokyo)

// 设置纽约时区
c = c.Timezone(carbon.NewYork)

// 设置伦敦时区
c = c.Timezone(carbon.London)

```

###### 获取当前时间
```go
c.Now().Format("Y-m-d H:i:s") // 2020-09-08 13:00:00
c.Now().Format("y-m-d h:i:s") // 20-09-08 01:00:00
c.Now().Format("Y/m/d") // 2020/09/08
c.Now().ToDateTimeString() // 2020-09-08 13:00:00
c.Now().ToDateString() // 2020-09-08
c.Now().ToTimeString() // 13:00:00
// 获取当前时间戳
c.Now().ToTimestamp() // 1599272433
```
###### 获取昨天、今天、明天时间
```go
c.Yesterday() // 2020-09-07 00:00:00
c.Today() // 2020-09-08 00:00:00
c.Tomorrow() // 2020-09-09 00:00:00
```

###### 获取第一天、最后一天
```go
c.Now().FirstDay() // 2020-09-01 00:00:00
c.Now().LastDay() // 2020-09-30 00:00:00
```

###### 数字转标准时间字符串
```go
// 时间戳 转成 标准时间字符串
c.CreateFromTimestamp(1599272433).Format("Y-m-d H:i:s") // 2020-09-08 13:00:00
// 年月日时分秒 转成 标准时间字符串
c.CreateFromDateTime(2020, 09, 08, 13, 00, 00).Format("Y-m-d H:i:s") // 2020-09-08 13:00:00
// 年月日 转成 标准时间字符串
c.CreateFromDate(2020, 09, 08).Format("Y-m-d H:i:s") // 2020-09-08 00:00:00
// 时分秒 转成 标准时间字符串(年月日默认为当前年月日)
c.CreateFromTime(13, 14, 15).Format("Y-m-d H:i:s") // 2020-09-08 13:14:15
```

###### 解析标准时间字符串
```go
c.Parse("2020-09-08 13:00:00").Format("YmdHis") // 20200908130000
c.Parse("2020-09-08 13:00:00").Format("Y-m-d") // 2020-09-08
c.Parse("2020-09-08").Format("Y/m/d H:i:s") // 2020/09/08 00:00:00
c.Parse("2020-09-08").Format("Y/m/d") // 2020/09/08
c.Parse("2020-09-08 13:00:00").ToDateTimeString() // 2020-09-05 13:00:00
c.Parse("2020-09-08 13:00:00").ToDateString() // 2020-09-08
c.Parse("2020-09-08 13:00:00").ToTimeString() // 13:00:00
c.Parse("2020-09-08 13:00:00").ToTimestamp() // 1599272433

c.Parse("2020/09/08 13:00:00").Format("Y-m-d H:i:s") // 2020-09-08 13:00:00
c.Parse("2020/09/08 13:00:00").Format("Y-m-d") // 2020-09-08
c.Parse("2020/09/08").Format("Y-m-d H:i:s") // 2020-09-08 00:00:00
c.Parse("2020/09/08").Format("Y-m-d") // 2020-09-08
c.Parse("2020/09/08 13:00:00").ToDateTimeString() // 2020-09-05 13:00:00
c.Parse("2020/09/08 13:00:00").ToDateString() // 2020-09-08
c.Parse("2020/09/08 13:00:00").ToTimeString() // 10:20:30
c.Parse("2020/09/08 13:00:00").ToTimestamp() // 1599272433

c.Parse("20200908130000").Format("Y-m-d H:i:s") // 2020-09-08 13:00:00
c.Parse("20200908130000").Format("Y-m-d") // 2020-09-08
c.Parse("20200908").Format("Y-m-d H:i:s") // 2020-09-08 00:00:00
c.Parse("20200908").Format("Y/m/d") // 2020/09/08
c.Parse("20200908130000").ToDateTimeString() // 2020-09-05 13:00:00
c.Parse("20200908130000").ToDateString() // 2020-09-08
c.Parse("20200908130000").ToTimeString() // 13:00:00
c.Parse("20200908130000").ToTimestamp() // 1599272433
```

###### 解析自定义格式时间字符串
```go
c.ParseByCustom("2020|09|08 13:00:00", "Y|m|d H:i:s").Format("Y-m-d H:i:s") // 2020-09-08 13:00:00
c.ParseByCustom("2020%09%08% 01%00%00", "Y年m月d日 h%i%s").Format("Y-m-d H:i:s") // 2020-09-08 13:00:00
c.ParseByCustom("2020年09月08日 13:00:00", "Y年m月d日 H:i:s").Format("Y-m-d H:i:s") // 2020-09-08 13:00:00
c.ParseByCustom("2020年09月08日 13时00分00秒", "Y年m月d日 H时i分s秒").ToDateTimeString() // 2020-09-08 13:20:30
c.ParseByCustom("2020年09月08日 13时00分00秒", "Y年m月d日 H时i分s秒").ToDateString() // 2020-09-08
c.ParseByCustom("2020年09月08日 13时00分00秒", "Y年m月d日 H时i分s秒").ToTimeString() // 13:00:00
c.ParseByCustom("2020年09月08日 13时00分00秒", "Y年m月d日 H时i分s秒").ToTimestamp() // 1599272433
```

###### 时间旅行
> 假设当前北京时间为2020-09-08 13:00:00
```go
// 三年后
c.Now().AddYears(3).ToDateTimeString() // 2023-09-08 13:00:00
// 一年后
c.Now().AddYear().ToDateTimeString() // 2021-09-08 13:00:00
// 三年前
c.Now().SubYears(3).ToDateTimeString() // 2017-09-08 13:00:00
// 一年前
c.Now().SubYear().ToDateTimeString() // 2019-09-08 13:00:00

// 三月后
c.Now().AddMonths(3).ToDateTimeString() // 2020-12-08 13:00:00
// 一月后
c.Now().AddMonth().ToDateTimeString() // 2020-10-08 13:00:00
// 三月前
c.Now().SubMonths(3).ToDateTimeString() // 2020-06-08 13:00:00
// 一月前
c.Now().SubMonth().ToDateTimeString() // 2020-08-08 13:00:00

// 三天后
c.Now().AddDays(3).ToDateTimeString() // 2020-09-11 13:00:00
// 一天后
c.Now().AddDay().ToDateTimeString() // 2020-09-09 13:00:00
// 三天前
c.Now().SubDays(3).ToDateTimeString() // 2020-09-05 13:00:00
// 一天前
c.Now().SubDay().ToDateTimeString() // 2020-08-07 13:00:00

// 三小时后
c.Now().AddHours(3).ToDateTimeString() // 2020-09-08 16:00:00
// 一小时后
c.Now().AddHoury().ToDateTimeString() // 2020-09-08 14:00:00
// 三小时前
c.Now().SubHours(3).ToDateTimeString() // 2020-09-08 10:00:00
// 一小时前
c.Now().SubHour().ToDateTimeString() // 2020-09-08 12:00:00

// 三分钟后
c.Now().AddMinutes(3).ToDateTimeString() // 2020-09-08 13:03:00
// 一分钟后
c.Now().AddMinute().ToDateTimeString() // 2020-09-08 13:01:00
// 三分钟前
c.Now().SubMinutes(3).ToDateTimeString() // 2020-09-08 12:57:00
// 一分钟前
c.Now().SubMinute().ToDateTimeString() // 2020-09-08 12:59:00

// 三秒钟后
c.Now().AddSeconds(3).ToDateTimeString() // 2020-09-08 13:00:03
// 一秒钟后
c.Now().AddSecond().ToDateTimeString() // 2020-09-08 13:00:01
// 三秒钟前
c.Now().SubSeconds(3).ToDateTimeString() // 2020-09-08 12:59:57
// 一秒钟前
c.Now().SubSecond().ToDateTimeString() // 2020-09-08 12:59:59
```

###### 日期判断
```go
// 是否是闰年
c.Now().IsLeapYear() // true

// 是否是周一
c.Now().IsMonday() // false
// 是否是周二
c.Now().IsTuesday() // true
// 是否是周三
c.Now().IsWednesday() // false
// 是否是周四
c.Now().IsThursday()  // false
// 是否是周五
c.Now().IsFriday() // false
// 是否是周六
c.Now().IsSaturday() // false
// 是否是周日
c.Now().IsSunday() // false

// 是否月初
c.Now().IsFirstDay() // false
//是否是月末
c.Now().IsLastDay() // false
```

 **更新日志** 
 
##### 2020-09-09
* 修复readme.md初始化的错误描述
* 添加单元测试
* 新增gorm时间格式化，支持多种输出格式
* 新增IsJanuary()方法判断是否第一月
* 新增IsFebruary()方法判断是否第二月
* 新增IsMarch()方法判断是否第三月
* 新增IsApril()方法判断是否第四月
* 新增IsMay()方法判断是否第五月
* 新增IsJune()方法判断是否第六月
* 新增IsJuly()方法判断是否第七月
* 新增IsAugust()方法判断是否第八月
* 新增IsSeptember()方法判断是否第九月
* 新增IsOctober()方法判断是否第十月
* 新增IsNovember()方法判断是否第十一月
* 新增IsDecember()方法判断是否第十二月
 
##### 2020-09-08
* 修复已知BUG
* 添加单元测试
* 新增FirstDay()方法获取第一天
* 新增FirstDay()方法获取最后一天
* 新增IsFirstDay()方法判断是否第一天
* 新增IsFirstDay()方法判断是否最后一天 
