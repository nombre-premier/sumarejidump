package main

import (
	"encoding/json"
)

type TransactionHead struct {
	TransactionHeadID              json.Number  `json:"transactionHeadId" csv:"transactionHeadId"`
	TransactionDateTime            string       `json:"transactionDateTime" csv:"transactionDateTime"`
	TransactionHeadDivision        string       `json:"transactionHeadDivision" csv:"transactionHeadDivision"`
	CancelDivision                 string       `json:"cancelDivision" csv:"cancelDivision"`
	UnitNonDiscountsubtotal        *json.Number `json:"unitNonDiscountsubtotal" csv:"unitNonDiscountsubtotal"`
	UnitDiscountsubtotal           *json.Number `json:"unitDiscountsubtotal" csv:"unitDiscountsubtotal"`
	Subtotal                       *json.Number `json:"subtotal" csv:"subtotal"`
	SubtotalForDiscount            *json.Number `json:"subtotalForDiscount" csv:"subtotalForDiscount"`
	SubtotalDiscountPrice          *json.Number `json:"subtotalDiscountPrice" csv:"subtotalDiscountPrice"`
	SubtotalDiscountRate           *json.Number `json:"subtotalDiscountRate" csv:"subtotalDiscountRate"`
	SubtotalDiscountDivision       *json.Number `json:"subtotalDiscountDivision" csv:"subtotalDiscountDivision"`
	PointDiscount                  *json.Number `json:"pointDiscount" csv:"pointDiscount"`
	CouponDiscount                 *json.Number `json:"couponDiscount" csv:"couponDiscount"`
	Total                          *json.Number `json:"total" csv:"total"`
	TaxExclude                     json.Number  `json:"taxExclude" csv:"taxExclude"`
	TaxInclude                     json.Number  `json:"taxInclude" csv:"taxInclude"`
	RoundingDivision               string       `json:"roundingDivision" csv:"roundingDivision"`
	RoundingPrice                  *json.Number `json:"roundingPrice" csv:"roundingPrice"`
	CashTotal                      *json.Number `json:"cashTotal" csv:"cashTotal"`
	CreditTotal                    *json.Number `json:"creditTotal" csv:"creditTotal"`
	Deposit                        *json.Number `json:"deposit" csv:"deposit"`
	DepositCash                    *json.Number `json:"depositCash" csv:"depositCash"`
	DepositCredit                  *json.Number `json:"depositCredit" csv:"depositCredit"`
	Charge                         *json.Number `json:"charge" csv:"charge"`
	ChangeDifference               json.Number  `json:"changeDifference" csv:"changeDifference"`
	Amount                         *json.Number `json:"amount" csv:"amount"`
	ReturnAmount                   *json.Number `json:"returnAmount" csv:"returnAmount"`
	CostTotal                      *json.Number `json:"costTotal" csv:"costTotal"`
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
	SpendPoint                     *json.Number `json:"spendPoint" csv:"spendPoint"`
	Point                          *json.Number `json:"point" csv:"point"`
	TotalPoint                     *json.Number `json:"totalPoint" csv:"totalPoint"`
	CurrentMile                    *json.Number `json:"currentMile" csv:"currentMile"`
	EarnMile                       *json.Number `json:"earnMile" csv:"earnMile"`
	TotalMile                      *json.Number `json:"totalMile" csv:"totalMile"`
	AdjustmentMile                 *json.Number `json:"adjustmentMile" csv:"adjustmentMile"`
	AdjustmentMileDivision         *string      `json:"adjustmentMileDivision" csv:"adjustmentMileDivision"`
	AdjustmentMileValue            *json.Number `json:"adjustmentMileValue" csv:"adjustmentMileValue"`
	StoreID                        *json.Number `json:"storeId" csv:"storeId"`
	StoreCode                      *string      `json:"storeCode" csv:"storeCode"`
	TerminalID                     *json.Number `json:"terminalId" csv:"terminalId"`
	CustomerID                     *json.Number `json:"customerId" csv:"customerId"`
	CustomerCode                   *string      `json:"customerCode" csv:"customerCode"`
	TerminalTranID                 *string      `json:"terminalTranId" csv:"terminalTranId"`
	TerminalTranDateTime           string       `json:"terminalTranDateTime" csv:"terminalTranDateTime"`
	SumDivision                    string       `json:"sumDivision" csv:"sumDivision"`
	AdjustmentDateTime             *string      `json:"adjustmentDateTime" csv:"adjustmentDateTime"`
	SumDateTime                    *string      `json:"sumDateTime" csv:"sumDateTime"`
	CustomerRank                   *json.Number `json:"customerRank" csv:"customerRank"`
	CustomerGroupID                *json.Number `json:"customerGroupId" csv:"customerGroupId"`
	CustomerGroupID2               *json.Number `json:"customerGroupId2" csv:"customerGroupId2"`
	CustomerGroupID3               *json.Number `json:"customerGroupId3" csv:"customerGroupId3"`
	CustomerGroupID4               *json.Number `json:"customerGroupId4" csv:"customerGroupId4"`
	CustomerGroupID5               *json.Number `json:"customerGroupId5" csv:"customerGroupId5"`
	StaffID                        *json.Number `json:"staffId" csv:"staffId"`
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
	GuestNumbers                   *json.Number `json:"guestNumbers" csv:"guestNumbers"`
	TaxFreeSalesDivision           string       `json:"taxFreeSalesDivision" csv:"taxFreeSalesDivision"`
	NetTaxFreeGeneralTaxInclude    *json.Number `json:"netTaxFreeGeneralTaxInclude" csv:"netTaxFreeGeneralTaxInclude"`
	NetTaxFreeGeneralTaxExclude    *json.Number `json:"netTaxFreeGeneralTaxExclude" csv:"netTaxFreeGeneralTaxExclude"`
	NetTaxFreeConsumableTaxInclude *json.Number `json:"netTaxFreeConsumableTaxInclude" csv:"netTaxFreeConsumableTaxInclude"`
	NetTaxFreeConsumableTaxExclude *json.Number `json:"netTaxFreeConsumableTaxExclude" csv:"netTaxFreeConsumableTaxExclude"`
	Tags                           *string      `json:"tags" csv:"tags"`
	PointGivingDivision            *string      `json:"pointGivingDivision" csv:"pointGivingDivision"`
	PointGivingUnitPrice           *json.Number `json:"pointGivingUnitPrice" csv:"pointGivingUnitPrice"`
	PointGivingUnit                *json.Number `json:"pointGivingUnit" csv:"pointGivingUnit"`
	PointSpendDivision             *string      `json:"pointSpendDivision" csv:"pointSpendDivision"`
	MileageDivision                *string      `json:"mileageDivision" csv:"mileageDivision"`
	MileageLabel                   *string      `json:"mileageLabel" csv:"mileageLabel"`
	CustomerPinCode                *string      `json:"customerPinCode" csv:"customerPinCode"`
	DisposeDivision                string       `json:"disposeDivision" csv:"disposeDivision"`
	DisposeServerTransactionHeadID *json.Number `json:"disposeServerTransactionHeadId" csv:"disposeServerTransactionHeadId"`
	CancelDateTime                 *string      `json:"cancelDateTime" csv:"cancelDateTime"`
	SellDivision                   string       `json:"sellDivision" csv:"sellDivision"`
	TaxRate                        *json.Number `json:"taxRate" csv:"taxRate"`
	TaxRounding                    *string      `json:"taxRounding" csv:"taxRounding"`
	DiscountRoundingDivision       string       `json:"discountRoundingDivision" csv:"discountRoundingDivision"`
	TransactionUUID                *string      `json:"transactionUuid" csv:"transactionUuid"`
	GiftReceiptValidDays           *json.Number `json:"giftReceiptValidDays" csv:"giftReceiptValidDays"`
	ReceiptIssueNumberOfTimes      json.Number  `json:"receiptIssueNumberOfTimes" csv:"receiptIssueNumberOfTimes"`
	PickupTransactionHeadID        *json.Number `json:"pickupTransactionHeadId" csv:"pickupTransactionHeadId"`
	PickUpDate                     *string      `json:"pickUpDate" csv:"pickUpDate"`
	PartPayment                    *json.Number `json:"partPayment" csv:"partPayment"`
	PartPaymentClass               *string      `json:"partPaymentClass" csv:"partPaymentClass"`
	LayawayServerTransactionHeadID *json.Number `json:"layawayServerTransactionHeadId" csv:"layawayServerTransactionHeadId"`
	DisabledEdit                   *string      `json:"disabledEdit" csv:"disabledEdit"`
	UpdDateTime                    string       `json:"updDateTime" csv:"updDateTime"`
}

