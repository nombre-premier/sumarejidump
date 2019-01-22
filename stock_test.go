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

func TestStock(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "stock")
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
		File("./testdata/stock.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{STOCK},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "Stock.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	stocks := []*Stock{}

	if err := gocsv.UnmarshalFile(csvFile, &stocks); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(stocks) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(stocks))
	}

	sto1 := Stock{
		StoreID:     json.Number("1"),
		ProductID:   json.Number("1"),
		StockAmount: "0",
		UpdDateTime: "2018-09-25 12:50:50",
	}

	if !reflect.DeepEqual(stocks[0], &sto1) {
		t.Fatalf("failed test expected: %v\n got: %v", sto1, *stocks[0])
	}

	sto2 := Stock{
		StoreID:     json.Number("2"),
		ProductID:   json.Number("1"),
		StockAmount: "0",
		UpdDateTime: "2018-07-28 12:38:42",
	}
	if !reflect.DeepEqual(stocks[1], &sto2) {
		t.Fatalf("failed test expected: %v\n got: %v", sto2, *stocks[1])
	}
}
