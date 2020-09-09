package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type getOrder struct {
	Status             string  `json:"status"`
	Side               string  `json:"side"`
	Price              float64 `json:"price"`
	Quantity           float64 `json:"quantity"`
	OrderID            string  `json:"order_id"`
	ClientOID          string  `json:"client_oid"`
	CreateTime         int64   `json:"create_time"`
	UpdateTime         int64   `json:"update_time"`
	Type               string  `json:"type"`
	InstrumentName     string  `json:"instrument_name"`
	AvgPrice           float64 `json:"avg_price"`
	CumulativeQuantity float64 `json:"cumulative_quantity"`
	CumulativeValue    float64 `json:"cumulative_value"`
	FeeCurrency        string  `json:"fee_currency"`
}

type getOrdersResult struct {
	Orders []getOrder `json:"order_list"`
}

type getOrdersResponse struct {
	ID     int    `json:"id"`
	Method string `json:"method"`
	Code   int    `json:"code"`
	Result getOrdersResult
}

func getOrdersMessage() []byte {

	p := make(map[string]string)
	// p["end_ts"] = fmt.Sprintf("%v", time.Now().Add(time.Hour*-2).UnixNano()/1000000)
	// p["instrument_name"] = "BTC_USDT"

	endpoint := "private/get-order-history"
	sig := getSignature(getSignatureRequest{
		ID:     1,
		Method: endpoint,
		Params: p,
		APIKey: os.Getenv("API_KEY"),
		Nonce:  time.Now().UnixNano() / 1000000,
	})

	j, _ := json.Marshal(sig)

	return j

}

func getCancelOrderMessage(id string, instrument string) []byte {
	p := make(map[string]string)
	p["order_id"] = id
	p["instrument_name"] = instrument

	endpoint := "private/cancel-order"
	sig := getSignature(getSignatureRequest{
		ID:     1,
		Method: endpoint,
		Params: p,
		APIKey: os.Getenv("API_KEY"),
		Nonce:  time.Now().UnixNano() / 1000000,
	})

	j, _ := json.Marshal(sig)

	return j
}

func getCreateOrderMessage(instrument string, side string, typer string, price float64, q float64) []byte {
	p := make(map[string]string)
	p["instrument_name"] = instrument
	p["side"] = side
	p["type"] = typer
	p["price"] = fmt.Sprintf("%0.2f", price)
	p["quantity"] = fmt.Sprintf("%v", q)

	// if typer == "MARKET" && side == "BUY" {
	// 	p["notional"] = fmt.Sprintf("%v", q)
	// }

	endpoint := "private/create-order"
	sig := getSignature(getSignatureRequest{
		ID:     2,
		Method: endpoint,
		Params: p,
		APIKey: os.Getenv("API_KEY"),
		Nonce:  time.Now().UnixNano() / 1000000,
	})

	j, _ := json.Marshal(sig)

	return j
}
