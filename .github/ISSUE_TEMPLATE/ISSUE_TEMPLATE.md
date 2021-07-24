---
name: Bug report
about: Something isn't working as expected
---
Hello,

I encountered an issue with the following code:
```go
carbon.Parse("2020-08-05 13:14:15").AddDays(3).ToDateTimeString()
```

carbon version: **PUT HERE YOUR CARBON VERSION**

golang version: **PUT HERE YOUR GOLANG VERSION**

I expected to get:

```
2020-08-08 13:14:15
```
<!--
    Always give your expectations. Each use has their owns.
    You may want daylight saving time to be taken into account,
    someone else want it to be ignored. You may want timezone
    to be used in comparisons, someone else may not, etc.
-->

But I actually get:

```
2020-08-07 13:14:15
```

Thanks!
