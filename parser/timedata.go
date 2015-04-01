package parser

import (
    "time"

    "github.com/cbrand/jiffyreporter/helpers"
)

// TimeData represents an entry parsed form the jiffy
// csv file.
type TimeData struct {
    // Customer is the string identification for this
    // entry.
    Customer string
    // Project ist the name of the Project being worked on.
    Project string
    // Task is the task in the project which has been worked on.
    Task string
    // Start is the time which the work progress has been started on.
    Start time.Time
    // End is the time which the work progress has been stopped at.
    End time.Time
    // Duration is the amount of time the person has worked on it.
    Duration time.Duration
    // Note is an additional information for this entry.
    Note string
}

func (self *TimeData) Date() time.Time {
    return helpers.DateFromTime(self.Start)
}
