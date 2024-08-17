package bot

import (
	tgb "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

func handleCallbackQuery(bot *tgb.BotAPI, callbackQuery *tgb.CallbackQuery) {
	data := callbackQuery.Data
	chatID := callbackQuery.Message.Chat.ID
	messageID := callbackQuery.Message.MessageID

	switch {
	case strings.HasPrefix(data, "web_"):
		msg := tgb.NewMessage(chatID, usageWeb)
		msg.ParseMode = tgb.ModeMarkdown
		bot.Send(msg)
		webMenu(bot, chatID, messageID)
		switch data {
		case "web_ip_info":
			forInput(bot, chatID, "Please input IP address", "wait_web_ip_info")
		case "web_ip_check":
			forInput(bot, chatID, "Please input IP address", "wait_web_ip_check")
		case "web_domain_info":
			forInput(bot, chatID, "Please input domain", "wait_web_domain_info")
		case "web_phone_info":
			forInput(bot, chatID, "Please input phone number", "wait_web_phone_info")
		case "web_phone_check":
			forInput(bot, chatID, "Please input phone number", "wait_web_phone_check")
		case "web_url_screen":
			forInput(bot, chatID, "Please input URL", "wait_web_url_screen")
		case "web_url_check_connection":
			forInput(bot, chatID, "Please input URL", "wait_web_url_check_connection")
		case "web_url_cut":
			forInput(bot, chatID, "Please input URL", "wait_web_url_cut")
		case "main_menu":
			editMessage(bot, chatID, messageID, "👋 *Welcome!*\n\nChoose an option:", createMainMenuKeyboard())
		}
	case strings.HasPrefix(data, "tg_"):
		msg := tgb.NewMessage(chatID, usageTelegram)
		msg.ParseMode = tgb.ModeMarkdown
		bot.Send(msg)
		telegramMenu(bot, chatID, messageID)
	case strings.HasPrefix(data, "files_"):
		msg := tgb.NewMessage(chatID, usageFiles)
		msg.ParseMode = tgb.ModeMarkdown
		bot.Send(msg)
		filesMenu(bot, chatID, messageID)
	case strings.HasPrefix(data, "search_"):
		msg := tgb.NewMessage(chatID, usageSearch)
		msg.ParseMode = tgb.ModeMarkdown
		bot.Send(msg)
		searchMenu(bot, chatID, messageID)
	case strings.HasPrefix(data, "main_menu"):
		editMessage(bot, chatID, messageID, "👋 *Welcome!*\n\nChoose an option:", createMainMenuKeyboard())
	}

	callback := tgb.NewCallback(callbackQuery.ID, "")
	_, err := bot.Request(callback)
	if err != nil {
		log.Println("Error sending callback response:", err)
	}
}
