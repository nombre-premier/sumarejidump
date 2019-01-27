package main

import (
	"encoding/json"

)

type TransactionHead struct {
	TransactionHeadID              json.Number  `json:"transactionHeadId" csv:"transactionHeadId"`
	TransactionDateTime            string       `json:"transactionDateTime" csv:"transactionDateTime"`
	TransactionHeadDivision        string       `json:"transactionHeadDivision" csv:"transactionHeadDivision"`
	CancelDivision                 string       `json:"cancelDivision" csv:"cancelDivision"`
	UnitNonDiscountsubtotal        json.Number `json:"unitNonDiscountsubtotal" csv:"unitNonDiscountsubtotal"`
	UnitDiscountsubtotal           json.Number  `json:"unitDiscountsubtotal" csv:"unitDiscountsubtotal"`
	Subtotal                       json.Number  `json:"subtotal" csv:"subtotal"`
	SubtotalDiscountPrice          json.Number  `json:"subtotalDiscountPrice" csv:"subtotalDiscountPrice"`
	SubtotalDiscountRate           json.Number  `json:"subtotalDiscountRate" csv:"subtotalDiscountRate"`
	SubtotalDiscountDivision       *json.Number `json:"subtotalDiscountDivision" csv:"subtotalDiscountDivision"`
	PointDiscount                  json.Number  `json:"pointDiscount" csv:"pointDiscount"`
	Total                          json.Number  `json:"total" csv:"total"`
	TaxExclude                     json.Number  `json:"taxExclude" csv:"taxExclude"`
	TaxInclude                     json.Number  `json:"taxInclude" csv:"taxInclude"`
	RoundingDivision               string       `json:"roundingDivision" csv:"roundingDivision"`
	RoundingPrice                  json.Number  `json:"roundingPrice" csv:"roundingPrice"`
	CashTotal                      json.Number  `json:"cashTotal" csv:"cashTotal"`
	CreditTotal                    json.Number  `json:"creditTotal" csv:"creditTotal"`
	Deposit                        json.Number  `json:"deposit" csv:"deposit"`
	DepositCash                    json.Number  `json:"depositCash" csv:"depositCash"`
	DepositCredit                  json.Number  `json:"depositCredit" csv:"depositCredit"`
	Charge                         json.Number  `json:"charge" csv:"charge"`
	ChangeDifference               json.Number  `json:"changeDifference" csv:"changeDifference"`
	Amount                         json.Number  `json:"amount" csv:"amount"`
	ReturnAmount                   json.Number  `json:"returnAmount" csv:"returnAmount"`
	CostTotal                      json.Number  `json:"costTotal" csv:"costTotal"`
	SalesHeadDivision              string       `json:"salesHeadDivision" csv:"salesHeadDivision"`
	InTaxSalesTotal                json.Number  `json:"inTaxSalesTotal" csv:"inTaxSalesTotal"`
	OutTaxSalesTotal               json.Number  `json:"outTaxSalesTotal" csv:"outTaxSalesTotal"`
	NonTaxSalesTotal               json.Number  `json:"nonTaxSalesTotal" csv:"nonTaxSalesTotal"`
	NonSalesTargetTotal            json.Number  `json:"nonSalesTargetTotal" csv:"nonSalesTargetTotal"`
	NonSalesTargetOutTaxTotal      json.Number  `json:"nonSalesTargetOutTaxTotal" csv:"nonSalesTargetOutTaxTotal"`
	NonSalesTargetInTaxTotal       json.Number  `json:"nonSalesTargetInTaxTotal" csv:"nonSalesTargetInTaxTotal"`
	NonSalesTargetTaxFreeTotal     json.Number  `json:"nonSalesTargetTaxFreeTotal" csv:"nonSalesTargetTaxFreeTotal"`
	NonSalesTargetCostTotal        json.Number  `json:"nonSalesTargetCostTotal" csv:"nonSalesTargetCostTotal"`
	NonSalesTargetAmount           json.Number  `json:"nonSalesTargetAmount" csv:"nonSalesTargetAmount"`
	NonSalesTargetReturnAmount     json.Number  `json:"nonSalesTargetReturnAmount" csv:"nonSalesTargetReturnAmount"`
	NewPoint                       json.Number  `json:"newPoint" csv:"newPoint"`
	SpendPoint                     json.Number  `json:"spendPoint" csv:"spendPoint"`
	Point                          json.Number  `json:"point" csv:"point"`
	TotalPoint                     json.Number  `json:"totalPoint" csv:"totalPoint"`
	CurrentMile                    *json.Number `json:"currentMile" csv:"currentMile"`
	EarnMile                       *json.Number `json:"earnMile" csv:"earnMile"`
	TotalMile                      *json.Number `json:"totalMile" csv:"totalMile"`
	AdjustmentMile                 *json.Number `json:"adjustmentMile" csv:"adjustmentMile"`
	AdjustmentMileDivision         *string      `json:"adjustmentMileDivision" csv:"adjustmentMileDivision"`
	AdjustmentMileValue            *json.Number `json:"adjustmentMileValue" csv:"adjustmentMileValue"`
	StoreID                        json.Number  `json:"storeId" csv:"storeID"`
	StoreCode                      string       `json:"storeCode" csv:"storeCode"`
	TerminalID                     json.Number  `json:"terminalId" csv:"terminalID"`
	CustomerID                     *json.Number `json:"customerId" csv:"customerID"`
	CustomerCode                   *string      `json:"customerCode" csv:"customerCode"`
	TerminalTranID                 string       `json:"terminalTranId" csv:"terminalTranID"`
	TerminalTranDateTime           string       `json:"terminalTranDateTime" csv:"terminalTranDateTime"`
	SumDivision                    string       `json:"sumDivision" csv:"sumDivision"`
	AdjustmentDateTime             *string      `json:"adjustmentDateTime" csv:"adjustmentDateTime"`
	SumDateTime                    string       `json:"sumDateTime" csv:"sumDateTime"`
	CustomerRank                   *json.Number `json:"customerRank" csv:"customerRank"`
	CustomerGroupID                *json.Number `json:"customerGroupId" csv:"customerGroupId"`
	CustomerGroupID2               *json.Number `json:"customerGroupId2" csv:"customerGroupId2"`
	CustomerGroupID3               *json.Number `json:"customerGroupId3" csv:"customerGroupId3"`
	CustomerGroupID4               *json.Number `json:"customerGroupId4" csv:"customerGroupId4"`
	CustomerGroupID5               *json.Number `json:"customerGroupId5" csv:"customerGroupId5"`
	StaffID                        *json.Number `json:"staffId" csv:"staffID"`
	StaffName                      *string      `json:"staffName" csv:"staffName"`
	StaffCode                      *string      `json:"staffCode" csv:"staffCode"`
	PaymentCount                   *string      `json:"paymentCount" csv:"paymentCount"`
	SlipNumber                     *string      `json:"slipNumber" csv:"slipNumber"`
	CancelSlipNumber               *string      `json:"cancelSlipNumber" csv:"cancelSlipNumber"`
	AuthNumber                     *string      `json:"authNumber" csv:"authNumber"`
	AuthDate                       *string      `json:"authDate" csv:"authDate"`
	CardCompany                    *string      `json:"cardCompany" csv:"cardCompany"`
	Memo                           *string      `json:"memo" csv:"memo"`
	ReceiptMemo                    *string      `json:"receiptMemo" csv:"receiptMemo"`
	PaymentMethodID1               *json.Number `json:"paymentMethodId1" csv:"paymentMethodId1"`
	PaymentMethodName1             *string      `json:"paymentMethodName1" csv:"paymentMethodName1"`
	DepositOthers1                 *json.Number `json:"depositOthers1" csv:"depositOthers1"`
	PaymentMethodID2               *json.Number `json:"paymentMethodId2" csv:"paymentMethodId2"`
	PaymentMethodName2             *string      `json:"paymentMethodName2" csv:"paymentMethodName2"`
	DepositOthers2                 *json.Number `json:"depositOthers2" csv:"depositOthers2"`
	PaymentMethodID3               *json.Number `json:"paymentMethodId3" csv:"paymentMethodId3"`
	PaymentMethodName3             *string      `json:"paymentMethodName3" csv:"paymentMethodName3"`
	DepositOthers3                 *json.Number `json:"depositOthers3" csv:"depositOthers3"`
	Carriage                       *string      `json:"carriage" csv:"carriage"`
	Commission                     *string      `json:"commission" csv:"commission"`
	GuestNumbers                   json.Number  `json:"guestNumbers" csv:"guestNumbers"`
	TaxFreeSalesDivision           string       `json:"taxFreeSalesDivision" csv:"taxFreeSalesDivision"`
	NetTaxFreeGeneralTaxInclude    json.Number  `json:"netTaxFreeGeneralTaxInclude" csv:"netTaxFreeGeneralTaxInclude"`
	NetTaxFreeGeneralTaxExclude    json.Number  `json:"netTaxFreeGeneralTaxExclude" csv:"netTaxFreeGeneralTaxExclude"`
	NetTaxFreeConsumableTaxInclude json.Number  `json:"netTaxFreeConsumableTaxInclude" csv:"netTaxFreeConsumableTaxInclude"`
	NetTaxFreeConsumableTaxExclude json.Number  `json:"netTaxFreeConsumableTaxExclude" csv:"netTaxFreeConsumableTaxExclude"`
	Tags                           string       `json:"tags" csv:"tags"`
	PointGivingDivision            string       `json:"pointGivingDivision" csv:"pointGivingDivision"`
	PointGivingUnitPrice           json.Number  `json:"pointGivingUnitPrice" csv:"pointGivingUnitPrice"`
	PointGivingUnit                json.Number  `json:"pointGivingUnit" csv:"pointGivingUnit"`
	PointSpendDivision             string       `json:"pointSpendDivision" csv:"pointSpendDivision"`
	MileageDivision                string       `json:"mileageDivision" csv:"mileageDivision"`
	MileageLabel                   string       `json:"mileageLabel" csv:"mileageLabel"`
	CustomerPinCode                *string      `json:"customerPinCode" csv:"customerPinCode"`
	DisposeDivision                string       `json:"disposeDivision" csv:"disposeDivision"`
	DisposeServerTransactionHeadID json.Number  `json:"disposeServerTransactionHeadId" csv:"disposeServerTransactionHeadID"`
	CancelDateTime                 *string      `json:"cancelDateTime" csv:"cancelDateTime"`
	SellDivision                   string       `json:"sellDivision" csv:"sellDivision"`
	TaxRate                        json.Number  `json:"taxRate" csv:"taxRate"`
	TaxRounding                    string       `json:"taxRounding" csv:"taxRounding"`
	DiscountRoundingDivision       string       `json:"discountRoundingDivision" csv:"discountRoundingDivision"`
	TransactionUUID                string       `json:"transactionUuid" csv:"transactionUUID"`
	GiftReceiptValidDays           *json.Number `json:"giftReceiptValidDays" csv:"giftReceiptValidDays"`
	ReceiptIssueNumberOfTimes      json.Number  `json:"receiptIssueNumberOfTimes" csv:"receiptIssueNumberOfTimes"`
	PickupTransactionHeadID        *json.Number `json:"pickupTransactionHeadId" csv:"pickupTransactionHeadID"`
	PickUpDate                     *string      `json:"pickUpDate" csv:"pickUpDate"`
	PartPayment                    *json.Number `json:"partPayment" csv:"partPayment"`
	PartPaymentClass               *string      `json:"partPaymentClass" csv:"partPaymentClass"`
	LayawayServerTransactionHeadID *json.Number `json:"layawayServerTransactionHeadId" csv:"layawayServerTransactionHeadID"`
	DisabledEdit                   *string      `json:"disabledEdit" csv:"disabledEdit"`
	UpdDateTime                    string       `json:"updDateTime" csv:"updDateTime"`
}

