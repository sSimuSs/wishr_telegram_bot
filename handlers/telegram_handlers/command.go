package telegram_handlers

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"
	"wishr_telegram_bot/configs"
	"wishr_telegram_bot/messages"
	"wishr_telegram_bot/models"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CommandHandler struct {
	Handler
}

func (h CommandHandler) Handle() interface{} {
	methodName := h.Update.Message.Command()
	caser := cases.Title(language.English, cases.NoLower)
	methodName = caser.String(methodName)
	st := reflect.ValueOf(h)
	Method := st.MethodByName(methodName)
	if !Method.IsValid() {
		return h.NotFoundMessage()
	} else {
		return reflect.ValueOf(&h).MethodByName(methodName).Call([]reflect.Value{})[0].Interface()
	}
}

func (h CommandHandler) Start() *tgbotapi.Message {
	var botUser = models.BotUser{}.GetByTelegramId(h.Update.Message.From.ID)
	if botUser.ID == 0 {
		newBotUser := models.BotUser{
			FirstName:             h.Update.Message.From.FirstName,
			LastName:              &h.Update.Message.From.LastName,
			Username:              &h.Update.Message.From.UserName,
			TelegramId:            strconv.FormatInt(h.Update.Message.From.ID, 10),
			LanguageCode:          &h.Update.Message.From.LanguageCode,
			DateCreated:           time.Now(),
			IsPremium:             h.Update.Message.From.IsPremium,
			AddedToAttachmentMenu: h.Update.Message.From.AddedToAttachmentMenu,
			AllowsWriteToPm:       true,
			Status:                true,
			HasLeft:               false,
		}
		if result := configs.DB.Create(&newBotUser); result.Error != nil {
			log.Printf(result.Error.Error())
		} else {
			botUser = newBotUser
		}
	} else if botUser.HasLeft {
		if result := configs.DB.Model(&botUser).Update("has_left", false); result.Error != nil {
			log.Printf("Error on making BotUser.has_left = false: %s", result.Error.Error())
		}
	}
	var keyboard [][]tgbotapi.InlineKeyboardButton
	keyboard = append(keyboard, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonWebApp(
			messages.LaunchInlineBtn[botUser.GetLang()],
			tgbotapi.WebAppInfo{URL: configs.MonifiMiniAppUrl},
		),
	))
	msg := tgbotapi.NewMessage(
		h.Update.Message.From.ID,
		messages.WelcomeToBot[botUser.GetLang()],
	)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(keyboard...)
	h.sendTelegram(msg)
	return nil
}

func (h CommandHandler) Info() *tgbotapi.Message {
	user := h.Update.Message.From
	name := user.FirstName
	if user.LastName != "" {
		name += " " + user.LastName
	}
	username := "-"
	if user.UserName != "" {
		username = fmt.Sprintf("@%s", user.UserName)
	}
	text := fmt.Sprintf("<b>%s</b>\n\nUsername: %s\nTelegram ID: <code>%d</code>",
		name, username, user.ID,
	)
	msg := tgbotapi.NewMessage(
		h.Update.Message.From.ID,
		text,
	)
	msg.ParseMode = tgbotapi.ModeHTML
	h.sendTelegram(msg)
	return nil
}
