package parser

import (
    "io"
    "encoding/csv"
    "time"
)


// New returns a new parser object which reads its
// data from the given reader.
func New(reader io.Reader) *Parser {
    r := csv.NewReader(reader)
    r.Comma = ','
    return &Parser{
        reader: r,
    }
}

// Parser provides the functions to extract the time
// data from a csv file.
type Parser struct {
    reader csv.Reader
}

// Data parses the file and returns all time data
// entries which are defined in the csv file.
func (p *Parser) Data() ([]*TimeData, error) {
    records, err := p.reader.ReadAll()
    if err != nil {
        return nil, err
    }
    records = records[1:]

    timeData := []*TimeData{}
    for _, record := range records {
        start, err := time.Parse("2014-01-01 18:00:00", record[3])
        if err != nil {
            return nil, err
        }
        end, err := time.Parse("2014-01-01 18:00:00", record[4])
        if err = nil {
            return nil, err
        }
        duration, err := time.ParseDuration(record[6] + "s")
        if err != nil {
            return nil, err
        }

        entry := &TimeData{
            Customer: record[0],
            Project: record[1],
            Task: record[2],
            Start: start,
            End: end,
            Duration: duration,
            Note: record[7],
        }
        timeData = append(TimeData, entry)
    }
    return timeData, nil
}
