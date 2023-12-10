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

type StorageParquetSchema struct {
	StorageID               int64   `json:",string" parquet:"name=storage_id, type=INT64"`
	StorageInfoID           int64   `json:",string" parquet:"name=storage_info_id, type=INT64"`
	SupplierID              int64   `json:",string" parquet:"name=supplier_id, type=INT64"`
	StorageStoreID          int64   `json:",string" parquet:"name=storage_store_id, type=INT64"`
	StorageExpectedDateFrom string  `parquet:"name=storage_expected_date_from, type=BYTE_ARRAY, convertedtype=UTF8"`
	StorageExpectedDateTo   string  `parquet:"name=storage_expected_date_to, type=BYTE_ARRAY, convertedtype=UTF8"`
	StorageDate             *string `parquet:"name=storage_date, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Memo                    string  `parquet:"name=memo, type=BYTE_ARRAY, convertedtype=UTF8"`
	Status                  string  `parquet:"name=status, type=BYTE_ARRAY, convertedtype=UTF8"`
	IdentificationNo        *string `parquet:"name=identification_no, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Modified                string  `parquet:"name=modified, type=BYTE_ARRAY, convertedtype=UTF8"`
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

type StorageDetailParquetSchema struct {
	StorageID              int64   `json:",string" parquet:"name=storage_id, type=INT64"`
	ProductID              int64   `json:",string" parquet:"name=product_id, type=INT64"`
	ProductCode            string  `parquet:"name=product_code, type=BYTE_ARRAY, convertedtype=UTF8"`
	ProductName            string  `parquet:"name=product_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	Size                   string  `parquet:"name=size, type=BYTE_ARRAY, convertedtype=UTF8"`
	Color                  string  `parquet:"name=color, type=BYTE_ARRAY, convertedtype=UTF8"`
	GroupCode              string  `parquet:"name=group_code, type=BYTE_ARRAY, convertedtype=UTF8"`
	SupplierProductNo      *string `parquet:"name=supplier_product_no, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Cost                   *int64  `json:",string" parquet:"name=cost, type=INT64, repetitiontype=OPTIONAL"`
	ScheduledQuantity      int64   `json:",string" parquet:"name=scheduled_quantity, type=INT64"`
	InspectionQuantity     int64   `json:",string" parquet:"name=inspection_quantity, type=INT64"`
	StockoutQuantity       int64   `json:",string" parquet:"name=stockout_quantity, type=INT64"`
	StockoutReason         string  `parquet:"name=stockout_reason, type=BYTE_ARRAY, convertedtype=UTF8"`
	InspectionDate         string  `parquet:"name=inspection_date, type=BYTE_ARRAY, convertedtype=UTF8"`
	CompulsoryCompleteFlag string  `parquet:"name=compulsory_complete_flag, type=BYTE_ARRAY, convertedtype=UTF8"`
	Status                 string  `parquet:"name=status, type=BYTE_ARRAY, convertedtype=UTF8"`
	Modified               string  `parquet:"name=modified, type=BYTE_ARRAY, convertedtype=UTF8"`
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
