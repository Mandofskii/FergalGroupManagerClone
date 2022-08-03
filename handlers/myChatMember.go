package handlers

import (
	"FergalManagerClone/database"
	"FergalManagerClone/functions"
	"FergalManagerClone/globals"
	"fmt"
	"strconv"

	"gopkg.in/telebot.v3"
)

func NewMyChatMemberHandler(ctx telebot.Context) error {
	var err error
	chatID := ctx.Chat().ID
	stringChatID := strconv.Itoa(int(chatID))
	baseGroupKey := fmt.Sprintf("group:%s:", stringChatID)
	if ctx.Bot().Me.Username == ctx.ChatMember().NewChatMember.User.Username {
		if ctx.ChatMember().NewChatMember.Role == "member" {
			if ctx.ChatMember().Chat.Type == "supergroup" {
				// Here configuring group
				admins, err := ctx.Bot().AdminsOf(ctx.Chat())
				functions.HandleError(err)
				for _, v := range admins {
					userID := strconv.Itoa(int(v.User.ID))
					if v.Role == "creator" || v.Role == "owner" {
						database.Set(baseGroupKey+"owner", userID)
						database.SAdd(baseGroupKey+"owners", userID)
					}
					database.SAdd(baseGroupKey+"admins", userID)
				}
				database.Set(baseGroupKey+"rudeMode", "0")
				database.SAdd("installedGroups", stringChatID)
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
