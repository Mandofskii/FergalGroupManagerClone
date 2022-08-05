package globals

import (
	"time"

	"FergalManagerClone/functions"

	"gopkg.in/telebot.v3"
)

// Bot
var (
	Token       string              = "5434857575:AAGmL7_QzF1EdRqsMrdpJ3mYNBJr0snVYAM"
	BotPoller   *telebot.LongPoller = &telebot.LongPoller{Timeout: 10 * time.Second}
	BotSettings telebot.Settings    = telebot.Settings{Token: Token, Poller: BotPoller}
	Bot         *telebot.Bot        = functions.ReturnBot(telebot.NewBot(BotSettings))
	DB          int                 = 9
)
