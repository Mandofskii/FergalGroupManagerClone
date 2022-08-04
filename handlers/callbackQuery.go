package handlers

import (
	"FergalManagerClone/database"
	"FergalManagerClone/functions"
	"FergalManagerClone/globals"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/telebot.v3"
)

func HandleNewCallbackQuery(ctx telebot.Context) error {
	var err error
	callbackQuery := ctx.Callback()
	callbackData := callbackQuery.Data
	chatID := callbackQuery.Message.Chat.ID
	stringChatID := strconv.Itoa(int(chatID))
	stringMessageID := strconv.Itoa(ctx.Callback().Message.ID)
	senderID := ctx.Callback().Sender.ID
	stringSenderID := strconv.Itoa(int(senderID))
	if callbackData == "about" {
		err = ctx.Edit(globals.AboutTeamAnswer, &telebot.SendOptions{ParseMode: "markdown", ReplyMarkup: globals.AboutKeyboard})
	} else if callbackData == "about_bot" {
		err = ctx.Edit(globals.StartAnswer, &telebot.SendOptions{ParseMode: "markdown", ReplyMarkup: globals.StartKeyboard})
	}
	panelOwner := database.Get("group:" + stringChatID + ":panel:" + stringMessageID + ":owner")
	if stringSenderID == panelOwner && database.SIsMember("group:"+stringChatID+":admins", panelOwner) {
		if callbackData == "back_help" {
			err = ctx.Edit(globals.BackToHelpAnswer, &telebot.SendOptions{ParseMode: "markdown", ReplyMarkup: globals.HelpKeyboard})
		} else if strings.HasSuffix(callbackData, "_help") {
			splittedVariableName := strings.Split(callbackData, "_")
			for k, v := range splittedVariableName {
				splittedVariableName[k] = cases.Title(language.Und, cases.NoLower).String(v)
			}
			variableName := strings.Join(splittedVariableName, "")
			err = ctx.Edit(globals.HelpTexts[variableName], &telebot.SendOptions{ParseMode: "markdown", ReplyMarkup: globals.BackToHelpKeyboard})
		}
	}
	functions.HandleError(err)
	return nil
}
