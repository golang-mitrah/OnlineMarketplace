package driver

import (
	"fmt"
	"onlinemarketplace/model"
	"os"

	"github.com/joho/godotenv"
	logging "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func NewDBConn(IsTest bool) {
	envpath := ".env"
	if IsTest {
		envpath = "../" + envpath
	}
	err := godotenv.Load(envpath)
	if err != nil {
		panic("Error while loading .env file")
	}

	userName := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	if IsTest {
		dbName = dbName + "_test"
	}
	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, userName, password, dbName, port)
	Db, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	logging.Info("Database connected succesfully", dbName)
	Db.AutoMigrate(&model.Product{}, &model.User{})
}

func GetDBConn() *gorm.DB {
	if Db == nil {
		panic("Database not connected")
	}
	return Db
}
