package app

import (
	"coinbase-indicators/exchange"
	"coinbase-indicators/indicator"
	"coinbase-indicators/types"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(config Configuration) {
	ctx, c := context.WithCancel(context.TODO())

	go handleApplicationClose(c)
	td := make(chan types.TradeData)

	// Initialise exchange connector
	connector := exchange.BuildExchange(config.EXCHANGE, config.WS_URI, config.INSTRUMENTS, ctx)
	connector.Feed(td)

	// Initialise configured indicator
	indicator := indicator.BuildIndicator(config.INDICATOR, config.WINDOW_SIZE)
	indicator.Receive(td)
}

func handleApplicationClose(appTerminate func()) {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	interrupt := <-exit
	log.Printf("Exiting due to %s", interrupt)

	appTerminate()
	os.Exit(0)
}
