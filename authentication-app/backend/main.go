package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/healthcheck", healthCheckHandler)

	// handler := cors.Default().Handler(mux)


}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello I am healthy")
}