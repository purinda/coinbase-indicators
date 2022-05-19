test:
	@go test -v ./tests 

debug:
	@go run . -Indicator=printer

run: 
	@go run . -Indicator=vwap

build:
	@go build -o coinbase-indicators .

clean:
	@rm -rf coinbase-indicators