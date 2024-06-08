package api

import (
	"app/api/handlers"
	"app/storage"
)

func Api(storage storage.StorageI) {
	h := handlers.NewHendlers(storage)

	h.CreateUser()
}