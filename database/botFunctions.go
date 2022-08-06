package database

import (
	"FergalManagerClone/functions"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"gopkg.in/telebot.v3"
)

func AddAdmin(userID, chatID int64) {
	stringUserID, stringChatID := functions.Int64ToString(userID), functions.Int64ToString(chatID)
	SAdd("group:"+stringChatID+":admins", stringUserID)
}

func IsAdmin(userID, chatID int64) bool {
	stringUserID, stringChatID := functions.Int64ToString(userID), functions.Int64ToString(chatID)
	return SIsMember("group:"+stringChatID+":admins", stringUserID)
}

func RemAdmin(userID, chatID int64) {
	stringUserID, stringChatID := functions.Int64ToString(userID), functions.Int64ToString(chatID)
	SRem("group:"+stringChatID+":admins", stringUserID)
}

func AddOwner(userID, chatID int64) {
	stringUserID, stringChatID := functions.Int64ToString(userID), functions.Int64ToString(chatID)
	SAdd("group:"+stringChatID+":owners", stringUserID)
}

func IsOwner(userID, chatID int64) bool {
	stringUserID, stringChatID := functions.Int64ToString(userID), functions.Int64ToString(chatID)
	return SIsMember("group:"+stringChatID+":owners", stringUserID)
}

func RemOwner(userID, chatID int64) {
	stringUserID, stringChatID := functions.Int64ToString(userID), functions.Int64ToString(chatID)
	SRem("group:"+stringChatID+":owners", stringUserID)
}

func IsInstalled(chatID int64) bool {
	stringChatID := functions.Int64ToString(chatID)
	return SIsMember("installedGroups", stringChatID)

}

func OpenPanel(ownerID, chatID int64, messageID int) {
	stringOwnerID, stringChatID, stringMessageID := functions.Int64ToString(ownerID), functions.Int64ToString(chatID), strconv.Itoa(messageID)
	Set("group:"+stringChatID+":panel:"+stringMessageID+":owner", stringOwnerID, 0)
}

func AddVip(userID, chatID int64) {
	stringUserID, stringChatID := functions.Int64ToString(userID), functions.Int64ToString(chatID)
	SAdd("group:"+stringChatID+":vips", stringUserID)
}

func IsVip(userID, chatID int64) bool {
	stringUserID, stringChatID := functions.Int64ToString(userID), functions.Int64ToString(chatID)
	return SIsMember("group:"+stringChatID+":vips", stringUserID)
}

func ListVip(chatID int64) []string {
	stringChatID := functions.Int64ToString(chatID)
	return SMembers("group:" + stringChatID + ":vips")
}

func CleanVip(chatID int64) {
	stringChatID := functions.Int64ToString(chatID)
	Rem("group:" + stringChatID + ":vips")
}

func RemVip(userID, chatID int64) {
	stringUserID, stringChatID := functions.Int64ToString(userID), functions.Int64ToString(chatID)
	SRem("group:"+stringChatID+":vips", stringUserID)
}

func InstallGroup(chatID int64) {
	baseGroupKey := "group:" + functions.Int64ToString(chatID) + ":"
	Set(baseGroupKey+"rudeMode", "0", 0)
	Set(baseGroupKey+"autoConfigure", "1", 0)
	SAdd("installedGroups", functions.Int64ToString(chatID))
}

func IsAutoConfigure(chatID int64) bool {
	baseGroupKey := "group:" + functions.Int64ToString(chatID) + ":"
	return Get(baseGroupKey+"autoConfigure") == "1"
}

func RemoveGroup(groupChatID string) {
	result, err := redisDatabase.Keys("group:" + groupChatID + ":*").Result()
	functions.HandleError(err)
	for _, v := range result {
		Rem(v)
	}
}

func AddMute(userID, chatID int64, hash string, until int) {
	fmt.Printf("%#v", ListMute(chatID))
	stringUserID, stringChatID := functions.Int64ToString(userID), functions.Int64ToString(chatID)
	SAdd("group:"+stringChatID+":mutes", stringUserID+"|"+hash)
	if until == 1 {
		until = 0
	}
	Set("group:"+stringChatID+":muted:hash:"+hash, "nothing special here", until)
}

func IsMute(userID, chatID int64) bool {
	stringUserID := functions.Int64ToString(userID)
	for _, muted := range ListMute(chatID) {
		fmt.Println(muted)
		if strings.Split(muted, "|")[0] == stringUserID {
			return true
		}
	}
	return false
}

func ListMute(chatID int64) []string {
	correctedList := []string{}
	stringChatID := functions.Int64ToString(chatID)
	result := SMembers("group:" + stringChatID + ":mutes")
	fmt.Printf("%#v\n", result)
	for _, muted := range result {
		if Get("group:"+stringChatID+":muted:hash:"+strings.Split(muted, "|")[1]) != "" {
			correctedList = append(correctedList, strings.Split(muted, "|")[0])
		} else {
			SRem("group:"+stringChatID+":mutes", muted)
		}
	}
	return correctedList
}

func CleanMute(chatID int64) {
	stringChatID := functions.Int64ToString(chatID)
	Rem("group:" + stringChatID + ":mutes")
}

