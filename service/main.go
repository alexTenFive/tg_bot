package main

import (
	"tg_gamble/config"
	"tg_gamble/telegram"
)

var (
	token = config.Cfg.TelegramToken
)

func main() {
	telegram.Init(token)
}
