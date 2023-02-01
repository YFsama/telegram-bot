package commands

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Getid(message tgbotapi.Message) string {
	if message.ReplyToMessage == nil {
		return "User ID: " + strconv.FormatInt(message.From.ID, 10)
	} else {
		return "User ID: " + strconv.FormatInt(message.ReplyToMessage.From.ID, 10)
	}
}
