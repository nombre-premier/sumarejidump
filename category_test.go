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

func TestCategory(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "category")
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
		File("./testdata/category.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{CATEGORY},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "Category.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	categories := []*Category{}

	if err := gocsv.UnmarshalFile(csvFile, &categories); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(categories) != 46 {
		t.Fatalf("Number of data should be 46. got=%d", len(categories))
	}

	emptyStr := ""
	emptyJSONNumber := json.Number("")

	cat1 := Category{
		CategoryID:         "1",
		CategoryCode:       "dbss",
		CategoryName:       "dbss",
		CategoryAbbr:       "dbss",
		CategoryGroupID:    &emptyJSONNumber,
		ParentCategoryID:   &emptyJSONNumber,
		Level:              "1",
		DisplaySequence:    "",
		DisplayFlag:        "1",
		PointNotApplicable: "0",
		TaxFreeDivision:    "0",
		Color:              &emptyStr,
		Tag:                &emptyStr,
		InsDateTime:        "2018-07-20 23:41:05",
		UpdDateTime:        "2018-07-20 23:41:05",
	}

	if !reflect.DeepEqual(categories[0], &cat1) {
		t.Fatalf("failed test expected: %q\n got: %q", cat1, *categories[0])
	}

	cat46 := Category{
		CategoryID:         "46",
		CategoryCode:       "marine",
		CategoryName:       "LA MARINE FRANCAISE",
		CategoryAbbr:       "Marine",
		CategoryGroupID:    &emptyJSONNumber,
		ParentCategoryID:   &emptyJSONNumber,
		Level:              "1",
		DisplaySequence:    "",
		DisplayFlag:        "1",
		PointNotApplicable: "0",
		TaxFreeDivision:    "1",
		Color:              &emptyStr,
		Tag:                &emptyStr,
		InsDateTime:        "2018-11-27 23:42:01",
		UpdDateTime:        "2018-11-30 19:16:02",
	}
	if !reflect.DeepEqual(categories[45], &cat46) {
		t.Fatalf("failed test expected: %q\n got: %q", cat46, *categories[45])
	}
}
