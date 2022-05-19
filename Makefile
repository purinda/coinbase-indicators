test-indicators:
	go test -v ./indicator

test-exchange:
	go test -v ./exchange

debug:
	go run . -Indicator=printer

run: 
	EXCHANGE=coinbase \
	WS_ENDPOINT="wss://ws-feed.exchange.coinbase.com" \

	go run . -Indicator=vwap

build:
	@go build -o coinbase-indicators .
	@echo "Built, run ./coinbase-indicators"

clean:
	@rm -rf coinbase-indicators
	@echo "Removed ./coinbase-indicators"
