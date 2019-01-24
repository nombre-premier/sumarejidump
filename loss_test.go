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

func TestLoss(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "loss")
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
		File("./testdata/loss.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{LOSS},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "Loss.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	defer csvFile.Close()

	ls := []*Loss{}

	if err := gocsv.UnmarshalFile(csvFile, &ls); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(ls) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(ls))
	}

	l1 := Loss{
		LossID:           json.Number("1"),
		StoreID:          json.Number("1"),
		Division:         "2",
		Memo:             "店内混雑で気づいたらなくなっていた。",
		LossDatetime:     "2019-01-01 00:00:00",
		IdentificationNo: str2Ptr(""),
		Modified:         "2019-01-24 10:40:45",
	}

	if !reflect.DeepEqual(ls[0], &l1) {
		t.Fatalf("failed test expected: %v\n got: %v", l1, *ls[0])
	}

	l2 := Loss{
		LossID:           json.Number("3"),
		StoreID:          json.Number("1"),
		Division:         "1",
		Memo:             "袖に穴が空いているため",
		LossDatetime:     "2019-01-23 00:00:00",
		IdentificationNo: str2Ptr(""),
		Modified:         "2019-01-24 10:42:36",
	}

	if !reflect.DeepEqual(ls[1], &l2) {
		t.Fatalf("failed test expected: %v\n got: %v", l2, *ls[1])
	}
}

func TestLossDetail(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "loss_detail")
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
		File("./testdata/loss_detail.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{LOSS_DETAIL},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "LossDetail.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	ld := []*LossDetail{}

	if err := gocsv.UnmarshalFile(csvFile, &ld); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(ld) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(ld))
	}

	l1 := LossDetail{
		LossID:            json.Number("1"),
		ProductID:         json.Number("29"),
		ProductCode:       "2000002000105",
		ProductName:       "コットントラウザーズ(SS)",
		Size:              "ＳＳ",
		Color:             "ベージュ",
		GroupCode:         "bottoms002",
		SupplierProductNo: str2Ptr(""),
		Quantity:          json.Number("1"),
		Modified:          "2019-01-24 10:40:45",
	}

	if !reflect.DeepEqual(ld[0], &l1) {
		t.Fatalf("failed test expected: %v\n got: %v", l1, *ld[0])
	}

	l2 := LossDetail{
		LossID:            json.Number("3"),
		ProductID:         json.Number("12"),
		ProductCode:       "2000001000120",
		ProductName:       "チェックシャツ(L)",
		Size:              "Ｌ",
		Color:             "ブルー",
		GroupCode:         "tops002",
		SupplierProductNo: str2Ptr(""),
		Quantity:          json.Number("1"),
		Modified:          "2019-01-24 10:42:36",
	}

	if !reflect.DeepEqual(ld[1], &l2) {
		t.Fatalf("failed test expected: %v\n got: %v", l2, *ld[1])
	}
}
