package main

import "encoding/json"

type Store struct {
	StoreID                  json.Number `json:"storeId" csv:"storeId"`
	StoreCode                string      `json:"storeCode" csv:"storeCode"`
	StoreName                string      `json:"storeName" csv:"storeName"`
	StoreAbbr                *string     `json:"storeAbbr" csv:"storeAbbr"`
	Division                 string      `json:"division" csv:"division"`
	PostCode                 *string     `json:"postCode" csv:"postCode"`
	Address                  *string     `json:"address" csv:"address"`
	PhoneNumber              *string     `json:"phoneNumber" csv:"phoneNumber"`
	FaxNumber                *string     `json:"faxNumber" csv:"faxNumber"`
	MailAddress              *string     `json:"mailAddress" csv:"mailAddress"`
	Homepage                 *string     `json:"homepage" csv:"homepage"`
	TempTranMailAddress      *string     `json:"tempTranMailAddress" csv:"tempTranMailAddress"`
	PriceChangeFlag          string      `json:"priceChangeFlag" csv:"priceChangeFlag"`
	SellDivision             string      `json:"sellDivision" csv:"sellDivision"`
	SumProcDivision          string      `json:"sumProcDivision" csv:"sumProcDivision"`
	SumDateChangeTime        *string     `json:"sumDateChangeTime" csv:"sumDateChangeTime"`
	SumRefColumn             string      `json:"sumRefColumn" csv:"sumRefColumn"`
	PointNotApplicable       string      `json:"pointNotApplicable" csv:"pointNotApplicable"`
	TaxFreeDivision          string      `json:"taxFreeDivision" csv:"taxFreeDivision"`
	MaxBundleProductCount    json.Number `json:"maxBundleProductCount" csv:"maxBundleProductCount"`
	RoundingDivision         string      `json:"roundingDivision" csv:"roundingDivision"`
	DiscountRoundingDivision string      `json:"discountRoundingDivision" csv:"discountRoundingDivision"`
	PauseFlag                json.Number `json:"pauseFlag" csv:"pauseFlag"`
	PointUseDivision         string      `json:"pointUseDivision" csv:"pointUseDivision"`
	InsDateTime              string      `json:"insDateTime" csv:"insDateTime"`
	UpdDateTime              string      `json:"updDateTime" csv:"updDateTime"`
}

type StoreParquetSchema struct {
	StoreID                  int64   `parquet:"name=store_id, type=INT64"`
	StoreCode                string  `parquet:"name=store_code, type=BYTE_ARRAY, convertedtype=UTF8"`
	StoreName                string  `parquet:"name=store_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	StoreAbbr                *string `parquet:"name=store_abbr, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Division                 string  `parquet:"name=division, type=BYTE_ARRAY, convertedtype=UTF8"`
	PostCode                 *string `parquet:"name=post_code, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Address                  *string `parquet:"name=address, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	PhoneNumber              *string `parquet:"name=phone_number, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	FaxNumber                *string `parquet:"name=fax_number, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	MailAddress              *string `parquet:"name=mail_address, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Homepage                 *string `parquet:"name=homepage, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	TempTranMailAddress      *string `parquet:"name=temp_tran_mail_address, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	PriceChangeFlag          string  `parquet:"name=price_change_flag, type=BYTE_ARRAY, convertedtype=UTF8"`
	SellDivision             string  `parquet:"name=sell_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	SumProcDivision          string  `parquet:"name=sum_proc_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	SumDateChangeTime        *string `parquet:"name=sum_date_change_time, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	SumRefColumn             string  `parquet:"name=sum_ref_column, type=BYTE_ARRAY, convertedtype=UTF8"`
	PointNotApplicable       string  `parquet:"name=point_not_applicable, type=BYTE_ARRAY, convertedtype=UTF8"`
	TaxFreeDivision          string  `parquet:"name=tax_free_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	MaxBundleProductCount    int64   `parquet:"name=max_bundle_product_count, type=INT64"`
	RoundingDivision         string  `parquet:"name=rounding_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	DiscountRoundingDivision string  `parquet:"name=discount_rounding_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	PauseFlag                int64   `parquet:"name=pause_flag, type=INT64"`
	PointUseDivision         string  `parquet:"name=point_use_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	InsDateTime              string  `parquet:"name=ins_date_time, type=BYTE_ARRAY, convertedtype=UTF8"`
	UpdDateTime              string  `parquet:"name=upd_date_time, type=BYTE_ARRAY, convertedtype=UTF8"`
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
