package main

import (
	"fiber_api/transaction"
	"fiber_api/user"
	"fiber_api/app"
	"fiber_api/database"
)

func main() {
	app := app.New()
	db := database.New()

	transaction.Routes(app.Group("transactions"), db)
	user.Routes(app.Group("users"), db)

	app.Listen(3333)
}
