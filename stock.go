package main

import "encoding/json"

type Stock struct {
	StoreID     int    `json:"storeId" csv:"storeId"`
	ProductID   int    `json:"productId" csv:"productId"`
	StockAmount string `json:"stockAmount" csv:"stockAmount"`
	UpdDateTime string `json:"updDateTime" csv:"updDateTime"`
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
	ID             int    `json:"id" csv:"id"`
	UpdDateTime    string `json:"updDateTime" csv:"updDateTime"`
	TargetDateTime string `json:"targetDateTime" csv:"targetDateTime"`
	ProductID      int    `json:"productId" csv:"productId"`
	StoreID        int    `json:"storeId" csv:"storeId"`
	Amount         int    `json:"amount" csv:"amount"`
	StockAmount    int    `json:"stockAmount" csv:"stockAmount"`
	StockDivision  string `json:"stockDivision" csv:"stockDivision"`
	FromStoreID    *int   `json:"fromStoreId" csv:"fromStoreId"`
	ToStoreID      *int   `json:"toStoreId" csv:"toStoreId"`
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
