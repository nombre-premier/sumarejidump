package main

type Category struct {
	CategoryID         int     `json:"categoryId" csv:"categoryID"`
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
