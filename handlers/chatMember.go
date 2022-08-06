package handlers

import (
	"FergalManagerClone/database"
	"FergalManagerClone/functions"
	"FergalManagerClone/globals"
	"fmt"

	"gopkg.in/telebot.v3"
)

func NewChatMemberHandler(ctx telebot.Context) error {
	oldRole := ctx.ChatMember().OldChatMember.Role
	newRole := ctx.ChatMember().NewChatMember.Role
	userID := ctx.ChatMember().NewChatMember.User.ID
	chatID := ctx.Chat().ID
	if database.IsAutoConfigure(chatID) {
		if oldRole == "administrator" && (newRole == "member" || newRole == "left") {
			database.RemAdmin(userID, chatID)
		} else if oldRole == "member" || oldRole == "left" {
			if newRole == "creator" {
				database.AddOwner(userID, chatID)
				database.AddAdmin(userID, chatID)
			}
			if newRole == "administrator" {
				database.AddAdmin(userID, chatID)
			}
		}
	}
	return nil
}

func NewMyChatMemberHandler(ctx telebot.Context) error {
	var err error
	chatID := ctx.Chat().ID
	stringChatID := functions.Int64ToString(chatID)
	baseGroupKey := fmt.Sprintf("group:%s:", stringChatID)
	if ctx.Bot().Me.Username == ctx.ChatMember().NewChatMember.User.Username {
		if ctx.ChatMember().NewChatMember.Role == "member" {
			if ctx.ChatMember().Chat.Type == "supergroup" {
				admins, err := ctx.Bot().AdminsOf(ctx.Chat())
				functions.HandleError(err)
				for _, v := range admins {
					userID := v.User.ID
					if v.Role == "creator" || v.Role == "owner" {
						database.Set(baseGroupKey+"owner", functions.Int64ToString(userID), 0)
						database.AddOwner(userID, chatID)
					}
					database.AddAdmin(userID, chatID)
				}

				database.InstallGroup(chatID)
				v := &telebot.Video{File: telebot.FromDisk("assets/installed.mp4")}
				v.Caption = globals.InstalledAnswer
				functions.HandleError(ctx.SendAlbum(telebot.Album{v}))
			} else {
				err = ctx.Send(globals.UpgradeToSuperGroup, &telebot.SendOptions{ParseMode: "markdown"})
			}
			functions.HandleError(err)
		} else if ctx.ChatMember().NewChatMember.Role == "left" {
			database.RemoveGroup(stringChatID)
		}
	}
	return nil
}
