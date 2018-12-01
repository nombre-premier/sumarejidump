package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	contractID := os.Getenv("SUMAREJI_CONTRACT_ID")
	accessToken := os.Getenv("SUMAREJI_ACCESS_TOKEN")
	client := NewSumarejiClient(contractID, accessToken)

	params := SumarejiRefParams{ProcName: "category_ref",
		Fields: []string{}, Conditions: []string{},
		Order: []string{}, Limit: 1000, Page: 1,
		TableName: "Category",
	}
	resp, _ := client.Request(params)

	empData := []*Category{}
	cw, err := NewCSVWriter(empData, "output/Category.csv")
	defer cw.Close()
	if err != nil {
		panic(err)
	}
	parsedData := []*Category{}

	for _, r := range resp.Result {
		c := Category{}
		json.Unmarshal([]byte(r.String()), &c)
		parsedData = append(parsedData, &c)
	}
	fmt.Println(parsedData)
	cw.Write(parsedData)
}
