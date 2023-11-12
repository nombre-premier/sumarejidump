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

type SrError struct {
	ErrorCode        int    `json:"error_code"`
	ErrorSummary     string `json:"error"`
	ErrroDescription string `json:"error_description"`
}

func (s *SrError) Error() string {
	jsonBytes, _ := json.MarshalIndent(s, "", "    ")
	return string(jsonBytes)
}

func NewSrErrorFromJSONBytes(jsonBytes []byte) *SrError {
	s := new(SrError)
	json.Unmarshal(jsonBytes, s)
	return s
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

	if resp.StatusCode != 200 {
		srErr := NewSrErrorFromJSONBytes(b)
		return nil, srErr
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

func (sc *SrClient) DumpTableToCSV(p SrRefParams) (*CSVWriter, error) {
	output := fmt.Sprintf("%s/%s.csv", sc.config.OutputDir, p.TableName)

	handler, err := chooseCSVHandler(p, output)
	if err != nil {
		return nil, err
	}
	defer handler.GetCSVWriter().Close()
	isAll := false
	// Conditionsの最初の要素に検索条件が入る。検索条件が無い場合は全件検索
	if len(p.Conditions[0]) == 0 {
		isAll = true
	}
	// StockHistory、TransactionHead、TransactionDetailテーブルで全件検索の場合は範囲指定をする
	if (p.TableName == "StockHistory" || p.TableName == "TransactionHead" || p.TableName == "TransactionDetail") && isAll {
		startNum := 0
		// 1回あたりの指定可能範囲は10万件まで。idで10万までを範囲指定する
		endNum := 100000
		for {
			var lastId int
			startNumStr := strconv.Itoa(startNum)
			endNumStr := strconv.Itoa(endNum)
			// 10万件の範囲指定＆id昇順にソートして取得
			if p.TableName == "StockHistory" {
				p.Order = []string{"id"}
				p.Conditions[0] = map[string]*string{"id >": &startNumStr, "id <=": &endNumStr}
			} else {
				p.Order = []string{"transactionHeadId"}
				p.Conditions[0] = map[string]*string{"transactionHeadId >": &startNumStr, "transactionHeadId <=": &endNumStr}
			}

			for {
				resp, err := sc.Request(p)
				if err != nil {
					return nil, err
				}
				handler.Write(resp)

				// 指定範囲の全件数よりLimit*page数が大きくなった場合、最後の要素のIDを保持しておく
				if resp.TotalCount <= p.Limit*p.Page {
					dat := map[string]string{}
					lastObj := resp.Result[len(resp.Result)-1]
					if err := json.Unmarshal([]byte(lastObj.String()), &dat); err != nil {
						panic(err)
					}
					if p.TableName == "StockHistory" {
						lastId, _ = strconv.Atoi(dat["id"])
					} else {
						lastId, _ = strconv.Atoi(dat["transactionHeadId"])
					}
					// page数をリセット
					p.Page = 1
					break
				} else {
					p.Page = p.Page + 1
				}
			}
			// 次の10万件取得用の範囲を設定
			if lastId == endNum {
				startNum = endNum
				endNum += 100000
			} else {
				break
			}
		}
	} else {
		for {
			resp, err := sc.Request(p)
			if err != nil {
				return nil, err
			}
			handler.Write(resp)

			if resp.TotalCount <= p.Limit*p.Page {
				break
			} else {
				p.Page = p.Page + 1
			}
		}
	}

	return handler.GetCSVWriter(), nil
}

func (sc *SrClient) DumpTableToParquet(p SrRefParams) (*ParquetWriter, error) {
	output := fmt.Sprintf("%s/%s.snappy.parquet", sc.config.OutputDir, p.TableName)

	handler, err := chooseParquetHandler(p, output)
	if err != nil {
		return nil, fmt.Errorf("failed to choose parquet handler: %w", err)
	}
	defer handler.GetParquetWriter().Close()

	for {
		resp, err := sc.Request(p)
		// TODO: remove this line
		fmt.Printf("Processing %d / %d\n", p.Limit*p.Page, resp.TotalCount)
		if err != nil {
			return nil, fmt.Errorf("failed to request with params: %w", err)
		}
		handler.Write(resp)

		if resp.TotalCount <= p.Limit*p.Page {
			break
		} else {
			p.Page = p.Page + 1
		}
	}

	return handler.GetParquetWriter(), nil
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
		return NewProductReserveItemCSV(p.Limit, output)
	case PRODUCT_RESERVE_ITEM_LABEL:
		return NewProductReserveItemLabelCSV(p.Limit, output)
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
	case SHIPPING:
		return NewShippingCSV(p.Limit, output)
	case SHIPPING_DETAIL:
		return NewShippingDetailCSV(p.Limit, output)
	case RECEIVING:
		return NewReceivingCSV(p.Limit, output)
	case RECEIVING_DETAIL:
		return NewReceivingDetailCSV(p.Limit, output)
	case STOCKTAKING_INFO:
		return NewStocktakingInfoCSV(p.Limit, output)
	case STOCKTAKING_HEAD:
		return NewStocktakingHeadCSV(p.Limit, output)
	case STOCKTAKING_DETAIL:
		return NewStocktakingDetailCSV(p.Limit, output)
	case LOSS:
		return NewLossCSV(p.Limit, output)
	case LOSS_DETAIL:
		return NewLossDetailCSV(p.Limit, output)
	case SHIPMENT:
		return NewShipmentCSV(p.Limit, output)
	case SHIPMENT_DETAIL:
		return NewShipmentDetailCSV(p.Limit, output)
	case STORAGE_INFO:
		return NewStorageInfoCSV(p.Limit, output)
	case STORAGE_INFO_DELIVERY:
		return NewStorageInfoDeliveryCSV(p.Limit, output)
	case STORAGE_INFO_PRODUCT:
		return NewStorageInfoProductCSV(p.Limit, output)
	case STORAGE_INFO_DELIVERY_PRODUCT:
		return NewStorageInfoDeliveryProductCSV(p.Limit, output)
	case STORAGE:
		return NewStorageCSV(p.Limit, output)
	case STORAGE_DETAIL:
		return NewStorageDetailCSV(p.Limit, output)
	case BUDGET_DAILY:
		return NewBudgetDailyCSV(p.Limit, output)
	default:
		return nil, errors.New("no table name is matched")
	}
}

func chooseParquetHandler(p SrRefParams, output string) (SrParquetHandlerIf, error) {
	switch p.TableName {
	case CATEGORY:
		return NewSrGenericParquet[CategoryParquetSchema](output)
	case STORE:
		return NewSrGenericParquet[StoreParquetSchema](output)
	case PRODUCT:
		return NewSrGenericParquet[ProductParquetSchema](output)
	case PRODUCT_PRICE:
		return NewSrGenericParquet[ProductPriceParquetSchema](output)
	case PRODUCT_RESERVE_ITEM:
		return NewSrGenericParquet[ProductReserveItemParquetSchema](output)
	case PRODUCT_RESERVE_ITEM_LABEL:
		return NewSrGenericParquet[ProductReserveItemLabelParquetSchema](output)
	case PRODUCT_STORE:
		return NewSrGenericParquet[ProductStoreParquetSchema](output)
	case PRODUCT_INVENTORY_RESERVATION:
		return NewSrGenericParquet[ProductInventoryReservationParquetSchema](output)
	case CUSTOMER:
		return NewSrGenericParquet[CustomerParquetSchema](output)
	// case STOCK:
	// case STOCK_HISTORY:
	// case TRANSACTION_HEAD:
	// case TRANSACTION_DETAIL:
	// case BARGAIN:
	// case BARGAIN_PRODUCT:
	// case BARGAIN_STORE:
	// case DAILY_SUM:
	// case SHIPPING:
	// case SHIPPING_DETAIL:
	// case RECEIVING:
	// case RECEIVING_DETAIL:
	// case STOCKTAKING_INFO:
	// case STOCKTAKING_HEAD:
	// case STOCKTAKING_DETAIL:
	// case LOSS:
	// case LOSS_DETAIL:
	// case SHIPMENT:
	// case SHIPMENT_DETAIL:
	// case STORAGE_INFO:
	// case STORAGE_INFO_DELIVERY:
	// case STORAGE_INFO_PRODUCT:
	// case STORAGE_INFO_DELIVERY_PRODUCT:
	// case STORAGE:
	// case STORAGE_DETAIL:
	case BUDGET_DAILY:
		return NewSrGenericParquet[BudgetDailyParquetSchema](output)
	default:
		return nil, errors.New("no table name is matched")
	}
}
