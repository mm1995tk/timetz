package timetz_test

import (
	"testing"
	"time"

	"github.com/mm1995tk/timetz"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var jst = AsiaTokyo{time.FixedZone("JST", 9*60*60)}

type AsiaTokyo struct{ location *time.Location }

func (tz AsiaTokyo) StdLocation() *time.Location { return tz.location }

type EnumTZ string

const (
	UTC EnumTZ = "UTC"
	JST EnumTZ = "JST"
)

func TestTime(t *testing.T) {
	t.Parallel()

	t.Run("StartOfDate", func(t *testing.T) {
		t.Parallel()
		d := time.Date(2025, 1, 8, 7, 0, 0, 0, jst.StdLocation())

		result := timetz.NewTime(d, jst).StartOfDate().Std()
		expected := time.Date(2025, 1, 8, 0, 0, 0, 0, jst.StdLocation())

		assert.Equal(t, expected, result)
	})

	t.Run("SameDay", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			a, b         time.Time
			expected     bool
			expectedDate string
			tz           EnumTZ
		}{
			{expected: true,
				a:  time.Date(2025, 1, 8, 7, 0, 0, 0, jst.StdLocation()),
				b:  time.Date(2025, 1, 8, 23, 59, 0, 0, jst.StdLocation()),
				tz: JST,
			},

			{expected: true,
				a:  time.Date(2025, 1, 8, 15, 0, 0, 0, time.UTC),
				b:  time.Date(2025, 1, 9, 7, 0, 0, 0, jst.StdLocation()),
				tz: JST,
			},
			{expected: true,
				a:  time.Date(2025, 1, 8, 15, 0, 0, 0, time.UTC),
				b:  time.Date(2025, 1, 9, 7, 0, 0, 0, jst.StdLocation()),
				tz: UTC,
			},

			{expected: true,
				a:  time.Date(2025, 1, 8, 11, 0, 0, 0, time.UTC),
				b:  time.Date(2025, 1, 8, 7, 0, 0, 0, jst.StdLocation()),
				tz: JST,
			},
			{expected: false,
				a:  time.Date(2025, 1, 8, 11, 0, 0, 0, time.UTC),
				b:  time.Date(2025, 1, 8, 7, 0, 0, 0, jst.StdLocation()),
				tz: UTC,
			},
		}

		for _, tt := range tests {
			t.Run(t.Name(), func(t *testing.T) {
				t.Parallel()

				var result bool
				switch tt.tz {
				case JST:
					result = timetz.NewTime(tt.a, jst).SameDay(timetz.NewTime(tt.b, jst))
				case UTC:
					result = timetz.NewTime(tt.a, timetz.UTC).SameDay(timetz.NewTime(tt.b, timetz.UTC))
				default:
					require.Fail(t, "unexpected timezone")
				}

				assert.Equal(t, tt.expected, result)
			})
		}

	})

}
