package models

import (
	"strconv"
	"time"
	"wishr_telegram_bot/configs"
	"wishr_telegram_bot/utils"
)

type BotUser struct {
	ID                    uint `gorm:"primarykey"`
	TelegramId            string
	FirstName             string
	LastName              *string
	Username              *string
	LanguageCode          *string
	IsPremium             bool
	AddedToAttachmentMenu bool
	AllowsWriteToPm       bool
	Status                bool
	HasLeft               bool `gorm:"default:false"`
	PersonalChat          *string
	BusinessConnection    *string
	DateCreated           time.Time
}

func (b BotUser) GetByTelegramId(tgID int64) BotUser {
	var user BotUser
	configs.DB.Where("telegram_id = ?", strconv.FormatInt(tgID, 10)).Find(&user)
	return user
}

func (b BotUser) TelegramIDAsInt64() int64 {
	i, err := strconv.ParseInt(b.TelegramId, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func (b BotUser) GetLang() string {
	if b.LanguageCode != nil && utils.StringInSlice(*b.LanguageCode, configs.LanguagesList) {
		return *b.LanguageCode
	} else {
		return configs.LangRu
	}
}

func (b BotUser) GetFullName() string {
	name := b.FirstName
	if b.LastName != nil {
		name += " " + *b.LastName
	}
	return name
}

func (BotUser) TableName() string {
	return "users_botuser"
}
