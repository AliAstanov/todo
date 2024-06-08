package main

import (
	"app/api"
	"app/config"
	"app/pkg/db"
	"app/storage"
	"log"
)

func main()  {
	cfg := config.Load()
	
	conn, err := db.ConnToDb(cfg)
	if err != nil {
		log.Println("error on connecting with database:",err)
		return
	}
	log.Println("succesfully connected with database!")

	storage := storage.NewStorage(conn)
	log.Println("storae succesfully")

	api.Api(storage)
	log.Println("Api succesfully")

}