func RemMute(userID, chatID int64) {
	stringUserID, stringChatID := functions.Int64ToString(userID), functions.Int64ToString(chatID)
	for _, muted := range SMembers("group:" + stringChatID + ":mutes") {
		if strings.Split(muted, "|")[0] == stringUserID {
			SRem("group:"+stringChatID+":mutes", muted)
		}
	}
}

func MuteUser(bot *telebot.Bot, chat *telebot.Chat, userID, timeTTL int64) {
	chatMember, _ := bot.ChatMemberOf(chat, functions.Int64ToString(userID))
	chatMember.CanSendMessages = false
	if timeTTL != 1 {
		chatMember.RestrictedUntil = time.Now().Unix() + int64(timeTTL)
	}
	bot.Restrict(chat, chatMember)
	hash := uuid.NewString()
	AddMute(userID, chat.ID, hash, int(timeTTL))
}

func UnmuteUser(bot *telebot.Bot, chat *telebot.Chat, userID int64) {
	chatMember, _ := bot.ChatMemberOf(chat, functions.Int64ToString(userID))
	chatMember.CanSendMessages = true
	bot.Restrict(chat, chatMember)
	RemMute(userID, chat.ID)
}

func GetUserIDByUsername(bot *telebot.Bot, username string, chatID int64) (string, int64) {
	uuidRandom := uuid.NewString()
	stringChatID := functions.Int64ToString(chatID)
	firstName, userID := "", int64(0)
	bot.Send(&telebot.User{ID: 5187419061}, "/getid "+username+"\n"+stringChatID+"\n"+uuidRandom)
	for {
		if Get("group:"+stringChatID+":hash:"+uuidRandom) != "" {
			userBase := strings.Split(Get("group:"+stringChatID+":hash:"+uuidRandom), "|")
			firstName = userBase[0]
			userID = functions.StringToInt64(userBase[1])
			break
		}
	}
	return firstName, userID
}

func IsFilter(chatID int64, word string) bool {
	return SIsMember("group:"+functions.Int64ToString(chatID)+":filters", word)
}

func AddFilter(chatID int64, word string) {
	SAdd("group:"+functions.Int64ToString(chatID)+":filters", word)
}

func RemFilter(chatID int64, word string) {
	SRem("group:"+functions.Int64ToString(chatID)+":filters", word)
}

func ListFilter(chatID int64) []string {
	return SMembers("group:" + functions.Int64ToString(chatID) + ":filters")
}

func HasFilteredWord(chatID int64, text string) bool {
	for _, word := range ListFilter(chatID) {
		if strings.Contains(text, word) {
			return true
		}
	}
	return false
}

func CleanFilter(chatID int64) {
	Rem("group:" + functions.Int64ToString(chatID) + ":filters")
}

func AddBan(userID, chatID int64, hash string, until int) {
	fmt.Printf("%#v", ListMute(chatID))
	stringUserID, stringChatID := functions.Int64ToString(userID), functions.Int64ToString(chatID)
	SAdd("group:"+stringChatID+":bans", stringUserID+"|"+hash)
	if until == 1 {
		until = 0
	}
	Set("group:"+stringChatID+":banned:hash:"+hash, "nothing special here", until)
}

func IsBan(userID, chatID int64) bool {
	stringUserID := functions.Int64ToString(userID)
	for _, banned := range ListBan(chatID) {
		fmt.Println(banned)
		if strings.Split(banned, "|")[0] == stringUserID {
			return true
		}
	}
	return false
}

func ListBan(chatID int64) []string {
	correctedList := []string{}
	stringChatID := functions.Int64ToString(chatID)
	result := SMembers("group:" + stringChatID + ":bans")
	fmt.Printf("%#v\n", result)
	for _, banned := range result {
		if Get("group:"+stringChatID+":banned:hash:"+strings.Split(banned, "|")[1]) != "" {
			correctedList = append(correctedList, strings.Split(banned, "|")[0])
		} else {
			SRem("group:"+stringChatID+":bans", banned)
		}
	}
	return correctedList
}

func CleanBan(chatID int64) {
	stringChatID := functions.Int64ToString(chatID)
	Rem("group:" + stringChatID + ":bans")
}

func RemBan(userID, chatID int64) {
	stringUserID, stringChatID := functions.Int64ToString(userID), functions.Int64ToString(chatID)
	for _, banned := range SMembers("group:" + stringChatID + ":bans") {
		if strings.Split(banned, "|")[0] == stringUserID {
			SRem("group:"+stringChatID+":bans", banned)
		}
	}
}

func BanUser(bot *telebot.Bot, chat *telebot.Chat, userID, timeTTL int64) {
	chatMember, _ := bot.ChatMemberOf(chat, functions.Int64ToString(userID))
	if timeTTL != 1 {
		chatMember.RestrictedUntil = time.Now().Unix() + int64(timeTTL)
	}
	revokeMessages := Get("group:"+functions.Int64ToString(chat.ID)+":delMsgBan") == "1"
	bot.Ban(chat, chatMember, revokeMessages)
	hash := uuid.NewString()
	AddBan(userID, chat.ID, hash, int(timeTTL))
}

func UnbanUser(bot *telebot.Bot, chat *telebot.Chat, userID int64) {
	err := bot.Unban(chat, functions.Int64ToString(userID))
	functions.HandleError(err)
	RemBan(userID, chat.ID)
}
