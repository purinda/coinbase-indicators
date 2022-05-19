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
	ws  *websocket.Conn
	ctx context.Context
}

type ExchangeFactory func(t string, ctx context.Context) Exchange

func create(t string, ctx context.Context) Exchange {
	switch t {
	case COINBASE:
		return &Coinbase{
			ws:  &websocket.Conn{},
			ctx: ctx,
		}
	default:
		log.Fatalf("Can't find the connector implementation for exchange: %s", t)
	}

	return nil
}

func BuildExchange(ex string, ctx context.Context) Exchange {
	var cf ExchangeFactory = create
	connector := cf(ex, ctx)

	return connector
}
