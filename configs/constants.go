package configs

import "os"

var BOT_TOKEN = os.Getenv("BOT_TOKEN")

var LangRu = "ru"
var LangEn = "en"
var LangUz = "uz"

var LanguagesList = []string{LangRu, LangEn, LangUz}

var MonifiMiniAppUrl = "https://wishr.uz/tg/"
