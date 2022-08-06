package handlers

import (
	"FergalManagerClone/database"
	"FergalManagerClone/functions"
	"FergalManagerClone/globals"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/telebot.v3"
)

func HandleNewMessages(ctx telebot.Context) error {
	var err error
	bot := ctx.Bot()

	message := ctx.Message()
	textMessage := message.Text

	chat := ctx.Chat()
	chatType := chat.Type
	chatID := chat.ID
	stringChatID := functions.Int64ToString(chatID)

	sender := message.Sender
	senderID := sender.ID
	stringSenderID := functions.Int64ToString(senderID)

	answer := ""
	if chatType == "private" {
		if textMessage == "/start" {
			err = ctx.Send(globals.StartAnswer, &telebot.SendOptions{ParseMode: "markdown", ReplyMarkup: globals.StartKeyboard})
		}
	} else {
		if database.IsInstalled(chatID) && database.IsAdmin(senderID, chatID) {
			switch lowerText := strings.ToLower(textMessage); {
			case lowerText == "راهنما":
				sendedMessage, err := bot.Send(chat, globals.HelpAnswer, &telebot.SendOptions{ReplyMarkup: globals.HelpKeyboard})
				database.Set("group:"+stringChatID+":panel:"+strconv.Itoa(sendedMessage.ID)+":owner", stringSenderID, 0)
				functions.HandleError(err)
			case lowerText == "clean vip list", lowerText == "پاکسازی لیست ویژه":
				if len(database.ListVip(chatID)) == 0 {
					answer = fmt.Sprintf(globals.ListAlreadyEmpty, "ویژه")
				} else {
					database.CleanVip(chatID)
					answer = fmt.Sprintf(globals.ListCleaned, "ویژه")
				}
				err = ctx.Send(answer)
			case lowerText == "clean filter list", lowerText == "پاکسازی لیست فیلتر":
				if len(database.ListFilter(chatID)) == 0 {
					answer = fmt.Sprintf(globals.ListAlreadyEmpty, "فیلتر")
				} else {
					database.CleanFilter(chatID)
					answer = fmt.Sprintf(globals.ListCleaned, "فیلتر")
				}
				err = ctx.Send(answer)
			case strings.Contains(lowerText, "unvip"), strings.Contains(lowerText, "حذف ویژه"):
				firstName := ""
				userID := int64(0)
				if message.ReplyTo != nil {
					firstName = message.ReplyTo.Sender.FirstName
					userID = message.ReplyTo.Sender.ID
				} else {
					base := strings.Split(textMessage, " ")
					username := base[len(base)-1]
					for _, v := range message.Entities {
						if v.Offset == len([]rune(textMessage))-len([]rune(username)) {
							if !strings.HasPrefix(username, "@") {
								firstName = v.User.FirstName
								userID = v.User.ID
							} else {
								firstName, userID = database.GetUserIDByUsername(bot, username, chatID)
							}
						}
					}
				}
				mention := functions.CreateMarkdownMention(userID, firstName)
				if !database.IsVip(userID, chatID) {
					answer = globals.AlreadyRemovedFromVip
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
				err = ctx.Send(answer, &telebot.SendOptions{ParseMode: "markdown"})
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
						if v.Offset == len([]rune(textMessage))-len([]rune(username)) {
							if !strings.HasPrefix(username, "@") {
								firstName = v.User.FirstName
								userID = v.User.ID
							} else {
								firstName, userID = database.GetUserIDByUsername(ctx.Bot(), username, chatID)
							}
						}
					}
				}
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
				err = ctx.Send(answer, &telebot.SendOptions{ParseMode: "markdown"})
			case strings.Contains(lowerText, "unfilter"), strings.Contains(lowerText, "حذف فیلتر"):
				muteRegexPattern := `(unfilter|حذف فیلتر) (.*)`
				r := regexp.MustCompile(muteRegexPattern)
				matches := r.FindStringSubmatch(lowerText)
				if matches[2] != "" {
					word := matches[2]
					if database.IsFilter(chatID, word) {
						database.RemFilter(chatID, word)
						answer = fmt.Sprintf(globals.RemovedFromFiltered, word)
					} else {
						answer = fmt.Sprintf(globals.AlreadyNotFiltered, word)
					}
					err = ctx.Send(answer, &telebot.SendOptions{ParseMode: "markdown"})
				}
			case strings.Contains(lowerText, "filter"), strings.Contains(lowerText, "فیلتر"):
				muteRegexPattern := `(filter|فیلتر) (.*)`
				r := regexp.MustCompile(muteRegexPattern)
				matches := r.FindStringSubmatch(lowerText)
				if matches[2] != "" {
					word := matches[2]
					if !database.IsFilter(chatID, word) {
						database.AddFilter(chatID, word)
						answer = fmt.Sprintf(globals.Filtered, word)
					} else {
						answer = fmt.Sprintf(globals.AlreadyFiltered, word)
					}
					err = ctx.Send(answer, &telebot.SendOptions{ParseMode: "markdown"})
				}
			}
			botChatMember := functions.IsBotAdmin(ctx.Bot(), chat)
			if botChatMember.Role == "administrator" && botChatMember.CanRestrictMembers {
				switch lowerText := strings.ToLower(textMessage); {
				case strings.Contains(lowerText, "clean mute list"), strings.Contains(lowerText, "پاکسازی لیست سکوت"):
					if len(database.ListMute(chatID)) == 0 {
						answer = fmt.Sprintf(globals.ListAlreadyEmpty, "سکوت")
					} else {
						for _, muted := range database.ListMute(chatID) {
							database.UnmuteUser(ctx.Bot(), chat, functions.StringToInt64(muted))
						}
						database.CleanMute(chatID)
						answer = fmt.Sprintf(globals.ListCleaned, "سکوت")
					}
					err = ctx.Send(answer)
				case strings.Contains(lowerText, "unmute"), strings.Contains(lowerText, "حذف سکوت"):
					firstName := ""
					userID := int64(0)
					if message.ReplyTo != nil {
						firstName = message.ReplyTo.Sender.FirstName
						userID = message.ReplyTo.Sender.ID
					} else {
						base := strings.Split(textMessage, " ")
						username := base[len(base)-1]
						for _, v := range message.Entities {
							if v.Offset == len([]rune(textMessage))-len([]rune(username)) {
								if !strings.HasPrefix(username, "@") {
									firstName = v.User.FirstName
									userID = v.User.ID
								} else {
									firstName, userID = database.GetUserIDByUsername(ctx.Bot(), username, chatID)
								}
							}
						}
					}
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
								} else {
									firstName, userID = database.GetUserIDByUsername(bot, username, chatID)
								}
							}
						}
					}
					mention := functions.CreateMarkdownMention(userID, firstName)
					if database.IsMute(userID, chatID) && last == "همیشه" {
						answer = fmt.Sprintf(globals.AlreadyAddedToMuteList, mention)
					} else {
						if database.IsOwner(userID, chatID) {
							answer = fmt.Sprintf(globals.HaveRole, mention, "مالک")
						} else if database.IsAdmin(userID, chatID) || functions.IsGAdmin(bot, chat, functions.Int64ToString(userID)) {
							answer = fmt.Sprintf(globals.HaveRole, mention, "مدیر")
						} else {
							database.MuteUser(bot, chat, userID, timeTTL)
							if last == "همیشه" {
								last = "برای همیشه سکوت شد 🔇"
							} else {
								last = fmt.Sprintf(globals.MutedForTime, last)
							}
							answer = fmt.Sprintf(globals.AddedToMuteList, mention, last)
						}
					}
					err = ctx.Send(answer, &telebot.SendOptions{ParseMode: "markdown"})
				}
			}
		}
		if !database.IsAdmin(senderID, chatID) && !functions.IsGAdmin(bot, chat, stringSenderID) && !database.IsVip(senderID, chatID) {
			if database.HasFilteredWord(chatID, strings.ToLower(textMessage)) {
				ctx.Delete()
			}
		}
	}
	functions.HandleError(err)
	return nil
}
