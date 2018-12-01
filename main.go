package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/antonholmquist/jason"
)

type SumarejiRefParams struct {
	Fields     []string `json:"fields"`
	Conditions []string `json:"conditions"`
	Order      []string `json:"order"`
	Limit      int      `json:"limit"`
	Page       int      `json:"page"`
	TableName  string   `json:"table_name"`
}

func main() {
	urlStr := "https://webapi.smaregi.jp/access/"
	contractID := os.Getenv("SUMAREJI_CONTRACT_ID")
	accessToken := os.Getenv("SUMAREJI_ACCESS_TOKEN")
	params := SumarejiRefParams{Fields: []string{}, Conditions: []string{}, Order: []string{}, Limit: 1000, Page: 1, TableName: "Category"}

	jsonBytes, err := json.Marshal(params)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}

	data := url.Values{}
	data.Set("proc_name", "category_ref")
	data.Add("params", string(jsonBytes))

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("X_contract_id", contractID)
	r.Header.Add("X_access_token", accessToken)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	fmt.Println(resp.Status)
	b, err := ioutil.ReadAll(resp.Body)
	v, err := jason.NewObjectFromBytes(b)
	result, err := v.GetObjectArray("result")

	empData := []*Category{}
	cw, err := NewCSVWriter(empData, "output/Category.csv")
	defer cw.Close()
	if err != nil {
		panic(err)
	}
	parsedData := []*Category{}

	for _, r := range result {
		c := Category{}
		json.Unmarshal([]byte(r.String()), &c)
		parsedData = append(parsedData, &c)
	}
	fmt.Println(parsedData)
	cw.Write(parsedData)

}
