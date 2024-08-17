package bot

import (
	"fmt"
	tgb "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

// Handle bot commands
func handleCommand(bot *tgb.BotAPI, message *tgb.Message) {
	switch message.Command() {
	case "info":
		msg := tgb.NewMessage(message.Chat.ID, "ℹ️ *Info command*\n\nThis command provides information.")
		msg.ParseMode = "Markdown"
		_, _ = bot.Send(msg)
	case "start":
		UserState[message.Chat.ID] = "main_menu"

		// Send command menu
		commandMenu(bot, message.Chat.ID)

		// Create and send the main menu message
		sendMainMenuMessage(bot, message.Chat.ID)
	case "help":
		msg := tgb.NewMessage(message.Chat.ID, "❓ *Help command*\n\nUse /start to begin interacting with the bot.")
		msg.ParseMode = "Markdown"
		_, _ = bot.Send(msg)
	case "clear":
		// Delete previous messages
		for i := message.MessageID - 1; i > 0; i-- {
			deleteMsg := tgb.NewDeleteMessage(message.Chat.ID, i)
			_, err := bot.Request(deleteMsg)
			if err != nil {
				log.Println("Error deleting message:", err)
				break
			}
		}

		// Delete the /clear command message
		clearCommandMsg := tgb.NewDeleteMessage(message.Chat.ID, message.MessageID)
		_, err := bot.Request(clearCommandMsg)
		if err != nil {
			log.Println("Error deleting /clear command message:", err)
		}
	default:
		msg := tgb.NewMessage(message.Chat.ID, fmt.Sprintf("❗*Unknown command*: %s\nTry /help for a list of commands.", message.Command()))
		msg.ParseMode = "Markdown"
		_, _ = bot.Send(msg)
	}
}
