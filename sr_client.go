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

func writeCSV(cw *CSVWriter, resultBuffer interface{}, resp *SrRefResponse) (*CSVWriter, error) {
	switch buffer := resultBuffer.(type) {
	case []Category:
		for i, r := range resp.Result {
			json.Unmarshal([]byte(r.String()), &buffer[i])
		}
		cw.Write(buffer[:len(resp.Result)])
	case []Store:
		for i, r := range resp.Result {
			json.Unmarshal([]byte(r.String()), &buffer[i])
		}
		cw.Write(buffer[:len(resp.Result)])
	case []Product:
		for i, r := range resp.Result {
			json.Unmarshal([]byte(r.String()), &buffer[i])
		}
		cw.Write(buffer[:len(resp.Result)])
	case []ProductPrice:
		for i, r := range resp.Result {
			json.Unmarshal([]byte(r.String()), &buffer[i])
		}
		cw.Write(buffer[:len(resp.Result)])
	case []ProductStore:
		for i, r := range resp.Result {
			json.Unmarshal([]byte(r.String()), &buffer[i])
		}
		cw.Write(buffer[:len(resp.Result)])
	case []ProductReseveItem:
		for i, r := range resp.Result {
			json.Unmarshal([]byte(r.String()), &buffer[i])
		}
		cw.Write(buffer[:len(resp.Result)])
	case []ProductReseveItemLabel:
		for i, r := range resp.Result {
			json.Unmarshal([]byte(r.String()), &buffer[i])
		}
		cw.Write(buffer[:len(resp.Result)])
	case []ProductInventoryReservation:
		for i, r := range resp.Result {
			json.Unmarshal([]byte(r.String()), &buffer[i])
		}
		cw.Write(buffer[:len(resp.Result)])
	case []Customer:
		for i, r := range resp.Result {
			json.Unmarshal([]byte(r.String()), &buffer[i])
		}
		cw.Write(buffer[:len(resp.Result)])
	case []Stock:
		for i, r := range resp.Result {
			json.Unmarshal([]byte(r.String()), &buffer[i])
		}
		cw.Write(buffer[:len(resp.Result)])
	case []StockHistory:
		for i, r := range resp.Result {
			json.Unmarshal([]byte(r.String()), &buffer[i])
		}
		cw.Write(buffer[:len(resp.Result)])
	case []TransactionHead:
		for i, r := range resp.Result {
			json.Unmarshal([]byte(r.String()), &buffer[i])
		}
		cw.Write(buffer[:len(resp.Result)])
	case []TransactionDetail:
		for i, r := range resp.Result {
			json.Unmarshal([]byte(r.String()), &buffer[i])
		}
		cw.Write(buffer[:len(resp.Result)])
	case []DailySum:
		for i, r := range resp.Result {
			json.Unmarshal([]byte(r.String()), &buffer[i])
		}
		cw.Write(buffer[:len(resp.Result)])
	case []Bargain:
		for i, r := range resp.Result {
			json.Unmarshal([]byte(r.String()), &buffer[i])
		}
		cw.Write(buffer[:len(resp.Result)])
	case []BargainProduct:
		for i, r := range resp.Result {
			json.Unmarshal([]byte(r.String()), &buffer[i])
		}
		cw.Write(buffer[:len(resp.Result)])
	case []BargainStore:
		for i, r := range resp.Result {
			json.Unmarshal([]byte(r.String()), &buffer[i])
		}
		cw.Write(buffer[:len(resp.Result)])
	default:
		return nil, errors.New("No buffer is matched")
	}

	return cw, nil
}

