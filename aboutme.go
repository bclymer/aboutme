package main

import (
	"aboutme/aboutme"
	"html/template"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
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

	http.HandleFunc("/", handler)
	http.HandleFunc("/stack/timeline", aboutme.StackTimeline)
	http.HandleFunc("/stack/me", aboutme.StackUser)
	http.HandleFunc("/github/events", aboutme.GithubEvents)
	http.HandleFunc("/github/me", aboutme.GithubUser)
	http.HandleFunc("/twitter/timeline", aboutme.TwitterTimeline)
	http.HandleFunc("/facebook/feed", aboutme.FacebookFeed)
	http.HandleFunc("/unsupported", aboutme.GithubUnsupported)
	http.HandleFunc("/js/config.js", config)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	log.Println("aboutme is running...")
	http.ListenAndServe(":42126", nil)
}
