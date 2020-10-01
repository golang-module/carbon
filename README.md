carbon 是一个轻量级、语义化、对IDE友好的日期时间处理库，是PHP Carbon库的Golang实现版本，初衷是为了摆脱Golang反人类的2006-01-02 15:04:05格式化时间设计，支持链式调用和gorm、xrom等主流orm

如果您觉得不错，请给个star吧

github:[github.com/golang-module/carbon](https://github.com/golang-module/carbon "github.com/golang-module/carbon")

gitee:[gitee.com/go-package/carbon](https://gitee.com/go-package/carbon "gitee.com/go-package/carbon")

#### 安装
```go
go get -u gitee.com/go-package/carbon
```
#### 初始化
```go
import (
    "gitee.com/go-package/carbon"
)

// 初始化，默认时区为Local，即服务器所在时区
c := carbon.New()
// 初始化并设置时区
c := carbon.New().Timezone(carbon.PRC)
```
>更多时区常量请查看[const.go](https://gitee.com/go-package/carbon/blob/master/const.go "const.go")文件

#### 通用用法

> 假设当前北京时间为2020-09-08 13:00:01

###### 获取当前时间
```go
c.Now().Format("Y-m-d H:i:s") // 2020-09-08 13:00:01
c.Now().Format("y-m-d h:i:s") // 20-09-08 01:00:01
c.Now().Format("Y/m/d") // 2020/09/08
c.Now().ToDateString() // 2020-09-08
c.Now().ToDateTimeString() // 2020-09-08 13:00:01
c.Now().ToDateStartString() // 2020-09-08 00:00:00
c.Now().ToDateEndString() // 2020-09-08 23:59:59
c.Now().ToTimeString() // 13:00:01
c.Now().ToTimeStartString() // 13:00:00
c.Now().ToTimeEndString() // 13:59:59
```

###### 获取当前时间戳
```go
c.Now().ToTimestamp() // 1599272433
```

###### 获取昨天、今天、明天
```go
c.Yesterday() // 2020-09-07
c.Today() // 2020-09-08
c.Tomorrow() // 2020-09-09

c.StartOfYesterday() // 2020-09-07 00:00:00
c.EndOfYesterday() // 2020-09-07 23:59:59
c.StartOfToday() // 2020-09-08 00:00:00
c.EndOfToday() // 2020-09-08 23:59:59
c.StartOfTomorrow() // 2020-09-09 00:00:00
c.EndOfTomorrow() // 2020-09-09 23:59:59
```

###### 获取开始时间、结束时间

```go
// 当前年的开始
c.Now().StartOfYear() // 2020-01-01 00:00:00
// 当前年的结束
c.Now().EndOfYear() // 2020-12-31 23:59:59
// 当前月的开始
c.Now().StartOfMonth() // 2020-09-01 00:00:00
// 当前月的结束
c.Now().EndOfMonth() // 2020-09-30 23:59:59
// 当前天的开始
c.Now().StartOfDay() // 2020-09-08 00:00:00
// 当前天的结束
c.Now().EndOfDay() // 2020-09-08 23:59:59
```

###### 获取第一天、最后一天

```go
// 当前年的第一天
c.Now().FirstOfYear() // 2020-01-01 00:00:00
// 当前年的最后一天
c.Now().LastOfYear() // 2020-12-31 00:00:00
// 当前月的第一天
c.Now().FirstOfMonth() // 2020-09-01 00:00:00
// 当前月的最后一天
c.Now().LastOfMonth() // 2020-09-30 00:00:00
```

###### 数字转标准时间字符串
```go
// 时间戳 转成 标准时间字符串
c.CreateFromTimestamp(1599272433).Format("Y-m-d H:i:s") // 2020-09-08 13:00:01
// 年月日时分秒 转成 标准时间字符串
c.CreateFromDateTime(2020, 09, 08, 13, 00, 01).Format("Y-m-d H:i:s") // 2020-09-08 13:00:01
// 年月日 转成 标准时间字符串
c.CreateFromDate(2020, 09, 08).Format("Y-m-d H:i:s") // 2020-09-08 00:00:00
// 时分秒 转成 标准时间字符串(年月日默认为当前年月日)
c.CreateFromTime(13, 14, 15).Format("Y-m-d H:i:s") // 2020-09-08 13:14:15
```

###### 解析标准格式时间字符串
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
c.ParseByFormat("2020|09|08 13:00:00", "Y|m|d H:i:s").Format("Y-m-d H:i:s") // 2020-09-08 13:00:00
c.ParseByFormat("2020%09%08% 01%00%00", "Y年m月d日 h%i%s").Format("Y-m-d H:i:s") // 2020-09-08 13:00:00
c.ParseByFormat("2020年09月08日 13:00:00", "Y年m月d日 H:i:s").Format("Y-m-d H:i:s") // 2020-09-08 13:00:00
c.ParseByFormat("2020年09月08日 13时00分00秒", "Y年m月d日 H时i分s秒").ToDateTimeString() // 2020-09-08 13:20:30
c.ParseByFormat("2020年09月08日 13时00分00秒", "Y年m月d日 H时i分s秒").ToDateString() // 2020-09-08
c.ParseByFormat("2020年09月08日 13时00分00秒", "Y年m月d日 H时i分s秒").ToTimeString() // 13:00:00
c.ParseByFormat("2020年09月08日 13时00分00秒", "Y年m月d日 H时i分s秒").ToTimestamp() // 1599272433
```

###### 解析相对时间字符串(相对于今天)
```go
// 十小时后
c.ParseByDuration("10h").ToDateTimeString() // 2020-09-09 01:00:01
// 十小时前
c.ParseByDuration("-10h").Format("Y-m-d H:i:s") // 2020-09-08 03:00:01
// 十分钟后
c.ParseByDuration("10m").Format("Y-m-d H:i:s") // 2020-09-08 13:10:01
// 十分钟前
c.ParseByDuration("-10m").Format("Y-m-d H:i:s") // 2020-09-08 12:50:01
// 十秒后
c.ParseByDuration("10s").Format("Y-m-d H:i:s") // 2020-09-08 13:00:11
// 十秒前
c.ParseByDuration("-10s").Format("Y-m-d H:i:s") // 2020-09-08 12:59:51
```

###### 时间旅行

```go
// 三年后
c.Now().AddYears(3).ToDateTimeString() // 2023-09-08 13:00:01
// 一年后
c.Now().AddYear().ToDateTimeString() // 2021-09-08 13:00:01
// 三年前
c.Now().SubYears(3).ToDateTimeString() // 2017-09-08 13:00:01
// 一年前
c.Now().SubYear().ToDateTimeString() // 2019-09-08 13:00:01

// 三月后
c.Now().AddMonths(3).ToDateTimeString() // 2020-12-08 13:00:01
// 一月后
c.Now().AddMonth().ToDateTimeString() // 2020-10-08 13:00:01
// 三月前
c.Now().SubMonths(3).ToDateTimeString() // 2020-06-08 13:00:01
// 一月前
c.Now().SubMonth().ToDateTimeString() // 2020-08-08 13:00:01

// 三周后
c.Now().AddWeeks(3).ToDateTimeString() // 2020-09-29 13:00:01
// 一周后
c.Now().AddWeek().ToDateTimeString() // 2020-09-15 13:00:01
// 三周前
c.Now().SubWeeks(3).ToDateTimeString() // 2020-08-17 13:00:01
// 一周前
c.Now().SubWeek().ToDateTimeString() // 2020-09-01 13:00:01

// 三天后
c.Now().AddDays(3).ToDateTimeString() // 2020-09-11 13:00:01
// 一天后
c.Now().AddDay().ToDateTimeString() // 2020-09-09 13:00:01
// 三天前
c.Now().SubDays(3).ToDateTimeString() // 2020-09-05 13:00:01
// 一天前
c.Now().SubDay().ToDateTimeString() // 2020-08-07 13:00:01

// 三小时后
c.Now().AddHours(3).ToDateTimeString() // 2020-09-08 16:00:01
// 一小时后
c.Now().AddHoury().ToDateTimeString() // 2020-09-08 14:00:01
// 三小时前
c.Now().SubHours(3).ToDateTimeString() // 2020-09-08 10:00:01
// 一小时前
c.Now().SubHour().ToDateTimeString() // 2020-09-08 12:00:01

// 三分钟后
c.Now().AddMinutes(3).ToDateTimeString() // 2020-09-08 13:03:01
// 一分钟后
c.Now().AddMinute().ToDateTimeString() // 2020-09-08 13:01:01
// 三分钟前
c.Now().SubMinutes(3).ToDateTimeString() // 2020-09-08 12:57:01
// 一分钟前
c.Now().SubMinute().ToDateTimeString() // 2020-09-08 12:59:01

// 三秒钟后
c.Now().AddSeconds(3).ToDateTimeString() // 2020-09-08 13:00:04
// 一秒钟后
c.Now().AddSecond().ToDateTimeString() // 2020-09-08 13:00:02
// 三秒钟前
c.Now().SubSeconds(3).ToDateTimeString() // 2020-09-08 12:59:56
// 一秒钟前
c.Now().SubSecond().ToDateTimeString() // 2020-09-08 13:00:00

```

###### 时间判断
```go
// 是否是闰年
c.Now().IsLeapYear() // true

// 是否是一月
c.Now().IsJanuary() // false
// 是否是二月
c.Now().IsFebruary() // false
// 是否是三月
c.Now().IsMarch() // false
// 是否是四月
c.Now().IsApril()  // false
// 是否是五月
c.Now().IsMay() // false
// 是否是六月
c.Now().IsJune() // false
// 是否是七月
c.Now().IsJuly() // false
// 是否是八月
c.Now().IsAugust() // false
// 是否是九月
c.Now().IsSeptember() // true
// 是否是十月
c.Now().IsOctober() // false
// 是否是十一月
c.Now().IsNovember() // false
// 是否是十二月
c.Now().IsDecember() // false

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

// 是否是今天
c.Now().IsToday() // true
// 是否是昨天
c.Now().IsYesterday() // false
// 是否是明天
c.Now().IsTomorrow() // false
// 是否是今天开始时间
c.Now().IsStartOfToday() // false
// 是否是今天结束时间
c.Now().IsEndOfToday // false
// 是否是明天开始时间
c.Now().IsStartOfTomorrow() // false
// 是否是明天结束时间
c.Now().IsEndOfTomorrow() // false
// 是否是昨天开始时间
c.Now().IsStartOfYesterday() // false
// 是否是昨天结束时间
c.Now().IsEndOfYesterday() // false

// 是否是年初
c.Now().IsFirstOfYear() // false
//是否是年末
c.Now().IsLastOfYear() // false
// 是否是月初
c.Now().IsFirstOfMonth() // false
//是否是月末
c.Now().IsLastOfMonth() // false
```

###### 农历支持
```go
// 获取生肖年
c.Now().ToAnimalYear() // 鼠
// 获取农历年
c.Now().ToLunarYear() // 庚子

// 是否是鼠年
c.Now().IsYearOfRat() // true
// 是否是牛年
c.Now().IsYearOfOx() // false
// 是否是虎年
c.Now().IsYearOfTiger() // false
// 是否是兔年
c.Now().IsYearOfRabbit() // false
// 是否是龙年
c.Now().IsYearOfDragon() // false
// 是否是蛇年
c.Now().IsYearOfSnake() // false
// 是否是马年
c.Now().IsYearOfHorse() // false
// 是否是羊年
c.Now().IsYearOfGoat() // false
// 是否是猴年
c.Now().IsYearOfMonkey() // false
// 是否是鸡年
c.Now().IsYearOfRooster() // false
// 是否是狗年
c.Now().IsYearOfDog() // false
// 是否是猪年
c.Now().IsYearOfPig() // false

```

#### 特殊用法
假设数据表为users，字段有id(int)、name(varchar)、age(int)、birthday(date)、created_at(datetime)、updated_at(datetime)、deleted_at(datetime)

##### 在gorm中的应用
gorm.Open时必须包括parseTime=True参数

```go
// 用法一，使用carbon.GormModel自动维护id、created_at、updated_at、deleted_at
type User struct {
	carbon.GormModel
	Name string `json:"name"`
	Age int `json:"age"`
	Birthday carbon.ToDateTimeString `json:"birthday"`
}
user := User {
    Name: "勾国印"
    Age: 18
    Birthday: "2012-09-09 00:00:00"
}
// json.Marshal(user)输出
{
    "id": 1, 
    "name": "勾国印", 
    "age": 18, 
    "birthday": "2012-09-09 00:00:00", 
    "created_at": "2020-09-09 12:13:14", 
    "updated_at": "2020-09-09 12:13:14", 
    "deleted_at": null
}

// 用法二，不使用carbon.GormModel
type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Birthday carbon.ToDateString `json:"birthday"`
	CreatedAt carbon.ToDateTimeString `json:"created_at"`
	UpdatedAt carbon.ToTimeString `json:"updated_at"`
	DeletedAt carbon.ToTimestamp `json:"deleted_at"`
}
user := User {
    Name: "勾国印"
    Age: 18
    Birthday: "2012-09-09 00:00:00"
}
// json.Marshal(user)输出
{
    "id": 1, 
    "name": "勾国印", 
    "age": 18, 
    "birthday": "2012-09-09", 
    "created_at": "2020-09-09 12:13:14", 
    "updated_at": "12:13:14", 
    "deleted_at": 1599272433
}
```

##### 在xorm中的应用
xorm.NewEngine时必须包括parseTime=True参数

```go
// 用法一，使用carbon.XormModel自动维护id、created_at、updated_at、deleted_at
type User struct {
	carbon.XormModel
	Name string `json:"name"`
	Age int `json:"age"`
	Birthday carbon.ToDateTimeString `json:"birthday"`
}
user := User {
    Name: "勾国印"
    Age: 18
    Birthday: "2012-09-09 00:00:00"
}
// json.Marshal(user)输出
{
    "id": 1, 
    "name": "勾国印", 
    "age": 18, 
    "birthday": "2012-09-09 00:00:00", 
    "created_at": "2020-09-09 12:13:14", 
    "updated_at": "2020-09-09 12:13:14", 
    "deleted_at": null
}

// 用法二，不使用carbon.XormModel
type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Birthday carbon.ToDateString `json:"birthday"`
	CreatedAt carbon.ToDateTimeString `json:"created_at"`
	UpdatedAt carbon.ToTimeString `json:"updated_at"`
	DeletedAt carbon.ToTimestamp `json:"deleted_at"`
}
user := User {
    Name: "勾国印"
    Age: 18
    Birthday: "2012-09-09 00:00:00"
}
// json.Marshal(user)输出
{
    "id": 1, 
    "name": "勾国印", 
    "age": 18, 
    "birthday": "2012-09-09", 
    "created_at": "2020-09-09 12:13:14", 
    "updated_at": "12:13:14", 
    "deleted_at": 1599272433
}
```

#### 更新日志
##### 2020-10-01
* 完善单元测试
* 修复 AddHours() 传入参数小于1天时变成浮点数的错误
* 修复 AddHour() 浮点数错误
* 修复 SubHours() 传入参数小于1天时变成浮点数的错误
* 修复 SubHour() 浮点数错误
* 修复 AddMinutes() 传入参数小于1天时变成浮点数的错误
* 修复 AddMinute() 浮点数错误
* 修复 SubMinutes() 传入参数小于1天时变成浮点数的错误
* 修复 SubMinute() 浮点数错误
* 修复 AddSeconds() 传入参数小于1天时变成浮点数的错误
* 修复 AddSecond() 浮点数错误
* 修复 SubSeconds() 传入参数小于1天时变成浮点数的错误
* 修复 SubSecond() 浮点数错误
* 修复orm中时间字段类型设置为carbon.ToDateTimeString时报错的BUG
* 改名解析自定义时间格式方法ParseByCustom() 为 ParseByFormat()
* 新增 ParseByDuration() 方法将持续时间字符串转化成时间实例
* 新增对农历的部分支持

##### 2020-09-14
* 完善单元测试
* 时区常量移到const.go文件里
* 新增StartOfYear()方法获取当年开始时间
* 新增EndOfYear()方法获取当年结束时间
* 新增StartOfMonth()方法获取当月开始时间
* 新增EndOfMonth()方法获取当月结束时间
* 新增StartOfDay()方法获取当天开始时间
* 新增EndOfDay()方法获取当天结束时间
* 新增StartOfYesterday()方法获取昨天开始时间
* 新增EndOfYesterday()方法获取昨天结束时间
* 新增StartOfToday()方法获取今天开始时间
* 新增EndOfToday()方法获取今天结束时间
* 新增StartOfTomorrow()方法获取明天开始时间
* 新增EndOfTomorrow方法获取明天结束时间
* 新增ToDateStartString方法转换成日期开始时间
* 新增ToDateEndString方法转换成日期结束时间
* 新增ToTimeStartString方法转换成小时开始时间
* 新增ToTimeEndString方法转换成小时结束时间

##### 2020-09-13
* 私有方法提取到独立文件private.go里
* 新增IsToday方法判断是否是今天
* 新增IsYesterday方法判断是否是昨天
* 新增IsTomorrow方法判断是否是明天
* 新增IsStartOfToday方法判断是否是今天开始时间
* 新增IsEndOfToday方法判断是否是今天结束时间
* 新增IsStartOfTomorrow方法判断是否是明天开始时间
* 新增IsEndOfTomorrow方法判断是否是明天结束时间
* 新增IsStartOfYesterday方法判断是否是昨天开始时间
* 新增IsEndOfYesterday方法判断是否是昨天结束时间

##### 2020-09-12
* 修复数据库中时间类型字段值为null或0000-00-00 00:00:00时，json格式化后为0001-01-01 00:00:00的BUG
* 完善单元测试
* 优化代码组织结构，精简代码
* 新增对xorm结构体的json输出时间格式化支持，支持输出多种标准时间格式
* 新增AddWeeks(N)方法获取N周后时间
* 新增AddWeek()方法获取1周后时间
* 新增SubWeeks(N)方法获取N周前时间
* 新增SubWeek()方法获取1周前时间

##### 2020-09-09
* 修复readme.md错误描述
* 完善单元测试
* 新增对gorm结构体的json输出时间格式化支持，支持输出多种标准时间格式
* 新增IsJanuary()方法判断是否是一月
* 新增IsFebruary()方法判断是否是二月
* 新增IsMarch()方法判断是否是三月
* 新增IsApril()方法判断是否是四月
* 新增IsMay()方法判断是否是五月
* 新增IsJune()方法判断是否是六月
* 新增IsJuly()方法判断是否是七月
* 新增IsAugust()方法判断是否是八月
* 新增IsSeptember()方法判断是否是九月
* 新增IsOctober()方法判断是否是十月
* 新增IsNovember()方法判断是否是十一月
* 新增IsDecember()方法判断是否是十二月
 
##### 2020-09-08
* 修复已知BUG
* 添加单元测试
* 新增FirstOfYear()方法获取年初第一天
* 新增LastOfYear()方法获取年末最后一天
* 新增FirstOfMonth()方法获取月初第一天
* 新增LastOfMonth()方法获取月末最后一天
* 新增IsFirstOfYaer()方法判断是否年初第一天
* 新增IsLastOfYear()方法判断是否年末最后一天
* 新增IsFirstOfMonth()方法判断是否月初第一天
* 新增IsLastOfMonth()方法判断是否月末最后一天
