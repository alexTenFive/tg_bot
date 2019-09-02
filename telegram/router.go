package telegram

import (
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

type userBalance struct {
	inited  bool
	balance int
}

var balances = make(map[int64]*userBalance)

func route(u tgbot.Update) {

	var msg tgbot.Chattable

	if u.Message == nil && u.CallbackQuery == nil {
		return
	}

	if u.CallbackQuery != nil {

		bet, _ := strconv.Atoi(u.CallbackQuery.Data)

		if bet > balances[u.CallbackQuery.Message.Chat.ID].balance {
			bot.AnswerCallbackQuery(tgbot.NewCallbackWithAlert(u.CallbackQuery.ID, "У Вас недостаточно денег для такой ставки"))
			return
		} else {
			_, err := bot.AnswerCallbackQuery(tgbot.NewCallback(u.CallbackQuery.ID, u.CallbackQuery.Data))
			if err != nil {
				log.Panic(err)
			}
		}

		msg = rollDices(u.CallbackQuery.Message.Chat.ID, u.CallbackQuery.Message.MessageID, bet)
	} else if u.Message != nil {

		if ub, f := balances[u.Message.Chat.ID]; !f || (f && !ub.inited) {
			balances[u.Message.Chat.ID] = &userBalance{true, 200}
		}

		msg = Run(u.Message.Text, u.Message.Chat.ID)
	}

	bot.Send(msg)
}
