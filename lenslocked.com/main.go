package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact</h1><p>To get in touch email me at <a href=\"mailto:stemahon@gmail.com\">stemahon@gmail.com</a>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=uft-8")
	fmt.Fprintf(w, "<h1>Frequently Asked Questions</h1>")
}

func notfoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>We could not find the page you were looking for!</h1>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact/":
		contactHandler(w, r)
	case "/faq/":
		faqHandler(w, r)
	default:
		notfoundHandler(w, r)
	}

}

func main() {
	//r := mux.NewRouter()
	http.HandleFunc("/", pathHandler)
	//r.HandleFunc("/", pathHandler)
	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", nil)
}
