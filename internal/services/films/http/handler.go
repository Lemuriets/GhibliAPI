package http

import (
	"github.com/Lemuriets/ghibliapi/internal/services/films"
)

type Handler struct {
	UseCase films.UseCase
}

func NewHandler(uc films.UseCase) *Handler {
	return &Handler{
		UseCase: uc,
	}
}
