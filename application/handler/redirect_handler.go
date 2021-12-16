package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/misaghrm/urlshortener/infrastructure/repository"
	"github.com/misaghrm/urlshortener/pkg/util"
	"log"
	"strings"
)

func Redirect(c *fiber.Ctx) error {
	if h := c.Get("Etag"); len(h) > 1 {
		log.Println("Etag value", h)
		return c.Redirect(strings.TrimLeft(c.Path(), "/"))
	}
	fmt.Println("raw Path:", c.Path())
	uniqueStr := strings.TrimLeft(c.Path(), "/")
	if len(uniqueStr) < 10 {
		c.Redirect("https://www.google.com")
	}
	log.Println(uniqueStr)
	Id := util.GetIntegerValue(uniqueStr)
	log.Println(Id)
	//e := `"` + uniqueStr + `"`
	url := repository.Find(Id)
	fmt.Println("url:", url)
	c.Response().Header.Set("Etag", url)
	c.Response().Header.Set("Cache-Control", "max-age=2592000")
	return c.Redirect(url)
}
