package main

import (
	"gohound/internal/app"
	"log"
	"net/http"
	"time"
)

func main() {
	api := app.NewAPI()

	srv := &http.Server{
		Handler: api.Router(),
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
