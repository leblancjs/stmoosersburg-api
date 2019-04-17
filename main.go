package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"github.com/leblancjs/stmoosersburg-api/hash"
	"github.com/leblancjs/stmoosersburg-api/user"
)

func main() {
	hashProvider := hash.NewBCryptProvider()
	hashSvc, err := hash.NewService(hashProvider)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := user.NewInMemoryRepository()
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
