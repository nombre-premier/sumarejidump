package main

import (
	"encoding/json"
)

type Product struct {
	ProductID              json.Number  `json:"productId" csv:"productId"`
	CategoryID             json.Number  `json:"categoryId" csv:"categoryId"`
	ProductCode            string       `json:"productCode" csv:"productCode"`
	ProductName            string       `json:"productName" csv:"productName"`
	ProductKana            *string      `json:"productKana" csv:"productKana"`
	TaxDivision            string       `json:"taxDivision" csv:"taxDivision"`
	ProductPriceDivision   string       `json:"productPriceDivision" csv:"productPriceDivision"`
	Price                  json.Number  `json:"price" csv:"price"`
	CustomerPrice          *json.Number `json:"customerPrice" csv:"customerPrice"`
	Cost                   json.Number  `json:"cost" csv:"cost"`
	Attribute              *string      `json:"attribute" csv:"attribute"`
	Description            *string      `json:"description" csv:"description"`
	CatchCopy              *string      `json:"catchCopy" csv:"catchCopy"`
	Size                   string       `json:"size" csv:"size"`
	Color                  string       `json:"color" csv:"color"`
	Tag                    *string      `json:"tag" csv:"tag"`
	GroupCode              *string      `json:"groupCode" csv:"groupCode"`
	URL                    *string      `json:"url" csv:"url"`
	DisplaySequence        *json.Number `json:"displaySequence" csv:"displaySequence"`
	SalesDivision          string       `json:"salesDivision" csv:"salesDivision"`
	StockControlDivision   string       `json:"stockControlDivision" csv:"stockControlDivision"`
	DisplayFlag            string       `json:"displayFlag" csv:"displayFlag"`
	Division               string       `json:"division" csv:"division"`
	ProductOptionGroupID   *json.Number `json:"productOptionGroupId" csv:"productOptionGroupId"`
	PointNotApplicable     string       `json:"pointNotApplicable" csv:"pointNotApplicable"`
	TaxFreeDivision        string       `json:"taxFreeDivision" csv:"taxFreeDivision"`
	SupplierProductNo      *string      `json:"supplierProductNo" csv:"supplierProductNo"`
	StaffDiscountRate      *json.Number `json:"staffDiscountRate" csv:"staffDiscountRate"`
	UseCategoryReduceTax   json.Number  `json:"useCategoryReduceTax" csv:"useCategoryReduceTax"`
	ReduceTaxId            *json.Number `json:"reduceTaxId" csv:"reduceTaxId"`
	ReduceTaxPrice         json.Number  `json:"reduceTaxPrice" csv:"reduceTaxPrice"`
	ReduceTaxCustomerPrice json.Number  `json:"reduceTaxCustomerPrice" csv:"reduceTaxCustomerPrice"`
	AppStartDateTime       *string      `json:"appStartDateTime" csv:"appStartDateTime"`
	InsDateTime            string       `json:"insDateTime" csv:"insDateTime"`
	UpdDateTime            string       `json:"updDateTime" csv:"updDateTime"`
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
	ProductID     json.Number `json:"productId" csv:"productId"`
	StoreID       string      `json:"storeId" csv:"storeId"`
	PriceDivision string      `json:"priceDivision" csv:"priceDivision"`
	StartDate     string      `json:"startDate" csv:"startDate"`
	EndDate       string      `json:"endDate" csv:"endDate"`
	Price         json.Number `json:"price" csv:"price"`
	InsDateTime   string      `json:"insDateTime" csv:"insDateTime"`
	UpdDateTime   string      `json:"updDateTime" csv:"updDateTime"`
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

type ProductReserveItem struct {
	ProductID json.Number `json:"productId" csv:"productId"`
	No        json.Number `json:"no" csv:"no"`
	Value     string      `json:"value" csv:"value"`
}

type ProductReserveItemCSV struct {
	*CSVHandler
	buf []ProductReserveItem
}

func NewProductReserveItemCSV(bufSize int, output string) (*ProductReserveItemCSV, error) {
	buf := make([]ProductReserveItem, bufSize)
	handler, err := NewCSVHandler([]ProductReserveItem{}, output)
	if err != nil {
		return nil, err
	}

	return &ProductReserveItemCSV{
		handler,
		buf,
	}, nil
}

func (pric *ProductReserveItemCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &pric.buf[i])
	}
	pric.CSVWriter.Write(pric.buf[:len(resp.Result)])
	return pric.CSVWriter
}

type ProductReserveItemLabel struct {
	No    json.Number `json:"no" csv:"no"`
	Label string      `json:"label" csv:"label"`
}

type ProductReserveItemLabelCSV struct {
	*CSVHandler
	buf []ProductReserveItemLabel
}

func NewProductReserveItemLabelCSV(bufSize int, output string) (*ProductReserveItemLabelCSV, error) {
	buf := make([]ProductReserveItemLabel, bufSize)
	handler, err := NewCSVHandler([]ProductReserveItemLabel{}, output)
	if err != nil {
		return nil, err
	}

	return &ProductReserveItemLabelCSV{
		handler,
		buf,
	}, nil
}

func (prilc *ProductReserveItemLabelCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &prilc.buf[i])
	}
	prilc.CSVWriter.Write(prilc.buf[:len(resp.Result)])
	return prilc.CSVWriter
}

type ProductStore struct {
	ProductID      json.Number `json:"productId" csv:"productId"`
	StoreID        json.Number `json:"storeId" csv:"storeId"`
	AssignDivision string      `json:"assignDivision" csv:"assignDivision"`
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
	ProductID            json.Number `json:"productId" csv:"productId"`
	ReservationProductID string      `json:"reservationProductId" csv:"reservationProductId"`
	ReservationAmount    json.Number `json:"reservationAmount" csv:"reservationAmount"`
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
