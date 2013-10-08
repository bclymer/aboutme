package main

import (
	"aboutme/aboutme"
	"github.com/AeroNotix/wedge"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

const (
	stackId = "650288"
)

var (
	curDir string
)

// Main page
func Index(w http.ResponseWriter, req *http.Request) (string, int) {
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		panic(err)
	}
	return string(content), http.StatusOK
}

func stack(w http.ResponseWriter, req *http.Request) (string, int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return aboutme.Get("http://api.stackexchange.com/users/" + stackId + "/timeline?site=stackoverflow.com"), http.StatusOK
}

func Page404(w http.ResponseWriter, req *http.Request) (string, int) {
	return "An oopsie!", http.StatusNotFound
}

func main() {
	if curDir == "" {
		curDir, _ = os.Getwd()
	}

	redisClient := aboutme.ConnectRedis()
	defer redisClient.Quit()

	App := wedge.NewAppServer("8080", 30)
	App.AddURLs(
		wedge.Favicon(filepath.Join(curDir, "static/images", "favicon.ico")),
		wedge.StaticFiles("/static/", filepath.Join(curDir, "static/")),
		wedge.URL("stack", "stackoverflow", stack, wedge.HTML),
		wedge.URL("^/$", "Index", Index, wedge.HTML), //, -1),
	)
	App.Handler404(Page404)
	//App.EnableStatTracking() // stat tracking on ^/statistics/?$
	App.Run()
}
