package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// NewsAggPage Structure
type NewsAggPage struct {
	Title string
	News  string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Woah, Go is neat!</h1>")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{
		Title: "Amazing News Aggregator",
		News:  "Some News Here",
	}

	t, _ := template.ParseFiles("basic-template.html")
	fmt.Println(t.Execute(w, p))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)

	http.ListenAndServe(":8000", nil)
}
