package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type quoteData struct {
	Instrument        string  `json:"i"`
	BestBidPrice      float64 `json:"b"`
	BestAskPrice      float64 `json:"k"`
	LastestTradePrice float64 `json:"a"`
	Highest           float64 `json:"h"`
	Lowest            float64 `json:"l"`
	Timestamp         int64   `json:"t"`
}

type quoteResult struct {
	InstrumentName string    `json:"instrument_name"`
	Data           quoteData `json:"data"`
}

type tick struct {
	Code   int         `json:"code"`
	Method string      `json:"method"`
	Result quoteResult `json:"result"`
}

func (t tick) average() float64 {
	return (t.Result.Data.Highest + t.Result.Data.Lowest) / 2
}

func (t tick) bestBidPrice() float64 {
	return t.Result.Data.BestBidPrice
}

func (t tick) bestAskPrice() float64 {
	return t.Result.Data.BestAskPrice
}

func getTick(from string, to string) *tick {
	endpoint := "public/get-ticker"
	param := fmt.Sprintf("/?instrument_name=%s_%s", from, to)
	r, err := http.NewRequest("GET", apiURL+endpoint+param, nil)
	c := &http.Client{}
	resp, err := c.Do(r)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	bs, _ := ioutil.ReadAll(resp.Body)
	b := &tick{}
	json.Unmarshal(bs, b)
	log.Println(from, "|| bid:", b.bestBidPrice(), "||", "ask", b.bestAskPrice(), "|| ", "avg", b.average(), "|| public/get-ticker")
	return b
}
