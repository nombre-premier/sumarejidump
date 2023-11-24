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

type DailySumParquetSchema struct {
	SumDate                          string   `parquet:"name=sum_date, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	StoreID                          int64    `parquet:"name=store_id, type=INT64"`
	CashDrawerID                     int64    `parquet:"name=cash_drawer_id, type=INT64"`
	Status                           string   `parquet:"name=status, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	SalesTotal                       float64  `parquet:"name=sales_total, type=DOUBLE"`
	Discount                         float64  `parquet:"name=discount, type=DOUBLE"`
	PointDiscount                    float64  `parquet:"name=point_discount, type=DOUBLE"`
	TaxExcludeReceive                float64  `parquet:"name=tax_exclude_receive, type=DOUBLE"`
	NonSalesTargetTotal              float64  `parquet:"name=non_sales_target_total, type=DOUBLE"`
	Total                            float64  `parquet:"name=total, type=DOUBLE"`
	TotalExcludTax                   float64  `parquet:"name=total_exclud_tax, type=DOUBLE"`
	InTaxSalesTotal                  float64  `parquet:"name=in_tax_sales_total, type=DOUBLE"`
	TaxInclude                       float64  `parquet:"name=tax_include, type=DOUBLE"`
	OutTaxSalesTotal                 float64  `parquet:"name=out_tax_sales_total, type=DOUBLE"`
	TaxExclude                       float64  `parquet:"name=tax_exclude, type=DOUBLE"`
	TaxTotal                         float64  `parquet:"name=tax_total, type=DOUBLE"`
	NonTaxSalesTotal                 float64  `parquet:"name=non_tax_sales_total, type=DOUBLE"`
	TaxFreeTotal                     float64  `parquet:"name=tax_free_total, type=DOUBLE"`
	NonSalesTargetTaxFreeTotal       float64  `parquet:"name=non_sales_target_tax_free_total, type=DOUBLE"`
	CostTotal                        float64  `parquet:"name=cost_total, type=DOUBLE"`
	GrossMargin                      float64  `parquet:"name=gross_margin, type=DOUBLE"`
	Amount                           float64  `parquet:"name=amount, type=DOUBLE"`
	TransactionCount                 float64  `parquet:"name=transaction_count, type=DOUBLE"`
	ReturnAmount                     float64  `parquet:"name=return_amount, type=DOUBLE"`
	Carriage                         float64  `parquet:"name=carriage, type=DOUBLE"`
	Commission                       float64  `parquet:"name=commission, type=DOUBLE"`
	PreparationCash                  float64  `parquet:"name=preparation_cash, type=DOUBLE"`
	CashSales                        float64  `parquet:"name=cash_sales, type=DOUBLE"`
	CreditSales                      float64  `parquet:"name=credit_sales, type=DOUBLE"`
	OtherSalseList                   []OtherSalesParquetSchema `parquet:"name=other_sales_list, type=LIST, convertedtype=LIST"`
	ChangeDifference                 float64  `parquet:"name=change_difference, type=DOUBLE"`
	PartPayment                      float64  `parquet:"name=part_payment, type=DOUBLE"`
	PartPaymentCash                  float64  `parquet:"name=part_payment_cash, type=DOUBLE"`
	PartPaymentCredit                float64  `parquet:"name=part_payment_credit, type=DOUBLE"`
	ReceivedDepositCash              float64  `parquet:"name=received_deposit_cash, type=DOUBLE"`
	ReceivedDepositCashTotal         float64  `parquet:"name=received_deposit_cash_total, type=DOUBLE"`
	ReceivedDepositCreditTotal       float64  `parquet:"name=received_deposit_credit_total, type=DOUBLE"`
	PartPaymentCancel                float64  `parquet:"name=part_payment_cancel, type=DOUBLE"`
	PartPaymentCashCancel            float64  `parquet:"name=part_payment_cash_cancel, type=DOUBLE"`
	PartPaymentCreditCancel          float64  `parquet:"name=part_payment_credit_cancel, type=DOUBLE"`
	Deposit                          float64  `parquet:"name=deposit, type=DOUBLE"`
	ReturnDeposit                    float64  `parquet:"name=return_deposit, type=DOUBLE"`
	Receipt                          float64  `parquet:"name=receipt, type=DOUBLE"`
	Payment                          float64  `parquet:"name=payment, type=DOUBLE"`
	NonSalesCashTotal                float64  `parquet:"name=non_sales_cash_total, type=DOUBLE"`
	NonSalesCreditTotal              float64  `parquet:"name=non_sales_credit_total, type=DOUBLE"`
	NonSalesOtherTotal               float64  `parquet:"name=non_sales_other_total, type=DOUBLE"`
	NonSalesTaxFreeTotal             float64  `parquet:"name=non_sales_tax_free_total, type=DOUBLE"`
	CalculateBalance                 float64  `parquet:"name=calculate_balance, type=DOUBLE"`
	RealBalance                      float64  `parquet:"name=real_balance, type=DOUBLE"`
	Difference                       float64  `parquet:"name=difference, type=DOUBLE"`
	Saving                           float64  `parquet:"name=saving, type=DOUBLE"`
	CarryOver                        float64  `parquet:"name=carry_over, type=DOUBLE"`
	TenThousandYen                   float64  `parquet:"name=ten_thousand_yen, type=DOUBLE"`
	FiveThousandYen                  float64  `parquet:"name=five_thousand_yen, type=DOUBLE"`
	TwoThousandYen                   float64  `parquet:"name=two_thousand_yen, type=DOUBLE"`
	OneThousandYen                   float64  `parquet:"name=one_thousand_yen, type=DOUBLE"`
	FiveHundredYen                   float64  `parquet:"name=five_hundred_yen, type=DOUBLE"`
	OneHundredYen                    float64  `parquet:"name=one_hundred_yen, type=DOUBLE"`
	FiftyYen                         float64  `parquet:"name=fifty_yen, type=DOUBLE"`
	TenYen                           float64  `parquet:"name=ten_yen, type=DOUBLE"`
	FiveYen                          float64  `parquet:"name=five_yen, type=DOUBLE"`
	OneYen                           float64  `parquet:"name=one_yen, type=DOUBLE"`
	Comment                          *string  `parquet:"name=comment, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	InsDateTime                      string   `parquet:"name=ins_date_time, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	UpdDateTime                      string   `parquet:"name=upd_date_time, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	SalesTotalNonSalesTargetDivision string   `parquet:"name=sales_total_non_sales_target_division, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	TotalTaxFreeDivision             string   `parquet:"name=total_tax_free_division, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
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
type OtherSalesParquetSchema struct {
	ID                        int64   `parquet:"name=id, type=INT64"`
	Name                      string  `parquet:"name=name, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	PaymentMethodDivision     string  `parquet:"name=payment_method_division, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	PaymentMethodDivisionName string  `parquet:"name=payment_method_division_name, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Sales                     float64 `parquet:"name=sales, type=DOUBLE"`
}