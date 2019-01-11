package main

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type TransactionHead struct {
	TransactionHeadID              json.Number      `json:"transactionHeadId" csv:"transactionHeadId"`
	TransactionDateTime            string           `json:"transactionDateTime" csv:"transactionDateTime"`
	TransactionHeadDivision        string           `json:"transactionHeadDivision" csv:"transactionHeadDivision"`
	CancelDivision                 string           `json:"cancelDivision" csv:"cancelDivision"`
	UnitNonDiscountsubtotal        decimal.Decimal  `json:"unitNonDiscountsubtotal" csv:"unitNonDiscountsubtotal"`
	UnitDiscountsubtotal           decimal.Decimal  `json:"unitDiscountsubtotal" csv:"unitDiscountsubtotal"`
	Subtotal                       decimal.Decimal  `json:"subtotal" csv:"subtotal"`
	SubtotalDiscountPrice          decimal.Decimal  `json:"subtotalDiscountPrice" csv:"subtotalDiscountPrice"`
	SubtotalDiscountRate           decimal.Decimal  `json:"subtotalDiscountRate" csv:"subtotalDiscountRate"`
	SubtotalDiscountDivision       *json.Number     `json:"subtotalDiscountDivision" csv:"subtotalDiscountDivision"`
	PointDiscount                  decimal.Decimal  `json:"pointDiscount" csv:"pointDiscount"`
	Total                          decimal.Decimal  `json:"total" csv:"total"`
	TaxExclude                     decimal.Decimal  `json:"taxExclude" csv:"taxExclude"`
	TaxInclude                     decimal.Decimal  `json:"taxInclude" csv:"taxInclude"`
	RoundingDivision               string           `json:"roundingDivision" csv:"roundingDivision"`
	RoundingPrice                  decimal.Decimal  `json:"roundingPrice" csv:"roundingPrice"`
	CashTotal                      decimal.Decimal  `json:"cashTotal" csv:"cashTotal"`
	CreditTotal                    decimal.Decimal  `json:"creditTotal" csv:"creditTotal"`
	Deposit                        decimal.Decimal  `json:"deposit" csv:"deposit"`
	DepositCash                    decimal.Decimal  `json:"depositCash" csv:"depositCash"`
	DepositCredit                  decimal.Decimal  `json:"depositCredit" csv:"depositCredit"`
	Charge                         decimal.Decimal  `json:"charge" csv:"charge"`
	ChangeDifference               decimal.Decimal  `json:"changeDifference" csv:"changeDifference"`
	Amount                         json.Number      `json:"amount" csv:"amount"`
	ReturnAmount                   json.Number      `json:"returnAmount" csv:"returnAmount"`
	CostTotal                      decimal.Decimal  `json:"costTotal" csv:"costTotal"`
	SalesHeadDivision              string           `json:"salesHeadDivision" csv:"salesHeadDivision"`
	InTaxSalesTotal                decimal.Decimal  `json:"inTaxSalesTotal" csv:"inTaxSalesTotal"`
	OutTaxSalesTotal               decimal.Decimal  `json:"outTaxSalesTotal" csv:"outTaxSalesTotal"`
	NonTaxSalesTotal               decimal.Decimal  `json:"nonTaxSalesTotal" csv:"nonTaxSalesTotal"`
	NonSalesTargetTotal            decimal.Decimal  `json:"nonSalesTargetTotal" csv:"nonSalesTargetTotal"`
	NonSalesTargetOutTaxTotal      decimal.Decimal  `json:"nonSalesTargetOutTaxTotal" csv:"nonSalesTargetOutTaxTotal"`
	NonSalesTargetInTaxTotal       decimal.Decimal  `json:"nonSalesTargetInTaxTotal" csv:"nonSalesTargetInTaxTotal"`
	NonSalesTargetTaxFreeTotal     decimal.Decimal  `json:"nonSalesTargetTaxFreeTotal" csv:"nonSalesTargetTaxFreeTotal"`
	NonSalesTargetCostTotal        decimal.Decimal  `json:"nonSalesTargetCostTotal" csv:"nonSalesTargetCostTotal"`
	NonSalesTargetAmount           json.Number      `json:"nonSalesTargetAmount" csv:"nonSalesTargetAmount"`
	NonSalesTargetReturnAmount     json.Number      `json:"nonSalesTargetReturnAmount" csv:"nonSalesTargetReturnAmount"`
	NewPoint                       decimal.Decimal  `json:"newPoint" csv:"newPoint"`
	SpendPoint                     decimal.Decimal  `json:"spendPoint" csv:"spendPoint"`
	Point                          decimal.Decimal  `json:"point" csv:"point"`
	TotalPoint                     decimal.Decimal  `json:"totalPoint" csv:"totalPoint"`
	CurrentMile                    *decimal.Decimal `json:"currentMile" csv:"currentMile"`
	EarnMile                       *decimal.Decimal `json:"earnMile" csv:"earnMile"`
	TotalMile                      *decimal.Decimal `json:"totalMile" csv:"totalMile"`
	AdjustmentMile                 *decimal.Decimal `json:"adjustmentMile" csv:"adjustmentMile"`
	AdjustmentMileDivision         *string          `json:"adjustmentMileDivision" csv:"adjustmentMileDivision"`
	AdjustmentMileValue            *json.Number     `json:"adjustmentMileValue" csv:"adjustmentMileValue"`
	StoreID                        json.Number      `json:"storeId" csv:"storeID"`
	StoreCode                      string           `json:"storeCode" csv:"storeCode"`
	TerminalID                     json.Number      `json:"terminalId" csv:"terminalID"`
	CustomerID                     *json.Number     `json:"customerId" csv:"customerID"`
	CustomerCode                   *string          `json:"customerCode" csv:"customerCode"`
	TerminalTranID                 string           `json:"terminalTranId" csv:"terminalTranID"`
	TerminalTranDateTime           string           `json:"terminalTranDateTime" csv:"terminalTranDateTime"`
	SumDivision                    string           `json:"sumDivision" csv:"sumDivision"`
	AdjustmentDateTime             *string          `json:"adjustmentDateTime" csv:"adjustmentDateTime"`
	SumDateTime                    string           `json:"sumDateTime" csv:"sumDateTime"`
	CustomerRank                   *json.Number     `json:"customerRank" csv:"customerRank"`
	CustomerGroupID                *json.Number     `json:"customerGroupId" csv:"customerGroupId"`
	CustomerGroupID2               *json.Number     `json:"customerGroupId2" csv:"customerGroupId2"`
	CustomerGroupID3               *json.Number     `json:"customerGroupId3" csv:"customerGroupId3"`
	CustomerGroupID4               *json.Number     `json:"customerGroupId4" csv:"customerGroupId4"`
	CustomerGroupID5               *json.Number     `json:"customerGroupId5" csv:"customerGroupId5"`
	StaffID                        *json.Number     `json:"staffId" csv:"staffID"`
	StaffName                      *string          `json:"staffName" csv:"staffName"`
	StaffCode                      *string          `json:"staffCode" csv:"staffCode"`
	PaymentCount                   *string          `json:"paymentCount" csv:"paymentCount"`
	SlipNumber                     *string          `json:"slipNumber" csv:"slipNumber"`
	CancelSlipNumber               *string          `json:"cancelSlipNumber" csv:"cancelSlipNumber"`
	AuthNumber                     *string          `json:"authNumber" csv:"authNumber"`
	AuthDate                       *string          `json:"authDate" csv:"authDate"`
	CardCompany                    *string          `json:"cardCompany" csv:"cardCompany"`
	Memo                           *string          `json:"memo" csv:"memo"`
	ReceiptMemo                    *string          `json:"receiptMemo" csv:"receiptMemo"`
	PaymentMethodID1               *json.Number     `json:"paymentMethodId1" csv:"paymentMethodId1"`
	PaymentMethodName1             *string          `json:"paymentMethodName1" csv:"paymentMethodName1"`
	DepositOthers1                 *decimal.Decimal `json:"depositOthers1" csv:"depositOthers1"`
	PaymentMethodID2               *json.Number     `json:"paymentMethodId2" csv:"paymentMethodId2"`
	PaymentMethodName2             *string          `json:"paymentMethodName2" csv:"paymentMethodName2"`
	DepositOthers2                 *decimal.Decimal `json:"depositOthers2" csv:"depositOthers2"`
	PaymentMethodID3               *json.Number     `json:"paymentMethodId3" csv:"paymentMethodId3"`
	PaymentMethodName3             *string          `json:"paymentMethodName3" csv:"paymentMethodName3"`
	DepositOthers3                 *decimal.Decimal `json:"depositOthers3" csv:"depositOthers3"`
	Carriage                       *string          `json:"carriage" csv:"carriage"`
	Commission                     *string          `json:"commission" csv:"commission"`
	GuestNumbers                   json.Number      `json:"guestNumbers" csv:"guestNumbers"`
	TaxFreeSalesDivision           string           `json:"taxFreeSalesDivision" csv:"taxFreeSalesDivision"`
	NetTaxFreeGeneralTaxInclude    decimal.Decimal  `json:"netTaxFreeGeneralTaxInclude" csv:"netTaxFreeGeneralTaxInclude"`
	NetTaxFreeGeneralTaxExclude    decimal.Decimal  `json:"netTaxFreeGeneralTaxExclude" csv:"netTaxFreeGeneralTaxExclude"`
	NetTaxFreeConsumableTaxInclude decimal.Decimal  `json:"netTaxFreeConsumableTaxInclude" csv:"netTaxFreeConsumableTaxInclude"`
	NetTaxFreeConsumableTaxExclude decimal.Decimal  `json:"netTaxFreeConsumableTaxExclude" csv:"netTaxFreeConsumableTaxExclude"`
	Tags                           string           `json:"tags" csv:"tags"`
	PointGivingDivision            string           `json:"pointGivingDivision" csv:"pointGivingDivision"`
	PointGivingUnitPrice           json.Number      `json:"pointGivingUnitPrice" csv:"pointGivingUnitPrice"`
	PointGivingUnit                decimal.Decimal  `json:"pointGivingUnit" csv:"pointGivingUnit"`
	PointSpendDivision             string           `json:"pointSpendDivision" csv:"pointSpendDivision"`
	MileageDivision                string           `json:"mileageDivision" csv:"mileageDivision"`
	MileageLabel                   string           `json:"mileageLabel" csv:"mileageLabel"`
	CustomerPinCode                *string          `json:"customerPinCode" csv:"customerPinCode"`
	DisposeDivision                string           `json:"disposeDivision" csv:"disposeDivision"`
	DisposeServerTransactionHeadID json.Number      `json:"disposeServerTransactionHeadId" csv:"disposeServerTransactionHeadID"`
	CancelDateTime                 *string          `json:"cancelDateTime" csv:"cancelDateTime"`
	SellDivision                   string           `json:"sellDivision" csv:"sellDivision"`
	TaxRate                        decimal.Decimal  `json:"taxRate" csv:"taxRate"`
	TaxRounding                    string           `json:"taxRounding" csv:"taxRounding"`
	DiscountRoundingDivision       string           `json:"discountRoundingDivision" csv:"discountRoundingDivision"`
	TransactionUUID                string           `json:"transactionUuid" csv:"transactionUUID"`
	GiftReceiptValidDays           *json.Number     `json:"giftReceiptValidDays" csv:"giftReceiptValidDays"`
	ReceiptIssueNumberOfTimes      json.Number      `json:"receiptIssueNumberOfTimes" csv:"receiptIssueNumberOfTimes"`
	PickupTransactionHeadID        *json.Number     `json:"pickupTransactionHeadId" csv:"pickupTransactionHeadID"`
	PickUpDate                     *string          `json:"pickUpDate" csv:"pickUpDate"`
	PartPayment                    *decimal.Decimal `json:"partPayment" csv:"partPayment"`
	PartPaymentClass               *string          `json:"partPaymentClass" csv:"partPaymentClass"`
	LayawayServerTransactionHeadID *json.Number     `json:"layawayServerTransactionHeadId" csv:"layawayServerTransactionHeadID"`
	DisabledEdit                   *string          `json:"disabledEdit" csv:"disabledEdit"`
	UpdDateTime                    string           `json:"updDateTime" csv:"updDateTime"`
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
	TransactionHeadID            json.Number      `json:"transactionHeadId" csv:"transactionHeadId"`
	TransactionDateTime          string           `json:"transactionDateTime" csv:"transactionDateTime"`
	TransactionDetailID          json.Number      `json:"transactionDetailId" csv:"transactionDetailId"`
	ParentTransactionDetailID    *json.Number     `json:"parentTransactionDetailId" csv:"parentTransactionDetailId"`
	TransactionDetailDivision    string           `json:"transactionDetailDivision" csv:"transactionDetailDivision"`
	ProductID                    json.Number      `json:"productId" csv:"productId"`
	ProductCode                  string           `json:"productCode" csv:"productCode"`
	ProductName                  string           `json:"productName" csv:"productName"`
	TaxDivision                  string           `json:"taxDivision" csv:"taxDivision"`
	Price                        decimal.Decimal  `json:"price" csv:"price"`
	SalesPrice                   decimal.Decimal  `json:"salesPrice" csv:"salesPrice"`
	UnitDiscountPrice            decimal.Decimal  `json:"unitDiscountPrice" csv:"unitDiscountPrice"`
	UnitDiscountRate             decimal.Decimal  `json:"unitDiscountRate" csv:"unitDiscountRate"`
	UnitDiscountDivision         *json.Number     `json:"unitDiscountDivision" csv:"unitDiscountDivision"`
	Cost                         *decimal.Decimal `json:"cost" csv:"cost"`
	Quantity                     json.Number      `json:"quantity" csv:"quantity"`
	UnitNonDiscountSum           decimal.Decimal  `json:"unitNonDiscountSum" csv:"unitNonDiscountSum"`
	UnitDiscountSum              decimal.Decimal  `json:"unitDiscountSum" csv:"unitDiscountSum"`
	UnitDiscountedSum            decimal.Decimal  `json:"unitDiscountedSum" csv:"unitDiscountedSum"`
	CostSum                      decimal.Decimal  `json:"costSum" csv:"costSum"`
	CategoryID                   json.Number      `json:"categoryId" csv:"categoryId"`
	CategoryName                 string           `json:"categoryName" csv:"categoryName"`
	DiscriminationNo             string           `json:"discriminationNo" csv:"discriminationNo"`
	SalesDivision                string           `json:"salesDivision" csv:"salesDivision"`
	ProductDivision              string           `json:"productDivision" csv:"productDivision"`
	PointNotApplicable           string           `json:"pointNotApplicable" csv:"pointNotApplicable"`
	TaxFreeDivision              string           `json:"taxFreeDivision" csv:"taxFreeDivision"`
	TaxFreeCommodityPrice        decimal.Decimal  `json:"taxFreeCommodityPrice" csv:"taxFreeCommodityPrice"`
	TaxFree                      decimal.Decimal  `json:"taxFree" csv:"taxFree"`
	ProductBundleGroupID         *json.Number     `json:"productBundleGroupId" csv:"productBundleGroupId"`
	DiscountPriceProportional    decimal.Decimal  `json:"discountPriceProportional" csv:"discountPriceProportional"`
	DiscountPointProportional    decimal.Decimal  `json:"discountPointProportional" csv:"discountPointProportional"`
	TaxIncludeProportional       decimal.Decimal  `json:"taxIncludeProportional" csv:"taxIncludeProportional"`
	TaxExcludeProportional       decimal.Decimal  `json:"taxExcludeProportional" csv:"taxExcludeProportional"`
	ProductBundleProportional    decimal.Decimal  `json:"productBundleProportional" csv:"productBundleProportional"`
	StaffDiscountProportional    decimal.Decimal  `json:"staffDiscountProportional" csv:"staffDiscountProportional"`
	BargainDiscountProportional  decimal.Decimal  `json:"bargainDiscountProportional" csv:"bargainDiscountProportional"`
	RoundingPriceProportional    decimal.Decimal  `json:"roundingPriceProportional" csv:"roundingPriceProportional"`
	InventoryReservationDivision string           `json:"inventoryReservationDivision" csv:"inventoryReservationDivision"`
	GroupCode                    string           `json:"groupCode" csv:"groupCode"`
	UpdDateTime                  string           `json:"updDateTime" csv:"updDateTime"`
	ProductStaffDiscountRate     *json.Number     `json:"productStaffDiscountRate" csv:"productStaffDiscountRate"`
	StaffRank                    *string          `json:"staffRank" csv:"staffRank"`
	StaffRankName                *string          `json:"staffRankName" csv:"staffRankName"`
	StaffDiscountRate            *decimal.Decimal `json:"staffDiscountRate" csv:"staffDiscountRate"`
	StaffDiscountDivision        *json.Number     `json:"staffDiscountDivision" csv:"staffDiscountDivision"`
	ApplyStaffDiscountRate       *decimal.Decimal `json:"applyStaffDiscountRate" csv:"applyStaffDiscountRate"`
	ApplyStaffDiscountPrice      *decimal.Decimal `json:"applyStaffDiscountPrice" csv:"applyStaffDiscountPrice"`
	BargainID                    *json.Number     `json:"bargainId" csv:"bargainID"`
	BargainName                  *string          `json:"bargainName" csv:"bargainName"`
	BargainDivision              *string          `json:"bargainDivision" csv:"bargainDivision"`
	BargainValue                 *decimal.Decimal `json:"bargainValue" csv:"bargainValue"`
	ApplyBargainValue            *decimal.Decimal `json:"applyBargainValue" csv:"applyBargainValue"`
	ApplyBargainDiscountPrice    *decimal.Decimal `json:"applyBargainDiscountPrice" csv:"applyBargainDiscountPrice"`
	TaxRate                      decimal.Decimal  `json:"taxRate" csv:"taxRate"`
	StandardTaxRate              decimal.Decimal  `json:"standardTaxRate" csv:"standardTaxRate"`
	ModifiedTaxRate              *decimal.Decimal `json:"modifiedTaxRate" csv:"modifiedTaxRate"`
	Color                        string           `json:"color" csv:"color"`
	Size                         string           `json:"size" csv:"size"`
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