type TransactionHeadCSV struct {
	*CSVHandler
	buf []TransactionHead
}

func NewTransactionHeadCSV(bufSize int, output string) (*TransactionHeadCSV, error) {
	buf := make([]TransactionHead, bufSize)
	handler, err := NewCSVHandler([]TransactionHead{}, output)
	if err != nil {
		return nil, err
	}

	return &TransactionHeadCSV{
		handler,
		buf,
	}, nil
}

func (thc *TransactionHeadCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &thc.buf[i])
	}
	thc.CSVWriter.Write(thc.buf[:len(resp.Result)])
	return thc.CSVWriter
}

type TransactionDetail struct {
	TransactionHeadID            json.Number  `json:"transactionHeadId" csv:"transactionHeadId"`
	TransactionDateTime          string       `json:"transactionDateTime" csv:"transactionDateTime"`
	TransactionDetailID          json.Number  `json:"transactionDetailId" csv:"transactionDetailId"`
	ParentTransactionDetailID    *json.Number `json:"parentTransactionDetailId" csv:"parentTransactionDetailId"`
	TransactionDetailDivision    string       `json:"transactionDetailDivision" csv:"transactionDetailDivision"`
	ProductID                    json.Number  `json:"productId" csv:"productId"`
	ProductCode                  string       `json:"productCode" csv:"productCode"`
	ProductName                  string       `json:"productName" csv:"productName"`
	TaxDivision                  string       `json:"taxDivision" csv:"taxDivision"`
	Price                        json.Number  `json:"price" csv:"price"`
	SalesPrice                   json.Number  `json:"salesPrice" csv:"salesPrice"`
	UnitDiscountPrice            json.Number  `json:"unitDiscountPrice" csv:"unitDiscountPrice"`
	UnitDiscountRate             json.Number  `json:"unitDiscountRate" csv:"unitDiscountRate"`
	UnitDiscountDivision         *json.Number `json:"unitDiscountDivision" csv:"unitDiscountDivision"`
	Cost                         *json.Number `json:"cost" csv:"cost"`
	Quantity                     json.Number  `json:"quantity" csv:"quantity"`
	UnitNonDiscountSum           json.Number  `json:"unitNonDiscountSum" csv:"unitNonDiscountSum"`
	UnitDiscountSum              json.Number  `json:"unitDiscountSum" csv:"unitDiscountSum"`
	UnitDiscountedSum            json.Number  `json:"unitDiscountedSum" csv:"unitDiscountedSum"`
	CostSum                      json.Number  `json:"costSum" csv:"costSum"`
	CategoryID                   json.Number  `json:"categoryId" csv:"categoryId"`
	CategoryName                 string       `json:"categoryName" csv:"categoryName"`
	DiscriminationNo             string       `json:"discriminationNo" csv:"discriminationNo"`
	SalesDivision                string       `json:"salesDivision" csv:"salesDivision"`
	ProductDivision              string       `json:"productDivision" csv:"productDivision"`
	PointNotApplicable           string       `json:"pointNotApplicable" csv:"pointNotApplicable"`
	TaxFreeDivision              string       `json:"taxFreeDivision" csv:"taxFreeDivision"`
	TaxFreeCommodityPrice        json.Number  `json:"taxFreeCommodityPrice" csv:"taxFreeCommodityPrice"`
	TaxFree                      json.Number  `json:"taxFree" csv:"taxFree"`
	ProductBundleGroupID         *json.Number `json:"productBundleGroupId" csv:"productBundleGroupId"`
	DiscountPriceProportional    json.Number  `json:"discountPriceProportional" csv:"discountPriceProportional"`
	DiscountPointProportional    json.Number  `json:"discountPointProportional" csv:"discountPointProportional"`
	TaxIncludeProportional       json.Number  `json:"taxIncludeProportional" csv:"taxIncludeProportional"`
	TaxExcludeProportional       json.Number  `json:"taxExcludeProportional" csv:"taxExcludeProportional"`
	ProductBundleProportional    json.Number  `json:"productBundleProportional" csv:"productBundleProportional"`
	StaffDiscountProportional    json.Number  `json:"staffDiscountProportional" csv:"staffDiscountProportional"`
	BargainDiscountProportional  json.Number  `json:"bargainDiscountProportional" csv:"bargainDiscountProportional"`
	RoundingPriceProportional    json.Number  `json:"roundingPriceProportional" csv:"roundingPriceProportional"`
	InventoryReservationDivision string       `json:"inventoryReservationDivision" csv:"inventoryReservationDivision"`
	GroupCode                    string       `json:"groupCode" csv:"groupCode"`
	UpdDateTime                  string       `json:"updDateTime" csv:"updDateTime"`
	ProductStaffDiscountRate     *json.Number `json:"productStaffDiscountRate" csv:"productStaffDiscountRate"`
	StaffRank                    *string      `json:"staffRank" csv:"staffRank"`
	StaffRankName                *string      `json:"staffRankName" csv:"staffRankName"`
	StaffDiscountRate            *json.Number `json:"staffDiscountRate" csv:"staffDiscountRate"`
	StaffDiscountDivision        *json.Number `json:"staffDiscountDivision" csv:"staffDiscountDivision"`
	ApplyStaffDiscountRate       *json.Number `json:"applyStaffDiscountRate" csv:"applyStaffDiscountRate"`
	ApplyStaffDiscountPrice      *json.Number `json:"applyStaffDiscountPrice" csv:"applyStaffDiscountPrice"`
	BargainID                    *json.Number `json:"bargainId" csv:"bargainID"`
	BargainName                  *string      `json:"bargainName" csv:"bargainName"`
	BargainDivision              *string      `json:"bargainDivision" csv:"bargainDivision"`
	BargainValue                 *json.Number `json:"bargainValue" csv:"bargainValue"`
	ApplyBargainValue            *json.Number `json:"applyBargainValue" csv:"applyBargainValue"`
	ApplyBargainDiscountPrice    *json.Number `json:"applyBargainDiscountPrice" csv:"applyBargainDiscountPrice"`
	TaxRate                      json.Number  `json:"taxRate" csv:"taxRate"`
	StandardTaxRate              json.Number  `json:"standardTaxRate" csv:"standardTaxRate"`
	ModifiedTaxRate              *json.Number `json:"modifiedTaxRate" csv:"modifiedTaxRate"`
	Color                        string       `json:"color" csv:"color"`
	Size                         string       `json:"size" csv:"size"`
}

type TransactionDetailCSV struct {
	*CSVHandler
	buf []TransactionDetail
}

func NewTransactionDetailCSV(bufSize int, output string) (*TransactionDetailCSV, error) {
	buf := make([]TransactionDetail, bufSize)
	handler, err := NewCSVHandler([]TransactionDetail{}, output)
	if err != nil {
		return nil, err
	}

	return &TransactionDetailCSV{
		handler,
		buf,
	}, nil
}

func (tdc *TransactionDetailCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &tdc.buf[i])
	}
	tdc.CSVWriter.Write(tdc.buf[:len(resp.Result)])
	return tdc.CSVWriter
}
