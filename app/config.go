package app

import (
	"flag"
	"os"
)

type Configuration struct {
	ENV         string
	WS_URI      string
	INSTRUMENTS string
	WINDOW_SIZE int
}

var configuration = Configuration{}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getArgs() (int, string) {
	ws := flag.Int("WindowSize", 200, "Number of data points to consider for indicators")
	instruments := flag.String("TradingPairs", "BTC-USD",
		"Specify trading instruments. You can comma separate for multiple instruments, ex: BTC-USD")

	flag.Parse()

	return *ws, *instruments
}

func LoadConfiguration() {
	configuration.ENV = getEnv("ENV", "dev")

	// Coinbase config
	configuration.WS_URI = getEnv("WS_ENDPOINT", "wss://ws-feed.prime.coinbase.com")

	// Load config from CLI args
	configuration.WINDOW_SIZE, configuration.INSTRUMENTS = getArgs()
}
