package main

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type Bargain struct {
	BargainID   int    `json:"bargainId"`
	BargainName string `json:"bargainName"`
	TermStart   string `json:"termStart"`
	TermEnd     string `json:"termEnd"`
}

type BargainCSV struct {
	*CSVHandler
	buf []Bargain
}

func NewBargainCSV(bufSize int, output string) (*BargainCSV, error) {
	buf := make([]Bargain, bufSize)
	handler, err := NewCSVHandler([]Bargain{}, output)
	if err != nil {
		return nil, err
	}

	return &BargainCSV{
		handler,
		buf,
	}, nil
}

func (bc *BargainCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &bc.buf[i])
	}
	bc.CSVWriter.Write(bc.buf[:len(resp.Result)])
	return bc.CSVWriter
}

type BargainProduct struct {
	BargainProductID int             `json:"bargainProductId"`
	BargainID        int             `json:"bargainId"`
	TargetDivision   string          `json:"targetDivision"`
	TargetID         string          `json:"targetId"`
	Division         string          `json:"division"`
	Value            decimal.Decimal `json:"value"`
}

type BargainProductCSV struct {
	*CSVHandler
	buf []BargainProduct
}

func NewBargainProductCSV(bufSize int, output string) (*BargainProductCSV, error) {
	buf := make([]BargainProduct, bufSize)
	handler, err := NewCSVHandler([]BargainProduct{}, output)
	if err != nil {
		return nil, err
	}

	return &BargainProductCSV{
		handler,
		buf,
	}, nil
}

func (bpc *BargainProductCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &bpc.buf[i])
	}
	bpc.CSVWriter.Write(bpc.buf[:len(resp.Result)])
	return bpc.CSVWriter
}

type BargainStore struct {
	BargainStoreID int `json:"bargainStoreId"`
	BargainID      int `json:"bargainId"`
	StoreID        int `json:"storeId"`
}

type BargainStoreCSV struct {
	*CSVHandler
	buf []BargainStore
}

func NewBargainStoreCSV(bufSize int, output string) (*BargainStoreCSV, error) {
	buf := make([]BargainStore, bufSize)
	handler, err := NewCSVHandler([]BargainStore{}, output)
	if err != nil {
		return nil, err
	}

	return &BargainStoreCSV{
		handler,
		buf,
	}, nil
}

func (bsc *BargainStoreCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &bsc.buf[i])
	}
	bsc.CSVWriter.Write(bsc.buf[:len(resp.Result)])
	return bsc.CSVWriter
}
