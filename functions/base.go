package functions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

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

func zip(a1, a2 []string) []string {
	r := make([]string, 2*len(a1))
	for i, e := range a1 {
		r[i*2] = e
		r[i*2+1] = a2[i]
	}
	return r
}

func GetMuteTime(text string) (int64, string, string) {
	timeTTL := 1
	muteRegexPattern := `(mute|سکوت)( (\d+) (m|h|d|روز|دقیقه|ساعت) ?)?(.*)`
	r := regexp.MustCompile(muteRegexPattern)
	matches := r.FindStringSubmatch(text)

	if matches[0] != matches[1] {
		if matches[3] != "" && matches[4] != "" {
			if matches[4] == "d" || matches[4] == "روز" {
				timeTTL, _ = strconv.Atoi(matches[3])
				timeTTL *= (24 * 60 * 60)
			} else if matches[4] == "h" || matches[4] == "ساعت" {
				timeTTL, _ = strconv.Atoi(matches[3])
				timeTTL *= (60 * 60)
			} else if matches[4] == "m" || matches[4] == "دقیقه" {
				timeTTL, _ = strconv.Atoi(matches[3])
				timeTTL *= 60
			}
		}
	}
	last := ""
	if matches[2] != "" {
		array1 := []string{"d", "h", "m"}
		array2 := []string{"روز", "ساعت", "دقیقه"}
		last = matches[3] + " " + strings.NewReplacer(zip(array1, array2)...).Replace(matches[4])
	} else {
		last = "همیشه"
	}
	return int64(timeTTL), matches[5], last
}

func GetBanTime(text string) (int64, string, string) {
	timeTTL := 1
	muteRegexPattern := `(ban|بن)( (\d+) (m|h|d|روز|دقیقه|ساعت) ?)?(.*)`
	r := regexp.MustCompile(muteRegexPattern)
	matches := r.FindStringSubmatch(text)

	if matches[0] != matches[1] {
		if matches[3] != "" && matches[4] != "" {
			if matches[4] == "d" || matches[4] == "روز" {
				timeTTL, _ = strconv.Atoi(matches[3])
				timeTTL *= (24 * 60 * 60)
			} else if matches[4] == "h" || matches[4] == "ساعت" {
				timeTTL, _ = strconv.Atoi(matches[3])
				timeTTL *= (60 * 60)
			} else if matches[4] == "m" || matches[4] == "دقیقه" {
				timeTTL, _ = strconv.Atoi(matches[3])
				timeTTL *= 60
			}
		}
	}
	last := ""
	if matches[2] != "" {
		array1 := []string{"d", "h", "m"}
		array2 := []string{"روز", "ساعت", "دقیقه"}
		last = matches[3] + " " + strings.NewReplacer(zip(array1, array2)...).Replace(matches[4])
	} else {
		last = "همیشه"
	}
	return int64(timeTTL), matches[5], last
}

func IsBotAdmin(bot *telebot.Bot, chat *telebot.Chat) *telebot.ChatMember {
	chatMember, err := bot.ChatMemberOf(chat, Int64ToString(bot.Me.ID))
	HandleError(err)
	return chatMember
}

func IsGAdmin(bot *telebot.Bot, chat *telebot.Chat, user string) bool {
	chatMember, err := bot.ChatMemberOf(chat, user)
	HandleError(err)
	return chatMember.Role == "administrator" || chatMember.Role == "creator"
}
