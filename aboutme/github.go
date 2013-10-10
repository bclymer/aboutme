package aboutme

import (
	"fmt"
	"net/http"
)

const (
	githubId = "bclymer"
)

func GithubEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, Get("https://api.github.com/users/"+githubId+"/events/public"))
}

func GithubUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, Get("https://api.github.com/users/"+githubId))
}
