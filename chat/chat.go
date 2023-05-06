package chat

import (
	"log"

	tgbodyapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	command_list = map[string]string{
		"help": "hello",
		"set1": "set 1",
		"set2": "set 2",
	}
)

func SendHelpMsg(new_bot *tgbodyapi.BotAPI, update tgbodyapi.Update) {
	msg := tgbodyapi.NewMessage(update.Message.Chat.ID, "Send /help")
	new_bot.Send(msg)
}

func StartChat(new_bot *tgbodyapi.BotAPI, updates tgbodyapi.UpdatesChannel) {

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			// log.Println(update.Message.Command())
			// log.Println(update.Message.CommandArguments())
			log.Println(command_list[update.Message.Command()])

		} else {
			SendHelpMsg(new_bot, update)
		}
	}
}
