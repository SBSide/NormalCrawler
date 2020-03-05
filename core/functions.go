package core

import (
	"net/http"
)

type RequestResult struct {
	Url    string
	Status string
}

func HitURL(url string, c chan<- RequestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- RequestResult{
		url,
		status,
	}
}
