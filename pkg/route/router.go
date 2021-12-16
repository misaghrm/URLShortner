package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/misaghrm/urlshortener/application/handler"
)

func SetupRoute(app *fiber.App) {
	r := app.Group("/")
	r.Get("/*", handler.Redirect)
	url := r.Group("/create")
	url.Post("/new", handler.NewURL)
}
