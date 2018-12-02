package main

import (
	"os"
)

func main() {
	contractID := os.Getenv("SUMAREJI_CONTRACT_ID")
	accessToken := os.Getenv("SUMAREJI_ACCESS_TOKEN")
	client := NewSrClient(contractID, accessToken)

	params, err := NewSrRefParamsWithTableName("Category")
	if err != nil {
		panic(err)
	}

	cw, err := client.DumpTableToCSV(params)
	if err != nil {
		panic(err)
	}
	defer cw.Close()

	params2, err := NewSrRefParamsWithTableName("Store")
	if err != nil {
		panic(err)
	}

	cw2, err := client.DumpTableToCSV(params2)
	if err != nil {
		panic(err)
	}
	defer cw2.Close()
}
