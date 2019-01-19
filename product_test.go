package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"reflect"
	"testing"

	"github.com/gocarina/gocsv"
	"gopkg.in/h2non/gock.v1"
)

func TestProduct(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "product")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir) // clean up

	gock.New("https://webapi.smaregi.jp").
		Post("/access").
		MatchHeader("X_contract_id", "(.*)").
		MatchHeader("X_access_token", "(.*)").
		MatchHeader("Content-Type", "application/x-www-form-urlencoded").
		Reply(200).
		File("./testdata/product.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{PRODUCT},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "Product.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	products := []*Product{}

	if err := gocsv.UnmarshalFile(csvFile, &products); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(products) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(products))
	}

	emptyStr := ""
	emptyJSONNumber := json.Number("")

	pro1 := Product{
		ProductID:            "1",
		CategoryID:           "2",
		ProductCode:          "2990000000001",
		ProductName:          "D181H223  Joined balmacaan denim coat",
		ProductKana:          &emptyStr,
		TaxDivision:          "1",
		ProductPriceDivision: "1",
		Price:                json.Number("69000"),
		CustomerPrice:        &emptyJSONNumber,
		Cost:                 json.Number("100000.00000"),
		Attribute:            &emptyStr,
		Description:          &emptyStr,
		CatchCopy:            &emptyStr,
		Size:                 "1",
		Color:                "87",
		Tag:                  &emptyStr,
		GroupCode:            &emptyStr,
		URL:                  &emptyStr,
		DisplaySequence:      &emptyJSONNumber,
		DisplayFlag:          "1",
		SalesDivision:        "0",
		Division:             "0",
		ProductOptionGroupID: &emptyJSONNumber,
		StockControlDivision: "0",
		PointNotApplicable:   "0",
		TaxFreeDivision:      "1",
		SupplierProductNo:    &emptyStr,
		StaffDiscountRate:    &emptyJSONNumber,
		AppStartDateTime:     &emptyStr,
		InsDateTime:          "2018-07-21 01:28:22",
		UpdDateTime:          "2018-12-06 16:40:43",
	}

	if !reflect.DeepEqual(products[0], &pro1) {
		t.Fatalf("failed test expected: %v\n got: %v", pro1, *products[0])
	}

	pro2 := Product{
		ProductID:            "2",
		CategoryID:           "2",
		ProductCode:          "2990000000002",
		ProductName:          "D181H223  Joined balmacaan denim coat",
		ProductKana:          &emptyStr,
		TaxDivision:          "1",
		ProductPriceDivision: "1",
		Price:                json.Number("69000"),
		CustomerPrice:        &emptyJSONNumber,
		Cost:                 json.Number("100000.00000"),
		Attribute:            &emptyStr,
		Description:          &emptyStr,
		CatchCopy:            &emptyStr,
		Size:                 "2",
		Color:                "87",
		Tag:                  &emptyStr,
		GroupCode:            &emptyStr,
		URL:                  &emptyStr,
		DisplaySequence:      &emptyJSONNumber,
		DisplayFlag:          "1",
		SalesDivision:        "0",
		Division:             "0",
		ProductOptionGroupID: &emptyJSONNumber,
		StockControlDivision: "0",
		PointNotApplicable:   "0",
		TaxFreeDivision:      "1",
		SupplierProductNo:    &emptyStr,
		StaffDiscountRate:    &emptyJSONNumber,
		AppStartDateTime:     &emptyStr,
		InsDateTime:          "2018-07-21 01:28:22",
		UpdDateTime:          "2018-12-06 16:40:43",
	}
	if !reflect.DeepEqual(products[1], &pro2) {
		t.Fatalf("failed test expected: %v\n got: %v", pro2, *products[1])
	}
}
