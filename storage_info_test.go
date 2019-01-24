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

func TestStorageInfo(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "storage_info")
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
		File("./testdata/storage_info.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{STORAGE_INFO},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "StorageInfo.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	si := []*StorageInfo{}

	if err := gocsv.UnmarshalFile(csvFile, &si); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(si) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(si))
	}

	s1 := StorageInfo{
		StorageInfoID:    json.Number("1"),
		RecipientOrderID: json.Number("1"),
		OrderedDate:      "2019-01-24",
		Status:           "2",
		IdentificationNo: str2Ptr(""),
		Modified:         "2019-01-24 11:38:20",
	}

	if !reflect.DeepEqual(si[0], &s1) {
		t.Fatalf("failed test expected: %v\n got: %v", s1, *si[0])
	}

	s2 := StorageInfo{
		StorageInfoID:    json.Number("2"),
		RecipientOrderID: json.Number("2"),
		OrderedDate:      "2019-01-03",
		Status:           "2",
		IdentificationNo: str2Ptr(""),
		Modified:         "2019-01-24 11:40:29",
	}

	if !reflect.DeepEqual(si[1], &s2) {
		t.Fatalf("failed test expected: %v\n got: %v", s2, *si[1])
	}
}

func TestStorageInfoDelivery(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "storage_info_delivery")
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
		File("./testdata/storage_info_delivery.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{STORAGE_INFO_DELIVERY},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "StorageInfoDelivery.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	sid := []*StorageInfoDelivery{}

	if err := gocsv.UnmarshalFile(csvFile, &sid); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(sid) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(sid))
	}

	s1 := StorageInfoDelivery{
		StorageInfoID:           json.Number("1"),
		StorageStoreID:          json.Number("1"),
		StorageExpectedDateFrom: "2019-01-30",
		StorageExpectedDateTo:   "2019-02-02",
		Modified:                "2019-01-24 11:38:20",
	}

	if !reflect.DeepEqual(sid[0], &s1) {
		t.Fatalf("failed test expected: %v\n got: %v", s1, *sid[0])
	}

	s2 := StorageInfoDelivery{
		StorageInfoID:           json.Number("2"),
		StorageStoreID:          json.Number("1"),
		StorageExpectedDateFrom: "2019-01-24",
		StorageExpectedDateTo:   "2019-01-30",
		Modified:                "2019-01-24 11:40:29",
	}

	if !reflect.DeepEqual(sid[1], &s2) {
		t.Fatalf("failed test expected: %v\n got: %v", s2, *sid[1])
	}
}

func TestStorageInfoProduct(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "storage_info_product")
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
		File("./testdata/storage_info_product.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{STORAGE_INFO_PRODUCT},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "StorageInfoProduct.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	sip := []*StorageInfoProduct{}

	if err := gocsv.UnmarshalFile(csvFile, &sip); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(sip) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(sip))
	}

	s1 := StorageInfoProduct{
		StorageInfoID:     json.Number("1"),
		ProductID:         json.Number("1"),
		ProductCode:       "2000001000014",
		ProductName:       "ユニセックススウェット(S)",
		Size:              "S",
		Color:             "レッド",
		GroupCode:         "tops001",
		SupplierProductNo: str2Ptr(""),
		Cost:              json.Number("12000.00000"),
		Quantity:          json.Number("1"),
		Modified:          "2019-01-24 11:38:20",
	}

	if !reflect.DeepEqual(sip[0], &s1) {
		t.Fatalf("failed test expected: %v\n got: %v", s1, *sip[0])
	}

	s2 := StorageInfoProduct{
		StorageInfoID:     json.Number("1"),
		ProductID:         json.Number("2"),
		ProductCode:       "2000001000021",
		ProductName:       "ユニセックススウェット(M)",
		Size:              "M",
		Color:             "レッド",
		GroupCode:         "tops001",
		SupplierProductNo: str2Ptr(""),
		Cost:              json.Number("12000.00000"),
		Quantity:          json.Number("1"),
		Modified:          "2019-01-24 11:38:20",
	}

	if !reflect.DeepEqual(sip[1], &s2) {
		t.Fatalf("failed test expected: %v\n got: %v", s2, *sip[1])
	}
}
