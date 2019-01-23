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

func TestStore(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "store")
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
		File("./testdata/store.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{STORE},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "Store.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	stores := []*Store{}

	if err := gocsv.UnmarshalFile(csvFile, &stores); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(stores) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(stores))
	}

	emptyStr := ""

	s1 := Store{
		StoreID:                  json.Number("1"),
		StoreCode:                "dbss",
		StoreName:                "DBSS Aoyama",
		StoreAbbr:                str2Ptr("DBSS aoyama"),
		Division:                 "1",
		PostCode:                 str2Ptr("107-0062"),
		Address:                  str2Ptr("東京都港区南青山5-4-30南青山澁澤ビル1F"),
		PhoneNumber:              str2Ptr("03-6427-7399"),
		FaxNumber:                &emptyStr,
		MailAddress:              &emptyStr,
		Homepage:                 &emptyStr,
		TempTranMailAddress:      &emptyStr,
		PriceChangeFlag:          "1",
		SellDivision:             "1",
		SumProcDivision:          "1",
		SumDateChangeTime:        str2Ptr("0400"),
		SumRefColumn:             "0",
		PointNotApplicable:       "0",
		TaxFreeDivision:          "0",
		MaxBundleProductCount:    json.Number("9"),
		RoundingDivision:         "00",
		DiscountRoundingDivision: "2",
		PauseFlag:                json.Number("0"),
		PointUseDivision:         "0",
		InsDateTime:              "2018-07-06 14:09:27",
		UpdDateTime:              "2018-07-26 12:50:19",
	}

	if !reflect.DeepEqual(stores[0], &s1) {
		t.Fatalf("failed test expected: %v\n got: %v", s1, *stores[0])
	}

	s2 := Store{
		StoreID:                  json.Number("2"),
		StoreCode:                "logistics",
		StoreName:                "LOGISTICS",
		StoreAbbr:                &emptyStr,
		Division:                 "2",
		PostCode:                 &emptyStr,
		Address:                  &emptyStr,
		PhoneNumber:              &emptyStr,
		FaxNumber:                &emptyStr,
		MailAddress:              &emptyStr,
		Homepage:                 &emptyStr,
		TempTranMailAddress:      &emptyStr,
		PriceChangeFlag:          "1",
		SellDivision:             "0",
		SumProcDivision:          "0",
		SumDateChangeTime:        &emptyStr,
		SumRefColumn:             "0",
		PointNotApplicable:       "0",
		TaxFreeDivision:          "0",
		MaxBundleProductCount:    json.Number("9"),
		RoundingDivision:         "00",
		DiscountRoundingDivision: "2",
		PauseFlag:                json.Number("0"),
		PointUseDivision:         "0",
		InsDateTime:              "2018-07-20 13:28:07",
		UpdDateTime:              "2018-11-29 15:28:35",
	}
	if !reflect.DeepEqual(stores[1], &s2) {
		t.Fatalf("failed test expected: %v\n got: %v", s2, *stores[1])
	}
}
