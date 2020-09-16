package main

import (
	"fiber_api/app"
	"fiber_api/database"
)

func main() {
	app := app.New()
	db := database.New()
	Router(app, db)

	app.Listen("0.0.0.0:3333")
}
