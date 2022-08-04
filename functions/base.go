package functions

import (
	"fmt"
	"strconv"

	"gopkg.in/telebot.v3"
)

func HandleError(err error) {
	if err != nil && err.Error() != "redis: nil" {
		fmt.Println(err.Error())
		// os.Exit(0)
	}
}

func ReturnBot(bot *telebot.Bot, err error) *telebot.Bot {
	HandleError(err)
	return bot
}

func Int64ToString(userID int64) string {
	return strconv.Itoa(int(userID))
}

func StringToInt64(userID string) int64 {
	result, err := strconv.Atoi(userID)
	HandleError(err)
	return int64(result)
}

func CreateMarkdownMention(userID int64, name string) string {
	return "[" + name + "](tg://user?id=" + Int64ToString(userID) + ")"
}
