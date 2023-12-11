package main

import "encoding/json"

type Stock struct {
	StoreID     json.Number `json:"storeId" csv:"storeId"`
	ProductID   json.Number `json:"productId" csv:"productId"`
	StockAmount string      `json:"stockAmount" csv:"stockAmount"`
	UpdDateTime string      `json:"updDateTime" csv:"updDateTime"`
}

type StockParquetSchema struct {
	StoreID     int64  `json:",string" parquet:"name=store_id, type=INT64"`
	ProductID   int64  `json:",string" parquet:"name=product_id, type=INT64"`
	StockAmount string `parquet:"name=stock_amount, type=BYTE_ARRAY, convertedtype=UTF8"`
	UpdDateTime string `parquet:"name=upd_date_time, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type StockCSV struct {
	*CSVHandler
	buf []Stock
}

func NewStockCSV(bufSize int, output string) (*StockCSV, error) {
	buf := make([]Stock, bufSize)
	handler, err := NewCSVHandler([]Stock{}, output)
	if err != nil {
		return nil, err
	}

	return &StockCSV{
		handler,
		buf,
	}, nil
}

func (sc *StockCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &sc.buf[i])
	}
	sc.CSVWriter.Write(sc.buf[:len(resp.Result)])
	return sc.CSVWriter
}

type StockHistory struct {
	ID             json.Number  `json:"id" csv:"id"`
	UpdDateTime    string       `json:"updDateTime" csv:"updDateTime"`
	TargetDateTime string       `json:"targetDateTime" csv:"targetDateTime"`
	ProductID      json.Number  `json:"productId" csv:"productId"`
	StoreID        json.Number  `json:"storeId" csv:"storeId"`
	Amount         json.Number  `json:"amount" csv:"amount"`
	StockAmount    json.Number  `json:"stockAmount" csv:"stockAmount"`
	StockDivision  string       `json:"stockDivision" csv:"stockDivision"`
	FromStoreID    *json.Number `json:"fromStoreId" csv:"fromStoreId"`
	ToStoreID      *json.Number `json:"toStoreId" csv:"toStoreId"`
}

type StockHistoryParquetSchema struct {
	ID             int64  `json:",string" parquet:"name=id, type=INT64"`
	UpdDateTime    string `parquet:"name=upd_date_time, type=BYTE_ARRAY, convertedtype=UTF8"`
	TargetDateTime string `parquet:"name=target_date_time, type=BYTE_ARRAY, convertedtype=UTF8"`
	ProductID      int64  `json:",string" parquet:"name=product_id, type=INT64"`
	StoreID        int64  `json:",string" parquet:"name=store_id, type=INT64"`
	Amount         int64  `json:",string" parquet:"name=amount, type=INT64"`
	StockAmount    int64  `json:",string" parquet:"name=stock_amount, type=INT64"`
	StockDivision  string `parquet:"name=stock_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	FromStoreID    *int64 `json:",string" parquet:"name=from_store_id, type=INT64, repetitiontype=OPTIONAL"`
	ToStoreID      *int64 `json:",string" parquet:"name=to_store_id, type=INT64, repetitiontype=OPTIONAL"`
}

type StockHistoryCSV struct {
	*CSVHandler
	buf []StockHistory
}

func NewStockHistoryCSV(bufSize int, output string) (*StockHistoryCSV, error) {
	buf := make([]StockHistory, bufSize)
	handler, err := NewCSVHandler([]StockHistory{}, output)
	if err != nil {
		return nil, err
	}

	return &StockHistoryCSV{
		handler,
		buf,
	}, nil
}

func (shc *StockHistoryCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &shc.buf[i])
	}
	shc.CSVWriter.Write(shc.buf[:len(resp.Result)])
	return shc.CSVWriter
}
