package main

import "encoding/json"

type Category struct {
	CategoryID         json.Number  `json:"categoryId" csv:"categoryId" parquet:"name=categoryId, type=INT64"`
	CategoryID         json.Number  `json:"categoryId" csv:"categoryId" parquet:"name=categoryId, type=INT64"`
	CategoryCode       string       `json:"categoryCode" csv:"categoryCode" parquet:"name=categoryCode, type=BYTE_ARRAY convertedtype=UTF8"`
	CategoryName       string       `json:"categoryName" csv:"categoryName" parquet:"name=categoryName, type=BYTE_ARRAY convertedtype=UTF8"`
	CategoryAbbr       string       `json:"categoryAbbr" csv:"categoryAbbr" parquet:"name=categoryAbbr, type=BYTE_ARRAY convertedtype=UTF8"`
	CategoryGroupID    *json.Number `json:"categoryGroupId" csv:"categoryGroupId" parquet:"name=categoryGroupId, type=INT64"`
	ParentCategoryID   *json.Number `json:"parentCategoryId" csv:"parentCategoryId" parquet:"name=parentCategoryId, type=INT64"`
	Level              json.Number  `json:"level" csv:"level" parquet:"name=level, type=INT64"`
	DisplaySequence    json.Number  `json:"displaySequence" csv:"displaySequence" parquet:"name=displaySequence, type=INT64"`
	DisplayFlag        string       `json:"displayFlag" csv:"displayFlag" parquet:"name=displayFlag, type=BYTE_ARRAY convertedtype=UTF8"`
	PointNotApplicable string       `json:"pointNotApplicable" csv:"pointNotApplicable" parquet:"name=pointNotApplicable, type=BYTE_ARRAY convertedtype=UTF8"`
	TaxFreeDivision    string       `json:"taxFreeDivision" csv:"taxFreeDivision" parquet:"name=taxFreeDivision, type=BYTE_ARRAY convertedtype=UTF8"`
	Color              *string      `json:"color" csv:"color" parquet:"name=color, type=BYTE_ARRAY convertedtype=UTF8"`
	Tag                *string      `json:"tag" csv:"tag" parquet:"name=tag, type=BYTE_ARRAY convertedtype=UTF8"`
	InsDateTime        string       `json:"insDateTime" csv:"insDateTime" parquet:"name=insDateTime, type=BYTE_ARRAY convertedtype=UTF8"`
	UpdDateTime        string       `json:"updDateTime" csv:"updDateTime" parquet:"name=updDateTime, type=BYTE_ARRAY convertedtype=UTF8"`
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

type CategoryParquet struct {
	*ParquetHandler
	buf []Category
}

func NewCategoryParquet(header interface{}, output string) (*CategoryParquet, error) {
	ph, err := NewParquetHandler(header, output)
	if err != nil {
		return nil, err
	}
	return &CategoryParquet{
		ParquetHandler: ph,
	}, nil
}

func (cp *CategoryParquet) Write(resp *SrRefResponse) error {
	for _, r := range resp.Result {
		var category Category
		json.Unmarshal([]byte(r.String()), &category)
		if err := cp.ParquetHandler.ParquetWriter.Write(category); err != nil {
			return err
		}
	}
	return nil
}

func (cc *CategoryCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &cc.buf[i])
	}
	cc.CSVWriter.Write(cc.buf[:len(resp.Result)])
	return cc.CSVWriter
}
