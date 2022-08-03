package functions

import (
	"fmt"
	"os"

	"gopkg.in/telebot.v3"
)

func HandleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func ReturnBot(bot *telebot.Bot, err error) *telebot.Bot {
	HandleError(err)
	return bot
}
