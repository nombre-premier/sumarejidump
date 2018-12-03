package main

type Stock struct {
	StoreID     int    `json:"storeId"`
	ProductID   int    `json:"productId"`
	StockAmount string `json:"stockAmount"`
	UpdDateTime string `json:"updDateTime"`
}

type StockHistory struct {
	ID             int         `json:"id"`
	UpdDateTime    string      `json:"updDateTime"`
	TargetDateTime string      `json:"targetDateTime"`
	ProductID      int         `json:"productId"`
	StoreID        int         `json:"storeId"`
	Amount         int         `json:"amount"`
	StockAmount    int         `json:"stockAmount"`
	StockDivision  string      `json:"stockDivision"`
	FromStoreID    interface{} `json:"fromStoreId"`
	ToStoreID      interface{} `json:"toStoreId"`
}
