package main

import (
	"github.com/NikCool98/authorization/config"
	"github.com/NikCool98/authorization/storage"
	"log"
)

func main() {
	db, err := storage.NewDB(config.StoragePath)
	if err != nil {
		log.Fatalf("Failed to init storage: %v", err)
	}
	defer db.Close()
	store := storage.NewStore(db)
	_ = store
}
