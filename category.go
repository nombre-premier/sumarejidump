package main

import (
	"encoding/json"
)

type Category struct {
	CategoryID         json.Number  `json:"categoryId" csv:"categoryId"`
	CategoryCode       string       `json:"categoryCode" csv:"categoryCode"`
	CategoryName       string       `json:"categoryName" csv:"categoryName"`
	CategoryAbbr       string       `json:"categoryAbbr" csv:"categoryAbbr"`
	CategoryGroupID    *json.Number `json:"categoryGroupId" csv:"categoryGroupId"`
	ParentCategoryID   *json.Number `json:"parentCategoryId" csv:"parentCategoryId"`
	Level              json.Number  `json:"level" csv:"level"`
	DisplaySequence    json.Number  `json:"displaySequence" csv:"displaySequence"`
	DisplayFlag        string       `json:"displayFlag" csv:"displayFlag"`
	PointNotApplicable string       `json:"pointNotApplicable" csv:"pointNotApplicable"`
	TaxFreeDivision    string       `json:"taxFreeDivision" csv:"taxFreeDivision"`
	Color              *string      `json:"color" csv:"color"`
	Tag                *string      `json:"tag" csv:"tag"`
	InsDateTime        string       `json:"insDateTime" csv:"insDateTime"`
	UpdDateTime        string       `json:"updDateTime" csv:"updDateTime"`
}

type CategoryParquetSchema struct {
	CategoryID         int64   `parquet:"name=category_id, type=INT64"`
	CategoryCode       string  `parquet:"name=category_code, type=BYTE_ARRAY, convertedtype=UTF8"`
	CategoryName       string  `parquet:"name=category_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	CategoryAbbr       string  `parquet:"name=category_abbr, type=BYTE_ARRAY, convertedtype=UTF8"`
	CategoryGroupID    *int64  `parquet:"name=category_group_id, type=INT64, repetitiontype=OPTIONAL"`
	ParentCategoryID   *int64  `parquet:"name=parent_category_id, type=INT64, repetitiontype=OPTIONAL"`
	Level              int32   `parquet:"name=level, type=INT32"`
	DisplaySequence    int32   `parquet:"name=display_sequence, type=INT32"`
	DisplayFlag        string  `parquet:"name=display_flag, type=BYTE_ARRAY, convertedtype=UTF8"`
	PointNotApplicable string  `parquet:"name=point_not_applicable, type=BYTE_ARRAY, convertedtype=UTF8"`
	TaxFreeDivision    string  `parquet:"name=tax_free_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	Color              *string `parquet:"name=color, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Tag                *string `parquet:"name=tag, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	InsDateTime        string  `parquet:"name=ins_date_time, type=BYTE_ARRAY, convertedtype=UTF8"`
	UpdDateTime        string  `parquet:"name=upd_date_time, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type CategoryCSV struct {
	*CSVHandler
	buf []Category
}

func NewCategoryCSV(bufSize int, output string) (*CategoryCSV, error) {
	buf := make([]Category, bufSize)
	handler, err := NewCSVHandler([]Category{}, output)
	if err != nil {
		return nil, err
	}

	return &CategoryCSV{
		handler,
		buf,
	}, nil
}

func (cc *CategoryCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &cc.buf[i])
	}
	cc.CSVWriter.Write(cc.buf[:len(resp.Result)])
	return cc.CSVWriter
}
