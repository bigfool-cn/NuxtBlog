package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"nuxt-blog-api/configs"
	"strconv"
	"time"
)

var RedisClient *redis.Client

func init()  {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     configs.Configs.Cache.Host + ":" +strconv.Itoa(configs.Configs.Cache.Port),
		Password: configs.Configs.Cache.Pwd, // no password set
		DB:       configs.Configs.Cache.Select,  // use default DB
	})

	pong, err := RedisClient.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
		panic(err)
	}
}

type Redis struct {
}

func (r Redis) Set(key string, val string, exp time.Duration) bool  {
	err := RedisClient.Set(key, val, exp).Err()
	if err != nil {
		return false
	}
	return true
}

func (r Redis) Get(key string) string  {
	val, err := RedisClient.Get(key).Result()
	if err != nil {
		return ""
	}
	return val
}

func (r Redis) Del(key string) bool  {
	err := RedisClient.Del(key).Err()
	if err != nil {
		return false
	}
	return true
}

var RedisCache = new(Redis)
