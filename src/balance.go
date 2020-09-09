package main

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// type currencyBalance struct {
// 	Currency  string `json:"currency"`
// 	Available string `json:"available"`
// 	Hold      string `json:"hold"`
// 	Balance   string `json:"balance"`
// }

// type balanceResponse struct {
// 	Code    string            `json:"code"`
// 	Data    []currencyBalance `json:"data"`
// 	Message string            `json:"message"`
// }

// func getBalance(currency string) currencyBalance {
// 	r, err := http.NewRequest("GET", apiURL+"/v1/account/getBalance", nil)

// 	// ms := getTimestamp()

// 	sign := getSha256(secretkey, "GET", "/v1/account/getBalance", "", ms)
// 	r.Header.Add("X-Nova-Access-Key", accesskey)
// 	r.Header.Add("X-Nova-Signature", sign)
// 	r.Header.Add("X-Nova-Timestamp", ms)
// 	c := &http.Client{}

// 	resp, err := c.Do(r)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer resp.Body.Close()

// 	bs, _ := ioutil.ReadAll(resp.Body)
// 	b := &balanceResponse{}
// 	json.Unmarshal(bs, b)

// 	for _, v := range b.Data {
// 		if v.Currency == currency {
// 			log.Printf("%+v %s", v, "|| account/getBalance/")
// 			return v
// 		}
// 		continue
// 	}

// 	return currencyBalance{}
// }
