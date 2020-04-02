package app

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    string `json:"expires_in"`
	Scope        string `json:"scope"`
}

var (
	state = "login"
	conf  = &oauth2.Config{
		ClientID:     "33d62682fd47b9b0152a4fa68c14d901",
		ClientSecret: "",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://kauth.kakao.com/oauth/authorize",
			TokenURL: "https://kauth.kakao.com/oauth/token",
		},
		RedirectURL: "http://localhost:9090/oauth/authorize",
		Scopes:      nil,
	}
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	u := conf.AuthCodeURL(state, oauth2.AccessTypeOffline)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func HandleCallBack(w http.ResponseWriter, r *http.Request) {

	authorizeCode := r.FormValue("code")

	fmt.Println(authorizeCode)

	parameters := url.Values{}
	parameters.Set("grant_type", "authorization_code")
	parameters.Set("client_id", conf.ClientID)
	parameters.Set("redirect_uri", "http://localhost:9090/oauth/authorize")
	parameters.Set("code", authorizeCode)

	resp, _ := http.PostForm("https://kauth.kakao.com/oauth/token", parameters)

	defer resp.Body.Close()

	token := new(Token)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	_ = json.Unmarshal(body, &token)
	fmt.Println(token)

	state := r.FormValue("state")
	if state != "login" {
		log.Fatal("Can't Access")
	}

	http.ServeFile(w, r, "./static/main.html")
}
