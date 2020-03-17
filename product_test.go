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
		ProductID:              "1",
		CategoryID:             "2",
		ProductCode:            "2990000000001",
		ProductName:            "D181H223  Joined balmacaan denim coat",
		ProductKana:            &emptyStr,
		TaxDivision:            "1",
		ProductPriceDivision:   "1",
		Price:                  json.Number("69000"),
		CustomerPrice:          &emptyJSONNumber,
		Cost:                   json.Number("100000.00000"),
		Attribute:              &emptyStr,
		Description:            &emptyStr,
		CatchCopy:              &emptyStr,
		Size:                   "1",
		Color:                  "87",
		Tag:                    &emptyStr,
		GroupCode:              &emptyStr,
		URL:                    &emptyStr,
		DisplaySequence:        &emptyJSONNumber,
		DisplayFlag:            "1",
		SalesDivision:          "0",
		Division:               "0",
		ProductOptionGroupID:   &emptyJSONNumber,
		StockControlDivision:   "0",
		PointNotApplicable:     "0",
		TaxFreeDivision:        "1",
		SupplierProductNo:      &emptyStr,
		StaffDiscountRate:      &emptyJSONNumber,
		UseCategoryReduceTax:   json.Number("0"),
		ReduceTaxId:            &emptyJSONNumber,
		ReduceTaxPrice:         json.Number("69000"),
		ReduceTaxCustomerPrice: json.Number("69000"),
		AppStartDateTime:       &emptyStr,
		InsDateTime:            "2018-07-21 01:28:22",
		UpdDateTime:            "2018-12-06 16:40:43",
	}

	if !reflect.DeepEqual(products[0], &pro1) {
		t.Fatalf("failed test expected: %v\n got: %v", pro1, *products[0])
	}

	pro2 := Product{
		ProductID:              "2",
		CategoryID:             "2",
		ProductCode:            "2990000000002",
		ProductName:            "D181H223  Joined balmacaan denim coat",
		ProductKana:            &emptyStr,
		TaxDivision:            "1",
		ProductPriceDivision:   "1",
		Price:                  json.Number("69000"),
		CustomerPrice:          &emptyJSONNumber,
		Cost:                   json.Number("100000.00000"),
		Attribute:              &emptyStr,
		Description:            &emptyStr,
		CatchCopy:              &emptyStr,
		Size:                   "2",
		Color:                  "87",
		Tag:                    &emptyStr,
		GroupCode:              &emptyStr,
		URL:                    &emptyStr,
		DisplaySequence:        &emptyJSONNumber,
		DisplayFlag:            "1",
		SalesDivision:          "0",
		Division:               "0",
		ProductOptionGroupID:   &emptyJSONNumber,
		StockControlDivision:   "0",
		PointNotApplicable:     "0",
		TaxFreeDivision:        "1",
		SupplierProductNo:      &emptyStr,
		StaffDiscountRate:      &emptyJSONNumber,
		UseCategoryReduceTax:   json.Number("0"),
		ReduceTaxId:            &emptyJSONNumber,
		ReduceTaxPrice:         json.Number("69000"),
		ReduceTaxCustomerPrice: json.Number("69000"),
		AppStartDateTime:       &emptyStr,
		InsDateTime:            "2018-07-21 01:28:22",
		UpdDateTime:            "2018-12-06 16:40:43",
	}
	if !reflect.DeepEqual(products[1], &pro2) {
		t.Fatalf("failed test expected: %v\n got: %v", pro2, *products[1])
	}
}

//==========ProductPrice=============================================================================================================================
func TestProductPrice(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "product_price")
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
		File("./testdata/product_price.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{PRODUCT_PRICE},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "ProductPrice.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	pp := []*ProductPrice{}

	if err := gocsv.UnmarshalFile(csvFile, &pp); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(pp) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(pp))
	}

	p1 := ProductPrice{
		ProductID:     json.Number("1"),
		StoreID:       "_ALL_",
		PriceDivision: "1",
		StartDate:     "2019-02-01",
		EndDate:       "2019-02-14",
		Price:         json.Number("12000"),
		InsDateTime:   "2019-01-25 18:29:48",
		UpdDateTime:   "2019-01-25 18:29:48",
	}

	if !reflect.DeepEqual(pp[0], &p1) {
		t.Fatalf("failed test expected: %v\n got: %v", p1, *pp[0])
	}

	p2 := ProductPrice{
		ProductID:     json.Number("2"),
		StoreID:       "_ALL_",
		PriceDivision: "1",
		StartDate:     "2019-02-01",
		EndDate:       "2019-02-14",
		Price:         json.Number("12000"),
		InsDateTime:   "2019-01-25 18:34:25",
		UpdDateTime:   "2019-01-25 18:34:25",
	}

	if !reflect.DeepEqual(pp[1], &p2) {
		t.Fatalf("failed test expected: %v\n got: %v", p2, *pp[1])
	}
}

