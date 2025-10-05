package messages

import "wishr_telegram_bot/configs"

var LaunchInlineBtn map[string]string = map[string]string{
	configs.LangRu: "Запустить",
	configs.LangUz: "Ishga tushirish",
	configs.LangEn: "Launch",
}

var CreateDigitalProductInlineBtn map[string]string = map[string]string{
	configs.LangRu: "Создать продукт",
	configs.LangUz: "Mahsulot yaratish",
	configs.LangEn: "Create product",
}

var CancelInlineBtn map[string]string = map[string]string{
	configs.LangRu: "✖️ Отменить",
	configs.LangUz: "✖️ Bekor qilish",
	configs.LangEn: "✖️ Cancel",
}

var BuyInlineBtn map[string]string = map[string]string{
	configs.LangRu: "Купить",
	configs.LangUz: "Harid qilish",
	configs.LangEn: "Buy",
}
