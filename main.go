package main

import (
	"publickey-repogitory/controller"
	"publickey-repogitory/db"
)

func main() {
	r := controller.Router()

	db.Init()
	r.Run()

	db.Close()
}
