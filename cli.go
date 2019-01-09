package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli"
)

func CreateCliApp() *cli.App {
	app := cli.NewApp()
	app.Name = "sumarejidump"
	app.Usage = "Just dump sumareji data"
	app.Version = "0.0.1"
	app.UsageText = "sumarejidump [command] [options] [table_name]"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "token, t",
			Usage:  "sumareji access token",
			EnvVar: "SUMAREJI_ACCESS_TOKEN",
		},
		cli.StringFlag{
			Name:   "contract_id, i",
			Usage:  "sumareji contract_id",
			EnvVar: "SUMAREJI_CONTRACT_ID",
		},
		// TODO: implement conditions
		cli.StringSliceFlag{
			Name:  "conditions, c",
			Usage: "filter data by given conditon(s)",
		},
		// TODO: implement dir config
		cli.StringSliceFlag{
			Name:  "output, o",
			Usage: "output dir name, default: yyyyMMDDhhmmss",
		},
	}
	app.Action = cliAction

	return app
}

func cliAction(c *cli.Context) error {
	t := time.Now()
	dirName := t.Format(dirFormat)
	if err := os.MkdirAll(fmt.Sprintf("%s", dirName), 0755); err != nil {
		panic(err)
	}

	config := SrConfig{
		ContractID:  c.String("contract_id"),
		AccessToken: c.String("token"),
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dirName,
	}

	client := NewSrClient(config)
	var tables []string = make([]string, 0, len(SrTableMetas))
	tableName := c.Args().Get(0)

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
			panic(err)
		}

		cw, err := client.DumpTableToCSV(params)
		if err != nil {
			panic(err)
		}
		cw.Close()
	}
	return nil

}
