package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	contractID := os.Getenv("SUMAREJI_CONTRACT_ID")
	accessToken := os.Getenv("SUMAREJI_ACCESS_TOKEN")

	t := time.Now()
	dirName := t.Format(dirFormat)
	if err := os.MkdirAll(fmt.Sprintf("%s", dirName), 0755); err != nil {
		panic(err)
	}

	config := SrConfig{
		ContractID:  contractID,
		AccessToken: accessToken,
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dirName,
	}

	client := NewSrClient(config)

	for _, t := range srTableMetas {
		params, err := NewSrRefParamsWithTableName(t.Name)
		if err != nil {
			panic(err)
		}

		_, err = client.DumpTableToCSV(params)
		if err != nil {
			panic(err)
		}
	}
}
