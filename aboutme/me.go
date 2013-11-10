package aboutme

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	events string
)

func MeInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	response := RedisCache("me_events", func() string {
		content, _ := ioutil.ReadFile("me_events.json")
		return string(content)
	})
	fmt.Fprint(w, response)
}
