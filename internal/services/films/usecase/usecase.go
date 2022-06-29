package usecase

import (
	"github.com/Lemuriets/ghibliapi/internal/services/films"
)

type UseCase struct {
	Repository films.Repository
}

func NewUseCase(repo films.Repository) UseCase {
	return UseCase{
		Repository: repo,
	}
}
