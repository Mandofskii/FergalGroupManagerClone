package functions

import (
	"fmt"
	"os"
	"strconv"
	"unsafe"

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

func CastStr(v *string) string {
	return fmt.Sprint(uintptr(unsafe.Pointer(v)))
}

func UncastStr(s string) string {
	p, _ := strconv.ParseInt(s, 10, 64)
	return *((*string)(unsafe.Pointer(uintptr(p))))
}
