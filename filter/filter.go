package filter

import (
    "time"

    "github.com/cbrand/jiffyreporter/parser"
)

// New returns a new Filter struct which can filter the
// timeData entries for given criteria.
func New(data chan *parser.TimeData) *Filter {
    return &Filter{
        data: data,
        filters: []filterFunc{},
    }
}

// NewFromArray initializes the Filter struct with the given passed
// TimeData array.
func NewFromArray(data []*parser.TimeData) *Filter {
    dataChannel := make(chan *parser.TimeData)
    go func() {
        for _, item := range data {
            dataChannel <- item
        }
    }()
    return New(dataChannel)
}

// Filter provides functions to filter the data for given
// criteria.
type Filter struct {
    data chan *parser.TimeData
    filters []filterFunc
}

// filterFunc is the structure of the function to return the channel.
type filterFunc func(data chan *parser.TimeData) chan *parser.TimeData

// addFilter adds the filter to the entry.
func (self *Filter) addFilter(filter filterFunc) {
    self.filters = append(self.filters, filter)
}

// ForCustomers filters for all TimeData entries for the passed
// customers.
func (self *Filter) ForCustomers(customers ...string) {
    filter := func(data chan *parser.TimeData) chan *parser.TimeData {
        res := make(chan *parser.TimeData)
        go func() {
            for timeData := range data {
                for _, customer := range customers {
                    if timeData.Customer == customer {
                        res <- timeData
                    }
                }
            }
            close(res)
        }()
        return res
    }
    self.addFilter(filter)
}

func (self *Filter) ForMonth(month time.Month) {
    filter := func(data chan *parser.TimeData) chan *parser.TimeData {
        res := make(chan *parser.TimeData)
        go func() {
            for timeData := range data {
                if timeData.Date().Month() == month {
                    res <- timeData
                }
            }
            close(res)
        }()
        return res
    }
    self.addFilter(filter)
}

// Data returns all filtered TimeData element.
func (self *Filter) Data() chan *parser.TimeData {
    res := self.data
    for _, filterFunc := range self.filters {
        res = filterFunc(res)
    }
    return res
}
