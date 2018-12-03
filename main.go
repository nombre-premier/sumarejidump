package main

import (
	"os"
)

func main() {
	contractID := os.Getenv("SUMAREJI_CONTRACT_ID")
	accessToken := os.Getenv("SUMAREJI_ACCESS_TOKEN")
	client := NewSrClient(contractID, accessToken)

	tables := []string{"Category", "Store",
		"Product", "ProductPrice",
		"ProductReserveItem", "ProductReserveItemLabel",
		"ProductStore", "ProductInventoryReservation",
		"Customer",
		"Stock", "StockHistory",
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
}
