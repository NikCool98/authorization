package main

import (
	"github.com/NikCool98/authorization/config"
	"github.com/NikCool98/authorization/handlers"
	"github.com/NikCool98/authorization/storage"
	"log"
	"net/http"
)

func main() {
	db, err := storage.NewDB(config.StoragePath)
	if err != nil {
		log.Fatalf("Failed to init storage: %v", err)
	}
	defer db.Close()
	store := storage.NewStore(db)
	http.HandleFunc("/", handlers.StartPageHandler(store))
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalf("Server run error: %v", err)
	}
}
