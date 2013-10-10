package aboutme

//package main

import (
	"bclymer/aboutme/aboutme"
	"log"
	"net/http"
)

const (
	urlPrefix    = "/me"
	folderPrefix = "aboutme/"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "aboutme/index.html")
}

func main() {
	aboutme.ConnectRedis()

	http.HandleFunc(urlPrefix+"/", handler)
	http.HandleFunc(urlPrefix+"/stack/timeline", aboutme.StackTimeline)
	http.HandleFunc(urlPrefix+"/stack/me", aboutme.StackUser)
	http.HandleFunc(urlPrefix+"/github/events", aboutme.GithubEvents)
	http.HandleFunc(urlPrefix+"/github/me", aboutme.GithubUser)
	http.HandleFunc(urlPrefix+"/unsupported", aboutme.GithubUnsupported)
	http.Handle(urlPrefix+"/static/", http.StripPrefix(urlPrefix+"/static", http.FileServer(http.Dir(folderPrefix+"static"))))
	log.Println("aboutme is running...")
}

func StartServer() {
	main()
}
