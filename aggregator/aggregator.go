package aggregator

import (
    "sort"

    "github.com/cbrand/jiffyreporter/parser"
)

// New returns a new Aggregator struct which can filter the
// timeData entries for given criteria.
func New(data chan *parser.TimeData) *Aggregator {
    return &Aggregator{
        data: data,
    }
}

// NewFromArray initializes the Aggregator struct with the given passed
// TimeData array.
func NewFromArray(data []*parser.TimeData) *Aggregator {
    dataChannel := make(chan *parser.TimeData)
    go func() {
        for _, item := range data {
            dataChannel <- item
        }
    }()
    return New(dataChannel)
}

// Aggregate function combines the data from the given channel.
func Aggregate(data chan *parser.TimeData) []*ReportTimeData {
    return New(data).Aggregate()
}

// AggregateArray combines the data from the given arrray.
func AggregateArray(data []*parser.TimeData) []*ReportTimeData {
    return NewFromArray(data).Aggregate()
}

// Aggregator struct is being used to convert the TimeData to ReportTimeData
// structs.
type Aggregator struct {
    data chan *parser.TimeData
}

// Aggregate aggregates the ReportTimeData entries from the aggregated
// data channels.
func (self *Aggregator) Aggregate() []*ReportTimeData {
    retData := ReportTimeDataSlice{}

    for timeData := range self.data {
        found := false
        var matchingTimeSheet *ReportTimeData
        for _, reportTimeData := range retData {
            if reportTimeData.MatchesTimeData(timeData) {
                matchingTimeSheet = reportTimeData
                found = true
            }
        }

        if !found {
            matchingTimeSheet := ReportTimeData{
                data: []*parser.TimeData{},
            }
            retData = append(retData, matchingTimeSheet)
        }
        matchingTimeSheet.AddTimeSheet(timeData)
    }
    sort.Sort(retData)
    return retData
}
