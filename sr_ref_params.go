package main

import "errors"

const (
	// category_ref
	CATEGORY_REF = "category_ref"
	CATEGORY     = "Category"

	// store_ref
	STORE_REF = "store_ref"
	STORE     = "Store"

	// product_ref
	PRODUCT_REF                   = "product_ref"
	PRODUCT                       = "Product"
	PRODUCT_PRICE                 = "ProductPrice"
	PRODUCT_RESERVE_ITEM          = "ProductReserveItem"
	PRODUCT_RESERVE_ITEM_LABEL    = "ProductReserveItemLabel"
	PRODUCT_STORE                 = "ProductStore"
	PRODUCT_INVENTORY_RESERVATION = "ProductInventoryReservation"

	// customer_ref
	CUSTOMER_REF = "customer_ref"
	CUSTOMER     = "Customer"

	// stock_ref
	STOCK_REF     = "stock_ref"
	STOCK         = "Stock"
	STOCK_HISTORY = "StockHistory"

	// transaction_ref
	TRANSACTION_REF    = "transaction_ref"
	TRANSACTION_HEAD   = "TransactionHead"
	TRANSACTION_DETAIL = "TransactionDetail"

	// daily_sum_ref
	DAILY_SUM_REF = "daily_sum_ref"
	DAILY_SUM     = "DailySum"

	// bargain_ref
	BARGAIN_REF     = "bargain_ref"
	BARGAIN         = "Bargain"
	BARGAIN_PRODUCT = "BargainProduct"
	BARGAIN_STORE   = "BargainStore"

	// shipping_ref
	SHIPPING_REF    = "shipping_ref"
	SHIPPING        = "Shipping"
	SHIPPING_DETAIL = "ShippingDetail"

	// receiving_ref
	RECEIVING_REF    = "receiving_ref"
	RECEIVING        = "Receiving"
	RECEIVING_DETAIL = "ReceivingDetail"

	// stocktaking_ref
	STOCKTAKING_REF    = "stocktaking_ref"
	STOCKTAKING_INFO   = "StocktakingInfo"
	STOCKTAKING_HEAD   = "StocktakingHead"
	STOCKTAKING_DETAIL = "StocktakingDetail"

	// loss_ref
	LOSS_REF    = "loss_ref"
	LOSS        = "Loss"
	LOSS_DETAIL = "LossDetail"

	//shipment_ref
	SHIPMENT_REF    = "shipment_ref"
	SHIPMENT        = "Shipment"
	SHIPMENT_DETAIL = "ShipmentDetail"
)

type SrRefParams struct {
	ProcName   string               `json:"-"`
	Fields     []string             `json:"fields"`
	Conditions []map[string]*string `json:"conditions"`
	Order      []string             `json:"order"`
	Limit      int                  `json:"limit"`
	Page       int                  `json:"page"`
	TableName  string               `json:"table_name"`
}

type SrTableMeta struct {
	Name     string
	Limit    int
	ProcName string
}

var SrTableMetas = []SrTableMeta{
	SrTableMeta{Name: CATEGORY, Limit: 1000, ProcName: CATEGORY_REF},
	SrTableMeta{Name: STORE, Limit: 1000, ProcName: STORE_REF},
	SrTableMeta{Name: PRODUCT, Limit: 1000, ProcName: PRODUCT_REF},
	SrTableMeta{Name: PRODUCT_PRICE, Limit: 1000, ProcName: PRODUCT_REF},
	SrTableMeta{Name: PRODUCT_RESERVE_ITEM, Limit: 1000, ProcName: PRODUCT_REF},
	SrTableMeta{Name: PRODUCT_RESERVE_ITEM_LABEL, Limit: 1000, ProcName: PRODUCT_REF},
	SrTableMeta{Name: PRODUCT_STORE, Limit: 1000, ProcName: PRODUCT_REF},
	SrTableMeta{Name: PRODUCT_INVENTORY_RESERVATION, Limit: 1000, ProcName: PRODUCT_REF},
	SrTableMeta{Name: CUSTOMER, Limit: 1000, ProcName: CUSTOMER_REF},
	SrTableMeta{Name: STOCK, Limit: 1000, ProcName: STOCK_REF},
	SrTableMeta{Name: STOCK_HISTORY, Limit: 1000, ProcName: STOCK_REF},
	SrTableMeta{Name: TRANSACTION_HEAD, Limit: 1000, ProcName: TRANSACTION_REF},
	SrTableMeta{Name: TRANSACTION_DETAIL, Limit: 1000, ProcName: TRANSACTION_REF},
	SrTableMeta{Name: DAILY_SUM, Limit: 100, ProcName: DAILY_SUM_REF},
	SrTableMeta{Name: BARGAIN, Limit: 1000, ProcName: BARGAIN_REF},
	SrTableMeta{Name: BARGAIN_PRODUCT, Limit: 1000, ProcName: BARGAIN_REF},
	SrTableMeta{Name: BARGAIN_STORE, Limit: 1000, ProcName: BARGAIN_REF},
	SrTableMeta{Name: SHIPPING, Limit: 1000, ProcName: SHIPPING_REF},
	SrTableMeta{Name: SHIPPING_DETAIL, Limit: 1000, ProcName: SHIPPING_REF},
	SrTableMeta{Name: RECEIVING, Limit: 1000, ProcName: RECEIVING_REF},
	SrTableMeta{Name: RECEIVING_DETAIL, Limit: 1000, ProcName: RECEIVING_REF},
	SrTableMeta{Name: STOCKTAKING_INFO, Limit: 1000, ProcName: STOCKTAKING_REF},
	SrTableMeta{Name: STOCKTAKING_HEAD, Limit: 1000, ProcName: STOCKTAKING_REF},
	SrTableMeta{Name: STOCKTAKING_DETAIL, Limit: 1000, ProcName: STOCKTAKING_REF},
	SrTableMeta{Name: LOSS, Limit: 1000, ProcName: LOSS_REF},
	SrTableMeta{Name: LOSS_DETAIL, Limit: 1000, ProcName: LOSS_REF},
	SrTableMeta{Name: SHIPMENT, Limit: 1000, ProcName: SHIPMENT_REF},
	SrTableMeta{Name: SHIPMENT_DETAIL, Limit: 1000, ProcName: SHIPMENT_REF},
}

func GetTableMetaByName(tableName string) *SrTableMeta {
	for _, t := range SrTableMetas {
		if t.Name == tableName {
			return &t
		}
	}
	return nil
}

func NewSrRefParamsWithConfig(tableName string, c SrConfig) (SrRefParams, error) {
	tableMeta := GetTableMetaByName(tableName)
	if tableMeta == nil {
		return SrRefParams{}, errors.New("non params is matched by table name")
	}

	return SrRefParams{ProcName: tableMeta.ProcName,
		Fields: []string{}, Conditions: []map[string]*string{c.Conditions},
		Order: []string{}, Limit: tableMeta.Limit, Page: 1,
		TableName: tableMeta.Name,
	}, nil
}
