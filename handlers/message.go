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
				// database.SAdd("group:"+stringChatID+":panels", stringMessageID)
				sendedMessage, err := ctx.Bot().Send(chat, globals.HelpAnswer, &telebot.SendOptions{ReplyMarkup: globals.HelpKeyboard})
				database.Set("group:"+stringChatID+":panel:"+strconv.Itoa(sendedMessage.ID)+":owner", stringSenderID)
				functions.HandleError(err)
			}
		}
	}
	functions.HandleError(err)
	return nil
}
