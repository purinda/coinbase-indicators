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

func Run() {
	ctx, c := context.WithCancel(context.TODO())

	go handleApplicationClose(c)
	td := make(chan types.TradeData)

	connector := exchange.BuildExchange(exchange.COINBASE, ctx)
	connector.Receive(td)

	indicator := indicator.BuildIndicator(indicator.PRINTER)
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
