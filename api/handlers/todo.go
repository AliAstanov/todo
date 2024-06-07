package handlers

import "app/models"

func (h *Handlers)CreateTodo() {
	h.storage.GetTodoRepo().CreateUser(ctx,models.Todo{})
}