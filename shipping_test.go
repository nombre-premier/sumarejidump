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

func TestShipping(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "shipping")
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
		File("./testdata/shipping.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{SHIPPING},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "Shipping.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	shippings := []*Shipping{}

	if err := gocsv.UnmarshalFile(csvFile, &shippings); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(shippings) != 2 {
		t.Fatalf("Number of data should be 46. got=%d", len(shippings))
	}

	emptyStr := ""

	s1 := Shipping{
		ShippingID:                json.Number("1"),
		ShippingStoreID:           json.Number("2"),
		ReceivingStoreID:          json.Number("1"),
		ReceivingExpectedDateFrom: "2018-08-08",
		ReceivingExpectedDateTo:   "2018-08-08",
		ShippingDate:              "2018-08-08",
		ReceivingDesiredDate:      &emptyStr,
		Memo:                      "在庫補填",
		Status:                    "3",
		ModificationRequestStatus:          "0",
		ModificationRequestDateTime:        &emptyStr,
		ModificationRequestCheckedDateTime: &emptyStr,
		IdentificationNo:                   &emptyStr,
		Modified:                           "2018-08-08 19:17:02",
	}
	if !reflect.DeepEqual(shippings[0], &s1) {
		t.Fatalf("failed test expected: %v\n got: %v", s1, *shippings[0])
	}

	s2 := Shipping{
		ShippingID:                json.Number("2"),
		ShippingStoreID:           json.Number("1"),
		ReceivingStoreID:          json.Number("2"),
		ReceivingExpectedDateFrom: "2018-08-10",
		ReceivingExpectedDateTo:   "2018-08-12",
		ShippingDate:              "2018-08-10",
		ReceivingDesiredDate:      &emptyStr,
		Memo:                      "中目黒より在庫移動依頼",
		Status:                    "3",
		ModificationRequestStatus:          "0",
		ModificationRequestDateTime:        &emptyStr,
		ModificationRequestCheckedDateTime: &emptyStr,
		IdentificationNo:                   &emptyStr,
		Modified:                           "2018-08-10 18:43:49",
	}
	if !reflect.DeepEqual(shippings[1], &s2) {
		t.Fatalf("failed test expected: %v\n got: %v", s2, *shippings[45])
	}
}
