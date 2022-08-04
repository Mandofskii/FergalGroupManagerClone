package database

import (
	"FergalManagerClone/functions"
	"strconv"
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
	Set("group:"+stringChatID+":panel:"+stringMessageID+":owner", stringOwnerID)
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
	SAdd("installedGroups", functions.Int64ToString(chatID))
}

func RemoveGroup(groupChatID string) {
	result, err := redisDatabase.Keys("group:" + groupChatID + ":*").Result()
	functions.HandleError(err)
	for _, v := range result {
		Rem(v)
	}
}
