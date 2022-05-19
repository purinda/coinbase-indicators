package main

import (
	"coinbase-indicators/app"
)

func main() {
	config := app.LoadConfiguration()

	app.Run(config)
}
