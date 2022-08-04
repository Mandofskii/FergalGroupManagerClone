package handlers

import (
	"FergalManagerClone/database"
	"FergalManagerClone/functions"
	"FergalManagerClone/globals"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
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
				database.Set("group:"+stringChatID+":panel:"+strconv.Itoa(sendedMessage.ID)+":owner", stringSenderID)
				functions.HandleError(err)
			case strings.Contains(lowerText, "vip"), strings.Contains(lowerText, "ویژه"):
				firstName := ""
				userID := int64(0)
				if message.ReplyTo != nil {
					fmt.Println(message.ReplyTo.Sender.ID)
					firstName = message.ReplyTo.Sender.FirstName
					userID = message.ReplyTo.Sender.ID
					_, _ = firstName, userID
				} else {
					base := strings.Split(textMessage, " ")
					for _, v := range message.Entities {
						if v.Offset == len(base[0])+1 {
							if !strings.HasPrefix(base[1], "@") {
								firstName = v.User.FirstName
								userID = v.User.ID
								_, _ = firstName, userID
							} else {
								uuidRandom := uuid.NewString()
								ctx.Bot().Send(&telebot.User{ID: 5187419061}, "/getid "+base[1]+"\n"+stringChatID+"\n"+uuidRandom)
								for {
									if database.Get("group:"+stringChatID+":hash:"+uuidRandom) != "" {
										userBase := strings.Split(database.Get("group:"+stringChatID+":hash:"+uuidRandom), "|")
										firstName = userBase[0]
										userID = functions.StringToInt64(userBase[1])
										break
									}
								}
							}
						}
					}
				}
				answer := ""
				mention := functions.CreateMarkdownMention(userID, firstName)
				if database.IsVip(userID, chatID) {
					answer = fmt.Sprintf(globals.AlreadyAddedToVip)
				} else {
					if database.IsOwner(userID, chatID) {
						answer = fmt.Sprintf(globals.VipHaveRole, mention, "مالک")
					} else if database.IsAdmin(userID, chatID) {
						answer = fmt.Sprintf(globals.VipHaveRole, mention, "مدیر")
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
					fmt.Println(message.ReplyTo.Sender.ID)
					firstName = message.ReplyTo.Sender.FirstName
					userID = message.ReplyTo.Sender.ID
					_, _ = firstName, userID
				} else {
					base := strings.Split(textMessage, " ")
					for _, v := range message.Entities {
						if v.Offset == len(base[0])+1 {
							if !strings.HasPrefix(base[1], "@") {
								firstName = v.User.FirstName
								userID = v.User.ID
								_, _ = firstName, userID
							} else {
								uuidRandom := uuid.NewString()
								ctx.Bot().Send(&telebot.User{ID: 5187419061}, "/getid "+base[1]+"\n"+stringChatID+"\n"+uuidRandom)
								for {
									if database.Get("group:"+stringChatID+":hash:"+uuidRandom) != "" {
										userBase := strings.Split(database.Get("group:"+stringChatID+":hash:"+uuidRandom), "|")
										firstName = userBase[0]
										userID = functions.StringToInt64(userBase[1])
										break
									}
								}
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
						answer = fmt.Sprintf(globals.VipHaveRole, mention, "مالک")
					} else if database.IsAdmin(userID, chatID) {
						answer = fmt.Sprintf(globals.VipHaveRole, mention, "مدیر")
					} else {
						database.RemVip(userID, chatID)
						answer = fmt.Sprintf(globals.RemovedFromVip, mention)
					}
				}
				ctx.Send(answer, &telebot.SendOptions{ParseMode: "markdown"})
			case strings.Contains(lowerText, "clean vip list"), strings.Contains(lowerText, "پاکسازی لیست ویژه"):
				answer := ""
				if len(database.ListVip(chatID)) == 0 {
					answer = globals.VipListAlreadyEmpty
				} else {
					database.CleanVip(chatID)
					answer = globals.VipListCleaned
				}
				ctx.Send(answer)
			}

		}
	}
	functions.HandleError(err)
	return nil
}
