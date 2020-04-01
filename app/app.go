package app

import (
	"net/http"
)

type App struct{}

func (a *App) Run() {

	//r := NewRouter()


	http.ListenAndServe(":9090",nil)
}
