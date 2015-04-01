package main

import (
    "fmt"
    "os"
    "strings"
    "time"

    "github.com/codegangsta/cli"

    "github.com/cbrand/jiffyreporter/parser"
    "github.com/cbrand/jiffyreporter/aggregator"
    "github.com/cbrand/jiffyreporter/dumper"
    "github.com/cbrand/jiffyreporter/filter"
)

const fileMode = 0660

func printError(err error) {
    fmt.Println("Error")
    fmt.Println(err.Error())
}

func extractCustomers(customers string) []string {
    cArray := strings.Split(customers, ",")
    resArray := []string{}

    for _, customerString := range cArray {
        customerString = strings.TrimSpace(customerString)
        if len(customerString) > 0 {
            resArray = append(resArray, customerString)
        }
    }
    return resArray
}

func main() {
    app := cli.NewApp()
    app.Name = "jiffyreporter"
    app.Usage = "converts jiffy report csvs to excel time sheets."
    app.Version = "1.0"
    app.Flags = []cli.Flag {
        cli.StringFlag{
            Name: "customer, c",
            Value: "",
            Usage: "comma separated names of the customers which should be used.",
        },
        cli.IntFlag{
            Name: "month, m",
            Value: -1,
            Usage: "The month which should be selected.",
        },
    }
    app.Action = func(c *cli.Context) {
        if len(c.Args()) != 2 {
            fmt.Println("Need two arguments (from) and (to).")
            return
        }
        fromCSV := c.Args().Get(0)
        toExcel := c.Args().Get(1)

        f, err := os.OpenFile(fromCSV, os.O_RDONLY, fileMode)
        if err != nil {
            printError(err)
            return
        }
        defer f.Close()
        dumpFile, err := os.OpenFile(toExcel, os.O_WRONLY | os.O_CREATE, fileMode)
        if err != nil {
            printError(err)
            return
        }
        defer dumpFile.Close()

        p := parser.New(f)
        timeData, err := p.Data()
        if err != nil {
            printError(err)
            return
        }

        filterData := filter.NewFromArray(timeData)

        customers := extractCustomers(c.String("customer"))
        if len(customers) > 0 {
            filterData.ForCustomers(customers...)
        }
        if c.Int("month") >= 1 && c.Int("month") <= 12 {
            filterData.ForMonth(time.Month(c.Int("month")))
        }

        aggregated := aggregator.Aggregate(filterData.Data())
        dumper.New(aggregated).Write(dumpFile)

    }

    app.Run(os.Args)
}
