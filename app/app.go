package app

import (
	"coinbase-indicators/exchange"
	"coinbase-indicators/types"
	"context"
	"fmt"
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

	for match := range td {
		fmt.Println(match.Instrument)
		fmt.Println(match.Price)
	}
}

func handleApplicationClose(appTerminate func()) {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	interrupt := <-exit
	log.Printf("Exiting due to %s", interrupt)

	appTerminate()
	os.Exit(0)
}
