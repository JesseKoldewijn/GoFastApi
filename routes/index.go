package routes

import "github.com/gofiber/fiber/v3"

func Root(App *fiber.App) {
	App.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello Mom!")
	})
}
