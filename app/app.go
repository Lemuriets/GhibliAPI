package app

import (
	"io"
	"log"
	"os"

	authhttp "github.com/Lemuriets/ghibliapi/internal/services/auth/http"
	filmshttp "github.com/Lemuriets/ghibliapi/internal/services/films/http"
	usershttp "github.com/Lemuriets/ghibliapi/internal/services/users/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	Router       *gin.Engine
	AuthHandler  *authhttp.Handler
	UsersHandler *usershttp.Handler
	FilmsHandler *filmshttp.Handler
}

func NewApp(db *gorm.DB) *App {
	return &App{
		Router:       gin.Default(),
		AuthHandler:  RegisterAuthService(db),
		UsersHandler: RegisterUsersService(db),
		FilmsHandler: RegisterFilmsService(db),
	}
}

func (app *App) Start() {
	configureGin()
	app.Router.Run(":8000")
}

func configureGin() {
	logFile, err := os.Create("gin.log")
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
}
