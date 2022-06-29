package main

import (
	// "log"

	"github.com/Lemuriets/ghibliapi/app"
	"github.com/Lemuriets/ghibliapi/internal/db"
	// "github.com/joho/godotenv"
)

func main() {
	// without docker:

	// if err := godotenv.Load(".env"); err != nil {
	// 	log.Fatal("Failed to load env: ", err)
	// }
	database := db.NewDB()
	db.InitDB(database)

	app := app.NewApp(database)
	app.Start()
}
