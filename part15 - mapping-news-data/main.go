package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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

func main() {
	var s SitemapIndex
	var n News

	newsMap := make(map[string]NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		var trimmedLocation = strings.TrimSpace(Location)

		resp, _ := http.Get(trimmedLocation)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		for idx := range n.Titles {
			newsMap[n.Titles[idx]] = NewsMap{
				n.Keywords[idx],
				n.Locations[idx],
			}
		}
	}

	for idx, data := range newsMap {
		fmt.Print("\n\n")
		fmt.Println(idx)
		fmt.Println(data.Keyword)
		fmt.Println(data.Location)
	}
}
