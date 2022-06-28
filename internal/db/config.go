package db

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Config gorm.Config = gorm.Config{
	Logger: logger.Default.LogMode(logger.Info),
}
