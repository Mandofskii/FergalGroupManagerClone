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
			switch lowerText := strings.ToLower(textMessage); {
			case lowerText == "راهنما":
				// database.SAdd("group:"+stringChatID+":panels", stringMessageID)
				sendedMessage, err := ctx.Bot().Send(chat, globals.HelpAnswer, &telebot.SendOptions{ReplyMarkup: globals.HelpKeyboard})
				database.Set("group:"+stringChatID+":panel:"+strconv.Itoa(sendedMessage.ID)+":owner", stringSenderID)
				functions.HandleError(err)
			case strings.Contains(lowerText, "vip"), strings.Contains(lowerText, "ویژه"):
				if message.ReplyTo != nil {
					firstName := message.ReplyTo.Sender.FirstName
					userID := message.ReplyTo.Sender.ID
					_, _ = firstName, userID
				} else {
					base := strings.Split(textMessage, " ")
					for _, v := range message.Entities {
						if v.Offset == len(base[0])+1 {
							if !strings.HasPrefix(base[1], "@") {
								firstName := v.User.FirstName
								userID := v.User.ID
								_, _ = firstName, userID
								fmt.Println(userID)
							} else {
								firstName := ""
								userID := int64(0)
								uuidRandom := uuid.NewString()
								ctx.Bot().Send(&telebot.User{ID: 5187419061}, "/getid "+base[1]+"\n"+stringChatID+"\n"+uuidRandom)
								ctx.Send("لطفا تا دریافت نتیجه کمی صبر کنید !")
								for {
									if database.Get("group:"+stringChatID+":hash:"+uuidRandom) != "" {
										userBase := database.Get("group:" + stringChatID + ":hash:" + uuidRandom)
										firstName = strings.Split(userBase, "|")[0]
										n, _ := strconv.ParseInt(strings.Split(userBase, "|")[1], 10, 64)
										userID = n
										break
									}
								}
								fmt.Println(firstName, userID)
							}
						}
					}

				}
			}

		}
	}
	functions.HandleError(err)
	return nil
}
