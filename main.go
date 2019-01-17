package main

import (
	"fmt"
	"os"
	"strings"
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
		params, err := NewSrRefParamsWithTableName(t)
		if err != nil {
			return err
		}

		cw, err := client.DumpTableToCSV(params)
		if err != nil {
			return err
		}
		cw.Close()
	}
	return nil
}

func main() {
	cliApp := CreateCliApp()
	err := cliApp.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
}
