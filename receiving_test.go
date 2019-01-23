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

func TestReceiving(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "receiving")
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
		File("./testdata/receiving.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{RECEIVING},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "Receiving.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	rs := []*Receiving{}

	if err := gocsv.UnmarshalFile(csvFile, &rs); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(rs) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(rs))
	}

	emptyStr := ""

	r1 := Receiving{
		ReceivingID:      json.Number("1"),
		ShippingID:       json.Number("1"),
		ShippingStoreID:  json.Number("2"),
		ReceivingStoreID: json.Number("1"),
		ReceivingDate:    "2018-08-08",
		Memo:             "在庫補填",
		Status:           "2",
		IdentificationNo: &emptyStr,
		Modified:         "2018-08-08 19:17:02",
	}
	if !reflect.DeepEqual(rs[0], &r1) {
		t.Fatalf("failed test expected: %v\n got: %v", r1, *rs[0])
	}

	r2 := Receiving{
		ReceivingID:      json.Number("2"),
		ShippingID:       json.Number("2"),
		ShippingStoreID:  json.Number("1"),
		ReceivingStoreID: json.Number("2"),
		ReceivingDate:    "2018-08-10",
		Memo:             "中目黒より在庫移動依頼",
		Status:           "2",
		IdentificationNo: &emptyStr,
		Modified:         "2018-08-10 18:43:49",
	}
	if !reflect.DeepEqual(rs[1], &r2) {
		t.Fatalf("failed test expected: %v\n got: %v", r2, *rs[1])
	}
}

func TestReceivingDetail(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "receiving_detail")
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
		File("./testdata/receiving_detail.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{RECEIVING_DETAIL},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "ReceivingDetail.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	rs := []*ReceivingDetail{}

	if err := gocsv.UnmarshalFile(csvFile, &rs); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(rs) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(rs))
	}

	emptyStr := ""

	r1 := ReceivingDetail{
		ReceivingID:        json.Number("1"),
		ProductID:          json.Number("621"),
		ProductCode:        "2000000001401",
		ProductName:        "Foo",
		Size:               "3",
		Color:              "90",
		GroupCode:          &emptyStr,
		SupplierProductNo:  &emptyStr,
		ScheduledQuantity:  json.Number("1"),
		InspectionQuantity: json.Number("1"),
		StockoutQuantity:   json.Number("0"),
		StockoutReason:     &emptyStr,
		InspectionDate:     "2018-08-08",
		Status:             "1",
		Modified:           "2018-08-08 19:17:02",
	}
	if !reflect.DeepEqual(rs[0], &r1) {
		t.Fatalf("failed test expected: %v\n got: %v", r1, *rs[0])
	}

	r2 := ReceivingDetail{
		ReceivingID:        json.Number("1"),
		ProductID:          json.Number("873"),
		ProductCode:        "2000000003849",
		ProductName:        "Bar",
		Size:               "3",
		Color:              "90",
		GroupCode:          &emptyStr,
		SupplierProductNo:  &emptyStr,
		ScheduledQuantity:  json.Number("1"),
		InspectionQuantity: json.Number("1"),
		StockoutQuantity:   json.Number("0"),
		StockoutReason:     &emptyStr,
		InspectionDate:     "2018-08-08",
		Status:             "1",
		Modified:           "2018-08-08 19:17:02",
	}
	if !reflect.DeepEqual(rs[1], &r2) {
		t.Fatalf("failed test expected: %v\n got: %v", r2, *rs[1])
	}
}
