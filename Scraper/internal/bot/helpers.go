package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func forInput(bot *tgbotapi.BotAPI, chatID int64, text, state string) {

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "markdown"
	_, _ = bot.Send(msg)
	UserState[chatID] = state
}
