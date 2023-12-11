package main

import (
	"encoding/json"
)

type Shipment struct {
	ShipmentID       json.Number `json:"shipmentId" csv:"shipmentId"`
	ShipmentStoreID  json.Number `json:"shipmentStoreId" csv:"shipmentStoreId"`
	RecipientType    string      `json:"recipientType" csv:"recipientType"`
	RecipientID      json.Number `json:"recipientId" csv:"recipientId"`
	RecipientName    string      `json:"recipientName" csv:"recipientName"`
	ShipmentDivision string      `json:"shipmentDivision" csv:"shipmentDivision"`
	ShipmentDate     string      `json:"shipmentDate" csv:"shipmentDate"`
	Status           string      `json:"status" csv:"status"`
	IdentificationNo *string     `json:"identificationNo" csv:"identificationNo"`
	Modified         string      `json:"modified" csv:"modified"`
}

type ShipmentParquetSchema struct {
	ShipmentID       int64  `json:",string" parquet:"name=shipment_id, type=INT64"`
	ShipmentStoreID  int64  `json:",string" parquet:"name=shipment_store_id, type=INT64"`
	RecipientType    string `parquet:"name=recipient_type, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	RecipientID      int64  `json:",string" parquet:"name=recipient_id, type=INT64"`
	RecipientName    string `parquet:"name=recipient_name, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	ShipmentDivision string `parquet:"name=shipment_division, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	ShipmentDate     string `parquet:"name=shipment_date, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Status           string `parquet:"name=status, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	IdentificationNo string `parquet:"name=identification_no, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Modified         string `parquet:"name=modified, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
}

type ShipmentCSV struct {
	*CSVHandler
	buf []Shipment
}

func NewShipmentCSV(bufSize int, output string) (*ShipmentCSV, error) {
	buf := make([]Shipment, bufSize)
	handler, err := NewCSVHandler([]Shipment{}, output)
	if err != nil {
		return nil, err
	}

	return &ShipmentCSV{
		handler,
		buf,
	}, nil
}

func (sh *ShipmentCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &sh.buf[i])
	}
	sh.CSVWriter.Write(sh.buf[:len(resp.Result)])
	return sh.CSVWriter
}

type ShipmentDetail struct {
	ShipmentID        json.Number `json:"shipmentId" csv:"shipmentId"`
	ProductID         json.Number `json:"productId" csv:"productId"`
	ProductCode       string      `json:"productCode" csv:"productCode"`
	ProductName       string      `json:"productName" csv:"productName"`
	Size              string      `json:"size" csv:"size"`
	Color             string      `json:"color" csv:"color"`
	GroupCode         string      `json:"groupCode" csv:"groupCode"`
	SupplierProductNo *string     `json:"supplierProductNo" csv:"supplierProductNo"`
	Cost              json.Number `json:"cost" csv:"cost"`
	Quantity          json.Number `json:"quantity" csv:"quantity"`
	Memo              string      `json:"memo" csv:"memo"`
	Modified          string      `json:"modified" csv:"modified"`
}

type ShipmentDetailParquetSchema struct {
	ShipmentID        int64  `json:",string" parquet:"name=shipment_id, type=INT64"`
	ProductID         int64  `json:",string" parquet:"name=product_id, type=INT64"`
	ProductCode       string `parquet:"name=product_code, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	ProductName       string `parquet:"name=product_name, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Size              string `parquet:"name=size, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Color             string `parquet:"name=color, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	GroupCode         string `parquet:"name=group_code, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	SupplierProductNo string `parquet:"name=supplier_product_no, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Cost              int64  `json:",string" parquet:"name=cost, type=INT64"`
	Quantity          int64  `json:",string" parquet:"name=quantity, type=INT64"`
	Memo              string `parquet:"name=memo, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Modified          string `parquet:"name=modified, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
}

type ShipmentDetailCSV struct {
	*CSVHandler
	buf []ShipmentDetail
}

func NewShipmentDetailCSV(bufSize int, output string) (*ShipmentDetailCSV, error) {
	buf := make([]ShipmentDetail, bufSize)
	handler, err := NewCSVHandler([]ShipmentDetail{}, output)
	if err != nil {
		return nil, err
	}

	return &ShipmentDetailCSV{
		handler,
		buf,
	}, nil
}

func (shd *ShipmentDetailCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &shd.buf[i])
	}
	shd.CSVWriter.Write(shd.buf[:len(resp.Result)])
	return shd.CSVWriter
}
