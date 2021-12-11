package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/misaghrm/urlshortener/application/handler"
)

func SetupRoute(app *fiber.App) {
	api := app.Group("/")
	api.Get("/", handler.Redirect)
	api.Post("/create", handler.NewURL)
}
