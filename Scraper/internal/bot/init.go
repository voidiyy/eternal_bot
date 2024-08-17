package bot

import (
	tgb "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var UserState = make(map[int64]string)

func InitBot(Token string) error {
	bot, err := tgb.NewBotAPI(Token)
	if err != nil {
		return err
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgb.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			handleMessage(bot, update.Message)
		} else if update.CallbackQuery != nil {
			handleCallbackQuery(bot, update.CallbackQuery)
		}
	}

	return nil
}
