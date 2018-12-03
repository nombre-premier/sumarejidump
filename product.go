package main

type Product struct {
	ProductID            string      `json:"productId"`
	CategoryID           string      `json:"categoryId"`
	ProductCode          string      `json:"productCode"`
	ProductName          string      `json:"productName"`
	ProductKana          string      `json:"productKana"`
	TaxDivision          string      `json:"taxDivision"`
	ProductPriceDivision string      `json:"productPriceDivision"`
	Price                string      `json:"price"`
	CustomerPrice        interface{} `json:"customerPrice"`
	Cost                 interface{} `json:"cost"`
	Attribute            interface{} `json:"attribute"`
	Description          string      `json:"description"`
	CatchCopy            interface{} `json:"catchCopy"`
	Size                 string      `json:"size"`
	Color                string      `json:"color"`
	Tag                  interface{} `json:"tag"`
	GroupCode            string      `json:"groupCode"`
	URL                  interface{} `json:"url"`
	DisplaySequence      interface{} `json:"displaySequence"`
	DisplayFlag          string      `json:"displayFlag"`
	SalesDivision        string      `json:"salesDivision"`
	Division             string      `json:"division"`
	ProductOptionGroupID interface{} `json:"productOptionGroupId"`
	StockControlDivision string      `json:"stockControlDivision"`
	PointNotApplicable   string      `json:"pointNotApplicable"`
	TaxFreeDivision      string      `json:"taxFreeDivision"`
	SupplierProductNo    interface{} `json:"supplierProductNo"`
	StaffDiscountRate    interface{} `json:"staffDiscountRate"`
	AppStartDateTime     interface{} `json:"appStartDateTime"`
	InsDateTime          string      `json:"insDateTime"`
	UpdDateTime          string      `json:"updDateTime"`
}
