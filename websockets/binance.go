package websockets

import (
	"fmt"
	"log"

	"github.com/farazsahebdel/exchange-websocket/setting"
	"github.com/tidwall/gjson"
	"golang.org/x/net/websocket"
)

var conn *websocket.Conn

type Result struct {
	Data   Info   `json;"data"`
	Result string `json:"result"`
}

type Info struct {
	LastPrice string `json:"c"`
	CloseTime int    `json:"C"`
	Symbol    string `json:"s"`
	Change    string `json:"P"`
}

func Binance() {
	connect()

	message := `{"method":"SUBSCRIBE", "params":["btcusdt@ticker", "ethusdt@ticker"], "id":1}`
	sendMessage(message)

	go receiverMessage()

}

func connect() {
	c, err := websocket.Dial(setting.C.SOCKET_BINANCE_URL, "", "http://localhost")
	if err != nil {
		log.Fatalln("Not connect to websocket binance", err)
	}
	conn = c
}

func sendMessage(message string) {
	err := websocket.JSON.Send(conn, gjson.Parse(message).Value())
	if err != nil {
		fmt.Println("Error Send Message Binance: ", err)
	}
}

func receiverMessage() {

	for {
		item := Result{}

		if err := websocket.JSON.Receive(conn, &item); err != nil {
			fmt.Println("Error Receive Message Binance", err)
			continue
		}

		if item.Data.Symbol == "" {
			continue
		}

		fmt.Println(item.Data.Symbol + " = " + item.Data.LastPrice)
	}
}
