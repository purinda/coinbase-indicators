package main

import (
	"coinbase-indicators/app"
)

func main() {
	app.LoadConfiguration()
	app.Run()
}
