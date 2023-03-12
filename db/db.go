package db

import (
	"github.com/natsukirigakure/publickey-repository/model/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	dsn := postgres.Open("host=0.0.0.0 port=5432 user=gorm dbname=gorm password=gorm sslmode=disable")
	db, err = gorm.Open(dsn)
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
