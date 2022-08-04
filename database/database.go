package database

import (
	"FergalManagerClone/functions"
	"FergalManagerClone/globals"

	"gopkg.in/redis.v5"
)

var redisDatabase *redis.Client = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", DB: globals.DB})

func SAdd(key string, value string) int {
	result, err := redisDatabase.SAdd(key, value).Result()
	functions.HandleError(err)
	return int(result)
}

func Set(key string, value string) string {
	result, err := redisDatabase.Set(key, value, 0).Result()
	functions.HandleError(err)
	return result
}

func Rem(key string) int {
	result, err := redisDatabase.Del(key).Result()
	functions.HandleError(err)
	return int(result)
}

func RemoveGroup(groupChatID string) int {
	result, err := redisDatabase.Keys("group:" + groupChatID + ":*").Result()
	functions.HandleError(err)
	for _, v := range result {
		Rem(v)
	}
	return 1
}

func SIsMember(key string, value string) bool {
	result, err := redisDatabase.SIsMember(key, value).Result()
	functions.HandleError(err)
	return result
}

func Get(key string) string {
	result, err := redisDatabase.Get(key).Result()
	functions.HandleError(err)
	return result
}
