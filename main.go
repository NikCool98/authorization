package main

import (
	"github.com/NikCool98/authorization/config"
	"github.com/NikCool98/authorization/storage"
	"os"
)

func main() {
	db := storage.NewDB(config.StoragePath)
	storage.InsertInDb(db)
	os.Exit(1)
}
