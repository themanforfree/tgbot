package main

import (
	"HautBot/tgbot"
	"flag"
)

var token = flag.String("t", "", "Telegram bot token")

var proxy = tgbot.ProxyFlag("p", "", "Telegram bot proxy")

func main() {
	flag.Parse()
	bs := tgbot.New(*token, *proxy)
	bs.Run()
}
