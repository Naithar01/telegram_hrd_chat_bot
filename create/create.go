package create

import (
	"os"

	tgbodyapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func NewBot() (*tgbodyapi.BotAPI, tgbodyapi.UpdatesChannel, error) {
	b, u, err := CreateNewUpdate()

	if err != nil {
		return nil, nil, err
	}

	return b, u, nil
}

func CreateNewBot() (*tgbodyapi.BotAPI, error) {
	// Load Env file
	err := godotenv.Load(".env")

	if err != nil {
		return nil, err
	}

	// Create new bot * use env ( api key )
	new_bot, err := tgbodyapi.NewBotAPI(os.Getenv("APIKEY"))

	if err != nil {
		return nil, err
	}

	// Debug mode on
	// new_bot.Debug = false

	// Print Bot Name || ex) @jeongho0730
	// log.Println(new_bot.Self.UserName)

	return new_bot, nil
}

func CreateNewUpdate() (*tgbodyapi.BotAPI, tgbodyapi.UpdatesChannel, error) {
	new_bot, err := CreateNewBot()

	if err != nil {
		return nil, nil, err
	}

	u := tgbodyapi.NewUpdate(0)
	u.Timeout = 60

	updates := new_bot.GetUpdatesChan(u)

	return new_bot, updates, nil
}
