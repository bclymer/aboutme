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
		Stack:   aboutme.StackId,
		Github:  aboutme.GithubId,
		Twitter: aboutme.TwitterId,
	}
	w.Header().Set("Content-Type", "text/javascript")
	t := template.New("Person template")
	t, _ = t.Parse(configJsConst)
	t.Execute(w, accountIds)
}

type AccountIds struct {
	Stack   string
	Github  string
	Twitter string
}

const configJsConst = `(function () {
	AboutMe.config = {
		stack: "{{.Stack}}",
		github: "{{.Github}}",
		twitter: "{{.Twitter}}",
	};
})();
`

func main() {
	aboutme.ConnectRedis()

	http.HandleFunc("/", handler)
	http.HandleFunc("/stack/timeline", aboutme.StackTimeline)
	http.HandleFunc("/stack/me", aboutme.StackUser)
	http.HandleFunc("/github/events", aboutme.GithubEvents)
	http.HandleFunc("/github/me", aboutme.GithubUser)
	http.HandleFunc("/twitter/timeline", aboutme.TwitterTimeline)
	http.HandleFunc("/me/info", aboutme.MeInfo)
	http.HandleFunc("/unsupported", aboutme.GithubUnsupported)
	http.HandleFunc("/js/config.js", config)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	log.Println("aboutme is running...")
	http.ListenAndServe(":42126", nil)
}
