package aboutme

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/go-oauth/oauth"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	TwitterId = "crdnilfan"
)

var (
	twitterClient TwitterApi
	oauthClient   = oauth.Client{
		TemporaryCredentialRequestURI: "http://api.twitter.com/oauth/request_token",
		ResourceOwnerAuthorizationURI: "http://api.twitter.com/oauth/authenticate",
		TokenRequestURI:               "http://api.twitter.com/oauth/access_token",
	}
)

type TwitterApi struct {
	Credentials *oauth.Credentials
}

type TwitterAuth struct {
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
}

func init() {
	var twitterAuth TwitterAuth
	content, err := ioutil.ReadFile("aboutme/twitterAuth.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &twitterAuth)
	if err != nil {
		panic(err)
	}
	oauthClient.Credentials.Token = twitterAuth.ConsumerKey
	oauthClient.Credentials.Secret = twitterAuth.ConsumerSecret
	twitterClient = TwitterApi{&oauth.Credentials{Token: twitterAuth.AccessToken, Secret: twitterAuth.AccessTokenSecret}}
}

func TwitterTimeline(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, getTwitterUrl("http://api.twitter.com/1.1/statuses/user_timeline.json"))
}

func getTwitterUrl(url string) string {
	responseChannel := make(chan (RedisResponse))
	redisRequest := RedisRequest{
		Get:             true,
		Key:             url,
		ResponseChannel: responseChannel,
	}
	RedisRequestChannel <- redisRequest
	redisResponse := <-redisRequest.ResponseChannel
	if redisResponse.Err != nil {
		log.Println("Cache miss -", url, redisResponse.Err)
		resp, err := oauthClient.Get(http.DefaultClient, twitterClient.Credentials, url, nil)
		if err != nil {
			log.Println("Twitter Err Get", err)
			return ""
		}
		if resp.StatusCode != 200 {
			p, _ := ioutil.ReadAll(resp.Body)
			fmt.Errorf("Get %s returned status %d, %s", resp.Request.URL, resp.StatusCode, p)
			return ""
		}
		responseBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Twitter Err Read", err)
			return ""
		}
		response := string(responseBytes)
		redisRequest = RedisRequest{
			Get:   false,
			Key:   url,
			Value: response,
		}
		RedisRequestChannel <- redisRequest
		return response
	}
	return redisResponse.Value
}
