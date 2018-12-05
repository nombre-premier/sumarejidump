package main

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type Product struct {
	ProductID            int             `json:"productId" csv:"productId"`
	CategoryID           int             `json:"categoryId" csv:"categoryId"`
	ProductCode          string          `json:"productCode" csv:"productCode"`
	ProductName          string          `json:"productName" csv:"productName"`
	ProductKana          string          `json:"productKana" csv:"productKana"`
	TaxDivision          string          `json:"taxDivision" csv:"taxDivision"`
	ProductPriceDivision string          `json:"productPriceDivision" csv:"productPriceDivision"`
	Price                decimal.Decimal `json:"price" csv:"price"`
	CustomerPrice        decimal.Decimal `json:"customerPrice" csv:"customerPrice"`
	Cost                 decimal.Decimal `json:"cost" csv:"cost"`
	Attribute            *string         `json:"attribute" csv:"attribute"`
	Description          string          `json:"description" csv:"description"`
	CatchCopy            *string         `json:"catchCopy" csv:"catchCopy"`
	Size                 string          `json:"size" csv:"size"`
	Color                string          `json:"color" csv:"color"`
	Tag                  *string         `json:"tag" csv:"tag"`
	GroupCode            string          `json:"groupCode" csv:"groupCode"`
	URL                  *string         `json:"url" csv:"url"`
	DisplaySequence      *int            `json:"displaySequence" csv:"displaySequence"`
	SalesDivision        string          `json:"salesDivision" csv:"salesDivision"`
	StockControlDivision string          `json:"stockControlDivision" csv:"stockControlDivision"`
	DisplayFlag          string          `json:"displayFlag" csv:"displayFlag"`
	Division             string          `json:"division" csv:"division"`
	ProductOptionGroupID *int            `json:"productOptionGroupId" csv:"productOptionGroupID"`
	PointNotApplicable   string          `json:"pointNotApplicable" csv:"pointNotApplicable"`
	TaxFreeDivision      string          `json:"taxFreeDivision" csv:"taxFreeDivision"`
	SupplierProductNo    *string         `json:"supplierProductNo" csv:"supplierProductNo"`
	StaffDiscountRate    *int            `json:"staffDiscountRate" csv:"staffDiscountRate"`
	AppStartDateTime     *string         `json:"appStartDateTime" csv:"appStartDateTime"`
	InsDateTime          string          `json:"insDateTime" csv:"insDateTime"`
	UpdDateTime          string          `json:"updDateTime" csv:"updDateTime"`
}

type ProductCSV struct {
	*CSVHandler
	buf []Product
}

func NewProductCSV(bufSize int, output string) (*ProductCSV, error) {
	buf := make([]Product, bufSize)
	handler, err := NewCSVHandler([]Product{}, output)
	if err != nil {
		return nil, err
	}

	return &ProductCSV{
		handler,
		buf,
	}, nil
}

func (pc *ProductCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &pc.buf[i])
	}
	pc.CSVWriter.Write(pc.buf[:len(resp.Result)])
	return pc.CSVWriter
}

type ProductPrice struct {
	ProductID     int             `json:"productId" csv:"productId"`
	StoreID       string          `json:"storeId" csv:"storeId"`
	PriceDivision string          `json:"priceDivision" csv:"priceDivision"`
	StartDate     string          `json:"startDate" csv:"startDate"`
	EndDate       string          `json:"endDate" csv:"endDate"`
	Price         decimal.Decimal `json:"price" csv:"price"`
	InsDateTime   string          `json:"insDateTime" csv:"insDateTime"`
	UpdDateTime   string          `json:"updDateTime" csv:"updDateTime"`
}

type ProductPriceCSV struct {
	*CSVHandler
	buf []ProductPrice
}

