package app

import (
	"log"
	"net/http"
)

func request(url string, accessToken string) (*http.Response, error) {
	rq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	rq.Header.Add("Authorization", "Bearer "+accessToken)
	client := http.Client{}

	return client.Do(rq)
}
