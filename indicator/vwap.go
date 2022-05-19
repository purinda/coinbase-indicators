package indicator

import (
	"coinbase-indicators/types"
	"container/list"
	"os"
	"reflect"

	"github.com/lensesio/tableprinter"
	"github.com/shopspring/decimal"
)

type VWAPPrintable struct {
	Instrument string `header:"Instrument"`
	VWAP       string `header:"VWAP"`
}
type VWAPData struct {
	CumulativeVolumeXPrice map[string]decimal.Decimal
	CumulativeVolume       map[string]decimal.Decimal
	VWAP                   map[string]decimal.Decimal
}

type Vwap struct {
	windowSize     int
	dataSeries     map[string]*list.List
	cumulativeData VWAPData
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
	c.cumulativeData = VWAPData{
		CumulativeVolumeXPrice: make(map[string]decimal.Decimal),
		CumulativeVolume:       make(map[string]decimal.Decimal),
		VWAP:                   make(map[string]decimal.Decimal),
	}

	c.dataSeries = make(map[string]*list.List)
	printer := c.printHeader()

	for trade := range td {
		// Ignore zero volume trades
		if trade.Volume.LessThanOrEqual(decimal.NewFromInt(0)) {
			continue
		}

		// Initialise cumulative data struct lists per instrument
		if c.cumulativeData.CumulativeVolume[trade.Instrument].IsZero() {
			c.dataSeries[trade.Instrument] = list.New()
		}

		c.dataSeries[trade.Instrument].PushBack(trade)

		if c.dataSeries[trade.Instrument].Len() == c.windowSize {
			e := c.dataSeries[trade.Instrument].Front()
			c.dataSeries[trade.Instrument].Remove(e)

		}

		c.aggregator(trade.Instrument, &c.cumulativeData)

		data := VWAPPrintable{
			Instrument: trade.Instrument,
			VWAP:       c.cumulativeData.VWAP[trade.Instrument].StringFixed(5),
		}

		c.printRow(printer, data)
	}

}

func (c *Vwap) printHeader() tableprinter.Printer {
	// Table printer to stdout
	printer := tableprinter.New(os.Stdout)
	v := reflect.ValueOf(VWAPPrintable{})

	// Render Table Header
	headers := tableprinter.StructParser.ParseHeaders(v)
	printer.Render(headers, nil, nil, false)

	return *printer
}

func (c *Vwap) printRow(printer tableprinter.Printer, data VWAPPrintable) {
	v := reflect.ValueOf(data)
	row, nums := tableprinter.StructParser.ParseRow(v)
	printer.RenderRow(row, nums)
}
