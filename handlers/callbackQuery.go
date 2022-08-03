package handlers

import (
	"FergalManagerClone/functions"
	"FergalManagerClone/globals"

	"gopkg.in/telebot.v3"
)

func HandleNewCallbackQuery(ctx telebot.Context) error {
	var err error
	callbackQuery := ctx.Callback()
	callbackData := callbackQuery.Data
	if callbackData == "about" {
		err = ctx.Edit(globals.AboutTeamAnswer, &telebot.SendOptions{ParseMode: "markdown", ReplyMarkup: globals.AboutKeyboard})
	} else if callbackData == "about_bot" {
		err = ctx.Edit(globals.StartAnswer, &telebot.SendOptions{ParseMode: "markdown", ReplyMarkup: globals.StartKeyboard})
	}
	functions.HandleError(err)
	return nil
}
