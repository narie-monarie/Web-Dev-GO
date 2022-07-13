package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to my site</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to my Contact site man</h1>")
}

//Raw path - used mostly in emails because its url is ?key=value
func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		//TODO: handle page not found
		http.Error(w, "Page not Found", http.StatusNotFound)
	}
}

func main() {
	//http.HandleFunc("/", pathHandler)
	//http.HandleFunc("/contact", contactHandler)
	fmt.Println("listening on localhost 3000...")
	http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
}
