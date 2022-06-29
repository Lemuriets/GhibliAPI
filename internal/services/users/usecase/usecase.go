package usecase

import (
	"github.com/Lemuriets/ghibliapi/internal/services/users"
)

type UseCase struct {
	Repository users.Repository
}

func NewUseCase(repo users.Repository) UseCase {
	return UseCase{
		Repository: repo,
	}
}
