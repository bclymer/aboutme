package aboutme

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	GithubId = "bclymer"
)

func GithubEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, Get("https://api.github.com/users/"+GithubId+"/events/public"))
}

func GithubUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, Get("https://api.github.com/users/"+GithubId))
}

func GithubUnsupported(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// let them know POST only
		return
	}
	file, err := os.OpenFile("aboutme/unsupported.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	file.Write(body)
	file.Write([]byte("\r\n"))
	file.Close()
}
