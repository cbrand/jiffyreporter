package helpers

import (
    "time"
)

// DateFromTime returns the date part of the given Time object.
func DateFromTime(t time.Time) time.Time {
    return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
