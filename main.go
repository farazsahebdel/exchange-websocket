package main

import (
	"github.com/farazsahebdel/exchange-websocket/setting"
	"github.com/farazsahebdel/exchange-websocket/websockets"
)

func main() {

	setting.LoadConfig()

	websockets.Binance()

	// To run forever if exists goroutine
	ch := make(chan int)
	<-ch
}
