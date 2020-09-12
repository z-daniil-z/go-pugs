package main

import (
	"fmt"
	"go-pugs/config"
	"go-pugs/internal/app"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func main() {
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "",
		DSN: fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=disable dbname=%s",
			conf.DataBase.User, conf.DataBase.Password, conf.DataBase.Ip, conf.DataBase.Port, conf.DataBase.DbName),
		PreferSimpleProtocol: false,
		Conn:                 nil,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	application, err := app.NewAPP(db)
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Handler: application.Router(),
		Addr:    fmt.Sprintf("%s:%s", conf.Ip, conf.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
