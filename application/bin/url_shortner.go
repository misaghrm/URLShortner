package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/misaghrm/urlshortner/pkg/route"
	"log"
	"os"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: false,
		AppName: "url Shortener:" + os.Getenv("ENV"),
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,Head,Put,DELETE,PATH",
		AllowHeaders: "",
	}))
	app.Use(recover.New())
	route.SetupRoute(app)
	err := app.Listen(":3000")
	if err != nil {
		log.Panicln(err)
	}
}
