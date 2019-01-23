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

func TestBargain(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "bargain")
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
		File("./testdata/bargain.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{BARGAIN},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "Bargain.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	bargains := []*Bargain{}

	if err := gocsv.UnmarshalFile(csvFile, &bargains); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(bargains) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(bargains))
	}

	b1 := Bargain{
		BargainID:   json.Number("1"),
		BargainName: "testSale",
		TermStart:   "2018-10-05",
		TermEnd:     "2018-10-31",
	}

	if !reflect.DeepEqual(bargains[0], &b1) {
		t.Fatalf("failed test expected: %v\n got: %v", b1, *bargains[0])
	}

	b2 := Bargain{
		BargainID:   json.Number("2"),
		BargainName: "testSale2",
		TermStart:   "2018-10-11",
		TermEnd:     "2019-10-31",
	}

	if !reflect.DeepEqual(bargains[1], &b2) {
		t.Fatalf("failed test expected: %v\n got: %v", b2, *bargains[1])
	}
}

func TestBargainProduct(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "bargain_product")
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
		File("./testdata/bargain_product.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{BARGAIN_PRODUCT},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "BargainProduct.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	bps := []*BargainProduct{}

	if err := gocsv.UnmarshalFile(csvFile, &bps); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(bps) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(bps))
	}

	b1 := BargainProduct{
		BargainProductID: json.Number("1"),
		BargainID:        json.Number("1"),
		TargetDivision:   "2",
		TargetID:         "1225",
		Division:         "1",
		Value:            json.Number("30"),
	}

	if !reflect.DeepEqual(bps[0], &b1) {
		t.Fatalf("failed test expected: %v\n got: %v", b1, *bps[0])
	}

	b2 := BargainProduct{
		BargainProductID: json.Number("2"),
		BargainID:        json.Number("2"),
		TargetDivision:   "2",
		TargetID:         "1234",
		Division:         "1",
		Value:            json.Number("50"),
	}

	if !reflect.DeepEqual(bps[1], &b2) {
		t.Fatalf("failed test expected: %v\n got: %v", b2, *bps[1])
	}
}
