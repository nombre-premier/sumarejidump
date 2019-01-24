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

func TestDailySum(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "daily_sum")
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
		File("./testdata/daily_sum.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{DAILY_SUM},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "DailySum.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	ds := []*DailySum{}

	if err := gocsv.UnmarshalFile(csvFile, &ds); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(ds) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(ds))
	}

	emptyStr := ""

	ds1 := DailySum{
		SumDate:                    "2018-07-23",
		StoreID:                    json.Number("2"),
		CashDrawerID:               json.Number("0"),
		Status:                     "2",
		SalesTotal:                 json.Number("0"),
		Discount:                   json.Number("0"),
		PointDiscount:              json.Number("0"),
		TaxExcludeReceive:          json.Number("0"),
		NonSalesTargetTotal:        json.Number("0"),
		Total:                      json.Number("0"),
		TotalExcludTax:             json.Number("0"),
		InTaxSalesTotal:            json.Number("0"),
		TaxInclude:                 json.Number("0"),
		OutTaxSalesTotal:           json.Number("0"),
		TaxExclude:                 json.Number("0"),
		TaxTotal:                   json.Number("0"),
		NonTaxSalesTotal:           json.Number("0"),
		TaxFreeTotal:               json.Number("0"),
		NonSalesTargetTaxFreeTotal: json.Number("0"),
		CostTotal:                  json.Number("0.00000"),
		GrossMargin:                json.Number("0"),
		Amount:                     json.Number("1"),
		TransactionCount:           json.Number("2"),
		ReturnAmount:               json.Number("1"),
		Carriage:                   json.Number("0"),
		Commission:                 json.Number("0"),
		PreparationCash:            json.Number("0"),
		CashSales:                  json.Number("0"),
		CreditSales:                json.Number("0"),
		// OtherSaleList is ignored right now
		OtherSalseList:                   nil,
		ChangeDifference:                 json.Number("0"),
		PartPayment:                      json.Number("0"),
		PartPaymentCash:                  json.Number("0"),
		PartPaymentCredit:                json.Number("0"),
		ReceivedDepositCash:              json.Number("0"),
		ReceivedDepositCashTotal:         json.Number("0"),
		ReceivedDepositCreditTotal:       json.Number("0"),
		PartPaymentCancel:                json.Number("0"),
		PartPaymentCashCancel:            json.Number("0"),
		PartPaymentCreditCancel:          json.Number("0"),
		Deposit:                          json.Number("0"),
		ReturnDeposit:                    json.Number("0"),
		Receipt:                          json.Number("0"),
		Payment:                          json.Number("0"),
		NonSalesCashTotal:                json.Number("0"),
		NonSalesCreditTotal:              json.Number("0"),
		NonSalesOtherTotal:               json.Number("0"),
		NonSalesTaxFreeTotal:             json.Number("0"),
		CalculateBalance:                 json.Number("0"),
		RealBalance:                      json.Number("0"),
		Difference:                       json.Number("0"),
		Saving:                           json.Number("0"),
		CarryOver:                        json.Number("0"),
		TenThousandYen:                   json.Number("0"),
		FiveThousandYen:                  json.Number("0"),
		TwoThousandYen:                   json.Number("0"),
		OneThousandYen:                   json.Number("0"),
		FiveHundredYen:                   json.Number("0"),
		OneHundredYen:                    json.Number("0"),
		FiftyYen:                         json.Number("0"),
		TenYen:                           json.Number("0"),
		FiveYen:                          json.Number("0"),
		OneYen:                           json.Number("0"),
		Comment:                          &emptyStr,
		InsDateTime:                      "2018-07-24 00:24:50",
		UpdDateTime:                      "2018-07-24 00:25:29",
		SalesTotalNonSalesTargetDivision: "1",
		TotalTaxFreeDivision:             "1",
	}

	if !reflect.DeepEqual(ds[0], &ds1) {
		t.Fatalf("failed test expected: %v\n got: %v", ds1, *ds[0])
	}

	ds2 := DailySum{
		SumDate:                          "2018-07-26",
		StoreID:                          json.Number("1"),
		CashDrawerID:                     json.Number("0"),
		Status:                           "2",
		SalesTotal:                       json.Number("0"),
		Discount:                         json.Number("0"),
		PointDiscount:                    json.Number("0"),
		TaxExcludeReceive:                json.Number("0"),
		NonSalesTargetTotal:              json.Number("0"),
		Total:                            json.Number("0"),
		TotalExcludTax:                   json.Number("0"),
		InTaxSalesTotal:                  json.Number("0"),
		TaxInclude:                       json.Number("0"),
		OutTaxSalesTotal:                 json.Number("0"),
		TaxExclude:                       json.Number("0"),
		TaxTotal:                         json.Number("0"),
		NonTaxSalesTotal:                 json.Number("0"),
		TaxFreeTotal:                     json.Number("0"),
		NonSalesTargetTaxFreeTotal:       json.Number("0"),
		CostTotal:                        json.Number("0.00000"),
		GrossMargin:                      json.Number("0"),
		Amount:                           json.Number("7"),
		TransactionCount:                 json.Number("14"),
		ReturnAmount:                     json.Number("7"),
		Carriage:                         json.Number("0"),
		Commission:                       json.Number("0"),
		PreparationCash:                  json.Number("0"),
		CashSales:                        json.Number("0"),
		CreditSales:                      json.Number("0"),
		OtherSalseList:                   nil,
		ChangeDifference:                 json.Number("0"),
		PartPayment:                      json.Number("0"),
		PartPaymentCash:                  json.Number("0"),
		PartPaymentCredit:                json.Number("0"),
		ReceivedDepositCash:              json.Number("0"),
		ReceivedDepositCashTotal:         json.Number("0"),
		ReceivedDepositCreditTotal:       json.Number("0"),
		PartPaymentCancel:                json.Number("0"),
		PartPaymentCashCancel:            json.Number("0"),
		PartPaymentCreditCancel:          json.Number("0"),
		Deposit:                          json.Number("0"),
		ReturnDeposit:                    json.Number("0"),
		Receipt:                          json.Number("0"),
		Payment:                          json.Number("0"),
		NonSalesCashTotal:                json.Number("0"),
		NonSalesCreditTotal:              json.Number("0"),
		NonSalesOtherTotal:               json.Number("0"),
		NonSalesTaxFreeTotal:             json.Number("0"),
		CalculateBalance:                 json.Number("0"),
		RealBalance:                      json.Number("0"),
		Difference:                       json.Number("0"),
		Saving:                           json.Number("0"),
		CarryOver:                        json.Number("0"),
		TenThousandYen:                   json.Number("0"),
		FiveThousandYen:                  json.Number("0"),
		TwoThousandYen:                   json.Number("0"),
		OneThousandYen:                   json.Number("0"),
		FiveHundredYen:                   json.Number("0"),
		OneHundredYen:                    json.Number("0"),
		FiftyYen:                         json.Number("0"),
		TenYen:                           json.Number("0"),
		FiveYen:                          json.Number("0"),
		OneYen:                           json.Number("0"),
		Comment:                          &emptyStr,
		InsDateTime:                      "2018-07-26 11:19:31",
		UpdDateTime:                      "2018-07-26 12:14:08",
		SalesTotalNonSalesTargetDivision: "1",
		TotalTaxFreeDivision:             "1",
	}
	if !reflect.DeepEqual(ds[1], &ds2) {
		t.Fatalf("failed test expected: %v\n got: %v", ds2, *ds[1])
	}
}
