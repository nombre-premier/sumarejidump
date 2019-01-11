package main

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type Bargain struct {
	BargainID   json.Number `json:"bargainId" csv:"bargainId"`
	BargainName string      `json:"bargainName" csv:"bargainName"`
	TermStart   string      `json:"termStart" csv:"termStart"`
	TermEnd     string      `json:"termEnd" csv:"termEnd"`
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
	BargainProductID json.Number     `json:"bargainProductId" csv:"bargainProductId"`
	BargainID        json.Number     `json:"bargainId" csv:"bargainId"`
	TargetDivision   string          `json:"targetDivision" csv:"targetDivision"`
	TargetID         string          `json:"targetId" csv:"targetId"`
	Division         string          `json:"division" csv:"division"`
	Value            decimal.Decimal `json:"value" csv:"value"`
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
	BargainStoreID json.Number `json:"bargainStoreId" csv:"bargainStoreId"`
	BargainID      json.Number `json:"bargainId" csv:"bargainId"`
	StoreID        json.Number `json:"storeId" csv:"storeId"`
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
