package main

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"

	"gopkg.in/h2non/gock.v1"
)

func TestError14(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "category")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir) // clean up

	gock.New("https://webapi.smaregi.jp").
		Post("/access").
		MatchHeader("X_contract_id", "").
		MatchHeader("X_access_token", "(.*)").
		MatchHeader("Content-Type", "application/x-www-form-urlencoded").
		Reply(400).
		File("./testdata/error_14.json")

	c := SrConfig{
		ContractID:  "",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{CATEGORY},
	}

	err = Main(c)

	srErr := SrError{
		ErrorCode:        14,
		ErrorSummary:     "契約IDが無効です。",
		ErrroDescription: "パラメータ不正です。(contractId is empty)",
	}

	if !reflect.DeepEqual(err, &srErr) {
		t.Fatalf("failed test expected: %v\n got: %v", srErr, err)
	}
}
