package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

type Doc struct {
	Id       int
	Keywords []string
}

var redisclient *redis.Client

func main() {
	docs := []*Doc{&Doc{Id: 1, Keywords: []string{"姜", "屎粑粑", "小宝宝"}}, &Doc{Id: 2, Keywords: []string{"麦麦", "薯条", "没有啦", "姜"}}}
	BuildInvertIndexv2(docs)
	fmt.Println(redisclient.LRange(context.TODO(), "姜", 0, -1).Val())
}

func BuildInvertIndexv1(docs []*Doc) map[string][]int {
	var InvertIndex map[string][]int = make(map[string][]int, 100)
	for _, doc := range docs {
		for _, keyword := range doc.Keywords {
			InvertIndex[keyword] = append(InvertIndex[keyword], doc.Id)
		}
	}
	return InvertIndex
}

func BuildInvertIndexv2(docs []*Doc) {
	InitRedisClient()
	for _, doc := range docs {
		for _, keyword := range doc.Keywords {
			redisclient.RPush(context.TODO(), keyword, doc.Id)
			redisclient.Expire(context.TODO(), keyword, time.Hour)
		}
	}
}

func InitRedisClient() {
	redisclient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	if err := redisclient.Ping(context.TODO()).Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
