package app

import (
	"gorm.io/gorm"

	authhttp "github.com/Lemuriets/ghibliapi/internal/services/auth/http"
	authrepo "github.com/Lemuriets/ghibliapi/internal/services/auth/repository"
	authuc "github.com/Lemuriets/ghibliapi/internal/services/auth/usecase"

	filmshttp "github.com/Lemuriets/ghibliapi/internal/services/films/http"
	filmsrepo "github.com/Lemuriets/ghibliapi/internal/services/films/repository"
	filmsuc "github.com/Lemuriets/ghibliapi/internal/services/films/usecase"

	usershttp "github.com/Lemuriets/ghibliapi/internal/services/users/http"
	usersrepo "github.com/Lemuriets/ghibliapi/internal/services/users/repository"
	usersuc "github.com/Lemuriets/ghibliapi/internal/services/users/usecase"
)

func RegisterAuthService(db *gorm.DB) *authhttp.Handler {
	repo := authrepo.NewRepository(db)
	uc := authuc.NewUseCase(repo)

	return authhttp.NewHandler(uc)
}

func RegisterFilmsService(db *gorm.DB) *filmshttp.Handler {
	repo := filmsrepo.NewRepository(db)
	uc := filmsuc.NewUseCase(repo)

	return filmshttp.NewHandler(uc)
}

func RegisterUsersService(db *gorm.DB) *usershttp.Handler {
	repo := usersrepo.NewRepository(db)
	uc := usersuc.NewUseCase(repo)

	return usershttp.NewHandler(uc)
}
