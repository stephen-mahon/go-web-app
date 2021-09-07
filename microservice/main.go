package main

import (
	"log"
	"net/http"

	"github.com/stephen-mahon/go-web-app/microservice/homepage"
	"github.com/stephen-mahon/go-web-app/microservice/server"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homepage.HomeHandler)

	srv := server.New(mux)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
