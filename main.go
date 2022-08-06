package main

import (
	"FergalManagerClone/globals"
	"FergalManagerClone/handlers"

	"gopkg.in/telebot.v3"
)

func main() {
	globals.Bot.Handle(telebot.OnText, handlers.HandleNewMessages)
	globals.Bot.Handle(telebot.OnCallback, handlers.HandleNewCallbackQuery)
	globals.Bot.Handle(telebot.OnChatMember, handlers.NewChatMemberHandler)
	globals.Bot.Handle(telebot.OnMyChatMember, handlers.NewMyChatMemberHandler)
	globals.Bot.Start()
}
