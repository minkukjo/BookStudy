package app

import (
	"bookstudy/model"
	"bookstudy/redis"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/jwtauth"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"time"
)

// Nested 처리는 아래와 같이 내부 구조체를 선언함으로써 가능

var (
	state = "login"

	defaultAuthCookieName = "user_token"
	defaultSessionExpire  = 6 * time.Hour

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

func getUserInform(accessToken string) model.User {

	resp, err := request("https://kapi.kakao.com/v2/user/me", accessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	userJson := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&userJson)
	if err != nil {
		log.Fatal(err)
	}
	properties := userJson["properties"].(map[string]interface{})

	fmt.Print()

	user := model.User{
		Id:          int(userJson["id"].(float64)),
		ConnectedAt: userJson["connected_at"].(string),
		Nickname:    properties["nickname"].(string),
		Token:       accessToken,
	}
	return user
}

// 지금은 userId를 최초로 parameter argument로만 넘겨서 인증하고 있음
// 이렇게하면 프론트에서 페이지를 이동할 때 마다 검사가 이뤄지지 않음.
// 결국 다른 방법을 생각해봐야함
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

	http.Redirect(w, r, "/temp?token="+token.AccessToken, http.StatusTemporaryRedirect)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/login.html")
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {

}

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		strings := r.Header["Cookie"]
		token, claims, err := jwtauth.FromContext(r.Context())
		fmt.Println("this is user id ------")
		fmt.Println(token)
		fmt.Println(strings)
		fmt.Println(claims["user_token"])
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		Id := r.FormValue("Id")

		if Id == "" {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		accessToken, _ := redis.RedisClient.Get(Id).Result()

		if accessToken == "" {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		resp, err := request("https://kapi.kakao.com/v1/user/access_token_info", accessToken)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		defer resp.Body.Close()

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}