func (sc *SrClient) DumpTableToCSV(params SrRefParams) (*CSVWriter, error) {
	switch params.TableName {
	case CATEGORY:
		resultBuffer := make([]Category, params.Limit)

		cw, err := NewCSVWriter(resultBuffer[:0], fmt.Sprintf("%s/%s.csv", sc.config.OutputDir, params.TableName))
		if err != nil {
			return nil, err
		}

		resp, err := sc.Request(params)
		if err != nil {
			return nil, err
		}

		return writeCSV(cw, resultBuffer, resp)
	case STORE:
		resultBuffer := make([]Store, params.Limit)

		cw, err := NewCSVWriter(resultBuffer[:0], fmt.Sprintf("%s/%s.csv", sc.config.OutputDir, params.TableName))
		if err != nil {
			return nil, err
		}

		resp, err := sc.Request(params)
		if err != nil {
			return nil, err
		}

		return writeCSV(cw, resultBuffer, resp)
	case PRODUCT:
		resultBuffer := make([]Product, params.Limit)

		cw, err := NewCSVWriter(resultBuffer[:0], fmt.Sprintf("%s/%s.csv", sc.config.OutputDir, params.TableName))
		if err != nil {
			return nil, err
		}

		resp, err := sc.Request(params)
		if err != nil {
			return nil, err
		}

		return writeCSV(cw, resultBuffer, resp)
	case PRODUCT_PRICE:
		resultBuffer := make([]ProductPrice, params.Limit)

		cw, err := NewCSVWriter(resultBuffer[:0], fmt.Sprintf("%s/%s.csv", sc.config.OutputDir, params.TableName))
		if err != nil {
			return nil, err
		}

		resp, err := sc.Request(params)
		if err != nil {
			return nil, err
		}
		return writeCSV(cw, resultBuffer, resp)
	case PRODUCT_RESERVE_ITEM:
		resultBuffer := make([]ProductReseveItem, params.Limit)

		cw, err := NewCSVWriter(resultBuffer[:0], fmt.Sprintf("%s/%s.csv", sc.config.OutputDir, params.TableName))
		if err != nil {
			return nil, err
		}

		resp, err := sc.Request(params)
		if err != nil {
			return nil, err
		}
		return writeCSV(cw, resultBuffer, resp)
	case PRODUCT_RESERVE_ITEM_LABEL:
		resultBuffer := make([]ProductReseveItemLabel, params.Limit)

		cw, err := NewCSVWriter(resultBuffer[:0], fmt.Sprintf("%s/%s.csv", sc.config.OutputDir, params.TableName))
		if err != nil {
			return nil, err
		}

		resp, err := sc.Request(params)
		if err != nil {
			return nil, err
		}

		return writeCSV(cw, resultBuffer, resp)
	case PRODUCT_STORE:
		resultBuffer := make([]ProductStore, params.Limit)

		cw, err := NewCSVWriter(resultBuffer[:0], fmt.Sprintf("%s/%s.csv", sc.config.OutputDir, params.TableName))
		if err != nil {
			return nil, err
		}

		resp, err := sc.Request(params)
		if err != nil {
			return nil, err
		}
		return writeCSV(cw, resultBuffer, resp)
	case PRODUCT_INVENTORY_RESERVATION:
		resultBuffer := make([]ProductInventoryReservation, params.Limit)

		cw, err := NewCSVWriter(resultBuffer[:0], fmt.Sprintf("%s/%s.csv", sc.config.OutputDir, params.TableName))
		if err != nil {
			return nil, err
		}

		resp, err := sc.Request(params)
		if err != nil {
			return nil, err
		}
		return writeCSV(cw, resultBuffer, resp)
	case CUSTOMER:
		resultBuffer := make([]Customer, params.Limit)

		cw, err := NewCSVWriter(resultBuffer[:0], fmt.Sprintf("%s/%s.csv", sc.config.OutputDir, params.TableName))
		if err != nil {
			return nil, err
		}

		resp, err := sc.Request(params)
		if err != nil {
			return nil, err
		}
		return writeCSV(cw, resultBuffer, resp)
	case STOCK:
		resultBuffer := make([]Stock, params.Limit)

		cw, err := NewCSVWriter(resultBuffer[:0], fmt.Sprintf("%s/%s.csv", sc.config.OutputDir, params.TableName))
		if err != nil {
			return nil, err
		}

		resp, err := sc.Request(params)
		if err != nil {
			return nil, err
		}

		return writeCSV(cw, resultBuffer, resp)
	case STOCK_HISTORY:
		resultBuffer := make([]StockHistory, params.Limit)

		cw, err := NewCSVWriter(resultBuffer[:0], fmt.Sprintf("%s/%s.csv", sc.config.OutputDir, params.TableName))
		if err != nil {
			return nil, err
		}

		resp, err := sc.Request(params)
		if err != nil {
			return nil, err
		}

		return writeCSV(cw, resultBuffer, resp)
	case TRANSACTION_HEAD:
		resultBuffer := make([]TransactionHead, params.Limit)

		cw, err := NewCSVWriter(resultBuffer[:0], fmt.Sprintf("%s/%s.csv", sc.config.OutputDir, params.TableName))
		if err != nil {
			return nil, err
		}

		resp, err := sc.Request(params)
		if err != nil {
			return nil, err
		}

		return writeCSV(cw, resultBuffer, resp)
	case TRANSACTION_DETAIL:
		resultBuffer := make([]TransactionDetail, params.Limit)

		cw, err := NewCSVWriter(resultBuffer[:0], fmt.Sprintf("%s/%s.csv", sc.config.OutputDir, params.TableName))
		if err != nil {
			return nil, err
		}

		resp, err := sc.Request(params)
		if err != nil {
			return nil, err
		}

		return writeCSV(cw, resultBuffer, resp)
	case BARGAIN:
		resultBuffer := make([]Bargain, params.Limit)

		cw, err := NewCSVWriter(resultBuffer[:0], fmt.Sprintf("%s/%s.csv", sc.config.OutputDir, params.TableName))
		if err != nil {
			return nil, err
		}

		resp, err := sc.Request(params)
		if err != nil {
			return nil, err
		}
		return writeCSV(cw, resultBuffer, resp)

	case BARGAIN_PRODUCT:
		resultBuffer := make([]BargainProduct, params.Limit)

		cw, err := NewCSVWriter(resultBuffer[:0], fmt.Sprintf("%s/%s.csv", sc.config.OutputDir, params.TableName))
		if err != nil {
			return nil, err
		}

		resp, err := sc.Request(params)
		if err != nil {
			return nil, err
		}

		return writeCSV(cw, resultBuffer, resp)
	case BARGAIN_STORE:
		resultBuffer := make([]BargainStore, params.Limit)

		cw, err := NewCSVWriter(resultBuffer[:0], fmt.Sprintf("%s/%s.csv", sc.config.OutputDir, params.TableName))
		if err != nil {
			return nil, err
		}

		resp, err := sc.Request(params)
		if err != nil {
			return nil, err
		}

		for i, r := range resp.Result {
			json.Unmarshal([]byte(r.String()), &resultBuffer[i])
		}
		cw.Write(resultBuffer[:len(resp.Result)])
		return cw, nil
	case DAILY_SUM:
		resultBuffer := make([]DailySum, params.Limit)

		cw, err := NewCSVWriter(resultBuffer[:0], fmt.Sprintf("%s/%s.csv", sc.config.OutputDir, params.TableName))
		if err != nil {
			return nil, err
		}

		resp, err := sc.Request(params)
		if err != nil {
			return nil, err
		}

		for i, r := range resp.Result {
			json.Unmarshal([]byte(r.String()), &resultBuffer[i])
		}
		cw.Write(resultBuffer[:len(resp.Result)])
		return cw, nil
	default:
		return nil, errors.New("No table name is matched")
	}
}
