package indicator

import (
	"coinbase-indicators/types"
	"os"
	"reflect"

	"github.com/lensesio/tableprinter"
	"github.com/shopspring/decimal"
)

type Printer struct{}

type PrinterData struct {
	Instrument string `header:"Instrument"`
	Volume     string `header:"Volume"`
	Price      string `header:"Price"`
}

func (c *Printer) Receive(td chan types.TradeData) {
	printer := tableprinter.New(os.Stdout)
	v := reflect.ValueOf(PrinterData{})

	// Render Table Header
	headers := tableprinter.StructParser.ParseHeaders(v)
	printer.Render(headers, nil, nil, false)

	for match := range td {
		if match.Volume.LessThanOrEqual(decimal.NewFromInt(0)) {
			continue
		}

		data := PrinterData{
			Instrument: match.Instrument,
			Volume:     match.Volume.String(),
			Price:      match.Price.String(),
		}

		v = reflect.ValueOf(data)
		row, nums := tableprinter.StructParser.ParseRow(v)
		printer.RenderRow(row, nums)
	}
}
