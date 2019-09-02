package menu

import tgbot "github.com/go-telegram-bot-api/telegram-bot-api"

// MainMenu Главное меню
var MainMenu = tgbot.NewReplyKeyboard(
	tgbot.NewKeyboardButtonRow(
		tgbot.NewKeyboardButton(MENU_RESTART),
		tgbot.NewKeyboardButton(MENU_NEW_GAME),
	),
)

// GameMenu Меню игры
var GameMenu = tgbot.NewReplyKeyboard(
	tgbot.NewKeyboardButtonRow(
		tgbot.NewKeyboardButton(GAME_MENU_BALANCE),
		tgbot.NewKeyboardButton(GAME_MENU_ROLL),
		tgbot.NewKeyboardButton(TO_MAIN_MENU),
	),
)

// MakeBetMenu for set bet on game
var MakeBetMenu = tgbot.NewInlineKeyboardMarkup(tgbot.NewInlineKeyboardRow(
	tgbot.NewInlineKeyboardButtonData("10", "10"),
	tgbot.NewInlineKeyboardButtonData("15", "15"),
	tgbot.NewInlineKeyboardButtonData("30", "30"),
),
	tgbot.NewInlineKeyboardRow(
		tgbot.NewInlineKeyboardButtonData("50", "50"),
		tgbot.NewInlineKeyboardButtonData("100", "100"),
		tgbot.NewInlineKeyboardButtonData("150", "150"),
	))
