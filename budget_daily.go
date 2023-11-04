package main

import (
	"encoding/json"
	"fmt"
)

type BudgetDaily struct {
	StoreID          json.Number `json:"storeId" csv:"storeId"`
	YM               string      `json:"ym" csv:"ym"`
	Day              string      `json:"day" csv:"day"`
	SalesTargetDaily json.Number `json:"salesTargetDaily" csv:"salesTargetDaily"`
}

type BudgetDailyParquetSchema struct {
	StoreID 				 int64  `parquet:"name=store_id, type=INT64"`
	YM      				 string `parquet:"name=ym, type=BYTE_ARRAY, convertedtype=UTF8", encoding=PLAIN_DICTIONARY`
	Day     				 string `parquet:"name=day, type=BYTE_ARRAY, convertedtype=UTF8", encoding=PLAIN_DICTIONARY`
	SalesTargetDaily int64  `parquet:"name=sales_target_daily, type=INT64"`
}

type BudgetDailyCSV struct {
	*CSVHandler
	buf []BudgetDaily
}

type BudgetDailyParquet struct {
	*ParquetHandler
	buf []BudgetDailyParquetSchema
}

func NewBudgetDailyCSV(bufSize int, output string) (*BudgetDailyCSV, error) {
	buf := make([]BudgetDaily, bufSize)
	handler, err := NewCSVHandler([]BudgetDaily{}, output)
	if err != nil {
		return nil, err
	}

	return &BudgetDailyCSV{
		handler,
		buf,
	}, nil
}

func NewBudgetDailyParquet(output string) (*BudgetDailyParquet, error) {
	ph, err := NewParquetHandler(new(BudgetDailyParquetSchema), output)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate ParquetHandler: %w\n", err)
	}
	return &BudgetDailyParquet{
		ParquetHandler: ph,
	}, nil
}

func (cp *BudgetDailyParquet) Write(resp *SrRefResponse) error {
	for _, r := range resp.Result {
		var budgetDaily BudgetDailyParquetSchema
		json.Unmarshal([]byte(r.String()), &budgetDaily)
		if err := cp.ParquetHandler.ParquetWriter.Write(budgetDaily); err != nil {
			return fmt.Errorf("failed to write parquet: %w\n", err)
		}
	}
	return nil
}

func (bdc *BudgetDailyCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &bdc.buf[i])
	}
	bdc.CSVWriter.Write(bdc.buf[:len(resp.Result)])
	return bdc.CSVWriter
}
