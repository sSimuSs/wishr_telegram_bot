package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"wishr_telegram_bot/configs"
	"wishr_telegram_bot/handlers"

	api "github.com/OvyFlash/telegram-bot-api"
)

func main() { webhookEcho() }

func webhookEcho() {
	configs.RunProject()

	bot, err := api.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	tokenParts := strings.Split(bot.Token, ":")
	// Listen for updates from the webhook
	updates := bot.ListenForWebhook("/tg_bot/" + tokenParts[0])
	go http.ListenAndServe("0.0.0.0:8101", nil)

	// Set the webhook
	//webHook, err := api.NewWebhook("https://monifi.uz/monifi_bot/" + tokenParts[0])
	//if err != nil {
	//	panic(err)
	//}
	//
	//apiResponse, err := bot.Request(webHook)
	//if err != nil {
	//	panic(err)
	//}
	//
	//if apiResponse.Ok {
	//	log.Printf("Webhook set successfully")
	//} else {
	//	log.Printf("Failed to set webhook: %s", apiResponse.Description)
	//}
	//
	//info, err := bot.GetWebhookInfo()
	//if err != nil {
	//	panic(err)
	//}
	//
	//if info.LastErrorDate != 0 {
	//	log.Printf("Failed to get webhook info: %s", info.LastErrorMessage)
	//}

	for update := range updates {
		handlers.HandleTelegramSingleUpdate(update, bot)
	}
}
