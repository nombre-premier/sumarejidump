package main

type Store struct {
	StoreID                  int     `json:"storeId" csv:"storeID"`
	StoreCode                string  `json:"storeCode" csv:"storeCode"`
	StoreName                string  `json:"storeName" csv:"storeName"`
	StoreAbbr                *string `json:"storeAbbr" csv:"storeAbbr"`
	Division                 string  `json:"division" csv:"division"`
	PostCode                 *string `json:"postCode" csv:"postCode"`
	Address                  *string `json:"address" csv:"address"`
	PhoneNumber              *string `json:"phoneNumber" csv:"phoneNumber"`
	FaxNumber                *string `json:"faxNumber" csv:"faxNumber"`
	MailAddress              *string `json:"mailAddress" csv:"mailAddress"`
	Homepage                 *string `json:"homepage" csv:"homepage"`
	TempTranMailAddress      *string `json:"tempTranMailAddress" csv:"tempTranMailAddress"`
	PriceChangeFlag          string  `json:"priceChangeFlag" csv:"priceChangeFlag"`
	SellDivision             string  `json:"sellDivision" csv:"sellDivision"`
	SumProcDivision          string  `json:"sumProcDivision" csv:"sumProcDivision"`
	SumDateChangeTime        string  `json:"sumDateChangeTime" csv:"sumDateChangeTime"`
	SumRefColumn             string  `json:"sumRefColumn" csv:"sumRefColumn"`
	PointNotApplicable       string  `json:"pointNotApplicable" csv:"pointNotApplicable"`
	TaxFreeDivision          string  `json:"taxFreeDivision" csv:"taxFreeDivision"`
	MaxBundleProductCount    int     `json:"maxBundleProductCount" csv:"maxBundleProductCount"`
	RoundingDivision         string  `json:"roundingDivision" csv:"roundingDivision"`
	DiscountRoundingDivision string  `json:"discountRoundingDivision" csv:"discountRoundingDivision"`
	PauseFlag                int     `json:"pauseFlag" csv:"pauseFlag"`
	PointUseDivision         string  `json:"pointUseDivision" csv:"pointUseDivision"`
	InsDateTime              string  `json:"insDateTime" csv:"insDateTime"`
	UpdDateTime              string  `json:"updDateTime" csv:"updDateTime"`
}
