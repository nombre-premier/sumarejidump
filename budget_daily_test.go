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

func TestBudgetDaily(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "budget_daily")
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
		File("./testdata/budget_daily.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{BUDGET_DAILY},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "BudgetDaily.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	budgetDailys := []*BudgetDaily{}

	if err := gocsv.UnmarshalFile(csvFile, &budgetDailys); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(budgetDailys) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(budgetDailys))
	}

	bd1 := BudgetDaily{
		StoreID:     json.Number("1"),
		YM:   json.Number("201908"),
		Day: "01",
		SalesTargetDaily: "100000",
	}

	if !reflect.DeepEqual(budgetDailys[0], &bd1) {
		t.Fatalf("failed test expected: %v\n got: %v", bd1, *budgetDailys[0])
	}

	bd2 := BudgetDaily{
		StoreID:     json.Number("1"),
		YM:   json.Number("201908"),
		Day: "02",
		SalesTargetDaily: "120000",
	}

	if !reflect.DeepEqual(budgetDailys[1], &bd2) {
		t.Fatalf("failed test expected: %v\n got: %v", bd2, *budgetDailys[1])
	}
}
