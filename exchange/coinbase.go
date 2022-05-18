package exchange

import (
	"encoding/json"
	"log"
	"time"

	"coinbase-indicators/types"

	"github.com/gorilla/websocket"
	"github.com/shopspring/decimal"
)

type Subscribe struct {
	Type       string   `json:"type"`
	ProductIDs []string `json:"product_ids"`
	Channels   []string `json:"channels"`
}

type MatchesData struct {
	Type         string          `json:"type"`
	TradeID      int             `json:"trade_id"`
	MakerOrderID string          `json:"maker_order_id"`
	TakerOrderID string          `json:"taker_order_id"`
	Side         string          `json:"side"`
	Size         decimal.Decimal `json:"size"`
	Price        decimal.Decimal `json:"price"`
	ProductID    string          `json:"product_id"`
	Sequence     uint64          `json:"sequence"`
	Time         time.Time       `json:"time"`
}

func (c *Coinbase) subscribe() {
	subReq := Subscribe{Type: "subscribe"}
	subReq.ProductIDs = append(subReq.ProductIDs, "ETH-USD")
	subReq.ProductIDs = append(subReq.ProductIDs, "BTC-USD")
	subReq.Channels = append(subReq.Channels, "matches")

	// Error handling for encoding and subscribe request
	json, err := json.Marshal(subReq)
	if err != nil {
		log.Fatalf("Error occured while preparing JSON request body: %s", err.Error())
	}

	err = c.ws.WriteMessage(websocket.TextMessage, json)
	if err != nil {
		log.Fatalf("Couldn't write to websocket connection: ", err.Error())
	}

}

func (c *Coinbase) Receive(td chan<- types.TradeData) {
	log.Printf("Connecting to %s", "ws-feed.exchange.coinbase.com")
	ws, _, err := websocket.DefaultDialer.Dial("wss://ws-feed.exchange.coinbase.com", nil)
	c.ws = ws

	if err != nil {
		log.Fatal("dial:", err)
	}

	// Send the initial subscription request
	c.subscribe()

	go func() {
		for {
			select {
			case <-c.ctx.Done():
				// Cleanly close the connection by sending a close message and then
				// waiting (with timeout) for the server to close the connection.
				err := c.ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					log.Println("write close:", err)
					return
				}

				defer c.ws.Close()
			default:
				var md MatchesData
				err := c.ws.ReadJSON(&md)
				if err != nil {
					log.Println("Response Error:", err)
					continue
				}

				td <- mapTradeData(md)
			}
		}
	}()
}

func (c *Coinbase) Disconnect() {
	err := c.ws.Close()

	if err != nil {
		log.Println("Error disconnecting Coinbase connector:", err)
		return
	}
}

func mapTradeData(md MatchesData) types.TradeData {
	return types.TradeData{
		Instrument: md.ProductID,
		Volume:     md.Size,
		Price:      md.Price,
	}
}
