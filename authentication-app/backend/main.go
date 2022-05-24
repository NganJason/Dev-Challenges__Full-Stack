package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/healthcheck", healthCheckHandler)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello I am healthy")
}