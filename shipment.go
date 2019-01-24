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
