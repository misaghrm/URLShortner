package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/misaghrm/urlshortener/application/handler"
)

func SetupRoute(app *fiber.App) {
	app.Use(handler.Redirect)
}
