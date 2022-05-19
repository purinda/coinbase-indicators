package indicator

import (
	"coinbase-indicators/types"
	"container/list"
	"fmt"

	"github.com/shopspring/decimal"
)

type Vwap struct {
	dataSeries map[string]*list.List
}

type VWAPData struct {
	CumulativeVolumeXPrice map[string]decimal.Decimal
	CumulativeVolume       map[string]decimal.Decimal
	VWAP                   map[string]decimal.Decimal
}

func (c *Vwap) aggregator(instrument string, d *VWAPData) {
	var cv decimal.Decimal
	var cvp decimal.Decimal

	for e := c.dataSeries[instrument].Front(); e != nil; e = e.Next() {
		trade := e.Value.(types.TradeData)
		cv = cv.Add(trade.Volume)
		cvp = cvp.Add(trade.Volume.Mul(trade.Price))
	}

	d.CumulativeVolume[instrument] = cv
	d.CumulativeVolumeXPrice[instrument] = cvp
	d.VWAP[instrument] = cvp.Div(cv)
}

func (c *Vwap) Receive(td chan types.TradeData) {
	// Struct to keep running totals
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

		c.aggregator(trade.Instrument, &cData)

		fmt.Printf("Instrument: %s, Trade Vol: %s , Cumulative Vol: %s , VWAP: %s \n",
			trade.Instrument,
			trade.Volume.String(),
			cData.CumulativeVolume[trade.Instrument].String(),
			cData.VWAP[trade.Instrument])
	}

}
