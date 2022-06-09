package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/processor"
	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/healthcheck", healthCheckHandler)

	mux.HandleFunc("/api/login/github", processor.GithubLoginProcessor)

	handler := cors.Default().Handler(mux)

	log.Println("Listening to port 8082")
	err := http.ListenAndServe(":8082", handler)
	if err != nil {
		log.Fatalf("error initiating server, %s", err.Error())
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello I am healthy")
}
