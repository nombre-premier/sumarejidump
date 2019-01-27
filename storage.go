package main

import (
	"encoding/json"
)

type Storage struct {
	StorageID               json.Number `json:"storageId" csv:"storageId"`
	StorageInfoID           json.Number `json:"storageInfoId" csv:"storageInfoId"`
	SupplierID              json.Number `json:"supplierId"  csv:"supplierId"`
	StorageStoreID          json.Number `json:"storageStoreId" csv:"storageStoreId"`
	StorageExpectedDateFrom string      `json:"storageExpectedDateFrom"  csv:"storageExpectedDateFrom"`
	StorageExpectedDateTo   string      `json:"storageExpectedDateTo"  csv:"storageExpectedDateTo"`
	StorageDate             *string     `json:"storageDate"   csv:"storageDate"`
	Memo                    string      `json:"memo"   csv:"memo"`
	Status                  string      `json:"status" csv:"status"`
	IdentificationNo        *string     `json:"identificationNo" csv:"identificationNo"`
	Modified                string      `json:"modified" csv:"modified"`
}

type StorageCSV struct {
	*CSVHandler
	buf []Storage
}

func NewStorageCSV(bufSize int, output string) (*StorageCSV, error) {
	buf := make([]Storage, bufSize)
	handler, err := NewCSVHandler([]Storage{}, output)
	if err != nil {
		return nil, err
	}

	return &StorageCSV{
		handler,
		buf,
	}, nil
}

func (st *StorageCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &st.buf[i])
	}
	st.CSVWriter.Write(st.buf[:len(resp.Result)])
	return st.CSVWriter
}

type StorageDetail struct {
	StorageID              json.Number  `json:"storageId"  csv:"storageId"`
	ProductID              json.Number  `json:"productId"  csv:"productId"`
	ProductCode            string       `json:"productCode" csv:"productCode"`
	ProductName            string       `json:"productName" csv:"productName"`
	Size                   string       `json:"size" csv:"size"`
	Color                  string       `json:"color"  csv:"color"`
	GroupCode              string       `json:"groupCode" csv:"groupCode"`
	SupplierProductNo      *string      `json:"supplierProductNo" csv:"supplierProductNo"`
	Cost                   *json.Number `json:"cost" csv:"cost"`
	ScheduledQuantity      json.Number  `json:"scheduledQuantity" csv:"scheduledQuantity"`
	InspectionQuantity     json.Number  `json:"inspectionQuantity" csv:"inspectionQuantity"`
	StockoutQuantity       json.Number  `json:"stockoutQuantity" csv:"stockoutQuantity"`
	StockoutReason         string       `json:"stockoutReason" csv:"stockoutReason"`
	InspectionDate         string       `json:"inspectionDate" csv:"inspectionDate"`
	CompulsoryCompleteFlag string       `json:"compulsoryCompleteFlag" csv:"compulsoryCompleteFlag"`
	Status                 string       `json:"status" csv:"status"`
	Modified               string       `json:"modified" csv:"modified"`
}

type StorageDetailCSV struct {
	*CSVHandler
	buf []StorageDetail
}

func NewStorageDetailCSV(bufSize int, output string) (*StorageDetailCSV, error) {
	buf := make([]StorageDetail, bufSize)
	handler, err := NewCSVHandler([]StorageDetail{}, output)
	if err != nil {
		return nil, err
	}

	return &StorageDetailCSV{
		handler,
		buf,
	}, nil
}

func (sds *StorageDetailCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &sds.buf[i])
	}
	sds.CSVWriter.Write(sds.buf[:len(resp.Result)])
	return sds.CSVWriter
}
