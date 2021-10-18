package eraser

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	url         = "https://twitter.com/i/api/1.1/statuses/destroy.json"
	method      = "POST"
	basePayload = "tweet_mode=extended&id="
)

type authData struct {
	authToken          string
	csrfToken          string
	authorizationToken string
}

type TweeterData struct {
	AuthData authData
}

var staticHeaders map[string]string = map[string]string{
	"authority":                 "twitter.com",
	"sec-ch-ua":                 "\"Google Chrome\";v=\"93\", \" Not;A Brand\";v=\"99\", \"Chromium\";v=\"93\"",
	"dnt":                       "1",
	"x-twitter-client-language": "fa",
	"sec-ch-ua-mobile":          "?0",
	"content-type":              "application/x-www-form-urlencoded",
	"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36",
	"x-twitter-auth-type":       "OAuth2Session",
	"x-twitter-active-user":     "yes",
	"sec-ch-ua-platform":        "\"macOS\"",
	"accept":                    "*/*",
	"origin":                    "https://twitter.com",
	"sec-fetch-site":            "same-origin",
	"sec-fetch-mode":            "cors",
	"sec-fetch-dest":            "empty",
	"referer":                   "https://twitter.com/home",
	"accept-language":           "en;q=0.9,fa;q=0.8,nl;q=0.7",
}

func NewTweeterData() *TweeterData {
	return &TweeterData{}
}

func (t *TweeterData) AuthToken(token string) {
	t.AuthData.authToken = token
}

func (t *TweeterData) CSRFToken(token string) {
	t.AuthData.csrfToken = token
}

func (t *TweeterData) AuthorizationToken(token string) {
	t.AuthData.authorizationToken = token
}

func (t *TweeterData) buildCookie() string {
	return "auth_token=" + t.AuthData.authToken + "; ct0=" + t.AuthData.csrfToken
}

func (t *TweeterData) DeleteTweet(TweetID string, client *http.Client) (int, []byte, error) {
	payload := strings.NewReader(basePayload + TweetID)

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return 0, nil, err
	}

	for k, v := range staticHeaders {
		req.Header.Add(k, v)
	}
	req.Header.Add("cookie", t.buildCookie())
	req.Header.Add("x-csrf-token", t.AuthData.csrfToken)
	req.Header.Add("authorization", "Bearer "+t.AuthData.authorizationToken)

	res, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}

	return res.StatusCode, body, nil
}
