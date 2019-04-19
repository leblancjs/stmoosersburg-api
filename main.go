package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	_ "github.com/lib/pq"

	"github.com/leblancjs/stmoosersburg-api/db"
	"github.com/leblancjs/stmoosersburg-api/hash"
	"github.com/leblancjs/stmoosersburg-api/user"
)

func main() {
	database, err := configureDatabase()
	if err != nil {
		log.Fatal(err)
	}
	err = database.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	hashProvider := hash.NewBCryptProvider()
	hashSvc, err := hash.NewService(hashProvider)
	if err != nil {
		log.Fatal(err)
	}

	userRepo, err := user.NewRepository(database)
	if err != nil {
		log.Fatal(err)
	}
	userSvc, err := user.NewService(userRepo, hashSvc)
	if err != nil {
		log.Fatal(err)
	}
	userHandler := user.MakeHandler(userSvc)

	mux := http.NewServeMux()
	mux.Handle("/v1/", userHandler)

	http.Handle("/", handlers.LoggingHandler(os.Stdout, mux))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func configureDatabase() (db.DB, error) {
	databaseType := os.Getenv("DB_TYPE")
	if databaseType == "" {
		databaseType = db.TypeInMemory
	}

	databaseConfig := db.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	database, err := db.New(databaseType, databaseConfig)
	if err != nil {
		return nil, err
	}

	return database, nil
}
