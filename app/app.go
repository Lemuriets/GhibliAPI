package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	Router *gin.Engine
	AuthHandler
	UsersHandler
	FilmsHandler
}

func NewApp(db *gorm.DB) *App {
	return &App{
		Router:       gin.Default(),
		AuthHandler:  RegisterAuthHandler(db),
		UsersHandler: RegisterUsersHandler(db),
		FilmsHandler: RegisterFilmsHandler(db),
	}
}

func (app *App) Start() {

}
