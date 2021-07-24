<!--
    🛑 DON'T REMOVE ME.
    This issue template apply to all
      - bug reports,
      - feature proposals,
      - and documentation requests

    Having all those informations will allow us to know exactly
    what you expect and answer you faster and precisely (answer
    that matches your carbon version, golang version and usage).
    
    Note: Comments between <!- - and - -> won't appear in the final
    issue (See [Preview] tab).
-->
Hello,

I encountered an issue with the following code:
```go
carbon.Parse("2020-08-05 13:14:15").AddDays(3).ToDateTimeString()
```

carbon version: **PUT HERE YOUR CARBON VERSION (exact version, not the range)**

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
