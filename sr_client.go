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

func (sc *SrClient) Request(params SrRefParams) (*SrRefResponse, error) {
	jsonBytes, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	data := url.Values{}
	data.Set("proc_name", params.ProcName)
	data.Add("params", string(jsonBytes))

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

func (sc *SrClient) DumpTableToCSV(params SrRefParams) (*CSVWriter, error) {
	var handler SrCSVHandlerIf
	var err error
	output := fmt.Sprintf("%s/%s.csv", sc.config.OutputDir, params.TableName)
	switch params.TableName {
	case CATEGORY:
		handler, err = NewCategoryCSV(params.Limit, output)
	case STORE:
		handler, err = NewStoreCSV(params.Limit, output)
	case PRODUCT:
		handler, err = NewProductCSV(params.Limit, output)
	case PRODUCT_PRICE:
		handler, err = NewProductPriceCSV(params.Limit, output)
	case PRODUCT_RESERVE_ITEM:
		handler, err = NewProductReseveItemCSV(params.Limit, output)
	case PRODUCT_RESERVE_ITEM_LABEL:
		handler, err = NewProductReseveItemLabelCSV(params.Limit, output)
	case PRODUCT_STORE:
		handler, err = NewProductStoreCSV(params.Limit, output)
	case PRODUCT_INVENTORY_RESERVATION:
		handler, err = NewProductInventoryReservationCSV(params.Limit, output)
	case CUSTOMER:
		handler, err = NewCustomerCSV(params.Limit, output)
	case STOCK:
		handler, err = NewStockCSV(params.Limit, output)
	case STOCK_HISTORY:
		handler, err = NewStockHistoryCSV(params.Limit, output)
	case TRANSACTION_HEAD:
		handler, err = NewTransactionHeadCSV(params.Limit, output)
	case TRANSACTION_DETAIL:
		handler, err = NewTransactionDetailCSV(params.Limit, output)
	case BARGAIN:
		handler, err = NewBargainCSV(params.Limit, output)
	case BARGAIN_PRODUCT:
		handler, err = NewBargainProductCSV(params.Limit, output)
	case BARGAIN_STORE:
		handler, err = NewBargainStoreCSV(params.Limit, output)
	case DAILY_SUM:
		handler, err = NewDailySumCSV(params.Limit, output)
	default:
		return nil, errors.New("No table name is matched")
	}
	if err != nil {
		return nil, err
	}

	resp, err := sc.Request(params)
	if err != nil {
		return nil, err
	}

	handler.Write(resp)
	return handler.GetCSVWriter(), nil
}
