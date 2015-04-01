package dumper

import (
    "io"
    "strconv"

    "github.com/tealeg/xlsx"
    "github.com/weidewang/go-strftime"

    "github.com/cbrand/jiffyreporter/aggregator"
)


type Dumper struct {
    data aggregator.ReportTimeDataSlice
    file *xlsx.File
    sheet *xlsx.Sheet
    headerStyle *xlsx.Style
}

func New(data aggregator.ReportTimeDataSlice) *Dumper {
    xlsxHeaderStyle := xlsx.NewStyle()
    xlsxHeaderStyle.Font.Bold = true

    return &Dumper{
        data: data,
        headerStyle: xlsxHeaderStyle,
    }
}

func (self *Dumper) Write(writer io.Writer) error {
    self.generate()
    return self.file.Write(writer)
}

// generate
func (self *Dumper) generate() {
    self.file = xlsx.NewFile()
    self.sheet = self.file.AddSheet("timedata")
    self.generateHeader()
    self.generateFromData()
    self.generateSummary()
}

// generateHeader prints the header entry.
func (self *Dumper) generateHeader() {
    row := self.sheet.AddRow()
    header := []string{
        "Project",
        "Date",
        "Time in Hours",
    }
    for _, headerText := range header {
        cell := row.AddCell()
        cell.Value = headerText
        cell.SetStyle(self.headerStyle)
    }
}

// generateFromData adds the data to the current sheet..
func (self *Dumper) generateFromData() {
    for _, item := range self.data {
        self.writeLine(item)
    }
}

// generateSummary adds the formula to the excel sheet to
// add the hours together.
func (self *Dumper) generateSummary() {
    row := self.sheet.AddRow()

    cell := row.AddCell()
    cell.Value = "Summary"
    cell.SetStyle(self.headerStyle)
    row.AddCell()
    cell = row.AddCell()
    cell.SetStyle(self.headerStyle)
    cell.SetFormula("C2:C" + strconv.FormatInt(int64(self.sheet.MaxRow), 10))
}

// writeLine generates the line for the passed data.
func (self *Dumper) writeLine(item *aggregator.ReportTimeData) {
    row := self.sheet.AddRow()
    cell := row.AddCell()
    cell.Value = item.Project()
    cell = row.AddCell()
    date := item.Date()
    cell.Value = strftime.Strftime(&date, "%Y-%m-%d")
    cell = row.AddCell()
    cell.SetFloatWithFormat(item.Duration().Hours(), "#,##0.00")
}
