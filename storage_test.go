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

func TestStorage(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "storage")
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
		File("./testdata/storage.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{STORAGE},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "Storage.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	defer csvFile.Close()

	st := []*Storage{}

	if err := gocsv.UnmarshalFile(csvFile, &st); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(st) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(st))
	}

	s1 := Storage{
		StorageID:               json.Number("2"),
		StorageInfoID:           json.Number("1"),
		SupplierID:              json.Number("1"),
		StorageStoreID:          json.Number("1"),
		StorageExpectedDateFrom: "2019-01-30",
		StorageExpectedDateTo:   "2019-02-02",
		StorageDate:             str2Ptr(""),
		Memo:                    "フォロー入荷分",
		Status:                  "0",
		IdentificationNo:        str2Ptr(""),
		Modified:                "2019-01-24 11:38:20",
	}

	if !reflect.DeepEqual(st[0], &s1) {
		t.Fatalf("failed test expected: %v\n got: %v", s1, *st[0])
	}

	s2 := Storage{
		StorageID:               json.Number("3"),
		StorageInfoID:           json.Number("2"),
		SupplierID:              json.Number("2"),
		StorageStoreID:          json.Number("1"),
		StorageExpectedDateFrom: "2019-01-24",
		StorageExpectedDateTo:   "2019-01-30",
		StorageDate:             str2Ptr(""),
		Memo:                    "イベントで配布のため、キャリー品を補填希望",
		Status:                  "0",
		IdentificationNo:        str2Ptr(""),
		Modified:                "2019-01-24 11:40:29",
	}

	if !reflect.DeepEqual(st[1], &s2) {
		t.Fatalf("failed test expected: %v\n got: %v", s2, *st[1])
	}
}

func TestStorageDetail(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "storage_detail")
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
		File("./testdata/storage_detail.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{STORAGE_DETAIL},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "StorageDetail.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	sds := []*StorageDetail{}

	if err := gocsv.UnmarshalFile(csvFile, &sds); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(sds) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(sds))
	}

	s1 := StorageDetail{
		StorageID:              json.Number("1"),
		ProductID:              json.Number("1"),
		ProductCode:            "2000001000014",
		ProductName:            "ユニセックススウェット(S)",
		Size:                   "S",
		Color:                  "レッド",
		GroupCode:              "tops001",
		SupplierProductNo:      str2Ptr(""),
		Cost:                   num2Ptr(""),
		ScheduledQuantity:      json.Number("2"),
		InspectionQuantity:     json.Number("2"),
		StockoutQuantity:       json.Number("0"),
		StockoutReason:         "",
		InspectionDate:         "2019-01-24",
		CompulsoryCompleteFlag: "0",
		Status:                 "1",
		Modified:               "2019-01-24 12:07:41",
	}

	if !reflect.DeepEqual(sds[0], &s1) {
		t.Fatalf("failed test expected: %v\n got: %v", s1, *sds[0])
	}

	s2 := StorageDetail{
		StorageID:              json.Number("1"),
		ProductID:              json.Number("2"),
		ProductCode:            "2000001000021",
		ProductName:            "ユニセックススウェット(M)",
		Size:                   "M",
		Color:                  "レッド",
		GroupCode:              "tops001",
		SupplierProductNo:      str2Ptr(""),
		Cost:                   num2Ptr(""),
		ScheduledQuantity:      json.Number("2"),
		InspectionQuantity:     json.Number("2"),
		StockoutQuantity:       json.Number("0"),
		StockoutReason:         "",
		InspectionDate:         "2019-01-24",
		CompulsoryCompleteFlag: "0",
		Status:                 "1",
		Modified:               "2019-01-24 12:07:41",
	}

	if !reflect.DeepEqual(sds[1], &s2) {
		t.Fatalf("failed test expected: %v\n got: %v", s2, *sds[1])
	}
}
