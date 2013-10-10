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
	RedisRequestChannel = make(chan (RedisRequest))
)

type RedisRequest struct {
	Get             bool
	Key             string
	Value           string
	ResponseChannel chan RedisResponse
}

type RedisResponse struct {
	Value string
	Err   error
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
			request := <-RedisRequestChannel
			if request.Get {
				value, err := redisGet(request.Key)
				response := RedisResponse{value, err}
				request.ResponseChannel <- response
			} else {
				redisPut(request.Key, request.Value)
			}
		}
	}()
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
