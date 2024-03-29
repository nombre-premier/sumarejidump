package main

import (
	"os"
	"strings"
	"time"

	"github.com/urfave/cli"
)

func CreateCliApp() *cli.App {
	app := cli.NewApp()
	app.Name = "sumarejidump"
	app.Usage = "Just dump sumareji data"
	app.Version = "0.1.1"
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
		cli.StringSliceFlag{
			Name:  "conditions, c",
			Usage: "filter data by given conditon(s)",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "output dir name, default: yyyyMMDDhhmmss",
		},
		cli.StringFlag{
      Name: "format, f",
      Usage: "output format, options: csv, parquet",
      Value: "csv",  // default value
    },
	}
	app.Action = cliAction

	return app
}

func cliAction(c *cli.Context) error {
	t := time.Now()
	dirName := ""
	if c.String("output") == "" {
		dirName = t.Format(dirFormat)
	} else {
		dirName = c.String("output")
	}
	if err := os.MkdirAll(dirName, 0755); err != nil {
		panic(err)
	}

	format := ""
	if c.String("format") == "parquet" {
		format = "parquet"
	} else {
		format = "csv"
	}

	config := SrConfig{
		ContractID:  c.String("contract_id"),
		AccessToken: c.String("token"),
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dirName,
		TableNames:  []string{c.Args().Get(0)},
		Conditions:  conditionParser(c.StringSlice("conditions")),
		FileFormat:  format,
	}
	return Main(config)
}

func conditionParser(conditions []string) map[string]*string {
	cond := make(map[string]*string)
	for _, c := range conditions {
		k := c[:strings.Index(c, ":")]
		v := c[strings.Index(c, ":")+1:]
		k = strings.Trim(k, " ")
		if v == "null" {
			cond[k] = nil
		} else {
			cond[k] = &v
		}
	}
	return cond
}
