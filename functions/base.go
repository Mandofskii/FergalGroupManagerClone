package functions

import (
	"fmt"

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
