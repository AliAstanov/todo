package handlers

import "app/storage"

type Handlers struct {
	storage storage.StorageI
}

func NewHendlers(storage storage.StorageI) Handlers {
	return Handlers{storage: storage}
}