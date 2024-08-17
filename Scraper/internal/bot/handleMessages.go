package bot

import tgb "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func handleMessage(bot *tgb.BotAPI, message *tgb.Message) {
	chatID := message.Chat.ID
	text := message.Text

	if message.IsCommand() {
		handleCommand(bot, message)
		return
	}

	switch UserState[chatID] {
	case "wait_web_ip_info":
		msg := tgb.NewMessage(chatID, "🔍 IP Information:\nYour input: "+text)
		msg.ParseMode = "Markdown"
		_, _ = bot.Send(msg)
		// Handle the IP info processing here
	case "wait_web_ip_check":
		msg := tgb.NewMessage(chatID, "🛑 IP Blacklist Check:\nYour input: "+text)
		msg.ParseMode = "Markdown"
		_, _ = bot.Send(msg)
		// Handle the IP check processing here
	case "wait_web_domain_info":
		msg := tgb.NewMessage(chatID, "🔍 Domain Information:\nYour input: "+text)
		msg.ParseMode = "Markdown"
		_, _ = bot.Send(msg)
		// Handle the domain info processing here
	case "wait_web_phone_info":
		msg := tgb.NewMessage(chatID, "📞 Phone Information:\nYour input: "+text)
		msg.ParseMode = "Markdown"
		_, _ = bot.Send(msg)
		// Handle the phone info processing here
	case "wait_web_phone_check":
		msg := tgb.NewMessage(chatID, "📵 Phone Blacklist Check:\nYour input: "+text)
		msg.ParseMode = "Markdown"
		_, _ = bot.Send(msg)
		// Handle the phone check processing here
	case "wait_web_url_screen":
		msg := tgb.NewMessage(chatID, "📸 URL Screenshot:\nYour input: "+text)
		msg.ParseMode = "Markdown"
		_, _ = bot.Send(msg)
		// Handle the URL screenshot processing here
	case "wait_web_url_check_connection":
		msg := tgb.NewMessage(chatID, "🔗 URL Connection Check:\nYour input: "+text)
		msg.ParseMode = "Markdown"
		_, _ = bot.Send(msg)
		// Handle the URL connection check processing here
	case "wait_web_url_cut":
		msg := tgb.NewMessage(chatID, "✂️ URL Shortener:\nYour input: "+text)
		msg.ParseMode = "Markdown"
		_, _ = bot.Send(msg)
		// Handle the URL shortening processing here
	}
}
