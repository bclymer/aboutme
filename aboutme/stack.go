package aboutme

import (
	"fmt"
	"net/http"
)

const (
	StackId = "650288"
)

func StackTimeline(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	url := "http://api.stackexchange.com/users/" + StackId + "/timeline?site=stackoverflow.com"
	response := RedisCache(url, func() string {
		return Get(url)
	})
	fmt.Fprint(w, response)
}

func StackUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	url := "http://api.stackexchange.com/users/" + StackId + "?site=stackoverflow.com"
	response := RedisCache(url, func() string {
		return Get(url)
	})
	fmt.Fprint(w, response)
}
