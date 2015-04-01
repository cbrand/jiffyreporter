package main

import (
    "fmt"
    "os"

    "github.com/codegangsta/cli"

    "github.com/cbrand/jiffyreporter/parser"
    "github.com/cbrand/jiffyreporter/aggregator"
    "github.com/cbrand/jiffyreporter/dumper"
)

const fileMode = 0660

func printError(err error) {
    fmt.Println("Error")
    fmt.Println(err.Error())
}

func main() {
    app := cli.NewApp()
    app.Name = "jiffyreporter"
    app.Usage = "converts jiffy report csvs to excel time sheets."
    app.Version = "1.0"
    app.Flags = []cli.Flag {
        cli.StringFlag{
            Name: "customer, c",
            Value: "customer",
            Usage: "comma separated names of the customers which should be used.",
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
        dumpFile, err := os.OpenFile(toExcel, os.O_WRONLY, fileMode)
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

        aggregated := aggregator.AggregateArray(timeData)
        dumper.New(aggregated).Write(dumpFile)

    }

    app.Run(os.Args)
}
