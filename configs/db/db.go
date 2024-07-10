package db

import (
	"fmt"
	"technical-test-be-abc/configs"
	"technical-test-be-abc/domains"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	cfg := configs.LoadConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", cfg.DB.DBHOST, cfg.DB.DBUSER, cfg.DB.DBPASSWORD, cfg.DB.DBNAME, cfg.DB.DBPORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
		panic(err)
	}
	db.AutoMigrate(&domains.Todo{}, &domains.User{})
	logrus.Println("db is connected")
	return db
}
