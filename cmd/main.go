package main

import (
	"app/config"
	"app/pkg/db"
	"log"
)

func main()  {
	cfg := config.Load()
	
	conn, err := db.ConnToDb(cfg)
	if err != nil {
		log.Println("error on connecting with database:",err)
		return
	}
	log.Println("succesfully connected with database!",conn)
}