package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/misaghrm/urlshortener/infrastructure/repository"
	"github.com/misaghrm/urlshortener/pkg/util"
)

func Redirect(c *fiber.Ctx) error {
	uniqueStr := c.Path()
	if len(uniqueStr) < 3 {
		c.Redirect("https://www.google.com")
	}
	Id := util.GetIntegerValue(uniqueStr)
	return c.Redirect(repository.Find(Id))
}
