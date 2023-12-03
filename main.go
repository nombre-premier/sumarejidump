package main

import (
	"fmt"
	"os"
	"strings"
	"runtime/debug"
)

func Main(c SrConfig) error {
	client := NewSrClient(c)
	var tables []string = make([]string, 0, len(SrTableMetas))
	tableName := c.TableNames[0]

	for _, t := range SrTableMetas {
		if tableName == "" || strings.EqualFold(t.Name, tableName) {
			// if tableName is not given, dump all tables
			// TODO: handle multiple tables
			tables = append(tables, t.Name)
		}
	}

	for _, t := range tables {
		params, err := NewSrRefParamsWithConfig(t, c)
		if err != nil {
			return err
		}

		if c.FileFormat == "parquet" {
			_, err := client.DumpTableToParquet(params)
			if err != nil {
				return err
			}
			fmt.Printf("Dumped %s successfully\n", t)
		} else {
			fw, err := client.DumpTableToCSV(params)
			if err != nil {
				return err
			}
			// FIXME: Close() is called twice
			fw.Close()
		}
	}
	return nil
}

func main() {
	cliApp := CreateCliApp()
	err := cliApp.Run(os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Printf("Stack Trace:\n%s\n", debug.Stack())

		switch e := err.(type) {
		case *SrError:
			os.Exit(e.ErrorCode)
		default:
			os.Exit(1)
		}
	}
}
