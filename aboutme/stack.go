package aboutme

import (
	"fmt"
	"net/http"
)

const (
	stackId = "650288"
)

func StackTimeline(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, Get("http://api.stackexchange.com/users/"+stackId+"/timeline?site=stackoverflow.com"))
}

func StackUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, Get("http://api.stackexchange.com/users/"+stackId+"?site=stackoverflow.com"))
}
