package main

import (
	"fmt"

	"github.com/kafelix496/learn-go/url_checker"
)

func main() {
	results := make(map[string]string)
	c := make(chan url_checker.UrlResult)
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.instagram.com",
		"https://www.youtube.com",
		"https://www.linkedin.com",
		"https://www.github.com",
		"https://www.medium.com",
		"https://www.reddit.com",
		"https://www.stackoverflow.com",
		"https://www.wikipedia.org",
		"https://www.quora.com",
		"https://www.pinterest.com",
	}

	for _, url := range urls {
		go url_checker.HitUrl(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.Url] = result.Status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}
