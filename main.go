package main

import (
	"os"
)

func main() {
	contractID := os.Getenv("SUMAREJI_CONTRACT_ID")
	accessToken := os.Getenv("SUMAREJI_ACCESS_TOKEN")
	client := NewSumarejiClient(contractID, accessToken)

	params := SumarejiRefParams{ProcName: "category_ref",
		Fields: []string{}, Conditions: []string{},
		Order: []string{}, Limit: 1000, Page: 1,
		TableName: "Category",
	}
	cw, err := client.DumpTableToCSV(params)
	if err != nil {
		panic(err)
	}
	defer cw.Close()

	params2 := SumarejiRefParams{ProcName: "store_ref",
		Fields: []string{}, Conditions: []string{},
		Order: []string{}, Limit: 1000, Page: 1,
		TableName: "Store",
	}
	cw2, err := client.DumpTableToCSV(params2)
	if err != nil {
		panic(err)
	}
	defer cw2.Close()
}
