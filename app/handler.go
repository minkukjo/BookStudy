package app

import (
	"bookstudy/db"
	"bookstudy/model"
	"bookstudy/redis"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	state = "login"

	defaultAuthCookieName = "user_id"
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

	resp, err := request("https://kapi.kakao.com/v2/user/me", accessToken, "GET")
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

	http.Redirect(w, r, "/session?token="+token.AccessToken, http.StatusTemporaryRedirect)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/login.html")
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	accessToken, err := redis.RedisClient.Get(cookie.Value).Result()

	if accessToken == "" || err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	_, err = request("https://kapi.kakao.com/v1/user/logout", accessToken, "POST")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	deleteSession(w)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("user_id")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		accessToken, err := redis.RedisClient.Get(cookie.Value).Result()

		if accessToken == "" || err != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		// 토큰 유효성 검사
		_, err = request("https://kapi.kakao.com/v1/user/access_token_info", accessToken, "GET")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}

// redis 저장, DB 저장
// redis의 용도는 토큰 유효성 검사 시 조금 더 효율적으로 하기 위함
func HandleSession(w http.ResponseWriter, r *http.Request) {

	accessToken := r.FormValue("token")

	user := getUserInform(accessToken)

	err := redis.RedisClient.Set(strconv.Itoa(user.Id), accessToken, 0).Err()
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	db.InsertUser(user)

	cookie := &http.Cookie{
		Name:     defaultAuthCookieName,
		Value:    strconv.Itoa(user.Id),
		Expires:  time.Now().Add(defaultSessionExpire),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func deleteSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     defaultAuthCookieName,
		Value:    "none",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
}

func HandleUserInform(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	id, err := strconv.Atoi(cookie.Value)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	user := model.User{
		Id: id,
	}

	if !db.FindFirstUser(&user) {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
}

func HandleWrite(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	id, err := strconv.Atoi(cookie.Value)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	user := model.User{
		Id: id,
	}

	if !db.FindFirstUser(&user) {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	post := model.Post{
		UserId: user.Id,
		Date:   time.Now().Format("2006-01-02"),
		Name:   user.Nickname,
	}

	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Fatal(err)
	}

	db.InsertPost(&post)
}

func HandlePosts(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("kind")

	posts := db.FindAllPosts("kind", param)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&posts)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
}
