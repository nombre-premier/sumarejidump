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

func TestStocktakingInfo(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "stocktaking_info")
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
		File("./testdata/stocktaking_info.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{STOCKTAKING_INFO},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "StocktakingInfo.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	shs := []*StocktakingInfo{}

	if err := gocsv.UnmarshalFile(csvFile, &shs); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(shs) != 15 {
		t.Fatalf("Number of data should be 15. got=%d", len(shs))
	}

	si1 := StocktakingInfo{
		StocktakingInfoID:       json.Number("3"),
		StoreID:                 json.Number("1"),
		DivisionUnit:            "0",
		TargetDate:              "2018-07-31",
		AdjustmentDate:          "2018-08-05 18:37:55",
		StocktakingCompleteDate: "2018-08-05 18:37:55",
		Status:                  "2",
		Created:                 "2018-07-29 12:44:13",
		Modified:                "2018-08-05 18:37:55",
	}

	if !reflect.DeepEqual(shs[0], &si1) {
		t.Fatalf("failed test expected: %v\n got: %v", si1, *shs[0])
	}

	si15 := StocktakingInfo{
		StocktakingInfoID:       json.Number("23"),
		StoreID:                 json.Number("1"),
		DivisionUnit:            "0",
		TargetDate:              "2018-12-31",
		AdjustmentDate:          "2019-01-10 15:41:16",
		StocktakingCompleteDate: "2019-01-10 15:41:16",
		Status:                  "2",
		Created:                 "2019-01-06 17:05:22",
		Modified:                "2019-01-10 15:41:16",
	}
	if !reflect.DeepEqual(shs[14], &si15) {
		t.Fatalf("failed test expected: %v\n got: %v", si15, *shs[14])
	}
}

func TestStocktakingHead(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "stocktaking_head")
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
		File("./testdata/stocktaking_head.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{STOCKTAKING_HEAD},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "StocktakingHead.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	shs := []*StocktakingHead{}

	if err := gocsv.UnmarshalFile(csvFile, &shs); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(shs) != 15 {
		t.Fatalf("Number of data should be 15. got=%d", len(shs))
	}

	sh1 := StocktakingHead{
		StocktakingInfoID: json.Number("3"),
		StocktakingHeadID: json.Number("3"),
		DivisionCode:      "0",
		Modified:          "2018-08-05 18:37:55",
	}

	if !reflect.DeepEqual(shs[0], &sh1) {
		t.Fatalf("failed test expected: %v\n got: %v", sh1, *shs[0])
	}

	sh15 := StocktakingHead{
		StocktakingInfoID: json.Number("23"),
		StocktakingHeadID: json.Number("85"),
		DivisionCode:      "0",
		Modified:          "2019-01-10 15:41:16",
	}

	if !reflect.DeepEqual(shs[14], &sh15) {
		t.Fatalf("failed test expected: %v\n got: %v", sh15, *shs[14])
	}
}

func TestStocktakingDetail(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "stocktaking_detail")
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
		File("./testdata/stocktaking_detail.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{STOCKTAKING_DETAIL},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "StocktakingDetail.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	sds := []*StocktakingDetail{}

	if err := gocsv.UnmarshalFile(csvFile, &sds); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(sds) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(sds))
	}

	emptyStr := ""

	sd1 := StocktakingDetail{
		StocktakingInfoID:             json.Number("3"),
		StocktakingHeadID:             json.Number("3"),
		ProductID:                     json.Number("1"),
		ProductCode:                   "30000000000001",
		ProductName:                   "Foo",
		Size:                          "1",
		Color:                         "87",
		GroupCode:                     &emptyStr,
		SupplierProductNo:             &emptyStr,
		StocktakingQuantity:           json.Number("0"),
		TransportationStockQuantity:   json.Number("0"),
		LayawayStockQuantity:          json.Number("0"),
		StockQuantityBeforeAdjustment: json.Number("0"),
		Cost:     json.Number("18815.00000"),
		Memo:     &emptyStr,
		Modified: "2018-08-05 18:37:55",
	}

	if !reflect.DeepEqual(sds[0], &sd1) {
		t.Fatalf("failed test expected: %v\n got: %v", sd1, *sds[0])
	}

	sd2 := StocktakingDetail{
		StocktakingInfoID:             json.Number("3"),
		StocktakingHeadID:             json.Number("3"),
		ProductID:                     json.Number("2"),
		ProductCode:                   "30000000000002",
		ProductName:                   "Bar",
		Size:                          "2",
		Color:                         "87",
		GroupCode:                     &emptyStr,
		SupplierProductNo:             &emptyStr,
		StocktakingQuantity:           json.Number("0"),
		TransportationStockQuantity:   json.Number("0"),
		LayawayStockQuantity:          json.Number("0"),
		StockQuantityBeforeAdjustment: json.Number("0"),
		Cost:     json.Number("18815.00000"),
		Memo:     &emptyStr,
		Modified: "2018-08-05 18:37:55",
	}

	if !reflect.DeepEqual(sds[1], &sd2) {
		t.Fatalf("failed test expected: %v\n got: %v", sd2, *sds[1])
	}
}