//==============ProductReserveItem=========================================================================================================================
func TestProductReserveItem(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "product_reserve_item")
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
		File("./testdata/product_reserve_item.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{PRODUCT_RESERVE_ITEM},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "ProductReserveItem.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	pri := []*ProductReserveItem{}

	if err := gocsv.UnmarshalFile(csvFile, &pri); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(pri) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(pri))
	}

	p1 := ProductReserveItem{
		ProductID: json.Number("1"),
		No:        json.Number("1"),
		Value:     "ゆったりデザイン",
	}

	if !reflect.DeepEqual(pri[0], &p1) {
		t.Fatalf("failed test expected: %v\n got: %v", p1, *pri[0])
	}

	p2 := ProductReserveItem{
		ProductID: json.Number("1"),
		No:        json.Number("2"),
		Value:     "★色落注意★",
	}

	if !reflect.DeepEqual(pri[1], &p2) {
		t.Fatalf("failed test expected: %v\n got: %v", p2, *pri[1])
	}
}

//=================ProductReserveItemLabel======================================================================================================================
func TestProductReserveItemLabel(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "product_reserve_item_label")
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
		File("./testdata/product_reserve_item_label.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{PRODUCT_RESERVE_ITEM_LABEL},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "ProductReserveItemLabel.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	pril := []*ProductReserveItemLabel{}

	if err := gocsv.UnmarshalFile(csvFile, &pril); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(pril) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(pril))
	}

	p1 := ProductReserveItemLabel{
		No:    "1",
		Label: "memo2",
	}

	if !reflect.DeepEqual(pril[0], &p1) {
		t.Fatalf("failed test expected: %v\n got: %v", p1, *pril[0])
	}

	p2 := ProductReserveItemLabel{
		No:    "2",
		Label: "memo3",
	}

	if !reflect.DeepEqual(pril[1], &p2) {
		t.Fatalf("failed test expected: %v\n got: %v", p2, *pril[1])
	}
}

//=====================ProductStore==================================================================================================================
func TestProductStore(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "product_store")
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
		File("./testdata/product_store.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{PRODUCT_STORE},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "ProductStore.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	ps := []*ProductStore{}

	if err := gocsv.UnmarshalFile(csvFile, &ps); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(ps) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(ps))
	}

	p1 := ProductStore{
		ProductID:      json.Number("1225"),
		StoreID:        json.Number("1"),
		AssignDivision: "0",
	}

	if !reflect.DeepEqual(ps[0], &p1) {
		t.Fatalf("failed test expected: %v\n got: %v", p1, *ps[0])
	}

	p2 := ProductStore{
		ProductID:      json.Number("1225"),
		StoreID:        json.Number("2"),
		AssignDivision: "0",
	}

	if !reflect.DeepEqual(ps[1], &p2) {
		t.Fatalf("failed test expected: %v\n got: %v", p2, *ps[1])
	}
}

//==========================ProductInventoryReservation=============================================================================================================
func TestProductInventoryReservation(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "product_inventory_reservation")
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
		File("./testdata/product_inventory_reservation.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{PRODUCT_INVENTORY_RESERVATION},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "ProductInventoryReservation.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	pir := []*ProductInventoryReservation{}

	if err := gocsv.UnmarshalFile(csvFile, &pir); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(pir) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(pir))
	}

	p1 := ProductInventoryReservation{
		ProductID:            json.Number("1203"),
		ReservationProductID: "553",
		ReservationAmount:    json.Number("1"),
	}

	if !reflect.DeepEqual(pir[0], &p1) {
		t.Fatalf("failed test expected: %v\n got: %v", p1, *pir[0])
	}

	p2 := ProductInventoryReservation{
		ProductID:            json.Number("1203"),
		ReservationProductID: "554",
		ReservationAmount:    json.Number("1"),
	}

	if !reflect.DeepEqual(pir[1], &p2) {
		t.Fatalf("failed test expected: %v\n got: %v", p2, *pir[1])
	}
}
