package telegram

import (
	"log"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Init initialize bot session
func Init(token string) {
	var err error
	// connect to bot
	bot, err = tgbot.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// init channel for updates from TG
	var ucfg = tgbot.NewUpdate(0)
	ucfg.Timeout = 60
	updates, err := bot.GetUpdatesChan(ucfg)

	// read updates from channel
	for {
		select {
		case update := <-updates:
			route(update)
		}
	}
}
