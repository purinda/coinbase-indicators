package exchange

import (
	"coinbase-indicators/types"
	"context"
	"log"

	"github.com/gorilla/websocket"
)

const (
	COINBASE = "coinbase"
	BINANCE  = "binance"
	IBKR     = "interactivebrokers"
)

type Exchange interface {
	Feed(td chan<- types.TradeData)
	Disconnect()
}

type Coinbase struct {
	ws          *websocket.Conn
	ws_url      string
	instruments []string
	ctx         context.Context
}

type ExchangeFactory func(ex string,
	ws_url string,
	instruments []string,
	ctx context.Context) Exchange

func create(t string, ws_url string, instruments []string, ctx context.Context) Exchange {
	switch t {
	case COINBASE:
		return &Coinbase{
			ws:          &websocket.Conn{},
			ws_url:      ws_url,
			instruments: instruments,
			ctx:         ctx,
		}
	default:
		log.Fatalf("Can't find the connector implementation for exchange: %s", t)
	}

	return nil
}

func BuildExchange(ex string, ws_url string, instruments []string, ctx context.Context) Exchange {
	var cf ExchangeFactory = create
	connector := cf(ex, ws_url, instruments, ctx)

	return connector
}
