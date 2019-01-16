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
		panic(err)
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	products := []*Product{}

	if err := gocsv.UnmarshalFile(csvFile, &products); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(products) != 1000 {
		t.Fatalf("Number of data should be 1000. got=%d", len(products))
	}

	emptyStr := ""
	emptyJSONNumber := json.Number("")

	pro1 := Product{
		ProductID:							"0",
		CategoryID:							"0",
		ProductCode:						"2990000000001",
		ProductName:						"D181H223  Joined balmacaan denim coat",
		ProductKana:						"",
		TaxDivision:						"1",
		ProductPriceDivision:		"1",
		Price:					"69000",
		CustomerPrice:	"0",
		cost:	"18815",
		attribute:	&emptyStr,
		description: &emptyStr,
		catchCopy: &emptyStr,
		size: 	"1",
		color:	"87",
		tag: &emptyStr,
		GroupCode: &emptyStr,
		url:&emptyStr,
		displaySequence:&emptyJSONNumber,
		salesDivision:	"0",
		stockControlDivision:	"0",
		displayFlag:	"1",
		division:	"0",
		productOptionGroupID:&emptyJSONNumber ,
		pointNotApplicable:	"0",
		taxFreeDivision:	"1",
		supplierProductNo:	&emptyStr,
		staffDiscountRate:&emptyStr,
		appStartDateTime:&emptyStr,
		insDateTime:"2018/7/21 1:28",
		updDateTime:"2018/12/6 16:40",
}

	if !reflect.DeepEqual(products[0], &pro1) {
		t.Fatalf("failed test expected: %v\n got: %v", pro1, *products[0])
	}

	pro1000 := Products{
		productId:	"0",
		categoryId:	"0",
		productCode:	"2990000000001",
		productName:	"D181H223  Joined balmacaan denim coat",
		productKana:	&emptyStr,
		taxDivision:	"1",
		productPriceDivision:	"1",
		price:	"69000",
		customerPrice:	"0",
		cost:	"18815",
		attribute:	&emptyStr,
		description: &emptyStr,
		catchCopy: &emptyStr,
		size: 	"1",
		color:	"87",
		tag: &emptyStr,
		groupCode: &emptyStr,
		url:&emptyStr,
		displaySequence:&emptyJSONNumber,
		salesDivision:	"0",
		stockControlDivision:	"0",
		displayFlag:	"1",
		division:	"0",
		productOptionGroupID:&emptyJSONNumber ,
		pointNotApplicable:	"0",
		taxFreeDivision:	"1",
		supplierProductNo:	&emptyStr,
		staffDiscountRate:&emptyStr,
		appStartDateTime:&emptyStr,
		insDateTime:"2018/7/21 1:28",
		updDateTime:"2018/12/6 16:40",
	}
	if !reflect.DeepEqual(products[45], &pro1000) {
		t.Fatalf("failed test expected: %v\n got: %v", pro1000, *products[45])
	}
}
