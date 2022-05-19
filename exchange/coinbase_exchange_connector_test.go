package exchange_test

import (
	"coinbase-indicators/exchange"
	"coinbase-indicators/types"
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectorBuilder(t *testing.T) {
	t.Run("Try initilising COINBASE connector", func(t *testing.T) {
		ctx, _ := context.WithCancel(context.TODO())

		instruments := make([]string, 3)
		instruments[0] = "BTC-USD"
		connector := exchange.BuildExchange(exchange.COINBASE, "wss://localhost", instruments, ctx)

		ct := reflect.TypeOf(connector)
		assert.Equal(t, ct.String(), "*exchange.Coinbase", "Coinbase connector initialisation")
	})
}

// End-to-end connection test
func TestCoinbaseConnection(t *testing.T) {
	t.Parallel()

	t.Run("Test COINBASE Feed", func(t *testing.T) {
		ctx, _ := context.WithCancel(context.TODO())

		instruments := make([]string, 3)
		instruments[0] = "BTC-USD"
		instruments[1] = "ETH-USD"
		coinbase := exchange.BuildExchange(exchange.COINBASE, "wss://ws-feed.exchange.coinbase.com", instruments, ctx)

		td := make(chan types.TradeData)
		coinbase.Feed(td)

		for tradeData := range td {
			if tradeData.Instrument == instruments[0] {
				assert.Equal(t, tradeData.Instrument, instruments[0], "Correct instrument symbol received")
				break
			}

			if tradeData.Instrument == instruments[1] {
				assert.Equal(t, tradeData.Instrument, instruments[1], "Correct instrument symbol received")
				break
			}
		}
	})
}
