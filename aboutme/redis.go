package aboutme

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"menteslibres.net/gosexy/redis"
)

const (
	cacheTime = 10 * 60 // 10 minutes
)

var (
	client              *redis.Client
	redisPrefix         = "aboutme_"
	redisRequestChannel = make(chan (RedisRequest))
)

type RedisRequest struct {
	Key             string
	CacheFunction   RedisCacheFunction
	ResponseChannel chan string
}

type RedisAuth struct {
	Password string `json:"password"`
}

func ConnectRedis() {
	var redisAuth RedisAuth
	content, err := ioutil.ReadFile("aboutme/redisAuth.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &redisAuth)
	if err != nil {
		panic(err)
	}
	client = redis.New()
	err = client.Connect("bclymer.com", 6379)
	if err != nil {
		panic(err)
	}
	_, err = client.Auth(redisAuth.Password)
	if err != nil {
		panic(err)
	}
	go func() {
		defer client.Quit()
		for {
			request := <-redisRequestChannel
			value := redisCache(request.Key, request.CacheFunction)
			request.ResponseChannel <- value
		}
	}()
}

type RedisCacheFunction func() string

func RedisCache(key string, redisCacheFunction RedisCacheFunction) string {
	responseChannel := make(chan (string))
	redisRequest := RedisRequest{
		Key:             key,
		CacheFunction:   redisCacheFunction,
		ResponseChannel: responseChannel,
	}
	redisRequestChannel <- redisRequest
	redisResponse := <-redisRequest.ResponseChannel
	return redisResponse
}

func redisCache(key string, redisCacheFunction RedisCacheFunction) string {
	value, err := redisGet(key)
	if err != nil && err == redis.ErrNilReply {
		log.Println("Cache miss -", key)
		value = redisCacheFunction()
		redisPut(key, value)
	} else {
		log.Println("Cache hit -", key)
	}
	return value
}

func redisPut(key, value string) {
	// check for client disconnect.
	fullKey := redisPrefix + key
	log.Println("SET -", fullKey)
	client.Set(fullKey, value)
	client.Expire(fullKey, cacheTime)
}

func redisGet(key string) (string, error) {
	// check for client disconnect
	fullKey := redisPrefix + key
	log.Println("GET -", fullKey)
	return client.Get(fullKey)
}
