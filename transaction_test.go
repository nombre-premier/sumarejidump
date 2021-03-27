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

func TestTransactionHead(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "transaction")
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
		File("./testdata/transaction_head.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{TRANSACTION_HEAD},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "TransactionHead.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	defer csvFile.Close()

	ths := []*TransactionHead{}

	if err := gocsv.UnmarshalFile(csvFile, &ths); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(ths) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(ths))
	}

	t1 := TransactionHead{
		TransactionHeadID:              json.Number("104"),
		TransactionDateTime:            "2018-09-02 13:34:34",
		TransactionHeadDivision:        "1",
		CancelDivision:                 "0",
		UnitNonDiscountsubtotal:        num2Ptr("26000"),
		UnitDiscountsubtotal:           num2Ptr("0"),
		Subtotal:                       num2Ptr("26000"),
		SubtotalForDiscount:            num2Ptr("26000"),
		SubtotalDiscountPrice:          num2Ptr("0"),
		SubtotalDiscountRate:           num2Ptr("0.00"),
		SubtotalDiscountDivision:       num2Ptr(""),
		PointDiscount:                  num2Ptr("0"),
		CouponDiscount:                 num2Ptr("0"),
		Total:                          num2Ptr("28080"),
		TaxExclude:                     json.Number("2080"),
		TaxInclude:                     json.Number("0"),
		RoundingDivision:               "00",
		RoundingPrice:                  num2Ptr("0"),
		CashTotal:                      num2Ptr("28080"),
		CreditTotal:                    num2Ptr("0"),
		Deposit:                        num2Ptr("28080"),
		DepositCash:                    num2Ptr("28080"),
		DepositCredit:                  num2Ptr("0"),
		Charge:                         num2Ptr("0"),
		ChangeDifference:               json.Number("0"),
		Amount:                         num2Ptr("1"),
		ReturnAmount:                   num2Ptr("0"),
		CostTotal:                      num2Ptr("100000.00000"),
		SalesHeadDivision:              "0",
		InTaxSalesTotal:                json.Number("0"),
		OutTaxSalesTotal:               json.Number("26000"),
		NonTaxSalesTotal:               json.Number("0"),
		NonSalesTargetTotal:            json.Number("0"),
		NonSalesTargetOutTaxTotal:      json.Number("0"),
		NonSalesTargetInTaxTotal:       json.Number("0"),
		NonSalesTargetTaxFreeTotal:     json.Number("0"),
		NonSalesTargetCostTotal:        json.Number("0.00000"),
		NonSalesTargetAmount:           json.Number("0"),
		NonSalesTargetReturnAmount:     json.Number("0"),
		NewPoint:                       json.Number("0"),
		SpendPoint:                     num2Ptr("0"),
		Point:                          num2Ptr("0"),
		TotalPoint:                     num2Ptr("0"),
		CurrentMile:                    num2Ptr(""),
		EarnMile:                       num2Ptr("0"),
		TotalMile:                      num2Ptr(""),
		AdjustmentMile:                 num2Ptr("0"),
		AdjustmentMileDivision:         str2Ptr(""),
		AdjustmentMileValue:            num2Ptr(""),
		StoreID:                        num2Ptr("4"),
		StoreCode:                      str2Ptr("dbssnagoya"),
		TerminalID:                     num2Ptr("4"),
		CustomerID:                     num2Ptr("41"),
		CustomerCode:                   str2Ptr("39"),
		TerminalTranID:                 str2Ptr("2"),
		TerminalTranDateTime:           "2018-09-02 13:34:32",
		SumDivision:                    "2",
		AdjustmentDateTime:             str2Ptr("2018-09-02 20:33:42"),
		SumDateTime:                    str2Ptr("2018-09-02 00:00:00"),
		CustomerRank:                   num2Ptr(""),
		CustomerGroupID:                num2Ptr(""),
		CustomerGroupID2:               num2Ptr(""),
		CustomerGroupID3:               num2Ptr(""),
		CustomerGroupID4:               num2Ptr(""),
		CustomerGroupID5:               num2Ptr(""),
		StaffID:                        num2Ptr("5"),
		StaffCode:                      str2Ptr(""),
		StaffName:                      str2Ptr("山田 太郎"),
		PaymentCount:                   str2Ptr(""),
		SlipNumber:                     str2Ptr(""),
		CancelSlipNumber:               str2Ptr(""),
		AuthNumber:                     str2Ptr(""),
		AuthDate:                       str2Ptr(""),
		CardCompany:                    str2Ptr(""),
		Memo:                           str2Ptr(""),
		ReceiptMemo:                    str2Ptr(""),
		PaymentMethodID1:               num2Ptr(""),
		PaymentMethodName1:             str2Ptr(""),
		DepositOthers1:                 num2Ptr(""),
		PaymentMethodID2:               num2Ptr(""),
		PaymentMethodName2:             str2Ptr(""),
		DepositOthers2:                 num2Ptr(""),
		PaymentMethodID3:               num2Ptr(""),
		PaymentMethodName3:             str2Ptr(""),
		DepositOthers3:                 num2Ptr(""),
		Carriage:                       str2Ptr(""),
		Commission:                     str2Ptr(""),
		GuestNumbers:                   num2Ptr("1"),
		TaxFreeSalesDivision:           "0",
		NetTaxFreeGeneralTaxInclude:    num2Ptr("0"),
		NetTaxFreeGeneralTaxExclude:    num2Ptr("0"),
		NetTaxFreeConsumableTaxInclude: num2Ptr("0"),
		NetTaxFreeConsumableTaxExclude: num2Ptr("0"),
		Tags:                           str2Ptr(""),
		PointGivingDivision:            str2Ptr("1"),
		PointGivingUnitPrice:           num2Ptr("1"),
		PointGivingUnit:                num2Ptr("0.00"),
		PointSpendDivision:             str2Ptr("1"),
		MileageDivision:                str2Ptr("0"),
		MileageLabel:                   str2Ptr("マイル"),
		CustomerPinCode:                str2Ptr(""),
		DisposeDivision:                "0",
		DisposeServerTransactionHeadID: num2Ptr("0"),
		CancelDateTime:                 str2Ptr(""),
		SellDivision:                   "1",
		TaxRate:                        num2Ptr("8.000"),
		TaxRounding:                    str2Ptr("1"),
		DiscountRoundingDivision:       "1",
		TransactionUUID:                str2Ptr("5358628720004"),
		GiftReceiptValidDays:           num2Ptr(""),
		ReceiptIssueNumberOfTimes:      json.Number("0"),
		PickupTransactionHeadID:        num2Ptr(""),
		PickUpDate:                     str2Ptr(""),
		PartPayment:                    num2Ptr(""),
		PartPaymentClass:               str2Ptr(""),
		LayawayServerTransactionHeadID: num2Ptr(""),
		DisabledEdit:                   str2Ptr(""),
		UpdDateTime:                    "2018-09-02 20:33:42",
	}

	if !reflect.DeepEqual(ths[0], &t1) {
		t.Fatalf("failed test expected: %v\n got: %v", t1, *ths[0])
	}

	t2 := TransactionHead{
		TransactionHeadID:              json.Number("1"),
		TransactionDateTime:            "2018-07-24 00:24:50",
		TransactionHeadDivision:        "1",
		CancelDivision:                 "0",
		UnitNonDiscountsubtotal:        num2Ptr("69000"),
		UnitDiscountsubtotal:           num2Ptr("0"),
		Subtotal:                       num2Ptr("69000"),
		SubtotalForDiscount:            num2Ptr("69000"),
		SubtotalDiscountPrice:          num2Ptr("0"),
		SubtotalDiscountRate:           num2Ptr("0.00"),
		SubtotalDiscountDivision:       num2Ptr(""),
		PointDiscount:                  num2Ptr("0"),
		CouponDiscount:                 num2Ptr("0"),
		Total:                          num2Ptr("74520"),
		TaxExclude:                     json.Number("5520"),
		TaxInclude:                     json.Number("0"),
		RoundingDivision:               "00",
		RoundingPrice:                  num2Ptr("0"),
		CashTotal:                      num2Ptr("74520"),
		CreditTotal:                    num2Ptr("0"),
		Deposit:                        num2Ptr("74520"),
		DepositCash:                    num2Ptr("74520"),
		DepositCredit:                  num2Ptr("0"),
		Charge:                         num2Ptr("0"),
		ChangeDifference:               json.Number("0"),
		Amount:                         num2Ptr("1"),
		ReturnAmount:                   num2Ptr("0"),
		CostTotal:                      num2Ptr("0.00000"),
		SalesHeadDivision:              "0",
		InTaxSalesTotal:                json.Number("0"),
		OutTaxSalesTotal:               json.Number("69000"),
		NonTaxSalesTotal:               json.Number("0"),
		NonSalesTargetTotal:            json.Number("0"),
		NonSalesTargetOutTaxTotal:      json.Number("0"),
		NonSalesTargetInTaxTotal:       json.Number("0"),
		NonSalesTargetTaxFreeTotal:     json.Number("0"),
		NonSalesTargetCostTotal:        json.Number("0.00000"),
		NonSalesTargetAmount:           json.Number("0"),
		NonSalesTargetReturnAmount:     json.Number("0"),
		NewPoint:                       json.Number("0"),
		SpendPoint:                     num2Ptr("0"),
		Point:                          num2Ptr("0"),
		TotalPoint:                     num2Ptr("0"),
		CurrentMile:                    num2Ptr(""),
		EarnMile:                       num2Ptr(""),
		TotalMile:                      num2Ptr(""),
		AdjustmentMile:                 num2Ptr(""),
		AdjustmentMileDivision:         str2Ptr(""),
		AdjustmentMileValue:            num2Ptr(""),
		StoreID:                        num2Ptr("2"),
		StoreCode:                      str2Ptr("dbssnakameguro"),
		TerminalID:                     num2Ptr("3"),
		CustomerID:                     num2Ptr(""),
		CustomerCode:                   str2Ptr(""),
		TerminalTranID:                 str2Ptr("1"),
		TerminalTranDateTime:           "2018-07-24 00:24:50",
		SumDivision:                    "2",
		AdjustmentDateTime:             str2Ptr(""),
		SumDateTime:                    str2Ptr("2018-07-23 00:00:00"),
		CustomerRank:                   num2Ptr(""),
		CustomerGroupID:                num2Ptr(""),
		CustomerGroupID2:               num2Ptr(""),
		CustomerGroupID3:               num2Ptr(""),
		CustomerGroupID4:               num2Ptr(""),
		CustomerGroupID5:               num2Ptr(""),
		StaffID:                        num2Ptr(""),
		StaffName:                      str2Ptr(""),
		StaffCode:                      str2Ptr(""),
		PaymentCount:                   str2Ptr(""),
		SlipNumber:                     str2Ptr(""),
		CancelSlipNumber:               str2Ptr(""),
		AuthNumber:                     str2Ptr(""),
		AuthDate:                       str2Ptr(""),
		CardCompany:                    str2Ptr(""),
		Memo:                           str2Ptr(""),
		ReceiptMemo:                    str2Ptr(""),
		PaymentMethodID1:               num2Ptr(""),
		PaymentMethodName1:             str2Ptr(""),
		DepositOthers1:                 num2Ptr(""),
		PaymentMethodID2:               num2Ptr(""),
		PaymentMethodName2:             str2Ptr(""),
		DepositOthers2:                 num2Ptr(""),
		PaymentMethodID3:               num2Ptr(""),
		PaymentMethodName3:             str2Ptr(""),
		DepositOthers3:                 num2Ptr(""),
		Carriage:                       str2Ptr(""),
		Commission:                     str2Ptr(""),
		GuestNumbers:                   num2Ptr("1"),
		TaxFreeSalesDivision:           "0",
		NetTaxFreeGeneralTaxInclude:    num2Ptr("0"),
		NetTaxFreeGeneralTaxExclude:    num2Ptr("0"),
		NetTaxFreeConsumableTaxInclude: num2Ptr("0"),
		NetTaxFreeConsumableTaxExclude: num2Ptr("0"),
		Tags:                           str2Ptr(""),
		PointGivingDivision:            str2Ptr("1"),
		PointGivingUnitPrice:           num2Ptr("1"),
		PointGivingUnit:                num2Ptr("0.00"),
		PointSpendDivision:             str2Ptr("1"),
		MileageDivision:                str2Ptr("0"),
		MileageLabel:                   str2Ptr("マイル"),
		CustomerPinCode:                str2Ptr(""),
		DisposeDivision:                "1",
		DisposeServerTransactionHeadID: num2Ptr("2"),
		CancelDateTime:                 str2Ptr(""),
		SellDivision:                   "1",
		TaxRate:                        num2Ptr("8.000"),
		TaxRounding:                    str2Ptr("1"),
		DiscountRoundingDivision:       "1",
		TransactionUUID:                str2Ptr("5323594900003"),
		GiftReceiptValidDays:           num2Ptr(""),
		ReceiptIssueNumberOfTimes:      json.Number("0"),
		PickupTransactionHeadID:        num2Ptr(""),
		PickUpDate:                     str2Ptr(""),
		PartPayment:                    num2Ptr(""),
		PartPaymentClass:               str2Ptr(""),
		LayawayServerTransactionHeadID: num2Ptr(""),
		DisabledEdit:                   str2Ptr(""),
		UpdDateTime:                    "2018-07-24 00:25:29",
	}

	if !reflect.DeepEqual(ths[1], &t2) {
		t.Fatalf("failed test expected: %v\n got: %v", t2, *ths[1])
	}
}

