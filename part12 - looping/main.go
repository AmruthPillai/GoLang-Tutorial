package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Looping Syntax
func syntax() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	x := 5
	for {
		fmt.Println("Do stuff", x)
		x += 3
		if x >= 25 {
			break
		}
	}

	for x := 5; x < 25; x += 3 {
		fmt.Println("Do stuff", x)
	}
}

// SitemapIndex - Sitemap Index Structure
type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

// Location - Location Structure
type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	// fmt.Println(s.Locations)

	for _, Location := range s.Locations {
		fmt.Printf("%s", Location)
	}
}
