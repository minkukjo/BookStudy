package app

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.NoCache)
		r.Use(Authenticator)

		r.HandleFunc("/main", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "public/index.html")
		})

		r.HandleFunc("/api/user", HandleUserInform)

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

		r.HandleFunc("/session", HandleSession)
	})

	r.Handle("/*", http.FileServer(http.Dir("public")))

	return r
}
