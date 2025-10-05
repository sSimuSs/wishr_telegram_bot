package utils

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"
)

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func TgMessageToJson(update tgbotapi.Message) *string {
	b, err := json.Marshal(&update)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	convert := string(b)
	return &convert
}

func ToJson(target interface{}) *string {
	b, err := json.Marshal(&target)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	convert := string(b)
	return &convert
}

func StructToMap(data interface{}) map[string]interface{} {
	// Marshal the struct to JSON
	jsonData, _ := json.Marshal(data)

	// Unmarshal JSON to map[string]interface{}
	var mapData map[string]interface{}
	json.Unmarshal(jsonData, &mapData)

	return mapData
}
