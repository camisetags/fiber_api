package main

import (
	"fiber_api/transaction"
	"fiber_api/user"
	
	"gorm.io/gorm"
	fiber "github.com/gofiber/fiber/v2"
)

func pingHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"ping": "pong",
	})
}

// Router routes
func Router(app *fiber.App, db *gorm.DB) {
	app.Get("/ping", pingHandler)

	transaction.Routes(app.Group("transactions"), db)
	user.Routes(app.Group("users"), db)
}
