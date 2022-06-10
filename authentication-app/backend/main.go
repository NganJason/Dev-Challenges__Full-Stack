package main

import (
	"log"
	"net/http"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/processor"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/wrapper"
	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()

	for _, proc := range processor.GetAllProcessors() {
		mux.HandleFunc(
			proc.Path,
			wrapper.WrapProcessor(
				proc.Processor,
				proc.Req,
				proc.Resp,
				proc.Cookie,
			),
		)
	}

	handler := cors.Default().Handler(mux)

	log.Println("Listening to port 8082")
	err := http.ListenAndServe(":8082", handler)
	if err != nil {
		log.Fatalf("error initiating server, %s", err.Error())
	}
}
