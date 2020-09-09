package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var (
	apiURL        = "https://api.crypto.com/v2/"
	wsAPIURL      = "stream.crypto.com:443"
	accesskey     = ""
	secretkey     = ""
	authenticated = false
)

func main() {
	godotenv.Load(".env")
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: wsAPIURL, Path: "v2/user"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			orders := &getOrdersResponse{}
			json.Unmarshal(message, orders)

			if len(orders.Result.Orders) > 0 {
				now := time.Now().UnixNano() / 1000000
				for _, o := range orders.Result.Orders {
					//delete those hanging for an hour || minute
					// if o.Status == "ACTIVE" && now-o.CreateTime > 1000*60*60 {
					if o.Status == "ACTIVE" && now-o.CreateTime > 1000*60 {
						err := c.WriteMessage(websocket.TextMessage, getCancelOrderMessage(o.OrderID, o.InstrumentName))
						if err != nil {
							log.Println("getCancelOrderMessage error:", err)
							continue
						}

						if authenticated {
							t := getTick("BTC", "USDT")
							err = c.WriteMessage(websocket.TextMessage, getCreateOrderMessage("BTC_USDT", "SELL", "LIMIT", t.bestAskPrice(), 0.001))
							if err != nil {
								log.Println("getCreateOrderMessage error:", err)
								return
							}

							err = c.WriteMessage(websocket.TextMessage, getCreateOrderMessage("BTC_USDT", "BUY", "LIMIT", t.bestBidPrice(), 0.001))
							if err != nil {
								log.Println("getCreateOrderMessage error:", err)
								return
							}
						}
					}
				}
			}

			// fmt.Printf("buy BTC for %v \n", t.bestBidPrice())
			// fmt.Printf("sell BTC for %v \n", t.bestAskPrice())

			// tick := &getTickResponse{}
			// json.Unmarshal(message, tick)
			// fmt.Printf("%+v", tick)

			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second * 6)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case _ = <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, getAuthMessage())
			if err != nil {
				log.Println("auth error:", err)
				return
			}
			authenticated = true

			err = c.WriteMessage(websocket.TextMessage, getOrdersMessage())
			if err != nil {
				log.Println("getOrders error:", err)
				return
			}

		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
				// case <-time.After(time.Second * 7):
			}
			return
		}
	}

}
