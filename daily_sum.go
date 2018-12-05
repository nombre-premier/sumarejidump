package main

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type DailySum struct {
	SumDate                          string          `json:"sumDate" csv:"sumDate"`
	StoreID                          int             `json:"storeId" csv:"storeId"`
	CashDrawerID                     int             `json:"cashDrawerId" csv:"cashDrawerId"`
	Status                           string          `json:"status" csv:"status"`
	SalesTotal                       decimal.Decimal `json:"salesTotal" csv:"salesTotal"`
	Discount                         decimal.Decimal `json:"discount" csv:"discount"`
	PointDiscount                    decimal.Decimal `json:"pointDiscount" csv:"pointDiscount"`
	TaxExcludeReceive                decimal.Decimal `json:"taxExcludeReceive" csv:"taxExcludeReceive"`
	NonSalesTargetTotal              decimal.Decimal `json:"nonSalesTargetTotal" csv:"nonSalesTargetTotal"`
	Total                            decimal.Decimal `json:"total" csv:"total"`
	TotalExcludTax                   decimal.Decimal `json:"totalExcludTax" csv:"totalExcludTax"`
	InTaxSalesTotal                  decimal.Decimal `json:"inTaxSalesTotal" csv:"inTaxSalesTotal"`
	TaxInclude                       decimal.Decimal `json:"taxInclude" csv:"taxInclude"`
	OutTaxSalesTotal                 decimal.Decimal `json:"outTaxSalesTotal" csv:"outTaxSalesTotal"`
	TaxExclude                       decimal.Decimal `json:"taxExclude" csv:"taxExclude"`
	TaxTotal                         decimal.Decimal `json:"taxTotal" csv:"taxTotal"`
	NonTaxSalesTotal                 decimal.Decimal `json:"nonTaxSalesTotal" csv:"nonTaxSalesTotal"`
	TaxFreeTotal                     decimal.Decimal `json:"taxFreeTotal" csv:"taxFreeTotal"`
	NonSalesTargetTaxFreeTotal       decimal.Decimal `json:"nonSalesTargetTaxFreeTotal" csv:"nonSalesTargetTaxFreeTotal"`
	CostTotal                        decimal.Decimal `json:"costTotal" csv:"costTotal"`
	GrossMargin                      decimal.Decimal `json:"grossMargin" csv:"grossMargin"`
	Amount                           int             `json:"amount" csv:"amount"`
	TransactionCount                 int             `json:"transactionCount" csv:"transactionCount"`
	ReturnAmount                     int             `json:"returnAmount" csv:"returnAmount"`
	Carriage                         decimal.Decimal `json:"carriage" csv:"carriage"`
	Commission                       decimal.Decimal `json:"commission" csv:"commission"`
	PreparationCash                  decimal.Decimal `json:"preparationCash" csv:"preparationCash"`
	CashSales                        decimal.Decimal `json:"cashSales" csv:"cashSales"`
	CreditSales                      decimal.Decimal `json:"creditSales" csv:"creditSales"`
	OtherSalseList                   []OtherSales    `json:"otherSalseList" csv:"-"`
	ChangeDifference                 decimal.Decimal `json:"changeDifference" csv:"changeDifference"`
	PartPayment                      decimal.Decimal `json:"partPayment" csv:"partPayment"`
	PartPaymentCash                  decimal.Decimal `json:"partPaymentCash" csv:"partPaymentCash"`
	PartPaymentCredit                decimal.Decimal `json:"partPaymentCredit" csv:"partPaymentCredit"`
	ReceivedDepositCash              decimal.Decimal `json:"receivedDepositCash" csv:"receivedDepositCash"`
	ReceivedDepositCashTotal         decimal.Decimal `json:"receivedDepositCashTotal" csv:"receivedDepositCashTotal"`
	ReceivedDepositCreditTotal       decimal.Decimal `json:"receivedDepositCreditTotal" csv:"receivedDepositCreditTotal"`
	PartPaymentCancel                decimal.Decimal `json:"partPaymentCancel" csv:"partPaymentCancel"`
	PartPaymentCashCancel            decimal.Decimal `json:"partPaymentCashCancel" csv:"partPaymentCashCancel"`
	PartPaymentCreditCancel          decimal.Decimal `json:"partPaymentCreditCancel" csv:"partPaymentCreditCancel"`
	Deposit                          decimal.Decimal `json:"deposit" csv:"deposit"`
	ReturnDeposit                    decimal.Decimal `json:"returnDeposit" csv:"returnDeposit"`
	Receipt                          decimal.Decimal `json:"receipt" csv:"receipt"`
	Payment                          decimal.Decimal `json:"payment" csv:"payment"`
	NonSalesCashTotal                decimal.Decimal `json:"nonSalesCashTotal" csv:"nonSalesCashTotal"`
	NonSalesCreditTotal              decimal.Decimal `json:"nonSalesCreditTotal" csv:"nonSalesCreditTotal"`
	NonSalesOtherTotal               decimal.Decimal `json:"nonSalesOtherTotal" csv:"nonSalesOtherTotal"`
	NonSalesTaxFreeTotal             decimal.Decimal `json:"nonSalesTaxFreeTotal" csv:"nonSalesTaxFreeTotal"`
	CalculateBalance                 decimal.Decimal `json:"calculateBalance" csv:"calculateBalance"`
	RealBalance                      decimal.Decimal `json:"realBalance" csv:"realBalance"`
	Difference                       decimal.Decimal `json:"difference" csv:"difference"`
	Saving                           decimal.Decimal `json:"saving" csv:"saving"`
	CarryOver                        decimal.Decimal `json:"carryOver" csv:"carryOver"`
	TenThousandYen                   int             `json:"tenThousandYen" csv:"tenThousandYen"`
	FiveThousandYen                  int             `json:"fiveThousandYen" csv:"fiveThousandYen"`
	TwoThousandYen                   int             `json:"twoThousandYen" csv:"twoThousandYen"`
	OneThousandYen                   int             `json:"oneThousandYen" csv:"oneThousandYen"`
	FiveHundredYen                   int             `json:"fiveHundredYen" csv:"fiveHundredYen"`
	OneHundredYen                    int             `json:"oneHundredYen" csv:"oneHundredYen"`
	FiftyYen                         int             `json:"fiftyYen" csv:"fiftyYen"`
	TenYen                           int             `json:"tenYen" csv:"tenYen"`
	FiveYen                          int             `json:"fiveYen" csv:"fiveYen"`
	OneYen                           int             `json:"oneYen" csv:"oneYen"`
	Comment                          *string         `json:"comment" csv:"comment"`
	InsDateTime                      string          `json:"insDateTime" csv:"insDateTime"`
	UpdDateTime                      string          `json:"updDateTime" csv:"updDateTime"`
	SalesTotalNonSalesTargetDivision string          `json:"salesTotalNonSalesTargetDivision" csv:"salesTotalNonSalesTargetDivision"`
	TotalTaxFreeDivision             string          `json:"totalTaxFreeDivision" csv:"totalTaxFreeDivision"`
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
	ID                        int             `json:"id" csv:"id"`
	Name                      string          `json:"name" csv:"name"`
	PaymentMethodDivision     string          `json:"paymentMethodDivision" csv:"paymentMethodDivision"`
	PaymentMethodDivisionName string          `json:"paymentMethodDivisionName" csv:"paymentMethodDivisionName"`
	Sales                     decimal.Decimal `json:"sales" csv:"sales"`
}
