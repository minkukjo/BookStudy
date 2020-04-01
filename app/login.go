package app

import "golang.org/x/oauth2"

var(
	conf = &oauth2.Config{
		ClientID:     "33d62682fd47b9b0152a4fa68c14d901",
		ClientSecret: "",
		Endpoint:     oauth2.Endpoint{
			AuthURL: "https://kauth.kakao.com/oauth/authorize",
			TokenURL: "https://kauth.kakao.com/oauth/token",
		},
		RedirectURL:  "/main",
		Scopes:       nil,
	}
)

func authorization() {

}