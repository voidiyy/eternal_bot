package bot

import (
	tgb "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

// Command menu at the bottom of the page
func commandMenu(bot *tgb.BotAPI, chatID int64) {
	commands := tgb.NewReplyKeyboard(
		tgb.NewKeyboardButtonRow(
			tgb.NewKeyboardButton("/start"),
			tgb.NewKeyboardButton("/help"),
			tgb.NewKeyboardButton("/info"),
		),
	)

	msg := tgb.NewMessage(chatID, "Use the buttons below to interact with the bot:")
	msg.ReplyMarkup = commands
	_, _ = bot.Send(msg)
}

// Main menu
func createMainMenuKeyboard() tgb.InlineKeyboardMarkup {
	menu := tgb.NewInlineKeyboardMarkup(
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("📱 Telegram Options", "tg_"),
			tgb.NewInlineKeyboardButtonData("📁 Files Options", "files_"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("🔍 Search Options", "search_"),
			tgb.NewInlineKeyboardButtonData("🌐 Web Options", "web_"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("👤 Profile", "profile"),
		),
	)
	return menu
}

func sendMainMenuMessage(bot *tgb.BotAPI, chatID int64) {
	text := "👋 *Welcome!*\n\nChoose an option:"
	menu := createMainMenuKeyboard()

	msg := tgb.NewMessage(chatID, text)
	msg.ParseMode = "Markdown"
	msg.ReplyMarkup = menu

	sentMsg, err := bot.Send(msg)
	if err != nil {
		log.Println("Error sending start menu message:", err)
	} else {
		// Save the message ID for future reference
		UserState[chatID] = strconv.Itoa(sentMsg.MessageID)
	}
}

// Files menu
func filesMenu(bot *tgb.BotAPI, chatID int64, messageID int) {
	text := "📁 *Files Menu* 📁\n\nChoose an option and then send the file to the bot:"
	menu := tgb.NewInlineKeyboardMarkup(
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("ℹ️ File Info", "files_info"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("🗜️ Create Archive", "files_archive_create"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("🗂️ Extract Archive", "files_archive_extract"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("🛡️ File Check via VirusTotal", "files_virus_total"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("🔢 File Hash Sum", "files_hash_sum"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("✏️ Write File Data to Chat", "files_write"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("🔙 Back", "main_menu"),
		),
	)

	editMessage(bot, chatID, messageID, text, menu)
}

// Search menu
func searchMenu(bot *tgb.BotAPI, chatID int64, messageID int) {
	text := "🔍 *Search Menu* 🔍\n\nChoose an option and provide the necessary information:"
	menu := tgb.NewInlineKeyboardMarkup(
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("📸 Search via Photo", "search_photo"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("🌐 Search in Social Media", "search_social_media"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("📞 Search Phone Info", "search_phone_info"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("🔍 Search Phone/Email in DBs", "search_phone_email_db"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("🔙 Back", "main_menu"),
		),
	)

	editMessage(bot, chatID, messageID, text, menu)
}

// Web menu
func webMenu(bot *tgb.BotAPI, chatID int64, messageID int) {
	text := "🌐 *Web Menu* 🌐\n\nChoose an option and then provide the necessary information:"
	menu := tgb.NewInlineKeyboardMarkup(
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("🌍 Check IP", "web_ip_info"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("🛑 IP Blacklist Check", "web_ip_check"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("🔍 Check Domain", "web_domain_info"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("📞 Check Phone", "web_phone_info"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("📵 Phone Blacklist Check", "web_phone_check"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("📸 URL Screenshot", "web_url_screen"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("🔗 URL Check Connection", "web_url_check_connection"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("✂️ URL Shortener", "web_url_cut"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("🔙 Back", "main_menu"),
		),
	)

	editMessage(bot, chatID, messageID, text, menu)
}

// Telegram menu
func telegramMenu(bot *tgb.BotAPI, chatID int64, messageID int) {
	text := "📱 *Telegram Menu* 📱\n\nChoose an option and then send the necessary information or forward a message:"
	menu := tgb.NewInlineKeyboardMarkup(
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("ℹ️ Probiv in Telegram", "tg_probiv"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("👤 User Page Info", "tg_user_info"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("🤖 Bot Info", "tg_bot_info"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("📢 Channels/Groups Info", "tg_channels_groups_info"),
		),
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("🔙 Back", "main_menu"),
		),
	)

	editMessage(bot, chatID, messageID, text, menu)
}

// Edit existing message
func editMessage(bot *tgb.BotAPI, chatID int64, messageID int, text string, menu tgb.InlineKeyboardMarkup) {
	editMsg := tgb.NewEditMessageTextAndMarkup(chatID, messageID, text, menu)
	editMsg.ParseMode = "Markdown"
	_, _ = bot.Send(editMsg)
}
