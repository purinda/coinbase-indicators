package exchange

import (
	"coinbase-indicators/types"
	"context"

	"github.com/gorilla/websocket"
)

const (
	COINBASE = "coinbase"
	IBKR     = "interactivebrokers"
)

type Exchange interface {
	Receive(td chan<- types.TradeData)
	Disconnect()
}

type Coinbase struct {
	ws  *websocket.Conn
	ctx context.Context
}

type ExchangeFactory func(ctx context.Context) Exchange

func create(ctx context.Context) Exchange {
	return &Coinbase{
		ws:  &websocket.Conn{},
		ctx: ctx,
	}
}

func BuildExchange(ex string, ctx context.Context) Exchange {
	var cf ExchangeFactory = create
	connector := cf(ctx)

	return connector
}
