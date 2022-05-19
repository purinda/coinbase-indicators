package tests

import (
	"coinbase-indicators/indicator"
	"coinbase-indicators/types"
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestIndicatorBuilder(t *testing.T) {
	t.Run("Try initilising VWAP indicator", func(t *testing.T) {
		indicator := indicator.BuildIndicator(indicator.VWAP)

		ct := reflect.TypeOf(indicator)
		assert.Equal(t, ct.String(), "*indicator.Vwap", "VWAP indicator initialisation")
	})
}

func TestCalculation(t *testing.T) {
	t.Run("Test VWAP aggregator calculation", func(t *testing.T) {
		vwap := indicator.BuildIndicator(indicator.VWAP)

		tradeData := []types.TradeData{
			{
				Instrument: "BTC-USD",
				Volume:     decimal.NewFromFloat(1),
				Price:      decimal.NewFromFloat(1),
			},
		}

		for _, td := range tradeData {
			vwap.Receive(<-td)
		}
	})
}
