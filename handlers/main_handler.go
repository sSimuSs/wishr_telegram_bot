package handlers

import (
	"wishr_telegram_bot/handlers/telegram_handlers"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"
)

func HandleTelegramUpdates(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
	for update := range updates {
		HandleTelegramSingleUpdate(update, bot)
	}
}

func HandleTelegramSingleUpdate(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	var mainHandler = telegram_handlers.Handler{
		Update: update,
		Bot:    bot,
	}
	if update.Message != nil {
		if update.Message.IsCommand() {
			handler := new(telegram_handlers.CommandHandler)
			handler.Handler = mainHandler
			handler.Handle()
		}
	}
}
