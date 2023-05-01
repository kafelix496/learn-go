package url_checker

import (
	"errors"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

type UrlResult struct {
	Url    string
	Status string
}

func HitUrl(url string, c chan<- UrlResult) {
	resp, err := http.Get(url)
	status := "SUCCESS"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}

	c <- UrlResult{Url: url, Status: status}
}
