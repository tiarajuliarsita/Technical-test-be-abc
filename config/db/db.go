package db

import (
	"fmt"
	"technical-test-be-abc/config"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	cfg := config.LoadDB()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", cfg.DBHOST, cfg.DBUSER, cfg.DBPASSWORD, cfg.DBNAME, cfg.DBPORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
		panic(err)
	}
	logrus.Println("db is connected")
	return db
}
