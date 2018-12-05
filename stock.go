package main

import "encoding/json"

type Stock struct {
	StoreID     int    `json:"storeId"`
	ProductID   int    `json:"productId"`
	StockAmount string `json:"stockAmount"`
	UpdDateTime string `json:"updDateTime"`
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
	ID             int         `json:"id"`
	UpdDateTime    string      `json:"updDateTime"`
	TargetDateTime string      `json:"targetDateTime"`
	ProductID      int         `json:"productId"`
	StoreID        int         `json:"storeId"`
	Amount         int         `json:"amount"`
	StockAmount    int         `json:"stockAmount"`
	StockDivision  string      `json:"stockDivision"`
	FromStoreID    interface{} `json:"fromStoreId"`
	ToStoreID      interface{} `json:"toStoreId"`
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
