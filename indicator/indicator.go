package indicator

import (
	"coinbase-indicators/types"
	"container/list"

	"log"
)

const (
	PRINTER = "printer"
	VWAP    = "vwap"
)

type Indicator interface {
	Receive(td chan types.TradeData)
}

type IndicatorFactory func(ex string, windowSize int) Indicator

func create(ex string, windowSize int) Indicator {
	switch ex {
	case PRINTER:
		return &Printer{}
	case VWAP:
		return &Vwap{
			windowSize:     windowSize,
			dataSeries:     map[string]*list.List{},
			cumulativeData: VWAPData{},
		}
	default:
		log.Fatalf("Can't find an indicator of type: %s", ex)
	}

	return nil
}

func BuildIndicator(ex string, windowSize int) Indicator {
	var cf IndicatorFactory = create
	i := cf(ex, windowSize)

	return i
}
