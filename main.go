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

	//	tables := []string{"Category",
	//		"Store",
	//		"Product", "ProductPrice",
	//		"ProductReserveItem", "ProductReserveItemLabel",
	//		"ProductStore", "ProductInventoryReservation",
	//		"Customer",
	//		"Stock", "StockHistory",
	//		"TransactionHead", "TransactionDetail",
	//		"Bargain", "BargainStore", "BargainProduct",
	//		"DailySum",
	//	}
	tables := []string{"Category"}

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
