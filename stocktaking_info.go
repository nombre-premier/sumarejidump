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
