package handlers

import (
	"app/models"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
)

var ctx context.Context

func (h *Handlers) CreateUser() {
	var User models.User

	User.UserID = uuid.New()

	fmt.Println("Enter Fullname: ")
	fmt.Scanln(&User.Fullname)

	fmt.Println("Enter Username: ")
	fmt.Scanln(&User.Username)

	fmt.Println("Enter Gmail: ")
	fmt.Scanln(&User.Gmail)

	fmt.Println("Enter Password: ")
	fmt.Scanln(&User.Password)

	log.Println("request in here")
	fmt.Println(h.storage.GetUserRepo())

	h.storage.GetUserRepo().CreateUser(ctx, models.User{})
	log.Println("request in here 2")

		
}
