package aboutme

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Get(url string) string {
	responseChannel := make(chan (RedisResponse))
	redisRequest := RedisRequest{
		Get:             true,
		Key:             url,
		ResponseChannel: responseChannel,
	}
	RedisRequestChannel <- redisRequest
	redisResponse := <-redisRequest.ResponseChannel
	if redisResponse.Err != nil {
		log.Println("Cache miss -", url, redisResponse.Err)
		response, err := http.Get(url)
		if err != nil {
			log.Printf("%s", err)
			return ""
		} else {
			defer response.Body.Close()
			contentBytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Printf("%s", err)
				return ""
			}
			response := string(contentBytes)
			redisRequest = RedisRequest{
				Get:   false,
				Key:   url,
				Value: response,
			}
			RedisRequestChannel <- redisRequest
			return response
		}
	}
	log.Println("Cache Hit -", url)
	return redisResponse.Value
}
