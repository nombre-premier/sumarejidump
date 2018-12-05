package main

import "encoding/json"

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

type StoreCSV struct {
	*CSVHandler
	buf []Store
}

func NewStoreCSV(bufSize int, output string) (*StoreCSV, error) {
	buf := make([]Store, bufSize)
	handler, err := NewCSVHandler([]Store{}, output)
	if err != nil {
		return nil, err
	}

	return &StoreCSV{
		handler,
		buf,
	}, nil
}

func (sc *StoreCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &sc.buf[i])
	}
	sc.CSVWriter.Write(sc.buf[:len(resp.Result)])
	return sc.CSVWriter
}
