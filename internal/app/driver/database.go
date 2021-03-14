package driver

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Instance *gorm.DB
}

type DBConfig struct {
	Host     string
	Port     string
	Name     string
	UserName string
	Password string
}

func NewDatabase(cfg DBConfig) (database *Database, err error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.UserName,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return
	}

	database = &Database{
		Instance: db,
	}

	return
}
