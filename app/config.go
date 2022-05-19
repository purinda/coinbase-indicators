package app

import (
	"flag"
	"os"
	"strings"
)

type Configuration struct {
	ENV         string
	WS_URI      string
	INDICATOR   string
	INSTRUMENTS []string
	WINDOW_SIZE int
}

var configuration = Configuration{}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getArgs() (int, string, string) {
	instruments := flag.String("TradingPairs", "BTC-USD",
		"Specify trading instruments. You can comma separate for multiple instruments, ex: BTC-USD")

	ws := flag.Int("WindowSize", 200, "Number of data points to consider for indicators")

	indicator := flag.String("Indicator", "printer", "Indicator to be used for trade data interpretation (possible values: `printer`, `vwap`)")

	flag.Parse()

	return *ws, *instruments, *indicator
}

func LoadConfiguration() Configuration {
	var instruments string
	configuration.ENV = getEnv("ENV", "dev")

	// Coinbase config
	configuration.WS_URI = getEnv("WS_ENDPOINT", "wss://ws-feed.exchange.coinbase.com")

	// Load config from CLI args
	configuration.WINDOW_SIZE, instruments, configuration.INDICATOR = getArgs()
	configuration.INSTRUMENTS = strings.Split(strings.ToUpper(instruments), ",")

	return configuration
}
