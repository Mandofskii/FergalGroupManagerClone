package handlers

import (
	"FergalManagerClone/database"
	"FergalManagerClone/functions"
	"FergalManagerClone/globals"
	"strconv"

	"gopkg.in/telebot.v3"
)

func HandleNewMessages(ctx telebot.Context) error {
	var err error
	message := ctx.Message()
	textMessage := message.Text
	chat := ctx.Chat()
	chatType := chat.Type
	chatID := chat.ID
	stringChatID := strconv.Itoa(int(chatID))
	sender := message.Sender
	senderID := sender.ID
	stringSenderID := strconv.Itoa(int(senderID))
	if chatType == "private" {
		if textMessage == "/start" {
			err = ctx.Send(globals.StartAnswer, &telebot.SendOptions{ParseMode: "markdown", ReplyMarkup: globals.StartKeyboard})
		}
	} else {
		if database.SIsMember("installedGroups", stringChatID) && database.SIsMember("group:"+stringChatID+":admins", stringSenderID) {
			if textMessage == "راهنما" {
				err = ctx.Send(globals.HelpAnswer, &telebot.SendOptions{ReplyMarkup: globals.HelpKeyboard})
			}
		}
	}
	functions.HandleError(err)
	return nil
}
