package app

import (
	"bookstudy/db"
	"bookstudy/redis"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"net/http"
	"strconv"
	"time"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.NoCache)
		r.Use(Authenticator)

		r.HandleFunc("/main", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "public/index.html")
		})

		r.HandleFunc("/logout", logoutHandler)
	})

	// Public routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.NoCache)

		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/main", http.StatusTemporaryRedirect)
		})

		r.HandleFunc("/login", loginHandler)

		r.HandleFunc("/kakaologin", HandleLogin)

		r.HandleFunc("/oauth/authorize", HandleCallBack)

		r.HandleFunc("/temp", func(w http.ResponseWriter, r *http.Request) {

			accessToken := r.FormValue("token")

			fmt.Println(accessToken)

			user := getUserInform(accessToken)

			err := redis.RedisClient.Set(strconv.Itoa(user.Id), accessToken, 0).Err()
			if err != nil {
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}

			db.InsertUser(user)

			_, _, err = tokenAuth.Encode(jwt.MapClaims{"user_id": user.Id})
			if err != nil {
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}

			cookie := &http.Cookie{
				Name:     defaultAuthCookieName,
				Value:    accessToken,
				Expires:  time.Now().Add(defaultSessionExpire),
				HttpOnly: true,
			}
			http.SetCookie(w, cookie)

			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		})
	})

	r.Handle("/*", http.FileServer(http.Dir("public")))

	return r
}
