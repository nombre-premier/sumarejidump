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

type ProductParquetSchema struct {
	ProductID              int64   `json:",string" parquet:"name=product_id, type=INT64"`
	CategoryID             int64   `json:",string" parquet:"name=category_id, type=INT64"`
	ProductCode            string  `parquet:"name=product_code, type=BYTE_ARRAY, convertedtype=UTF8"`
	ProductName            string  `parquet:"name=product_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	ProductKana            *string `parquet:"name=product_kana, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	TaxDivision            string  `parquet:"name=tax_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	ProductPriceDivision   string  `parquet:"name=product_price_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	Price                  int64   `json:",string" parquet:"name=price, type=INT64"`
	CustomerPrice          *int64  `json:",string" parquet:"name=customer_price, type=INT64, repetitiontype=OPTIONAL"`
	Cost                   int64   `json:",string" parquet:"name=cost, type=INT64"`
	Attribute              *string `parquet:"name=attribute, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Description            *string `parquet:"name=description, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	CatchCopy              *string `parquet:"name=catch_copy, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Size                   string  `parquet:"name=size, type=BYTE_ARRAY, convertedtype=UTF8"`
	Color                  string  `parquet:"name=color, type=BYTE_ARRAY, convertedtype=UTF8"`
	Tag                    *string `parquet:"name=tag, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	GroupCode              *string `parquet:"name=group_code, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	URL                    *string `parquet:"name=url, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	DisplaySequence        *int64  `json:",string" parquet:"name=display_sequence, type=INT64, repetitiontype=OPTIONAL"`
	SalesDivision          string  `parquet:"name=sales_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	StockControlDivision   string  `parquet:"name=stock_control_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	DisplayFlag            string  `parquet:"name=display_flag, type=BYTE_ARRAY, convertedtype=UTF8"`
	Division               string  `parquet:"name=division, type=BYTE_ARRAY, convertedtype=UTF8"`
	ProductOptionGroupID   *int64  `json:",string" parquet:"name=product_option_group_id, type=INT64, repetitiontype=OPTIONAL"`
	PointNotApplicable     string  `parquet:"name=point_not_applicable, type=BYTE_ARRAY, convertedtype=UTF8"`
	TaxFreeDivision        string  `parquet:"name=tax_free_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	SupplierProductNo      *string `parquet:"name=supplier_product_no, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	StaffDiscountRate      *int64  `json:",string" parquet:"name=staff_discount_rate, type=INT64, repetitiontype=OPTIONAL"`
	UseCategoryReduceTax   int64   `json:",string" parquet:"name=use_category_reduce_tax, type=INT64"`
	ReduceTaxId            *int64  `json:",string" parquet:"name=reduce_tax_id, type=INT64, repetitiontype=OPTIONAL"`
	ReduceTaxPrice         int64   `json:",string" parquet:"name=reduce_tax_price, type=INT64"`
	ReduceTaxCustomerPrice int64   `json:",string" parquet:"name=reduce_tax_customer_price, type=INT64"`
	AppStartDateTime       *string `parquet:"name=app_start_date_time, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	InsDateTime            string  `parquet:"name=ins_date_time, type=BYTE_ARRAY, convertedtype=UTF8"`
	UpdDateTime            string  `parquet:"name=upd_date_time, type=BYTE_ARRAY, convertedtype=UTF8"`
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

type ProductPriceParquetSchema struct {
	ProductID     int64  `json:",string" parquet:"name=product_id, type=INT64"`
	StoreID       string `parquet:"name=store_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	PriceDivision string `parquet:"name=price_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	StartDate     string `parquet:"name=start_date, type=BYTE_ARRAY, convertedtype=UTF8"`
	EndDate       string `parquet:"name=end_date, type=BYTE_ARRAY, convertedtype=UTF8"`
	Price         int64  `json:",string" parquet:"name=price, type=INT64"`
	InsDateTime   string `parquet:"name=ins_date_time, type=BYTE_ARRAY, convertedtype=UTF8"`
	UpdDateTime   string `parquet:"name=upd_date_time, type=BYTE_ARRAY, convertedtype=UTF8"`
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

type ProductReserveItemParquetSchema struct {
	ProductID int64  `json:",string" parquet:"name=product_id, type=INT64"`
	No        int64  `json:",string" parquet:"name=no, type=INT64"`
	Value     string `parquet:"name=value, type=BYTE_ARRAY, convertedtype=UTF8"`
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

type ProductReserveItemLabelParquetSchema struct {
	No    int64  `json:",string" parquet:"name=no, type=INT64"`
	Label string `parquet:"name=label, type=BYTE_ARRAY, convertedtype=UTF8"`
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

type ProductStoreParquetSchema struct {
	ProductID      int64  `json:",string" parquet:"name=product_id, type=INT64"`
	StoreID        int64  `json:",string" parquet:"name=store_id, type=INT64"`
	AssignDivision string `parquet:"name=assign_division, type=BYTE_ARRAY, convertedtype=UTF8"`
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

type ProductInventoryReservationParquetSchema struct {
	ProductID            int64  `json:",string" parquet:"name=product_id, type=INT64"`
	ReservationProductID string `parquet:"name=reservation_product_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	ReservationAmount    int64  `json:",string" parquet:"name=reservation_amount, type=INT64"`
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
