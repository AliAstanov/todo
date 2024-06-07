package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	UserID   uuid.UUID `json:"user_id"`
	Fullname string    `json:"fullname"`
	Username string    `json:"username"`
	Gmail    string    `json:"gmail"`
	Password string    `json:"password"`
}

// Todo modeli
type Todo struct {
	TodoID      uuid.UUID `json:"todo_id"`
	Task        string    `json:"task"`
	CreatedAt   time.Time `json:"created_at"`
	IsComplated bool      `json:"is_complated"`
}

type GetTodosResp struct {
	Todos []Todo
	Count int
}