type TransactionHeadParquetSchema struct {
	TransactionHeadID              int64   `json:",string" parquet:"name=transaction_head_id, type=INT64"`
	TransactionDateTime            string  `parquet:"name=transaction_date_time, type=BYTE_ARRAY, convertedtype=UTF8"`
	TransactionHeadDivision        string  `parquet:"name=transaction_head_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	CancelDivision                 string  `parquet:"name=cancel_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	UnitNonDiscountsubtotal        *int64  `json:",string" parquet:"name=unit_non_discountsubtotal, type=INT64, repetitiontype=OPTIONAL"`
	UnitDiscountsubtotal           *int64  `json:",string" parquet:"name=unit_discountsubtotal, type=INT64, repetitiontype=OPTIONAL"`
	Subtotal                       *int64  `json:",string" parquet:"name=subtotal, type=INT64, repetitiontype=OPTIONAL"`
	SubtotalForDiscount            *int64  `json:",string" parquet:"name=subtotal_for_discount, type=INT64, repetitiontype=OPTIONAL"`
	SubtotalDiscountPrice          *int64  `json:",string" parquet:"name=subtotal_discount_price, type=INT64, repetitiontype=OPTIONAL"`
	SubtotalDiscountRate           *int64  `json:",string" parquet:"name=subtotal_discount_rate, type=INT64, repetitiontype=OPTIONAL"`
	SubtotalDiscountDivision       *int64  `json:",string" parquet:"name=subtotal_discount_division, type=INT64, repetitiontype=OPTIONAL"`
	PointDiscount                  *int64  `json:",string" parquet:"name=point_discount, type=INT64, repetitiontype=OPTIONAL"`
	CouponDiscount                 *int64  `json:",string" parquet:"name=coupon_discount, type=INT64, repetitiontype=OPTIONAL"`
	Total                          *int64  `json:",string" parquet:"name=total, type=INT64, repetitiontype=OPTIONAL"`
	TaxExclude                     int64   `json:",string" parquet:"name=tax_exclude, type=INT64"`
	TaxInclude                     int64   `json:",string" parquet:"name=tax_include, type=INT64"`
	RoundingDivision               string  `parquet:"name=rounding_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	RoundingPrice                  *int64  `json:",string" parquet:"name=rounding_price, type=INT64, repetitiontype=OPTIONAL"`
	CashTotal                      *int64  `json:",string" parquet:"name=cash_total, type=INT64, repetitiontype=OPTIONAL"`
	CreditTotal                    *int64  `json:",string" parquet:"name=credit_total, type=INT64, repetitiontype=OPTIONAL"`
	Deposit                        *int64  `json:",string" parquet:"name=deposit, type=INT64, repetitiontype=OPTIONAL"`
	DepositCash                    *int64  `json:",string" parquet:"name=deposit_cash, type=INT64, repetitiontype=OPTIONAL"`
	DepositCredit                  *int64  `json:",string" parquet:"name=deposit_credit, type=INT64, repetitiontype=OPTIONAL"`
	Charge                         *int64  `json:",string" parquet:"name=charge, type=INT64, repetitiontype=OPTIONAL"`
	ChangeDifference               int64   `json:",string" parquet:"name=change_difference, type=INT64"`
	Amount                         *int64  `json:",string" parquet:"name=amount, type=INT64, repetitiontype=OPTIONAL"`
	ReturnAmount                   *int64  `json:",string" parquet:"name=return_amount, type=INT64, repetitiontype=OPTIONAL"`
	CostTotal                      *int64  `json:",string" parquet:"name=cost_total, type=INT64, repetitiontype=OPTIONAL"`
	SalesHeadDivision              string  `parquet:"name=sales_head_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	InTaxSalesTotal                int64   `json:",string" parquet:"name=in_tax_sales_total, type=INT64"`
	OutTaxSalesTotal               int64   `json:",string" parquet:"name=out_tax_sales_total, type=INT64"`
	NonTaxSalesTotal               int64   `json:",string" parquet:"name=non_tax_sales_total, type=INT64"`
	NonSalesTargetTotal            int64   `json:",string" parquet:"name=non_sales_target_total, type=INT64"`
	NonSalesTargetOutTaxTotal      int64   `json:",string" parquet:"name=non_sales_target_out_tax_total, type=INT64"`
	NonSalesTargetInTaxTotal       int64   `json:",string" parquet:"name=non_sales_target_in_tax_total, type=INT64"`
	NonSalesTargetTaxFreeTotal     int64   `json:",string" parquet:"name=non_sales_target_tax_free_total, type=INT64"`
	NonSalesTargetCostTotal        int64   `json:",string" parquet:"name=non_sales_target_cost_total, type=INT64"`
	NonSalesTargetAmount           int64   `json:",string" parquet:"name=non_sales_target_amount, type=INT64"`
	NonSalesTargetReturnAmount     int64   `json:",string" parquet:"name=non_sales_target_return_amount, type=INT64"`
	NewPoint                       int64   `json:",string" parquet:"name=new_point, type=INT64"`
	SpendPoint                     *int64  `json:",string" parquet:"name=spend_point, type=INT64, repetitiontype=OPTIONAL"`
	Point                          *int64  `json:",string" parquet:"name=point, type=INT64, repetitiontype=OPTIONAL"`
	TotalPoint                     *int64  `json:",string" parquet:"name=total_point, type=INT64, repetitiontype=OPTIONAL"`
	CurrentMile                    *int64  `json:",string" parquet:"name=current_mile, type=INT64, repetitiontype=OPTIONAL"`
	EarnMile                       *int64  `json:",string" parquet:"name=earn_mile, type=INT64, repetitiontype=OPTIONAL"`
	TotalMile                      *int64  `json:",string" parquet:"name=total_mile, type=INT64, repetitiontype=OPTIONAL"`
	AdjustmentMile                 *int64  `json:",string" parquet:"name=adjustment_mile, type=INT64, repetitiontype=OPTIONAL"`
	AdjustmentMileDivision         *string `parquet:"name=adjustment_mile_division, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	AdjustmentMileValue            *int64  `json:",string" parquet:"name=adjustment_mile_value, type=INT64, repetitiontype=OPTIONAL"`
	StoreID                        *int64  `json:",string" parquet:"name=store_id, type=INT64, repetitiontype=OPTIONAL"`
	StoreCode                      *string `parquet:"name=store_code, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	TerminalID                     *int64  `json:",string" parquet:"name=terminal_id, type=INT64, repetitiontype=OPTIONAL"`
	CustomerID                     *int64  `json:",string" parquet:"name=customer_id, type=INT64, repetitiontype=OPTIONAL"`
	CustomerCode                   *string `parquet:"name=customer_code, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	TerminalTranID                 *string `parquet:"name=terminal_tran_id, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	TerminalTranDateTime           string  `parquet:"name=terminal_tran_date_time, type=BYTE_ARRAY, convertedtype=UTF8"`
	SumDivision                    string  `parquet:"name=sum_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	AdjustmentDateTime             *string `parquet:"name=adjustment_date_time, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	SumDateTime                    *string `parquet:"name=sum_date_time, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	CustomerRank                   *int64  `json:",string" parquet:"name=customer_rank, type=INT64, repetitiontype=OPTIONAL"`
	CustomerGroupID                *int64  `json:",string" parquet:"name=customer_group_id, type=INT64, repetitiontype=OPTIONAL"`
	CustomerGroupID2               *int64  `json:",string" parquet:"name=customer_group_id2, type=INT64, repetitiontype=OPTIONAL"`
	CustomerGroupID3               *int64  `json:",string" parquet:"name=customer_group_id3, type=INT64, repetitiontype=OPTIONAL"`
	CustomerGroupID4               *int64  `json:",string" parquet:"name=customer_group_id4, type=INT64, repetitiontype=OPTIONAL"`
	CustomerGroupID5               *int64  `json:",string" parquet:"name=customer_group_id5, type=INT64, repetitiontype=OPTIONAL"`
	StaffID                        *int64  `json:",string" parquet:"name=staff_id, type=INT64, repetitiontype=OPTIONAL"`
	StaffName                      *string `parquet:"name=staff_name, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	StaffCode                      *string `parquet:"name=staff_code, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	PaymentCount                   *string `parquet:"name=payment_count, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	SlipNumber                     *string `parquet:"name=slip_number, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	CancelSlipNumber               *string `parquet:"name=cancel_slip_number, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	AuthNumber                     *string `parquet:"name=auth_number, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	AuthDate                       *string `parquet:"name=auth_date, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	CardCompany                    *string `parquet:"name=card_company, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Memo                           *string `parquet:"name=memo, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	ReceiptMemo                    *string `parquet:"name=receipt_memo, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	PaymentMethodID1               *int64  `json:",string" parquet:"name=payment_method_id1, type=INT64, repetitiontype=OPTIONAL"`
	PaymentMethodName1             *string `parquet:"name=payment_method_name1, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	DepositOthers1                 *int64  `json:",string" parquet:"name=deposit_others1, type=INT64, repetitiontype=OPTIONAL"`
	PaymentMethodID2               *int64  `json:",string" parquet:"name=payment_method_id2, type=INT64, repetitiontype=OPTIONAL"`
	PaymentMethodName2             *string `parquet:"name=payment_method_name2, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	DepositOthers2                 *int64  `json:",string" parquet:"name=deposit_others2, type=INT64, repetitiontype=OPTIONAL"`
	PaymentMethodID3               *int64  `json:",string" parquet:"name=payment_method_id3, type=INT64, repetitiontype=OPTIONAL"`
	PaymentMethodName3             *string `parquet:"name=payment_method_name3, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	DepositOthers3                 *int64  `json:",string" parquet:"name=deposit_others3, type=INT64, repetitiontype=OPTIONAL"`
	Carriage                       *string `parquet:"name=carriage, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Commission                     *string `parquet:"name=commission, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	GuestNumbers                   *int64  `json:",string" parquet:"name=guest_numbers, type=INT64, repetitiontype=OPTIONAL"`
	TaxFreeSalesDivision           string  `parquet:"name=tax_free_sales_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	NetTaxFreeGeneralTaxInclude    *int64  `json:",string" parquet:"name=net_tax_free_general_tax_include, type=INT64, repetitiontype=OPTIONAL"`
	NetTaxFreeGeneralTaxExclude    *int64  `json:",string" parquet:"name=net_tax_free_general_tax_exclude, type=INT64, repetitiontype=OPTIONAL"`
	NetTaxFreeConsumableTaxInclude *int64  `json:",string" parquet:"name=net_tax_free_consumable_tax_include, type=INT64, repetitiontype=OPTIONAL"`
	NetTaxFreeConsumableTaxExclude *int64  `json:",string" parquet:"name=net_tax_free_consumable_tax_exclude, type=INT64, repetitiontype=OPTIONAL"`
	Tags                           *string `parquet:"name=tags, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	PointGivingDivision            *string `parquet:"name=point_giving_division, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	PointGivingUnitPrice           *int64  `json:",string" parquet:"name=point_giving_unit_price, type=INT64, repetitiontype=OPTIONAL"`
	PointGivingUnit                *int64  `json:",string" parquet:"name=point_giving_unit, type=INT64, repetitiontype=OPTIONAL"`
	PointSpendDivision             *string `parquet:"name=point_spend_division, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	MileageDivision                *string `parquet:"name=mileage_division, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	MileageLabel                   *string `parquet:"name=mileage_label, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	CustomerPinCode                *string `parquet:"name=customer_pin_code, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	DisposeDivision                string  `parquet:"name=dispose_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	DisposeServerTransactionHeadID *int64  `json:",string" parquet:"name=dispose_server_transaction_head_id, type=INT64, repetitiontype=OPTIONAL"`
	CancelDateTime                 *string `parquet:"name=cancel_date_time, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	SellDivision                   string  `parquet:"name=sell_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	TaxRate                        *int64  `json:",string" parquet:"name=tax_rate, type=INT64, repetitiontype=OPTIONAL"`
	TaxRounding                    *string `parquet:"name=tax_rounding, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	DiscountRoundingDivision       string  `parquet:"name=discount_rounding_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	TransactionUUID                *string `parquet:"name=transaction_uuid, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	GiftReceiptValidDays           *int64  `json:",string" parquet:"name=gift_receipt_valid_days, type=INT64, repetitiontype=OPTIONAL"`
	ReceiptIssueNumberOfTimes      int64   `json:",string" parquet:"name=receipt_issue_number_of_times, type=INT64"`
	PickupTransactionHeadID        *int64  `json:",string" parquet:"name=pickup_transaction_head_id, type=INT64, repetitiontype=OPTIONAL"`
	PickUpDate                     *string `parquet:"name=pick_up_date, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	PartPayment                    *int64  `json:",string" parquet:"name=part_payment, type=INT64, repetitiontype=OPTIONAL"`
	PartPaymentClass               *string `parquet:"name=part_payment_class, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	LayawayServerTransactionHeadID *int64  `json:",string" parquet:"name=layaway_server_transaction_head_id, type=INT64, repetitiontype=OPTIONAL"`
	DisabledEdit                   *string `parquet:"name=disabled_edit, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	UpdDateTime                    string  `parquet:"name=upd_date_time, type=BYTE_ARRAY, convertedtype=UTF8"`
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
	ProductID                    *json.Number `json:"productId" csv:"productId"`
	ProductCode                  *string      `json:"productCode" csv:"productCode"`
	ProductName                  *string      `json:"productName" csv:"productName"`
	TaxDivision                  string       `json:"taxDivision" csv:"taxDivision"`
	Price                        json.Number  `json:"price" csv:"price"`
	SalesPrice                   json.Number  `json:"salesPrice" csv:"salesPrice"`
	UnitDiscountPrice            json.Number  `json:"unitDiscountPrice" csv:"unitDiscountPrice"`
	UnitDiscountRate             *json.Number `json:"unitDiscountRate" csv:"unitDiscountRate"`
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
	CalcDiscount                 string       `json:"calcDiscount" csv:"calcDiscount"`
	TaxFreeDivision              string       `json:"taxFreeDivision" csv:"taxFreeDivision"`
	TaxFreeCommodityPrice        json.Number  `json:"taxFreeCommodityPrice" csv:"taxFreeCommodityPrice"`
	TaxFree                      json.Number  `json:"taxFree" csv:"taxFree"`
	ProductBundleGroupID         *json.Number `json:"productBundleGroupId" csv:"productBundleGroupId"`
	DiscountPriceProportional    json.Number  `json:"discountPriceProportional" csv:"discountPriceProportional"`
	DiscountPointProportional    json.Number  `json:"discountPointProportional" csv:"discountPointProportional"`
	DiscountCouponProportional   json.Number  `json:"discountCouponProportional" csv:"discountCouponProportional"`
	TaxIncludeProportional       json.Number  `json:"taxIncludeProportional" csv:"taxIncludeProportional"`
	TaxExcludeProportional       json.Number  `json:"taxExcludeProportional" csv:"taxExcludeProportional"`
	ProductBundleProportional    json.Number  `json:"productBundleProportional" csv:"productBundleProportional"`
	StaffDiscountProportional    json.Number  `json:"staffDiscountProportional" csv:"staffDiscountProportional"`
	BargainDiscountProportional  json.Number  `json:"bargainDiscountProportional" csv:"bargainDiscountProportional"`
	RoundingPriceProportional    json.Number  `json:"roundingPriceProportional" csv:"roundingPriceProportional"`
	InventoryReservationDivision string       `json:"inventoryReservationDivision" csv:"inventoryReservationDivision"`
	GroupCode                    *string      `json:"groupCode" csv:"groupCode"`
	UpdDateTime                  string       `json:"updDateTime" csv:"updDateTime"`
	ProductStaffDiscountRate     *json.Number `json:"productStaffDiscountRate" csv:"productStaffDiscountRate"`
	StaffRank                    *string      `json:"staffRank" csv:"staffRank"`
	StaffRankName                *string      `json:"staffRankName" csv:"staffRankName"`
	StaffDiscountRate            *json.Number `json:"staffDiscountRate" csv:"staffDiscountRate"`
	StaffDiscountDivision        *json.Number `json:"staffDiscountDivision" csv:"staffDiscountDivision"`
	ApplyStaffDiscountRate       *json.Number `json:"applyStaffDiscountRate" csv:"applyStaffDiscountRate"`
	ApplyStaffDiscountPrice      *json.Number `json:"applyStaffDiscountPrice" csv:"applyStaffDiscountPrice"`
	BargainID                    *json.Number `json:"bargainId" csv:"bargainId"`
	BargainName                  *string      `json:"bargainName" csv:"bargainName"`
	BargainDivision              *string      `json:"bargainDivision" csv:"bargainDivision"`
	BargainValue                 *json.Number `json:"bargainValue" csv:"bargainValue"`
	ApplyBargainValue            *json.Number `json:"applyBargainValue" csv:"applyBargainValue"`
	ApplyBargainDiscountPrice    *json.Number `json:"applyBargainDiscountPrice" csv:"applyBargainDiscountPrice"`
	TaxRate                      json.Number  `json:"taxRate" csv:"taxRate"`
	StandardTaxRate              json.Number  `json:"standardTaxRate" csv:"standardTaxRate"`
	ModifiedTaxRate              *json.Number `json:"modifiedTaxRate" csv:"modifiedTaxRate"`
	ReduceTaxID                  *json.Number `json:"reduceTaxId" csv:"reduceTaxId"`
	ReduceTaxName                *string      `json:"reduceTaxName" csv:"reduceTaxName"`
	ReduceTaxRate                *json.Number `json:"reduceTaxRate" csv:"reduceTaxRate"`
	Color                        *string      `json:"color" csv:"color"`
	Size                         *string      `json:"size" csv:"size"`
}

type TransactionDetailParquetSchema struct {
	TransactionHeadID            int64    `json:",string" parquet:"name=transaction_head_id, type=INT64"`
	TransactionDateTime          *string  `parquet:"name=transaction_date_time, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	TransactionDetailID          int64    `json:",string" parquet:"name=transaction_detail_id, type=INT64"`
	ParentTransactionDetailID    *int64   `json:",string" parquet:"name=parent_transaction_detail_id, type=INT64, repetitiontype=OPTIONAL"`
	TransactionDetailDivision    string   `parquet:"name=transaction_detail_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	ProductID                    *int64   `json:",string" parquet:"name=product_id, type=INT64, repetitiontype=OPTIONAL"`
	ProductCode                  *string  `parquet:"name=product_code, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	ProductName                  *string  `parquet:"name=product_name, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	TaxDivision                  string   `parquet:"name=tax_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	Price                        int64    `json:",string" parquet:"name=price, type=INT64"`
	SalesPrice                   int64    `json:",string" parquet:"name=sales_price, type=INT64"`
	UnitDiscountPrice            int64    `json:",string" parquet:"name=unit_discount_price, type=INT64"`
	UnitDiscountRate             *float64 `parquet:"name=unit_discount_rate, type=DOUBLE, repetitiontype=OPTIONAL"`
	UnitDiscountDivision         *int64   `json:",string" parquet:"name=unit_discount_division, type=INT64, repetitiontype=OPTIONAL"`
	Cost                         *int64   `json:",string" parquet:"name=cost, type=INT64, repetitiontype=OPTIONAL"`
	Quantity                     int64    `json:",string" parquet:"name=quantity, type=INT64"`
	UnitNonDiscountSum           int64    `json:",string" parquet:"name=unit_non_discount_sum, type=INT64"`
	UnitDiscountSum              int64    `json:",string" parquet:"name=unit_discount_sum, type=INT64"`
	UnitDiscountedSum            int64    `json:",string" parquet:"name=unit_discounted_sum, type=INT64"`
	CostSum                      int64    `json:",string" parquet:"name=cost_sum, type=INT64"`
	CategoryID                   int64    `json:",string" parquet:"name=category_id, type=INT64"`
	CategoryName                 string   `parquet:"name=category_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	DiscriminationNo             *string  `parquet:"name=discrimination_no, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	SalesDivision                string   `parquet:"name=sales_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	ProductDivision              string   `parquet:"name=product_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	PointNotApplicable           string   `parquet:"name=point_not_applicable, type=BYTE_ARRAY, convertedtype=UTF8"`
	CalcDiscount                 string   `parquet:"name=calc_discount, type=BYTE_ARRAY, convertedtype=UTF8"`
	TaxFreeDivision              string   `parquet:"name=tax_free_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	TaxFreeCommodityPrice        int64    `json:",string" parquet:"name=tax_free_commodity_price, type=INT64"`
	TaxFree                      int64    `json:",string" parquet:"name=tax_free, type=INT64"`
	ProductBundleGroupID         *int64   `json:",string" parquet:"name=product_bundle_group_id, type=INT64, repetitiontype=OPTIONAL"`
	DiscountPriceProportional    int64    `json:",string" parquet:"name=discount_price_proportional, type=INT64"`
	DiscountPointProportional    int64    `json:",string" parquet:"name=discount_point_proportional, type=INT64"`
	DiscountCouponProportional   int64    `json:",string" parquet:"name=discount_coupon_proportional, type=INT64"`
	TaxIncludeProportional       int64    `json:",string" parquet:"name=tax_include_proportional, type=INT64"`
	TaxExcludeProportional       int64    `json:",string" parquet:"name=tax_exclude_proportional, type=INT64"`
	ProductBundleProportional    int64    `json:",string" parquet:"name=product_bundle_proportional, type=INT64"`
	StaffDiscountProportional    int64    `json:",string" parquet:"name=staff_discount_proportional, type=INT64"`
	BargainDiscountProportional  int64    `json:",string" parquet:"name=bargain_discount_proportional, type=INT64"`
	RoundingPriceProportional    int64    `json:",string" parquet:"name=rounding_price_proportional, type=INT64"`
	InventoryReservationDivision string   `parquet:"name=inventory_reservation_division, type=BYTE_ARRAY, convertedtype=UTF8"`
	GroupCode                    *string  `parquet:"name=group_code, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	UpdDateTime                  *string  `parquet:"name=upd_date_time, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	ProductStaffDiscountRate     *float64 `parquet:"name=product_staff_discount_rate, type=DOUBLE, repetitiontype=OPTIONAL"`
	StaffRank                    *string  `parquet:"name=staff_rank, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	StaffRankName                *string  `parquet:"name=staff_rank_name, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	StaffDiscountRate            *float64 `parquet:"name=staff_discount_rate, type=DOUBLE, repetitiontype=OPTIONAL"`
	StaffDiscountDivision        *int64   `json:",string" parquet:"name=staff_discount_division, type=INT64, repetitiontype=OPTIONAL"`
	ApplyStaffDiscountRate       *float64 `parquet:"name=apply_staff_discount_rate, type=DOUBLE, repetitiontype=OPTIONAL"`
	ApplyStaffDiscountPrice      *float64 `parquet:"name=apply_staff_discount_price, type=DOUBLE, repetitiontype=OPTIONAL"`
	BargainID                    *int64   `json:",string" parquet:"name=bargain_id, type=INT64, repetitiontype=OPTIONAL"`
	BargainName                  *string  `parquet:"name=bargain_name, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	BargainDivision              *string  `parquet:"name=bargain_division, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	BargainValue                 *int64   `json:",string" parquet:"name=bargain_value, type=INT64, repetitiontype=OPTIONAL"`
	ApplyBargainValue            *int64   `json:",string" parquet:"name=apply_bargain_value, type=INT64, repetitiontype=OPTIONAL"`
	ApplyBargainDiscountPrice    *int64   `json:",string" parquet:"name=apply_bargain_discount_price, type=INT64, repetitiontype=OPTIONAL"`
	TaxRate                      int64    `json:",string" parquet:"name=tax_rate, type=INT64"`
	StandardTaxRate              int64    `json:",string" parquet:"name=standard_tax_rate, type=INT64"`
	ModifiedTaxRate              *float64 `parquet:"name=modified_tax_rate, type=DOUBLE, repetitiontype=OPTIONAL"`
	ReduceTaxID                  *int64   `json:",string" parquet:"name=reduce_tax_id, type=INT64, repetitiontype=OPTIONAL"`
	ReduceTaxName                *string  `parquet:"name=reduce_tax_name, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	ReduceTaxRate                *float64 `parquet:"name=reduce_tax_rate, type=DOUBLE, repetitiontype=OPTIONAL"`
	Color                        *string  `parquet:"name=color, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Size                         *string  `parquet:"name=size, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
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
