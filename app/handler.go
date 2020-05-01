package app

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"time"
)

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

func getUserInform(token *oauth2.Token) User {
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
	state := r.FormValue("state")
	if state != "login" {
		log.Fatal("Can't Access")
	}
	c := r.FormValue("code")

	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx := context.WithValue(context.TODO(), oauth2.HTTPClient, httpClient)

	token, err := conf.Exchange(ctx, c)
	if err != nil {
		log.Fatal(err)
	}
	client := conf.Client(ctx, token)
	_ = client

	user := getUserInform(token)

	err = RedisClient.Set(token.AccessToken, user.Properties.Nickname, 0).Err()
	if err != nil {
		log.Fatal(err)
	}
	//http.SetCookie(w,&http.Cookie{
	//	Name: "access_token",
	//	Value: token.AccessToken,
	//	Expires: time.Now().Add(15 * time.Second),
	//})

	http.Redirect(w, r, "/main"+"?access_token="+token.AccessToken, http.StatusTemporaryRedirect)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/login.html")
}

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.FormValue("access_token")
		if token == "" {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		userName, _ := RedisClient.Get(token).Result()

		if userName == "" {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}
