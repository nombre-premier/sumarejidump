package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antonholmquist/jason"
)

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

	totalCount, err := v.GetInt64("total_count")
	if err != nil {
		return nil, err
	}

	refResponse := SumarejiRefResponse{TotalCount: int(totalCount), Result: result}

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
