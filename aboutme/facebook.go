package aboutme

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	FacebookId = "crdnilfan"
)

var (
	accessToken string
)

type FacebookAuth struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func FacebookFeed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, Get("https://graph.facebook.com/crdnilfan/feed?access_token="+accessToken))
}

func Setup(notifier chan bool) {
	log.Println("Setting up facebook")
	var facebookAuth FacebookAuth
	content, err := ioutil.ReadFile("aboutme/facebookAuth.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &facebookAuth)
	if err != nil {
		panic(err)
	}
	go func() {
		response := Get(fmt.Sprintf("https://graph.facebook.com/oauth/access_token?client_id=%s&client_secret=%s&grant_type=client_credentials", facebookAuth.ClientId, facebookAuth.ClientSecret))
		accessKey := strings.Split(response, "=")
		if len(accessKey) >= 2 {
			if accessKey[0] == "access_token" {
				accessToken = accessKey[1]
			} else {
				log.Fatalln("Can't start facebook!!")
			}
		} else {
			log.Fatalln("Can't start facebook!!")
		}
		notifier <- true
	}()
}
