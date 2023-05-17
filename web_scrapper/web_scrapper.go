package web_scrapper

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ExtractedJob struct {
	title    string
	location string
}

func GetPagesCount(url string) int {
	pagesCount := 0
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pagesCount = s.Find("a").Length()
	})

	return pagesCount
}

func GetTargetPageUrl(url string, page int) string {
	return url + "&recruitPage=" + strconv.Itoa(page)
}

func GetPage(url string, mainC chan<- []ExtractedJob) {
	var jobs []ExtractedJob
	c := make(chan ExtractedJob)

	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")
	searchCards.Each(func(i int, s *goquery.Selection) {
		go extractJob(s, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func WriteJobs(jobs []ExtractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Title", "Location"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{job.title, job.location}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func extractJob(s *goquery.Selection, c chan<- ExtractedJob) {
	title := cleanString(s.Find("h2.job_tit").Text())
	location := cleanString(s.Find(".job_condition span").First().Text())

	c <- ExtractedJob{title, location}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Error: ", res.StatusCode)
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
