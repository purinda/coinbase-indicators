# Coinbase Indicator 

Purpose of the project is to evaluates price action indicators on real-time trading data which gets produced on Coinbase exchange. 

# Usage

```
🚀  ./coinbase-indicators --help
Usage of ./coinbase-indicators:
  -Indicator printer
        Indicator to be used for trade data interpretation (possible values: printer, `vwap`) (default "printer")
  -TradingPairs string
        Specify trading instruments. You can comma separate for multiple instruments, ex: BTC-USD (default "BTC-USD")
  -WindowSize int
        Number of data points to consider for indicators (default 200)
```

## How to run

Build first

        make build

### Runnable Configurations

1. Run without configuration. This mode uses `Printer` indicator to output the websocket feed comes from `Coinbase` for `BTC-USD` instrument.

        ./coinbase-indicators

2. Run with `vwap` indicator with data sampling with `200` data points and use data from `Coinbase` feed.

        ./coinbase-indicators -Indicator=vwap

3. Run with `vwap` indicator with data sampling with `200` data points and use data from `Coinbase` feed for `BTC-USD, ETH-USD, ETH-BTC` instruments.

        ./coinbase-indicators -Indicator=vwap -WindowSize=200 -TradingPairs="btc-usd,eth-usd,eth-btc"

4. Running tests

        make test-exchange   # runs tests for exchange module
        make test-indicators # runs tests for indicators
# Design

Application design is done in a way to satisfy the requirement of needing to connect to `datasource` to retrieve course of sale or
"trade data" and run that data through single or multiple `indicators` which can be used for processing received data then redirect to stdout or other data streams.

# Exchange Connector and Indicator

Concept of a `Exchange Connector` and `Indicator` is used as modules which can be either changed easily using application configuration or extended to add functionality without modifying underlying application framework.

- **Exchange Connector**

  This allows application to consume trade data from existing connector such as `Coinbase` websocket connection or build a similar connector for another provider and switch between the connectors without needing to modify the application.

- **Indicator**

  Indicator is a module based approach to apply a mathematical function on a stream of data to process and feed it back to another stream. For the simplicity of the application, currently uses the `os.stdout` as output stream.

  Supported indicators: 
   - `vwap`: calculated the Volume Weighted Average Price of the data feed based on a specific period / window.
   - `printer`: debug indicator which doesn't apply any function, rather it outputs to `stdout` to be viewed and for debugging 
                data streams.

# Improvements

- This is my first Go application and some Go concepts applied in the application development can be improved (such as methods on structs and use of interfaces)
- Need additional tests further code coverage.
- Application was written without knowing how Golang test framework work, hence ran into trouble of implementing some tests to cover indicator implementation.