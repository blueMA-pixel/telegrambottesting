package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"telegramBotTesting/cmd"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	token := "7870899044:AAEJ8iZ43olWFZlHpQrMtH_0naZZR7LTQkw"
	encodedToken := url.QueryEscape(token)
	fmt.Println(encodedToken)
	bot, err := tgbotapi.NewBotAPI(encodedToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			cmd.RootCmd.SetArgs(strings.Split(update.Message.Text, " "))

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, cmd.Execute())
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
