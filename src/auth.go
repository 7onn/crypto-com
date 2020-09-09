package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type getSignatureRequest struct {
	ID     int               `json:"id"`
	Method string            `json:"method"`
	APIKey string            `json:"api_key"`
	Params map[string]string `json:"params"`
	Nonce  int64             `json:"nonce"`
}

type getSignatureResponse struct {
	ID     int               `json:"id"`
	Method string            `json:"method"`
	APIKey string            `json:"api_key"`
	Params map[string]string `json:"params"`
	Nonce  int64             `json:"nonce"`
	Sig    string            `json:"sig"`
}

func getSignature(sr getSignatureRequest) getSignatureResponse {
	ps := ""
	// p := make(map[string]string)
	// json.Unmarshal([]byte(sr.Params), &p)
	// for k, v := range p {
	for k, v := range sr.Params {
		ps = ps + k
		ps = ps + v
	}
	mac := hmac.New(sha256.New, []byte(os.Getenv("SECRET")))
	bs := []byte(fmt.Sprintf("%s%v%s%s%v", sr.Method, sr.ID, sr.APIKey, ps, sr.Nonce))
	// fmt.Printf("%v\n", string(bs))
	mac.Write(bs)
	// fmt.Printf("getSignature %+v \n", string(bs))
	// if sr.Method == "private/create-order" {
	// 	return getSignatureResponse{
	// 		ID:     sr.ID,
	// 		Method: sr.Method,
	// 		APIKey: sr.APIKey,
	// 		Params: sr.Params,
	// 		Nonce:  sr.Nonce,
	// 	}
	// }
	return getSignatureResponse{
		ID:     sr.ID,
		Method: sr.Method,
		APIKey: sr.APIKey,
		Sig:    hex.EncodeToString(mac.Sum(nil)),
		Params: sr.Params,
		Nonce:  sr.Nonce,
	}
}

func getAuthMessage() []byte {
	p := make(map[string]string)
	// p["end_ts"] = fmt.Sprintf("%v", time.Now().Add(time.Hour*-2).Unix())
	// p["instrument_name"] = "BTC_USDT"

	endpoint := "public/auth"
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
