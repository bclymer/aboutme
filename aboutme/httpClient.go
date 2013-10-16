package aboutme

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Get(url string) string {
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
		return response
	}
}
