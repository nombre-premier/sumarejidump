package main

import (
	"encoding/json"
)

type StorageInfo struct {
	StorageInfoID    json.Number `json:"storageInfoId" csv:"storageInfoId"`
	RecipientOrderID json.Number `json:"recipientOrderId" csv:"recipientOrderId"`
	OrderedDate      string      `json:"orderedDate" csv:"orderedDate"`
	Status           string      `json:"status" csv:"status"`
	IdentificationNo *string     `json:"identificationNo" csv:"identificationNo"`
	Modified         string      `json:"modified" csv:"modified"`
}

type StorageInfoCSV struct {
	*CSVHandler
	buf []StorageInfo
}

func NewStorageInfoCSV(bufSize int, output string) (*StorageInfoCSV, error) {
	buf := make([]StorageInfo, bufSize)
	handler, err := NewCSVHandler([]StorageInfo{}, output)
	if err != nil {
		return nil, err
	}

	return &StorageInfoCSV{
		handler,
		buf,
	}, nil
}

func (si *StorageInfoCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &si.buf[i])
	}
	si.CSVWriter.Write(si.buf[:len(resp.Result)])
	return si.CSVWriter
}

type StorageInfoDelivery struct {
	StorageInfoID           json.Number `json:"storageInfoId" csv:"storageInfoId"`
	StorageStoreID          json.Number `json:"storageStoreId" csv:"storageStoreId"`
	StorageExpectedDateFrom string      `json:"storageExpectedDateFrom" csv:"storageExpectedDateFrom"`
	StorageExpectedDateTo   string      `json:"storageExpectedDateTo" csv:"storageExpectedDateTo"`
	Modified                string      `json:"modified" csv:"modified"`
}

type StorageInfoDeliveryCSV struct {
	*CSVHandler
	buf []StorageInfoDelivery
}

func NewStorageInfoDeliveryCSV(bufSize int, output string) (*StorageInfoDeliveryCSV, error) {
	buf := make([]StorageInfoDelivery, bufSize)
	handler, err := NewCSVHandler([]StorageInfoDelivery{}, output)
	if err != nil {
		return nil, err
	}

	return &StorageInfoDeliveryCSV{
		handler,
		buf,
	}, nil
}

func (sid *StorageInfoDeliveryCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &sid.buf[i])
	}
	sid.CSVWriter.Write(sid.buf[:len(resp.Result)])
	return sid.CSVWriter
}

type StorageInfoProduct struct {
	StorageInfoID     json.Number `json:"storageInfoId" csv:"storageInfoId"`
	ProductID         json.Number `json:"productId" csv:"productId"`
	ProductCode       string      `json:"productCode" csv:"productCode"`
	ProductName       string      `json:"productName" csv:"productName"`
	Size              string      `json:"size" csv:"size"`
	Color             string      `json:"color" csv:"color"`
	GroupCode         string      `json:"groupCode" csv:"groupCode"`
	SupplierProductNo *string     `json:"supplierProductNo" csv:"supplierProductNo"`
	Cost              json.Number `json:"cost" csv:"cost"`
	Quantity          json.Number `json:"quantity" csv:"quantity"`
	Modified          string      `json:"modified" csv:"modified"`
}

type StorageInfoProductCSV struct {
	*CSVHandler
	buf []StorageInfoProduct
}

func NewStorageInfoProductCSV(bufSize int, output string) (*StorageInfoProductCSV, error) {
	buf := make([]StorageInfoProduct, bufSize)
	handler, err := NewCSVHandler([]StorageInfoProduct{}, output)
	if err != nil {
		return nil, err
	}

	return &StorageInfoProductCSV{
		handler,
		buf,
	}, nil
}

func (sip *StorageInfoProductCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &sip.buf[i])
	}
	sip.CSVWriter.Write(sip.buf[:len(resp.Result)])
	return sip.CSVWriter
}

type StorageInfoDeliveryProduct struct {
	StorageInfoID     json.Number `json:"storageInfoId" csv:"storageInfoId"`
	StoreID           json.Number `json:"storeId" csv:"storeId"`
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

type StorageInfoDeliveryProductCSV struct {
	*CSVHandler
	buf []StorageInfoDeliveryProduct
}

func NewStorageInfoDeliveryProductCSV(bufSize int, output string) (*StorageInfoDeliveryProductCSV, error) {
	buf := make([]StorageInfoDeliveryProduct, bufSize)
	handler, err := NewCSVHandler([]StorageInfoDeliveryProduct{}, output)
	if err != nil {
		return nil, err
	}

	return &StorageInfoDeliveryProductCSV{
		handler,
		buf,
	}, nil
}

func (sidp *StorageInfoDeliveryProductCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &sidp.buf[i])
	}
	sidp.CSVWriter.Write(sidp.buf[:len(resp.Result)])
	return sidp.CSVWriter
}
