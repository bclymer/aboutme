package aboutme

//package main

import (
	"bclymer/aboutme/aboutme"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	urlPrefix    = "/me"
	folderPrefix = "aboutme/"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "aboutme/index.html")
}

func unsupported(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// let them know post only
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
	file.Close()
}

func main() {
	aboutme.ConnectRedis()

	http.HandleFunc(urlPrefix+"/", handler)
	http.HandleFunc(urlPrefix+"/stack/timeline", aboutme.StackTimeline)
	http.HandleFunc(urlPrefix+"/stack/me", aboutme.StackUser)
	http.HandleFunc(urlPrefix+"/github/events", aboutme.GithubEvents)
	http.HandleFunc(urlPrefix+"/github/me", aboutme.GithubUser)
	http.HandleFunc(urlPrefix+"/unsupported", unsupported)
	http.Handle(urlPrefix+"/static/", http.StripPrefix(urlPrefix+"/static", http.FileServer(http.Dir(folderPrefix+"static"))))
	log.Println("aboutme is running...")
}

func StartServer() {
	main()
}
