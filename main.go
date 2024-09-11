package main

import (
	"log"

	"bitbucket.org/isbtotogroup/sdsb4d-backend/db"
	"bitbucket.org/isbtotogroup/sdsb4d-backend/routers"
)

func main() {
	db.Init()
	app := routers.Init()
	log.Fatal(app.Listen(":9091"))
}
