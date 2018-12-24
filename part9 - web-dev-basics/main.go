package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hey, there!</h1>")
	fmt.Fprintf(w, "<p>Go is fast and simple!</p>")
	fmt.Fprintf(w, "<p>You %s even add %s!</p>", "can", "<strong>variables</strong>")

	fmt.Fprintf(w, `
		<h1>Hey, there!</h1>
		<p>Go is fast and simple!</p>
	`)
}

func main() {
	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8000", nil)
}
