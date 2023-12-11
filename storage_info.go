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

type StorageInfoParquetSchema struct {
	StorageInfoID    int64   `json:",string" parquet:"name=storage_info_id, type=INT64"`
	RecipientOrderID int64   `json:",string" parquet:"name=recipient_order_id, type=INT64"`
	OrderedDate      string  `parquet:"name=ordered_date, type=BYTE_ARRAY, convertedtype=UTF8"`
	Status           string  `parquet:"name=status, type=BYTE_ARRAY, convertedtype=UTF8"`
	IdentificationNo *string `parquet:"name=identification_no, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Modified         string  `parquet:"name=modified, type=BYTE_ARRAY, convertedtype=UTF8"`
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

type StorageInfoDeliveryParquetSchema struct {
	StorageInfoID           int64  `json:",string" parquet:"name=storage_info_id, type=INT64"`
	StorageStoreID          int64  `json:",string" parquet:"name=storage_store_id, type=INT64"`
	StorageExpectedDateFrom string `parquet:"name=storage_expected_date_from, type=BYTE_ARRAY, convertedtype=UTF8"`
	StorageExpectedDateTo   string `parquet:"name=storage_expected_date_to, type=BYTE_ARRAY, convertedtype=UTF8"`
	Modified                string `parquet:"name=modified, type=BYTE_ARRAY, convertedtype=UTF8"`
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

type StorageInfoProductParquetSchema struct {
	StorageInfoID     int64   `json:",string" parquet:"name=storage_info_id, type=INT64"`
	ProductID         int64   `json:",string" parquet:"name=product_id, type=INT64"`
	ProductCode       string  `parquet:"name=product_code, type=BYTE_ARRAY, convertedtype=UTF8"`
	ProductName       string  `parquet:"name=product_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	Size              string  `parquet:"name=size, type=BYTE_ARRAY, convertedtype=UTF8"`
	Color             string  `parquet:"name=color, type=BYTE_ARRAY, convertedtype=UTF8"`
	GroupCode         string  `parquet:"name=group_code, type=BYTE_ARRAY, convertedtype=UTF8"`
	SupplierProductNo *string `parquet:"name=supplier_product_no, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Cost              int64   `json:",string" parquet:"name=cost, type=INT64"`
	Quantity          int64   `json:",string" parquet:"name=quantity, type=INT64"`
	Modified          string  `parquet:"name=modified, type=BYTE_ARRAY, convertedtype=UTF8"`
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

type StorageInfoDeliveryProductParquetSchema struct {
	StorageInfoID     int64   `json:",string" parquet:"name=storage_info_id, type=INT64"`
	StoreID           int64   `json:",string" parquet:"name=store_id, type=INT64"`
	ProductID         int64   `json:",string" parquet:"name=product_id, type=INT64"`
	ProductCode       string  `parquet:"name=product_code, type=BYTE_ARRAY, convertedtype=UTF8"`
	ProductName       string  `parquet:"name=product_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	Size              string  `parquet:"name=size, type=BYTE_ARRAY, convertedtype=UTF8"`
	Color             string  `parquet:"name=color, type=BYTE_ARRAY, convertedtype=UTF8"`
	GroupCode         string  `parquet:"name=group_code, type=BYTE_ARRAY, convertedtype=UTF8"`
	SupplierProductNo *string `parquet:"name=supplier_product_no, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Quantity          int64   `json:",string" parquet:"name=quantity, type=INT64"`
	Modified          string  `parquet:"name=modified, type=BYTE_ARRAY, convertedtype=UTF8"`
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
