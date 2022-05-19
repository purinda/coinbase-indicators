package indicator

import (
	"coinbase-indicators/types"
	"container/list"
	"fmt"

	"github.com/shopspring/decimal"
)

type VWAPData struct {
	CumulativeVolumeXPrice map[string]decimal.Decimal
	CumulativeVolume       map[string]decimal.Decimal
	VWAP                   map[string]decimal.Decimal
}

func (c *Vwap) sum(instrument string, d *VWAPData) {
	var result decimal.Decimal

	for e := c.dataSeries[instrument].Front(); e != nil; e = e.Next() {
		result = result.Add(e.Value.(types.TradeData).Volume)
	}

	d.CumulativeVolume[instrument] = result
}

func (c *Vwap) Receive(td chan types.TradeData) {
	// Struct to keep cumulative values
	var cData = VWAPData{
		CumulativeVolumeXPrice: make(map[string]decimal.Decimal),
		CumulativeVolume:       make(map[string]decimal.Decimal),
		VWAP:                   make(map[string]decimal.Decimal),
	}

	c.dataSeries = make(map[string]*list.List)

	for trade := range td {
		// Ignore zero volume trades
		if trade.Volume.LessThanOrEqual(decimal.NewFromInt(0)) {
			continue
		}

		// Initialise cumulative data struct lists per instrument
		if cData.CumulativeVolume[trade.Instrument].IsZero() {
			c.dataSeries[trade.Instrument] = list.New()
		}

		c.dataSeries[trade.Instrument].PushBack(trade)

		if c.dataSeries[trade.Instrument].Len() == 10 {
			e := c.dataSeries[trade.Instrument].Front()
			c.dataSeries[trade.Instrument].Remove(e)

		}

		c.sum(trade.Instrument, &cData)

		fmt.Printf("Index: %d, Trade Vol: %s , Cumulative Vol: %s \n", c.dataSeries[trade.Instrument].Len(),
			trade.Volume.String(),
			cData.CumulativeVolume[trade.Instrument].String())
	}

}
