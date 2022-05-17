package types

import "github.com/shopspring/decimal"

/*
 * Struct type to be used for exchanging data between exchange specific and
 * core application logic data requirements.
 */
type TradeData struct {
	Instrument string
	Volume     decimal.Decimal
	Price      decimal.Decimal
}
