package aboutme

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Get(url string) string {
	cachedResponse, err := RedisGet(url)
	if err != nil {
		log.Println("Cache miss -", url, err)
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
			RedisPut(url, response)
			return response
		}
	}
	log.Println("Cache Hit -", url)
	return cachedResponse
}
