package main

import (
	"encoding/json"
)

type Loss struct {
	LossID           json.Number `json:"lossId" csv:"lossId"`
	StoreID          json.Number `json:"storeId" csv:"storeId"`
	Division         string      `json:"division" csv:"division"`
	Memo             string      `json:"memo" csv:"memo"`
	LossDatetime     string      `json:"lossDatetime" csv:"lossDatetime"`
	IdentificationNo *string     `json:"identificationNo" csv:"identificationNo"`
	Modified         string      `json:"modified" csv:"modified"`
}

type LossParquetSchema struct {
	LossID           int64  `parquet:"name=loss_id, type=INT64"`
	StoreID          int64  `parquet:"name=store_id, type=INT64"`
	Division         string `parquet:"name=division, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Memo             string `parquet:"name=memo, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	LossDatetime     string `parquet:"name=loss_datetime, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	IdentificationNo string `parquet:"name=identification_no, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Modified         string `parquet:"name=modified, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
}

type LossCSV struct {
	*CSVHandler
	buf []Loss
}

func NewLossCSV(bufSize int, output string) (*LossCSV, error) {
	buf := make([]Loss, bufSize)
	handler, err := NewCSVHandler([]Loss{}, output)
	if err != nil {
		return nil, err
	}

	return &LossCSV{
		handler,
		buf,
	}, nil
}

func (ls *LossCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &ls.buf[i])
	}
	ls.CSVWriter.Write(ls.buf[:len(resp.Result)])
	return ls.CSVWriter
}

type LossDetail struct {
	LossID            json.Number `json:"lossId" csv:"lossId"`
	ProductID         json.Number `json:"productId" csv:"productId"`
	ProductCode       string      `json:"productCode" csv:"productCode"`
	ProductName       string      `json:"productName" csv:"productName"`
	Size              string      `json:"size" csv:"size"`
	Color             string      `json:"color" csv:"color"`
	GroupCode         string      `json:"groupCode" csv:"groupCode"`
	SupplierProductNo *string     `json:"supplierProductNo" csv:"supplierProductNo"`
	Quantity          json.Number `json:"quantity" csv:"quantity"`
	Modified          string      `json:"modified" csv:"modified"`
}

type LossDetailParquetSchema struct {
	LossID            int64  `parquet:"name=loss_id, type=INT64"`
	ProductID         int64  `parquet:"name=product_id, type=INT64"`
	ProductCode       string `parquet:"name=product_code, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	ProductName       string `parquet:"name=product_name, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Size              string `parquet:"name=size, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Color             string `parquet:"name=color, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	GroupCode         string `parquet:"name=group_code, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	SupplierProductNo string `parquet:"name=supplier_product_no, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Quantity          int64  `parquet:"name=quantity, type=INT64"`
	Modified          string `parquet:"name=modified, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
}

type LossDetailCSV struct {
	*CSVHandler
	buf []LossDetail
}

func NewLossDetailCSV(bufSize int, output string) (*LossDetailCSV, error) {
	buf := make([]LossDetail, bufSize)
	handler, err := NewCSVHandler([]LossDetail{}, output)
	if err != nil {
		return nil, err
	}

	return &LossDetailCSV{
		handler,
		buf,
	}, nil
}

func (ldl *LossDetailCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &ldl.buf[i])
	}
	ldl.CSVWriter.Write(ldl.buf[:len(resp.Result)])
	return ldl.CSVWriter
}