func NewProductPriceCSV(bufSize int, output string) (*ProductPriceCSV, error) {
	buf := make([]ProductPrice, bufSize)
	handler, err := NewCSVHandler([]ProductPrice{}, output)
	if err != nil {
		return nil, err
	}

	return &ProductPriceCSV{
		handler,
		buf,
	}, nil
}

func (ppc *ProductPriceCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &ppc.buf[i])
	}
	ppc.CSVWriter.Write(ppc.buf[:len(resp.Result)])
	return ppc.CSVWriter
}

type ProductReseveItem struct {
	ProductID string `json:"productId" csv:"productId"`
	No        string `json:"no" csv:"no"`
	Value     string `json:"value" csv:"value"`
}

type ProductReseveItemCSV struct {
	*CSVHandler
	buf []ProductReseveItem
}

func NewProductReseveItemCSV(bufSize int, output string) (*ProductReseveItemCSV, error) {
	buf := make([]ProductReseveItem, bufSize)
	handler, err := NewCSVHandler([]ProductReseveItem{}, output)
	if err != nil {
		return nil, err
	}

	return &ProductReseveItemCSV{
		handler,
		buf,
	}, nil
}

func (pric *ProductReseveItemCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &pric.buf[i])
	}
	pric.CSVWriter.Write(pric.buf[:len(resp.Result)])
	return pric.CSVWriter
}

type ProductReseveItemLabel struct {
	No    string `json:"no" csv:"no"`
	Label string `json:"label" csv:"label"`
}

type ProductReseveItemLabelCSV struct {
	*CSVHandler
	buf []ProductReseveItemLabel
}

func NewProductReseveItemLabelCSV(bufSize int, output string) (*ProductReseveItemLabelCSV, error) {
	buf := make([]ProductReseveItemLabel, bufSize)
	handler, err := NewCSVHandler([]ProductReseveItemLabel{}, output)
	if err != nil {
		return nil, err
	}

	return &ProductReseveItemLabelCSV{
		handler,
		buf,
	}, nil
}

func (prilc *ProductReseveItemLabelCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &prilc.buf[i])
	}
	prilc.CSVWriter.Write(prilc.buf[:len(resp.Result)])
	return prilc.CSVWriter
}

type ProductStore struct {
	ProductID      int    `json:"productId" csv:"productId"`
	StoreID        int    `json:"storeId" csv:"storeId"`
	AssignDivision string `json:"assignDivision" csv:"assignDivision"`
}

type ProductStoreCSV struct {
	*CSVHandler
	buf []ProductStore
}

func NewProductStoreCSV(bufSize int, output string) (*ProductStoreCSV, error) {
	buf := make([]ProductStore, bufSize)
	handler, err := NewCSVHandler([]ProductStore{}, output)
	if err != nil {
		return nil, err
	}

	return &ProductStoreCSV{
		handler,
		buf,
	}, nil
}

func (psc *ProductStoreCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &psc.buf[i])
	}
	psc.CSVWriter.Write(psc.buf[:len(resp.Result)])
	return psc.CSVWriter
}

type ProductInventoryReservation struct {
	ProductID            int    `json:"productId" csv:"productId"`
	ReservationProductID string `json:"reservationProductId" csv:"reservationProductId"`
	ReservationAmount    int    `json:"reservationAmount" csv:"reservationAmount"`
}

type ProductInventoryReservationCSV struct {
	*CSVHandler
	buf []ProductInventoryReservation
}

func NewProductInventoryReservationCSV(bufSize int, output string) (*ProductInventoryReservationCSV, error) {
	buf := make([]ProductInventoryReservation, bufSize)
	handler, err := NewCSVHandler([]ProductInventoryReservation{}, output)
	if err != nil {
		return nil, err
	}

	return &ProductInventoryReservationCSV{
		handler,
		buf,
	}, nil
}

func (pirc *ProductInventoryReservationCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &pirc.buf[i])
	}
	pirc.CSVWriter.Write(pirc.buf[:len(resp.Result)])
	return pirc.CSVWriter
}
