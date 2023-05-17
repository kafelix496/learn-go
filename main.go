package main

import (
	"github.com/kafelix496/learn-go/web_scrapper"
)

type extractedJob struct {
	title    string
	location string
}

var baseURL = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python"

func main() {
	totalPageCount := web_scrapper.GetPagesCount(baseURL)
	var jobs []web_scrapper.ExtractedJob
	c := make(chan []web_scrapper.ExtractedJob)

	for i := 0; i < totalPageCount; i++ {
		pageUrl := web_scrapper.GetTargetPageUrl(baseURL, i)

		go web_scrapper.GetPage(pageUrl, c)
	}

	for i := 0; i < totalPageCount; i++ {
		job := <-c
		jobs = append(jobs, job...)
	}

	web_scrapper.WriteJobs(jobs)
}
