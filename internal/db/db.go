package db

import (
	"log"
	"os"

	"github.com/Lemuriets/ghibliapi/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := os.Getenv("dsn")
	db, err := gorm.Open(postgres.Open(dsn), &Config)

	if err != nil {
		log.Fatal("Failed to connect db: ", err)
	}
	return db
}

func InitDB(db *gorm.DB) error {
	createSuperUser(db)
	return db.AutoMigrate(
		&model.User{},
		&model.Film{},
	)
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
