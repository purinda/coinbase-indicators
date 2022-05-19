package indicator_test

import (
	"coinbase-indicators/indicator"
	"coinbase-indicators/types"
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestIndicatorBuilder(t *testing.T) {
	t.Run("Try initilising VWAP indicator", func(t *testing.T) {
		indicator := indicator.BuildIndicator(indicator.VWAP, 10)

		ct := reflect.TypeOf(indicator)
		assert.Equal(t, ct.String(), "*indicator.Vwap", "VWAP indicator initialisation")
	})
}

func TestCalculation(t *testing.T) {
	t.Run("Test VWAP aggregator calculation", func(t *testing.T) {
		v := indicator.BuildIndicator(indicator.VWAP, 10)
		td := make(chan types.TradeData, 6)

		// Sample data from https://school.stockcharts.com/doku.php?id=technical_indicators:vwap_intraday
		tradeData := []types.TradeData{
			{
				Instrument: "BTC-USD",
				Volume:     decimal.NewFromFloat(89329),
				Price:      decimal.NewFromFloat(127.21),
			},
			{
				Instrument: "BTC-USD",
				Volume:     decimal.NewFromFloat(16137),
				Price:      decimal.NewFromFloat(127.17),
			},
			{
				Instrument: "BTC-USD",
				Volume:     decimal.NewFromFloat(23945),
				Price:      decimal.NewFromFloat(127.16),
			},
			{
				Instrument: "BTC-USD",
				Volume:     decimal.NewFromFloat(20679),
				Price:      decimal.NewFromFloat(127.04),
			},
			{
				Instrument: "BTC-USD",
				Volume:     decimal.NewFromFloat(27252),
				Price:      decimal.NewFromFloat(127.01),
			},
			{
				Instrument: "BTC-USD",
				Volume:     decimal.NewFromFloat(20915),
				Price:      decimal.NewFromFloat(127.08),
			},
		}

		go func() {
			for _, data := range tradeData {
				td <- data
			}

			// Couldn't find a better way to retrieve private "unexported" fields from a package
			field := reflect.ValueOf(v).Elem().FieldByName("cumulativeData")
			VWAP := reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface().(indicator.VWAPData).VWAP
			fmt.Println(VWAP)
		}()

		for {
			v.Receive(td)
		}
	})
}
