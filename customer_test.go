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

func str2Ptr(s string) *string {
	return &s
}

func num2Ptr(j json.Number) *json.Number {
	return &j
}

func TestCustomer(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "customer")
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
		File("./testdata/customer.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{CUSTOMER},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "Customer.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	customers := []*Customer{}

	if err := gocsv.UnmarshalFile(csvFile, &customers); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(customers) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(customers))
	}

	emptyStr := ""
	emptyJSONNumber := json.Number("")

	cus1 := Customer{
		CustomerID:           json.Number("1"),
		CustomerCode:         "01",
		CustomerNo:           str2Ptr(""),
		Rank:                 &emptyStr,
		StaffRank:            &emptyStr,
		LastName:             "山田",
		FirstName:            "花子",
		LastKana:             "ヤマダ",
		FirstKana:            "ハナコ",
		PostCode:             str2Ptr("1230045"),
		Address:              str2Ptr("東京都渋谷区恵比寿南1-1-1"),
		PhoneNumber:          str2Ptr("0337930220"),
		FaxNumber:            str2Ptr("0337930250"),
		MobileNumber:         str2Ptr("00000000000"),
		MailAddress:          str2Ptr("aa@hoge.co.jp"),
		MailAddress2:         &emptyStr,
		MailAddress3:         &emptyStr,
		CompanyName:          &emptyStr,
		DepartmentName:       &emptyStr,
		ManagerialPosition:   &emptyStr,
		Sex:                  "2",
		BirthDate:            str2Ptr("2000-05-28"),
		Mile:                 &emptyJSONNumber,
		Point:                json.Number("0"),
		PointExpireDate:      &emptyStr,
		LastComeDateTime:     &emptyStr,
		EntryDate:            "2018-08-19",
		LeaveDate:            &emptyStr,
		PointGivingUnitPrice: &emptyStr,
		PointGivingUnit:      &emptyStr,
		PinCode:              &emptyStr,
		PassportNo:           &emptyStr,
		Nationality:          &emptyStr,
		AlphabetName:         &emptyStr,
		MailReceiveFlag:      "1",
		Note:                 str2Ptr("スカパラのファンでそこからDBSSに初来店。初回は中目黒でルーズフィット黒を紹介購入。その後初回来店時に提案したシャギーニットをご購入。その後ルーズフィット黒2本目を取寄せ購入。\nサイズ2"),
		Note2:                &emptyStr,
		FavoriteList:         &emptyStr,
		BrowsingList:         &emptyStr,
		Status:               "0",
		StoreID:              num2Ptr("1"),
		InsDateTime:          "2018-08-19 15:40:50",
		UpdDateTime:          "2019-01-03 13:30:58",
	}

	if !reflect.DeepEqual(customers[0], &cus1) {
		t.Fatalf("failed test expected: %v\n got: %v", cus1, *customers[0])
	}

	cus2 := Customer{
		CustomerID:           "2",
		CustomerCode:         "02",
		CustomerNo:           &emptyStr,
		Rank:                 &emptyStr,
		StaffRank:            &emptyStr,
		LastName:             "田中",
		FirstName:            "太郎",
		LastKana:             "タナカ",
		FirstKana:            "タロウ",
		PostCode:             str2Ptr("4608501"),
		Address:              str2Ptr("愛知県名古屋市中区三の丸三丁目1番2号"),
		PhoneNumber:          str2Ptr(""),
		FaxNumber:            str2Ptr(""),
		MobileNumber:         str2Ptr("08000000000"),
		MailAddress:          str2Ptr("bbbb@hoge.com"),
		MailAddress2:         &emptyStr,
		MailAddress3:         &emptyStr,
		CompanyName:          &emptyStr,
		DepartmentName:       &emptyStr,
		ManagerialPosition:   &emptyStr,
		Sex:                  "1",
		BirthDate:            str2Ptr("1987-12-16"),
		Mile:                 &emptyJSONNumber,
		Point:                "0",
		PointExpireDate:      &emptyStr,
		LastComeDateTime:     &emptyStr,
		EntryDate:            "2018-08-19",
		LeaveDate:            &emptyStr,
		PointGivingUnitPrice: &emptyStr,
		PointGivingUnit:      &emptyStr,
		PinCode:              &emptyStr,
		PassportNo:           &emptyStr,
		Nationality:          &emptyStr,
		AlphabetName:         &emptyStr,
		MailReceiveFlag:      "1",
		Note:                 str2Ptr("2月にデニムシャツをご購入。\nZOZOでトラックフーディーのセットアップをご購入。\n5月に軍鶏アロハシャツとロングタンクトップをご購入。\n18AWの入荷連もいただいている。"),
		Note2:                &emptyStr,
		FavoriteList:         &emptyStr,
		BrowsingList:         &emptyStr,
		Status:               "0",
		StoreID:              num2Ptr("1"),
		InsDateTime:          "2018-08-19 15:43:03",
		UpdDateTime:          "2018-08-24 12:37:57",
	}
	if !reflect.DeepEqual(customers[1], &cus2) {
		t.Fatalf("failed test expected: %v\n got: %v", cus2, *customers[1])
	}
}
