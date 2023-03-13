package db

import (
	"fmt"
	"github.com/natsukirigakure/publickey-repository/model/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	dbType := "postgres"
	dbHost := "127.0.0.1"
	dbPort := 5432
	dbUser := "gorm"
	dbPass := "gorm"
	dbName := "gorm"
	dbSslMode := "disable"

	switch dbType {
	case "postgres":
		dsn := postgres.Open(fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", dbHost, dbPort, dbUser, dbName, dbPass, dbSslMode))
		db, err = gorm.Open(dsn)
	}
	if err != nil {
		panic(err)
	}

	autoMigration()
}

func GetDB() *gorm.DB {
	return db
}

/*func Close() {
	db.Close
	if err := db.Close(); err != nil {
		panic(err)
	}
}*/

func autoMigration() {
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.PublicKey{})
}
