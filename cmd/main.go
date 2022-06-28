package main

import (
	"github.com/Lemuriets/ghibliapi/app"
	"github.com/Lemuriets/ghibliapi/internal/db"
)

func main() {
	database := db.NewDB()
	db.InitDB(database)

	app := app.NewApp(database)
	app.Start()
}
