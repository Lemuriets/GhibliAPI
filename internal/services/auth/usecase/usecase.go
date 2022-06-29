package usecase

import (
	"github.com/Lemuriets/ghibliapi/internal/services/auth"
)

type UseCase struct {
	Repository auth.Repository
}

func NewUseCase(repo auth.Repository) UseCase {
	return UseCase{
		Repository: repo,
	}
}