func TestTransactionDetail(t *testing.T) {
	defer gock.Off()

	dir, err := ioutil.TempDir("", "transaction_detail")
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
		File("./testdata/transaction_detail.json")

	c := SrConfig{
		ContractID:  "sumareji_id",
		AccessToken: "sumareji_token",
		EndPoint:    "https://webapi.smaregi.jp/access/",
		OutputDir:   dir,
		TableNames:  []string{TRANSACTION_DETAIL},
	}

	err = Main(c)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	csvFile, err := os.OpenFile(path.Join(dir, "TransactionDetail.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer csvFile.Close()

	tds := []*TransactionDetail{}

	if err := gocsv.UnmarshalFile(csvFile, &tds); err != nil { // Load clients from file
		t.Fatalf("failed test %#v", err)
	}

	if len(tds) != 2 {
		t.Fatalf("Number of data should be 2. got=%d", len(tds))
	}

	t1 := TransactionDetail{
		TransactionHeadID:            json.Number("1"),
		TransactionDateTime:          "2018-07-24 00:24:50",
		TransactionDetailID:          json.Number("1"),
		ParentTransactionDetailID:    num2Ptr(""),
		TransactionDetailDivision:    "1",
		ProductID:                    num2Ptr("2"),
		ProductCode:                  str2Ptr("2990000000002"),
		ProductName:                  str2Ptr("D181H223  Joined balmacaan denim coat"),
		TaxDivision:                  "1",
		Price:                        json.Number("69000"),
		SalesPrice:                   json.Number("69000"),
		UnitDiscountPrice:            json.Number("0"),
		UnitDiscountRate:             num2Ptr(""),
		UnitDiscountDivision:         num2Ptr(""),
		Cost:                         num2Ptr(""),
		Quantity:                     json.Number("1"),
		UnitNonDiscountSum:           json.Number("69000"),
		UnitDiscountSum:              json.Number("0"),
		UnitDiscountedSum:            json.Number("69000"),
		CostSum:                      json.Number("0.00000"),
		CategoryID:                   json.Number("2"),
		CategoryName:                 "コート",
		DiscriminationNo:             "1-008",
		SalesDivision:                "0",
		ProductDivision:              "0",
		PointNotApplicable:           "0",
		CalcDiscount:                 "1",
		TaxFreeDivision:              "0",
		TaxFreeCommodityPrice:        json.Number("0"),
		TaxFree:                      json.Number("0"),
		ProductBundleGroupID:         num2Ptr(""),
		DiscountPriceProportional:    json.Number("0"),
		DiscountPointProportional:    json.Number("0"),
		DiscountCouponProportional:   json.Number("0"),
		TaxIncludeProportional:       json.Number("0"),
		TaxExcludeProportional:       json.Number("5520"),
		ProductBundleProportional:    json.Number("0"),
		StaffDiscountProportional:    json.Number("0"),
		BargainDiscountProportional:  json.Number("0"),
		RoundingPriceProportional:    json.Number("0"),
		InventoryReservationDivision: "0",
		GroupCode:                    str2Ptr(""),
		UpdDateTime:                  "2018-07-24 00:25:29",
		ProductStaffDiscountRate:     num2Ptr(""),
		StaffRank:                    str2Ptr(""),
		StaffRankName:                str2Ptr(""),
		StaffDiscountRate:            num2Ptr(""),
		StaffDiscountDivision:        num2Ptr(""),
		ApplyStaffDiscountRate:       num2Ptr(""),
		ApplyStaffDiscountPrice:      num2Ptr(""),
		BargainID:                    num2Ptr(""),
		BargainName:                  str2Ptr(""),
		BargainDivision:              str2Ptr(""),
		BargainValue:                 num2Ptr(""),
		ApplyBargainValue:            num2Ptr(""),
		ApplyBargainDiscountPrice:    num2Ptr(""),
		TaxRate:                      json.Number("8.000"),
		StandardTaxRate:              json.Number("8.000"),
		ModifiedTaxRate:              num2Ptr(""),
		ReduceTaxID:                  num2Ptr(""),
		ReduceTaxName:                str2Ptr(""),
		ReduceTaxRate:                num2Ptr(""),
		Color:                        str2Ptr("87"),
		Size:                         str2Ptr("2"),
	}

	if !reflect.DeepEqual(tds[0], &t1) {
		t.Fatalf("failed test expected: %v\n got: %v", t1, *tds[0])
	}

	t2 := TransactionDetail{
		TransactionHeadID:            json.Number("2"),
		TransactionDateTime:          "2018-07-24 00:25:29",
		TransactionDetailID:          json.Number("1"),
		ParentTransactionDetailID:    num2Ptr(""),
		TransactionDetailDivision:    "2",
		ProductID:                    num2Ptr("2"),
		ProductCode:                  str2Ptr("2990000000002"),
		ProductName:                  str2Ptr("D181H223  Joined balmacaan denim coat"),
		TaxDivision:                  "1",
		Price:                        json.Number("69000"),
		SalesPrice:                   json.Number("69000"),
		UnitDiscountPrice:            json.Number("0"),
		UnitDiscountRate:             num2Ptr(""),
		UnitDiscountDivision:         num2Ptr(""),
		Cost:                         num2Ptr(""),
		Quantity:                     json.Number("1"),
		UnitNonDiscountSum:           json.Number("69000"),
		UnitDiscountSum:              json.Number("0"),
		UnitDiscountedSum:            json.Number("69000"),
		CostSum:                      json.Number("0.00000"),
		CategoryID:                   json.Number("2"),
		CategoryName:                 "コート",
		DiscriminationNo:             "1-009",
		SalesDivision:                "0",
		ProductDivision:              "0",
		PointNotApplicable:           "0",
		CalcDiscount:                 "1",
		TaxFreeDivision:              "0",
		TaxFreeCommodityPrice:        json.Number("0"),
		TaxFree:                      json.Number("0"),
		ProductBundleGroupID:         num2Ptr(""),
		DiscountPriceProportional:    json.Number("0"),
		DiscountPointProportional:    json.Number("0"),
		DiscountCouponProportional:   json.Number("0"),
		TaxIncludeProportional:       json.Number("0"),
		TaxExcludeProportional:       json.Number("5520"),
		ProductBundleProportional:    json.Number("0"),
		StaffDiscountProportional:    json.Number("0"),
		BargainDiscountProportional:  json.Number("0"),
		RoundingPriceProportional:    json.Number("0"),
		InventoryReservationDivision: "0",
		GroupCode:                    str2Ptr(""),
		UpdDateTime:                  "2018-07-24 00:25:29",
		ProductStaffDiscountRate:     num2Ptr(""),
		StaffRank:                    str2Ptr(""),
		StaffRankName:                str2Ptr(""),
		StaffDiscountRate:            num2Ptr(""),
		StaffDiscountDivision:        num2Ptr(""),
		ApplyStaffDiscountRate:       num2Ptr(""),
		ApplyStaffDiscountPrice:      num2Ptr(""),
		BargainID:                    num2Ptr(""),
		BargainName:                  str2Ptr(""),
		BargainDivision:              str2Ptr(""),
		BargainValue:                 num2Ptr(""),
		ApplyBargainValue:            num2Ptr(""),
		ApplyBargainDiscountPrice:    num2Ptr(""),
		TaxRate:                      json.Number("8.000"),
		StandardTaxRate:              json.Number("8.000"),
		ModifiedTaxRate:              num2Ptr(""),
		ReduceTaxID:                  num2Ptr(""),
		ReduceTaxName:                str2Ptr(""),
		ReduceTaxRate:                num2Ptr(""),
		Color:                        str2Ptr("87"),
		Size:                         str2Ptr("2"),
	}

	if !reflect.DeepEqual(tds[1], &t2) {
		t.Fatalf("failed test expected: %v\n got: %v", t2, *tds[1])
	}
}
