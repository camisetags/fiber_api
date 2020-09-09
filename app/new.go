package app

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func pingHandler(ctx *fiber.Ctx) {
	ctx.JSON(fiber.Map{
		"ping": "pong",
	})
}

// New returns a new instance of app
func New() *fiber.App {
	app := fiber.New(&fiber.Settings{
		Prefork:       true,
		StrictRouting: false,
		CaseSensitive: true,
	})

	app.Use(middleware.Recover())

	app.Use(func(ctx *fiber.Ctx) {
		ctx.Accepts("application/json")
		ctx.Next()
	})

	app.Get("/ping", pingHandler)

	return app
}
