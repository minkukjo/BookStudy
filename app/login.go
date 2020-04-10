package app

import (
	"encoding/json"
	"golang.org/x/oauth2"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

// Nested 처리는 아래와 같이 내부 구조체를 선언함으로써 가능
type User struct {
	Id          int    `json:"id"`
	ConnectedAt string `json:"connected_at"`
	Properties  struct {
		Nickname string `json:"nickname"`
	} `json:"properties"`
	KakaoAccount struct {
		ProfileNeedsAgreement bool `json:"profile_needs_agreement"`
	}
}

var (
	state = "login"

	conf = &oauth2.Config{
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

func getToken(r *http.Request) Token {
	authorizeCode := r.FormValue("code")

	parameters := url.Values{}
	parameters.Set("grant_type", "authorization_code")
	parameters.Set("client_id", conf.ClientID)
	parameters.Set("redirect_uri", "http://localhost:9090/oauth/authorize")
	parameters.Set("code", authorizeCode)

	resp, err := http.PostForm("https://kauth.kakao.com/oauth/token", parameters)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	token := Token{}
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		log.Fatal(err)
	}

	return token
}

func getUserInform(token Token) User {
	rq, err := http.NewRequest("GET", "https://kapi.kakao.com/v2/user/me", nil)
	if err != nil {
		log.Fatal(err)
	}

	rq.Header.Add("Authorization", "Bearer "+token.AccessToken)
	client := http.Client{}
	resp, err := client.Do(rq)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	user := User{}
	json.NewDecoder(resp.Body).Decode(&user)
	return user
}

func HandleCallBack(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("./static/main.html")
	if err != nil {
		log.Fatal(err)
	}

	token := getToken(r)
	user := getUserInform(token)

	state := r.FormValue("state")
	if state != "login" {
		log.Fatal("Can't Access")
	}

	t.Execute(w, user)
}
