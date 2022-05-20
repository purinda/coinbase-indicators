package types

import "github.com/shopspring/decimal"

/*
 * Struct type to be used for exchanging data between exchange specific and
 * core application logic data requirements.
 *
 * Also a way of generalising data comes from various exchanges to be used across
 * downsteam modules such as indicators which ultimately rely on volume and price data.
 */
type TradeData struct {
	Instrument string
	Volume     decimal.Decimal
	Price      decimal.Decimal
}
