package main

import (
	"log"
	"net/http"

	"github.com/NganJason/Dev-Challenges__Full-Stack/tree/master/image-uploader/handler"
	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/img/upload", handler.UploadHandler)
	mux.HandleFunc("/api/img/get", handler.GetImgHandler)
	mux.HandleFunc("/api/healthcheck", handler.HealthCheckHandler)

	handler := cors.Default().Handler(mux)

	log.Println("Listening to port 8082")
	err := http.ListenAndServe(":8082", handler)
	if err != nil {
		log.Fatalf("error initiating server, %s", err.Error())
	}
}