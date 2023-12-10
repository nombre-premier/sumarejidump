package main

import "encoding/json"

type StocktakingInfo struct {
	StocktakingInfoID       json.Number `json:"stocktakingInfoId" csv:"stocktakingInfoId"`
	StoreID                 json.Number `json:"storeId" csv:"storeId"`
	DivisionUnit            string      `json:"divisionUnit" csv:"divisionUnit"`
	TargetDate              string      `json:"targetDate" csv:"targetDate"`
	AdjustmentDate          string      `json:"adjustmentDate" csv:"adjustmentDate"`
	StocktakingCompleteDate string      `json:"stocktakingCompleteDate" csv:"stocktakingCompleteDate"`
	Status                  string      `json:"status" csv:"status"`
	Created                 string      `json:"created" csv:"created"`
	Modified                string      `json:"modified" csv:"modified"`
}

type StocktakingInfoParquetSchema struct {
	StocktakingInfoID       int64  `json:",string" parquet:"name=stocktaking_info_id, type=INT64"`
	StoreID                 int64  `json:",string" parquet:"name=store_id, type=INT64"`
	DivisionUnit            string `parquet:"name=division_unit, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	TargetDate              string `parquet:"name=target_date, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	AdjustmentDate          string `parquet:"name=adjustment_date, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	StocktakingCompleteDate string `parquet:"name=stocktaking_complete_date, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Status                  string `parquet:"name=status, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Created                 string `parquet:"name=created, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Modified                string `parquet:"name=modified, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
}

type StocktakingInfoCSV struct {
	*CSVHandler
	buf []StocktakingInfo
}

func NewStocktakingInfoCSV(bufSize int, output string) (*StocktakingInfoCSV, error) {
	buf := make([]StocktakingInfo, bufSize)
	handler, err := NewCSVHandler([]StocktakingInfo{}, output)
	if err != nil {
		return nil, err
	}

	return &StocktakingInfoCSV{
		handler,
		buf,
	}, nil
}

func (cc *StocktakingInfoCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &cc.buf[i])
	}
	cc.CSVWriter.Write(cc.buf[:len(resp.Result)])
	return cc.CSVWriter
}

type StocktakingHead struct {
	StocktakingInfoID json.Number `json:"stocktakingInfoId" csv:"stocktakingInfoId"`
	StocktakingHeadID json.Number `json:"stocktakingHeadId" csv:"stocktakingHeadId"`
	DivisionCode      string      `json:"divisionCode" csv:"divisionCode"`
	Modified          string      `json:"modified" csv:"modified"`
}

type StocktakingHeadParquetSchema struct {
	StocktakingInfoID int64  `json:",string" parquet:"name=stocktaking_info_id, type=INT64"`
	StocktakingHeadID int64  `json:",string" parquet:"name=stocktaking_head_id, type=INT64"`
	DivisionCode      string `parquet:"name=division_code, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Modified          string `parquet:"name=modified, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
}

type StocktakingHeadCSV struct {
	*CSVHandler
	buf []StocktakingHead
}

func NewStocktakingHeadCSV(bufSize int, output string) (*StocktakingHeadCSV, error) {
	buf := make([]StocktakingHead, bufSize)
	handler, err := NewCSVHandler([]StocktakingHead{}, output)
	if err != nil {
		return nil, err
	}

	return &StocktakingHeadCSV{
		handler,
		buf,
	}, nil
}

func (cc *StocktakingHeadCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &cc.buf[i])
	}
	cc.CSVWriter.Write(cc.buf[:len(resp.Result)])
	return cc.CSVWriter
}

type StocktakingDetail struct {
	StocktakingInfoID             json.Number `json:"stocktakingInfoId" csv:"stocktakingInfoId"`
	StocktakingHeadID             json.Number `json:"stocktakingHeadId" csv:"stocktakingHeadId"`
	ProductID                     json.Number `json:"productId" csv:"productId"`
	ProductCode                   string      `json:"productCode" csv:"productCode"`
	ProductName                   string      `json:"productName" csv:"productName"`
	Size                          string      `json:"size" csv:"size"`
	Color                         string      `json:"color" csv:"color"`
	GroupCode                     *string     `json:"groupCode" csv:"groupCode"`
	SupplierProductNo             *string     `json:"supplierProductNo" csv:"supplierProductNo"`
	StocktakingQuantity           json.Number `json:"stocktakingQuantity" csv:"stocktakingQuantity"`
	TransportationStockQuantity   json.Number `json:"transportationStockQuantity" csv:"transportationStockQuantity"`
	LayawayStockQuantity          json.Number `json:"layawayStockQuantity" csv:"layawayStockQuantity"`
	StockQuantityBeforeAdjustment json.Number `json:"stockQuantityBeforeAdjustment" csv:"stockQuantityBeforeAdjustment"`
	Cost                          json.Number `json:"cost" csv:"cost"`
	Memo                          *string     `json:"memo" csv:"memo"`
	Modified                      string      `json:"modified" csv:"modified"`
}

type StocktakingDetailParquetSchema struct {
	StocktakingInfoID             int64  `json:",string" parquet:"name=stocktaking_info_id, type=INT64"`
	StocktakingHeadID             int64  `json:",string" parquet:"name=stocktaking_head_id, type=INT64"`
	ProductID                     int64  `json:",string" parquet:"name=product_id, type=INT64"`
	ProductCode                   string `parquet:"name=product_code, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	ProductName                   string `parquet:"name=product_name, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Size                          string `parquet:"name=size, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Color                         string `parquet:"name=color, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	GroupCode                     string `parquet:"name=group_code, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	SupplierProductNo             string `parquet:"name=supplier_product_no, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	StocktakingQuantity           int64  `json:",string" parquet:"name=stocktaking_quantity, type=INT64"`
	TransportationStockQuantity   int64  `json:",string" parquet:"name=transportation_stock_quantity, type=INT64"`
	LayawayStockQuantity          int64  `json:",string" parquet:"name=layaway_stock_quantity, type=INT64"`
	StockQuantityBeforeAdjustment int64  `json:",string" parquet:"name=stock_quantity_before_adjustment, type=INT64"`
	Cost                          int64  `json:",string" parquet:"name=cost, type=INT64"`
	Memo                          string `parquet:"name=memo, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Modified                      string `parquet:"name=modified, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
}

type StocktakingDetailCSV struct {
	*CSVHandler
	buf []StocktakingDetail
}

func NewStocktakingDetailCSV(bufSize int, output string) (*StocktakingDetailCSV, error) {
	buf := make([]StocktakingDetail, bufSize)
	handler, err := NewCSVHandler([]StocktakingDetail{}, output)
	if err != nil {
		return nil, err
	}

	return &StocktakingDetailCSV{
		handler,
		buf,
	}, nil
}

func (cc *StocktakingDetailCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &cc.buf[i])
	}
	cc.CSVWriter.Write(cc.buf[:len(resp.Result)])
	return cc.CSVWriter
}
