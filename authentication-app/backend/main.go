package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/processor"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/clog"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/wrapper"
	"github.com/rs/cors"
)

func main() {
	ctx := context.Background()
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

	clog.SetMinLogLevel(clog.LevelInfo)
	handler := cors.Default().Handler(mux)

	clog.Info(ctx, "Listening to port 8082")
	err := http.ListenAndServe(":8082", handler)
	if err != nil {
		clog.Fatal(ctx, fmt.Sprintf("error init server, %s", err.Error()))
	}
}
