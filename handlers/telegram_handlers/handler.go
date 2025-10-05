package telegram_handlers

import (
	"fmt"
	"log"
	"wishr_telegram_bot/configs"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"
)

type Handler struct {
	Update tgbotapi.Update
	Bot    *tgbotapi.BotAPI
}

func (h Handler) NotFoundMessage() *tgbotapi.Message {
	msg := tgbotapi.NewMessage(h.Update.Message.Chat.ID, h.Update.Message.Text)
	msg.Text = "What? 🤔"
	msg.ParseMode = tgbotapi.ModeMarkdownV2
	return h.sendTelegram(msg)
}

func (h Handler) GetUserAvatarUrl(user tgbotapi.User) *string {
	userAvas, err := h.Bot.GetUserProfilePhotos(tgbotapi.NewUserProfilePhotos(user.ID))
	if err != nil {
		log.Printf(err.Error())
	} else {
		if userAvas.TotalCount > 0 {
			fileURL := h.GetTelegramFileUrl(userAvas.Photos[0][1].FileID)
			if fileURL != nil {
				return fileURL
			}
		}
	}
	return nil
}

func (h Handler) GetTelegramFileUrl(FileId string) *string {
	firstAvaFile, err := h.Bot.GetFile(tgbotapi.FileConfig{FileID: FileId})
	if err != nil {
		log.Printf(err.Error())
	} else {
		avaUrl := firstAvaFile.Link(configs.BOT_TOKEN)
		return &avaUrl
	}
	return nil
}

func (h Handler) GetMessageFileId(message tgbotapi.Message) *string {
	var messageFileIdTG *string
	if len(message.Photo) > 0 {
		messageFileIdTG = &message.Photo[1].FileID
	} else if message.Voice != nil {
		messageFileIdTG = &message.Voice.FileID
	} else if message.Audio != nil {
		messageFileIdTG = &message.Audio.FileID
	} else if message.Document != nil {
		messageFileIdTG = &message.Document.FileID
	} else if message.Animation != nil {
		messageFileIdTG = &message.Animation.FileID
	} else if message.Video != nil {
		messageFileIdTG = &message.Video.FileID
	} else if message.VideoNote != nil {
		messageFileIdTG = &message.VideoNote.FileID
	} else if message.Sticker != nil {
		messageFileIdTG = &message.Sticker.FileID
	}
	return messageFileIdTG
}

func (h Handler) sendTelegram(msg tgbotapi.Chattable) *tgbotapi.Message {
	//if os.Getenv("APP_ENV") == constants.AppEnvTest {
	//	return nil
	//}
	if respMsg, err := h.Bot.Send(msg); err != nil {
		log.Printf(err.Error())
		fmt.Printf("%v", msg)
	} else {
		return &respMsg
	}
	return nil
}

func (h Handler) sendTelegramRequest(msg tgbotapi.Chattable, additionalEvent bool) {
	if _, err := h.Bot.Request(msg); err != nil {
		fmt.Println(err.Error())
		fmt.Printf("%v", msg)
	}
}
