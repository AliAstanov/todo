package handlers

import (
	"app/models"
	"context"
)

var ctx context.Context

func (h *Handlers) CreateUser() {
	h.storage.GetUserRepo().CreateUser(ctx,models.User{})
}