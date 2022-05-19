test:
	@go test -v ./indicator ./exchange

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
