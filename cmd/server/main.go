package main

import (
	"fmt"
	"go-pugs/config"
	"go-pugs/internal/app"
	"log"
	"net/http"
	"time"
)

func main() {
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	api := app.NewAPI()

	srv := &http.Server{
		Handler: api.Router(),
		Addr:    fmt.Sprintf("%s:%s", conf.Ip, conf.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
