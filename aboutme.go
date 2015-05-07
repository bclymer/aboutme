package main

import (
	"github.com/bclymer/aboutme"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func editInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "editInfo.html")
	} else {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprint(w, "Couldn't read post body")
			log.Println("Couldn't read post body", err)
			return
		}
		auth := r.FormValue("auth")
		correctAuth, err := ioutil.ReadFile("auth")
		if err != nil {
			fmt.Fprint(w, "Couldn't auth file")
			log.Println("Couldn't read auth file", err)
			return
		}
		if auth != string(correctAuth) {
			fmt.Fprint(w, "My security is bad and I should feel bad, but you still failed :P")
			log.Println("My security is bad and I should feel bad, but you still failed :P")
			return
		}
		log.Println("POST body", string(body))
		ioutil.WriteFile("me_events.json", body, 7777)
		aboutme.RedisDelete("me_events")
		fmt.Fprint(w, "Updated successfully and cache cleared")
	}
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
	http.HandleFunc("/edit/info", editInfo)
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
