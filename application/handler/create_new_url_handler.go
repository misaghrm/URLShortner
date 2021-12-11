package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/misaghrm/urlshortener/application/model"
	"github.com/misaghrm/urlshortener/infrastructure/repository"
	"github.com/misaghrm/urlshortener/pkg/util"
	"net/http"
	"net/url"
)

func NewURL(c *fiber.Ctx) error {
	request := model.CreateURLRequest{}
	//response := model.ResponseModel{}
	err := c.BodyParser(&request)
	if err != nil {
		return c.JSON(model.ResponseModel{
			Code:    http.StatusBadRequest,
			Message: "request body is wrong",
		})
	}
	if !util.IsURLValid(request.URL) {
		return c.JSON(model.ResponseModel{
			Code:    http.StatusBadRequest,
			Message: "url is wrong",
		})
	}
	parsedUrl, _ := url.Parse(request.URL)
	marshal, err := json.Marshal(parsedUrl)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(marshal))
	id := repository.InsertNewURL(parsedUrl)
	uniquestring := util.GetStringValue(id)
	e := `"` + uniquestring + `"`
	c.Response().Header.Set("Etag", e)
	c.Response().Header.Set("Cache-Control", "max-age=2592000")
	return c.JSON(model.ResponseModel{
		Data: map[string]string{"URL": uniquestring},
		Code: http.StatusOK,
	})
}
