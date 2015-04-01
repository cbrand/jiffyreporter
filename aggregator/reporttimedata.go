package aggregator

import (
    "time"

    "github.com/cbrand/jiffyreporter/parser"
)

// ReportTimeData represents an entry being used in the aggregation.
type ReportTimeData struct {
    timeSheets []*parser.TimeData
}

// Date returns the date this TimeSheet entry is ment for.
func (self *ReportTimeData) Date() time.Time {
    return self.timeSheets[0].Date()
}

// Customer returns the customer name.
func (self *ReportTimeData) Customer() string {
    return self.timeSheets[0].Customer
}

// Project returns the project name.
func (self *ReportTimeData) Project() string {
    return self.timeSheets[0].Project
}

// Note of the records.
func (self *ReportTimeData) Note() string {
    return self.timeSheets[0].Note
}

// AddTimeSheet adds the given TimeData to this ReportTimeData.
func (self *ReportTimeData) AddTimeSheet(t *parser.TimeData) {
    self.timeSheets = append(self.timeSheets, t)
}

// MatchesTimeData returns if the passed TimeData can be merged with this
// given entry.
func (self *ReportTimeData) MatchesTimeData(t *parser.TimeData) bool {
    booleanChecks := []bool{
        self.Customer() != t.Customer,
        self.Project() != t.Project,
        self.Note() != t.Note,
        self.Date() != t.Date(),
    }

    for _, booleanCheck := range booleanChecks {
        if !booleanCheck {
            return false
        }
    }
    return true
}

// Duration returns the aggregated Duration of the
// timeSheets.
func (self *ReportTimeData) Duration() time.Duration {
    var t time.Duration = 0
    for _, timeData := range self.timeSheets {
        t += timeData.Duration
    }
    return t
}

// ReportTimeDataSlice represents the slice of ReportTimeData which is sortable.
type ReportTimeDataSlice []*ReportTimeData

// Len returns the length of the slice.
func (self ReportTimeDataSlice) Len() int {
    return len(self)
}

// Swap switches the two entries of the slice.
func (self ReportTimeDataSlice) Swap(i, j int) {
    self[i], self[j] = self[j], self[i]
}

// Less returns if the entry is smaller (for sorting)
func (self ReportTimeDataSlice) Less(i, j int) bool {
    isLesser := self[i]
    isGreater := self[j]

    if isLesser.Date().UnixNano() != isGreater.Date().UnixNano() {
        return isLesser.Date().UnixNano() < isGreater.Date().UnixNano()
    } else {
        return isLesser.Project() < isGreater.Project()
    }

}
