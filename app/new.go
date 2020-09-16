package app

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// New returns a new instance of app
func New() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		StrictRouting: false,
		CaseSensitive: true,
	})

	app.Use(recover.New())

	app.Use(func(ctx *fiber.Ctx) error {
		ctx.Accepts("application/json")
		return ctx.Next()
	})

	return app
}
