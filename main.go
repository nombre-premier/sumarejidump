package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Category struct {
	CategoryID         int    `json:"categoryId" csv:"categoryID"`
	CategoryCode       string `json:"categoryCode" csv:"categoryCode"`
	CategoryName       string `json:"categoryName" csv:"categoryName"`
	CategoryAbbr       string `json:"categoryAbbr" csv:"categoryAbbr"`
	CategoryGroupID    int    `json:"categoryGroupId" csv:"categoryGroupID"`
	ParentCategoryID   int    `json:"parentCategoryId" csv:"parentCategoryID"`
	Level              int    `json:"level" csv:"level"`
	DisplaySequence    int    `json:"displaySequence" csv:"displaySequence"`
	DisplayFlag        string `json:"displayFlag" csv:"displayFlag"`
	PointNotApplicable string `json:"pointNotApplicable" csv:"pointNotApplicable"`
	TaxFreeDivision    string `json:"taxFreeDivision" csv:"taxFreeDivision"`
	Color              string `json:"color" csv:"color"`
	Tag                string `json:"tag" csv:"tag"`
	InsDateTime        string `json:"insDateTime" csv:"insDateTime"`
	UpdDateTime        string `json:"updDateTime" csv:"updDateTime"`
}

type SumarejiRefParams struct {
	Fields    []string
	Condtions []string
	Order     []string
	Limit     int
	Page      int
	TableNmae string
}

func main() {
	urlStr := "https://webapi.smaregi.jp/access/"
	contractID := os.Getenv("SUMAREJI_CONTRACT_ID")
	accessToken := os.Getenv("SUMAREJI_ACCESS_TOKEN")

	data := url.Values{}
	data.Set("proc_nmae", "category_ref")
	data.Add("params", "bar")

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("X_contract_id", contractID)
	r.Header.Add("X_access_token", accessToken)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	fmt.Println(resp.Status)
}
