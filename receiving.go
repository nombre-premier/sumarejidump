package main

import "encoding/json"

type Receiving struct {
	ReceivingID      string      `json:"receivingId" csv:"receivingId"`
	ShippingID       string      `json:"shippingId" csv:"shippingId"`
	ShippingStoreID  string      `json:"shippingStoreId" csv:"shippingStoreId"`
	ReceivingStoreID string      `json:"receivingStoreId" csv:"receivingStoreId"`
	ReceivingDate    string      `json:"receivingDate" csv:"receivingDate"`
	Memo             string      `json:"memo" csv:"memo"`
	Status           string      `json:"status" csv:"status"`
	IdentificationNo interface{} `json:"identificationNo" csv:"identificationNo"`
	Modified         string      `json:"modified" csv:"modified"`
}

type ReceivingCSV struct {
	*CSVHandler
	buf []Receiving
}

func NewReceivingCSV(bufSize int, output string) (*ReceivingCSV, error) {
	buf := make([]Receiving, bufSize)
	handler, err := NewCSVHandler([]Receiving{}, output)
	if err != nil {
		return nil, err
	}

	return &ReceivingCSV{
		handler,
		buf,
	}, nil
}

func (cc *ReceivingCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &cc.buf[i])
	}
	cc.CSVWriter.Write(cc.buf[:len(resp.Result)])
	return cc.CSVWriter
}

type ReceivingDetail struct {
	ReceivingID        int     `json:"receivingId" csv:"receivingId"`
	ProductID          int     `json:"productId" csv:"productId"`
	ProductCode        string  `json:"productCode" csv:"productCode"`
	ProductName        string  `json:"productName" csv:"productName"`
	Size               string  `json:"size" csv:"size"`
	Color              string  `json:"color" csv:"color"`
	GroupCode          *string `json:"groupCode" csv:"groupCode"`
	SupplierProductNo  *string `json:"supplierProductNo" csv:"supplierProductNo"`
	ScheduledQuantity  string  `json:"scheduledQuantity" csv:"scheduledQuantity"`
	InspectionQuantity string  `json:"inspectionQuantity" csv:"inspectionQuantity"`
	StockoutQuantity   string  `json:"stockoutQuantity" csv:"stockoutQuantity"`
	StockoutReason     *string `json:"stockoutReason" csv:"stockoutReason"`
	InspectionDate     string  `json:"inspectionDate" csv:"inspectionDate"`
	Status             string  `json:"status" csv:"status"`
	Modified           string  `json:"modified" csv:"modified"`
}

type ReceivingDetailCSV struct {
	*CSVHandler
	buf []ReceivingDetail
}

func NewReceivingDetailCSV(bufSize int, output string) (*ReceivingDetailCSV, error) {
	buf := make([]ReceivingDetail, bufSize)
	handler, err := NewCSVHandler([]ReceivingDetail{}, output)
	if err != nil {
		return nil, err
	}

	return &ReceivingDetailCSV{
		handler,
		buf,
	}, nil
}

func (cc *ReceivingDetailCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &cc.buf[i])
	}
	cc.CSVWriter.Write(cc.buf[:len(resp.Result)])
	return cc.CSVWriter
}
