package aboutme

//package main

import (
	"bclymer/aboutme/aboutme"
	"html/template"
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

func config(w http.ResponseWriter, r *http.Request) {
	accountIds := AccountIds{
		Stack:    aboutme.StackId,
		Github:   aboutme.GithubId,
		Twitter:  aboutme.TwitterId,
		Facebook: aboutme.FacebookId,
	}
	w.Header().Set("Content-Type", "text/javascript")
	t := template.New("Person template")
	t, _ = t.Parse(configJsConst)
	t.Execute(w, accountIds)
}

type AccountIds struct {
	Stack    string
	Github   string
	Twitter  string
	Facebook string
}

const configJsConst = `(function () {
	AboutMe.config = {
		stack: "{{.Stack}}",
		github: "{{.Github}}",
		twitter: "{{.Twitter}}",
		facebook: "{{.Facebook}}",
	};
})();
`

func main() {
	aboutme.ConnectRedis()

	notifier := make(chan bool)
	aboutme.Setup(notifier)
	<-notifier

	http.HandleFunc(urlPrefix+"/", handler)
	http.HandleFunc(urlPrefix+"/stack/timeline", aboutme.StackTimeline)
	http.HandleFunc(urlPrefix+"/stack/me", aboutme.StackUser)
	http.HandleFunc(urlPrefix+"/github/events", aboutme.GithubEvents)
	http.HandleFunc(urlPrefix+"/github/me", aboutme.GithubUser)
	http.HandleFunc(urlPrefix+"/twitter/timeline", aboutme.TwitterTimeline)
	http.HandleFunc(urlPrefix+"/facebook/feed", aboutme.FacebookFeed)
	http.HandleFunc(urlPrefix+"/unsupported", aboutme.GithubUnsupported)
	http.HandleFunc(urlPrefix+"/js/config.js", config)
	http.Handle(urlPrefix+"/static/", http.StripPrefix(urlPrefix+"/static", http.FileServer(http.Dir(folderPrefix+"static"))))
	log.Println("aboutme is running...")
}

func StartServer() {
	main()
}
