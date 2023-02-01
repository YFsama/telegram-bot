package main

import (
	"log"
	"os"
	"telegram-bot/src/commands"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = os.Getenv("TELEGRAM_DEBUG") == "true"

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = commands.Help()
		case "deepl":
			msg.Text = commands.Deepl(*update.Message)
		case "getid":
			msg.Text = commands.Getid(*update.Message)
			// case "whois":
			// msg.Text = commands.Whois(update.Message.CommandArguments())
		}

		if _, err := bot.Send(msg); err != nil {
			// log.Panic(err)
		}
	}
}
