# timetz

timetz is a thin wrapper of time. timetz can recieve timezone as a type parameter so that you can get typesafe programming experience.

## how to use

### import

```go
import "github.com/mm1995tk/timetz"
```

### define timezones you use in your project.

```go
var JST = AsiaTokyo{time.FixedZone("JST", 9*60*60)}

type AsiaTokyo struct{ location *time.Location }

// impl timetz.TimeZone
func (tz AsiaTokyo) StdLocation() *time.Location { return tz.location }
```

### create instance from time pkg.

```go
now := time.Now()

nowTZ := timetz.NewTime(now, JST) // timetz.Time[AsiaTokyo]
nowUTC := nowTZ.UTC() // timetz.Time[timetz.EtcUTC]
convertedToTime := nowUTC.Std() // time.Time

d := nowTZ.Date() // timetz.Date[AsiaTokyo]

convertedToTime = d.
  StartOfDate(). // timetz.Time[AsiaTokyo]
  Std() // time.Time
```
