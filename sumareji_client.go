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

const CATEGORY = "Category"
const STORE = "Store"

type SumarejiClient struct {
	endpoint    string
	contractID  string
	accessToken string
	client      *http.Client
}

type SumarejiRefParams struct {
	ProcName   string   `json:"-"`
	Fields     []string `json:"fields"`
	Conditions []string `json:"conditions"`
	Order      []string `json:"order"`
	Limit      int      `json:"limit"`
	Page       int      `json:"page"`
	TableName  string   `json:"table_name"`
}

type SumarejiRefResponse struct {
	TotalCount int
	Result     []*jason.Object
}

func NewSumarejiClient(contractID string, accessToken string) SumarejiClient {
	return SumarejiClient{
		endpoint:    "https://webapi.smaregi.jp/access/",
		contractID:  contractID,
		accessToken: accessToken,
		client:      &http.Client{},
	}
}

func parseRefResponse(resp *http.Response) (*SumarejiRefResponse, error) {
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
	refResponse := SumarejiRefResponse{TotalCount: totalCount, Result: result}

	return &refResponse, nil
}

func (sc *SumarejiClient) Request(params SumarejiRefParams) (*SumarejiRefResponse, error) {
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

func (sc *SumarejiClient) DumpTableToCSV(params SumarejiRefParams) (*CSVWriter, error) {
	switch params.TableName {
	case CATEGORY:
		empData := []*Category{}
		resultBuffer := make([]*Category, params.Limit, params.Limit)

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
	default:
		return nil, errors.New("No table name is matched")
	}
}
