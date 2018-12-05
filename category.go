package main

import "encoding/json"

type Category struct {
	CategoryID         int     `json:"categoryId" csv:"categoryId"`
	CategoryCode       string  `json:"categoryCode" csv:"categoryCode"`
	CategoryName       string  `json:"categoryName" csv:"categoryName"`
	CategoryAbbr       string  `json:"categoryAbbr" csv:"categoryAbbr"`
	CategoryGroupID    *int    `json:"categoryGroupId" csv:"categoryGroupID"`
	ParentCategoryID   *int    `json:"parentCategoryId" csv:"parentCategoryID"`
	Level              int     `json:"level" csv:"level"`
	DisplaySequence    int     `json:"displaySequence" csv:"displaySequence"`
	DisplayFlag        string  `json:"displayFlag" csv:"displayFlag"`
	PointNotApplicable string  `json:"pointNotApplicable" csv:"pointNotApplicable"`
	TaxFreeDivision    string  `json:"taxFreeDivision" csv:"taxFreeDivision"`
	Color              *string `json:"color" csv:"color"`
	Tag                *string `json:"tag" csv:"tag"`
	InsDateTime        string  `json:"insDateTime" csv:"insDateTime"`
	UpdDateTime        string  `json:"updDateTime" csv:"updDateTime"`
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
