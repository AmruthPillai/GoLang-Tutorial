package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"text/template"
)

var wg sync.WaitGroup

// SitemapIndex Structure
type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

// News Structure
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

// NewsMap Structure
type NewsMap struct {
	Keyword  string
	Location string
}

// NewsAggPage Structure
type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Woah, Go is neat!</h1>")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s SitemapIndex
	var n News

	newsMap := make(map[string]NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	resp.Body.Close()

	queue := make(chan News, 30)

	for _, Location := range s.Locations {
		wg.Add(1)
		var trimmedLocation = strings.TrimSpace(Location)
		go newsRoutine(queue, trimmedLocation)
	}

	wg.Wait()
	close(queue)

	for elem := range queue {
		for idx := range n.Titles {
			newsMap[elem.Titles[idx]] = NewsMap{
				elem.Keywords[idx],
				elem.Locations[idx],
			}
		}
	}

	p := NewsAggPage{Title: "Amazing News Aggregator", News: newsMap}
	t, _ := template.ParseFiles("news-template.html")
	fmt.Println(t.Execute(w, p))
}

func newsRoutine(c chan News, Location string) {
	defer wg.Done()

	var n News
	resp, _ := http.Get(Location)
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()

	c <- n
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)

	http.ListenAndServe(":8000", nil)
}
