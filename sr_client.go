package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/antonholmquist/jason"
)

const dirFormat = "20060102150405"

type SrClient struct {
	config SrConfig
	client *http.Client
}

type SrRefResponse struct {
	TotalCount int
	Result     []*jason.Object
}

func NewSrClient(config SrConfig) SrClient {
	return SrClient{
		config: config,
		client: &http.Client{},
	}
}

func parseRefResponse(resp *http.Response) (*SrRefResponse, error) {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	v, err := jason.NewObjectFromBytes(b)
	if err != nil {
		return nil, err
	}

	result, err := v.GetObjectArray("result")
	if err != nil {
		return nil, err
	}

	totalCountStr, err := v.GetString("total_count")
	if err != nil {
		return nil, err
	}

	totalCount, err := strconv.Atoi(totalCountStr)
	if err != nil {
		return nil, err
	}
	refResponse := SrRefResponse{TotalCount: totalCount, Result: result}

	return &refResponse, nil
}

func (sc *SrClient) Request(p SrRefParams) (*SrRefResponse, error) {
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	data := url.Values{}
	data.Set("proc_name", p.ProcName)
	data.Add("p", string(jsonBytes))

	r, err := http.NewRequest("POST", sc.config.EndPoint, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	r.Header.Add("X_contract_id", sc.config.ContractID)
	r.Header.Add("X_access_token", sc.config.AccessToken)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := sc.client.Do(r)
	if err != nil {
		return nil, err
	}
	return parseRefResponse(resp)
}

func (sc *SrClient) DumpTableToCSV(p SrRefParams) (*CSVWriter, error) {
	output := fmt.Sprintf("%s/%s.csv", sc.config.OutputDir, p.TableName)
	handler, err := chooseCSVHandler(p, output)
	if err != nil {
		return nil, err
	}
	defer handler.GetCSVWriter().Close()

	resp, err := sc.Request(p)
	if err != nil {
		return nil, err
	}

	handler.Write(resp)
	return handler.GetCSVWriter(), nil
}

func chooseCSVHandler(p SrRefParams, output string) (SrCSVHandlerIf, error) {
	switch p.TableName {
	case CATEGORY:
		return NewCategoryCSV(p.Limit, output)
	case STORE:
		return NewStoreCSV(p.Limit, output)
	case PRODUCT:
		return NewProductCSV(p.Limit, output)
	case PRODUCT_PRICE:
		return NewProductPriceCSV(p.Limit, output)
	case PRODUCT_RESERVE_ITEM:
		return NewProductReseveItemCSV(p.Limit, output)
	case PRODUCT_RESERVE_ITEM_LABEL:
		return NewProductReseveItemLabelCSV(p.Limit, output)
	case PRODUCT_STORE:
		return NewProductStoreCSV(p.Limit, output)
	case PRODUCT_INVENTORY_RESERVATION:
		return NewProductInventoryReservationCSV(p.Limit, output)
	case CUSTOMER:
		return NewCustomerCSV(p.Limit, output)
	case STOCK:
		return NewStockCSV(p.Limit, output)
	case STOCK_HISTORY:
		return NewStockHistoryCSV(p.Limit, output)
	case TRANSACTION_HEAD:
		return NewTransactionHeadCSV(p.Limit, output)
	case TRANSACTION_DETAIL:
		return NewTransactionDetailCSV(p.Limit, output)
	case BARGAIN:
		return NewBargainCSV(p.Limit, output)
	case BARGAIN_PRODUCT:
		return NewBargainProductCSV(p.Limit, output)
	case BARGAIN_STORE:
		return NewBargainStoreCSV(p.Limit, output)
	case DAILY_SUM:
		return NewDailySumCSV(p.Limit, output)
	default:
		return nil, errors.New("No table name is matched")
	}
}
