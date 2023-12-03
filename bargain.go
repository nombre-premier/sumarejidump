package main

import (
	"encoding/json"
)

type Bargain struct {
	BargainID   json.Number `json:"bargainId" csv:"bargainId"`
	BargainName string      `json:"bargainName" csv:"bargainName"`
	TermStart   string      `json:"termStart" csv:"termStart"`
	TermEnd     string      `json:"termEnd" csv:"termEnd"`
}

type BargainParquetSchema struct {
	BargainID   int64  `parquet:"name=bargain_id, type=INT64"`
	BargainName string `parquet:"name=bargain_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	TermStart   string `parquet:"name=term_start, type=BYTE_ARRAY, convertedtype=UTF8"`
	TermEnd     string `parquet:"name=term_end, type=BYTE_ARRAY, convertedtype=UTF8"`
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
	BargainProductID json.Number `json:"bargainProductId" csv:"bargainProductId"`
	BargainID        json.Number `json:"bargainId" csv:"bargainId"`
	TargetDivision   string      `json:"targetDivision" csv:"targetDivision"`
	TargetID         string      `json:"targetId" csv:"targetId"`
	Division         string      `json:"division" csv:"division"`
	Value            json.Number `json:"value" csv:"value"`
}

type BargainProductParquetSchema struct {
	BargainProductID int64  `parquet:"name=bargain_product_id, type=INT64"`
	BargainID        int64  `parquet:"name=bargain_id, type=INT64"`
	TargetDivision   string `parquet:"name=target_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	TargetID         string `parquet:"name=target_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Division         string `parquet:"name=division, type=BYTE_ARRAY, convertedtype=UTF8"`
	Value            int64  `parquet:"name=value, type=INT64"`
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

type BargainStoreParquetSchema struct {
	BargainStoreID int64 `parquet:"name=bargain_store_id, type=INT64"`
	BargainID      int64 `parquet:"name=bargain_id, type=INT64"`
	StoreID        int64 `parquet:"name=store_id, type=INT64"`
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
