package route

import "github.com/gofiber/fiber/v2"

func SetupRoute(app *fiber.App) {
	api := app.Group("/")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("It Works!")
	})
	api.Post("/create")
}
