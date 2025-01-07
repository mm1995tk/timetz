package timetz

import "time"

// Time represents a time with a timezone.
type Time[TZ TimeZone] struct {
	t  time.Time
	tz TZ
}

type TimeZone interface{ StdLocation() *time.Location }

// Std returns the `time.Time`.
func (t Time[TZ]) Std() time.Time { return t.t }

// Add is a wrapper of `time.Time.Add`.
func (t Time[TZ]) Add(d time.Duration) Time[TZ] { return NewTime(t.t.Add(d), t.tz) }

// AddDate is a wrapper of `time.Time.AddDate`.
func (t Time[TZ]) AddDate(years int, months int, days int) Time[TZ] {
	return NewTime(t.t.AddDate(years, months, days), t.tz)
}

// SameDay returns `true` if `t` and `u` are the same day in TZ.
func (t Time[TZ]) SameDay(u Time[TZ]) bool { return t.t.YearDay() == u.t.YearDay() }

// StartOfDate returns the start of the day in TZ.
//
// For example, if `t` is 2021-01-02T03:04:05Z, then `t.StartOfDate()` is 2021-01-02T00:00:00Z.
func (t Time[TZ]) StartOfDate() Time[TZ] {
	return NewTime(time.Date(t.t.Year(), t.t.Month(), t.t.Day(), 0, 0, 0, 0, t.tz.StdLocation()), t.tz)
}

// Date returns a `Date` instance.
func (t Time[TZ]) Date() Date[TZ] {
	return Date[TZ]{
		t:  t.StartOfDate(),
		tz: t.tz,
	}
}

// UTC represents Universal Coordinated Time (UTC).
var UTC = EtcUTC{time.UTC}

// EtcUTC represents the UTC timezone.
type EtcUTC struct{ location *time.Location }

func (tz EtcUTC) StdLocation() *time.Location { return tz.location }

// UTC returns t with the location set to UTC.
func (t Time[TZ]) UTC() Time[EtcUTC] { return NewTime(t.t, UTC) }

// NewTime returns a new `Time` instance.
func NewTime[TZ TimeZone](t time.Time, tz TZ) Time[TZ] {
	return Time[TZ]{
		t:  t.In(tz.StdLocation()),
		tz: tz,
	}
}

// Date represents a date with a timezone.
type Date[TZ TimeZone] struct {
	t  Time[TZ]
	tz TZ
}

// StartOfDate returns the start of the day in TZ.
func (d Date[TZ]) StartOfDate() Time[TZ] { return d.t }

// Add is a wrapper of `time.Time.Add`.
func (d Date[TZ]) Add(years int, months int, days int) Date[TZ] {
	return d.t.AddDate(years, months, days).Date()
}

// Before is a wrapper of `time.Time.Before`.
func (d Date[TZ]) Before(u Date[TZ]) bool { return d.t.t.Before(u.t.t) }

// After is a wrapper of `time.Time.After`.
func (d Date[TZ]) After(u Date[TZ]) bool { return d.t.t.After(u.t.t) }

// Equal is a wrapper of `xtime.Time.SameDay`.
func (d Date[TZ]) Equal(u Date[TZ]) bool { return d.t.SameDay(u.t) }

func (d Date[TZ]) SameDay(u time.Time) bool { return d.t.SameDay(NewTime(u, d.tz)) }
