package app

import (
	"log"
	"net/http"
)

func request(url string, accessToken string, method string) (*http.Response, error) {
	rq, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	rq.Header.Add("Authorization", "Bearer "+accessToken)
	client := http.Client{}

	return client.Do(rq)
}
