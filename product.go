package main

import (
	"github.com/shopspring/decimal"
)

type Product struct {
	ProductID            int             `json:"productId" csv:"productID"`
	CategoryID           int             `json:"categoryId" csv:"categoryID"`
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

type ProductPrice struct {
	ProductID     int             `json:"productId" csv:"product_id"`
	StoreID       string          `json:"storeId" csv:"store_id"`
	PriceDivision string          `json:"priceDivision" csv:"price_division"`
	StartDate     string          `json:"startDate" csv:"start_date"`
	EndDate       string          `json:"endDate" csv:"end_date"`
	Price         decimal.Decimal `json:"price" csv:"price"`
	InsDateTime   string          `json:"insDateTime" csv:"ins_date_time"`
	UpdDateTime   string          `json:"updDateTime" csv:"upd_date_time"`
}

type ProductReseveItem struct {
	ProductID string `json:"productId" csv:"product_id"`
	No        string `json:"no" csv:"no"`
	Value     string `json:"value" csv:"value"`
}

type ProductReseveItemLabel struct {
	No    string `json:"no" csv:"no"`
	Label string `json:"label" csv:"label"`
}

type ProductStore struct {
	ProductID      int    `json:"productId" csv:"product_id"`
	StoreID        int    `json:"storeId" csv:"store_id"`
	AssignDivision string `json:"assignDivision" csv:"assign_division"`
}

type ProductInventoryReservation struct {
	ProductID            int    `json:"productId" csv:"product_id"`
	ReservationProductID string `json:"reservationProductId" csv:"reservation_product_id"`
	ReservationAmount    int    `json:"reservationAmount" csv:"reservation_amount"`
}
