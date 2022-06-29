package http

import (
	"github.com/Lemuriets/ghibliapi/internal/services/auth"
)

type Handler struct {
	UseCase auth.UseCase
}

func NewHandler(uc auth.UseCase) *Handler {
	return &Handler{
		UseCase: uc,
	}
}
