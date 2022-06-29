package http

import (
	"github.com/Lemuriets/ghibliapi/internal/services/users"
)

type Handler struct {
	UseCase users.UseCase
}

func NewHandler(uc users.UseCase) *Handler {
	return &Handler{
		UseCase: uc,
	}
}
