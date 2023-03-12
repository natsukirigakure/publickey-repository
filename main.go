package main

import (
	"github.com/natsukirigakure/publickey-repogitory/controller"
	"github.com/natsukirigakure/publickey-repogitory/db"
)

func main() {
	r := controller.Router()

	db.Init()
	r.Run()
}
