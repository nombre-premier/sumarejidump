package main

import (
	"encoding/json"
)

type DailySum struct {
	SumDate                          string       `json:"sumDate" csv:"sumDate"`
	StoreID                          json.Number  `json:"storeId" csv:"storeId"`
	CashDrawerID                     json.Number  `json:"cashDrawerId" csv:"cashDrawerId"`
	Status                           string       `json:"status" csv:"status"`
	SalesTotal                       json.Number  `json:"salesTotal" csv:"salesTotal"`
	Discount                         json.Number  `json:"discount" csv:"discount"`
	PointDiscount                    json.Number  `json:"pointDiscount" csv:"pointDiscount"`
	TaxExcludeReceive                json.Number  `json:"taxExcludeReceive" csv:"taxExcludeReceive"`
	NonSalesTargetTotal              json.Number  `json:"nonSalesTargetTotal" csv:"nonSalesTargetTotal"`
	Total                            json.Number  `json:"total" csv:"total"`
	TotalExcludTax                   json.Number  `json:"totalExcludTax" csv:"totalExcludTax"`
	InTaxSalesTotal                  json.Number  `json:"inTaxSalesTotal" csv:"inTaxSalesTotal"`
	TaxInclude                       json.Number  `json:"taxInclude" csv:"taxInclude"`
	OutTaxSalesTotal                 json.Number  `json:"outTaxSalesTotal" csv:"outTaxSalesTotal"`
	TaxExclude                       json.Number  `json:"taxExclude" csv:"taxExclude"`
	TaxTotal                         json.Number  `json:"taxTotal" csv:"taxTotal"`
	NonTaxSalesTotal                 json.Number  `json:"nonTaxSalesTotal" csv:"nonTaxSalesTotal"`
	TaxFreeTotal                     json.Number  `json:"taxFreeTotal" csv:"taxFreeTotal"`
	NonSalesTargetTaxFreeTotal       json.Number  `json:"nonSalesTargetTaxFreeTotal" csv:"nonSalesTargetTaxFreeTotal"`
	CostTotal                        json.Number  `json:"costTotal" csv:"costTotal"`
	GrossMargin                      json.Number  `json:"grossMargin" csv:"grossMargin"`
	Amount                           json.Number  `json:"amount" csv:"amount"`
	TransactionCount                 json.Number  `json:"transactionCount" csv:"transactionCount"`
	ReturnAmount                     json.Number  `json:"returnAmount" csv:"returnAmount"`
	Carriage                         json.Number  `json:"carriage" csv:"carriage"`
	Commission                       json.Number  `json:"commission" csv:"commission"`
	PreparationCash                  json.Number  `json:"preparationCash" csv:"preparationCash"`
	CashSales                        json.Number  `json:"cashSales" csv:"cashSales"`
	CreditSales                      json.Number  `json:"creditSales" csv:"creditSales"`
	OtherSalseList                   []OtherSales `json:"otherSalseList" csv:"-"`
	ChangeDifference                 json.Number  `json:"changeDifference" csv:"changeDifference"`
	PartPayment                      json.Number  `json:"partPayment" csv:"partPayment"`
	PartPaymentCash                  json.Number  `json:"partPaymentCash" csv:"partPaymentCash"`
	PartPaymentCredit                json.Number  `json:"partPaymentCredit" csv:"partPaymentCredit"`
	ReceivedDepositCash              json.Number  `json:"receivedDepositCash" csv:"receivedDepositCash"`
	ReceivedDepositCashTotal         json.Number  `json:"receivedDepositCashTotal" csv:"receivedDepositCashTotal"`
	ReceivedDepositCreditTotal       json.Number  `json:"receivedDepositCreditTotal" csv:"receivedDepositCreditTotal"`
	PartPaymentCancel                json.Number  `json:"partPaymentCancel" csv:"partPaymentCancel"`
	PartPaymentCashCancel            json.Number  `json:"partPaymentCashCancel" csv:"partPaymentCashCancel"`
	PartPaymentCreditCancel          json.Number  `json:"partPaymentCreditCancel" csv:"partPaymentCreditCancel"`
	Deposit                          json.Number  `json:"deposit" csv:"deposit"`
	ReturnDeposit                    json.Number  `json:"returnDeposit" csv:"returnDeposit"`
	Receipt                          json.Number  `json:"receipt" csv:"receipt"`
	Payment                          json.Number  `json:"payment" csv:"payment"`
	NonSalesCashTotal                json.Number  `json:"nonSalesCashTotal" csv:"nonSalesCashTotal"`
	NonSalesCreditTotal              json.Number  `json:"nonSalesCreditTotal" csv:"nonSalesCreditTotal"`
	NonSalesOtherTotal               json.Number  `json:"nonSalesOtherTotal" csv:"nonSalesOtherTotal"`
	NonSalesTaxFreeTotal             json.Number  `json:"nonSalesTaxFreeTotal" csv:"nonSalesTaxFreeTotal"`
	CalculateBalance                 json.Number  `json:"calculateBalance" csv:"calculateBalance"`
	RealBalance                      json.Number  `json:"realBalance" csv:"realBalance"`
	Difference                       json.Number  `json:"difference" csv:"difference"`
	Saving                           json.Number  `json:"saving" csv:"saving"`
	CarryOver                        json.Number  `json:"carryOver" csv:"carryOver"`
	TenThousandYen                   json.Number  `json:"tenThousandYen" csv:"tenThousandYen"`
	FiveThousandYen                  json.Number  `json:"fiveThousandYen" csv:"fiveThousandYen"`
	TwoThousandYen                   json.Number  `json:"twoThousandYen" csv:"twoThousandYen"`
	OneThousandYen                   json.Number  `json:"oneThousandYen" csv:"oneThousandYen"`
	FiveHundredYen                   json.Number  `json:"fiveHundredYen" csv:"fiveHundredYen"`
	OneHundredYen                    json.Number  `json:"oneHundredYen" csv:"oneHundredYen"`
	FiftyYen                         json.Number  `json:"fiftyYen" csv:"fiftyYen"`
	TenYen                           json.Number  `json:"tenYen" csv:"tenYen"`
	FiveYen                          json.Number  `json:"fiveYen" csv:"fiveYen"`
	OneYen                           json.Number  `json:"oneYen" csv:"oneYen"`
	Comment                          *string      `json:"comment" csv:"comment"`
	InsDateTime                      string       `json:"insDateTime" csv:"insDateTime"`
	UpdDateTime                      string       `json:"updDateTime" csv:"updDateTime"`
	SalesTotalNonSalesTargetDivision string       `json:"salesTotalNonSalesTargetDivision" csv:"salesTotalNonSalesTargetDivision"`
	TotalTaxFreeDivision             string       `json:"totalTaxFreeDivision" csv:"totalTaxFreeDivision"`
}

type DailySumCSV struct {
	*CSVHandler
	buf []DailySum
}

func NewDailySumCSV(bufSize int, output string) (*DailySumCSV, error) {
	buf := make([]DailySum, bufSize)
	handler, err := NewCSVHandler([]DailySum{}, output)
	if err != nil {
		return nil, err
	}

	return &DailySumCSV{
		handler,
		buf,
	}, nil
}

func (dsc *DailySumCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &dsc.buf[i])
	}
	dsc.CSVWriter.Write(dsc.buf[:len(resp.Result)])
	return dsc.CSVWriter
}

type OtherSales struct {
	ID                        json.Number `json:"id" csv:"id"`
	Name                      string      `json:"name" csv:"name"`
	PaymentMethodDivision     string      `json:"paymentMethodDivision" csv:"paymentMethodDivision"`
	PaymentMethodDivisionName string      `json:"paymentMethodDivisionName" csv:"paymentMethodDivisionName"`
	Sales                     json.Number `json:"sales" csv:"sales"`
}
