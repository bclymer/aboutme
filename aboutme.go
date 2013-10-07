package main

import (
	"aboutme/aboutme"
	"github.com/AeroNotix/wedge"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
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
		wedge.CacheURL("^/$", "Index", Index, wedge.HTML, -1),
	)
	App.Handler404(Page404)
	//App.EnableStatTracking() // stat tracking on ^/statistics/?$
	App.Run()
}
