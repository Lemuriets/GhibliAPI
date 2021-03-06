package db

import (
	"fmt"
	"log"
	"os"

	"github.com/Lemuriets/ghibliapi/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s",
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &Config)

	if err != nil {
		log.Fatal("Failed to connect db: ", err)
	}
	return db
}

func InitDB(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.User{},
		&model.Film{},
	)
	createSuperUser(db)
	return err
}

func createSuperUser(db *gorm.DB) error {
	isSuperUserExists, err := checkSuperUser(db)

	if err != nil {
		return err
	}
	if !isSuperUserExists {
		superUser := model.User{
			Login:    "Lemuriets",
			Password: os.Getenv("superuserpassw"),
		}
		return db.Create(&superUser).Error
	}
	return nil
}

func checkSuperUser(db *gorm.DB) (bool, error) {
	var user model.User
	err := db.First(&user, 1).Error

	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return false, nil
	}
	return true, nil
}
