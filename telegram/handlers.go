package telegram

import (
	"fmt"
	"log"
	"sync"
	"tg_gamble/games/dices"
	"tg_gamble/telegram/menu"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

type handler func(int64) tgbot.Chattable

var mutex sync.Mutex

var handlers = make(map[string]handler)

func init() {
	handlers[menu.MENU_NEW_GAME] = newGame
	handlers[menu.MENU_RESTART] = restartGame

	handlers[menu.GAME_MENU_ROLL] = throwDices
	handlers[menu.GAME_MENU_BALANCE] = checkBalance

	handlers[menu.TO_MAIN_MENU] = backToMain
}

// Run is method for handling incoming message from bot
func Run(text string, chatID int64) tgbot.Chattable {
	log.Printf("Run button [Text: %s][CharID: %d]\n", text, chatID)

	method, exists := handlers[text]
	if exists {
		return method(chatID)
	}

	msg := tgbot.NewMessage(chatID, text)
	msg.ReplyMarkup = menu.MainMenu

	return msg
}

func newGame(chatID int64) tgbot.Chattable {
	mutex.Lock()
	defer mutex.Unlock()

	text := fmt.Sprintf(
		`
		Добро пожаловать в игру!
			
		Баланс: %d 💵
		`, balances[chatID].balance)
	msg := tgbot.NewMessage(chatID, text)
	msg.ReplyMarkup = menu.GameMenu

	return msg
}

func restartGame(chatID int64) tgbot.Chattable {
	mutex.Lock()
	defer mutex.Unlock()

	balances[chatID].balance = 200

	text := fmt.Sprintf(
		`
	Игровые фишки сброшены!
	Наслаждайтесь игрой.
	
	Баланс: %d 💵
	`, balances[chatID].balance)

	msg := tgbot.NewMessage(chatID, text)
	msg.ReplyMarkup = menu.MainMenu

	return msg
}

func checkBalance(chatID int64) tgbot.Chattable {
	mutex.Lock()
	defer mutex.Unlock()

	text := fmt.Sprintf(
		`
	Ваш баланс: %d 💵
		`, balances[chatID].balance)
	msg := tgbot.NewMessage(chatID, text)
	msg.ReplyMarkup = menu.GameMenu

	return msg
}

func backToMain(chatID int64) tgbot.Chattable {
	msg := tgbot.NewMessage(chatID, `
	Назад
	`)
	msg.ReplyMarkup = menu.MainMenu

	return msg
}

func throwDices(chatID int64) tgbot.Chattable {
	msg := tgbot.NewMessage(chatID, `
	Выберите ставку
	`)
	msg.ReplyMarkup = menu.MakeBetMenu

	return msg
}

func rollDices(chatID int64, messageID int, bet int) tgbot.Chattable {
	var (
		resText string
	)

	computer := dices.Dices{}
	user := dices.Dices{}

	computer.Throw()
	user.Throw()

	if user.Result() > computer.Result() {
		resText = fmt.Sprintf("Вы выиграли %d 💵!", bet*2)

		{
			mutex.Lock()
			balances[chatID].balance += bet
			mutex.Unlock()
		}

	} else {
		resText = fmt.Sprintf("Вы проиграли %d 💵:(", bet)

		{
			mutex.Lock()
			balances[chatID].balance -= bet
			mutex.Unlock()
		}
	}

	text := fmt.Sprintf(
		`
	Вы поставили по %d 💵.

	🤜🏾 Компьютер бросает кости...
	Выпало: 🎲 %d и 🎲 %d = 🎲🎲 %d
	
	
	🤜 Вы бросаете кости...
	Выпало: 🎲 %d и 🎲 %d = 🎲🎲 %d
	
	%s
		`,
		bet,
		computer.First, computer.Second, computer.Result(),
		user.First, user.Second, user.Result(),
		resText)

	edit := tgbot.NewEditMessageText(
		chatID,
		messageID,
		text)

	return edit
}
