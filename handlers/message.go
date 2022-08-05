package handlers

import (
	"FergalManagerClone/database"
	"FergalManagerClone/functions"
	"FergalManagerClone/globals"
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/telebot.v3"
)

func HandleNewMessages(ctx telebot.Context) error {
	var err error
	message := ctx.Message()
	textMessage := message.Text
	chat := ctx.Chat()
	chatType := chat.Type
	chatID := chat.ID
	stringChatID := functions.Int64ToString(chatID)
	sender := message.Sender
	senderID := sender.ID
	stringSenderID := functions.Int64ToString(senderID)
	if chatType == "private" {
		if textMessage == "/start" {
			err = ctx.Send(globals.StartAnswer, &telebot.SendOptions{ParseMode: "markdown", ReplyMarkup: globals.StartKeyboard})
		}
	} else {
		if database.IsInstalled(chatID) && database.IsAdmin(senderID, chatID) {
			switch lowerText := strings.ToLower(textMessage); {
			case lowerText == "راهنما":
				// database.SAdd("group:"+stringChatID+":panels", stringMessageID)
				sendedMessage, err := ctx.Bot().Send(chat, globals.HelpAnswer, &telebot.SendOptions{ReplyMarkup: globals.HelpKeyboard})
				database.Set("group:"+stringChatID+":panel:"+strconv.Itoa(sendedMessage.ID)+":owner", stringSenderID, 0)
				functions.HandleError(err)
			case strings.Contains(lowerText, "clean vip list"), strings.Contains(lowerText, "پاکسازی لیست ویژه"):
				answer := ""
				if len(database.ListVip(chatID)) == 0 {
					answer = globals.VipListAlreadyEmpty
				} else {
					database.CleanVip(chatID)
					answer = globals.VipListCleaned
				}
				ctx.Send(answer)
			case strings.Contains(lowerText, "vip"), strings.Contains(lowerText, "ویژه"):
				firstName := ""
				userID := int64(0)
				if message.ReplyTo != nil {
					firstName = message.ReplyTo.Sender.FirstName
					userID = message.ReplyTo.Sender.ID
					_, _ = firstName, userID
				} else {
					base := strings.Split(textMessage, " ")
					username := base[len(base)-1]
					for _, v := range message.Entities {
						if v.Offset == len(textMessage)-len(username) {
							if !strings.HasPrefix(username, "@") {
								firstName = v.User.FirstName
								userID = v.User.ID
								_, _ = firstName, userID
							} else {
								firstName, userID = database.GetUserIDByUsername(ctx.Bot(), username, chatID)
							}
						}
					}
				}
				answer := ""
				mention := functions.CreateMarkdownMention(userID, firstName)
				if database.IsVip(userID, chatID) {
					answer = fmt.Sprintf(globals.AlreadyAddedToVip, mention)
				} else {
					if database.IsOwner(userID, chatID) {
						answer = fmt.Sprintf(globals.HaveRole, mention, "مالک")
					} else if database.IsAdmin(userID, chatID) {
						answer = fmt.Sprintf(globals.HaveRole, mention, "مدیر")
					} else {
						database.AddVip(userID, chatID)
						answer = fmt.Sprintf(globals.AddedToVip, mention)
					}
				}
				ctx.Send(answer, &telebot.SendOptions{ParseMode: "markdown"})
			case strings.Contains(lowerText, "unvip"), strings.Contains(lowerText, "حذف ویژه"):
				firstName := ""
				userID := int64(0)
				if message.ReplyTo != nil {
					firstName = message.ReplyTo.Sender.FirstName
					userID = message.ReplyTo.Sender.ID
					_, _ = firstName, userID
				} else {
					base := strings.Split(textMessage, " ")
					username := base[len(base)-1]
					for _, v := range message.Entities {
						if v.Offset == len(textMessage)-len(username) {
							if !strings.HasPrefix(username, "@") {
								firstName = v.User.FirstName
								userID = v.User.ID
								_, _ = firstName, userID
							} else {
								firstName, userID = database.GetUserIDByUsername(ctx.Bot(), username, chatID)
							}
						}
					}
				}
				answer := ""
				mention := functions.CreateMarkdownMention(userID, firstName)
				if !database.IsVip(userID, chatID) {
					answer = fmt.Sprintf(globals.AlreadyRemovedFromVip)
				} else {
					if database.IsOwner(userID, chatID) {
						answer = fmt.Sprintf(globals.HaveRole, mention, "مالک")
					} else if database.IsAdmin(userID, chatID) {
						answer = fmt.Sprintf(globals.HaveRole, mention, "مدیر")
					} else {
						database.RemVip(userID, chatID)
						answer = fmt.Sprintf(globals.RemovedFromVip, mention)
					}
				}
				ctx.Send(answer, &telebot.SendOptions{ParseMode: "markdown"})
			}
			chatMember := functions.IsBotAdmin(ctx.Bot(), chat)
			if chatMember.Role == "administrator" && chatMember.CanRestrictMembers {
				switch lowerText := strings.ToLower(textMessage); {
				case strings.Contains(lowerText, "clean mute list"), strings.Contains(lowerText, "پاکسازی لیست سکوت"):
					answer := ""
					if len(database.ListMute(chatID)) == 0 {
						answer = globals.MuteListAlreadyEmpty
					} else {

						for _, muted := range database.ListMute(chatID) {
							database.UnmuteUser(ctx.Bot(), chat, functions.StringToInt64(muted))
						}
						database.CleanMute(chatID)
						answer = globals.MuteListCleaned
					}
					ctx.Send(answer)
				case strings.Contains(lowerText, "unmute"), strings.Contains(lowerText, "حذف سکوت"):
					firstName := ""
					userID := int64(0)
					if message.ReplyTo != nil {
						firstName = message.ReplyTo.Sender.FirstName
						userID = message.ReplyTo.Sender.ID
						_, _ = firstName, userID
					} else {
						base := strings.Split(textMessage, " ")
						username := base[len(base)-1]
						for _, v := range message.Entities {
							if v.Offset == len([]rune(textMessage))-len([]rune(username)) {
								if !strings.HasPrefix(username, "@") {
									firstName = v.User.FirstName
									userID = v.User.ID
									_, _ = firstName, userID
								} else {
									firstName, userID = database.GetUserIDByUsername(ctx.Bot(), username, chatID)
								}
							}
						}
					}
					answer := ""
					mention := functions.CreateMarkdownMention(userID, firstName)
					if database.IsMute(userID, chatID) {
						if database.IsOwner(userID, chatID) {
							answer = fmt.Sprintf(globals.HaveRole, mention, "مالک")
						} else if database.IsAdmin(userID, chatID) || functions.IsGAdmin(ctx.Bot(), chat, functions.Int64ToString(userID)) {
							database.UnmuteUser(ctx.Bot(), chat, userID)
							answer = fmt.Sprintf(globals.HaveRole, mention, "مدیر")
						} else {
							database.UnmuteUser(ctx.Bot(), chat, userID)
							answer = fmt.Sprintf(globals.RemovedFromMuteList, mention)
						}
					} else {
						answer = fmt.Sprintf(globals.AlreadyRemovedFromMuteList, mention)
					}

					ctx.Send(answer, &telebot.SendOptions{ParseMode: "markdown"})
				case strings.Contains(lowerText, "mute"), strings.Contains(lowerText, "سکوت"):
					firstName := ""
					userID := int64(0)
					timeTTL, username, last := functions.GetMuteTime(lowerText)
					if message.ReplyTo != nil {
						firstName = message.ReplyTo.Sender.FirstName
						userID = message.ReplyTo.Sender.ID
					} else {
						for _, v := range message.Entities {
							if v.Offset == len([]rune(textMessage))-len([]rune(username)) {

								if !strings.HasPrefix(username, "@") {
									firstName = v.User.FirstName
									userID = v.User.ID
									_, _ = firstName, userID

								} else {
									firstName, userID = database.GetUserIDByUsername(ctx.Bot(), username, chatID)
								}
							}
						}
					}
					answer := ""
					mention := functions.CreateMarkdownMention(userID, firstName)
					if database.IsMute(userID, chatID) && last == "همیشه" {
						answer = fmt.Sprintf(globals.AlreadyAddedToMuteList, mention)
					} else {
						if database.IsOwner(userID, chatID) {
							answer = fmt.Sprintf(globals.HaveRole, mention, "مالک")
						} else if database.IsAdmin(userID, chatID) || functions.IsGAdmin(ctx.Bot(), chat, functions.Int64ToString(userID)) {
							answer = fmt.Sprintf(globals.HaveRole, mention, "مدیر")
						} else {
							database.MuteUser(ctx.Bot(), chat, userID, timeTTL)
							if last == "همیشه" {
								last = "برای همیشه سکوت شد 🔇"
							} else {
								last = fmt.Sprintf(globals.MutedForTime, last)
							}
							answer = fmt.Sprintf(globals.AddedToMuteList, mention, last)
						}
					}
					ctx.Send(answer, &telebot.SendOptions{ParseMode: "markdown"})
				}
			}
		}
	}
	functions.HandleError(err)
	return nil
}
