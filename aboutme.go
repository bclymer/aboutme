package aboutme

import (
	"bclymer/aboutme/aboutme"
	"fmt"
	"log"
	"net/http"
)

const (
	stackId = "650288"
)

var (
	curDir string
)

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "aboutme/index.html")
}

func stack(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, aboutme.Get("http://api.stackexchange.com/users/"+stackId+"/timeline?site=stackoverflow.com"))
}

func Page404(w http.ResponseWriter, req *http.Request) (string, int) {
	return "An oopsie!", http.StatusNotFound
}

func StartServer(urlPrefix string) {
	redisClient := aboutme.ConnectRedis()
	defer redisClient.Quit()

	if urlPrefix != "" {
		urlPrefix = "/" + urlPrefix
	}

	http.HandleFunc(urlPrefix+"/", Index)
	http.HandleFunc(urlPrefix+"/stack", stack)
	http.Handle(urlPrefix+"/static/", http.StripPrefix(urlPrefix+"/static", http.FileServer(http.Dir("aboutme/static"))))
	log.Println("aboutme is running...")
}
