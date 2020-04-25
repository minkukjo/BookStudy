package app

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"net/http"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("765678d6"), nil)
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(middleware.NoCache)
		r.Use(jwtauth.Verifier(tokenAuth))

		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "static/index.html")
		})

		r.HandleFunc("/login", HandleLogin)

		r.HandleFunc("/oauth/authorize", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/home", 301)
		})
		r.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "static/home.html")
		})

	})

	r.Handle("/*", http.FileServer(http.Dir("static")))

	return r
}
