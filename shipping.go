package main

import "encoding/json"

type Shipping struct {
	ShippingID                         json.Number `json:"shippingId" csv:"shippingId"`
	ShippingStoreID                    json.Number `json:"shippingStoreId" csv:"shippingStoreId"`
	ReceivingStoreID                   json.Number `json:"receivingStoreId" csv:"receivingStoreId"`
	ReceivingExpectedDateFrom          string      `json:"receivingExpectedDateFrom" csv:"receivingExpectedDateFrom"`
	ReceivingExpectedDateTo            string      `json:"receivingExpectedDateTo" csv:"receivingExpectedDateTo"`
	ShippingDate                       string      `json:"shippingDate" csv:"shippingDate"`
	ReceivingDesiredDate               *string     `json:"receivingDesiredDate" csv:"receivingDesiredDate"`
	Memo                               string      `json:"memo" csv:"memo"`
	Status                             string      `json:"status" csv:"status"`
	ModificationRequestStatus          string      `json:"modificationRequestStatus" csv:"modificationRequestStatus"`
	ModificationRequestDateTime        *string     `json:"modificationRequestDateTime" csv:"modificationRequestDateTime"`
	ModificationRequestCheckedDateTime *string     `json:"modificationRequestCheckedDateTime" csv:"modificationRequestCheckedDateTime"`
	IdentificationNo                   *string     `json:"identificationNo" csv:"identificationNo"`
	Modified                           string      `json:"modified" csv:"modified"`
}

type ShippingCSV struct {
	*CSVHandler
	buf []Shipping
}

func NewShippingCSV(bufSize int, output string) (*ShippingCSV, error) {
	buf := make([]Shipping, bufSize)
	handler, err := NewCSVHandler([]Shipping{}, output)
	if err != nil {
		return nil, err
	}

	return &ShippingCSV{
		handler,
		buf,
	}, nil
}

func (cc *ShippingCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &cc.buf[i])
	}
	cc.CSVWriter.Write(cc.buf[:len(resp.Result)])
	return cc.CSVWriter
}

type ShippingDetail struct {
	ShippingID        json.Number `json:"shippingId"`
	ProductID         json.Number `json:"productId"`
	ProductCode       string      `json:"productCode"`
	ProductName       string      `json:"productName"`
	Size              string      `json:"size"`
	Color             string      `json:"color"`
	GroupCode         *string     `json:"groupCode"`
	SupplierProductNo *string     `json:"supplierProductNo"`
	RequestQuantity   *string     `json:"requestQuantity"`
	Quantity          string      `json:"quantity"`
	Modified          string      `json:"modified"`
}

type ShippingDetailCSV struct {
	*CSVHandler
	buf []ShippingDetail
}

func NewShippingDetailCSV(bufSize int, output string) (*ShippingDetailCSV, error) {
	buf := make([]ShippingDetail, bufSize)
	handler, err := NewCSVHandler([]ShippingDetail{}, output)
	if err != nil {
		return nil, err
	}

	return &ShippingDetailCSV{
		handler,
		buf,
	}, nil
}

func (cc *ShippingDetailCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &cc.buf[i])
	}
	cc.CSVWriter.Write(cc.buf[:len(resp.Result)])
	return cc.CSVWriter
}
