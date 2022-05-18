package indicator

import (
	"coinbase-indicators/types"

	"log"
)

const (
	PRINTER = "printer"
	VWAP    = "vwap"
)

type Indicator interface {
	Receive(td chan types.TradeData)
}

type Printer struct{}
type Vwap struct{}

type IndicatorFactory func(t string) Indicator

func create(t string) Indicator {
	switch t {
	case PRINTER:
		return &Printer{}
	case VWAP:
		return &Vwap{}
	default:
		log.Fatalf("Can't find an indicator of type: %s", t)
	}

	return nil
}

func BuildIndicator(ex string) Indicator {
	var cf IndicatorFactory = create
	i := cf(ex)

	return i
}
