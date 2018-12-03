package main

import (
	"github.com/shopspring/decimal"
)

type Bargain struct {
	BargainID   int    `json:"bargainId"`
	BargainName string `json:"bargainName"`
	TermStart   string `json:"termStart"`
	TermEnd     string `json:"termEnd"`
}

type BargainProduct struct {
	BargainProductID int             `json:"bargainProductId"`
	BargainID        int             `json:"bargainId"`
	TargetDivision   string          `json:"targetDivision"`
	TargetID         string          `json:"targetId"`
	Division         string          `json:"division"`
	Value            decimal.Decimal `json:"value"`
}

type BargainStore struct {
	BargainStoreID int `json:"bargainStoreId"`
	BargainID      int `json:"bargainId"`
	StoreID        int `json:"storeId"`
}
