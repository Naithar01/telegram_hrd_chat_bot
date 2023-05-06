package main

import (
	"log"

	"github.com/Naithar01/go_hrd_telegram_bot/chat"
	"github.com/Naithar01/go_hrd_telegram_bot/create"
)

func main() {
	bot, updates, err := create.NewBot()

	if err != nil {
		log.Fatal(err.Error())
	}

	chat.StartChat(bot, updates)
	return
}
