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

type SrClient struct {
	endpoint    string
	contractID  string
	accessToken string
	client      *http.Client
}

type SrRefResponse struct {
	TotalCount int
	Result     []*jason.Object
}

func NewSrClient(contractID string, accessToken string) SrClient {
	return SrClient{
		endpoint:    "https://webapi.smaregi.jp/access/",
		contractID:  contractID,
		accessToken: accessToken,
		client:      &http.Client{},
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

	r, err := http.NewRequest("POST", sc.endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	r.Header.Add("X_contract_id", sc.contractID)
	r.Header.Add("X_access_token", sc.accessToken)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := sc.client.Do(r)
	if err != nil {
		return nil, err
	}
	return parseRefResponse(resp)
}

func (sc *SrClient) DumpTableToCSV(params SrRefParams) (*CSVWriter, error) {
	switch params.TableName {
	case CATEGORY:
		empData := []*Category{}
		resultBuffer := make([]*Category, params.Limit)

		cw, err := NewCSVWriter(empData, fmt.Sprintf("output/%s.csv", params.TableName))
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
	case STORE:
		empData := []*Store{}
		resultBuffer := make([]Store, 1, params.Limit)

		cw, err := NewCSVWriter(empData, fmt.Sprintf("output/%s.csv", params.TableName))
		if err != nil {
			panic(err)
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
	case PRODUCT:
		empData := []*Product{}
		resultBuffer := make([]Product, params.Limit)

		cw, err := NewCSVWriter(empData, fmt.Sprintf("output/%s.csv", params.TableName))
		if err != nil {
			panic(err)
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

	case PRODUCT_PRICE:
		empData := []*ProductPrice{}
		resultBuffer := make([]ProductPrice, params.Limit)

		cw, err := NewCSVWriter(empData, fmt.Sprintf("output/%s.csv", params.TableName))
		if err != nil {
			panic(err)
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
	case PRODUCT_RESERVE_ITEM:
		empData := []*ProductReseveItem{}
		resultBuffer := make([]ProductReseveItem, 1, params.Limit)

		cw, err := NewCSVWriter(empData, fmt.Sprintf("output/%s.csv", params.TableName))
		if err != nil {
			panic(err)
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
	case PRODUCT_RESERVE_ITEM_LABEL:
		empData := []*ProductReseveItemLabel{}
		resultBuffer := make([]ProductReseveItemLabel, params.Limit)

		cw, err := NewCSVWriter(empData, fmt.Sprintf("output/%s.csv", params.TableName))
		if err != nil {
			panic(err)
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
	case PRODUCT_STORE:
		empData := []*ProductStore{}
		resultBuffer := make([]ProductStore, params.Limit)

		cw, err := NewCSVWriter(empData, fmt.Sprintf("output/%s.csv", params.TableName))
		if err != nil {
			panic(err)
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
	case PRODUCT_INVENTORY_RESERVATION:
		empData := []*ProductInventoryReservation{}
		resultBuffer := make([]ProductInventoryReservation, params.Limit)

		cw, err := NewCSVWriter(empData, fmt.Sprintf("output/%s.csv", params.TableName))
		if err != nil {
			panic(err)
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
	case CUSTOMER:
		empData := []*Customer{}
		resultBuffer := make([]Customer, params.Limit)

		cw, err := NewCSVWriter(empData, fmt.Sprintf("output/%s.csv", params.TableName))
		if err != nil {
			panic(err)
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
