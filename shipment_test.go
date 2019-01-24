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

func TestShipment(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "shipment")
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
		File("./testdata/shipment.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{SHIPMENT},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "Shipment.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	shipments := []*Shipment{}

	if err := gocsv.UnmarshalFile(csvFile, &shipments); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(shipments) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(shipments))
	}

	s1 := Shipment{
		ShipmentID:       json.Number("2"),
		ShipmentStoreID:  json.Number("3"),
		RecipientType:    "3",
		RecipientID:      json.Number("1"),
		RecipientName:    "本社",
		ShipmentDivision: "0",
		ShipmentDate:     "2019-01-22",
		Status:           "2",
		IdentificationNo: str2Ptr(""),
		Modified:         "2019-01-24 11:25:00",
	}

	if !reflect.DeepEqual(shipments[0], &s1) {
		t.Fatalf("failed test expected: %v\n got: %v", s1, *shipments[0])
	}

	s2 := Shipment{
		ShipmentID:       json.Number("4"),
		ShipmentStoreID:  json.Number("3"),
		RecipientType:    "3",
		RecipientID:      json.Number("2"),
		RecipientName:    "本社",
		ShipmentDivision: "2",
		ShipmentDate:     "2019-01-29",
		Status:           "2",
		IdentificationNo: str2Ptr(""),
		Modified:         "2019-01-24 11:29:59",
	}

	if !reflect.DeepEqual(shipments[1], &s2) {
		t.Fatalf("failed test expected: %v\n got: %v", s2, *shipments[1])
	}
}

func TestShipmentDetail(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "shipment_detail")
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
		File("./testdata/shipment_detail.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{SHIPMENT_DETAIL},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "ShipmentDetail.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	sds := []*ShipmentDetail{}

	if err := gocsv.UnmarshalFile(csvFile, &sds); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(sds) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(sds))
	}

	s1 := ShipmentDetail{
		ShipmentID:        json.Number("1"),
		ProductID:         json.Number("47"),
		ProductCode:       "2000003000012",
		ProductName:       "サンダル(24)",
		Size:              "24",
		Color:             "ブラック",
		GroupCode:         "foot001",
		SupplierProductNo: str2Ptr(""),
		Cost:              json.Number("3000.00000"),
		Quantity:          json.Number("1"),
		Memo:              "",
		Modified:          "2019-01-24 11:23:10",
	}

	if !reflect.DeepEqual(sds[0], &s1) {
		t.Fatalf("failed test expected: %v\n got: %v", s1, *sds[0])
	}

	s2 := ShipmentDetail{
		ShipmentID:        json.Number("1"),
		ProductID:         json.Number("48"),
		ProductCode:       "2000003000029",
		ProductName:       "サンダル(26)",
		Size:              "26",
		Color:             "ブラック",
		GroupCode:         "foot001",
		SupplierProductNo: str2Ptr(""),
		Cost:              json.Number("3000.00000"),
		Quantity:          json.Number("1"),
		Memo:              "",
		Modified:          "2019-01-24 11:23:10",
	}

	if !reflect.DeepEqual(sds[1], &s2) {
		t.Fatalf("failed test expected: %v\n got: %v", s2, *sds[1])
	}
}
